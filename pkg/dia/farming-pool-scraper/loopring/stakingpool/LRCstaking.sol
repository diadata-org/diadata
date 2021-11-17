/**
 *Submitted for verification at Etherscan.io on 2019-11-20
*/

/**
Author: Loopring Foundation (Loopring Project Ltd)
*/

pragma solidity ^0.5.11;


contract Ownable {
    address public owner;

    event OwnershipTransferred(
        address indexed previousOwner,
        address indexed newOwner
    );

    
    
    constructor()
        public
    {
        owner = msg.sender;
    }

    
    modifier onlyOwner()
    {
        require(msg.sender == owner, "UNAUTHORIZED");
        _;
    }

    
    
    
    function transferOwnership(
        address newOwner
        )
        public
        onlyOwner
    {
        require(newOwner != address(0), "ZERO_ADDRESS");
        emit OwnershipTransferred(owner, newOwner);
        owner = newOwner;
    }

    function renounceOwnership()
        public
        onlyOwner
    {
        emit OwnershipTransferred(owner, address(0));
        owner = address(0);
    }
}

contract Claimable is Ownable
{
    address public pendingOwner;

    
    modifier onlyPendingOwner() {
        require(msg.sender == pendingOwner, "UNAUTHORIZED");
        _;
    }

    
    
    function transferOwnership(
        address newOwner
        )
        public
        onlyOwner
    {
        require(newOwner != address(0) && newOwner != owner, "INVALID_ADDRESS");
        pendingOwner = newOwner;
    }

    
    function claimOwnership()
        public
        onlyPendingOwner
    {
        emit OwnershipTransferred(owner, pendingOwner);
        owner = pendingOwner;
        pendingOwner = address(0);
    }
}

contract ERC20 {
    function totalSupply()
        public
        view
        returns (uint);

    function balanceOf(
        address who
        )
        public
        view
        returns (uint);

    function allowance(
        address owner,
        address spender
        )
        public
        view
        returns (uint);

    function transfer(
        address to,
        uint value
        )
        public
        returns (bool);

    function transferFrom(
        address from,
        address to,
        uint    value
        )
        public
        returns (bool);

    function approve(
        address spender,
        uint    value
        )
        public
        returns (bool);
}

library ERC20SafeTransfer {
    function safeTransferAndVerify(
        address token,
        address to,
        uint    value
        )
        internal
    {
        safeTransferWithGasLimitAndVerify(
            token,
            to,
            value,
            gasleft()
        );
    }

    function safeTransfer(
        address token,
        address to,
        uint    value
        )
        internal
        returns (bool)
    {
        return safeTransferWithGasLimit(
            token,
            to,
            value,
            gasleft()
        );
    }

    function safeTransferWithGasLimitAndVerify(
        address token,
        address to,
        uint    value,
        uint    gasLimit
        )
        internal
    {
        require(
            safeTransferWithGasLimit(token, to, value, gasLimit),
            "TRANSFER_FAILURE"
        );
    }

    function safeTransferWithGasLimit(
        address token,
        address to,
        uint    value,
        uint    gasLimit
        )
        internal
        returns (bool)
    {
        
        
        

        
        bytes memory callData = abi.encodeWithSelector(
            bytes4(0xa9059cbb),
            to,
            value
        );
        (bool success, ) = token.call.gas(gasLimit)(callData);
        return checkReturnValue(success);
    }

    function safeTransferFromAndVerify(
        address token,
        address from,
        address to,
        uint    value
        )
        internal
    {
        safeTransferFromWithGasLimitAndVerify(
            token,
            from,
            to,
            value,
            gasleft()
        );
    }

    function safeTransferFrom(
        address token,
        address from,
        address to,
        uint    value
        )
        internal
        returns (bool)
    {
        return safeTransferFromWithGasLimit(
            token,
            from,
            to,
            value,
            gasleft()
        );
    }

    function safeTransferFromWithGasLimitAndVerify(
        address token,
        address from,
        address to,
        uint    value,
        uint    gasLimit
        )
        internal
    {
        bool result = safeTransferFromWithGasLimit(
            token,
            from,
            to,
            value,
            gasLimit
        );
        require(result, "TRANSFER_FAILURE");
    }

    function safeTransferFromWithGasLimit(
        address token,
        address from,
        address to,
        uint    value,
        uint    gasLimit
        )
        internal
        returns (bool)
    {
        
        
        

        
        bytes memory callData = abi.encodeWithSelector(
            bytes4(0x23b872dd),
            from,
            to,
            value
        );
        (bool success, ) = token.call.gas(gasLimit)(callData);
        return checkReturnValue(success);
    }

    function checkReturnValue(
        bool success
        )
        internal
        pure
        returns (bool)
    {
        
        
        
        if (success) {
            assembly {
                switch returndatasize()
                
                case 0 {
                    success := 1
                }
                
                case 32 {
                    returndatacopy(0, 0, 32)
                    success := mload(0)
                }
                
                default {
                    success := 0
                }
            }
        }
        return success;
    }
}

