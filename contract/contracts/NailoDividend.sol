// SPDX-License-Identifier: MIT
pragma solidity ^0.8.34;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

interface INailoNodeDividend {
    function ownerOf(uint256 tokenId) external view returns (address);
    function totalSupply() external view returns (uint256);
    function balanceOf(address owner) external view returns (uint256);
    function tokenOfOwnerByIndex(address owner, uint256 index) external view returns (uint256);
    function nextTokenId() external view returns (uint256);
}

contract NailoDividend is AccessControl, ReentrancyGuard {
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    uint256 public constant PRECISION = 1e12;

    INailoNodeDividend public immutable nodeToken;

    uint256 public accDividendPerToken;
    mapping(uint256 => uint256) public lastClaimed;

    event Deposited(address indexed from, uint256 amount, uint256 accDividendPerToken);
    event Claimed(address indexed owner, uint256 indexed tokenId, uint256 amount);

    constructor(address _nodeToken) {
        require(_nodeToken != address(0), "Dividend: invalid node token");
        nodeToken = INailoNodeDividend(_nodeToken);

        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);
    }

    function deposit() public payable {
        require(msg.value > 0, "Dividend: zero deposit");

        uint256 supply = nodeToken.totalSupply();
        if (supply > 0) {
            accDividendPerToken += (msg.value * PRECISION) / supply;
        }

        emit Deposited(msg.sender, msg.value, accDividendPerToken);
    }

    function pendingDividend(uint256 tokenId) public view returns (uint256) {
        if (tokenId >= nodeToken.nextTokenId()) return 0;
        uint256 accumulated = accDividendPerToken;
        if (accumulated <= lastClaimed[tokenId]) return 0;
        return (accumulated - lastClaimed[tokenId]) / PRECISION;
    }

    function claim() external nonReentrant {
        uint256 count = nodeToken.balanceOf(msg.sender);
        uint256 totalPending;

        for (uint256 i = 0; i < count; i++) {
            uint256 tokenId = nodeToken.tokenOfOwnerByIndex(msg.sender, i);
            uint256 pending = pendingDividend(tokenId);
            if (pending > 0) {
                lastClaimed[tokenId] = accDividendPerToken;
                totalPending += pending;
                emit Claimed(msg.sender, tokenId, pending);
            }
        }

        require(totalPending > 0, "Dividend: nothing to claim");
        _safeTransferETH(msg.sender, totalPending);
    }

    function getMyTokenIds() external view returns (uint256[] memory) {
        uint256 count = nodeToken.balanceOf(msg.sender);
        uint256[] memory ids = new uint256[](count);
        for (uint256 i = 0; i < count; i++) {
            ids[i] = nodeToken.tokenOfOwnerByIndex(msg.sender, i);
        }
        return ids;
    }

    function pendingDividendBatch(uint256[] calldata tokenIds) external view returns (uint256[] memory) {
        uint256[] memory result = new uint256[](tokenIds.length);
        for (uint256 i = 0; i < tokenIds.length; i++) {
            result[i] = pendingDividend(tokenIds[i]);
        }
        return result;
    }

    function _safeTransferETH(address to, uint256 amount) internal {
        (bool success, ) = payable(to).call{value: amount}("");
        require(success, "Dividend: ETH transfer failed");
    }

    receive() external payable {
        deposit();
    }
}
