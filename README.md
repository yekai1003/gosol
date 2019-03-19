# 说明
本框架是为了自动部署以太坊智能合约，以及自动测试合约代码所用，基于go语言编写。
## 1 以太坊智能合约编译
以太坊智能合约编写使用solidity语言，一般情况下我们会在remix环境下进行编译测试，在线环境相对比较稳定。如果不想用在线环境，那我们就需要自己动手来编译代码，并且进行测试。我们都需要准备哪些工具呢？
- 编译器，solidity语言需要安装solc编译器
- 以太坊节点，部署合约时需要用到

### 1.1 编译器安装和准备

安装solc以及solcjs
```
sudo npm install -g solc solc-cli --save-dev
```
solc可以将sol代码编译为go语言，solcjs可以将sol代码编译为abi。


solc 使用方式如下：
```
ykdeMac-mini:abi yekai$ solc  -h
Usage:
  solc [OPTIONS] [ARGS]

Options: 
      --out-dir [PATH]   Output directory for the compiled contracts (Default is ./contracts)
      --optimise         If present activate the solc optimiser
  -k, --no-color         Omit color from output
      --debug            Show debug information
  -h, --help             Display help and usage details

```

不过本人经过实验，这种solc的版本问题仍然会给编译带来很大困扰，因为我还是推荐安装以太坊官方的solidity库，编译就会到solc可执行程序。


下载源码
```
git clone https://github.com/ethereum/solidity
```
编译源码，注意此处编译的时候需要cmake，如果系统中不存在，可以提前安装一下。
如果需要安装cmake，可以参考:
- for ubuntu
```
sudo apt-get install cmake
```
- for mac-os
```
brew install cmake
```
cmake存在后可以编译源码了
```
cd solidity
mkdir build
cd build
cmake .. && make
```

编译好solc之后需要将它放到$PATH环境变量所对应的某个路径下，本人是直接拷贝到了/usr/local/bin下，当然你如果想建立一个软连接也可以，直接连接的路径是准确的就行！
```
cp solc /usr/local/bin/
```
当然我们本篇文章重点讨论的是go语言的问题，所以go语言的安装环境也需要搞定！


### 1.2 以太坊节点安装和部署

