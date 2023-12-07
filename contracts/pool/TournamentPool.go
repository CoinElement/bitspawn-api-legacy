// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pool

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BitspawnABI is the input ABI used to generate the binding from.
const BitspawnABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wat\",\"type\":\"bool\"}],\"name\":\"trust\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name_\",\"type\":\"bytes32\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"trusted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"pull\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wat\",\"type\":\"bool\"}],\"name\":\"Trust\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BitspawnBin is the compiled bytecode used for deploying new contracts.
const BitspawnBin = `0x608060405260126006557f5350574e000000000000000000000000000000000000000000000000000000006007557f426974737061776e20546f6b656e00000000000000000000000000000000000060085534801561005d57600080fd5b5033600081815260016020526040808220829055818055600480546001600160a01b03191684179055517fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed949190a2610f9f806100ba6000396000f3fe6080604052600436106101815760003560e01c80637a9e5e4b116100d1578063a9059cbb1161008a578063be9a655511610064578063be9a655514610580578063bf7e214f14610588578063dd62ed3e1461059d578063f2d5d56b146105d857610181565b8063a9059cbb146104cb578063b753a98c14610504578063bb35783b1461053d57610181565b80637a9e5e4b146103b45780637aa3295b146103e75780638da5cb5b1461042257806395d89b41146104535780639dc29fac14610468578063a0712d68146104a157610181565b806323b872dd1161013e57806342966c681161011857806342966c68146103185780635ac801fe1461034257806370a082311461036c57806375f12b211461039f57610181565b806323b872dd14610287578063313ce567146102ca57806340c10f19146102df57610181565b806306262f1b1461018657806306fdde03146101c357806307da68f5146101ea578063095ea7b3146101f257806313af40351461023f57806318160ddd14610272575b600080fd5b34801561019257600080fd5b506101c1600480360360408110156101a957600080fd5b506001600160a01b0381351690602001351515610611565b005b3480156101cf57600080fd5b506101d8610696565b60408051918252519081900360200190f35b6101c161069c565b3480156101fe57600080fd5b5061022b6004803603604081101561021557600080fd5b506001600160a01b038135169060200135610736565b604080519115158252519081900360200190f35b34801561024b57600080fd5b506101c16004803603602081101561026257600080fd5b50356001600160a01b0316610763565b34801561027e57600080fd5b506101d86107d2565b34801561029357600080fd5b5061022b600480360360608110156102aa57600080fd5b506001600160a01b038135811691602081013590911690604001356107d8565b3480156102d657600080fd5b506101d861093a565b3480156102eb57600080fd5b506101c16004803603604081101561030257600080fd5b506001600160a01b038135169060200135610940565b34801561032457600080fd5b506101c16004803603602081101561033b57600080fd5b5035610a47565b34801561034e57600080fd5b506101c16004803603602081101561036557600080fd5b5035610a54565b34801561037857600080fd5b506101d86004803603602081101561038f57600080fd5b50356001600160a01b0316610a78565b3480156103ab57600080fd5b5061022b610a93565b3480156103c057600080fd5b506101c1600480360360208110156103d757600080fd5b50356001600160a01b0316610aa3565b3480156103f357600080fd5b5061022b6004803603604081101561040a57600080fd5b506001600160a01b0381358116916020013516610b12565b34801561042e57600080fd5b50610437610b40565b604080516001600160a01b039092168252519081900360200190f35b34801561045f57600080fd5b506101d8610b4f565b34801561047457600080fd5b506101c16004803603604081101561048b57600080fd5b506001600160a01b038135169060200135610b55565b3480156104ad57600080fd5b506101c1600480360360208110156104c457600080fd5b5035610c5c565b3480156104d757600080fd5b5061022b600480360360408110156104ee57600080fd5b506001600160a01b038135169060200135610c66565b34801561051057600080fd5b506101c16004803603604081101561052757600080fd5b506001600160a01b038135169060200135610c73565b34801561054957600080fd5b506101c16004803603606081101561056057600080fd5b506001600160a01b03813581169160208101359091169060400135610c83565b6101c1610c94565b34801561059457600080fd5b50610437610d28565b3480156105a957600080fd5b506101d8600480360360408110156105c057600080fd5b506001600160a01b0381358116916020013516610d37565b3480156105e457600080fd5b506101c1600480360360408110156105fb57600080fd5b506001600160a01b038135169060200135610d62565b600454600160a01b900460ff161561062857600080fd5b3360008181526005602090815260408083206001600160a01b03871680855290835292819020805460ff1916861515908117909155815190815290519293927ff184148577730b253ecb4339c543a564af420f3d32ed12a1c62ae83d67d65fe3929181900390910190a35050565b60085481565b6106b2336000356001600160e01b031916610d6d565b6106bb57600080fd5b604080513480825260208201838152369383018490526004359360243593849386933393600080356001600160e01b03191694909260608201848480828437600083820152604051601f909101601f1916909201829003965090945050505050a4505060048054600160a01b60ff021916600160a01b179055565b600454600090600160a01b900460ff161561075057600080fd5b61075a8383610e57565b90505b92915050565b610779336000356001600160e01b031916610d6d565b61078257600080fd5b600480546001600160a01b0319166001600160a01b0383811691909117918290556040519116907fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9490600090a250565b60005490565b600454600090600160a01b900460ff16156107f257600080fd5b6001600160a01b038416331480159061082f57506001600160a01b038416600090815260056020908152604080832033845290915290205460ff16155b15610887576001600160a01b03841660009081526002602090815260408083203384529091529020546108629083610ebd565b6001600160a01b03851660009081526002602090815260408083203384529091529020555b6001600160a01b0384166000908152600160205260409020546108aa9083610ebd565b6001600160a01b0380861660009081526001602052604080822093909355908516815220546108d99083610f18565b6001600160a01b0380851660008181526001602090815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a35060019392505050565b60065481565b610956336000356001600160e01b031916610d6d565b61095f57600080fd5b600454600160a01b900460ff161561097657600080fd5b6001600160a01b0382166000908152600160205260409020546109999082610f18565b6001600160a01b038316600090815260016020526040812091909155546109c09082610f18565b6000556040805182815290516001600160a01b038416917f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885919081900360200190a26040805182815290516001600160a01b038416916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b610a513382610b55565b50565b610a6a336000356001600160e01b031916610d6d565b610a7357600080fd5b600855565b6001600160a01b031660009081526001602052604090205490565b600454600160a01b900460ff1681565b610ab9336000356001600160e01b031916610d6d565b610ac257600080fd5b600380546001600160a01b0319166001600160a01b0383811691909117918290556040519116907f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada490600090a250565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6004546001600160a01b031681565b60075481565b610b6b336000356001600160e01b031916610d6d565b610b7457600080fd5b600454600160a01b900460ff1615610b8b57600080fd5b6001600160a01b038216600090815260016020526040902054610bae9082610ebd565b6001600160a01b03831660009081526001602052604081209190915554610bd59082610ebd565b6000556040805182815290516001600160a01b038416917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a26040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b610a513382610940565b600061075a3384846107d8565b610c7e3383836107d8565b505050565b610c8e8383836107d8565b50505050565b610caa336000356001600160e01b031916610d6d565b610cb357600080fd5b604080513480825260208201838152369383018490526004359360243593849386933393600080356001600160e01b03191694909260608201848480828437600083820152604051601f909101601f1916909201829003965090945050505050a4505060048054600160a01b60ff0219169055565b6003546001600160a01b031681565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b610c7e8233836107d8565b60006001600160a01b038316301415610d885750600161075d565b6004546001600160a01b0384811691161415610da65750600161075d565b6003546001600160a01b0316610dbe5750600061075d565b60035460408051600160e01b63b70096130281526001600160a01b0386811660048301523060248301526001600160e01b0319861660448301529151919092169163b7009613916064808301926020929190829003018186803b158015610e2457600080fd5b505afa158015610e38573d6000803e3d6000fd5b505050506040513d6020811015610e4e57600080fd5b5051905061075d565b3360008181526002602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b8082038281111561075d5760408051600160e51b62461bcd02815260206004820152601560248201527f64732d6d6174682d7375622d756e646572666c6f770000000000000000000000604482015290519081900360640190fd5b8082018281101561075d5760408051600160e51b62461bcd02815260206004820152601460248201527f64732d6d6174682d6164642d6f766572666c6f77000000000000000000000000604482015290519081900360640190fdfea165627a7a7230582078a69ba806d3b7bcd37a09a2756e4e79c1c21369312a6366c1bfebcacf2ccfeb0029`

