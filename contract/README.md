# Nailo 智能合约

基于Hardhat + TypeScript开发的Nailo ERC-20智能合约。

## 合约信息

- **名称**: Nailo
- **符号**: NAILO
- **小数位数**: 18
- **初始供应量**: 1,000,000,000 NAILO

## 安装依赖

```bash
npm install
```

## 编译合约

```bash
npm run compile
# 或
npx hardhat compile
```

## 运行测试

```bash
npm run test
# 或
npx hardhat test
```

## 部署合约

### 本地Hardhat网络（临时，用于测试）

```bash
npm run deploy
# 或
npx hardhat run scripts/deploy.ts
```

### 本地节点（持久）

1. 启动本地节点：
```bash
npm run node
# 或
npx hardhat node
```

2. 在新终端中部署到本地节点：
```bash
npm run deploy:localhost
# 或
npx hardhat run scripts/deploy.ts --network localhost
```

### 测试网/主网部署

1. 复制 `.env.example` 为 `.env` 并填写配置：
```bash
cp .env.example .env
# 编辑 .env 文件，填写私钥和RPC URL
```

2. 部署到Sepolia测试网：
```bash
npx hardhat run scripts/deploy.ts --network sepolia
```

3. 部署到以太坊主网：
```bash
npx hardhat run scripts/deploy.ts --network mainnet
```

## 验证合约（Etherscan）

在 `.env` 中配置 `ETHERSCAN_API_KEY` 后：

```bash
npx hardhat verify --network sepolia <合约地址>
```

## 常用Hardhat命令

```bash
# 查看帮助
npx hardhat help

# 查看可用账户
npx hardhat accounts

# 清理编译产物
npx hardhat clean

# 查看gas报告
REPORT_GAS=true npx hardhat test

# 运行Hardhat控制台
npx hardhat console
```

## 技术栈

- **Hardhat**: 智能合约开发框架
- **Solidity**: 0.8.24
- **TypeScript**: 类型安全
- **OpenZeppelin Contracts**: 5.0.2
- **Ethers.js**: 6.x
- **Mocha + Chai**: 测试框架