library MathUint {
    function mul(
        uint a,
        uint b
        )
        internal
        pure
        returns (uint c)
    {
        c = a * b;
        require(a == 0 || c / a == b, "MUL_OVERFLOW");
    }

    function sub(
        uint a,
        uint b
        )
        internal
        pure
        returns (uint)
    {
        require(b <= a, "SUB_UNDERFLOW");
        return a - b;
    }

    function add(
        uint a,
        uint b
        )
        internal
        pure
        returns (uint c)
    {
        c = a + b;
        require(c >= a, "ADD_OVERFLOW");
    }

    function decodeFloat(
        uint f
        )
        internal
        pure
        returns (uint value)
    {
        uint numBitsMantissa = 23;
        uint exponent = f >> numBitsMantissa;
        uint mantissa = f & ((1 << numBitsMantissa) - 1);
        value = mantissa * (10 ** exponent);
    }
}

contract ReentrancyGuard {
    
    uint private _guardValue;

    
    modifier nonReentrant()
    {
        
        require(_guardValue == 0, "REENTRANCY");

        
        _guardValue = 1;

        
        _;

        
        _guardValue = 0;
    }
}

contract IProtocolFeeVault {
    uint public constant REWARD_PERCENTAGE      = 70;
    uint public constant DAO_PERDENTAGE         = 20;

    address public userStakingPoolAddress;
    address public lrcAddress;
    address public tokenSellerAddress;
    address public daoAddress;

    uint claimedReward;
    uint claimedDAOFund;
    uint claimedBurn;

    event LRCClaimed(uint amount);
    event DAOFunded(uint amountDAO, uint amountBurn);
    event TokenSold(address token, uint amount);
    event SettingsUpdated(uint time);

    
    
    
    
    function updateSettings(
        address _userStakingPoolAddress,
        address _tokenSellerAddress,
        address _daoAddress
        )
        external;

    
    
    
    
    
    function claimStakingReward(uint amount) external;

    
    
    function fundDAO() external;

    
    
    
    
    function sellTokenForLRC(
        address token,
        uint    amount
        )
        external;

    
    
    
    
    
    
    
    
    
    function getProtocolFeeStats()
        public
        view
        returns (
            uint accumulatedFees,
            uint accumulatedBurn,
            uint accumulatedDAOFund,
            uint accumulatedReward,
            uint remainingFees,
            uint remainingBurn,
            uint remainingDAOFund,
            uint remainingReward
        );
}

contract IUserStakingPool {
    uint public constant MIN_CLAIM_DELAY        = 90 days;
    uint public constant MIN_WITHDRAW_DELAY     = 90 days;

    address public lrcAddress;
    address public protocolFeeVaultAddress;

    uint    public numAddresses;

    event ProtocolFeeVaultChanged (address feeVaultAddress);

    event LRCStaked       (address indexed user,  uint amount);
    event LRCWithdrawn    (address indexed user,  uint amount);
    event LRCRewarded     (address indexed user,  uint amount);

    
    
    function setProtocolFeeVault(address _protocolFeeVaultAddress)
        external;

    
    function getTotalStaking()
        public
        view
        returns (uint);

    
    
    
    
    
    
    function getUserStaking(address user)
        public
        view
        returns (
            uint withdrawalWaitTime,
            uint rewardWaitTime,
            uint balance,
            uint pendingReward
        );

    
    
    
    function stake(uint amount)
        external;

    
    
    function withdraw(uint amount)
        external;

    
    
    
    function claim()
        external
        returns (uint claimedAmount);
}

