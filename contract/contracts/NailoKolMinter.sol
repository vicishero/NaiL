// SPDX-License-Identifier: MIT
pragma solidity ^0.8.34;

import "./NailoKOL.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract NailoKolMinter is AccessControl {
    using SafeERC20 for IERC20;

    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    NailoKOL public immutable kolToken;
    IERC20 public usdt;
    uint256 public price = 200 * 10**6;

    event Minted(address indexed to, uint256 indexed tokenId);
    event PriceUpdated(uint256 oldPrice, uint256 newPrice);
    event UsdtAddressUpdated(address oldUsdt, address newUsdt);

    constructor(address _kolToken) {
        require(_kolToken != address(0), "Minter: invalid token address");
        kolToken = NailoKOL(_kolToken);
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(MINTER_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);
    }

    function mint(address to) public onlyRole(MINTER_ROLE) {
        uint256 tokenId = kolToken.nextTokenId();
        kolToken.safeMint(to);
        emit Minted(to, tokenId);
    }

    function buy() public {
        require(address(usdt) != address(0), "Minter: USDT not set");
        usdt.safeTransferFrom(msg.sender, address(this), price);
        uint256 tokenId = kolToken.nextTokenId();
        kolToken.safeMint(msg.sender);
        emit Minted(msg.sender, tokenId);
    }

    function setPrice(uint256 _price) public onlyRole(ADMIN_ROLE) {
        emit PriceUpdated(price, _price);
        price = _price;
    }

    function setUsdtAddress(address _usdt) public onlyRole(ADMIN_ROLE) {
        emit UsdtAddressUpdated(address(usdt), _usdt);
        usdt = IERC20(_usdt);
    }

    function withdrawUsdt(address to, uint256 amount) public onlyRole(ADMIN_ROLE) {
        usdt.safeTransfer(to, amount);
    }

    function addMinter(address account) public onlyRole(DEFAULT_ADMIN_ROLE) {
        grantRole(MINTER_ROLE, account);
    }

    function removeMinter(address account) public onlyRole(DEFAULT_ADMIN_ROLE) {
        revokeRole(MINTER_ROLE, account);
    }
}
