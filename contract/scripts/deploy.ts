import { ethers } from "hardhat";

async function main() {
  const [deployer] = await ethers.getSigners();

  console.log("使用账户部署合约:", deployer.address);
  console.log("账户余额:", ethers.formatEther(await deployer.provider.getBalance(deployer.address)));

  const Nailo = await ethers.getContractFactory("Nailo");
  const nailo = await Nailo.deploy();
  await nailo.waitForDeployment();

  const contractAddress = await nailo.getAddress();
  console.log("Nailo合约部署成功!");
  console.log("合约地址:", contractAddress);

  // 验证部署
  const name = await nailo.name();
  const symbol = await nailo.symbol();
  const totalSupply = await nailo.totalSupply();
  const deployerBalance = await nailo.balanceOf(deployer.address);

  console.log("\n合约信息:");
  console.log("名称:", name);
  console.log("符号:", symbol);
  console.log("总供应量:", ethers.formatEther(totalSupply), symbol);
  console.log("部署者余额:", ethers.formatEther(deployerBalance), symbol);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
