// SPDX-License-Identifier: MIT
pragma solidity ^0.8.34;

import "./NailoKOL.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract NailoKolMinter is AccessControl {
    using SafeERC20 for IERC20;

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    NailoKOL public immutable kolToken;
    IERC20 public usdt;
    address public treasury;
    uint256 public monthlyPrice = 10 * 10**6;

    event Minted(address indexed user, uint256 indexed tokenId, uint256 months);
    event Renewed(address indexed user, uint256 months);
    event MonthlyPriceUpdated(uint256 oldPrice, uint256 newPrice);
    event UsdtAddressUpdated(address oldUsdt, address newUsdt);
    event TreasuryUpdated(address oldTreasury, address newTreasury);

    constructor(address _kolToken, address _treasury) {
        require(_kolToken != address(0), "Minter: invalid token address");
        require(_treasury != address(0), "Minter: invalid treasury address");
        kolToken = NailoKOL(_kolToken);
        treasury = _treasury;
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);
    }

    function buy(uint256 months) public {
        require(months > 0, "Minter: months must be greater than 0");
        require(address(usdt) != address(0), "Minter: USDT not set");

        uint256 amount = monthlyPrice * months;
        usdt.safeTransferFrom(msg.sender, treasury, amount);

        if (kolToken.balanceOf(msg.sender) == 0) {
            uint256 tokenId = kolToken.nextTokenId();
            kolToken.mint(msg.sender);
            kolToken.setExpireTime(msg.sender, block.timestamp + months * 30 days);
            emit Minted(msg.sender, tokenId, months);
        } else {
            uint256 currentExpire = kolToken.expireTime(msg.sender);
            uint256 newExpire = currentExpire > block.timestamp ? currentExpire : block.timestamp;
            kolToken.setExpireTime(msg.sender, newExpire + months * 30 days);
            emit Renewed(msg.sender, months);
        }
    }

    function setMonthlyPrice(uint256 _monthlyPrice) public onlyRole(ADMIN_ROLE) {
        emit MonthlyPriceUpdated(monthlyPrice, _monthlyPrice);
        monthlyPrice = _monthlyPrice;
    }

    function setUsdtAddress(address _usdt) public onlyRole(ADMIN_ROLE) {
        emit UsdtAddressUpdated(address(usdt), _usdt);
        usdt = IERC20(_usdt);
    }

    function setTreasury(address _treasury) public onlyRole(ADMIN_ROLE) {
        emit TreasuryUpdated(treasury, _treasury);
        treasury = _treasury;
    }
}
