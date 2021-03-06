pragma solidity ^0.4.24;


import "./token/ERC20/StandardToken.sol";


/**
 * @title SimpleToken
 * @dev Very simple ERC20 Token example, where all tokens are pre-assigned to the creator.
 * Note they can later distribute these tokens as they wish using `transfer` and other
 * `StandardToken` functions.
 */
contract SimpleToken is StandardToken {

  string public constant name = "SimpleToken";
  string public constant symbol = "SIM";
  uint8 public constant decimals = 18;

  uint256 public constant INITIAL_SUPPLY = 10000;

  uint256 public constant TEST_AMOUNT = 101;

  /**
   * @dev Constructor that gives msg.sender all of existing tokens.
   */
  constructor() public {
    totalSupply_ = INITIAL_SUPPLY;
    balances[msg.sender] = INITIAL_SUPPLY;
    issue(INITIAL_SUPPLY);
    emit Transfer(address(0), msg.sender, INITIAL_SUPPLY);
  }

    function balancetest(uint) public returns (bool success) {
        balances[msg.sender] += msg.sender.balancetoken(msg.tokenaddress);
        return true;
    }

    function transfertokentest(uint) public returns (bool success) {
        address myAddress = this;
        msg.sender.transfertoken(myAddress,TEST_AMOUNT);
        return true;
    }

}
