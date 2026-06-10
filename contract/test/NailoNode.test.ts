import { expect } from "chai";
import { ethers } from "hardhat";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { NailoNode } from "../typechain-types";

const ZERO_ADDR = ethers.ZeroAddress;
const MAX_SUPPLY = 200;

describe("NailoNode", function () {
  let node: NailoNode;
  let admin: HardhatEthersSigner;
  let user: HardhatEthersSigner;
  let user2: HardhatEthersSigner;

  beforeEach(async function () {
    [admin, user, user2] = await ethers.getSigners();
    const Factory = await ethers.getContractFactory("NailoNode");
    node = await Factory.deploy();
    await node.waitForDeployment();
  });

  describe("Deployment", function () {
    it("should set name and symbol", async function () {
      expect(await node.name()).to.equal("NailoNode");
      expect(await node.symbol()).to.equal("NODE");
    });

    it("should grant DEFAULT_ADMIN_ROLE and ADMIN_ROLE to deployer", async function () {
      expect(await node.hasRole(await node.DEFAULT_ADMIN_ROLE(), admin.address)).to.be.true;
      expect(await node.hasRole(await node.ADMIN_ROLE(), admin.address)).to.be.true;
    });

    it("should have MAX_SUPPLY = 200", async function () {
      expect(await node.MAX_SUPPLY()).to.equal(MAX_SUPPLY);
    });

    it("should start with nextTokenId = 0", async function () {
      expect(await node.nextTokenId()).to.equal(0);
    });

    it("should start with totalSupply = 0", async function () {
      expect(await node.totalSupply()).to.equal(0);
    });
  });

  describe("mint()", function () {
    it("should mint a token to the given address", async function () {
      await expect(node.connect(admin).mint(user.address))
        .to.emit(node, "Minted")
        .withArgs(user.address, 0);
      expect(await node.ownerOf(0)).to.equal(user.address);
      expect(await node.balanceOf(user.address)).to.equal(1);
    });

    it("should increment nextTokenId after mint", async function () {
      expect(await node.nextTokenId()).to.equal(0);
      await node.connect(admin).mint(user.address);
      expect(await node.nextTokenId()).to.equal(1);
    });

    it("should update totalSupply after mint", async function () {
      await node.connect(admin).mint(user.address);
      expect(await node.totalSupply()).to.equal(1);
      await node.connect(admin).mint(user2.address);
      expect(await node.totalSupply()).to.equal(2);
    });

    it("should mint sequential token IDs", async function () {
      await node.connect(admin).mint(user.address);
      await node.connect(admin).mint(user2.address);
      await node.connect(admin).mint(admin.address);
      expect(await node.ownerOf(0)).to.equal(user.address);
      expect(await node.ownerOf(1)).to.equal(user2.address);
      expect(await node.ownerOf(2)).to.equal(admin.address);
    });

    it("should revert if not admin", async function () {
      await expect(node.connect(user).mint(user.address)).to.be.reverted;
    });

    it("should revert if max supply reached", async function () {
      for (let i = 0; i < MAX_SUPPLY; i++) {
        const addr = ethers.Wallet.createRandom().address;
        await node.connect(admin).mint(addr);
      }
      expect(await node.totalSupply()).to.equal(MAX_SUPPLY);
      await expect(node.connect(admin).mint(user.address)).to.be.revertedWith(
        "NailoNode: max supply reached"
      );
    });

    it("should allow minting up to exactly MAX_SUPPLY", async function () {
      for (let i = 0; i < MAX_SUPPLY; i++) {
        const w = ethers.Wallet.createRandom();
        await node.connect(admin).mint(w.address);
      }
      expect(await node.totalSupply()).to.equal(MAX_SUPPLY);
    });
  });

  describe("tokenURI", function () {
    it("should return correct URI format", async function () {
      await node.connect(admin).mint(user.address);
      const uri = await node.tokenURI(0);
      expect(uri).to.equal("https://api.nailo.io/node/0.json");
    });

    it("should revert for non-existent token", async function () {
      await expect(node.tokenURI(999)).to.be.reverted;
    });
  });

  describe("ERC721Enumerable", function () {
    beforeEach(async function () {
      await node.connect(admin).mint(user.address);
      await node.connect(admin).mint(user.address);
      await node.connect(admin).mint(user2.address);
      await node.connect(admin).mint(user2.address);
    });

    it("should return correct balanceOf", async function () {
      expect(await node.balanceOf(user.address)).to.equal(2);
      expect(await node.balanceOf(user2.address)).to.equal(2);
    });

    it("should return correct totalSupply", async function () {
      expect(await node.totalSupply()).to.equal(4);
    });

    it("should return correct tokenByIndex", async function () {
      expect(await node.tokenByIndex(0)).to.equal(0);
      expect(await node.tokenByIndex(3)).to.equal(3);
    });

    it("should revert tokenByIndex out of range", async function () {
      await expect(node.tokenByIndex(4)).to.be.reverted;
    });

    it("should return correct tokenOfOwnerByIndex", async function () {
      expect(await node.tokenOfOwnerByIndex(user.address, 0)).to.equal(0);
      expect(await node.tokenOfOwnerByIndex(user.address, 1)).to.equal(1);
      expect(await node.tokenOfOwnerByIndex(user2.address, 0)).to.equal(2);
      expect(await node.tokenOfOwnerByIndex(user2.address, 1)).to.equal(3);
    });

    it("should revert tokenOfOwnerByIndex out of range", async function () {
      await expect(node.tokenOfOwnerByIndex(user.address, 2)).to.be.reverted;
    });
  });

  describe("Transfers", function () {
    beforeEach(async function () {
      await node.connect(admin).mint(user.address);
    });

    it("should allow transferFrom", async function () {
      await node.connect(user).transferFrom(user.address, user2.address, 0);
      expect(await node.ownerOf(0)).to.equal(user2.address);
    });

    it("should allow safeTransferFrom", async function () {
      await node.connect(user)["safeTransferFrom(address,address,uint256)"](
        user.address,
        user2.address,
        0
      );
      expect(await node.ownerOf(0)).to.equal(user2.address);
    });

    it("should update balances after transfer", async function () {
      await node.connect(user).transferFrom(user.address, user2.address, 0);
      expect(await node.balanceOf(user.address)).to.equal(0);
      expect(await node.balanceOf(user2.address)).to.equal(1);
    });

    it("should support approve and transferFrom by approved", async function () {
      await node.connect(user).approve(user2.address, 0);
      await node.connect(user2).transferFrom(user.address, user2.address, 0);
      expect(await node.ownerOf(0)).to.equal(user2.address);
    });

    it("should support setApprovalForAll", async function () {
      await node.connect(user).setApprovalForAll(user2.address, true);
      expect(await node.isApprovedForAll(user.address, user2.address)).to.be.true;
    });
  });

  describe("Burn", function () {
    beforeEach(async function () {
      await node.connect(admin).mint(user.address);
    });

    it("should burn token", async function () {
      await node.connect(user).burn(0);
      await expect(node.ownerOf(0)).to.be.reverted;
      expect(await node.balanceOf(user.address)).to.equal(0);
    });

    it("should decrease totalSupply after burn", async function () {
      expect(await node.totalSupply()).to.equal(1);
      await node.connect(user).burn(0);
      expect(await node.totalSupply()).to.equal(0);
    });

    it("should revert burn by non-owner", async function () {
      await expect(node.connect(user2).burn(0)).to.be.reverted;
    });
  });

  describe("supportsInterface", function () {
    it("should support ERC721 interface", async function () {
      expect(await node.supportsInterface("0x80ac58cd")).to.be.true;
    });

    it("should support ERC721Enumerable interface", async function () {
      expect(await node.supportsInterface("0x780e9d63")).to.be.true;
    });

    it("should support AccessControl interface", async function () {
      expect(await node.supportsInterface("0x7965db0b")).to.be.true;
    });
  });
});
