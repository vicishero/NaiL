// SPDX-License-Identifier: MIT
pragma solidity ^0.8.34;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract MockPancakeRouter {
    using SafeERC20 for IERC20;

    IERC20 public immutable usdt;
    uint256 public rateN = 1;
    uint256 public rateD = 1;
    uint256 public swapPenaltyBps = 0;

    constructor(address _usdt) {
        usdt = IERC20(_usdt);
    }

    function setRate(uint256 _rateN, uint256 _rateD) external {
        rateN = _rateN;
        rateD = _rateD;
    }

    function setSwapPenaltyBps(uint256 _swapPenaltyBps) external {
        swapPenaltyBps = _swapPenaltyBps;
    }

    function getAmountsOut(uint256 amountIn, address[] calldata) external view returns (uint256[] memory amounts) {
        amounts = new uint256[](2);
        amounts[0] = amountIn;
        amounts[1] = amountIn * rateN / rateD;
    }

    function getAmountsIn(uint256 amountOut, address[] calldata) external view returns (uint256[] memory amounts) {
        amounts = new uint256[](2);
        amounts[0] = (amountOut * rateD + rateN - 1) / rateN;
        amounts[1] = amountOut;
    }

    function swapExactTokensForTokens(
        uint256 amountIn,
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256
    ) external returns (uint256[] memory amounts) {
        uint256 grossOut = amountIn * rateN / rateD;
        uint256 amountOut = (grossOut * (10000 - swapPenaltyBps)) / 10000;
        require(amountOut >= amountOutMin, "INSUFFICIENT_OUTPUT_AMOUNT");

        IERC20(path[0]).safeTransferFrom(msg.sender, address(this), amountIn);
        usdt.safeTransfer(to, amountOut);

        amounts = new uint256[](2);
        amounts[0] = amountIn;
        amounts[1] = amountOut;
    }

    function fund(uint256 amount) external {
        usdt.safeTransferFrom(msg.sender, address(this), amount);
    }
}