contract UserStakingPool is Claimable, ReentrancyGuard, IUserStakingPool
{
    using ERC20SafeTransfer for address;
    using MathUint          for uint;

    struct Staking {
        uint   balance;        
        uint64 depositedAt;
        uint64 claimedAt;      
    }

    Staking public total;
    mapping (address => Staking) public stakings;

    constructor(address _lrcAddress)
        Claimable()
        public
    {
        require(_lrcAddress != address(0), "ZERO_ADDRESS");
        lrcAddress = _lrcAddress;
    }

    function setProtocolFeeVault(address _protocolFeeVaultAddress)
        external
        nonReentrant
        onlyOwner
    {
        
        protocolFeeVaultAddress = _protocolFeeVaultAddress;
        emit ProtocolFeeVaultChanged(protocolFeeVaultAddress);
    }

    function getTotalStaking()
        public
        view
        returns (uint)
    {
        return total.balance;
    }

    function getUserStaking(address user)
        public
        view
        returns (
            uint withdrawalWaitTime,
            uint rewardWaitTime,
            uint balance,
            uint pendingReward
        )
    {
        withdrawalWaitTime = getUserWithdrawalWaitTime(user);
        rewardWaitTime = getUserClaimWaitTime(user);
        balance = stakings[user].balance;
        (, , pendingReward) = getUserPendingReward(user);
    }

    function stake(uint amount)
        external
        nonReentrant
    {
        require(amount > 0, "ZERO_VALUE");

        
        lrcAddress.safeTransferFromAndVerify(msg.sender, address(this), amount);

        Staking storage user = stakings[msg.sender];

        if (user.balance == 0) {
            numAddresses += 1;
        }

        
        uint balance = user.balance.add(amount);

        user.depositedAt = uint64(
            user.balance
                .mul(user.depositedAt)
                .add(amount.mul(now)) / balance
        );

        user.claimedAt = uint64(
            user.balance
                .mul(user.claimedAt)
                .add(amount.mul(now)) / balance
        );

        user.balance = balance;

        
        balance = total.balance.add(amount);

        total.claimedAt = uint64(
            total.balance
                .mul(total.claimedAt)
                .add(amount.mul(now)) / balance
        );

        total.balance = balance;

        emit LRCStaked(msg.sender, amount);
    }

    function withdraw(uint amount)
        external
        nonReentrant
    {
        require(getUserWithdrawalWaitTime(msg.sender) == 0, "NEED_TO_WAIT");

        
        if (protocolFeeVaultAddress != address(0) &&
            getUserClaimWaitTime(msg.sender) == 0) {
            claimReward();
        }

        Staking storage user = stakings[msg.sender];

        uint _amount = (amount == 0 || amount > user.balance) ? user.balance : amount;
        require(_amount > 0, "ZERO_BALANCE");

        total.balance = total.balance.sub(_amount);
        user.balance = user.balance.sub(_amount);

        if (user.balance == 0) {
            numAddresses -= 1;
            delete stakings[msg.sender];
        }

        
        lrcAddress.safeTransferAndVerify(msg.sender, _amount);

        emit LRCWithdrawn(msg.sender, _amount);
    }

    function claim()
        external
        nonReentrant
        returns (uint claimedAmount)
    {
        return claimReward();
    }

    

    function claimReward()
        private
        returns (uint claimedAmount)
    {
        require(protocolFeeVaultAddress != address(0), "ZERO_ADDRESS");
        require(getUserClaimWaitTime(msg.sender) == 0, "NEED_TO_WAIT");

        uint totalPoints;
        uint userPoints;

        (totalPoints, userPoints, claimedAmount) = getUserPendingReward(msg.sender);

        if (claimedAmount > 0) {
            IProtocolFeeVault(protocolFeeVaultAddress).claimStakingReward(claimedAmount);

            total.balance = total.balance.add(claimedAmount);

            total.claimedAt = uint64(
                now.sub(totalPoints.sub(userPoints) / total.balance)
            );

            Staking storage user = stakings[msg.sender];
            user.balance = user.balance.add(claimedAmount);
            user.claimedAt = uint64(now);
        }
        emit LRCRewarded(msg.sender, claimedAmount);
    }

    function getUserWithdrawalWaitTime(address user)
        private
        view
        returns (uint)
    {
        uint depositedAt = stakings[user].depositedAt;
        if (depositedAt == 0) {
            return MIN_WITHDRAW_DELAY;
        } else {
            uint time = depositedAt + MIN_WITHDRAW_DELAY;
            return (time <= now) ? 0 : time.sub(now);
        }
    }

    function getUserClaimWaitTime(address user)
        private
        view
        returns (uint)
    {
        uint claimedAt = stakings[user].claimedAt;
        if (claimedAt == 0) {
            return MIN_CLAIM_DELAY;
        } else {
            uint time = stakings[user].claimedAt + MIN_CLAIM_DELAY;
            return (time <= now) ? 0 : time.sub(now);
        }
    }

    function getUserPendingReward(address user)
        private
        view
        returns (
            uint totalPoints,
            uint userPoints,
            uint pendingReward
        )
    {
        Staking storage staking = stakings[user];

        
        totalPoints = total.balance.mul(now.sub(total.claimedAt).add(1));
        userPoints = staking.balance.mul(now.sub(staking.claimedAt));

        
        if (totalPoints < userPoints) {
            userPoints = totalPoints;
        }

        if (protocolFeeVaultAddress != address(0) &&
            totalPoints != 0 &&
            userPoints != 0) {
            (, , , , , , , pendingReward) = IProtocolFeeVault(
                protocolFeeVaultAddress
            ).getProtocolFeeStats();
            pendingReward = pendingReward.mul(userPoints) / totalPoints;
        }
    }
}
