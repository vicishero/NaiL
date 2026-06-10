// SPDX-License-Identifier: MIT
pragma solidity ^0.8.34;

import "./NailoKOL.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract NailoKolMinter is AccessControl {
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    NailoKOL public immutable kolToken;

    event Minted(address indexed to, uint256 indexed tokenId);
    event BatchMinted(address indexed to, uint256[] tokenIds);

    constructor(address _kolToken) {
        require(_kolToken != address(0), "Minter: invalid token address");
        kolToken = NailoKOL(_kolToken);
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(MINTER_ROLE, msg.sender);
    }

    function mint(address to) public onlyRole(MINTER_ROLE) {
        uint256 tokenId = kolToken.nextTokenId();
        kolToken.safeMint(to);
        emit Minted(to, tokenId);
    }

    function batchMint(address to, uint256 amount) public onlyRole(MINTER_ROLE) {
        require(amount > 0, "Minter: amount must be greater than 0");
        require(amount <= 100, "Minter: amount exceeds limit");

        uint256[] memory tokenIds = new uint256[](amount);
        for (uint256 i = 0; i < amount; i++) {
            tokenIds[i] = kolToken.nextTokenId();
            kolToken.safeMint(to);
        }
        emit BatchMinted(to, tokenIds);
    }

    function addMinter(address account) public onlyRole(DEFAULT_ADMIN_ROLE) {
        grantRole(MINTER_ROLE, account);
    }

    function removeMinter(address account) public onlyRole(DEFAULT_ADMIN_ROLE) {
        revokeRole(MINTER_ROLE, account);
    }
}
