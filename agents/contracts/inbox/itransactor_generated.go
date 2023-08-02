// autogenerated file

package inbox

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IInboxTransactor ...
type IInboxTransactor interface {
	// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
	//
	// Solidity: function initialize(address agentManager_, address origin_, address destination_, address summit_) returns()
	Initialize(opts *bind.TransactOpts, agentManager_ common.Address, origin_ common.Address, destination_ common.Address, summit_ common.Address) (*types.Transaction, error)
	// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
	//
	// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
	Multicall(opts *bind.TransactOpts, calls []MultiCallableCall) (*types.Transaction, error)
	// PassReceipt is a paid mutator transaction binding the contract method 0x6b47b3bc.
	//
	// Solidity: function passReceipt(uint32 attNotaryIndex, uint32 attNonce, uint256 paddedTips, bytes rcptPayload) returns(bool wasAccepted)
	PassReceipt(opts *bind.TransactOpts, attNotaryIndex uint32, attNonce uint32, paddedTips *big.Int, rcptPayload []byte) (*types.Transaction, error)
	// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
	//
	// Solidity: function renounceOwnership() returns()
	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)
	// SubmitReceipt is a paid mutator transaction binding the contract method 0xb2a4b455.
	//
	// Solidity: function submitReceipt(bytes rcptPayload, bytes rcptSignature, uint256 paddedTips, bytes32 headerHash, bytes32 bodyHash) returns(bool wasAccepted)
	SubmitReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte, paddedTips *big.Int, headerHash [32]byte, bodyHash [32]byte) (*types.Transaction, error)
	// SubmitReceiptReport is a paid mutator transaction binding the contract method 0x89246503.
	//
	// Solidity: function submitReceiptReport(bytes rcptPayload, bytes rcptSignature, bytes rrSignature) returns(bool wasAccepted)
	SubmitReceiptReport(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte, rrSignature []byte) (*types.Transaction, error)
	// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
	//
	// Solidity: function submitSnapshot(bytes snapPayload, bytes snapSignature) returns(bytes attPayload, bytes32 agentRoot_, uint256[] snapGas)
	SubmitSnapshot(opts *bind.TransactOpts, snapPayload []byte, snapSignature []byte) (*types.Transaction, error)
	// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x0b6b985c.
	//
	// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
	SubmitStateReportWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error)
	// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x62389709.
	//
	// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
	SubmitStateReportWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error)
	// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
	//
	// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
	SubmitStateReportWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error)
	// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
	//
	// Solidity: function transferOwnership(address newOwner) returns()
	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)
	// VerifyAttestation is a paid mutator transaction binding the contract method 0x0ca77473.
	//
	// Solidity: function verifyAttestation(bytes attPayload, bytes attSignature) returns(bool isValidAttestation)
	VerifyAttestation(opts *bind.TransactOpts, attPayload []byte, attSignature []byte) (*types.Transaction, error)
	// VerifyAttestationReport is a paid mutator transaction binding the contract method 0x31e8df5a.
	//
	// Solidity: function verifyAttestationReport(bytes attPayload, bytes arSignature) returns(bool isValidReport)
	VerifyAttestationReport(opts *bind.TransactOpts, attPayload []byte, arSignature []byte) (*types.Transaction, error)
	// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
	//
	// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
	VerifyReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error)
	// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
	//
	// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
	VerifyReceiptReport(opts *bind.TransactOpts, rcptPayload []byte, rrSignature []byte) (*types.Transaction, error)
	// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
	//
	// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
	VerifyStateReport(opts *bind.TransactOpts, statePayload []byte, srSignature []byte) (*types.Transaction, error)
	// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
	//
	// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
	VerifyStateWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error)
	// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
	//
	// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
	VerifyStateWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error)
	// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
	//
	// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
	VerifyStateWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error)
}