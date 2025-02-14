package walletcontroller

import (
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	notifier "github.com/lightningnetwork/lnd/chainntnfs"
)

type TxStatus int

const (
	TxNotFound TxStatus = iota
	TxInMemPool
	TxInChain
)

type WalletController interface {
	UnlockWallet(timeoutSecs int64) error
	AddressPublicKey(address btcutil.Address) (*btcec.PublicKey, error)
	DumpPrivateKey(address btcutil.Address) (*btcec.PrivateKey, error)
	ImportPrivKey(privKeyWIF *btcutil.WIF) error
	NetworkName() string
	CreateTransaction(
		outputs []*wire.TxOut,
		feeRatePerKb btcutil.Amount,
		changeScript btcutil.Address) (*wire.MsgTx, error)
	SignRawTransaction(tx *wire.MsgTx) (*wire.MsgTx, bool, error)
	// requires wallet to be unlocked
	CreateAndSignTx(
		output []*wire.TxOut,
		feeRatePerKb btcutil.Amount,
		changeAddress btcutil.Address,
	) (*wire.MsgTx, error)
	SendRawTransaction(tx *wire.MsgTx, allowHighFees bool) (*chainhash.Hash, error)
	ListOutputs(onlySpendable bool) ([]Utxo, error)
	TxDetails(txHash *chainhash.Hash, pkScript []byte) (*notifier.TxConfirmation, TxStatus, error)
	SignBip322NativeSegwit(msg []byte, address btcutil.Address) (wire.TxWitness, error)
}
