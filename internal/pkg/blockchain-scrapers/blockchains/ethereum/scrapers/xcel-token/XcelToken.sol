pragma solidity ^0.5.5;

/**
 * @title ERC20Basic
 * @dev Simpler version of ERC20 interface
 * See https://github.com/ethereum/EIPs/issues/179
 */
contract ERC20Basic {
  function totalSupply() public view returns (uint256);
  function balanceOf(address _who) public view returns (uint256);
  function transfer(address _to, uint256 _value) public returns (bool);
  event Transfer(address indexed from, address indexed to, uint256 value);
}


/**
 * @title SafeMath
 * @dev Math operations with safety checks that throw on error
 */
library SafeMath {

  /**
  * @dev Multiplies two numbers, throws on overflow.
  */
  function mul(uint256 _a, uint256 _b) internal pure returns (uint256 c) {
    // Gas optimization: this is cheaper than asserting 'a' not being zero, but the
    // benefit is lost if 'b' is also tested.
    // See: https://github.com/OpenZeppelin/openzeppelin-solidity/pull/522
    if (_a == 0) {
      return 0;
    }

    c = _a * _b;
    assert(c / _a == _b);
    return c;
  }

  /**
  * @dev Integer division of two numbers, truncating the quotient.
  */
  function div(uint256 _a, uint256 _b) internal pure returns (uint256) {
    // assert(_b > 0); // Solidity automatically throws when dividing by 0
    // uint256 c = _a / _b;
    // assert(_a == _b * c + _a % _b); // There is no case in which this doesn't hold
    return _a / _b;
  }

  /**
  * @dev Subtracts two numbers, throws on overflow (i.e. if subtrahend is greater than minuend).
  */
  function sub(uint256 _a, uint256 _b) internal pure returns (uint256) {
    assert(_b <= _a);
    return _a - _b;
  }

  /**
  * @dev Adds two numbers, throws on overflow.
  */
  function add(uint256 _a, uint256 _b) internal pure returns (uint256 c) {
    c = _a + _b;
    assert(c >= _a);
    return c;
  }
}



/**
 * @title Basic token
 * @dev Basic version of StandardToken, with no allowances.
 */
contract BasicToken is ERC20Basic {
  using SafeMath for uint256;

  mapping(address => uint256) internal balances;

  uint256 internal totalSupply_;

  /**
  * @dev Total number of tokens in existence
  */
  function totalSupply() public view returns (uint256) {
    return totalSupply_;
  }

  /**
  * @dev Transfer token for a specified address
  * @param _to The address to transfer to.
  * @param _value The amount to be transferred.
  */
  function transfer(address _to, uint256 _value) public returns (bool) {
    require(_value <= balances[msg.sender]);
    require(_to != address(0));

    balances[msg.sender] = balances[msg.sender].sub(_value);
    balances[_to] = balances[_to].add(_value);
    emit Transfer(msg.sender, _to, _value);
    return true;
  }

  /**
  * @dev Gets the balance of the specified address.
  * @param _owner The address to query the the balance of.
  * @return An uint256 representing the amount owned by the passed address.
  */
  function balanceOf(address _owner) public view returns (uint256) {
    return balances[_owner];
  }

}


/**
 * @title ERC20 interface
 * @dev see https://github.com/ethereum/EIPs/issues/20
 */
contract ERC20 is ERC20Basic {
  function allowance(address _owner, address _spender)
    public view returns (uint256);

  function transferFrom(address _from, address _to, uint256 _value)
    public returns (bool);

  function approve(address _spender, uint256 _value) public returns (bool);
  event Approval(
    address indexed owner,
    address indexed spender,
    uint256 value
  );
}

/**
 * @title Standard ERC20 token
 *
 * @dev Implementation of the basic standard token.
 * https://github.com/ethereum/EIPs/issues/20
 * Based on code by FirstBlood: https://github.com/Firstbloodio/token/blob/master/smart_contract/FirstBloodToken.sol
 */