这个网上也有很多文章，我们暂且不展开讨论。也可以到这里来查看[区块链技术百科](http://pdjedu.com/wiki.html)


启动geth，开发时使用开发者模式，相对比较容易一些，开发者模式会提供一个默认账户，这个账户有很多钱，而且使用此账户也无需解锁，但是不要想太多，这个是私链的以太币！
 ```
 geth  --rpc --rpcport "8545" --rpccorsdomain "*" --datadir ./data/ --nodiscover --networkid 18  --rpcapi "db,eth,net,web3,personal" --gasprice 0 --dev --dev.period 1   console 2> 1.log
 ```
 启动的目录很关键，后面要用到，其目录下会有data目录，私钥就存放在其中的一个子目录下。


## 2 自动化生成测试代码

下载自动化测试代码框架，这个是基于go语言编写的，目前尚在不断完善中！

```
cd $GOPATH/src

git clone https://github.com/yekai1003/gosol
```
进入到目录，可以进行测试
```
cd gosol
```

在sol目录下有一个示例智能合约，一个银行存取款业务。

```
pragma solidity^0.5.0;

contract pdbank {
    address public  owner;
    mapping(address=>uint256) public balances;
    uint256 public totalAmount;
    string public bankName;
    //构造函数
    constructor(string memory _bankName) public  {
        owner = msg.sender;
        bankName = _bankName;
    }
    //充值
    function deposit() public payable {
        totalAmount += msg.value;
        balances[msg.sender] += msg.value;
    }
    //提现
    function withdraw(uint256 _amount) public payable {
        if(balances[msg.sender] > _amount) {
            balances[msg.sender]  -=  _amount;  
            msg.sender.transfer(_amount);
            totalAmount -= _amount;
        }
    }
}
```

### 2.1 编译合约代码

本框架可以自动测试合约代码是否有效，集编译、部署、功能测试于一体。

编译本工程
```
go build -i
mkdir build
```
templates/templates.toml这个文件是最初的配置文件

```
[common]
connstr = "http://localhost:8545"  #以太坊地址
keydir = "/Users/yekai/eth/data/keystore"  #私钥所在目录
contractaddr = "0xe98a031548fc8e8580e1d2d67637045a82b5250c" #合约地址，部署后拿到
buildpath = "build" #合约地址，部署后拿到
Codepath  =  "callsol.go"  #生成目标代码路径
TestCodepath  =  "testFunc.go" #生成调用时的目标代码路径
TestConfigpath  =  "config.toml" #生成调用时的目标代码路径


[Deploy]
contractname =  "Pdbank"  #部署合约函数，合约编译成go语言后，可以得到
fromaddr  =  "0x70c53a4c94ccce9ce56effbfcb89b221f986cd41" #部署用的账户地址，需要消耗gas
pass  = "123"  #账户地址对应的密码
AbiFile =  "contracts/pdbank.abi"  #编译后ABI文件所在路径
```
其中连接地址默认私链可以不必改动，keydir是代表私钥存在路径，这个需要修改为自己本机中的位置，还记得geth的启动目录吧，在那里可以找到。
AbiFile这个文件是在编译合约后得到的，所以编译之后可以修改此处！


### 2.2 生成目标测试代码

编译合约代码

```
localhost:gosol yekai$ ./gosol comp
templates init ok
[pdbank.sol] <nil>
pdbank.sol
/usr/local/bin/abigen
run complie go ok!! <nil>
run BuildAbi ok!! <nil>

```
此后可以在contracts看到2个文件

```
localhost:gosol yekai$ ls contracts/
pdbank.abi	pdbank.go
```
如果合约名字不一样，可以修改一下之前的配置文件地方，对abi文件进行替换。

上面只是完成了第一步，编译代码，接着需要生成自动化的测试代码。

```
localhost:gosol yekai$ ./gosol build
templates init ok
deploy config begin......
deploy config ok......
ImplRunCode() begin
ImplRunCode() end
```
可以在build目录得到目标的代码，进入到build目录，可以进行合约代码测试工作了。

```
cd build 
go build -i
```
可以查看一下命令帮助

```
localhost:build yekai$ ./build 
call build config.init
call build config.init ok
1 - help   :./build  1
2 - deploy  :./build  2
3 - test totalAmount:./build  3
4 - test bankName:./build  4
5 - test balances:./build  5
6 - test withdraw:./build  6
7 - test owner:./build  7
8 - test deposit:./build  8
```
之前的合约一共有这么多个函数，所以自动化生成这些测试命令，需要先确认一下配置文件config.toml，这里面的fromaddr和pass是部署合约使用，必须是正确的。

```
[common]
connstr = "http://localhost:8545"         #以太坊连接地址
keydir = "/Users/yekai/eth/data/keystore"           #以太坊私钥目录
ContractAddr = "0xe98a031548fc8e8580e1d2d67637045a82b5250c"  #合约地址，部署后得到
Buildpath  =  "build"   #编译目标目录，默认为build
Codepath  =  "callsol.go"     #编译后的代码文件名
TestCodepath  =  "testFunc.go"  #自动生成的测试用代码文件名 
TestConfigpath  =  "config.toml" #自动生成的配置文件名

[Deploy]
func =  "DeployPdbank"  #部署函数入口
fromaddr  =  "0x70c53a4c94ccce9ce56effbfcb89b221f986cd41"   #部署用账户
pass  = "123"            #部署账户密码
TestAddr = "0x70c53a4c94ccce9ce56effbfcb89b221f986cd41"     #测试用账户
TestPass = "123"     #测试用账户密码
```


部署合约试试，可以得到合约地址和收据hash
```
localhost:build yekai$ ./build 2
call build config.init
call build config.init ok
0 0x70a4124d3fa01c05c608ae001f135eae06d500295a6907cde74563a2eeba5e0a 36
0xf375DB17BD4CFF8ff9E16D6665DbF75dED3CB567
```

合约地址拿到，别忘了去修改配置文件的信息,TestAddr与TestPass也可以一起设置下，否则后面测试的账户是空的

查看一下银行的名称：可以看到yekai
```
localhost:build yekai$ ./build 4
call build config.init
call build config.init ok
yekai <nil>
```

充值试试，这个代码callsol.go需要改一下 ，因为默认充值金额为0.
```
func CallDeposit(addr, pass string) (*types.Transaction, error) {

	instance, err := contracts.NewPdbank(common.HexToAddress(Config.Common.ContractAddr), testclient)
	if err != nil {
		fmt.Println("failed to get contract instance", err)
		return nil, err
	}
	auth, err := MakeAuth(addr, pass)
	if err != nil {
		fmt.Println("failed to makeAuth", err)
		return nil, err
	}
	auth.Value = big.NewInt(100010)
	ts,err := instance.Deposit(auth)
	if err != nil {
		fmt.Println("failed to call ", err)
		return nil, err
	}
	fmt.Println(ts.ChainId(), ts.Hash().Hex(), ts.Nonce())
	return ts , err
}
```


再重新编译一下，调用充值Deposit
```
go build -i

localhost:build yekai$ ./build 8
call build config.init
call build config.init ok
0 0x3083d6000959762c4c8911e4aaca8de9f8700140dd115ca96b1233d7bccd399d 38
localhost:build yekai$ ./build 5
call build config.init
call build config.init ok
100010 <nil>

```
