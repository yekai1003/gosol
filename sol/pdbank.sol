pragma solidity^0.5.0;

contract pdbank {
    address public  owner;
    mapping(address=>uint256) public balances;
    uint256 public totalAmount;
    string public bankName;
    constructor(string memory _bankName) public  {
        owner = msg.sender;
        bankName = _bankName;
    }
    function deposit() public payable {
        //do nothing
        totalAmount += msg.value;
        balances[msg.sender] += msg.value;
    }
    function withdraw(uint256 _amount) public payable {
        if(balances[msg.sender] > _amount) {
            balances[msg.sender]  -=  _amount;  
            msg.sender.transfer(_amount);
            totalAmount -= _amount;
        }
    }
}
