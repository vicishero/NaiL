// SPDX-License-Identifier: MIT
pragma solidity ^0.8.34;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract NailoKOL is ERC721, ERC721Burnable, Ownable {
    uint256 private _nextTokenId;

    function nextTokenId() public view returns (uint256) {
        return _nextTokenId;
    }

    constructor() ERC721("NailoKOL", "KOL") Ownable(msg.sender) {}

    function safeMint(address to) public onlyOwner {
        uint256 tokenId = _nextTokenId++;
        _safeMint(to, tokenId);
    }

    function tokenURI(uint256 tokenId) public view override returns (string memory) {
        _requireOwned(tokenId);
        return string(abi.encodePacked("https://api.nailo.io/kol/", Strings.toString(tokenId), ".json"));
    }
}