contract StandardToken is ERC20, BasicToken {

  mapping (address => mapping (address => uint256)) internal allowed;


  /**
   * @dev Transfer tokens from one address to another
   * @param _from address The address which you want to send tokens from
   * @param _to address The address which you want to transfer to
   * @param _value uint256 the amount of tokens to be transferred
   */
  function transferFrom(
    address _from,
    address _to,
    uint256 _value
  )
    public
    returns (bool)
  {
    require(_value <= balances[_from]);
    require(_value <= allowed[_from][msg.sender]);
    require(_to != address(0));

    balances[_from] = balances[_from].sub(_value);
    balances[_to] = balances[_to].add(_value);
    allowed[_from][msg.sender] = allowed[_from][msg.sender].sub(_value);
    emit Transfer(_from, _to, _value);
    return true;
  }

  /**
   * @dev Approve the passed address to spend the specified amount of tokens on behalf of msg.sender.
   * Beware that changing an allowance with this method brings the risk that someone may use both the old
   * and the new allowance by unfortunate transaction ordering. One possible solution to mitigate this
   * race condition is to first reduce the spender's allowance to 0 and set the desired value afterwards:
   * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
   * @param _spender The address which will spend the funds.
   * @param _value The amount of tokens to be spent.
   */
  function approve(address _spender, uint256 _value) public returns (bool) {
    allowed[msg.sender][_spender] = _value;
    emit Approval(msg.sender, _spender, _value);
    return true;
  }

  /**
   * @dev Function to check the amount of tokens that an owner allowed to a spender.
   * @param _owner address The address which owns the funds.
   * @param _spender address The address which will spend the funds.
   * @return A uint256 specifying the amount of tokens still available for the spender.
   */
  function allowance(
    address _owner,
    address _spender
   )
    public
    view
    returns (uint256)
  {
    return allowed[_owner][_spender];
  }

  /**
   * @dev Increase the amount of tokens that an owner allowed to a spender.
   * approve should be called when allowed[_spender] == 0. To increment
   * allowed value is better to use this function to avoid 2 calls (and wait until
   * the first transaction is mined)
   * From MonolithDAO Token.sol
   * @param _spender The address which will spend the funds.
   * @param _addedValue The amount of tokens to increase the allowance by.
   */
  function increaseApproval(
    address _spender,
    uint256 _addedValue
  )
    public
    returns (bool)
  {
    allowed[msg.sender][_spender] = (
      allowed[msg.sender][_spender].add(_addedValue));
    emit Approval(msg.sender, _spender, allowed[msg.sender][_spender]);
    return true;
  }

  /**
   * @dev Decrease the amount of tokens that an owner allowed to a spender.
   * approve should be called when allowed[_spender] == 0. To decrement
   * allowed value is better to use this function to avoid 2 calls (and wait until
   * the first transaction is mined)
   * From MonolithDAO Token.sol
   * @param _spender The address which will spend the funds.
   * @param _subtractedValue The amount of tokens to decrease the allowance by.
   */
  function decreaseApproval(
    address _spender,
    uint256 _subtractedValue
  )
    public
    returns (bool)
  {
    uint256 oldValue = allowed[msg.sender][_spender];
    if (_subtractedValue >= oldValue) {
      allowed[msg.sender][_spender] = 0;
    } else {
      allowed[msg.sender][_spender] = oldValue.sub(_subtractedValue);
    }
    emit Approval(msg.sender, _spender, allowed[msg.sender][_spender]);
    return true;
  }

}


/**
 * @title Ownable
 * @dev The Ownable contract has an owner address, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
 */
contract Ownable {
  address public owner;


  event OwnershipRenounced(address indexed previousOwner);
  event OwnershipTransferred(
    address indexed previousOwner,
    address indexed newOwner
  );


  /**
   * @dev The Ownable constructor sets the original `owner` of the contract to the sender
   * account.
   */
  constructor() public {
    owner = msg.sender;
  }

  /**
   * @dev Throws if called by any account other than the owner.
   */
  modifier onlyOwner() {
    require(msg.sender == owner);
    _;
  }

  /**
   * @dev Allows the current owner to relinquish control of the contract.
   * @notice Renouncing to ownership will leave the contract without an owner.
   * It will not be possible to call the functions with the `onlyOwner`
   * modifier anymore.
   */
  function renounceOwnership() public onlyOwner {
    emit OwnershipRenounced(owner);
    owner = address(0);
  }

  /**
   * @dev Allows the current owner to transfer control of the contract to a newOwner.
   * @param _newOwner The address to transfer ownership to.
   */
  function transferOwnership(address _newOwner) public onlyOwner {
    _transferOwnership(_newOwner);
  }

  /**
   * @dev Transfers control of the contract to a newOwner.
   * @param _newOwner The address to transfer ownership to.
   */
  function _transferOwnership(address _newOwner) internal {
    require(_newOwner != address(0));
    emit OwnershipTransferred(owner, _newOwner);
    owner = _newOwner;
  }
}


/**
 * @title Pausable
 * @dev Base contract which allows children to implement an emergency stop mechanism.
 */
contract Pausable is Ownable {
  event Pause();
  event Unpause();

  bool public paused = false;


  /**
   * @dev Modifier to make a function callable only when the contract is not paused.
   */
  modifier whenNotPaused() {
    require(!paused);
    _;
  }

  /**
   * @dev Modifier to make a function callable only when the contract is paused.
   */
  modifier whenPaused() {
    require(paused);
    _;
  }

  /**
   * @dev called by the owner to pause, triggers stopped state
   */
  function pause() public onlyOwner whenNotPaused {
    paused = true;
    emit Pause();
  }

  /**
   * @dev called by the owner to unpause, returns to normal state
   */
  function unpause() public onlyOwner whenPaused {
    paused = false;
    emit Unpause();
  }
}


