// SPDX-License-Identifier: MIT
pragma solidity ^0.8.34;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

interface INailoInvitation {
    function getInviterChain(address _user) external view returns (address[10] memory);
}

interface INailoKOL {
    function ownerOf(uint256 tokenId) external view returns (address);
}

interface IPancakeRouter02 {
    function WETH() external pure returns (address);

    function swapExactETHForTokens(
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external payable returns (uint256[] memory amounts);

    function addLiquidityETH(
        address token,
        uint256 amountTokenDesired,
        uint256 amountTokenMin,
        uint256 amountETHMin,
        address to,
        uint256 deadline
    ) external payable returns (uint256 amountToken, uint256 amountETH, uint256 liquidity);
}

contract NailoSocial is AccessControl, ReentrancyGuard {
    using SafeERC20 for IERC20;

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    uint256 public constant BASIS_POINTS = 10000;
    uint256 public constant KOL_SHARE = 200;
    uint256 public constant SERVER_SHARE = 300;
    uint256 public constant REFERRER_SHARE = 2000;
    uint256[10] public REFERRER_SPLITS = [500, 400, 300, 200, 100, 100, 100, 100, 100, 100];
    uint256 public constant REWARD_PRECISION = 1e12;

    INailoInvitation public immutable invitation;
    INailoKOL public immutable kolToken;
    IERC20 public immutable nailoToken;
    IPancakeRouter02 public immutable pancakeRouter;

    address public serverFeeAddress;
    address public marketingAddress;
    uint256 public subscriptionPrice = 0.08 ether;
    uint256 public minTipAmount = 0.03 ether;
    uint256 public slippageBps = 500;

    mapping(address => uint256) public subscriptionExpiry;
    mapping(address => uint256) public userLP;
    uint256 public totalLP;
    uint256 public accNailoPerLP;
    mapping(address => uint256) public rewardDebt;

    event Subscribed(address indexed user, uint256 indexed kolTokenId, uint256 months, uint256 amount, uint256 expiry);
    event Tipped(address indexed from, uint256 indexed kolTokenId, uint256 amount);
    event ReferrerPaid(address indexed referrer, uint256 amount, uint256 level);
    event MarketingReceived(uint256 amount);
    event LPCreated(address indexed user, uint256 bnbAmount, uint256 nailoAmount, uint256 lpAmount);
    event NailoDeposited(uint256 amount);
    event RewardClaimed(address indexed user, uint256 amount);
    event SubscriptionPriceUpdated(uint256 oldPrice, uint256 newPrice);
    event MinTipAmountUpdated(uint256 oldAmount, uint256 newAmount);
    event ServerFeeAddressUpdated(address oldAddr, address newAddr);
    event MarketingAddressUpdated(address oldAddr, address newAddr);
    event SlippageUpdated(uint256 oldBps, uint256 newBps);

    constructor(
        address _invitation,
        address _kolToken,
        address _nailoToken,
        address _pancakeRouter,
        address _serverFeeAddress,
        address _marketingAddress
    ) {
        require(_invitation != address(0), "Social: invalid invitation");
        require(_kolToken != address(0), "Social: invalid KOL token");
        require(_nailoToken != address(0), "Social: invalid nailo token");
        require(_pancakeRouter != address(0), "Social: invalid router");
        require(_serverFeeAddress != address(0), "Social: invalid server fee address");
        require(_marketingAddress != address(0), "Social: invalid marketing address");

        invitation = INailoInvitation(_invitation);
        kolToken = INailoKOL(_kolToken);
        nailoToken = IERC20(_nailoToken);
        pancakeRouter = IPancakeRouter02(_pancakeRouter);
        serverFeeAddress = _serverFeeAddress;
        marketingAddress = _marketingAddress;

        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);
    }

    function subscribe(uint256 kolTokenId, uint256 months) external payable nonReentrant {
        require(months == 1 || months == 6 || months == 12, "Social: months must be 1, 6, or 12");

        uint256 price = subscriptionPrice * months;
        require(msg.value >= price, "Social: insufficient BNB");

        address kolAddress = kolToken.ownerOf(kolTokenId);
        require(kolAddress != address(0), "Social: KOL token does not exist");

        _processPayment(msg.sender, price, kolAddress);

        uint256 currentExpiry = subscriptionExpiry[msg.sender];
        uint256 baseTime = currentExpiry > block.timestamp ? currentExpiry : block.timestamp;
        subscriptionExpiry[msg.sender] = baseTime + months * 30 days;

        emit Subscribed(msg.sender, kolTokenId, months, price, subscriptionExpiry[msg.sender]);

        if (msg.value > price) {
            _safeTransferETH(msg.sender, msg.value - price);
        }
    }

    function tip(uint256 kolTokenId) external payable nonReentrant {
        require(msg.value >= minTipAmount, "Social: tip below minimum");

        address kolAddress = kolToken.ownerOf(kolTokenId);
        require(kolAddress != address(0), "Social: KOL token does not exist");

        _processPayment(msg.sender, msg.value, kolAddress);

        emit Tipped(msg.sender, kolTokenId, msg.value);
    }

    function depositNailoReward(uint256 amount) external onlyRole(ADMIN_ROLE) {
        nailoToken.safeTransferFrom(msg.sender, address(this), amount);
        if (totalLP > 0) {
            accNailoPerLP += (amount * REWARD_PRECISION) / totalLP;
        }
        emit NailoDeposited(amount);
    }

    function pendingNailoReward(address user) public view returns (uint256) {
        uint256 accumulated = (userLP[user] * accNailoPerLP) / REWARD_PRECISION;
        if (accumulated <= rewardDebt[user]) return 0;
        return accumulated - rewardDebt[user];
    }

