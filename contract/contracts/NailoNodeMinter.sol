// SPDX-License-Identifier: MIT
pragma solidity ^0.8.34;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

interface INailoNode {
    function mint(address to) external;
    function nextTokenId() external view returns (uint256);
}

contract NailoNodeMinter is AccessControl {
    using SafeERC20 for IERC20;

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    INailoNode public immutable nodeToken;
    IERC20 public immutable usdt;
    address public treasury;
    uint256 public price = 300 * 10**18;

    event Purchased(address indexed user, uint256 indexed tokenId, uint256 amount);
    event PriceUpdated(uint256 oldPrice, uint256 newPrice);
    event TreasuryUpdated(address oldTreasury, address newTreasury);

    constructor(address _nodeToken, address _usdt, address _treasury) {
        require(_nodeToken != address(0), "Minter: invalid node token");
        require(_usdt != address(0), "Minter: invalid USDT address");
        require(_treasury != address(0), "Minter: invalid treasury");

        nodeToken = INailoNode(_nodeToken);
        usdt = IERC20(_usdt);
        treasury = _treasury;

        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);
    }

    function buy() external {
        usdt.safeTransferFrom(msg.sender, treasury, price);

        uint256 tokenId = nodeToken.nextTokenId();
        nodeToken.mint(msg.sender);

        emit Purchased(msg.sender, tokenId, price);
    }

    function setPrice(uint256 _price) external onlyRole(ADMIN_ROLE) {
        emit PriceUpdated(price, _price);
        price = _price;
    }

    function setTreasury(address _treasury) external onlyRole(ADMIN_ROLE) {
        require(_treasury != address(0), "Minter: invalid treasury");
        emit TreasuryUpdated(treasury, _treasury);
        treasury = _treasury;
    }
}