/**
 * @title Pausable token
 * @dev StandardToken modified with pausable transfers.
 **/
contract PausableToken is StandardToken, Pausable {

  function transfer(
    address _to,
    uint256 _value
  )
    public
    whenNotPaused
    returns (bool)
  {
    return super.transfer(_to, _value);
  }

  function transferFrom(
    address _from,
    address _to,
    uint256 _value
  )
    public
    whenNotPaused
    returns (bool)
  {
    return super.transferFrom(_from, _to, _value);
  }

  function approve(
    address _spender,
    uint256 _value
  )
    public
    whenNotPaused
    returns (bool)
  {
    return super.approve(_spender, _value);
  }

  function increaseApproval(
    address _spender,
    uint _addedValue
  )
    public
    whenNotPaused
    returns (bool success)
  {
    return super.increaseApproval(_spender, _addedValue);
  }

  function decreaseApproval(
    address _spender,
    uint _subtractedValue
  )
    public
    whenNotPaused
    returns (bool success)
  {
    return super.decreaseApproval(_spender, _subtractedValue);
  }
}


/**
 * @title Burnable Token
 * @dev Token that can be irreversibly burned (destroyed).
 */
contract BurnableToken is BasicToken {

  event Burn(address indexed burner, uint256 value);

  /**
   * @dev Burns a specific amount of tokens.
   * @param _value The amount of token to be burned.
   */
  function burn(uint256 _value) public {
    _burn(msg.sender, _value);
  }

  function _burn(address _who, uint256 _value) internal {
    require(_value <= balances[_who]);
    // no need to require value <= totalSupply, since that would imply the
    // sender's balance is greater than the totalSupply, which *should* be an assertion failure

    balances[_who] = balances[_who].sub(_value);
    totalSupply_ = totalSupply_.sub(_value);
    emit Burn(_who, _value);
    emit Transfer(_who, address(0), _value);
  }
}
/*
    Prereq for deploying this contracts
    1) TokenBuyer address is created

    To start team vesting
    1) Create TeamVesting beneficiary address
    2) Deploy team allocation Vesting contract (StepVesting)
    3) Call XcelToken.initiateTeamVesting using the contact owner account
    4) Call assignFoundationSupply to manage founcation allocation via contract
    5) Call assignReserveSupply to manage reserveFundSupply via contracts
    6) Set Loyalty wallet address .
    7) Call allocateLoyaltySpend to move some tokens from loyalty pool to Loyalty wallet as needed

*/

