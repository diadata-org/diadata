pragma solidity 0.4.24;
import "zos-lib/contracts/Initializable.sol";
import "openzeppelin-eth/contracts/token/ERC20/ERC20.sol";
import "openzeppelin-eth/contracts/token/ERC20/ERC20Mintable.sol";
import "openzeppelin-eth/contracts/ownership/Ownable.sol";

contract DIAToken is Initializable, ERC20, ERC20Mintable, Ownable {
	function initialize(address _owner) public initializer()  {
		Ownable.initialize(_owner);
		ERC20Mintable.initialize(_owner);
	}
}