    function claimNailoReward() external nonReentrant {
        uint256 pending = pendingNailoReward(msg.sender);
        require(pending > 0, "Social: no rewards to claim");

        rewardDebt[msg.sender] = (userLP[msg.sender] * accNailoPerLP) / REWARD_PRECISION;

        nailoToken.safeTransfer(msg.sender, pending);
        emit RewardClaimed(msg.sender, pending);
    }

    function setSubscriptionPrice(uint256 _price) external onlyRole(ADMIN_ROLE) {
        emit SubscriptionPriceUpdated(subscriptionPrice, _price);
        subscriptionPrice = _price;
    }

    function setMinTipAmount(uint256 _amount) external onlyRole(ADMIN_ROLE) {
        emit MinTipAmountUpdated(minTipAmount, _amount);
        minTipAmount = _amount;
    }

    function setServerFeeAddress(address _addr) external onlyRole(ADMIN_ROLE) {
        require(_addr != address(0), "Social: invalid address");
        emit ServerFeeAddressUpdated(serverFeeAddress, _addr);
        serverFeeAddress = _addr;
    }

    function setMarketingAddress(address _addr) external onlyRole(ADMIN_ROLE) {
        require(_addr != address(0), "Social: invalid address");
        emit MarketingAddressUpdated(marketingAddress, _addr);
        marketingAddress = _addr;
    }

    function setSlippageBps(uint256 _bps) external onlyRole(ADMIN_ROLE) {
        require(_bps <= 1000, "Social: slippage too high");
        emit SlippageUpdated(slippageBps, _bps);
        slippageBps = _bps;
    }

    function isSubscribed(address user) public view returns (bool) {
        return subscriptionExpiry[user] > block.timestamp;
    }

    function getUserLP(address user) external view returns (uint256) {
        return userLP[user];
    }

    function getTotalLP() external view returns (uint256) {
        return totalLP;
    }

    function _processPayment(address user, uint256 amount, address kolAddress) internal {
        uint256 kolAmount = (amount * KOL_SHARE) / BASIS_POINTS;
        if (kolAmount > 0) {
            _safeTransferETH(kolAddress, kolAmount);
        }

        uint256 serverAmount = (amount * SERVER_SHARE) / BASIS_POINTS;
        if (serverAmount > 0) {
            _safeTransferETH(serverFeeAddress, serverAmount);
        }

        uint256 referrerPool = (amount * REFERRER_SHARE) / BASIS_POINTS;
        if (referrerPool > 0) {
            _distributeToReferrers(user, referrerPool);
        }

        uint256 lpBnb = amount - kolAmount - serverAmount - referrerPool;
        if (lpBnb > 0) {
            _addLiquidity(user, lpBnb);
        }
    }

    function _distributeToReferrers(address user, uint256 pool) internal {
        address[10] memory chain = invitation.getInviterChain(user);
        uint256 totalDistributed;

        for (uint256 i = 0; i < 10; i++) {
            address referrer = chain[i];
            if (referrer == address(0)) break;
            if (!isSubscribed(referrer)) continue;

            uint256 share = (pool * REFERRER_SPLITS[i]) / REFERRER_SHARE;
            if (share > 0) {
                _safeTransferETH(referrer, share);
                totalDistributed += share;
                emit ReferrerPaid(referrer, share, i);
            }
        }

        if (totalDistributed < pool) {
            uint256 remainder = pool - totalDistributed;
            _safeTransferETH(marketingAddress, remainder);
            emit MarketingReceived(remainder);
        }
    }

    function _addLiquidity(address user, uint256 bnbAmount) internal {
        require(bnbAmount > 0, "Social: zero LP amount");

        IERC20 nailo = nailoToken;
        IPancakeRouter02 router = pancakeRouter;

        uint256 bnbForSwap = bnbAmount / 2;
        uint256 bnbForLP = bnbAmount - bnbForSwap;

        uint256 nailoBefore = nailo.balanceOf(address(this));

        if (bnbForSwap > 0) {
            address[] memory path = new address[](2);
            path[0] = router.WETH();
            path[1] = address(nailo);

            router.swapExactETHForTokens{value: bnbForSwap}(
                (bnbForSwap * (BASIS_POINTS - slippageBps)) / BASIS_POINTS,
                path,
                address(this),
                block.timestamp + 300
            );
        }

        uint256 nailoReceived = nailo.balanceOf(address(this)) - nailoBefore;

        if (nailoReceived > 0 && bnbForLP > 0) {
            nailo.forceApprove(address(router), nailoReceived);

            try router.addLiquidityETH{value: bnbForLP}(
                address(nailo),
                nailoReceived,
                (nailoReceived * (BASIS_POINTS - slippageBps)) / BASIS_POINTS,
                (bnbForLP * (BASIS_POINTS - slippageBps)) / BASIS_POINTS,
                address(this),
                block.timestamp + 300
            ) returns (uint256, uint256, uint256 liquidity) {
                userLP[user] += liquidity;
                totalLP += liquidity;
                _updateRewardDebt(user);
                emit LPCreated(user, bnbAmount, nailoReceived, liquidity);
            } catch {
                revert("Social: add liquidity failed");
            }
        } else if (bnbForLP > 0) {
            revert("Social: no nailo for LP");
        }
    }

    function _updateRewardDebt(address user) internal {
        rewardDebt[user] = (userLP[user] * accNailoPerLP) / REWARD_PRECISION;
    }

    function _safeTransferETH(address to, uint256 amount) internal {
        require(to != address(0), "Social: transfer to zero");
        (bool success, ) = payable(to).call{value: amount}("");
        require(success, "Social: ETH transfer failed");
    }

    receive() external payable {}
}