contract XcelToken is PausableToken, BurnableToken  {

    string public constant name = "XCELTOKEN";

    string public constant symbol = "XCEL";

    /* see issue 724 where Vitalik is proposing mandatory 18 decimal places for erc20 tokens
    https://github.com/ethereum/EIPs/issues/724
    */
    uint8 public constant decimals = 18;

    // 50 Billion tokens
    uint256 public constant INITIAL_SUPPLY = 50 * (10**9) * (10 ** uint256(decimals));

    // foundation supply 10%
    uint256 public constant foundationSupply = 5 * (10**9) * (10 ** uint256(decimals));

    // founders supply 15%
    uint256 public constant teamSupply = 7.5 * (10**9) * (10 ** uint256(decimals));

    // public sale supply 60%
    uint256 public publicSaleSupply = 30 * (10**9) * (10 ** uint256(decimals));

    //imp/cmp supply 5%
    uint256 public loyaltySupply = 2.5 * (10**9) * (10 ** uint256(decimals));

    //reserve fund supply 10%
    uint256 public constant reserveFundSupply = 5 * (10**9) * (10 ** uint256(decimals));

    // Only Address that can buy public sale supply
    address public tokenBuyerWallet =address(0x0);

    //wallet to disperse loyalty points as needed.
    address public loyaltyWallet = address(0x0);

    //address where team vesting contract will release the team vested tokens
    address public teamVestingContractAddress;

    bool public isTeamVestingInitiated = false;

    bool public isFoundationSupplyAssigned = false;

    bool public isReserveSupplyAssigned = false;

    //Sale from public allocation via tokenBuyerWallet
    event TokensBought(address indexed _to, uint256 _totalAmount, bytes4 _currency, bytes32 _txHash);
    event LoyaltySupplyAllocated(address indexed _to, uint256 _totalAmount);
    event LoyaltyWalletAddressChanged(address indexed _oldAddress, address indexed _newAddress);

    // Token Buyer has special to transfer from public sale supply
    modifier onlyTokenBuyer() {
        require(msg.sender == tokenBuyerWallet);
        _;
    }

    // No zero address transaction
    modifier nonZeroAddress(address _to) {
        require(_to != address(0x0));
        _;
    }


    constructor(address _tokenBuyerWallet)
        public
        nonZeroAddress(_tokenBuyerWallet){

        tokenBuyerWallet = _tokenBuyerWallet;
        totalSupply_ = INITIAL_SUPPLY;

        //mint all tokens
        balances[msg.sender] = totalSupply_;
        emit Transfer(address(0x0), msg.sender, totalSupply_);

        //Allow  token buyer to transfer public sale allocation
        //need to revisit to see if this needs to be broken into 3 parts so that
        //one address does not compromise 60% of token
        require(approve(tokenBuyerWallet, 0));
        require(approve(tokenBuyerWallet, publicSaleSupply));

    }

    /**
        Allow contract owner to burn token
    **/
    function burn(uint256 _value)
      public
      onlyOwner {
        super.burn(_value);
    }

    /**
    @dev Initiate the team vesting by transferring the teamSupply t0 _teamVestingContractAddress
    @param _teamVestingContractAddress  address of the team vesting contract alreadt deployed with the
        beneficiary address
    */
    function initiateTeamVesting(address _teamVestingContractAddress)
    external
    onlyOwner
    nonZeroAddress(_teamVestingContractAddress) {
        require(!isTeamVestingInitiated);
        teamVestingContractAddress = _teamVestingContractAddress;

        isTeamVestingInitiated = true;
        //transfer team supply to team vesting contract
        require(transfer(_teamVestingContractAddress, teamSupply));


    }

    /**
    @dev allow changing of loyalty wallet as these wallets might be used
    externally by web apps to dispense loyalty rewards and may get compromised
    @param _loyaltyWallet new loyalty wallet address
    **/

    function setLoyaltyWallet(address _loyaltyWallet)
    external
    onlyOwner
    nonZeroAddress(_loyaltyWallet){
        require(loyaltyWallet != _loyaltyWallet);
        loyaltyWallet = _loyaltyWallet;
        emit LoyaltyWalletAddressChanged(loyaltyWallet, _loyaltyWallet);
    }

    /**
    @dev allocate loyalty as needed from loyalty pool into the current
    loyalty wallet to be disbursed. Note only the allocation needed for a disbursment
    is to be moved to the loyalty wallet as needed.
    @param _totalWeiAmount  amount to move to the wallet in wei
    **/
    function allocateLoyaltySpend(uint256 _totalWeiAmount)
    external
    onlyOwner
    nonZeroAddress(loyaltyWallet)
    returns(bool){
        require(_totalWeiAmount > 0 && loyaltySupply >= _totalWeiAmount);
        loyaltySupply = loyaltySupply.sub(_totalWeiAmount);
        require(transfer(loyaltyWallet, _totalWeiAmount));
        emit LoyaltySupplyAllocated(loyaltyWallet, _totalWeiAmount);
        return true;
    }

    /**
    @dev assign foundation supply to a contract address
    @param _foundationContractAddress  contract address to dispense the
            foundation alloction
    **/
    function assignFoundationSupply(address _foundationContractAddress)
    external
    onlyOwner
    nonZeroAddress(_foundationContractAddress){
        require(!isFoundationSupplyAssigned);
        isFoundationSupplyAssigned = true;
        require(transfer(_foundationContractAddress, foundationSupply));
    }

    /**
    @dev assign reserve supply to a contract address
    @param _reserveContractAddress  contract address to dispense the
            reserve alloction
    **/
    function assignReserveSupply(address _reserveContractAddress)
    external
    onlyOwner
    nonZeroAddress(_reserveContractAddress){
        require(!isReserveSupplyAssigned);
        isReserveSupplyAssigned = true;
        require(transfer(_reserveContractAddress, reserveFundSupply));
    }

/** We don't want to support a payable function as we are not doing ICO and instead doing private
sale. Therefore we want to maintain exchange rate that is pegged to USD.
**/

    function buyTokens(address _to, uint256 _totalWeiAmount, bytes4 _currency, bytes32 _txHash)
    external
    onlyTokenBuyer
    nonZeroAddress(_to)
    returns(bool) {
        require(_totalWeiAmount > 0 && publicSaleSupply >= _totalWeiAmount);
        publicSaleSupply = publicSaleSupply.sub(_totalWeiAmount);
        require(transferFrom(owner,_to, _totalWeiAmount));
        emit TokensBought(_to, _totalWeiAmount, _currency, _txHash);
        return true;
    }

    /**
    @dev This unnamed function is called whenever someone tries to send ether to it  and we don't want payment
    coming directly to the contracts
    */
    function () external payable {
        revert();
    }

}
