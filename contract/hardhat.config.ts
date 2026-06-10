import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "dotenv/config";

const PRIVATE_KEY = process.env.PRIVATE_KEY || "";
const SEPOLIA_RPC_URL = process.env.SEPOLIA_RPC_URL || "";
const MAINNET_RPC_URL = process.env.MAINNET_RPC_URL || "";
const ETHERSCAN_API_KEY = process.env.ETHERSCAN_API_KEY || "";

const networks: HardhatUserConfig["networks"] = {
  hardhat: {
    chainId: 31337
  },
  localhost: {
    url: "http://127.0.0.1:8545",
    chainId: 31337
  }
};

if (PRIVATE_KEY && SEPOLIA_RPC_URL) {
  networks.sepolia = {
    url: SEPOLIA_RPC_URL,
    accounts: [PRIVATE_KEY],
    chainId: 11155111
  };
}

if (PRIVATE_KEY && MAINNET_RPC_URL) {
  networks.mainnet = {
    url: MAINNET_RPC_URL,
    accounts: [PRIVATE_KEY],
    chainId: 1
  };
}

const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.34",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200
      },
      evmVersion: "cancun"
    }
  },
  networks,
  etherscan: {
    apiKey: ETHERSCAN_API_KEY
  },
  paths: {
    sources: "./contracts",
    tests: "./test",
    cache: "./cache",
    artifacts: "./artifacts"
  }
};

export default config;
