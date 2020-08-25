pragma solidity ^0.6.0;

  library SafeMath {
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
      assert(b <= a);
      return a - b;
    }
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
      uint256 c = a + b;
      assert(c >= a);
      return c;
    }
  }

  contract Neotoken {
    using SafeMath for uint256;
    string public constant name = "TUSDTEST";
    string public constant symbol = "USDT";
    uint256 public constant decimals = 18;
    uint256 public constant totalSupply = 100000000000000000000000000;
    //address private founder = 0x0;
    uint256 private distributed = 0;
    mapping (address => uint256) private balances;
    mapping (address => mapping (address => uint256)) private allowed;
    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
    constructor() public {
      balances[msg.sender] = totalSupply;
    }

    fallback () payable external {}
    receive () payable external {}

    function balanceOf(address _owner) public view returns (uint256 balance) {
      return balances[_owner];
    }
    function transfer(address _to, uint256 _value) public returns (bool success) {
     // require (_to != 0x0, "");
      require((balances[msg.sender] >= _value), "");
      balances[msg.sender] = balances[msg.sender].sub(_value);
      balances[_to] = balances[_to].add(_value);
      emit Transfer(msg.sender, _to, _value);
      return true;
    }
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
      require (_to != address(0), "");

      require(balances[_from] >= _value && allowed[_from][msg.sender] >= _value, "not eng aaa");
      allowed[_from][msg.sender] = allowed[_from][msg.sender].sub(_value);
      balances[_from] = balances[_from].sub(_value);
      balances[_to] = balances[_to].add(_value);
      emit Transfer(_from, _to, _value);
      return true;
    }
    function approve(address _spender, uint256 _value) public returns (bool success) {
      allowed[msg.sender][_spender] = _value;
      emit Approval(msg.sender, _spender, _value);
      return true;
    }
    function allowance(address _owner, address _spender) public view returns (uint256 remaining) {
      return allowed[_owner][_spender];
    }
  }