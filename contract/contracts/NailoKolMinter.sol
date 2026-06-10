// SPDX-License-Identifier: MIT
pragma solidity ^0.8.34;

import "./NailoKOL.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

interface IPancakeRouter02 {
    function swapExactTokensForTokens(
        uint256 amountIn,
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external returns (uint256[] memory amounts);

    function getAmountsOut(
        uint256 amountIn,
        address[] calldata path
    ) external view returns (uint256[] memory amounts);

    function getAmountsIn(
        uint256 amountOut,
        address[] calldata path
    ) external view returns (uint256[] memory amounts);
}

contract NailoKolMinter is AccessControl {
    using SafeERC20 for IERC20;

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    struct PaymentToken {
        IERC20 token;
        address projectParty;
        address[] swapPath;
        bool active;
    }

    NailoKOL public immutable kolToken;
    IERC20 public immutable usdt;
    IPancakeRouter02 public immutable pancakeRouter;
    address public treasury;
    uint256 public monthlyPrice = 10 * 10**18;
    uint256 public slippageBps = 500;

    mapping(address => PaymentToken) private _paymentTokens;

    event Minted(address indexed user, uint256 indexed tokenId, uint256 months);
    event Renewed(address indexed user, uint256 months);
    event MonthlyPriceUpdated(uint256 oldPrice, uint256 newPrice);
    event TreasuryUpdated(address oldTreasury, address newTreasury);
    event SlippageUpdated(uint256 oldSlippageBps, uint256 newSlippageBps);
    event PaymentTokenAdded(address indexed token, address indexed projectParty);
    event PaymentTokenUpdated(address indexed token, address indexed projectParty);
    event PaymentTokenRemoved(address indexed token);
    event TokenPurchased(
        address indexed user,
        address indexed paymentToken,
        uint256 amountPaid,
        uint256 months,
        uint256 usdtReceived
    );

    constructor(address _kolToken, address _usdt, address _pancakeRouter, address _treasury) {
        require(_kolToken != address(0), "Minter: invalid kol token");
        require(_usdt != address(0), "Minter: invalid USDT address");
        require(_pancakeRouter != address(0), "Minter: invalid router address");
        require(_treasury != address(0), "Minter: invalid treasury");

        kolToken = NailoKOL(_kolToken);
        usdt = IERC20(_usdt);
        pancakeRouter = IPancakeRouter02(_pancakeRouter);
        treasury = _treasury;

        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);
    }

    function buy(uint256 months) external {
        require(months > 0, "Minter: months must be greater than 0");

        uint256 amount = monthlyPrice * months;
        usdt.safeTransferFrom(msg.sender, treasury, amount);

        _processPurchase(msg.sender, months);

        emit TokenPurchased(msg.sender, address(usdt), amount, months, amount);
    }

    function buyWithToken(address tokenAddress, uint256 months) external {
        require(months > 0, "Minter: months must be greater than 0");

        PaymentToken storage pt = _paymentTokens[tokenAddress];
        require(pt.active, "Minter: token not configured");
        require(pt.swapPath.length >= 2, "Minter: swap path not configured");

        uint256 usdtValue = monthlyPrice * months;
        uint256[] memory amountsIn = pancakeRouter.getAmountsIn(usdtValue, pt.swapPath);
        uint256 amount = amountsIn[0];

        uint256 amountWithBuffer = (amount * (10000 + slippageBps)) / 10000;
        uint256 half = amountWithBuffer / 2;

        pt.token.safeTransferFrom(msg.sender, address(this), amountWithBuffer);

        if (half > 0) {
            pt.token.safeTransfer(pt.projectParty, half);
        }

        uint256 swapAmount = amountWithBuffer - half;
        uint256[] memory amountsOut = pancakeRouter.getAmountsOut(swapAmount, pt.swapPath);
        uint256 expectedOut = amountsOut[amountsOut.length - 1];
        uint256 minOut = (expectedOut * (10000 - slippageBps)) / 10000;

        pt.token.forceApprove(address(pancakeRouter), swapAmount);
        pancakeRouter.swapExactTokensForTokens(
            swapAmount,
            minOut,
            pt.swapPath,
            treasury,
            block.timestamp + 300
        );

        _processPurchase(msg.sender, months);

        emit TokenPurchased(msg.sender, tokenAddress, amountWithBuffer, months, expectedOut);
    }

