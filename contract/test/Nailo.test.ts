import { expect } from "chai";
import { ethers } from "hardhat";
import { Nailo } from "../typechain-types";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";

describe("Nailo Token", function () {
  let nailo: Nailo;
  let owner: HardhatEthersSigner;
  let addr1: HardhatEthersSigner;
  let addr2: HardhatEthersSigner;

  const INITIAL_SUPPLY = ethers.parseEther("1000000000"); // 10亿

  beforeEach(async function () {
    // 获取测试账户
    [owner, addr1, addr2] = await ethers.getSigners();

    // 部署合约
    const NailoFactory = await ethers.getContractFactory("Nailo");
    nailo = await NailoFactory.deploy();
    await nailo.waitForDeployment();
  });

  describe("部署验证", function () {
    it("应该设置正确的代币名称", async function () {
      expect(await nailo.name()).to.equal("Nailo");
    });

    it("应该设置正确的代币符号", async function () {
      expect(await nailo.symbol()).to.equal("NAILO");
    });

    it("应该设置正确的小数位数", async function () {
      expect(await nailo.decimals()).to.equal(18);
    });

    it("应该铸造正确的初始总供应量", async function () {
      expect(await nailo.totalSupply()).to.equal(INITIAL_SUPPLY);
    });

    it("应该将初始供应量分配给部署者", async function () {
      expect(await nailo.balanceOf(owner.address)).to.equal(INITIAL_SUPPLY);
    });
  });

  describe("转账功能", function () {
    it("应该能正常转账", async function () {
      const amount = ethers.parseEther("100");

      await nailo.transfer(addr1.address, amount);

      expect(await nailo.balanceOf(owner.address)).to.equal(INITIAL_SUPPLY - amount);
      expect(await nailo.balanceOf(addr1.address)).to.equal(amount);
    });

    it("应该触发Transfer事件", async function () {
      const amount = ethers.parseEther("100");

      await expect(nailo.transfer(addr1.address, amount))
        .to.emit(nailo, "Transfer")
        .withArgs(owner.address, addr1.address, amount);
    });

    it("转账超过余额应该失败", async function () {
      const balance = await nailo.balanceOf(addr1.address);
      expect(balance).to.equal(0);

      await expect(nailo.connect(addr1).transfer(owner.address, 1))
        .to.be.revertedWithCustomError;
    });

    it("支持转账金额为0", async function () {
      await expect(nailo.transfer(addr1.address, 0))
        .to.emit(nailo, "Transfer")
        .withArgs(owner.address, addr1.address, 0);
    });
  });

  describe("授权与代理转账", function () {
    it("应该能设置授权额度", async function () {
      const amount = ethers.parseEther("100");

      await nailo.approve(addr1.address, amount);

      expect(await nailo.allowance(owner.address, addr1.address)).to.equal(amount);
    });

    it("应该能在授权额度内代理转账", async function () {
      const amount = ethers.parseEther("100");

      await nailo.approve(addr1.address, amount);
      await nailo.connect(addr1).transferFrom(owner.address, addr2.address, amount);

      expect(await nailo.balanceOf(addr2.address)).to.equal(amount);
      expect(await nailo.allowance(owner.address, addr1.address)).to.equal(0);
    });

    it("超出授权额度转账应该失败", async function () {
      const approvedAmount = ethers.parseEther("100");
      const transferAmount = ethers.parseEther("200");

      await nailo.approve(addr1.address, approvedAmount);

      await expect(nailo.connect(addr1).transferFrom(owner.address, addr2.address, transferAmount))
        .to.be.revertedWithCustomError;
    });
  });
});
