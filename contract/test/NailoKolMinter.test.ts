import { expect } from "chai";
import { ethers } from "hardhat";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { NailoKOL, NailoKolMinter, MockPancakeRouter } from "../typechain-types";

const ZERO_ADDR = ethers.ZeroAddress;
const MONTH = 30 * 24 * 60 * 60;
const PRICE = ethers.parseEther("10");

describe("NailoKolMinter", function () {
  let kol: NailoKOL;
  let usdt: NailoKOL;
  let router: MockPancakeRouter;
  let minter: NailoKolMinter;
  let paymentToken: NailoKOL;

  let kolAddr: string;
  let usdtAddr: string;
  let routerAddr: string;
  let paymentTokenAddr: string;

  let admin: HardhatEthersSigner;
  let treasury: HardhatEthersSigner;
  let user: HardhatEthersSigner;
  let projectParty: HardhatEthersSigner;
  let other: HardhatEthersSigner;

  async function deployERC20(name: string, symbol: string): Promise<NailoKOL> {
    const f = await ethers.getContractFactory("Nailo");
    const c = await f.deploy();
    await c.waitForDeployment();
    return c as unknown as NailoKOL;
  }

  async function getBlockTimestamp(): Promise<number> {
    const block = await ethers.provider.getBlock("latest");
    return block!.timestamp;
  }

  async function advanceTime(seconds: number): Promise<void> {
    await ethers.provider.send("evm_increaseTime", [seconds]);
    await ethers.provider.send("evm_mine", []);
  }

  beforeEach(async function () {
    [admin, treasury, user, projectParty, other] = await ethers.getSigners();

    usdt = await deployERC20("USDT", "USDT");
    paymentToken = await deployERC20("PaymentToken", "PT");
    usdtAddr = await usdt.getAddress();
    paymentTokenAddr = await paymentToken.getAddress();

    const KolFactory = await ethers.getContractFactory("NailoKOL");
    kol = await KolFactory.deploy();
    await kol.waitForDeployment();
    kolAddr = await kol.getAddress();

    const RouterFactory = await ethers.getContractFactory("MockPancakeRouter");
    router = await RouterFactory.deploy(usdtAddr);
    await router.waitForDeployment();
    routerAddr = await router.getAddress();

    const MinterFactory = await ethers.getContractFactory("NailoKolMinter");
    minter = await MinterFactory.deploy(kolAddr, usdtAddr, routerAddr, treasury.address);
    await minter.waitForDeployment();

    await kol.grantRole(await kol.ADMIN_ROLE(), await minter.getAddress());

    const amount = ethers.parseEther("1000000");
    await usdt.transfer(user.address, amount);
    await usdt.transfer(admin.address, amount);
    await paymentToken.transfer(user.address, amount);
    await paymentToken.transfer(admin.address, amount);

    const minterAddr = await minter.getAddress();
    await usdt.connect(user).approve(minterAddr, amount);
    await usdt.connect(admin).approve(minterAddr, amount);
    await paymentToken.connect(user).approve(minterAddr, amount);
    await paymentToken.connect(admin).approve(minterAddr, amount);
    await usdt.approve(routerAddr, amount);
    await router.fund(amount);
  });

  describe("Deployment", function () {
    it("should set immutable variables correctly", async function () {
      expect(await minter.kolToken()).to.equal(kolAddr);
      expect(await minter.usdt()).to.equal(usdtAddr);
      expect(await minter.pancakeRouter()).to.equal(routerAddr);
    });

    it("should set treasury from constructor", async function () {
      expect(await minter.treasury()).to.equal(treasury.address);
    });

    it("should set default monthlyPrice to 10 USDT (18 decimals)", async function () {
      expect(await minter.monthlyPrice()).to.equal(PRICE);
    });

    it("should set default slippageBps to 500 (5%)", async function () {
      expect(await minter.slippageBps()).to.equal(500);
    });

    it("should grant DEFAULT_ADMIN_ROLE and ADMIN_ROLE to deployer", async function () {
      expect(await minter.hasRole(await minter.DEFAULT_ADMIN_ROLE(), admin.address)).to.be.true;
      expect(await minter.hasRole(await minter.ADMIN_ROLE(), admin.address)).to.be.true;
    });

    it("should revert if kolToken is zero address", async function () {
      const f = await ethers.getContractFactory("NailoKolMinter");
      await expect(f.deploy(ZERO_ADDR, usdtAddr, routerAddr, treasury.address)).to.be.revertedWith(
        "Minter: invalid kol token"
      );
    });

    it("should revert if usdt is zero address", async function () {
      const f = await ethers.getContractFactory("NailoKolMinter");
      await expect(f.deploy(kolAddr, ZERO_ADDR, routerAddr, treasury.address)).to.be.revertedWith(
        "Minter: invalid USDT address"
      );
    });

    it("should revert if router is zero address", async function () {
      const f = await ethers.getContractFactory("NailoKolMinter");
      await expect(f.deploy(kolAddr, usdtAddr, ZERO_ADDR, treasury.address)).to.be.revertedWith(
        "Minter: invalid router address"
      );
    });

    it("should revert if treasury is zero address", async function () {
      const f = await ethers.getContractFactory("NailoKolMinter");
      await expect(f.deploy(kolAddr, usdtAddr, routerAddr, ZERO_ADDR)).to.be.revertedWith(
        "Minter: invalid treasury"
      );
    });
  });

  describe("buy()", function () {
    it("should mint a new KOL token on first purchase", async function () {
      await expect(minter.connect(user).buy(1)).to.emit(minter, "Minted");
    });

    it("should set correct expire time", async function () {
      const tx = await minter.connect(user).buy(2);
      const receipt = await tx.wait();
      const block = await ethers.provider.getBlock(receipt!.blockNumber);
      const expectedExpire = BigInt(block!.timestamp) + BigInt(2 * MONTH);
      expect(await kol.expireTime(user.address)).to.equal(expectedExpire);
    });

    it("should transfer correct USDT amount to treasury", async function () {
      const balanceBefore = await usdt.balanceOf(treasury.address);
      await minter.connect(user).buy(3);
      const balanceAfter = await usdt.balanceOf(treasury.address);
      expect(balanceAfter - balanceBefore).to.equal(PRICE * 3n);
    });

    it("should emit TokenPurchased event", async function () {
      await expect(minter.connect(user).buy(1))
        .to.emit(minter, "TokenPurchased")
        .withArgs(user.address, usdtAddr, PRICE, 1, PRICE);
    });

    it("should renew existing token on second purchase", async function () {
      await minter.connect(user).buy(1);
      await advanceTime(5 * 86400);
      await expect(minter.connect(user).buy(1)).to.emit(minter, "Renewed");
    });

    it("should extend from current expiry when not expired", async function () {
      await minter.connect(user).buy(1);
      const firstExpire = await kol.expireTime(user.address);
      await advanceTime(5 * 86400);
      await minter.connect(user).buy(1);
      expect(await kol.expireTime(user.address)).to.equal(firstExpire + BigInt(1 * MONTH));
    });

    it("should extend from block.timestamp when expired", async function () {
      await minter.connect(user).buy(1);
      const firstExpire = await kol.expireTime(user.address);
      const now = await getBlockTimestamp();
      await advanceTime(Number(firstExpire) - now + 3600);

      const tx = await minter.connect(user).buy(1);
      const receipt = await tx.wait();
      const block = await ethers.provider.getBlock(receipt!.blockNumber);
      expect(await kol.expireTime(user.address)).to.equal(
        BigInt(block!.timestamp) + BigInt(1 * MONTH)
      );
    });

    it("should emit Minted with correct months", async function () {
      await expect(minter.connect(user).buy(3))
        .to.emit(minter, "Minted")
        .withArgs(user.address, 0, 3);
    });

    it("should emit Renewed with correct months", async function () {
      await minter.connect(user).buy(1);
      await advanceTime(1 * 86400);
      await expect(minter.connect(user).buy(2))
        .to.emit(minter, "Renewed")
        .withArgs(user.address, 2);
    });

    it("should revert if months == 0", async function () {
      await expect(minter.connect(user).buy(0)).to.be.revertedWith(
        "Minter: months must be greater than 0"
      );
    });

    it("should revert if user has insufficient USDT allowance", async function () {
      await usdt.connect(other).approve(await minter.getAddress(), ethers.parseEther("1"));
      await expect(minter.connect(other).buy(10)).to.be.reverted;
    });

    it("should handle multiple months proportionally", async function () {
      const balanceBefore = await usdt.balanceOf(treasury.address);
      await minter.connect(user).buy(6);
      const balanceAfter = await usdt.balanceOf(treasury.address);
      expect(balanceAfter - balanceBefore).to.equal(PRICE * 6n);
    });

    it("should allow admin to buy for themselves", async function () {
      await expect(minter.connect(admin).buy(1)).to.emit(minter, "Minted");
    });
  });

  describe("buyWithToken()", function () {
    let swapPath: [string, string];

    beforeEach(async function () {
      swapPath = [paymentTokenAddr, usdtAddr];
      await minter.connect(admin).addPaymentToken(paymentTokenAddr, projectParty.address, swapPath);
      await paymentToken
        .connect(user)
        .approve(await minter.getAddress(), ethers.parseEther("1000000"));
    });

    it("should mint a KOL token on first purchase with payment token", async function () {
      await expect(minter.connect(user).buyWithToken(paymentTokenAddr, 1)).to.emit(
        minter,
        "Minted"
      );
    });

    it("should set correct expire time", async function () {
      const tx = await minter.connect(user).buyWithToken(paymentTokenAddr, 1);
      const receipt = await tx.wait();
      const block = await ethers.provider.getBlock(receipt!.blockNumber);
      const expected = BigInt(block!.timestamp) + BigInt(1 * MONTH);
      expect(await kol.expireTime(user.address)).to.equal(expected);
    });

    it("should send 50% of payment tokens to projectParty", async function () {
      const balanceBefore = await paymentToken.balanceOf(projectParty.address);
      await minter.connect(user).buyWithToken(paymentTokenAddr, 1);
      const balanceAfter = await paymentToken.balanceOf(projectParty.address);
      expect(balanceAfter - balanceBefore).to.be.gt(0);

      const withBuffer = (PRICE * 10500n) / 10000n;
      const expectedHalf = withBuffer / 2n;
      expect(balanceAfter - balanceBefore).to.equal(expectedHalf);
    });

    it("should swap remaining tokens to USDT and send to treasury", async function () {
      const treasuryBalBefore = await usdt.balanceOf(treasury.address);
      await minter.connect(user).buyWithToken(paymentTokenAddr, 1);
      const treasuryBalAfter = await usdt.balanceOf(treasury.address);
      expect(treasuryBalAfter - treasuryBalBefore).to.be.gt(0);
    });

    it("should emit TokenPurchased event", async function () {
      await expect(minter.connect(user).buyWithToken(paymentTokenAddr, 2)).to.emit(
        minter,
        "TokenPurchased"
      );
    });

    it("should renew existing token on second purchase", async function () {
      await minter.connect(user).buyWithToken(paymentTokenAddr, 1);
      await advanceTime(5 * 86400);
      await expect(minter.connect(user).buyWithToken(paymentTokenAddr, 1)).to.emit(
        minter,
        "Renewed"
      );
    });

    it("should extend from current expiry when not expired", async function () {
      await minter.connect(user).buyWithToken(paymentTokenAddr, 1);
      const firstExpire = await kol.expireTime(user.address);
      await advanceTime(5 * 86400);
      await minter.connect(user).buyWithToken(paymentTokenAddr, 1);
      expect(await kol.expireTime(user.address)).to.equal(firstExpire + BigInt(1 * MONTH));
    });

    it("should extend from block.timestamp when expired", async function () {
      await minter.connect(user).buyWithToken(paymentTokenAddr, 1);
      const firstExpire = await kol.expireTime(user.address);
      const now = await getBlockTimestamp();
      await advanceTime(Number(firstExpire) - now + 3600);

      const tx = await minter.connect(user).buyWithToken(paymentTokenAddr, 1);
      const receipt = await tx.wait();
      const block = await ethers.provider.getBlock(receipt!.blockNumber);
      expect(await kol.expireTime(user.address)).to.equal(
        BigInt(block!.timestamp) + BigInt(1 * MONTH)
      );
    });

    it("should emit Minted with correct months", async function () {
      await expect(minter.connect(user).buyWithToken(paymentTokenAddr, 3))
        .to.emit(minter, "Minted")
        .withArgs(user.address, 0, 3);
    });

    it("should emit Renewed with correct months", async function () {
      await minter.connect(user).buyWithToken(paymentTokenAddr, 1);
      await advanceTime(1 * 86400);
      await expect(minter.connect(user).buyWithToken(paymentTokenAddr, 2))
        .to.emit(minter, "Renewed")
        .withArgs(user.address, 2);
    });

    it("should revert if token is not configured", async function () {
      await expect(minter.connect(user).buyWithToken(other.address, 1)).to.be.revertedWith(
        "Minter: token not configured"
      );
    });

    it("should revert if months == 0", async function () {
      await expect(minter.connect(user).buyWithToken(paymentTokenAddr, 0)).to.be.revertedWith(
        "Minter: months must be greater than 0"
      );
    });

    it("should revert if token was removed", async function () {
      await minter.connect(admin).removePaymentToken(paymentTokenAddr);
      await expect(minter.connect(user).buyWithToken(paymentTokenAddr, 1)).to.be.revertedWith(
        "Minter: token not configured"
      );
    });

    it("should handle odd amounts correctly", async function () {
      await router.setRate(3, 2);
      await expect(minter.connect(user).buyWithToken(paymentTokenAddr, 1)).to.not.be.reverted;
    });

    it("should revert if slippage check fails", async function () {
      await router.setSwapPenaltyBps(1000);
      await expect(
        minter.connect(user).buyWithToken(paymentTokenAddr, 1)
      ).to.be.revertedWith("INSUFFICIENT_OUTPUT_AMOUNT");
    });
  });

  describe("Admin", function () {
    describe("setMonthlyPrice", function () {
      it("should update monthlyPrice", async function () {
        const newPrice = ethers.parseEther("20");
        await minter.connect(admin).setMonthlyPrice(newPrice);
        expect(await minter.monthlyPrice()).to.equal(newPrice);
      });

      it("should emit MonthlyPriceUpdated event", async function () {
        const newPrice = ethers.parseEther("20");
        await expect(minter.connect(admin).setMonthlyPrice(newPrice))
          .to.emit(minter, "MonthlyPriceUpdated")
          .withArgs(PRICE, newPrice);
      });

      it("should revert if not admin", async function () {
        await expect(
          minter.connect(user).setMonthlyPrice(ethers.parseEther("20"))
        ).to.be.reverted;
      });
    });

    describe("setTreasury", function () {
      it("should update treasury", async function () {
        await minter.connect(admin).setTreasury(other.address);
        expect(await minter.treasury()).to.equal(other.address);
      });

      it("should emit TreasuryUpdated event", async function () {
        await expect(minter.connect(admin).setTreasury(other.address))
          .to.emit(minter, "TreasuryUpdated")
          .withArgs(treasury.address, other.address);
      });

      it("should revert if zero address", async function () {
        await expect(minter.connect(admin).setTreasury(ZERO_ADDR)).to.be.revertedWith(
          "Minter: invalid treasury"
        );
      });

      it("should revert if not admin", async function () {
        await expect(minter.connect(user).setTreasury(other.address)).to.be.reverted;
      });

      it("should route future USDT payments to new treasury", async function () {
        await minter.connect(admin).setTreasury(other.address);
        const balBefore = await usdt.balanceOf(other.address);
        await minter.connect(user).buy(1);
        const balAfter = await usdt.balanceOf(other.address);
        expect(balAfter - balBefore).to.equal(PRICE);
      });
    });

    describe("setSlippageBps", function () {
      it("should update slippageBps", async function () {
        await minter.connect(admin).setSlippageBps(300);
        expect(await minter.slippageBps()).to.equal(300);
      });

      it("should emit SlippageUpdated event", async function () {
        await expect(minter.connect(admin).setSlippageBps(300))
          .to.emit(minter, "SlippageUpdated")
          .withArgs(500, 300);
      });

      it("should revert if > 5000 (50%)", async function () {
        await expect(minter.connect(admin).setSlippageBps(5001)).to.be.revertedWith(
          "Minter: slippage too high"
        );
      });

      it("should allow 5000 (50%)", async function () {
        await expect(minter.connect(admin).setSlippageBps(5000)).to.not.be.reverted;
      });

      it("should revert if not admin", async function () {
        await expect(minter.connect(user).setSlippageBps(300)).to.be.reverted;
      });
    });

    describe("addPaymentToken", function () {
      it("should add a payment token", async function () {
        await minter.connect(admin).addPaymentToken(paymentTokenAddr, projectParty.address, [
          paymentTokenAddr,
          usdtAddr,
        ]);
        expect(await minter.isTokenSupported(paymentTokenAddr)).to.be.true;
      });

      it("should emit PaymentTokenAdded event", async function () {
        await expect(
          minter
            .connect(admin)
            .addPaymentToken(paymentTokenAddr, projectParty.address, [paymentTokenAddr, usdtAddr])
        )
          .to.emit(minter, "PaymentTokenAdded")
          .withArgs(paymentTokenAddr, projectParty.address);
      });

      it("should store correct token data", async function () {
        await minter
          .connect(admin)
          .addPaymentToken(paymentTokenAddr, projectParty.address, [paymentTokenAddr, usdtAddr]);
        const pt = await minter.getPaymentToken(paymentTokenAddr);
        expect(pt.token).to.equal(paymentTokenAddr);
        expect(pt.projectParty).to.equal(projectParty.address);
        expect(pt.active).to.be.true;
      });

      it("should revert if token already added", async function () {
        await minter
          .connect(admin)
          .addPaymentToken(paymentTokenAddr, projectParty.address, [paymentTokenAddr, usdtAddr]);
        await expect(
          minter
            .connect(admin)
            .addPaymentToken(paymentTokenAddr, projectParty.address, [paymentTokenAddr, usdtAddr])
        ).to.be.revertedWith("Minter: token already added");
      });

      it("should revert if token is zero address", async function () {
        await expect(
          minter
            .connect(admin)
            .addPaymentToken(ZERO_ADDR, projectParty.address, [ZERO_ADDR, usdtAddr])
        ).to.be.revertedWith("Minter: invalid token");
      });

      it("should revert if swap path is too short", async function () {
        await expect(
          minter
            .connect(admin)
            .addPaymentToken(paymentTokenAddr, projectParty.address, [paymentTokenAddr])
        ).to.be.revertedWith("Minter: swap path too short");
      });

      it("should revert if swap path does not end with USDT", async function () {
        await expect(
          minter
            .connect(admin)
            .addPaymentToken(paymentTokenAddr, projectParty.address, [
              paymentTokenAddr,
              other.address,
            ])
        ).to.be.revertedWith("Minter: path must end with USDT");
      });

      it("should revert if not admin", async function () {
        await expect(
          minter
            .connect(user)
            .addPaymentToken(paymentTokenAddr, projectParty.address, [paymentTokenAddr, usdtAddr])
        ).to.be.reverted;
      });

      it("should allow re-adding a previously removed token", async function () {
        await minter
          .connect(admin)
          .addPaymentToken(paymentTokenAddr, projectParty.address, [paymentTokenAddr, usdtAddr]);
        await minter.connect(admin).removePaymentToken(paymentTokenAddr);
        await expect(
          minter
            .connect(admin)
            .addPaymentToken(paymentTokenAddr, projectParty.address, [paymentTokenAddr, usdtAddr])
        ).to.not.be.reverted;
        expect(await minter.isTokenSupported(paymentTokenAddr)).to.be.true;
      });
    });

    describe("updatePaymentToken", function () {
      beforeEach(async function () {
        await minter
          .connect(admin)
          .addPaymentToken(paymentTokenAddr, projectParty.address, [paymentTokenAddr, usdtAddr]);
      });

      it("should update projectParty", async function () {
        await minter
          .connect(admin)
          .updatePaymentToken(paymentTokenAddr, other.address, [paymentTokenAddr, usdtAddr]);
        const pt = await minter.getPaymentToken(paymentTokenAddr);
        expect(pt.projectParty).to.equal(other.address);
      });

      it("should emit PaymentTokenUpdated event", async function () {
        await expect(
          minter
            .connect(admin)
            .updatePaymentToken(paymentTokenAddr, other.address, [paymentTokenAddr, usdtAddr])
        )
          .to.emit(minter, "PaymentTokenUpdated")
          .withArgs(paymentTokenAddr, other.address);
      });

      it("should revert if token not found", async function () {
        await expect(
          minter
            .connect(admin)
            .updatePaymentToken(other.address, projectParty.address, [other.address, usdtAddr])
        ).to.be.revertedWith("Minter: token not found");
      });

      it("should revert if swap path too short", async function () {
        await expect(
          minter
            .connect(admin)
            .updatePaymentToken(paymentTokenAddr, projectParty.address, [paymentTokenAddr])
        ).to.be.revertedWith("Minter: swap path too short");
      });

      it("should revert if swap path does not end with USDT", async function () {
        await expect(
          minter
            .connect(admin)
            .updatePaymentToken(paymentTokenAddr, projectParty.address, [
              paymentTokenAddr,
              other.address,
            ])
        ).to.be.revertedWith("Minter: path must end with USDT");
      });

      it("should revert if not admin", async function () {
        await expect(
          minter
            .connect(user)
            .updatePaymentToken(paymentTokenAddr, other.address, [paymentTokenAddr, usdtAddr])
        ).to.be.reverted;
      });
    });

    describe("removePaymentToken", function () {
      beforeEach(async function () {
        await minter
          .connect(admin)
          .addPaymentToken(paymentTokenAddr, projectParty.address, [paymentTokenAddr, usdtAddr]);
      });

      it("should deactivate the token", async function () {
        await minter.connect(admin).removePaymentToken(paymentTokenAddr);
        expect(await minter.isTokenSupported(paymentTokenAddr)).to.be.false;
      });

      it("should emit PaymentTokenRemoved event", async function () {
        await expect(minter.connect(admin).removePaymentToken(paymentTokenAddr))
          .to.emit(minter, "PaymentTokenRemoved")
          .withArgs(paymentTokenAddr);
      });

      it("should revert if token not found", async function () {
        await expect(
          minter.connect(admin).removePaymentToken(other.address)
        ).to.be.revertedWith("Minter: token not found");
      });

      it("should revert if not admin", async function () {
        await expect(minter.connect(user).removePaymentToken(paymentTokenAddr)).to.be.reverted;
      });

      it("should make getPaymentToken revert for removed token", async function () {
        await minter.connect(admin).removePaymentToken(paymentTokenAddr);
        await expect(minter.getPaymentToken(paymentTokenAddr)).to.be.revertedWith(
          "Minter: token not found"
        );
      });
    });
  });

  describe("Views", function () {
    beforeEach(async function () {
      await minter
        .connect(admin)
        .addPaymentToken(paymentTokenAddr, projectParty.address, [paymentTokenAddr, usdtAddr]);
    });

    describe("getPaymentToken", function () {
      it("should return correct payment token data", async function () {
        const pt = await minter.getPaymentToken(paymentTokenAddr);
        expect(pt.token).to.equal(paymentTokenAddr);
        expect(pt.projectParty).to.equal(projectParty.address);
        expect(pt.active).to.be.true;
        expect(pt.swapPath).to.deep.equal([paymentTokenAddr, usdtAddr]);
      });

      it("should revert for inactive token", async function () {
        await minter.connect(admin).removePaymentToken(paymentTokenAddr);
        await expect(minter.getPaymentToken(paymentTokenAddr)).to.be.revertedWith(
          "Minter: token not found"
        );
      });

      it("should revert for never-added token", async function () {
        await expect(minter.getPaymentToken(other.address)).to.be.revertedWith(
          "Minter: token not found"
        );
      });
    });

    describe("isTokenSupported", function () {
      it("should return true for active token", async function () {
        expect(await minter.isTokenSupported(paymentTokenAddr)).to.be.true;
      });

      it("should return false for never-added token", async function () {
        expect(await minter.isTokenSupported(other.address)).to.be.false;
      });

      it("should return false for removed token", async function () {
        await minter.connect(admin).removePaymentToken(paymentTokenAddr);
        expect(await minter.isTokenSupported(paymentTokenAddr)).to.be.false;
      });
    });

    describe("getRequiredTokens", function () {
      it("should return token amount with buffer for given months", async function () {
        const required = await minter.getRequiredTokens(paymentTokenAddr, 1);
        expect(required).to.equal((PRICE * 10500n) / 10000n);
      });

      it("should scale with months", async function () {
        const r1 = await minter.getRequiredTokens(paymentTokenAddr, 1);
        const r3 = await minter.getRequiredTokens(paymentTokenAddr, 3);
        expect(r3).to.equal(r1 * 3n);
      });

      it("should revert if token not configured", async function () {
        await expect(minter.getRequiredTokens(other.address, 1)).to.be.revertedWith(
          "Minter: token not configured"
        );
      });

      it("should revert if months == 0", async function () {
        await expect(minter.getRequiredTokens(paymentTokenAddr, 0)).to.be.revertedWith(
          "Minter: months must be greater than 0"
        );
      });

      it("should respect rate changes from router", async function () {
        await router.setRate(2, 1);
        const required = await minter.getRequiredTokens(paymentTokenAddr, 1);
        expect(required).to.equal(((PRICE * 1n) / 2n) * 10500n / 10000n);
      });
    });

    describe("estimateSwap", function () {
      it("should return USDT output for given token amount", async function () {
        const amount = ethers.parseEther("100");
        const estimated = await minter.estimateSwap(paymentTokenAddr, amount);
        expect(estimated).to.equal(amount);
      });

      it("should revert for inactive token", async function () {
        await minter.connect(admin).removePaymentToken(paymentTokenAddr);
        await expect(
          minter.estimateSwap(paymentTokenAddr, ethers.parseEther("10"))
        ).to.be.revertedWith("Minter: token not configured");
      });
    });
  });

  describe("Integration & Edge Cases", function () {
    beforeEach(async function () {
      await minter
        .connect(admin)
        .addPaymentToken(paymentTokenAddr, projectParty.address, [paymentTokenAddr, usdtAddr]);
    });

    it("should allow mixing USDT and token payments", async function () {
      await minter.connect(user).buy(1);
      await advanceTime(10 * 86400);
      await minter.connect(user).buyWithToken(paymentTokenAddr, 1);
      expect(await kol.balanceOf(user.address)).to.equal(1);
    });

    it("should mint sequential tokenIds for different users", async function () {
      await minter.connect(user).buy(1);
      await minter.connect(admin).buy(1);

      const userEvents = await minter.queryFilter(minter.filters.Minted(user.address));
      const adminEvents = await minter.queryFilter(minter.filters.Minted(admin.address));
      expect(userEvents[0].args.tokenId).to.equal(0);
      expect(adminEvents[0].args.tokenId).to.equal(1);
    });

    it("should correctly handle renew after mix of payment methods", async function () {
      await minter.connect(user).buy(1);
      const firstExpire = await kol.expireTime(user.address);
      await advanceTime(10 * 86400);
      await minter.connect(user).buyWithToken(paymentTokenAddr, 1);
      expect(await kol.expireTime(user.address)).to.equal(firstExpire + BigInt(1 * MONTH));
    });

    it("should handle large month value without overflow", async function () {
      await expect(minter.connect(user).buy(120)).to.not.be.reverted;
    });

    it("should emit Minted with correct user and tokenId", async function () {
      await expect(minter.connect(user).buy(1))
        .to.emit(minter, "Minted")
        .withArgs(user.address, 0, 1);
    });

    it("should allow burn via KOL contract", async function () {
      await minter.connect(user).buy(1);
      const mintedEvents = await minter.queryFilter(minter.filters.Minted(user.address));
      const tokenId = mintedEvents[0].args.tokenId;
      await kol.connect(user).burn(tokenId);
      expect(await kol.balanceOf(user.address)).to.equal(0);
    });

    it("should allow minting again after burn", async function () {
      await minter.connect(user).buy(1);
      const mintedEvents = await minter.queryFilter(minter.filters.Minted(user.address));
      const tokenId = mintedEvents[0].args.tokenId;
      await kol.connect(user).burn(tokenId);
      await expect(minter.connect(user).buy(1)).to.emit(minter, "Minted");
    });

    it("should not allow two tokens for same address via direct mint bypass", async function () {
      await minter.connect(user).buy(1);
      await expect(kol.connect(admin).mint(user.address)).to.be.revertedWith(
        "NailoKOL: already has a token"
      );
    });
  });
});
