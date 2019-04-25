
// Code generated by typeRegistryGenerator.go. DO NOT EDIT.

package kkProto

import (
	"reflect"
)

func init() {
	typeRegistry["Initialize"] = reflect.TypeOf(Initialize{})
	typeRegistry["GetFeatures"] = reflect.TypeOf(GetFeatures{})
	typeRegistry["Features"] = reflect.TypeOf(Features{})
	typeRegistry["GetCoinTable"] = reflect.TypeOf(GetCoinTable{})
	typeRegistry["CoinTable"] = reflect.TypeOf(CoinTable{})
	typeRegistry["ClearSession"] = reflect.TypeOf(ClearSession{})
	typeRegistry["ApplySettings"] = reflect.TypeOf(ApplySettings{})
	typeRegistry["ChangePin"] = reflect.TypeOf(ChangePin{})
	typeRegistry["Ping"] = reflect.TypeOf(Ping{})
	typeRegistry["Success"] = reflect.TypeOf(Success{})
	typeRegistry["Failure"] = reflect.TypeOf(Failure{})
	typeRegistry["ButtonRequest"] = reflect.TypeOf(ButtonRequest{})
	typeRegistry["ButtonAck"] = reflect.TypeOf(ButtonAck{})
	typeRegistry["PinMatrixRequest"] = reflect.TypeOf(PinMatrixRequest{})
	typeRegistry["PinMatrixAck"] = reflect.TypeOf(PinMatrixAck{})
	typeRegistry["Cancel"] = reflect.TypeOf(Cancel{})
	typeRegistry["PassphraseRequest"] = reflect.TypeOf(PassphraseRequest{})
	typeRegistry["PassphraseAck"] = reflect.TypeOf(PassphraseAck{})
	typeRegistry["GetEntropy"] = reflect.TypeOf(GetEntropy{})
	typeRegistry["Entropy"] = reflect.TypeOf(Entropy{})
	typeRegistry["GetPublicKey"] = reflect.TypeOf(GetPublicKey{})
	typeRegistry["PublicKey"] = reflect.TypeOf(PublicKey{})
	typeRegistry["GetAddress"] = reflect.TypeOf(GetAddress{})
	typeRegistry["EthereumGetAddress"] = reflect.TypeOf(EthereumGetAddress{})
	typeRegistry["Address"] = reflect.TypeOf(Address{})
	typeRegistry["EthereumAddress"] = reflect.TypeOf(EthereumAddress{})
	typeRegistry["WipeDevice"] = reflect.TypeOf(WipeDevice{})
	typeRegistry["LoadDevice"] = reflect.TypeOf(LoadDevice{})
	typeRegistry["ResetDevice"] = reflect.TypeOf(ResetDevice{})
	typeRegistry["EntropyRequest"] = reflect.TypeOf(EntropyRequest{})
	typeRegistry["EntropyAck"] = reflect.TypeOf(EntropyAck{})
	typeRegistry["RecoveryDevice"] = reflect.TypeOf(RecoveryDevice{})
	typeRegistry["WordRequest"] = reflect.TypeOf(WordRequest{})
	typeRegistry["WordAck"] = reflect.TypeOf(WordAck{})
	typeRegistry["CharacterRequest"] = reflect.TypeOf(CharacterRequest{})
	typeRegistry["CharacterAck"] = reflect.TypeOf(CharacterAck{})
	typeRegistry["SignMessage"] = reflect.TypeOf(SignMessage{})
	typeRegistry["VerifyMessage"] = reflect.TypeOf(VerifyMessage{})
	typeRegistry["MessageSignature"] = reflect.TypeOf(MessageSignature{})
	typeRegistry["EncryptMessage"] = reflect.TypeOf(EncryptMessage{})
	typeRegistry["EncryptedMessage"] = reflect.TypeOf(EncryptedMessage{})
	typeRegistry["DecryptMessage"] = reflect.TypeOf(DecryptMessage{})
	typeRegistry["DecryptedMessage"] = reflect.TypeOf(DecryptedMessage{})
	typeRegistry["CipherKeyValue"] = reflect.TypeOf(CipherKeyValue{})
	typeRegistry["CipheredKeyValue"] = reflect.TypeOf(CipheredKeyValue{})
	typeRegistry["EstimateTxSize"] = reflect.TypeOf(EstimateTxSize{})
	typeRegistry["TxSize"] = reflect.TypeOf(TxSize{})
	typeRegistry["SignTx"] = reflect.TypeOf(SignTx{})
	typeRegistry["TxRequest"] = reflect.TypeOf(TxRequest{})
	typeRegistry["TxAck"] = reflect.TypeOf(TxAck{})
	typeRegistry["RawTxAck"] = reflect.TypeOf(RawTxAck{})
	typeRegistry["EthereumSignTx"] = reflect.TypeOf(EthereumSignTx{})
	typeRegistry["EthereumTxRequest"] = reflect.TypeOf(EthereumTxRequest{})
	typeRegistry["EthereumTxAck"] = reflect.TypeOf(EthereumTxAck{})
	typeRegistry["EthereumSignMessage"] = reflect.TypeOf(EthereumSignMessage{})
	typeRegistry["EthereumVerifyMessage"] = reflect.TypeOf(EthereumVerifyMessage{})
	typeRegistry["EthereumMessageSignature"] = reflect.TypeOf(EthereumMessageSignature{})
	typeRegistry["SignIdentity"] = reflect.TypeOf(SignIdentity{})
	typeRegistry["SignedIdentity"] = reflect.TypeOf(SignedIdentity{})
	typeRegistry["ApplyPolicies"] = reflect.TypeOf(ApplyPolicies{})
	typeRegistry["FlashHash"] = reflect.TypeOf(FlashHash{})
	typeRegistry["FlashWrite"] = reflect.TypeOf(FlashWrite{})
	typeRegistry["FlashHashResponse"] = reflect.TypeOf(FlashHashResponse{})
	typeRegistry["DebugLinkFlashDump"] = reflect.TypeOf(DebugLinkFlashDump{})
	typeRegistry["DebugLinkFlashDumpResponse"] = reflect.TypeOf(DebugLinkFlashDumpResponse{})
	typeRegistry["SoftReset"] = reflect.TypeOf(SoftReset{})
	typeRegistry["FirmwareErase"] = reflect.TypeOf(FirmwareErase{})
	typeRegistry["FirmwareUpload"] = reflect.TypeOf(FirmwareUpload{})
	typeRegistry["DebugLinkDecision"] = reflect.TypeOf(DebugLinkDecision{})
	typeRegistry["DebugLinkGetState"] = reflect.TypeOf(DebugLinkGetState{})
	typeRegistry["DebugLinkState"] = reflect.TypeOf(DebugLinkState{})
	typeRegistry["DebugLinkStop"] = reflect.TypeOf(DebugLinkStop{})
	typeRegistry["DebugLinkLog"] = reflect.TypeOf(DebugLinkLog{})
	typeRegistry["DebugLinkFillConfig"] = reflect.TypeOf(DebugLinkFillConfig{})
	typeRegistry["HDNodeType"] = reflect.TypeOf(HDNodeType{})
	typeRegistry["HDNodePathType"] = reflect.TypeOf(HDNodePathType{})
	typeRegistry["CoinType"] = reflect.TypeOf(CoinType{})
	typeRegistry["MultisigRedeemScriptType"] = reflect.TypeOf(MultisigRedeemScriptType{})
	typeRegistry["TxInputType"] = reflect.TypeOf(TxInputType{})
	typeRegistry["TxOutputType"] = reflect.TypeOf(TxOutputType{})
	typeRegistry["TxOutputBinType"] = reflect.TypeOf(TxOutputBinType{})
	typeRegistry["TransactionType"] = reflect.TypeOf(TransactionType{})
	typeRegistry["RawTransactionType"] = reflect.TypeOf(RawTransactionType{})
	typeRegistry["TxRequestDetailsType"] = reflect.TypeOf(TxRequestDetailsType{})
	typeRegistry["TxRequestSerializedType"] = reflect.TypeOf(TxRequestSerializedType{})
	typeRegistry["IdentityType"] = reflect.TypeOf(IdentityType{})
	typeRegistry["PolicyType"] = reflect.TypeOf(PolicyType{})
	typeRegistry["ExchangeType"] = reflect.TypeOf(ExchangeType{})
	typeRegistry["ExchangeAddress"] = reflect.TypeOf(ExchangeAddress{})
	typeRegistry["ExchangeResponseV2"] = reflect.TypeOf(ExchangeResponseV2{})
	typeRegistry["SignedExchangeResponse"] = reflect.TypeOf(SignedExchangeResponse{})
	typeRegistry["ExchangeResponse"] = reflect.TypeOf(ExchangeResponse{})
	typeRegistry["Storage"] = reflect.TypeOf(Storage{})
}

var typeRegistry = make(map[string]reflect.Type)

// TypeRegistry returns the reflect.Type associated with a given string type-name
// The entries registered in the map should be structs implementing proto.Message
func TypeRegistry(t string) (reflect.Type, bool) {
	v, ok := typeRegistry[t]
	return v, ok
}