// DeployBitspawn deploys a new Ethereum contract, binding an instance of Bitspawn to it.
func DeployBitspawn(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bitspawn, error) {
	parsed, err := abi.JSON(strings.NewReader(BitspawnABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BitspawnBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bitspawn{BitspawnCaller: BitspawnCaller{contract: contract}, BitspawnTransactor: BitspawnTransactor{contract: contract}, BitspawnFilterer: BitspawnFilterer{contract: contract}}, nil
}

// Bitspawn is an auto generated Go binding around an Ethereum contract.
type Bitspawn struct {
	BitspawnCaller     // Read-only binding to the contract
	BitspawnTransactor // Write-only binding to the contract
	BitspawnFilterer   // Log filterer for contract events
}

// BitspawnCaller is an auto generated read-only Go binding around an Ethereum contract.
type BitspawnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitspawnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BitspawnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitspawnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BitspawnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BitspawnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BitspawnSession struct {
	Contract     *Bitspawn         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BitspawnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BitspawnCallerSession struct {
	Contract *BitspawnCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BitspawnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BitspawnTransactorSession struct {
	Contract     *BitspawnTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BitspawnRaw is an auto generated low-level Go binding around an Ethereum contract.
type BitspawnRaw struct {
	Contract *Bitspawn // Generic contract binding to access the raw methods on
}

// BitspawnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BitspawnCallerRaw struct {
	Contract *BitspawnCaller // Generic read-only contract binding to access the raw methods on
}

// BitspawnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BitspawnTransactorRaw struct {
	Contract *BitspawnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBitspawn creates a new instance of Bitspawn, bound to a specific deployed contract.
func NewBitspawn(address common.Address, backend bind.ContractBackend) (*Bitspawn, error) {
	contract, err := bindBitspawn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bitspawn{BitspawnCaller: BitspawnCaller{contract: contract}, BitspawnTransactor: BitspawnTransactor{contract: contract}, BitspawnFilterer: BitspawnFilterer{contract: contract}}, nil
}

// NewBitspawnCaller creates a new read-only instance of Bitspawn, bound to a specific deployed contract.
func NewBitspawnCaller(address common.Address, caller bind.ContractCaller) (*BitspawnCaller, error) {
	contract, err := bindBitspawn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BitspawnCaller{contract: contract}, nil
}

// NewBitspawnTransactor creates a new write-only instance of Bitspawn, bound to a specific deployed contract.
func NewBitspawnTransactor(address common.Address, transactor bind.ContractTransactor) (*BitspawnTransactor, error) {
	contract, err := bindBitspawn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BitspawnTransactor{contract: contract}, nil
}

// NewBitspawnFilterer creates a new log filterer instance of Bitspawn, bound to a specific deployed contract.
func NewBitspawnFilterer(address common.Address, filterer bind.ContractFilterer) (*BitspawnFilterer, error) {
	contract, err := bindBitspawn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BitspawnFilterer{contract: contract}, nil
}

// bindBitspawn binds a generic wrapper to an already deployed contract.
func bindBitspawn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BitspawnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bitspawn *BitspawnRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bitspawn.Contract.BitspawnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bitspawn *BitspawnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bitspawn.Contract.BitspawnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bitspawn *BitspawnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bitspawn.Contract.BitspawnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bitspawn *BitspawnCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bitspawn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bitspawn *BitspawnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bitspawn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bitspawn *BitspawnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bitspawn.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_Bitspawn *BitspawnCaller) Allowance(opts *bind.CallOpts, src common.Address, guy common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bitspawn.contract.Call(opts, out, "allowance", src, guy)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_Bitspawn *BitspawnSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _Bitspawn.Contract.Allowance(&_Bitspawn.CallOpts, src, guy)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_Bitspawn *BitspawnCallerSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _Bitspawn.Contract.Allowance(&_Bitspawn.CallOpts, src, guy)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Bitspawn *BitspawnCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bitspawn.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Bitspawn *BitspawnSession) Authority() (common.Address, error) {
	return _Bitspawn.Contract.Authority(&_Bitspawn.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Bitspawn *BitspawnCallerSession) Authority() (common.Address, error) {
	return _Bitspawn.Contract.Authority(&_Bitspawn.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_Bitspawn *BitspawnCaller) BalanceOf(opts *bind.CallOpts, src common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bitspawn.contract.Call(opts, out, "balanceOf", src)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_Bitspawn *BitspawnSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _Bitspawn.Contract.BalanceOf(&_Bitspawn.CallOpts, src)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_Bitspawn *BitspawnCallerSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _Bitspawn.Contract.BalanceOf(&_Bitspawn.CallOpts, src)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Bitspawn *BitspawnCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bitspawn.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Bitspawn *BitspawnSession) Decimals() (*big.Int, error) {
	return _Bitspawn.Contract.Decimals(&_Bitspawn.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Bitspawn *BitspawnCallerSession) Decimals() (*big.Int, error) {
	return _Bitspawn.Contract.Decimals(&_Bitspawn.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_Bitspawn *BitspawnCaller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Bitspawn.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_Bitspawn *BitspawnSession) Name() ([32]byte, error) {
	return _Bitspawn.Contract.Name(&_Bitspawn.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(bytes32)
func (_Bitspawn *BitspawnCallerSession) Name() ([32]byte, error) {
	return _Bitspawn.Contract.Name(&_Bitspawn.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bitspawn *BitspawnCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bitspawn.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bitspawn *BitspawnSession) Owner() (common.Address, error) {
	return _Bitspawn.Contract.Owner(&_Bitspawn.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bitspawn *BitspawnCallerSession) Owner() (common.Address, error) {
	return _Bitspawn.Contract.Owner(&_Bitspawn.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Bitspawn *BitspawnCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bitspawn.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Bitspawn *BitspawnSession) Stopped() (bool, error) {
	return _Bitspawn.Contract.Stopped(&_Bitspawn.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Bitspawn *BitspawnCallerSession) Stopped() (bool, error) {
	return _Bitspawn.Contract.Stopped(&_Bitspawn.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_Bitspawn *BitspawnCaller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Bitspawn.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_Bitspawn *BitspawnSession) Symbol() ([32]byte, error) {
	return _Bitspawn.Contract.Symbol(&_Bitspawn.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_Bitspawn *BitspawnCallerSession) Symbol() ([32]byte, error) {
	return _Bitspawn.Contract.Symbol(&_Bitspawn.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Bitspawn *BitspawnCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bitspawn.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Bitspawn *BitspawnSession) TotalSupply() (*big.Int, error) {
	return _Bitspawn.Contract.TotalSupply(&_Bitspawn.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Bitspawn *BitspawnCallerSession) TotalSupply() (*big.Int, error) {
	return _Bitspawn.Contract.TotalSupply(&_Bitspawn.CallOpts)
}

// Trusted is a free data retrieval call binding the contract method 0x7aa3295b.
//
// Solidity: function trusted(src address, guy address) constant returns(bool)
func (_Bitspawn *BitspawnCaller) Trusted(opts *bind.CallOpts, src common.Address, guy common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bitspawn.contract.Call(opts, out, "trusted", src, guy)
	return *ret0, err
}

// Trusted is a free data retrieval call binding the contract method 0x7aa3295b.
//
// Solidity: function trusted(src address, guy address) constant returns(bool)
func (_Bitspawn *BitspawnSession) Trusted(src common.Address, guy common.Address) (bool, error) {
	return _Bitspawn.Contract.Trusted(&_Bitspawn.CallOpts, src, guy)
}

// Trusted is a free data retrieval call binding the contract method 0x7aa3295b.
//
// Solidity: function trusted(src address, guy address) constant returns(bool)
func (_Bitspawn *BitspawnCallerSession) Trusted(src common.Address, guy common.Address) (bool, error) {
	return _Bitspawn.Contract.Trusted(&_Bitspawn.CallOpts, src, guy)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_Bitspawn *BitspawnTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_Bitspawn *BitspawnSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.Contract.Approve(&_Bitspawn.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_Bitspawn *BitspawnTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.Contract.Approve(&_Bitspawn.TransactOpts, guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(guy address, wad uint256) returns()
func (_Bitspawn *BitspawnTransactor) Burn(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.contract.Transact(opts, "burn", guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(guy address, wad uint256) returns()
func (_Bitspawn *BitspawnSession) Burn(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.Contract.Burn(&_Bitspawn.TransactOpts, guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(guy address, wad uint256) returns()
func (_Bitspawn *BitspawnTransactorSession) Burn(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.Contract.Burn(&_Bitspawn.TransactOpts, guy, wad)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(wad uint256) returns()
func (_Bitspawn *BitspawnTransactor) Mint(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.contract.Transact(opts, "mint", wad)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(wad uint256) returns()
func (_Bitspawn *BitspawnSession) Mint(wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.Contract.Mint(&_Bitspawn.TransactOpts, wad)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(wad uint256) returns()
func (_Bitspawn *BitspawnTransactorSession) Mint(wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.Contract.Mint(&_Bitspawn.TransactOpts, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(src address, dst address, wad uint256) returns()
func (_Bitspawn *BitspawnTransactor) Move(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.contract.Transact(opts, "move", src, dst, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(src address, dst address, wad uint256) returns()
func (_Bitspawn *BitspawnSession) Move(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.Contract.Move(&_Bitspawn.TransactOpts, src, dst, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(src address, dst address, wad uint256) returns()
func (_Bitspawn *BitspawnTransactorSession) Move(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.Contract.Move(&_Bitspawn.TransactOpts, src, dst, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(src address, wad uint256) returns()
func (_Bitspawn *BitspawnTransactor) Pull(opts *bind.TransactOpts, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.contract.Transact(opts, "pull", src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(src address, wad uint256) returns()
func (_Bitspawn *BitspawnSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.Contract.Pull(&_Bitspawn.TransactOpts, src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(src address, wad uint256) returns()
func (_Bitspawn *BitspawnTransactorSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.Contract.Pull(&_Bitspawn.TransactOpts, src, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(dst address, wad uint256) returns()
func (_Bitspawn *BitspawnTransactor) Push(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.contract.Transact(opts, "push", dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(dst address, wad uint256) returns()
func (_Bitspawn *BitspawnSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.Contract.Push(&_Bitspawn.TransactOpts, dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(dst address, wad uint256) returns()
func (_Bitspawn *BitspawnTransactorSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.Contract.Push(&_Bitspawn.TransactOpts, dst, wad)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Bitspawn *BitspawnTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Bitspawn.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Bitspawn *BitspawnSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Bitspawn.Contract.SetAuthority(&_Bitspawn.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Bitspawn *BitspawnTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Bitspawn.Contract.SetAuthority(&_Bitspawn.TransactOpts, authority_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(name_ bytes32) returns()
func (_Bitspawn *BitspawnTransactor) SetName(opts *bind.TransactOpts, name_ [32]byte) (*types.Transaction, error) {
	return _Bitspawn.contract.Transact(opts, "setName", name_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(name_ bytes32) returns()
func (_Bitspawn *BitspawnSession) SetName(name_ [32]byte) (*types.Transaction, error) {
	return _Bitspawn.Contract.SetName(&_Bitspawn.TransactOpts, name_)
}

// SetName is a paid mutator transaction binding the contract method 0x5ac801fe.
//
// Solidity: function setName(name_ bytes32) returns()
func (_Bitspawn *BitspawnTransactorSession) SetName(name_ [32]byte) (*types.Transaction, error) {
	return _Bitspawn.Contract.SetName(&_Bitspawn.TransactOpts, name_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Bitspawn *BitspawnTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Bitspawn.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Bitspawn *BitspawnSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Bitspawn.Contract.SetOwner(&_Bitspawn.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Bitspawn *BitspawnTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Bitspawn.Contract.SetOwner(&_Bitspawn.TransactOpts, owner_)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Bitspawn *BitspawnTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bitspawn.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Bitspawn *BitspawnSession) Start() (*types.Transaction, error) {
	return _Bitspawn.Contract.Start(&_Bitspawn.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Bitspawn *BitspawnTransactorSession) Start() (*types.Transaction, error) {
	return _Bitspawn.Contract.Start(&_Bitspawn.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Bitspawn *BitspawnTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bitspawn.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Bitspawn *BitspawnSession) Stop() (*types.Transaction, error) {
	return _Bitspawn.Contract.Stop(&_Bitspawn.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Bitspawn *BitspawnTransactorSession) Stop() (*types.Transaction, error) {
	return _Bitspawn.Contract.Stop(&_Bitspawn.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_Bitspawn *BitspawnTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_Bitspawn *BitspawnSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.Contract.Transfer(&_Bitspawn.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_Bitspawn *BitspawnTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.Contract.Transfer(&_Bitspawn.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_Bitspawn *BitspawnTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_Bitspawn *BitspawnSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.Contract.TransferFrom(&_Bitspawn.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_Bitspawn *BitspawnTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Bitspawn.Contract.TransferFrom(&_Bitspawn.TransactOpts, src, dst, wad)
}

// Trust is a paid mutator transaction binding the contract method 0x06262f1b.
//
// Solidity: function trust(guy address, wat bool) returns()
func (_Bitspawn *BitspawnTransactor) Trust(opts *bind.TransactOpts, guy common.Address, wat bool) (*types.Transaction, error) {
	return _Bitspawn.contract.Transact(opts, "trust", guy, wat)
}

// Trust is a paid mutator transaction binding the contract method 0x06262f1b.
//
// Solidity: function trust(guy address, wat bool) returns()
func (_Bitspawn *BitspawnSession) Trust(guy common.Address, wat bool) (*types.Transaction, error) {
	return _Bitspawn.Contract.Trust(&_Bitspawn.TransactOpts, guy, wat)
}

// Trust is a paid mutator transaction binding the contract method 0x06262f1b.
//
// Solidity: function trust(guy address, wat bool) returns()
func (_Bitspawn *BitspawnTransactorSession) Trust(guy common.Address, wat bool) (*types.Transaction, error) {
	return _Bitspawn.Contract.Trust(&_Bitspawn.TransactOpts, guy, wat)
}

// BitspawnApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bitspawn contract.
type BitspawnApprovalIterator struct {
	Event *BitspawnApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BitspawnApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitspawnApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BitspawnApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BitspawnApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitspawnApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitspawnApproval represents a Approval event raised by the Bitspawn contract.
type BitspawnApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(src indexed address, guy indexed address, wad uint256)
func (_Bitspawn *BitspawnFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*BitspawnApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Bitspawn.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &BitspawnApprovalIterator{contract: _Bitspawn.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(src indexed address, guy indexed address, wad uint256)
func (_Bitspawn *BitspawnFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BitspawnApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Bitspawn.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitspawnApproval)
				if err := _Bitspawn.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BitspawnBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Bitspawn contract.
type BitspawnBurnIterator struct {
	Event *BitspawnBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BitspawnBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitspawnBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BitspawnBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BitspawnBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitspawnBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitspawnBurn represents a Burn event raised by the Bitspawn contract.
type BitspawnBurn struct {
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(guy indexed address, wad uint256)
func (_Bitspawn *BitspawnFilterer) FilterBurn(opts *bind.FilterOpts, guy []common.Address) (*BitspawnBurnIterator, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Bitspawn.contract.FilterLogs(opts, "Burn", guyRule)
	if err != nil {
		return nil, err
	}
	return &BitspawnBurnIterator{contract: _Bitspawn.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(guy indexed address, wad uint256)
func (_Bitspawn *BitspawnFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BitspawnBurn, guy []common.Address) (event.Subscription, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Bitspawn.contract.WatchLogs(opts, "Burn", guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitspawnBurn)
				if err := _Bitspawn.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BitspawnLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the Bitspawn contract.
type BitspawnLogSetAuthorityIterator struct {
	Event *BitspawnLogSetAuthority // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BitspawnLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitspawnLogSetAuthority)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BitspawnLogSetAuthority)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BitspawnLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitspawnLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitspawnLogSetAuthority represents a LogSetAuthority event raised by the Bitspawn contract.
type BitspawnLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: e LogSetAuthority(authority indexed address)
func (_Bitspawn *BitspawnFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*BitspawnLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Bitspawn.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &BitspawnLogSetAuthorityIterator{contract: _Bitspawn.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: e LogSetAuthority(authority indexed address)
func (_Bitspawn *BitspawnFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *BitspawnLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _Bitspawn.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitspawnLogSetAuthority)
				if err := _Bitspawn.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BitspawnLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the Bitspawn contract.
type BitspawnLogSetOwnerIterator struct {
	Event *BitspawnLogSetOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BitspawnLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitspawnLogSetOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BitspawnLogSetOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BitspawnLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitspawnLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitspawnLogSetOwner represents a LogSetOwner event raised by the Bitspawn contract.
type BitspawnLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: e LogSetOwner(owner indexed address)
func (_Bitspawn *BitspawnFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*BitspawnLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Bitspawn.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &BitspawnLogSetOwnerIterator{contract: _Bitspawn.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: e LogSetOwner(owner indexed address)
func (_Bitspawn *BitspawnFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *BitspawnLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Bitspawn.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitspawnLogSetOwner)
				if err := _Bitspawn.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BitspawnMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Bitspawn contract.
type BitspawnMintIterator struct {
	Event *BitspawnMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BitspawnMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitspawnMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BitspawnMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BitspawnMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitspawnMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitspawnMint represents a Mint event raised by the Bitspawn contract.
type BitspawnMint struct {
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(guy indexed address, wad uint256)
func (_Bitspawn *BitspawnFilterer) FilterMint(opts *bind.FilterOpts, guy []common.Address) (*BitspawnMintIterator, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Bitspawn.contract.FilterLogs(opts, "Mint", guyRule)
	if err != nil {
		return nil, err
	}
	return &BitspawnMintIterator{contract: _Bitspawn.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(guy indexed address, wad uint256)
func (_Bitspawn *BitspawnFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *BitspawnMint, guy []common.Address) (event.Subscription, error) {

	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Bitspawn.contract.WatchLogs(opts, "Mint", guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitspawnMint)
				if err := _Bitspawn.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BitspawnTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bitspawn contract.
type BitspawnTransferIterator struct {
	Event *BitspawnTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BitspawnTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitspawnTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BitspawnTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BitspawnTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitspawnTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitspawnTransfer represents a Transfer event raised by the Bitspawn contract.
type BitspawnTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(src indexed address, dst indexed address, wad uint256)
func (_Bitspawn *BitspawnFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*BitspawnTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Bitspawn.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &BitspawnTransferIterator{contract: _Bitspawn.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(src indexed address, dst indexed address, wad uint256)
func (_Bitspawn *BitspawnFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BitspawnTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Bitspawn.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitspawnTransfer)
				if err := _Bitspawn.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BitspawnTrustIterator is returned from FilterTrust and is used to iterate over the raw logs and unpacked data for Trust events raised by the Bitspawn contract.
type BitspawnTrustIterator struct {
	Event *BitspawnTrust // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BitspawnTrustIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BitspawnTrust)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BitspawnTrust)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BitspawnTrustIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BitspawnTrustIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BitspawnTrust represents a Trust event raised by the Bitspawn contract.
type BitspawnTrust struct {
	Src common.Address
	Guy common.Address
	Wat bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTrust is a free log retrieval operation binding the contract event 0xf184148577730b253ecb4339c543a564af420f3d32ed12a1c62ae83d67d65fe3.
//
// Solidity: e Trust(src indexed address, guy indexed address, wat bool)
func (_Bitspawn *BitspawnFilterer) FilterTrust(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*BitspawnTrustIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Bitspawn.contract.FilterLogs(opts, "Trust", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &BitspawnTrustIterator{contract: _Bitspawn.contract, event: "Trust", logs: logs, sub: sub}, nil
}

// WatchTrust is a free log subscription operation binding the contract event 0xf184148577730b253ecb4339c543a564af420f3d32ed12a1c62ae83d67d65fe3.
//
// Solidity: e Trust(src indexed address, guy indexed address, wat bool)
func (_Bitspawn *BitspawnFilterer) WatchTrust(opts *bind.WatchOpts, sink chan<- *BitspawnTrust, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Bitspawn.contract.WatchLogs(opts, "Trust", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BitspawnTrust)
				if err := _Bitspawn.contract.UnpackLog(event, "Trust", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DSAuthABI is the input ABI used to generate the binding from.
const DSAuthABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// DSAuthBin is the compiled bytecode used for deploying new contracts.
const DSAuthBin = `0x608060405234801561001057600080fd5b50600180546001600160a01b031916339081179091556040517fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9490600090a26102db8061005e6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806313af4035146100515780637a9e5e4b146100795780638da5cb5b1461009f578063bf7e214f146100c3575b600080fd5b6100776004803603602081101561006757600080fd5b50356001600160a01b03166100cb565b005b6100776004803603602081101561008f57600080fd5b50356001600160a01b031661013a565b6100a76101a5565b604080516001600160a01b039092168252519081900360200190f35b6100a76101b4565b6100e1336000356001600160e01b0319166101c3565b6100ea57600080fd5b600180546001600160a01b0319166001600160a01b0383811691909117918290556040519116907fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9490600090a250565b610150336000356001600160e01b0319166101c3565b61015957600080fd5b600080546001600160a01b0319166001600160a01b03838116919091178083556040519116917f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada491a250565b6001546001600160a01b031681565b6000546001600160a01b031681565b60006001600160a01b0383163014156101de575060016102a9565b6001546001600160a01b03848116911614156101fc575060016102a9565b6000546001600160a01b0316610214575060006102a9565b60005460408051600160e01b63b70096130281526001600160a01b0386811660048301523060248301526001600160e01b0319861660448301529151919092169163b7009613916064808301926020929190829003018186803b15801561027a57600080fd5b505afa15801561028e573d6000803e3d6000fd5b505050506040513d60208110156102a457600080fd5b505190505b9291505056fea165627a7a72305820efb720b527a6e879d9578a258579cc0395670bf18751809554b0a9f8c0ebaa770029`

// DeployDSAuth deploys a new Ethereum contract, binding an instance of DSAuth to it.
func DeployDSAuth(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DSAuth, error) {
	parsed, err := abi.JSON(strings.NewReader(DSAuthABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSAuthBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSAuth{DSAuthCaller: DSAuthCaller{contract: contract}, DSAuthTransactor: DSAuthTransactor{contract: contract}, DSAuthFilterer: DSAuthFilterer{contract: contract}}, nil
}

// DSAuth is an auto generated Go binding around an Ethereum contract.
type DSAuth struct {
	DSAuthCaller     // Read-only binding to the contract
	DSAuthTransactor // Write-only binding to the contract
	DSAuthFilterer   // Log filterer for contract events
}

// DSAuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSAuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSAuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSAuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSAuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DSAuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSAuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSAuthSession struct {
	Contract     *DSAuth           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSAuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSAuthCallerSession struct {
	Contract *DSAuthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DSAuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSAuthTransactorSession struct {
	Contract     *DSAuthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSAuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSAuthRaw struct {
	Contract *DSAuth // Generic contract binding to access the raw methods on
}

// DSAuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSAuthCallerRaw struct {
	Contract *DSAuthCaller // Generic read-only contract binding to access the raw methods on
}

// DSAuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSAuthTransactorRaw struct {
	Contract *DSAuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSAuth creates a new instance of DSAuth, bound to a specific deployed contract.
func NewDSAuth(address common.Address, backend bind.ContractBackend) (*DSAuth, error) {
	contract, err := bindDSAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSAuth{DSAuthCaller: DSAuthCaller{contract: contract}, DSAuthTransactor: DSAuthTransactor{contract: contract}, DSAuthFilterer: DSAuthFilterer{contract: contract}}, nil
}

// NewDSAuthCaller creates a new read-only instance of DSAuth, bound to a specific deployed contract.
func NewDSAuthCaller(address common.Address, caller bind.ContractCaller) (*DSAuthCaller, error) {
	contract, err := bindDSAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DSAuthCaller{contract: contract}, nil
}

// NewDSAuthTransactor creates a new write-only instance of DSAuth, bound to a specific deployed contract.
func NewDSAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*DSAuthTransactor, error) {
	contract, err := bindDSAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DSAuthTransactor{contract: contract}, nil
}

// NewDSAuthFilterer creates a new log filterer instance of DSAuth, bound to a specific deployed contract.
func NewDSAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*DSAuthFilterer, error) {
	contract, err := bindDSAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DSAuthFilterer{contract: contract}, nil
}

// bindDSAuth binds a generic wrapper to an already deployed contract.
func bindDSAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSAuthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSAuth *DSAuthRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSAuth.Contract.DSAuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSAuth *DSAuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSAuth.Contract.DSAuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSAuth *DSAuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSAuth.Contract.DSAuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSAuth *DSAuthCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSAuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSAuth *DSAuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSAuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSAuth *DSAuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSAuth.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSAuth *DSAuthCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSAuth.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSAuth *DSAuthSession) Authority() (common.Address, error) {
	return _DSAuth.Contract.Authority(&_DSAuth.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSAuth *DSAuthCallerSession) Authority() (common.Address, error) {
	return _DSAuth.Contract.Authority(&_DSAuth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSAuth *DSAuthCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSAuth.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSAuth *DSAuthSession) Owner() (common.Address, error) {
	return _DSAuth.Contract.Owner(&_DSAuth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSAuth *DSAuthCallerSession) Owner() (common.Address, error) {
	return _DSAuth.Contract.Owner(&_DSAuth.CallOpts)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_DSAuth *DSAuthTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _DSAuth.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_DSAuth *DSAuthSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _DSAuth.Contract.SetAuthority(&_DSAuth.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_DSAuth *DSAuthTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _DSAuth.Contract.SetAuthority(&_DSAuth.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_DSAuth *DSAuthTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _DSAuth.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_DSAuth *DSAuthSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _DSAuth.Contract.SetOwner(&_DSAuth.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_DSAuth *DSAuthTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _DSAuth.Contract.SetOwner(&_DSAuth.TransactOpts, owner_)
}

// DSAuthLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the DSAuth contract.
type DSAuthLogSetAuthorityIterator struct {
	Event *DSAuthLogSetAuthority // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DSAuthLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSAuthLogSetAuthority)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DSAuthLogSetAuthority)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DSAuthLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSAuthLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSAuthLogSetAuthority represents a LogSetAuthority event raised by the DSAuth contract.
type DSAuthLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: e LogSetAuthority(authority indexed address)
func (_DSAuth *DSAuthFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*DSAuthLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _DSAuth.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &DSAuthLogSetAuthorityIterator{contract: _DSAuth.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: e LogSetAuthority(authority indexed address)
func (_DSAuth *DSAuthFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *DSAuthLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _DSAuth.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSAuthLogSetAuthority)
				if err := _DSAuth.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DSAuthLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the DSAuth contract.
type DSAuthLogSetOwnerIterator struct {
	Event *DSAuthLogSetOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DSAuthLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSAuthLogSetOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DSAuthLogSetOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DSAuthLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSAuthLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSAuthLogSetOwner represents a LogSetOwner event raised by the DSAuth contract.
type DSAuthLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: e LogSetOwner(owner indexed address)
func (_DSAuth *DSAuthFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*DSAuthLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DSAuth.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &DSAuthLogSetOwnerIterator{contract: _DSAuth.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: e LogSetOwner(owner indexed address)
func (_DSAuth *DSAuthFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *DSAuthLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DSAuth.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSAuthLogSetOwner)
				if err := _DSAuth.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DSAuthEventsABI is the input ABI used to generate the binding from.
const DSAuthEventsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// DSAuthEventsBin is the compiled bytecode used for deploying new contracts.
const DSAuthEventsBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3fe6080604052600080fdfea165627a7a723058203e1acd63c18515ca472059ac0f698ab2ab3151f30d1dbce422674b11509a354a0029`

// DeployDSAuthEvents deploys a new Ethereum contract, binding an instance of DSAuthEvents to it.
func DeployDSAuthEvents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DSAuthEvents, error) {
	parsed, err := abi.JSON(strings.NewReader(DSAuthEventsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSAuthEventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSAuthEvents{DSAuthEventsCaller: DSAuthEventsCaller{contract: contract}, DSAuthEventsTransactor: DSAuthEventsTransactor{contract: contract}, DSAuthEventsFilterer: DSAuthEventsFilterer{contract: contract}}, nil
}

// DSAuthEvents is an auto generated Go binding around an Ethereum contract.
type DSAuthEvents struct {
	DSAuthEventsCaller     // Read-only binding to the contract
	DSAuthEventsTransactor // Write-only binding to the contract
	DSAuthEventsFilterer   // Log filterer for contract events
}

// DSAuthEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSAuthEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSAuthEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSAuthEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSAuthEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DSAuthEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSAuthEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSAuthEventsSession struct {
	Contract     *DSAuthEvents     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSAuthEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSAuthEventsCallerSession struct {
	Contract *DSAuthEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DSAuthEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSAuthEventsTransactorSession struct {
	Contract     *DSAuthEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DSAuthEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSAuthEventsRaw struct {
	Contract *DSAuthEvents // Generic contract binding to access the raw methods on
}

// DSAuthEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSAuthEventsCallerRaw struct {
	Contract *DSAuthEventsCaller // Generic read-only contract binding to access the raw methods on
}

// DSAuthEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSAuthEventsTransactorRaw struct {
	Contract *DSAuthEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSAuthEvents creates a new instance of DSAuthEvents, bound to a specific deployed contract.
func NewDSAuthEvents(address common.Address, backend bind.ContractBackend) (*DSAuthEvents, error) {
	contract, err := bindDSAuthEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSAuthEvents{DSAuthEventsCaller: DSAuthEventsCaller{contract: contract}, DSAuthEventsTransactor: DSAuthEventsTransactor{contract: contract}, DSAuthEventsFilterer: DSAuthEventsFilterer{contract: contract}}, nil
}

// NewDSAuthEventsCaller creates a new read-only instance of DSAuthEvents, bound to a specific deployed contract.
func NewDSAuthEventsCaller(address common.Address, caller bind.ContractCaller) (*DSAuthEventsCaller, error) {
	contract, err := bindDSAuthEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DSAuthEventsCaller{contract: contract}, nil
}

// NewDSAuthEventsTransactor creates a new write-only instance of DSAuthEvents, bound to a specific deployed contract.
func NewDSAuthEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*DSAuthEventsTransactor, error) {
	contract, err := bindDSAuthEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DSAuthEventsTransactor{contract: contract}, nil
}

// NewDSAuthEventsFilterer creates a new log filterer instance of DSAuthEvents, bound to a specific deployed contract.
func NewDSAuthEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*DSAuthEventsFilterer, error) {
	contract, err := bindDSAuthEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DSAuthEventsFilterer{contract: contract}, nil
}

// bindDSAuthEvents binds a generic wrapper to an already deployed contract.
func bindDSAuthEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSAuthEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSAuthEvents *DSAuthEventsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSAuthEvents.Contract.DSAuthEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSAuthEvents *DSAuthEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSAuthEvents.Contract.DSAuthEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSAuthEvents *DSAuthEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSAuthEvents.Contract.DSAuthEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSAuthEvents *DSAuthEventsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSAuthEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSAuthEvents *DSAuthEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSAuthEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSAuthEvents *DSAuthEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSAuthEvents.Contract.contract.Transact(opts, method, params...)
}

// DSAuthEventsLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the DSAuthEvents contract.
type DSAuthEventsLogSetAuthorityIterator struct {
	Event *DSAuthEventsLogSetAuthority // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DSAuthEventsLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSAuthEventsLogSetAuthority)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DSAuthEventsLogSetAuthority)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DSAuthEventsLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSAuthEventsLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSAuthEventsLogSetAuthority represents a LogSetAuthority event raised by the DSAuthEvents contract.
type DSAuthEventsLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: e LogSetAuthority(authority indexed address)
func (_DSAuthEvents *DSAuthEventsFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*DSAuthEventsLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _DSAuthEvents.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &DSAuthEventsLogSetAuthorityIterator{contract: _DSAuthEvents.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: e LogSetAuthority(authority indexed address)
func (_DSAuthEvents *DSAuthEventsFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *DSAuthEventsLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _DSAuthEvents.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSAuthEventsLogSetAuthority)
				if err := _DSAuthEvents.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DSAuthEventsLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the DSAuthEvents contract.
type DSAuthEventsLogSetOwnerIterator struct {
	Event *DSAuthEventsLogSetOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DSAuthEventsLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSAuthEventsLogSetOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DSAuthEventsLogSetOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DSAuthEventsLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSAuthEventsLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSAuthEventsLogSetOwner represents a LogSetOwner event raised by the DSAuthEvents contract.
type DSAuthEventsLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: e LogSetOwner(owner indexed address)
func (_DSAuthEvents *DSAuthEventsFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*DSAuthEventsLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DSAuthEvents.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &DSAuthEventsLogSetOwnerIterator{contract: _DSAuthEvents.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: e LogSetOwner(owner indexed address)
func (_DSAuthEvents *DSAuthEventsFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *DSAuthEventsLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DSAuthEvents.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSAuthEventsLogSetOwner)
				if err := _DSAuthEvents.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DSAuthorityABI is the input ABI used to generate the binding from.
const DSAuthorityABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"canCall\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DSAuthorityBin is the compiled bytecode used for deploying new contracts.
const DSAuthorityBin = `0x`

// DeployDSAuthority deploys a new Ethereum contract, binding an instance of DSAuthority to it.
func DeployDSAuthority(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DSAuthority, error) {
	parsed, err := abi.JSON(strings.NewReader(DSAuthorityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSAuthorityBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSAuthority{DSAuthorityCaller: DSAuthorityCaller{contract: contract}, DSAuthorityTransactor: DSAuthorityTransactor{contract: contract}, DSAuthorityFilterer: DSAuthorityFilterer{contract: contract}}, nil
}

// DSAuthority is an auto generated Go binding around an Ethereum contract.
type DSAuthority struct {
	DSAuthorityCaller     // Read-only binding to the contract
	DSAuthorityTransactor // Write-only binding to the contract
	DSAuthorityFilterer   // Log filterer for contract events
}

// DSAuthorityCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSAuthorityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSAuthorityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSAuthorityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSAuthorityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DSAuthorityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSAuthoritySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSAuthoritySession struct {
	Contract     *DSAuthority      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSAuthorityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSAuthorityCallerSession struct {
	Contract *DSAuthorityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DSAuthorityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSAuthorityTransactorSession struct {
	Contract     *DSAuthorityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DSAuthorityRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSAuthorityRaw struct {
	Contract *DSAuthority // Generic contract binding to access the raw methods on
}

// DSAuthorityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSAuthorityCallerRaw struct {
	Contract *DSAuthorityCaller // Generic read-only contract binding to access the raw methods on
}

// DSAuthorityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSAuthorityTransactorRaw struct {
	Contract *DSAuthorityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSAuthority creates a new instance of DSAuthority, bound to a specific deployed contract.
func NewDSAuthority(address common.Address, backend bind.ContractBackend) (*DSAuthority, error) {
	contract, err := bindDSAuthority(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSAuthority{DSAuthorityCaller: DSAuthorityCaller{contract: contract}, DSAuthorityTransactor: DSAuthorityTransactor{contract: contract}, DSAuthorityFilterer: DSAuthorityFilterer{contract: contract}}, nil
}

// NewDSAuthorityCaller creates a new read-only instance of DSAuthority, bound to a specific deployed contract.
func NewDSAuthorityCaller(address common.Address, caller bind.ContractCaller) (*DSAuthorityCaller, error) {
	contract, err := bindDSAuthority(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DSAuthorityCaller{contract: contract}, nil
}

// NewDSAuthorityTransactor creates a new write-only instance of DSAuthority, bound to a specific deployed contract.
func NewDSAuthorityTransactor(address common.Address, transactor bind.ContractTransactor) (*DSAuthorityTransactor, error) {
	contract, err := bindDSAuthority(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DSAuthorityTransactor{contract: contract}, nil
}

// NewDSAuthorityFilterer creates a new log filterer instance of DSAuthority, bound to a specific deployed contract.
func NewDSAuthorityFilterer(address common.Address, filterer bind.ContractFilterer) (*DSAuthorityFilterer, error) {
	contract, err := bindDSAuthority(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DSAuthorityFilterer{contract: contract}, nil
}

// bindDSAuthority binds a generic wrapper to an already deployed contract.
func bindDSAuthority(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSAuthorityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSAuthority *DSAuthorityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSAuthority.Contract.DSAuthorityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSAuthority *DSAuthorityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSAuthority.Contract.DSAuthorityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSAuthority *DSAuthorityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSAuthority.Contract.DSAuthorityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSAuthority *DSAuthorityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSAuthority.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSAuthority *DSAuthorityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSAuthority.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSAuthority *DSAuthorityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSAuthority.Contract.contract.Transact(opts, method, params...)
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(src address, dst address, sig bytes4) constant returns(bool)
func (_DSAuthority *DSAuthorityCaller) CanCall(opts *bind.CallOpts, src common.Address, dst common.Address, sig [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DSAuthority.contract.Call(opts, out, "canCall", src, dst, sig)
	return *ret0, err
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(src address, dst address, sig bytes4) constant returns(bool)
func (_DSAuthority *DSAuthoritySession) CanCall(src common.Address, dst common.Address, sig [4]byte) (bool, error) {
	return _DSAuthority.Contract.CanCall(&_DSAuthority.CallOpts, src, dst, sig)
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(src address, dst address, sig bytes4) constant returns(bool)
func (_DSAuthority *DSAuthorityCallerSession) CanCall(src common.Address, dst common.Address, sig [4]byte) (bool, error) {
	return _DSAuthority.Contract.CanCall(&_DSAuthority.CallOpts, src, dst, sig)
}

// DSMathABI is the input ABI used to generate the binding from.
const DSMathABI = "[]"

// DSMathBin is the compiled bytecode used for deploying new contracts.
const DSMathBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3fe6080604052600080fdfea165627a7a723058208c70c3c717cb5dd7605536dd47795524ce4f11b982bc355b8335f31ed20ca6ca0029`

// DeployDSMath deploys a new Ethereum contract, binding an instance of DSMath to it.
func DeployDSMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DSMath, error) {
	parsed, err := abi.JSON(strings.NewReader(DSMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSMath{DSMathCaller: DSMathCaller{contract: contract}, DSMathTransactor: DSMathTransactor{contract: contract}, DSMathFilterer: DSMathFilterer{contract: contract}}, nil
}

// DSMath is an auto generated Go binding around an Ethereum contract.
type DSMath struct {
	DSMathCaller     // Read-only binding to the contract
	DSMathTransactor // Write-only binding to the contract
	DSMathFilterer   // Log filterer for contract events
}

// DSMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DSMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSMathSession struct {
	Contract     *DSMath           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSMathCallerSession struct {
	Contract *DSMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DSMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSMathTransactorSession struct {
	Contract     *DSMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSMathRaw struct {
	Contract *DSMath // Generic contract binding to access the raw methods on
}

// DSMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSMathCallerRaw struct {
	Contract *DSMathCaller // Generic read-only contract binding to access the raw methods on
}

// DSMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSMathTransactorRaw struct {
	Contract *DSMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSMath creates a new instance of DSMath, bound to a specific deployed contract.
func NewDSMath(address common.Address, backend bind.ContractBackend) (*DSMath, error) {
	contract, err := bindDSMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSMath{DSMathCaller: DSMathCaller{contract: contract}, DSMathTransactor: DSMathTransactor{contract: contract}, DSMathFilterer: DSMathFilterer{contract: contract}}, nil
}

// NewDSMathCaller creates a new read-only instance of DSMath, bound to a specific deployed contract.
func NewDSMathCaller(address common.Address, caller bind.ContractCaller) (*DSMathCaller, error) {
	contract, err := bindDSMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DSMathCaller{contract: contract}, nil
}

// NewDSMathTransactor creates a new write-only instance of DSMath, bound to a specific deployed contract.
func NewDSMathTransactor(address common.Address, transactor bind.ContractTransactor) (*DSMathTransactor, error) {
	contract, err := bindDSMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DSMathTransactor{contract: contract}, nil
}

// NewDSMathFilterer creates a new log filterer instance of DSMath, bound to a specific deployed contract.
func NewDSMathFilterer(address common.Address, filterer bind.ContractFilterer) (*DSMathFilterer, error) {
	contract, err := bindDSMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DSMathFilterer{contract: contract}, nil
}

// bindDSMath binds a generic wrapper to an already deployed contract.
func bindDSMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSMath *DSMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSMath.Contract.DSMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSMath *DSMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSMath.Contract.DSMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSMath *DSMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSMath.Contract.DSMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSMath *DSMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSMath *DSMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSMath *DSMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSMath.Contract.contract.Transact(opts, method, params...)
}

// DSNoteABI is the input ABI used to generate the binding from.
const DSNoteABI = "[{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"}]"

// DSNoteBin is the compiled bytecode used for deploying new contracts.
const DSNoteBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3fe6080604052600080fdfea165627a7a723058207c1fa1d72eebf41bdc0ec1b861e2b57496c01676fafa9f23a0b203b04427de8e0029`

// DeployDSNote deploys a new Ethereum contract, binding an instance of DSNote to it.
func DeployDSNote(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DSNote, error) {
	parsed, err := abi.JSON(strings.NewReader(DSNoteABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSNoteBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSNote{DSNoteCaller: DSNoteCaller{contract: contract}, DSNoteTransactor: DSNoteTransactor{contract: contract}, DSNoteFilterer: DSNoteFilterer{contract: contract}}, nil
}

// DSNote is an auto generated Go binding around an Ethereum contract.
type DSNote struct {
	DSNoteCaller     // Read-only binding to the contract
	DSNoteTransactor // Write-only binding to the contract
	DSNoteFilterer   // Log filterer for contract events
}

// DSNoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSNoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSNoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSNoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSNoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DSNoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSNoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSNoteSession struct {
	Contract     *DSNote           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSNoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSNoteCallerSession struct {
	Contract *DSNoteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DSNoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSNoteTransactorSession struct {
	Contract     *DSNoteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSNoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSNoteRaw struct {
	Contract *DSNote // Generic contract binding to access the raw methods on
}

// DSNoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSNoteCallerRaw struct {
	Contract *DSNoteCaller // Generic read-only contract binding to access the raw methods on
}

// DSNoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSNoteTransactorRaw struct {
	Contract *DSNoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSNote creates a new instance of DSNote, bound to a specific deployed contract.
func NewDSNote(address common.Address, backend bind.ContractBackend) (*DSNote, error) {
	contract, err := bindDSNote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSNote{DSNoteCaller: DSNoteCaller{contract: contract}, DSNoteTransactor: DSNoteTransactor{contract: contract}, DSNoteFilterer: DSNoteFilterer{contract: contract}}, nil
}

// NewDSNoteCaller creates a new read-only instance of DSNote, bound to a specific deployed contract.
func NewDSNoteCaller(address common.Address, caller bind.ContractCaller) (*DSNoteCaller, error) {
	contract, err := bindDSNote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DSNoteCaller{contract: contract}, nil
}

// NewDSNoteTransactor creates a new write-only instance of DSNote, bound to a specific deployed contract.
func NewDSNoteTransactor(address common.Address, transactor bind.ContractTransactor) (*DSNoteTransactor, error) {
	contract, err := bindDSNote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DSNoteTransactor{contract: contract}, nil
}

// NewDSNoteFilterer creates a new log filterer instance of DSNote, bound to a specific deployed contract.
func NewDSNoteFilterer(address common.Address, filterer bind.ContractFilterer) (*DSNoteFilterer, error) {
	contract, err := bindDSNote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DSNoteFilterer{contract: contract}, nil
}

// bindDSNote binds a generic wrapper to an already deployed contract.
func bindDSNote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSNoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSNote *DSNoteRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSNote.Contract.DSNoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSNote *DSNoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSNote.Contract.DSNoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSNote *DSNoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSNote.Contract.DSNoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSNote *DSNoteCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSNote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSNote *DSNoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSNote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSNote *DSNoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSNote.Contract.contract.Transact(opts, method, params...)
}

// DSStopABI is the input ABI used to generate the binding from.
const DSStopABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"}]"

// DSStopBin is the compiled bytecode used for deploying new contracts.
const DSStopBin = `0x60806040819052600180546001600160a01b03191633908117909155907fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9490600090a26104aa806100516000396000f3fe6080604052600436106100705760003560e01c80637a9e5e4b1161004e5780637a9e5e4b146100db5780638da5cb5b1461010e578063be9a65551461013f578063bf7e214f1461014757610070565b806307da68f51461007557806313af40351461007f57806375f12b21146100b2575b600080fd5b61007d61015c565b005b34801561008b57600080fd5b5061007d600480360360208110156100a257600080fd5b50356001600160a01b03166101f6565b3480156100be57600080fd5b506100c7610265565b604080519115158252519081900360200190f35b3480156100e757600080fd5b5061007d600480360360208110156100fe57600080fd5b50356001600160a01b0316610275565b34801561011a57600080fd5b506101236102e0565b604080516001600160a01b039092168252519081900360200190f35b61007d6102ef565b34801561015357600080fd5b50610123610383565b610172336000356001600160e01b031916610392565b61017b57600080fd5b604080513480825260208201838152369383018490526004359360243593849386933393600080356001600160e01b03191694909260608201848480828437600083820152604051601f909101601f1916909201829003965090945050505050a4505060018054600160a01b60ff021916600160a01b179055565b61020c336000356001600160e01b031916610392565b61021557600080fd5b600180546001600160a01b0319166001600160a01b0383811691909117918290556040519116907fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9490600090a250565b600154600160a01b900460ff1681565b61028b336000356001600160e01b031916610392565b61029457600080fd5b600080546001600160a01b0319166001600160a01b03838116919091178083556040519116917f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada491a250565b6001546001600160a01b031681565b610305336000356001600160e01b031916610392565b61030e57600080fd5b604080513480825260208201838152369383018490526004359360243593849386933393600080356001600160e01b03191694909260608201848480828437600083820152604051601f909101601f1916909201829003965090945050505050a4505060018054600160a01b60ff0219169055565b6000546001600160a01b031681565b60006001600160a01b0383163014156103ad57506001610478565b6001546001600160a01b03848116911614156103cb57506001610478565b6000546001600160a01b03166103e357506000610478565b60005460408051600160e01b63b70096130281526001600160a01b0386811660048301523060248301526001600160e01b0319861660448301529151919092169163b7009613916064808301926020929190829003018186803b15801561044957600080fd5b505afa15801561045d573d6000803e3d6000fd5b505050506040513d602081101561047357600080fd5b505190505b9291505056fea165627a7a723058207a37f05e0571dc7d70768d306d2a9487e0d44f1937dd7581565569097a1ef31d0029`

// DeployDSStop deploys a new Ethereum contract, binding an instance of DSStop to it.
func DeployDSStop(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DSStop, error) {
	parsed, err := abi.JSON(strings.NewReader(DSStopABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSStopBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSStop{DSStopCaller: DSStopCaller{contract: contract}, DSStopTransactor: DSStopTransactor{contract: contract}, DSStopFilterer: DSStopFilterer{contract: contract}}, nil
}

// DSStop is an auto generated Go binding around an Ethereum contract.
type DSStop struct {
	DSStopCaller     // Read-only binding to the contract
	DSStopTransactor // Write-only binding to the contract
	DSStopFilterer   // Log filterer for contract events
}

// DSStopCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSStopCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSStopTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSStopTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSStopFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DSStopFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSStopSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSStopSession struct {
	Contract     *DSStop           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSStopCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSStopCallerSession struct {
	Contract *DSStopCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DSStopTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSStopTransactorSession struct {
	Contract     *DSStopTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSStopRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSStopRaw struct {
	Contract *DSStop // Generic contract binding to access the raw methods on
}

// DSStopCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSStopCallerRaw struct {
	Contract *DSStopCaller // Generic read-only contract binding to access the raw methods on
}

// DSStopTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSStopTransactorRaw struct {
	Contract *DSStopTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSStop creates a new instance of DSStop, bound to a specific deployed contract.
func NewDSStop(address common.Address, backend bind.ContractBackend) (*DSStop, error) {
	contract, err := bindDSStop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSStop{DSStopCaller: DSStopCaller{contract: contract}, DSStopTransactor: DSStopTransactor{contract: contract}, DSStopFilterer: DSStopFilterer{contract: contract}}, nil
}

// NewDSStopCaller creates a new read-only instance of DSStop, bound to a specific deployed contract.
func NewDSStopCaller(address common.Address, caller bind.ContractCaller) (*DSStopCaller, error) {
	contract, err := bindDSStop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DSStopCaller{contract: contract}, nil
}

// NewDSStopTransactor creates a new write-only instance of DSStop, bound to a specific deployed contract.
func NewDSStopTransactor(address common.Address, transactor bind.ContractTransactor) (*DSStopTransactor, error) {
	contract, err := bindDSStop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DSStopTransactor{contract: contract}, nil
}

// NewDSStopFilterer creates a new log filterer instance of DSStop, bound to a specific deployed contract.
func NewDSStopFilterer(address common.Address, filterer bind.ContractFilterer) (*DSStopFilterer, error) {
	contract, err := bindDSStop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DSStopFilterer{contract: contract}, nil
}

// bindDSStop binds a generic wrapper to an already deployed contract.
func bindDSStop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSStopABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSStop *DSStopRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSStop.Contract.DSStopCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSStop *DSStopRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSStop.Contract.DSStopTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSStop *DSStopRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSStop.Contract.DSStopTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSStop *DSStopCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSStop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSStop *DSStopTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSStop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSStop *DSStopTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSStop.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSStop *DSStopCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSStop.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSStop *DSStopSession) Authority() (common.Address, error) {
	return _DSStop.Contract.Authority(&_DSStop.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_DSStop *DSStopCallerSession) Authority() (common.Address, error) {
	return _DSStop.Contract.Authority(&_DSStop.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSStop *DSStopCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSStop.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSStop *DSStopSession) Owner() (common.Address, error) {
	return _DSStop.Contract.Owner(&_DSStop.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DSStop *DSStopCallerSession) Owner() (common.Address, error) {
	return _DSStop.Contract.Owner(&_DSStop.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_DSStop *DSStopCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DSStop.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_DSStop *DSStopSession) Stopped() (bool, error) {
	return _DSStop.Contract.Stopped(&_DSStop.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_DSStop *DSStopCallerSession) Stopped() (bool, error) {
	return _DSStop.Contract.Stopped(&_DSStop.CallOpts)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_DSStop *DSStopTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _DSStop.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_DSStop *DSStopSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _DSStop.Contract.SetAuthority(&_DSStop.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_DSStop *DSStopTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _DSStop.Contract.SetAuthority(&_DSStop.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_DSStop *DSStopTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _DSStop.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_DSStop *DSStopSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _DSStop.Contract.SetOwner(&_DSStop.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_DSStop *DSStopTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _DSStop.Contract.SetOwner(&_DSStop.TransactOpts, owner_)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_DSStop *DSStopTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSStop.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_DSStop *DSStopSession) Start() (*types.Transaction, error) {
	return _DSStop.Contract.Start(&_DSStop.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_DSStop *DSStopTransactorSession) Start() (*types.Transaction, error) {
	return _DSStop.Contract.Start(&_DSStop.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_DSStop *DSStopTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSStop.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_DSStop *DSStopSession) Stop() (*types.Transaction, error) {
	return _DSStop.Contract.Stop(&_DSStop.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_DSStop *DSStopTransactorSession) Stop() (*types.Transaction, error) {
	return _DSStop.Contract.Stop(&_DSStop.TransactOpts)
}

// DSStopLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the DSStop contract.
type DSStopLogSetAuthorityIterator struct {
	Event *DSStopLogSetAuthority // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DSStopLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSStopLogSetAuthority)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DSStopLogSetAuthority)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DSStopLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSStopLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSStopLogSetAuthority represents a LogSetAuthority event raised by the DSStop contract.
type DSStopLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: e LogSetAuthority(authority indexed address)
func (_DSStop *DSStopFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*DSStopLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _DSStop.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &DSStopLogSetAuthorityIterator{contract: _DSStop.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: e LogSetAuthority(authority indexed address)
func (_DSStop *DSStopFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *DSStopLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _DSStop.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSStopLogSetAuthority)
				if err := _DSStop.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DSStopLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the DSStop contract.
type DSStopLogSetOwnerIterator struct {
	Event *DSStopLogSetOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DSStopLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSStopLogSetOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DSStopLogSetOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DSStopLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSStopLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSStopLogSetOwner represents a LogSetOwner event raised by the DSStop contract.
type DSStopLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: e LogSetOwner(owner indexed address)
func (_DSStop *DSStopFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*DSStopLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DSStop.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &DSStopLogSetOwnerIterator{contract: _DSStop.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: e LogSetOwner(owner indexed address)
func (_DSStop *DSStopFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *DSStopLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _DSStop.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSStopLogSetOwner)
				if err := _DSStop.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DSTokenBaseABI is the input ABI used to generate the binding from.
const DSTokenBaseABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"supply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// DSTokenBaseBin is the compiled bytecode used for deploying new contracts.
const DSTokenBaseBin = `0x608060405234801561001057600080fd5b506040516020806105778339810180604052602081101561003057600080fd5b505133600090815260016020526040812082905555610523806100546000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063095ea7b31461006757806318160ddd146100a757806323b872dd146100c157806370a08231146100f7578063a9059cbb1461011d578063dd62ed3e14610149575b600080fd5b6100936004803603604081101561007d57600080fd5b506001600160a01b038135169060200135610177565b604080519115158252519081900360200190f35b6100af6101de565b60408051918252519081900360200190f35b610093600480360360608110156100d757600080fd5b506001600160a01b038135811691602081013590911690604001356101e4565b6100af6004803603602081101561010d57600080fd5b50356001600160a01b03166103e7565b6100936004803603604081101561013357600080fd5b506001600160a01b038135169060200135610402565b6100af6004803603604081101561015f57600080fd5b506001600160a01b0381358116916020013516610416565b3360008181526002602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35060015b92915050565b60005490565b60006001600160a01b03841633146102c4576001600160a01b03841660009081526002602090815260408083203384529091529020548211156102715760408051600160e51b62461bcd02815260206004820152601e60248201527f64732d746f6b656e2d696e73756666696369656e742d617070726f76616c0000604482015290519081900360640190fd5b6001600160a01b038416600090815260026020908152604080832033845290915290205461029f9083610441565b6001600160a01b03851660009081526002602090815260408083203384529091529020555b6001600160a01b0384166000908152600160205260409020548211156103345760408051600160e51b62461bcd02815260206004820152601d60248201527f64732d746f6b656e2d696e73756666696369656e742d62616c616e6365000000604482015290519081900360640190fd5b6001600160a01b0384166000908152600160205260409020546103579083610441565b6001600160a01b038086166000908152600160205260408082209390935590851681522054610386908361049c565b6001600160a01b0380851660008181526001602090815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a35060019392505050565b6001600160a01b031660009081526001602052604090205490565b600061040f3384846101e4565b9392505050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b808203828111156101d85760408051600160e51b62461bcd02815260206004820152601560248201527f64732d6d6174682d7375622d756e646572666c6f770000000000000000000000604482015290519081900360640190fd5b808201828110156101d85760408051600160e51b62461bcd02815260206004820152601460248201527f64732d6d6174682d6164642d6f766572666c6f77000000000000000000000000604482015290519081900360640190fdfea165627a7a72305820439e61c65bfda7f65f5c5fabb35a574b8711befb4bc2452716a94e40354cc6670029`

// DeployDSTokenBase deploys a new Ethereum contract, binding an instance of DSTokenBase to it.
func DeployDSTokenBase(auth *bind.TransactOpts, backend bind.ContractBackend, supply *big.Int) (common.Address, *types.Transaction, *DSTokenBase, error) {
	parsed, err := abi.JSON(strings.NewReader(DSTokenBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSTokenBaseBin), backend, supply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSTokenBase{DSTokenBaseCaller: DSTokenBaseCaller{contract: contract}, DSTokenBaseTransactor: DSTokenBaseTransactor{contract: contract}, DSTokenBaseFilterer: DSTokenBaseFilterer{contract: contract}}, nil
}

// DSTokenBase is an auto generated Go binding around an Ethereum contract.
type DSTokenBase struct {
	DSTokenBaseCaller     // Read-only binding to the contract
	DSTokenBaseTransactor // Write-only binding to the contract
	DSTokenBaseFilterer   // Log filterer for contract events
}

// DSTokenBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSTokenBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSTokenBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSTokenBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSTokenBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DSTokenBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSTokenBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSTokenBaseSession struct {
	Contract     *DSTokenBase      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSTokenBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSTokenBaseCallerSession struct {
	Contract *DSTokenBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DSTokenBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSTokenBaseTransactorSession struct {
	Contract     *DSTokenBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DSTokenBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSTokenBaseRaw struct {
	Contract *DSTokenBase // Generic contract binding to access the raw methods on
}

// DSTokenBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSTokenBaseCallerRaw struct {
	Contract *DSTokenBaseCaller // Generic read-only contract binding to access the raw methods on
}

// DSTokenBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSTokenBaseTransactorRaw struct {
	Contract *DSTokenBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSTokenBase creates a new instance of DSTokenBase, bound to a specific deployed contract.
func NewDSTokenBase(address common.Address, backend bind.ContractBackend) (*DSTokenBase, error) {
	contract, err := bindDSTokenBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSTokenBase{DSTokenBaseCaller: DSTokenBaseCaller{contract: contract}, DSTokenBaseTransactor: DSTokenBaseTransactor{contract: contract}, DSTokenBaseFilterer: DSTokenBaseFilterer{contract: contract}}, nil
}

// NewDSTokenBaseCaller creates a new read-only instance of DSTokenBase, bound to a specific deployed contract.
func NewDSTokenBaseCaller(address common.Address, caller bind.ContractCaller) (*DSTokenBaseCaller, error) {
	contract, err := bindDSTokenBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DSTokenBaseCaller{contract: contract}, nil
}

// NewDSTokenBaseTransactor creates a new write-only instance of DSTokenBase, bound to a specific deployed contract.
func NewDSTokenBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*DSTokenBaseTransactor, error) {
	contract, err := bindDSTokenBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DSTokenBaseTransactor{contract: contract}, nil
}

// NewDSTokenBaseFilterer creates a new log filterer instance of DSTokenBase, bound to a specific deployed contract.
func NewDSTokenBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*DSTokenBaseFilterer, error) {
	contract, err := bindDSTokenBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DSTokenBaseFilterer{contract: contract}, nil
}

// bindDSTokenBase binds a generic wrapper to an already deployed contract.
func bindDSTokenBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSTokenBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSTokenBase *DSTokenBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSTokenBase.Contract.DSTokenBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSTokenBase *DSTokenBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSTokenBase.Contract.DSTokenBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSTokenBase *DSTokenBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSTokenBase.Contract.DSTokenBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSTokenBase *DSTokenBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSTokenBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSTokenBase *DSTokenBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSTokenBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSTokenBase *DSTokenBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSTokenBase.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_DSTokenBase *DSTokenBaseCaller) Allowance(opts *bind.CallOpts, src common.Address, guy common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DSTokenBase.contract.Call(opts, out, "allowance", src, guy)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_DSTokenBase *DSTokenBaseSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _DSTokenBase.Contract.Allowance(&_DSTokenBase.CallOpts, src, guy)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_DSTokenBase *DSTokenBaseCallerSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _DSTokenBase.Contract.Allowance(&_DSTokenBase.CallOpts, src, guy)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_DSTokenBase *DSTokenBaseCaller) BalanceOf(opts *bind.CallOpts, src common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DSTokenBase.contract.Call(opts, out, "balanceOf", src)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_DSTokenBase *DSTokenBaseSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _DSTokenBase.Contract.BalanceOf(&_DSTokenBase.CallOpts, src)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_DSTokenBase *DSTokenBaseCallerSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _DSTokenBase.Contract.BalanceOf(&_DSTokenBase.CallOpts, src)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DSTokenBase *DSTokenBaseCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DSTokenBase.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DSTokenBase *DSTokenBaseSession) TotalSupply() (*big.Int, error) {
	return _DSTokenBase.Contract.TotalSupply(&_DSTokenBase.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DSTokenBase *DSTokenBaseCallerSession) TotalSupply() (*big.Int, error) {
	return _DSTokenBase.Contract.TotalSupply(&_DSTokenBase.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.Contract.Approve(&_DSTokenBase.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.Contract.Approve(&_DSTokenBase.TransactOpts, guy, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.Contract.Transfer(&_DSTokenBase.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.Contract.Transfer(&_DSTokenBase.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.Contract.TransferFrom(&_DSTokenBase.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_DSTokenBase *DSTokenBaseTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DSTokenBase.Contract.TransferFrom(&_DSTokenBase.TransactOpts, src, dst, wad)
}

// DSTokenBaseApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DSTokenBase contract.
type DSTokenBaseApprovalIterator struct {
	Event *DSTokenBaseApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DSTokenBaseApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSTokenBaseApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DSTokenBaseApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DSTokenBaseApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSTokenBaseApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSTokenBaseApproval represents a Approval event raised by the DSTokenBase contract.
type DSTokenBaseApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(src indexed address, guy indexed address, wad uint256)
func (_DSTokenBase *DSTokenBaseFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*DSTokenBaseApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _DSTokenBase.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &DSTokenBaseApprovalIterator{contract: _DSTokenBase.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(src indexed address, guy indexed address, wad uint256)
func (_DSTokenBase *DSTokenBaseFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DSTokenBaseApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _DSTokenBase.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSTokenBaseApproval)
				if err := _DSTokenBase.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DSTokenBaseTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DSTokenBase contract.
type DSTokenBaseTransferIterator struct {
	Event *DSTokenBaseTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DSTokenBaseTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSTokenBaseTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DSTokenBaseTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DSTokenBaseTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSTokenBaseTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSTokenBaseTransfer represents a Transfer event raised by the DSTokenBase contract.
type DSTokenBaseTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(src indexed address, dst indexed address, wad uint256)
func (_DSTokenBase *DSTokenBaseFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*DSTokenBaseTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _DSTokenBase.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &DSTokenBaseTransferIterator{contract: _DSTokenBase.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(src indexed address, dst indexed address, wad uint256)
func (_DSTokenBase *DSTokenBaseFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DSTokenBaseTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _DSTokenBase.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSTokenBaseTransfer)
				if err := _DSTokenBase.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, src common.Address, guy common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", src, guy)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, src, guy)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, src, guy)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(guy address) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, guy common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", guy)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(guy address) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(guy common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, guy)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(guy address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(guy common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, guy)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_ERC20 *ERC20Session) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, guy, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_ERC20 *ERC20Session) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, src, dst, wad)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(src indexed address, guy indexed address, wad uint256)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*ERC20ApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(src indexed address, guy indexed address, wad uint256)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(src indexed address, dst indexed address, wad uint256)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*ERC20TransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(src indexed address, dst indexed address, wad uint256)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ERC20EventsABI is the input ABI used to generate the binding from.
const ERC20EventsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20EventsBin is the compiled bytecode used for deploying new contracts.
const ERC20EventsBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3fe6080604052600080fdfea165627a7a72305820592aa4e5c1452d350f1f1fd6941fee4871082542696ce5c6a59dfc58c6f05cd80029`

// DeployERC20Events deploys a new Ethereum contract, binding an instance of ERC20Events to it.
func DeployERC20Events(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Events, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20EventsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20EventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Events{ERC20EventsCaller: ERC20EventsCaller{contract: contract}, ERC20EventsTransactor: ERC20EventsTransactor{contract: contract}, ERC20EventsFilterer: ERC20EventsFilterer{contract: contract}}, nil
}

// ERC20Events is an auto generated Go binding around an Ethereum contract.
type ERC20Events struct {
	ERC20EventsCaller     // Read-only binding to the contract
	ERC20EventsTransactor // Write-only binding to the contract
	ERC20EventsFilterer   // Log filterer for contract events
}

// ERC20EventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20EventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20EventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20EventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20EventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20EventsSession struct {
	Contract     *ERC20Events      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20EventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20EventsCallerSession struct {
	Contract *ERC20EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC20EventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20EventsTransactorSession struct {
	Contract     *ERC20EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC20EventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20EventsRaw struct {
	Contract *ERC20Events // Generic contract binding to access the raw methods on
}

// ERC20EventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20EventsCallerRaw struct {
	Contract *ERC20EventsCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20EventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20EventsTransactorRaw struct {
	Contract *ERC20EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Events creates a new instance of ERC20Events, bound to a specific deployed contract.
func NewERC20Events(address common.Address, backend bind.ContractBackend) (*ERC20Events, error) {
	contract, err := bindERC20Events(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Events{ERC20EventsCaller: ERC20EventsCaller{contract: contract}, ERC20EventsTransactor: ERC20EventsTransactor{contract: contract}, ERC20EventsFilterer: ERC20EventsFilterer{contract: contract}}, nil
}

// NewERC20EventsCaller creates a new read-only instance of ERC20Events, bound to a specific deployed contract.
func NewERC20EventsCaller(address common.Address, caller bind.ContractCaller) (*ERC20EventsCaller, error) {
	contract, err := bindERC20Events(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20EventsCaller{contract: contract}, nil
}

// NewERC20EventsTransactor creates a new write-only instance of ERC20Events, bound to a specific deployed contract.
func NewERC20EventsTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20EventsTransactor, error) {
	contract, err := bindERC20Events(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20EventsTransactor{contract: contract}, nil
}

// NewERC20EventsFilterer creates a new log filterer instance of ERC20Events, bound to a specific deployed contract.
func NewERC20EventsFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20EventsFilterer, error) {
	contract, err := bindERC20Events(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20EventsFilterer{contract: contract}, nil
}

// bindERC20Events binds a generic wrapper to an already deployed contract.
func bindERC20Events(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20EventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Events *ERC20EventsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Events.Contract.ERC20EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Events *ERC20EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Events.Contract.ERC20EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Events *ERC20EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Events.Contract.ERC20EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Events *ERC20EventsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Events *ERC20EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Events *ERC20EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Events.Contract.contract.Transact(opts, method, params...)
}

// ERC20EventsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Events contract.
type ERC20EventsApprovalIterator struct {
	Event *ERC20EventsApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20EventsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20EventsApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20EventsApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20EventsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20EventsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20EventsApproval represents a Approval event raised by the ERC20Events contract.
type ERC20EventsApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(src indexed address, guy indexed address, wad uint256)
func (_ERC20Events *ERC20EventsFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*ERC20EventsApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _ERC20Events.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &ERC20EventsApprovalIterator{contract: _ERC20Events.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(src indexed address, guy indexed address, wad uint256)
func (_ERC20Events *ERC20EventsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20EventsApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _ERC20Events.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20EventsApproval)
				if err := _ERC20Events.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ERC20EventsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Events contract.
type ERC20EventsTransferIterator struct {
	Event *ERC20EventsTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20EventsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20EventsTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20EventsTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20EventsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20EventsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20EventsTransfer represents a Transfer event raised by the ERC20Events contract.
type ERC20EventsTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(src indexed address, dst indexed address, wad uint256)
func (_ERC20Events *ERC20EventsFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*ERC20EventsTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _ERC20Events.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &ERC20EventsTransferIterator{contract: _ERC20Events.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(src indexed address, dst indexed address, wad uint256)
func (_ERC20Events *ERC20EventsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20EventsTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _ERC20Events.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20EventsTransfer)
				if err := _ERC20Events.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// FeeABI is the input ABI used to generate the binding from.
const FeeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"nextOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"take\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseAllFunds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenContractAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"LogPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogTake\",\"type\":\"event\"}]"

// FeeBin is the compiled bytecode used for deploying new contracts.
const FeeBin = `0x608060405234801561001057600080fd5b506040516020806105278339810180604052602081101561003057600080fd5b5051600080546001600160a01b03199081163317909155600180546001600160a01b03909316929091169190911790556104b88061006f6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806313af40351461006757806355a373d61461008f578063732c3297146100b35780638da5cb5b146100df57806390a82fe9146100e7578063c4076876146100ef575b600080fd5b61008d6004803603602081101561007d57600080fd5b50356001600160a01b031661011b565b005b610097610189565b604080516001600160a01b039092168252519081900360200190f35b61008d600480360360408110156100c957600080fd5b50803590602001356001600160a01b0316610198565b6100976102a6565b61008d6102b5565b61008d6004803603604081101561010557600080fd5b506001600160a01b038135169060200135610398565b6000546001600160a01b0316331461016757604051600160e51b62461bcd02815260040180806020018281038252602281526020018061046b6022913960400191505060405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b031681565b6000546001600160a01b031633146101e457604051600160e51b62461bcd02815260040180806020018281038252602281526020018061046b6022913960400191505060405180910390fd5b60015460408051600160e01b6323b872dd0281523060048201526001600160a01b03848116602483015260448201869052915191909216916323b872dd9160648083019260209291908290030181600087803b15801561024357600080fd5b505af1158015610257573d6000803e3d6000fd5b505050506040513d602081101561026d57600080fd5b50506040805183815290517f01b939a07dbe41ee1e401d3b79344dbbb2580adef8e270f6519ba4ebd15861599181900360200190a15050565b6000546001600160a01b031681565b6000546001600160a01b0316331461030157604051600160e51b62461bcd02815260040180806020018281038252602281526020018061046b6022913960400191505060405180910390fd5b60015460408051600160e01b6370a0823102815230600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b15801561034f57600080fd5b505afa158015610363573d6000803e3d6000fd5b505050506040513d602081101561037957600080fd5b50516000549091506103959082906001600160a01b0316610198565b50565b60015460408051600160e01b6323b872dd0281526001600160a01b03858116600483015230602483015260448201859052915191909216916323b872dd9160648083019260209291908290030181600087803b1580156103f757600080fd5b505af115801561040b573d6000803e3d6000fd5b505050506040513d602081101561042157600080fd5b5050604080518281526001600160a01b038416602082015281517fffae32d7bbdcf36142b1abb9d02cd7e5aa1ac4bfbc53355b7671b44f4d3ab720929181900390910190a1505056fe6d75737420626520746865206f776e6572206f66207468697320636f6e7472616374a165627a7a7230582017b1b8fe21b746b0d47907e5f05b66337f38d2158887b8b6b60f0503e4262d850029`

// DeployFee deploys a new Ethereum contract, binding an instance of Fee to it.
func DeployFee(auth *bind.TransactOpts, backend bind.ContractBackend, tokenContractAddr common.Address) (common.Address, *types.Transaction, *Fee, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FeeBin), backend, tokenContractAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Fee{FeeCaller: FeeCaller{contract: contract}, FeeTransactor: FeeTransactor{contract: contract}, FeeFilterer: FeeFilterer{contract: contract}}, nil
}

// Fee is an auto generated Go binding around an Ethereum contract.
type Fee struct {
	FeeCaller     // Read-only binding to the contract
	FeeTransactor // Write-only binding to the contract
	FeeFilterer   // Log filterer for contract events
}

// FeeCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeSession struct {
	Contract     *Fee              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeCallerSession struct {
	Contract *FeeCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FeeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeTransactorSession struct {
	Contract     *FeeTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeRaw struct {
	Contract *Fee // Generic contract binding to access the raw methods on
}

// FeeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeCallerRaw struct {
	Contract *FeeCaller // Generic read-only contract binding to access the raw methods on
}

// FeeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeTransactorRaw struct {
	Contract *FeeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFee creates a new instance of Fee, bound to a specific deployed contract.
func NewFee(address common.Address, backend bind.ContractBackend) (*Fee, error) {
	contract, err := bindFee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fee{FeeCaller: FeeCaller{contract: contract}, FeeTransactor: FeeTransactor{contract: contract}, FeeFilterer: FeeFilterer{contract: contract}}, nil
}

// NewFeeCaller creates a new read-only instance of Fee, bound to a specific deployed contract.
func NewFeeCaller(address common.Address, caller bind.ContractCaller) (*FeeCaller, error) {
	contract, err := bindFee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeCaller{contract: contract}, nil
}

// NewFeeTransactor creates a new write-only instance of Fee, bound to a specific deployed contract.
func NewFeeTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeTransactor, error) {
	contract, err := bindFee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeTransactor{contract: contract}, nil
}

// NewFeeFilterer creates a new log filterer instance of Fee, bound to a specific deployed contract.
func NewFeeFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeFilterer, error) {
	contract, err := bindFee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeFilterer{contract: contract}, nil
}

// bindFee binds a generic wrapper to an already deployed contract.
func bindFee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fee *FeeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Fee.Contract.FeeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fee *FeeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fee.Contract.FeeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fee *FeeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fee.Contract.FeeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fee *FeeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Fee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fee *FeeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fee *FeeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fee.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Fee *FeeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Fee.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Fee *FeeSession) Owner() (common.Address, error) {
	return _Fee.Contract.Owner(&_Fee.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Fee *FeeCallerSession) Owner() (common.Address, error) {
	return _Fee.Contract.Owner(&_Fee.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_Fee *FeeCaller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Fee.contract.Call(opts, out, "tokenContract")
	return *ret0, err
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_Fee *FeeSession) TokenContract() (common.Address, error) {
	return _Fee.Contract.TokenContract(&_Fee.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_Fee *FeeCallerSession) TokenContract() (common.Address, error) {
	return _Fee.Contract.TokenContract(&_Fee.CallOpts)
}

// Pay is a paid mutator transaction binding the contract method 0xc4076876.
//
// Solidity: function pay(sender address, amount uint256) returns()
func (_Fee *FeeTransactor) Pay(opts *bind.TransactOpts, sender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fee.contract.Transact(opts, "pay", sender, amount)
}

// Pay is a paid mutator transaction binding the contract method 0xc4076876.
//
// Solidity: function pay(sender address, amount uint256) returns()
func (_Fee *FeeSession) Pay(sender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fee.Contract.Pay(&_Fee.TransactOpts, sender, amount)
}

// Pay is a paid mutator transaction binding the contract method 0xc4076876.
//
// Solidity: function pay(sender address, amount uint256) returns()
func (_Fee *FeeTransactorSession) Pay(sender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fee.Contract.Pay(&_Fee.TransactOpts, sender, amount)
}

// ReleaseAllFunds is a paid mutator transaction binding the contract method 0x90a82fe9.
//
// Solidity: function releaseAllFunds() returns()
func (_Fee *FeeTransactor) ReleaseAllFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fee.contract.Transact(opts, "releaseAllFunds")
}

// ReleaseAllFunds is a paid mutator transaction binding the contract method 0x90a82fe9.
//
// Solidity: function releaseAllFunds() returns()
func (_Fee *FeeSession) ReleaseAllFunds() (*types.Transaction, error) {
	return _Fee.Contract.ReleaseAllFunds(&_Fee.TransactOpts)
}

// ReleaseAllFunds is a paid mutator transaction binding the contract method 0x90a82fe9.
//
// Solidity: function releaseAllFunds() returns()
func (_Fee *FeeTransactorSession) ReleaseAllFunds() (*types.Transaction, error) {
	return _Fee.Contract.ReleaseAllFunds(&_Fee.TransactOpts)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(nextOwner address) returns()
func (_Fee *FeeTransactor) SetOwner(opts *bind.TransactOpts, nextOwner common.Address) (*types.Transaction, error) {
	return _Fee.contract.Transact(opts, "setOwner", nextOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(nextOwner address) returns()
func (_Fee *FeeSession) SetOwner(nextOwner common.Address) (*types.Transaction, error) {
	return _Fee.Contract.SetOwner(&_Fee.TransactOpts, nextOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(nextOwner address) returns()
func (_Fee *FeeTransactorSession) SetOwner(nextOwner common.Address) (*types.Transaction, error) {
	return _Fee.Contract.SetOwner(&_Fee.TransactOpts, nextOwner)
}

// Take is a paid mutator transaction binding the contract method 0x732c3297.
//
// Solidity: function take(amount uint256, to address) returns()
func (_Fee *FeeTransactor) Take(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Fee.contract.Transact(opts, "take", amount, to)
}

// Take is a paid mutator transaction binding the contract method 0x732c3297.
//
// Solidity: function take(amount uint256, to address) returns()
func (_Fee *FeeSession) Take(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Fee.Contract.Take(&_Fee.TransactOpts, amount, to)
}

// Take is a paid mutator transaction binding the contract method 0x732c3297.
//
// Solidity: function take(amount uint256, to address) returns()
func (_Fee *FeeTransactorSession) Take(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Fee.Contract.Take(&_Fee.TransactOpts, amount, to)
}

// FeeLogPaymentIterator is returned from FilterLogPayment and is used to iterate over the raw logs and unpacked data for LogPayment events raised by the Fee contract.
type FeeLogPaymentIterator struct {
	Event *FeeLogPayment // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeeLogPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeLogPayment)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeeLogPayment)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeeLogPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeLogPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeLogPayment represents a LogPayment event raised by the Fee contract.
type FeeLogPayment struct {
	Amount      *big.Int
	Participant common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogPayment is a free log retrieval operation binding the contract event 0xffae32d7bbdcf36142b1abb9d02cd7e5aa1ac4bfbc53355b7671b44f4d3ab720.
//
// Solidity: e LogPayment(amount uint256, participant address)
func (_Fee *FeeFilterer) FilterLogPayment(opts *bind.FilterOpts) (*FeeLogPaymentIterator, error) {

	logs, sub, err := _Fee.contract.FilterLogs(opts, "LogPayment")
	if err != nil {
		return nil, err
	}
	return &FeeLogPaymentIterator{contract: _Fee.contract, event: "LogPayment", logs: logs, sub: sub}, nil
}

// WatchLogPayment is a free log subscription operation binding the contract event 0xffae32d7bbdcf36142b1abb9d02cd7e5aa1ac4bfbc53355b7671b44f4d3ab720.
//
// Solidity: e LogPayment(amount uint256, participant address)
func (_Fee *FeeFilterer) WatchLogPayment(opts *bind.WatchOpts, sink chan<- *FeeLogPayment) (event.Subscription, error) {

	logs, sub, err := _Fee.contract.WatchLogs(opts, "LogPayment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeLogPayment)
				if err := _Fee.contract.UnpackLog(event, "LogPayment", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// FeeLogTakeIterator is returned from FilterLogTake and is used to iterate over the raw logs and unpacked data for LogTake events raised by the Fee contract.
type FeeLogTakeIterator struct {
	Event *FeeLogTake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeeLogTakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeLogTake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeeLogTake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeeLogTakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeLogTakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeLogTake represents a LogTake event raised by the Fee contract.
type FeeLogTake struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogTake is a free log retrieval operation binding the contract event 0x01b939a07dbe41ee1e401d3b79344dbbb2580adef8e270f6519ba4ebd1586159.
//
// Solidity: e LogTake(amount uint256)
func (_Fee *FeeFilterer) FilterLogTake(opts *bind.FilterOpts) (*FeeLogTakeIterator, error) {

	logs, sub, err := _Fee.contract.FilterLogs(opts, "LogTake")
	if err != nil {
		return nil, err
	}
	return &FeeLogTakeIterator{contract: _Fee.contract, event: "LogTake", logs: logs, sub: sub}, nil
}

// WatchLogTake is a free log subscription operation binding the contract event 0x01b939a07dbe41ee1e401d3b79344dbbb2580adef8e270f6519ba4ebd1586159.
//
// Solidity: e LogTake(amount uint256)
func (_Fee *FeeFilterer) WatchLogTake(opts *bind.WatchOpts, sink chan<- *FeeLogTake) (event.Subscription, error) {

	logs, sub, err := _Fee.contract.WatchLogs(opts, "LogTake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeLogTake)
				if err := _Fee.contract.UnpackLog(event, "LogTake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x6080604052348015600f57600080fd5b50600080546001600160a01b03191633179055608a806100306000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80638da5cb5b14602d575b600080fd5b6033604f565b604080516001600160a01b039092168252519081900360200190f35b6000546001600160a01b03168156fea165627a7a72305820e1270c1692ca92a33eafb528d1b5ea711a5a644d305c9cbb01289d79c5dc5d910029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TournamentPoolABI is the input ABI used to generate the binding from.
const TournamentPoolABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"feeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"participants\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"backAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"finalPlacements\",\"type\":\"address[]\"},{\"name\":\"prizeAllocationPerTenThousand\",\"type\":\"uint256[]\"}],\"name\":\"completeTournament\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cancelTournament\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minEntryFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"organizerPercentage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numPlayers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeContractPercentage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"players\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isCompleted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenContractAddress\",\"type\":\"address\"},{\"name\":\"feeContractAddress\",\"type\":\"address\"},{\"name\":\"_feeContractPercentage\",\"type\":\"uint256\"},{\"name\":\"_organizerPercentage\",\"type\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_minEntryFee\",\"type\":\"uint256\"},{\"name\":\"initialParticipants\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"LogRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"LogUnregister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"payout\",\"type\":\"uint256\"}],\"name\":\"LogPayout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"backer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogBacker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"backer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogComplete\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogCancel\",\"type\":\"event\"}]"

// TournamentPoolBin is the compiled bytecode used for deploying new contracts.
const TournamentPoolBin = `0x608060405234801561001057600080fd5b50604051611360380380611360833981018060405260e081101561003357600080fd5b8151602083015160408401516060850151608086015160a087015160c088018051969895979496939592949193918301929164010000000081111561007757600080fd5b8201602081018481111561008a57600080fd5b81518560208202830111640100000000821117156100a757600080fd5b5050600580546001600160a01b03808d166001600160a01b031992831617909255600680548c841690831617905560038a90556004899055600080549289169290911691909117815560028690559093509150505b815181101561017157600082828151811061011357fe5b6020026020010151905061012c8161017e60201b60201c565b604080516001600160a01b038316815290517f4c2cd98156e6a27c8f70b5432e1dc7cee6bc43471c4bb8446cadbd44fbfbf2ba9181900360200190a1506001016100fc565b5050505050505050610207565b6001600160a01b03811660009081526008602052604090205460ff16610204576001600160a01b0381166000818152600860205260408120805460ff19166001908117909155600a805491820181559091527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a80180546001600160a01b03191690911790555b50565b61114a806102166000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80638da5cb5b11610097578063e2eb41ff11610066578063e2eb41ff14610319578063e79a198f1461033f578063f207564e14610347578063fa391c641461036457610100565b80638da5cb5b146102e457806397b2f556146102ec578063997953c6146102f4578063ca1d209d146102fc57610100565b80635c05b26e116100d35780635c05b26e146101a3578063646156aa146102cc57806376e5b862146102d457806377b2eae0146102dc57610100565b806306e297121461010557806309e69ede146101295780631fd882d11461016357806355a373d61461019b575b600080fd5b61010d61036c565b604080516001600160a01b039092168252519081900360200190f35b61014f6004803603602081101561013f57600080fd5b50356001600160a01b031661037b565b604080519115158252519081900360200190f35b6101896004803603602081101561017957600080fd5b50356001600160a01b0316610390565b60408051918252519081900360200190f35b61010d6103a2565b6102ca600480360360408110156101b957600080fd5b8101906020810181356401000000008111156101d457600080fd5b8201836020820111156101e657600080fd5b8035906020019184602083028401116401000000008311171561020857600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929594936020810193503591505064010000000081111561025857600080fd5b82018360208201111561026a57600080fd5b8035906020019184602083028401116401000000008311171561028c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506103b1945050505050565b005b6102ca6105f9565b610189610869565b61018961086f565b61010d610875565b610189610884565b61018961088a565b6102ca6004803603602081101561031257600080fd5b5035610890565b61014f6004803603602081101561032f57600080fd5b50356001600160a01b03166108f6565b6102ca61090b565b6102ca6004803603602081101561035d57600080fd5b5035610a48565b61014f610b93565b6006546001600160a01b031681565b60086020526000908152604090205460ff1681565b60076020526000908152604090205481565b6005546001600160a01b031681565b6000546001600160a01b031633146104135760408051600160e51b62461bcd02815260206004820152601e60248201527f6e6f7420746865206f776e6572206f66207468697320636f6e74726163740000604482015290519081900360640190fd5b600654600160a01b900460ff16156104635760408051600160e51b62461bcd02815260206004820152601f60248201526000805160206110a3833981519152604482015290519081900360640190fd5b80518251146104a657604051600160e51b62461bcd0281526004018080602001828103825260398152602001806110c36039913960400191505060405180910390fd5b6104b08282610ba3565b60055460065460408051600160e01b6370a08231028152306004820181905291516001600160a01b03948516946323b872dd94169185916370a0823191602480820192602092909190829003018186803b15801561050d57600080fd5b505afa158015610521573d6000803e3d6000fd5b505050506040513d602081101561053757600080fd5b50516040805163ffffffff861660e01b81526001600160a01b0394851660048201529290931660248301526044820152905160648083019260209291908290030181600087803b15801561058a57600080fd5b505af115801561059e573d6000803e3d6000fd5b505050506040513d60208110156105b457600080fd5b505060068054600160a01b60ff021916600160a01b1790556040517fbbf93173200b07feb5b440c63ea96d13e6fe53f2550e2a359eeb993f1d9f523890600090a15050565b6000546001600160a01b0316331461065b5760408051600160e51b62461bcd02815260206004820152601e60248201527f6e6f7420746865206f776e6572206f66207468697320636f6e74726163740000604482015290519081900360640190fd5b600654600160a01b900460ff16156106ab5760408051600160e51b62461bcd02815260206004820152601f60248201526000805160206110a3833981519152604482015290519081900360640190fd5b60005b600a54811015610721576000600a82815481106106c757fe5b60009182526020808320909101546001600160a01b0316808352600790915260409091205490915015610718576001600160a01b038116600090815260076020526040902054610718908290610df5565b506001016106ae565b5060055460005460408051600160e01b6370a08231028152306004820181905291516001600160a01b03948516946323b872dd94169185916370a0823191602480820192602092909190829003018186803b15801561077f57600080fd5b505afa158015610793573d6000803e3d6000fd5b505050506040513d60208110156107a957600080fd5b50516040805163ffffffff861660e01b81526001600160a01b0394851660048201529290931660248301526044820152905160648083019260209291908290030181600087803b1580156107fc57600080fd5b505af1158015610810573d6000803e3d6000fd5b505050506040513d602081101561082657600080fd5b505060068054600160a01b60ff021916600160a01b1790556040517f1086688fc15d38f80096d3c109f05d1b696290c27e5f6263e8e8d28a501a0c5d90600090a1565b60025481565b60045481565b6000546001600160a01b031681565b60015481565b60035481565b600654600160a01b900460ff16156108e05760408051600160e51b62461bcd02815260206004820152601f60248201526000805160206110a3833981519152604482015290519081900360640190fd5b6108e933610f37565b6108f33382610fbd565b50565b60096020526000908152604090205460ff1681565b600654600160a01b900460ff161561095b5760408051600160e51b62461bcd02815260206004820152601f60248201526000805160206110a3833981519152604482015290519081900360640190fd5b3360008181526009602052604090205460ff166109c25760408051600160e51b62461bcd02815260206004820152601a60248201527f7061727469636970616e74206e6f742072656769737465726564000000000000604482015290519081900360640190fd5b6001600160a01b0381166000908152600760205260409020546109e6908290610df5565b6001600160a01b038116600081815260096020908152604091829020805460ff1916905560018054600019019055815192835290517f11854d1b3c0aa24c7c879af700c0089a48a48e9280bac11f5370b90b7cca481c9281900390910190a150565b600654600160a01b900460ff1615610a985760408051600160e51b62461bcd02815260206004820152601f60248201526000805160206110a3833981519152604482015290519081900360640190fd5b600254811015610af25760408051600160e51b62461bcd02815260206004820152601960248201527f70617920746865206d696e696d756d20656e7472792066656500000000000000604482015290519081900360640190fd5b33610afc81610f37565b610b068183610fbd565b6001600160a01b03811660009081526009602052604090205460ff16610b53576001600160a01b0381166000908152600960205260409020805460ff191660019081179091558054810190555b604080516001600160a01b038316815290517f4c2cd98156e6a27c8f70b5432e1dc7cee6bc43471c4bb8446cadbd44fbfbf2ba9181900360200190a15050565b600654600160a01b900460ff1681565b60055460408051600160e01b6370a0823102815230600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b158015610bf157600080fd5b505afa158015610c05573d6000803e3d6000fd5b505050506040513d6020811015610c1b57600080fd5b5051600480546005546000805460408051600160e01b6323b872dd02815230968101969096526001600160a01b0391821660248701526064948702859004604487018190529051969750959216936323b872dd9381810193602093909283900390910190829087803b158015610c9057600080fd5b505af1158015610ca4573d6000803e3d6000fd5b505050506040513d6020811015610cba57600080fd5b50600090505b8451811015610dee576000858281518110610cd757fe5b602002602001015190506000620f4240868481518110610cf357fe5b60200260200101516003546004546064030387020281610d0f57fe5b60055460408051600160e01b6323b872dd0281523060048201526001600160a01b038781166024830152949093046044840181905290519094509216916323b872dd916064808201926020929091908290030181600087803b158015610d7457600080fd5b505af1158015610d88573d6000803e3d6000fd5b505050506040513d6020811015610d9e57600080fd5b5050604080516001600160a01b03841681526020810183905281517f7d6963073a93103d7a336ea21c7f029ad757c9ef4c0c6ad0ff86c689480e0ed8929181900390910190a15050600101610cc0565b5050505050565b6001600160a01b038216600090815260076020526040902054811115610e4f57604051600160e51b62461bcd0281526004018080602001828103825260238152602001806110fc6023913960400191505060405180910390fd5b60055460408051600160e01b6323b872dd0281523060048201526001600160a01b03858116602483015260448201859052915191909216916323b872dd9160648083019260209291908290030181600087803b158015610eae57600080fd5b505af1158015610ec2573d6000803e3d6000fd5b505050506040513d6020811015610ed857600080fd5b50506001600160a01b038216600081815260076020908152604091829020805485900390558151928352820183905280517fb6c0eca8138e097d71e2dd31e19a1266487f0553f170b7260ffe68bcbe9ff8a79281900390910190a15050565b6001600160a01b03811660009081526008602052604090205460ff166108f3576001600160a01b03166000818152600860205260408120805460ff19166001908117909155600a805491820181559091527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a80180546001600160a01b0319169091179055565b6001600160a01b0380831660008181526007602090815260408083208054870190556005548151600160e01b6323b872dd02815260048101959095523060248601526044850187905290519416936323b872dd93606480820194918390030190829087803b15801561102e57600080fd5b505af1158015611042573d6000803e3d6000fd5b505050506040513d602081101561105857600080fd5b5050604080516001600160a01b03841681526020810183905281517fda1f9d13deae0af1312bfbdc9f9b3ce4690b9abb55e276994ab00a5d79b18989929181900390910190a1505056fe746f75726e616d656e7420697320616c726561647920636f6d706c65746564006c656e677468206f662066696e616c506c6163656d656e7420616e64207072697a65416c6c6f636174696f6e20646f206e6f74206d617463686261636b656420616d6f756e74206e6f7420656e6f75676820666f7220726566756e64a165627a7a7230582016523f388683562e578aced6e71caa2eada09ce2907ffbf07b42b5aa3d3417c70029`

// DeployTournamentPool deploys a new Ethereum contract, binding an instance of TournamentPool to it.
func DeployTournamentPool(auth *bind.TransactOpts, backend bind.ContractBackend, tokenContractAddress common.Address, feeContractAddress common.Address, _feeContractPercentage *big.Int, _organizerPercentage *big.Int, _owner common.Address, _minEntryFee *big.Int, initialParticipants []common.Address) (common.Address, *types.Transaction, *TournamentPool, error) {
	parsed, err := abi.JSON(strings.NewReader(TournamentPoolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TournamentPoolBin), backend, tokenContractAddress, feeContractAddress, _feeContractPercentage, _organizerPercentage, _owner, _minEntryFee, initialParticipants)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TournamentPool{TournamentPoolCaller: TournamentPoolCaller{contract: contract}, TournamentPoolTransactor: TournamentPoolTransactor{contract: contract}, TournamentPoolFilterer: TournamentPoolFilterer{contract: contract}}, nil
}

// TournamentPool is an auto generated Go binding around an Ethereum contract.
type TournamentPool struct {
	TournamentPoolCaller     // Read-only binding to the contract
	TournamentPoolTransactor // Write-only binding to the contract
	TournamentPoolFilterer   // Log filterer for contract events
}

// TournamentPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type TournamentPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TournamentPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TournamentPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TournamentPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TournamentPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TournamentPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TournamentPoolSession struct {
	Contract     *TournamentPool   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TournamentPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TournamentPoolCallerSession struct {
	Contract *TournamentPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TournamentPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TournamentPoolTransactorSession struct {
	Contract     *TournamentPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TournamentPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type TournamentPoolRaw struct {
	Contract *TournamentPool // Generic contract binding to access the raw methods on
}

// TournamentPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TournamentPoolCallerRaw struct {
	Contract *TournamentPoolCaller // Generic read-only contract binding to access the raw methods on
}

// TournamentPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TournamentPoolTransactorRaw struct {
	Contract *TournamentPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTournamentPool creates a new instance of TournamentPool, bound to a specific deployed contract.
func NewTournamentPool(address common.Address, backend bind.ContractBackend) (*TournamentPool, error) {
	contract, err := bindTournamentPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TournamentPool{TournamentPoolCaller: TournamentPoolCaller{contract: contract}, TournamentPoolTransactor: TournamentPoolTransactor{contract: contract}, TournamentPoolFilterer: TournamentPoolFilterer{contract: contract}}, nil
}

// NewTournamentPoolCaller creates a new read-only instance of TournamentPool, bound to a specific deployed contract.
func NewTournamentPoolCaller(address common.Address, caller bind.ContractCaller) (*TournamentPoolCaller, error) {
	contract, err := bindTournamentPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TournamentPoolCaller{contract: contract}, nil
}

// NewTournamentPoolTransactor creates a new write-only instance of TournamentPool, bound to a specific deployed contract.
func NewTournamentPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*TournamentPoolTransactor, error) {
	contract, err := bindTournamentPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TournamentPoolTransactor{contract: contract}, nil
}

// NewTournamentPoolFilterer creates a new log filterer instance of TournamentPool, bound to a specific deployed contract.
func NewTournamentPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*TournamentPoolFilterer, error) {
	contract, err := bindTournamentPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TournamentPoolFilterer{contract: contract}, nil
}

// bindTournamentPool binds a generic wrapper to an already deployed contract.
func bindTournamentPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TournamentPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TournamentPool *TournamentPoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TournamentPool.Contract.TournamentPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TournamentPool *TournamentPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TournamentPool.Contract.TournamentPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TournamentPool *TournamentPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TournamentPool.Contract.TournamentPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TournamentPool *TournamentPoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TournamentPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TournamentPool *TournamentPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TournamentPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TournamentPool *TournamentPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TournamentPool.Contract.contract.Transact(opts, method, params...)
}

// BackAmount is a free data retrieval call binding the contract method 0x1fd882d1.
//
// Solidity: function backAmount( address) constant returns(uint256)
func (_TournamentPool *TournamentPoolCaller) BackAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TournamentPool.contract.Call(opts, out, "backAmount", arg0)
	return *ret0, err
}

// BackAmount is a free data retrieval call binding the contract method 0x1fd882d1.
//
// Solidity: function backAmount( address) constant returns(uint256)
func (_TournamentPool *TournamentPoolSession) BackAmount(arg0 common.Address) (*big.Int, error) {
	return _TournamentPool.Contract.BackAmount(&_TournamentPool.CallOpts, arg0)
}

// BackAmount is a free data retrieval call binding the contract method 0x1fd882d1.
//
// Solidity: function backAmount( address) constant returns(uint256)
func (_TournamentPool *TournamentPoolCallerSession) BackAmount(arg0 common.Address) (*big.Int, error) {
	return _TournamentPool.Contract.BackAmount(&_TournamentPool.CallOpts, arg0)
}

// FeeContract is a free data retrieval call binding the contract method 0x06e29712.
//
// Solidity: function feeContract() constant returns(address)
func (_TournamentPool *TournamentPoolCaller) FeeContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TournamentPool.contract.Call(opts, out, "feeContract")
	return *ret0, err
}

// FeeContract is a free data retrieval call binding the contract method 0x06e29712.
//
// Solidity: function feeContract() constant returns(address)
func (_TournamentPool *TournamentPoolSession) FeeContract() (common.Address, error) {
	return _TournamentPool.Contract.FeeContract(&_TournamentPool.CallOpts)
}

// FeeContract is a free data retrieval call binding the contract method 0x06e29712.
//
// Solidity: function feeContract() constant returns(address)
func (_TournamentPool *TournamentPoolCallerSession) FeeContract() (common.Address, error) {
	return _TournamentPool.Contract.FeeContract(&_TournamentPool.CallOpts)
}

// FeeContractPercentage is a free data retrieval call binding the contract method 0x997953c6.
//
// Solidity: function feeContractPercentage() constant returns(uint256)
func (_TournamentPool *TournamentPoolCaller) FeeContractPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TournamentPool.contract.Call(opts, out, "feeContractPercentage")
	return *ret0, err
}

// FeeContractPercentage is a free data retrieval call binding the contract method 0x997953c6.
//
// Solidity: function feeContractPercentage() constant returns(uint256)
func (_TournamentPool *TournamentPoolSession) FeeContractPercentage() (*big.Int, error) {
	return _TournamentPool.Contract.FeeContractPercentage(&_TournamentPool.CallOpts)
}

// FeeContractPercentage is a free data retrieval call binding the contract method 0x997953c6.
//
// Solidity: function feeContractPercentage() constant returns(uint256)
func (_TournamentPool *TournamentPoolCallerSession) FeeContractPercentage() (*big.Int, error) {
	return _TournamentPool.Contract.FeeContractPercentage(&_TournamentPool.CallOpts)
}

// IsCompleted is a free data retrieval call binding the contract method 0xfa391c64.
//
// Solidity: function isCompleted() constant returns(bool)
func (_TournamentPool *TournamentPoolCaller) IsCompleted(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TournamentPool.contract.Call(opts, out, "isCompleted")
	return *ret0, err
}

// IsCompleted is a free data retrieval call binding the contract method 0xfa391c64.
//
// Solidity: function isCompleted() constant returns(bool)
func (_TournamentPool *TournamentPoolSession) IsCompleted() (bool, error) {
	return _TournamentPool.Contract.IsCompleted(&_TournamentPool.CallOpts)
}

// IsCompleted is a free data retrieval call binding the contract method 0xfa391c64.
//
// Solidity: function isCompleted() constant returns(bool)
func (_TournamentPool *TournamentPoolCallerSession) IsCompleted() (bool, error) {
	return _TournamentPool.Contract.IsCompleted(&_TournamentPool.CallOpts)
}

// MinEntryFee is a free data retrieval call binding the contract method 0x76e5b862.
//
// Solidity: function minEntryFee() constant returns(uint256)
func (_TournamentPool *TournamentPoolCaller) MinEntryFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TournamentPool.contract.Call(opts, out, "minEntryFee")
	return *ret0, err
}

// MinEntryFee is a free data retrieval call binding the contract method 0x76e5b862.
//
// Solidity: function minEntryFee() constant returns(uint256)
func (_TournamentPool *TournamentPoolSession) MinEntryFee() (*big.Int, error) {
	return _TournamentPool.Contract.MinEntryFee(&_TournamentPool.CallOpts)
}

// MinEntryFee is a free data retrieval call binding the contract method 0x76e5b862.
//
// Solidity: function minEntryFee() constant returns(uint256)
func (_TournamentPool *TournamentPoolCallerSession) MinEntryFee() (*big.Int, error) {
	return _TournamentPool.Contract.MinEntryFee(&_TournamentPool.CallOpts)
}

// NumPlayers is a free data retrieval call binding the contract method 0x97b2f556.
//
// Solidity: function numPlayers() constant returns(uint256)
func (_TournamentPool *TournamentPoolCaller) NumPlayers(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TournamentPool.contract.Call(opts, out, "numPlayers")
	return *ret0, err
}

// NumPlayers is a free data retrieval call binding the contract method 0x97b2f556.
//
// Solidity: function numPlayers() constant returns(uint256)
func (_TournamentPool *TournamentPoolSession) NumPlayers() (*big.Int, error) {
	return _TournamentPool.Contract.NumPlayers(&_TournamentPool.CallOpts)
}

// NumPlayers is a free data retrieval call binding the contract method 0x97b2f556.
//
// Solidity: function numPlayers() constant returns(uint256)
func (_TournamentPool *TournamentPoolCallerSession) NumPlayers() (*big.Int, error) {
	return _TournamentPool.Contract.NumPlayers(&_TournamentPool.CallOpts)
}

// OrganizerPercentage is a free data retrieval call binding the contract method 0x77b2eae0.
//
// Solidity: function organizerPercentage() constant returns(uint256)
func (_TournamentPool *TournamentPoolCaller) OrganizerPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TournamentPool.contract.Call(opts, out, "organizerPercentage")
	return *ret0, err
}

// OrganizerPercentage is a free data retrieval call binding the contract method 0x77b2eae0.
//
// Solidity: function organizerPercentage() constant returns(uint256)
func (_TournamentPool *TournamentPoolSession) OrganizerPercentage() (*big.Int, error) {
	return _TournamentPool.Contract.OrganizerPercentage(&_TournamentPool.CallOpts)
}

// OrganizerPercentage is a free data retrieval call binding the contract method 0x77b2eae0.
//
// Solidity: function organizerPercentage() constant returns(uint256)
func (_TournamentPool *TournamentPoolCallerSession) OrganizerPercentage() (*big.Int, error) {
	return _TournamentPool.Contract.OrganizerPercentage(&_TournamentPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TournamentPool *TournamentPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TournamentPool.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TournamentPool *TournamentPoolSession) Owner() (common.Address, error) {
	return _TournamentPool.Contract.Owner(&_TournamentPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TournamentPool *TournamentPoolCallerSession) Owner() (common.Address, error) {
	return _TournamentPool.Contract.Owner(&_TournamentPool.CallOpts)
}

// Participants is a free data retrieval call binding the contract method 0x09e69ede.
//
// Solidity: function participants( address) constant returns(bool)
func (_TournamentPool *TournamentPoolCaller) Participants(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TournamentPool.contract.Call(opts, out, "participants", arg0)
	return *ret0, err
}

// Participants is a free data retrieval call binding the contract method 0x09e69ede.
//
// Solidity: function participants( address) constant returns(bool)
func (_TournamentPool *TournamentPoolSession) Participants(arg0 common.Address) (bool, error) {
	return _TournamentPool.Contract.Participants(&_TournamentPool.CallOpts, arg0)
}

// Participants is a free data retrieval call binding the contract method 0x09e69ede.
//
// Solidity: function participants( address) constant returns(bool)
func (_TournamentPool *TournamentPoolCallerSession) Participants(arg0 common.Address) (bool, error) {
	return _TournamentPool.Contract.Participants(&_TournamentPool.CallOpts, arg0)
}

// Players is a free data retrieval call binding the contract method 0xe2eb41ff.
//
// Solidity: function players( address) constant returns(bool)
func (_TournamentPool *TournamentPoolCaller) Players(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TournamentPool.contract.Call(opts, out, "players", arg0)
	return *ret0, err
}

// Players is a free data retrieval call binding the contract method 0xe2eb41ff.
//
// Solidity: function players( address) constant returns(bool)
func (_TournamentPool *TournamentPoolSession) Players(arg0 common.Address) (bool, error) {
	return _TournamentPool.Contract.Players(&_TournamentPool.CallOpts, arg0)
}

// Players is a free data retrieval call binding the contract method 0xe2eb41ff.
//
// Solidity: function players( address) constant returns(bool)
func (_TournamentPool *TournamentPoolCallerSession) Players(arg0 common.Address) (bool, error) {
	return _TournamentPool.Contract.Players(&_TournamentPool.CallOpts, arg0)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_TournamentPool *TournamentPoolCaller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TournamentPool.contract.Call(opts, out, "tokenContract")
	return *ret0, err
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_TournamentPool *TournamentPoolSession) TokenContract() (common.Address, error) {
	return _TournamentPool.Contract.TokenContract(&_TournamentPool.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_TournamentPool *TournamentPoolCallerSession) TokenContract() (common.Address, error) {
	return _TournamentPool.Contract.TokenContract(&_TournamentPool.CallOpts)
}

// CancelTournament is a paid mutator transaction binding the contract method 0x646156aa.
//
// Solidity: function cancelTournament() returns()
func (_TournamentPool *TournamentPoolTransactor) CancelTournament(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TournamentPool.contract.Transact(opts, "cancelTournament")
}

// CancelTournament is a paid mutator transaction binding the contract method 0x646156aa.
//
// Solidity: function cancelTournament() returns()
func (_TournamentPool *TournamentPoolSession) CancelTournament() (*types.Transaction, error) {
	return _TournamentPool.Contract.CancelTournament(&_TournamentPool.TransactOpts)
}

// CancelTournament is a paid mutator transaction binding the contract method 0x646156aa.
//
// Solidity: function cancelTournament() returns()
func (_TournamentPool *TournamentPoolTransactorSession) CancelTournament() (*types.Transaction, error) {
	return _TournamentPool.Contract.CancelTournament(&_TournamentPool.TransactOpts)
}

// CompleteTournament is a paid mutator transaction binding the contract method 0x5c05b26e.
//
// Solidity: function completeTournament(finalPlacements address[], prizeAllocationPerTenThousand uint256[]) returns()
func (_TournamentPool *TournamentPoolTransactor) CompleteTournament(opts *bind.TransactOpts, finalPlacements []common.Address, prizeAllocationPerTenThousand []*big.Int) (*types.Transaction, error) {
	return _TournamentPool.contract.Transact(opts, "completeTournament", finalPlacements, prizeAllocationPerTenThousand)
}

// CompleteTournament is a paid mutator transaction binding the contract method 0x5c05b26e.
//
// Solidity: function completeTournament(finalPlacements address[], prizeAllocationPerTenThousand uint256[]) returns()
func (_TournamentPool *TournamentPoolSession) CompleteTournament(finalPlacements []common.Address, prizeAllocationPerTenThousand []*big.Int) (*types.Transaction, error) {
	return _TournamentPool.Contract.CompleteTournament(&_TournamentPool.TransactOpts, finalPlacements, prizeAllocationPerTenThousand)
}

// CompleteTournament is a paid mutator transaction binding the contract method 0x5c05b26e.
//
// Solidity: function completeTournament(finalPlacements address[], prizeAllocationPerTenThousand uint256[]) returns()
func (_TournamentPool *TournamentPoolTransactorSession) CompleteTournament(finalPlacements []common.Address, prizeAllocationPerTenThousand []*big.Int) (*types.Transaction, error) {
	return _TournamentPool.Contract.CompleteTournament(&_TournamentPool.TransactOpts, finalPlacements, prizeAllocationPerTenThousand)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(amount uint256) returns()
func (_TournamentPool *TournamentPoolTransactor) Fund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TournamentPool.contract.Transact(opts, "fund", amount)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(amount uint256) returns()
func (_TournamentPool *TournamentPoolSession) Fund(amount *big.Int) (*types.Transaction, error) {
	return _TournamentPool.Contract.Fund(&_TournamentPool.TransactOpts, amount)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(amount uint256) returns()
func (_TournamentPool *TournamentPoolTransactorSession) Fund(amount *big.Int) (*types.Transaction, error) {
	return _TournamentPool.Contract.Fund(&_TournamentPool.TransactOpts, amount)
}

// Register is a paid mutator transaction binding the contract method 0xf207564e.
//
// Solidity: function register(amount uint256) returns()
func (_TournamentPool *TournamentPoolTransactor) Register(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TournamentPool.contract.Transact(opts, "register", amount)
}

// Register is a paid mutator transaction binding the contract method 0xf207564e.
//
// Solidity: function register(amount uint256) returns()
func (_TournamentPool *TournamentPoolSession) Register(amount *big.Int) (*types.Transaction, error) {
	return _TournamentPool.Contract.Register(&_TournamentPool.TransactOpts, amount)
}

// Register is a paid mutator transaction binding the contract method 0xf207564e.
//
// Solidity: function register(amount uint256) returns()
func (_TournamentPool *TournamentPoolTransactorSession) Register(amount *big.Int) (*types.Transaction, error) {
	return _TournamentPool.Contract.Register(&_TournamentPool.TransactOpts, amount)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_TournamentPool *TournamentPoolTransactor) Unregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TournamentPool.contract.Transact(opts, "unregister")
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_TournamentPool *TournamentPoolSession) Unregister() (*types.Transaction, error) {
	return _TournamentPool.Contract.Unregister(&_TournamentPool.TransactOpts)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_TournamentPool *TournamentPoolTransactorSession) Unregister() (*types.Transaction, error) {
	return _TournamentPool.Contract.Unregister(&_TournamentPool.TransactOpts)
}

// TournamentPoolLogBackerIterator is returned from FilterLogBacker and is used to iterate over the raw logs and unpacked data for LogBacker events raised by the TournamentPool contract.
type TournamentPoolLogBackerIterator struct {
	Event *TournamentPoolLogBacker // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TournamentPoolLogBackerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TournamentPoolLogBacker)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TournamentPoolLogBacker)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TournamentPoolLogBackerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TournamentPoolLogBackerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TournamentPoolLogBacker represents a LogBacker event raised by the TournamentPool contract.
type TournamentPoolLogBacker struct {
	Backer common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogBacker is a free log retrieval operation binding the contract event 0xda1f9d13deae0af1312bfbdc9f9b3ce4690b9abb55e276994ab00a5d79b18989.
//
// Solidity: e LogBacker(backer address, amount uint256)
func (_TournamentPool *TournamentPoolFilterer) FilterLogBacker(opts *bind.FilterOpts) (*TournamentPoolLogBackerIterator, error) {

	logs, sub, err := _TournamentPool.contract.FilterLogs(opts, "LogBacker")
	if err != nil {
		return nil, err
	}
	return &TournamentPoolLogBackerIterator{contract: _TournamentPool.contract, event: "LogBacker", logs: logs, sub: sub}, nil
}

// WatchLogBacker is a free log subscription operation binding the contract event 0xda1f9d13deae0af1312bfbdc9f9b3ce4690b9abb55e276994ab00a5d79b18989.
//
// Solidity: e LogBacker(backer address, amount uint256)
func (_TournamentPool *TournamentPoolFilterer) WatchLogBacker(opts *bind.WatchOpts, sink chan<- *TournamentPoolLogBacker) (event.Subscription, error) {

	logs, sub, err := _TournamentPool.contract.WatchLogs(opts, "LogBacker")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TournamentPoolLogBacker)
				if err := _TournamentPool.contract.UnpackLog(event, "LogBacker", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TournamentPoolLogCancelIterator is returned from FilterLogCancel and is used to iterate over the raw logs and unpacked data for LogCancel events raised by the TournamentPool contract.
type TournamentPoolLogCancelIterator struct {
	Event *TournamentPoolLogCancel // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TournamentPoolLogCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TournamentPoolLogCancel)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TournamentPoolLogCancel)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TournamentPoolLogCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TournamentPoolLogCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TournamentPoolLogCancel represents a LogCancel event raised by the TournamentPool contract.
type TournamentPoolLogCancel struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogCancel is a free log retrieval operation binding the contract event 0x1086688fc15d38f80096d3c109f05d1b696290c27e5f6263e8e8d28a501a0c5d.
//
// Solidity: e LogCancel()
func (_TournamentPool *TournamentPoolFilterer) FilterLogCancel(opts *bind.FilterOpts) (*TournamentPoolLogCancelIterator, error) {

	logs, sub, err := _TournamentPool.contract.FilterLogs(opts, "LogCancel")
	if err != nil {
		return nil, err
	}
	return &TournamentPoolLogCancelIterator{contract: _TournamentPool.contract, event: "LogCancel", logs: logs, sub: sub}, nil
}

// WatchLogCancel is a free log subscription operation binding the contract event 0x1086688fc15d38f80096d3c109f05d1b696290c27e5f6263e8e8d28a501a0c5d.
//
// Solidity: e LogCancel()
func (_TournamentPool *TournamentPoolFilterer) WatchLogCancel(opts *bind.WatchOpts, sink chan<- *TournamentPoolLogCancel) (event.Subscription, error) {

	logs, sub, err := _TournamentPool.contract.WatchLogs(opts, "LogCancel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TournamentPoolLogCancel)
				if err := _TournamentPool.contract.UnpackLog(event, "LogCancel", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TournamentPoolLogCompleteIterator is returned from FilterLogComplete and is used to iterate over the raw logs and unpacked data for LogComplete events raised by the TournamentPool contract.
type TournamentPoolLogCompleteIterator struct {
	Event *TournamentPoolLogComplete // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TournamentPoolLogCompleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TournamentPoolLogComplete)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TournamentPoolLogComplete)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TournamentPoolLogCompleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TournamentPoolLogCompleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TournamentPoolLogComplete represents a LogComplete event raised by the TournamentPool contract.
type TournamentPoolLogComplete struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogComplete is a free log retrieval operation binding the contract event 0xbbf93173200b07feb5b440c63ea96d13e6fe53f2550e2a359eeb993f1d9f5238.
//
// Solidity: e LogComplete()
func (_TournamentPool *TournamentPoolFilterer) FilterLogComplete(opts *bind.FilterOpts) (*TournamentPoolLogCompleteIterator, error) {

	logs, sub, err := _TournamentPool.contract.FilterLogs(opts, "LogComplete")
	if err != nil {
		return nil, err
	}
	return &TournamentPoolLogCompleteIterator{contract: _TournamentPool.contract, event: "LogComplete", logs: logs, sub: sub}, nil
}

// WatchLogComplete is a free log subscription operation binding the contract event 0xbbf93173200b07feb5b440c63ea96d13e6fe53f2550e2a359eeb993f1d9f5238.
//
// Solidity: e LogComplete()
func (_TournamentPool *TournamentPoolFilterer) WatchLogComplete(opts *bind.WatchOpts, sink chan<- *TournamentPoolLogComplete) (event.Subscription, error) {

	logs, sub, err := _TournamentPool.contract.WatchLogs(opts, "LogComplete")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TournamentPoolLogComplete)
				if err := _TournamentPool.contract.UnpackLog(event, "LogComplete", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TournamentPoolLogPayoutIterator is returned from FilterLogPayout and is used to iterate over the raw logs and unpacked data for LogPayout events raised by the TournamentPool contract.
type TournamentPoolLogPayoutIterator struct {
	Event *TournamentPoolLogPayout // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TournamentPoolLogPayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TournamentPoolLogPayout)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TournamentPoolLogPayout)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TournamentPoolLogPayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TournamentPoolLogPayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TournamentPoolLogPayout represents a LogPayout event raised by the TournamentPool contract.
type TournamentPoolLogPayout struct {
	Participant common.Address
	Payout      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogPayout is a free log retrieval operation binding the contract event 0x7d6963073a93103d7a336ea21c7f029ad757c9ef4c0c6ad0ff86c689480e0ed8.
//
// Solidity: e LogPayout(participant address, payout uint256)
func (_TournamentPool *TournamentPoolFilterer) FilterLogPayout(opts *bind.FilterOpts) (*TournamentPoolLogPayoutIterator, error) {

	logs, sub, err := _TournamentPool.contract.FilterLogs(opts, "LogPayout")
	if err != nil {
		return nil, err
	}
	return &TournamentPoolLogPayoutIterator{contract: _TournamentPool.contract, event: "LogPayout", logs: logs, sub: sub}, nil
}

// WatchLogPayout is a free log subscription operation binding the contract event 0x7d6963073a93103d7a336ea21c7f029ad757c9ef4c0c6ad0ff86c689480e0ed8.
//
// Solidity: e LogPayout(participant address, payout uint256)
func (_TournamentPool *TournamentPoolFilterer) WatchLogPayout(opts *bind.WatchOpts, sink chan<- *TournamentPoolLogPayout) (event.Subscription, error) {

	logs, sub, err := _TournamentPool.contract.WatchLogs(opts, "LogPayout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TournamentPoolLogPayout)
				if err := _TournamentPool.contract.UnpackLog(event, "LogPayout", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TournamentPoolLogRefundIterator is returned from FilterLogRefund and is used to iterate over the raw logs and unpacked data for LogRefund events raised by the TournamentPool contract.
type TournamentPoolLogRefundIterator struct {
	Event *TournamentPoolLogRefund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TournamentPoolLogRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TournamentPoolLogRefund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TournamentPoolLogRefund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TournamentPoolLogRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TournamentPoolLogRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TournamentPoolLogRefund represents a LogRefund event raised by the TournamentPool contract.
type TournamentPoolLogRefund struct {
	Backer common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogRefund is a free log retrieval operation binding the contract event 0xb6c0eca8138e097d71e2dd31e19a1266487f0553f170b7260ffe68bcbe9ff8a7.
//
// Solidity: e LogRefund(backer address, amount uint256)
func (_TournamentPool *TournamentPoolFilterer) FilterLogRefund(opts *bind.FilterOpts) (*TournamentPoolLogRefundIterator, error) {

	logs, sub, err := _TournamentPool.contract.FilterLogs(opts, "LogRefund")
	if err != nil {
		return nil, err
	}
	return &TournamentPoolLogRefundIterator{contract: _TournamentPool.contract, event: "LogRefund", logs: logs, sub: sub}, nil
}

// WatchLogRefund is a free log subscription operation binding the contract event 0xb6c0eca8138e097d71e2dd31e19a1266487f0553f170b7260ffe68bcbe9ff8a7.
//
// Solidity: e LogRefund(backer address, amount uint256)
func (_TournamentPool *TournamentPoolFilterer) WatchLogRefund(opts *bind.WatchOpts, sink chan<- *TournamentPoolLogRefund) (event.Subscription, error) {

	logs, sub, err := _TournamentPool.contract.WatchLogs(opts, "LogRefund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TournamentPoolLogRefund)
				if err := _TournamentPool.contract.UnpackLog(event, "LogRefund", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TournamentPoolLogRegistrationIterator is returned from FilterLogRegistration and is used to iterate over the raw logs and unpacked data for LogRegistration events raised by the TournamentPool contract.
type TournamentPoolLogRegistrationIterator struct {
	Event *TournamentPoolLogRegistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TournamentPoolLogRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TournamentPoolLogRegistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TournamentPoolLogRegistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TournamentPoolLogRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TournamentPoolLogRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TournamentPoolLogRegistration represents a LogRegistration event raised by the TournamentPool contract.
type TournamentPoolLogRegistration struct {
	Participant common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogRegistration is a free log retrieval operation binding the contract event 0x4c2cd98156e6a27c8f70b5432e1dc7cee6bc43471c4bb8446cadbd44fbfbf2ba.
//
// Solidity: e LogRegistration(participant address)
func (_TournamentPool *TournamentPoolFilterer) FilterLogRegistration(opts *bind.FilterOpts) (*TournamentPoolLogRegistrationIterator, error) {

	logs, sub, err := _TournamentPool.contract.FilterLogs(opts, "LogRegistration")
	if err != nil {
		return nil, err
	}
	return &TournamentPoolLogRegistrationIterator{contract: _TournamentPool.contract, event: "LogRegistration", logs: logs, sub: sub}, nil
}

// WatchLogRegistration is a free log subscription operation binding the contract event 0x4c2cd98156e6a27c8f70b5432e1dc7cee6bc43471c4bb8446cadbd44fbfbf2ba.
//
// Solidity: e LogRegistration(participant address)
func (_TournamentPool *TournamentPoolFilterer) WatchLogRegistration(opts *bind.WatchOpts, sink chan<- *TournamentPoolLogRegistration) (event.Subscription, error) {

	logs, sub, err := _TournamentPool.contract.WatchLogs(opts, "LogRegistration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TournamentPoolLogRegistration)
				if err := _TournamentPool.contract.UnpackLog(event, "LogRegistration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TournamentPoolLogUnregisterIterator is returned from FilterLogUnregister and is used to iterate over the raw logs and unpacked data for LogUnregister events raised by the TournamentPool contract.
type TournamentPoolLogUnregisterIterator struct {
	Event *TournamentPoolLogUnregister // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TournamentPoolLogUnregisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TournamentPoolLogUnregister)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TournamentPoolLogUnregister)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TournamentPoolLogUnregisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TournamentPoolLogUnregisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TournamentPoolLogUnregister represents a LogUnregister event raised by the TournamentPool contract.
type TournamentPoolLogUnregister struct {
	Participant common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogUnregister is a free log retrieval operation binding the contract event 0x11854d1b3c0aa24c7c879af700c0089a48a48e9280bac11f5370b90b7cca481c.
//
// Solidity: e LogUnregister(participant address)
func (_TournamentPool *TournamentPoolFilterer) FilterLogUnregister(opts *bind.FilterOpts) (*TournamentPoolLogUnregisterIterator, error) {

	logs, sub, err := _TournamentPool.contract.FilterLogs(opts, "LogUnregister")
	if err != nil {
		return nil, err
	}
	return &TournamentPoolLogUnregisterIterator{contract: _TournamentPool.contract, event: "LogUnregister", logs: logs, sub: sub}, nil
}

// WatchLogUnregister is a free log subscription operation binding the contract event 0x11854d1b3c0aa24c7c879af700c0089a48a48e9280bac11f5370b90b7cca481c.
//
// Solidity: e LogUnregister(participant address)
func (_TournamentPool *TournamentPoolFilterer) WatchLogUnregister(opts *bind.WatchOpts, sink chan<- *TournamentPoolLogUnregister) (event.Subscription, error) {

	logs, sub, err := _TournamentPool.contract.WatchLogs(opts, "LogUnregister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TournamentPoolLogUnregister)
				if err := _TournamentPool.contract.UnpackLog(event, "LogUnregister", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
