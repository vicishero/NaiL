// SPDX-License-Identifier: MIT
pragma solidity ^0.8.34;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract Nailo is ERC20 {
    constructor() ERC20("Nailo", "NAILO") {
        _mint(msg.sender, 1000000000 * 10 ** decimals());
    }
}
