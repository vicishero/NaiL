// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @title Nailo ERC-20 Token
 * @dev 标准ERC-20代币，初始供应量10亿，部署时全部铸造给部署者
 */
contract Nailo is ERC20 {
    /**
     * @dev 构造函数，初始化代币并铸造初始供应量
     * 代币名称: Nailo
     * 代币符号: NAILO
     * 小数位数: 18 (ERC20标准)
     * 初始供应量: 1,000,000,000 NAILO
     */
    constructor() ERC20("Nailo", "NAILO") {
        _mint(msg.sender, 1000000000 * 10 ** decimals());
    }
}