    function setMonthlyPrice(uint256 _monthlyPrice) external onlyRole(ADMIN_ROLE) {
        emit MonthlyPriceUpdated(monthlyPrice, _monthlyPrice);
        monthlyPrice = _monthlyPrice;
    }

    function setTreasury(address _treasury) external onlyRole(ADMIN_ROLE) {
        require(_treasury != address(0), "Minter: invalid treasury");
        emit TreasuryUpdated(treasury, _treasury);
        treasury = _treasury;
    }

    function setSlippageBps(uint256 _slippageBps) external onlyRole(ADMIN_ROLE) {
        require(_slippageBps <= 5000, "Minter: slippage too high");
        emit SlippageUpdated(slippageBps, _slippageBps);
        slippageBps = _slippageBps;
    }

    function addPaymentToken(
        address tokenAddress,
        address projectParty,
        address[] calldata swapPath
    ) external onlyRole(ADMIN_ROLE) {
        require(tokenAddress != address(0), "Minter: invalid token");
        require(!_paymentTokens[tokenAddress].active, "Minter: token already added");
        require(swapPath.length >= 2, "Minter: swap path too short");
        require(swapPath[swapPath.length - 1] == address(usdt), "Minter: path must end with USDT");

        _paymentTokens[tokenAddress] = PaymentToken({
            token: IERC20(tokenAddress),
            projectParty: projectParty,
            swapPath: swapPath,
            active: true
        });

        emit PaymentTokenAdded(tokenAddress, projectParty);
    }

    function updatePaymentToken(
        address tokenAddress,
        address projectParty,
        address[] calldata swapPath
    ) external onlyRole(ADMIN_ROLE) {
        require(_paymentTokens[tokenAddress].active, "Minter: token not found");
        require(swapPath.length >= 2, "Minter: swap path too short");
        require(swapPath[swapPath.length - 1] == address(usdt), "Minter: path must end with USDT");

        _paymentTokens[tokenAddress].projectParty = projectParty;
        _paymentTokens[tokenAddress].swapPath = swapPath;

        emit PaymentTokenUpdated(tokenAddress, projectParty);
    }

    function removePaymentToken(address tokenAddress) external onlyRole(ADMIN_ROLE) {
        require(_paymentTokens[tokenAddress].active, "Minter: token not found");
        _paymentTokens[tokenAddress].active = false;
        emit PaymentTokenRemoved(tokenAddress);
    }

    function getPaymentToken(address tokenAddress) external view returns (PaymentToken memory) {
        require(_paymentTokens[tokenAddress].active, "Minter: token not found");
        return _paymentTokens[tokenAddress];
    }

    function isTokenSupported(address tokenAddress) external view returns (bool) {
        return _paymentTokens[tokenAddress].active;
    }

    function getRequiredTokens(address tokenAddress, uint256 months) external view returns (uint256) {
        PaymentToken storage pt = _paymentTokens[tokenAddress];
        require(pt.active, "Minter: token not configured");
        require(months > 0, "Minter: months must be greater than 0");
        require(pt.swapPath.length >= 2, "Minter: no swap path");

        uint256 usdtValue = monthlyPrice * months;
        uint256[] memory amountsIn = pancakeRouter.getAmountsIn(usdtValue, pt.swapPath);
        return (amountsIn[0] * (10000 + slippageBps)) / 10000;
    }

    function estimateSwap(address tokenAddress, uint256 amountIn) external view returns (uint256) {
        PaymentToken storage pt = _paymentTokens[tokenAddress];
        require(pt.active, "Minter: token not configured");
        require(pt.swapPath.length >= 2, "Minter: no swap path");
        uint256[] memory amounts = pancakeRouter.getAmountsOut(amountIn, pt.swapPath);
        return amounts[amounts.length - 1];
    }

    function _processPurchase(address user, uint256 months) internal {
        uint256 expire = block.timestamp + months * 30 days;

        if (kolToken.balanceOf(user) == 0) {
            uint256 tokenId = kolToken.nextTokenId();
            kolToken.mint(user);
            kolToken.setExpireTime(user, expire);
            emit Minted(user, tokenId, months);
        } else {
            uint256 currentExpire = kolToken.expireTime(user);
            uint256 baseTime = currentExpire > block.timestamp ? currentExpire : block.timestamp;
            kolToken.setExpireTime(user, baseTime + months * 30 days);
            emit Renewed(user, months);
        }
    }
}
