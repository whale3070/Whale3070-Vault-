/*
 * 版权所有 (C) [2026] [whale3070/ Whale-Valut-NFT团队]
 * 本项目基于 CC BY-NC 4.0 协议开源，禁止第三方商用（详见仓库 LICENSE 文件）。
 * 著作权人保留本项目的商业使用权。
 */

pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract QuickNFT is ERC721, Ownable {
    uint256 private _nextTokenId;

    // 构造函数：Ownable(msg.sender) 确保部署者就是管理员
    constructor() ERC721("Whale Vault Badge", "WVB") Ownable(msg.sender) {}

    /**
     * @dev 核心铸造函数
     * 修正：func 改为 function
     */
    function mint(address to) public onlyOwner {
        uint256 tokenId = _nextTokenId;
        _safeMint(to, tokenId);
        _nextTokenId++; // 逻辑清晰：先发再加
    }

    /**
     * @dev 增加一个查询接口，方便你用 cast call 检查
     */
    function nextTokenId() public view returns (uint256) {
        return _nextTokenId;
    }

    // 基础 URI
    function _baseURI() internal pure override returns (string memory) {
        return "https://api.whale3070.com/metadata/";
    }
}
