import { expect } from "chai";
import { ethers } from "hardhat";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { NailoNode, NailoDividend } from "../typechain-types";

const ZERO_ADDR = ethers.ZeroAddress;
const MAX_SUPPLY = 200;

describe("NailoDividend", function () {
  let node: NailoNode;
  let dividend: NailoDividend;
  let admin: HardhatEthersSigner;
  let user: HardhatEthersSigner;
  let user2: HardhatEthersSigner;

  beforeEach(async function () {
    [admin, user, user2] = await ethers.getSigners();

    const NodeFactory = await ethers.getContractFactory("NailoNode");
    node = await NodeFactory.deploy();
    await node.waitForDeployment();

    const DivFactory = await ethers.getContractFactory("NailoDividend");
    dividend = await DivFactory.deploy(await node.getAddress());
    await dividend.waitForDeployment();

    await node.grantRole(await node.ADMIN_ROLE(), admin.address);
  });

  describe("Deployment", function () {
    it("should set nodeToken immutable", async function () {
      expect(await dividend.nodeToken()).to.equal(await node.getAddress());
    });

    it("should grant DEFAULT_ADMIN_ROLE and ADMIN_ROLE to deployer", async function () {
      expect(await dividend.hasRole(await dividend.DEFAULT_ADMIN_ROLE(), admin.address)).to.be.true;
      expect(await dividend.hasRole(await dividend.ADMIN_ROLE(), admin.address)).to.be.true;
    });

    it("should start with accDividendPerToken = 0", async function () {
      expect(await dividend.accDividendPerToken()).to.equal(0);
    });

    it("should revert if node token is zero address", async function () {
      const F = await ethers.getContractFactory("NailoDividend");
      await expect(F.deploy(ZERO_ADDR)).to.be.revertedWith("Dividend: invalid node token");
    });
  });

  describe("deposit()", function () {
    it("should accept BNB deposit", async function () {
      const amount = ethers.parseEther("1");
      await expect(dividend.connect(user).deposit({ value: amount }))
        .to.emit(dividend, "Deposited")
        .withArgs(user.address, amount, 0);
    });

    it("should revert zero deposit", async function () {
      await expect(dividend.connect(user).deposit({ value: 0 })).to.be.revertedWith(
        "Dividend: zero deposit"
      );
    });

    it("should update accDividendPerToken when supply > 0", async function () {
      await node.connect(admin).mint(user.address);
      const amount = ethers.parseEther("1");
      await dividend.connect(user).deposit({ value: amount });

      const expected = (amount * BigInt(1e12)) / 1n;
      expect(await dividend.accDividendPerToken()).to.equal(expected);
    });

    it("should not update accDividendPerToken when supply = 0", async function () {
      const amount = ethers.parseEther("1");
      await dividend.connect(user).deposit({ value: amount });
      expect(await dividend.accDividendPerToken()).to.equal(0);
    });

    it("should work via receive() fallback", async function () {
      const amount = ethers.parseEther("0.5");
      await user.sendTransaction({
        to: await dividend.getAddress(),
        value: amount,
      });
      expect(await ethers.provider.getBalance(await dividend.getAddress())).to.equal(amount);
    });
  });

  describe("claim()", function () {
    beforeEach(async function () {
      await node.connect(admin).mint(user.address);
      await node.connect(admin).mint(user.address);
      await node.connect(admin).mint(user2.address);
    });

    it("should claim dividends for all owned tokens", async function () {
      const depositAmount = ethers.parseEther("3");
      await dividend.connect(admin).deposit({ value: depositAmount });

      const balanceBefore = await ethers.provider.getBalance(user.address);
      const tx = await dividend.connect(user).claim();
      const receipt = await tx.wait();
      const gasCost = receipt!.gasUsed * receipt!.gasPrice;
      const balanceAfter = await ethers.provider.getBalance(user.address);

      // user owns 2 of 3 tokens, should get 2/3 of 3 ETH = 2 ETH
      const expected = (depositAmount * 2n) / 3n;
      expect(balanceAfter - balanceBefore + gasCost).to.equal(expected);
    });

    it("should handle multiple deposits correctly", async function () {
      await dividend.connect(admin).deposit({ value: ethers.parseEther("1") });
      await dividend.connect(admin).deposit({ value: ethers.parseEther("2") });

      // user owns 2 of 3 tokens, total deposits = 3 ETH, share = 2 ETH
      const balanceBefore = await ethers.provider.getBalance(user.address);
      const tx = await dividend.connect(user).claim();
      const receipt = await tx.wait();
      const gasCost = receipt!.gasUsed * receipt!.gasPrice;
      const balanceAfter = await ethers.provider.getBalance(user.address);

      const diff = balanceAfter - balanceBefore + gasCost;
      const expected = ethers.parseEther("2");
      expect(diff >= expected - 5n && diff <= expected + 5n).to.be.true;
    });

    it("should emit Claimed event for each token", async function () {
      await dividend.connect(admin).deposit({ value: ethers.parseEther("3") });
      await expect(dividend.connect(user).claim())
        .to.emit(dividend, "Claimed")
        .withArgs(user.address, 0, ethers.parseEther("1"))
        .and.to.emit(dividend, "Claimed")
        .withArgs(user.address, 1, ethers.parseEther("1"));
    });

    it("should revert if nothing to claim", async function () {
      await expect(dividend.connect(user).claim()).to.be.revertedWith(
        "Dividend: nothing to claim"
      );
    });

    it("should not claim twice for the same deposit", async function () {
      await dividend.connect(admin).deposit({ value: ethers.parseEther("3") });
      await dividend.connect(user).claim();
      await expect(dividend.connect(user).claim()).to.be.revertedWith(
        "Dividend: nothing to claim"
      );
    });

    it("should allow claiming after additional deposits", async function () {
      await dividend.connect(admin).deposit({ value: ethers.parseEther("3") });
      await dividend.connect(user).claim();

      await dividend.connect(admin).deposit({ value: ethers.parseEther("3") });
      const balanceBefore = await ethers.provider.getBalance(user.address);
      const tx = await dividend.connect(user).claim();
      const receipt = await tx.wait();
      const gasCost = receipt!.gasUsed * receipt!.gasPrice;
      const balanceAfter = await ethers.provider.getBalance(user.address);

      expect(balanceAfter - balanceBefore + gasCost).to.equal(ethers.parseEther("2"));
    });

    it("should transfer dividends to new owner after token transfer", async function () {
      await dividend.connect(admin).deposit({ value: ethers.parseEther("3") });

      // Transfer token 0 to user2, user2 now has token 0 and token 2 = 2 of 3
      await node.connect(user).transferFrom(user.address, user2.address, 0);

      const balanceBefore = await ethers.provider.getBalance(user2.address);
      const tx = await dividend.connect(user2).claim();
      const receipt = await tx.wait();
      const gasCost = receipt!.gasUsed * receipt!.gasPrice;
      const balanceAfter = await ethers.provider.getBalance(user2.address);

      const diff = balanceAfter - balanceBefore + gasCost;
      const expected = ethers.parseEther("2");
      expect(diff >= expected - 5n && diff <= expected + 5n).to.be.true;
    });
  });

  describe("pendingDividend()", function () {
    beforeEach(async function () {
      await node.connect(admin).mint(user.address);
      await node.connect(admin).mint(user2.address);
    });

    it("should return 0 before any deposit", async function () {
      expect(await dividend.pendingDividend(0)).to.equal(0);
    });

    it("should return correct pending amount after deposit", async function () {
      await dividend.connect(admin).deposit({ value: ethers.parseEther("2") });
      expect(await dividend.pendingDividend(0)).to.equal(ethers.parseEther("1"));
    });

    it("should return 0 for non-existent token", async function () {
      await dividend.connect(admin).deposit({ value: ethers.parseEther("1") });
      expect(await dividend.pendingDividend(999)).to.equal(0);
    });

    it("should reflect unclaimed amount between deposits", async function () {
      await dividend.connect(admin).deposit({ value: ethers.parseEther("2") });
      expect(await dividend.pendingDividend(0)).to.equal(ethers.parseEther("1"));

      await dividend.connect(admin).deposit({ value: ethers.parseEther("2") });
      expect(await dividend.pendingDividend(0)).to.equal(ethers.parseEther("2"));
    });
  });

  describe("getMyTokenIds()", function () {
    it("should return empty for address with no tokens", async function () {
      const ids = await dividend.connect(user).getMyTokenIds();
      expect(ids.length).to.equal(0);
    });

    it("should return all token IDs for owner", async function () {
      await node.connect(admin).mint(user.address);
      await node.connect(admin).mint(user.address);
      await node.connect(admin).mint(user.address);

      const ids = await dividend.connect(user).getMyTokenIds();
      expect(ids.length).to.equal(3);
      expect(ids[0]).to.equal(0);
      expect(ids[1]).to.equal(1);
      expect(ids[2]).to.equal(2);
    });
  });

  describe("pendingDividendBatch()", function () {
    beforeEach(async function () {
      await node.connect(admin).mint(user.address);
      await node.connect(admin).mint(user.address);
      await dividend.connect(admin).deposit({ value: ethers.parseEther("2") });
    });

    it("should return array of pending amounts", async function () {
      const result = await dividend.pendingDividendBatch([0, 1]);
      expect(result[0]).to.equal(ethers.parseEther("1"));
      expect(result[1]).to.equal(ethers.parseEther("1"));
    });

    it("should return 0 for unowned token IDs", async function () {
      const result = await dividend.pendingDividendBatch([0, 999]);
      expect(result[0]).to.equal(ethers.parseEther("1"));
      expect(result[1]).to.equal(0);
    });
  });

  describe("Precision", function () {
    it("should handle small deposits correctly", async function () {
      await node.connect(admin).mint(user.address);
      await node.connect(admin).mint(user2.address);

      // 1 wei deposit with 2 holders => 0 each (truncated)
      await dividend.connect(admin).deposit({ value: 1 });
      expect(await dividend.pendingDividend(0)).to.equal(0);
    });

    it("should accumulate small amounts over multiple deposits", async function () {
      await node.connect(admin).mint(user.address);

      // Deposit 2 wei, 10 times = 20 wei total
      for (let i = 0; i < 10; i++) {
        await dividend.connect(admin).deposit({ value: 2 });
      }
      expect(await dividend.pendingDividend(0)).to.equal(20);
    });

    it("should handle zero-supply edge case gracefully", async function () {
      await dividend.connect(admin).deposit({ value: ethers.parseEther("100") });
      expect(await dividend.accDividendPerToken()).to.equal(0);

      // Mint after deposit, old deposit is not distributed
      await node.connect(admin).mint(user.address);
      expect(await dividend.pendingDividend(0)).to.equal(0);

      // New deposit works normally
      await dividend.connect(admin).deposit({ value: ethers.parseEther("1") });
      expect(await dividend.pendingDividend(0)).to.equal(ethers.parseEther("1"));
    });
  });
});
