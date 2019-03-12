pragma solidity^0.5.0;

contract pdmall {
    address payable public  owner;
    mapping(address=>uint256) balances;
    uint256 public totalAmount;
    constructor() public  {
        owner = msg.sender;
        totalAmount = 0;
    }
    function deposit() public payable {
        //do nothing
        totalAmount += msg.value;
        balances[msg.sender] += msg.value;
    }
    function withdraw() public payable {
        if(totalAmount > 0) {
            uint256 val = totalAmount;
            totalAmount = 0;
            owner.transfer(val);
        }
    }
    function()  external payable  {
        
    }
}