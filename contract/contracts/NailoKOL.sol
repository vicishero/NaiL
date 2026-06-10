// SPDX-License-Identifier: MIT
pragma solidity ^0.8.34;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract NailoKOL is ERC721, ERC721Burnable, AccessControl {
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    uint256 private _nextTokenId;
    mapping(address => uint256) public expireTime;

    event ExpireTimeUpdated(address indexed user, uint256 expireTime);

    constructor() ERC721("NailoKOL", "KOL") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);
    }

    function mint(address to) public onlyRole(ADMIN_ROLE) {
        require(balanceOf(to) == 0, "NailoKOL: already has a token");
        uint256 tokenId = _nextTokenId++;
        _safeMint(to, tokenId);
    }

    function setExpireTime(address user, uint256 time) public onlyRole(ADMIN_ROLE) {
        expireTime[user] = time;
        emit ExpireTimeUpdated(user, time);
    }

    function nextTokenId() public view returns (uint256) {
        return _nextTokenId;
    }

    function tokenURI(uint256 tokenId) public view override returns (string memory) {
        _requireOwned(tokenId);
        return string(abi.encodePacked("https://api.nailo.io/kol/", Strings.toString(tokenId), ".json"));
    }

    function _update(address to, uint256 tokenId, address auth) internal override returns (address) {
        address from = _ownerOf(tokenId);
        if (from != address(0) && to != address(0)) {
            revert("NailoKOL: transfer is disabled");
        }
        return super._update(to, tokenId, auth);
    }

    function approve(address to, uint256 tokenId) public pure override {
        revert("NailoKOL: approve is disabled");
    }

    function setApprovalForAll(address operator, bool approved) public pure override {
        revert("NailoKOL: setApprovalForAll is disabled");
    }

    function supportsInterface(bytes4 interfaceId) public view override(ERC721, AccessControl) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
