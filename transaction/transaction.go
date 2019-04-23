package transaction

import (
	"encoding/binary"
	"encoding/hex"
	"io"

	"spamp2p/log"

	"github.com/KickSeason/neo-go-sdk/neotransaction"
)

type TxWrapper struct {
	NTx *neotransaction.NeoTransaction
}

var logger = log.NewLogger("transaction")

func (t TxWrapper) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, t.NTx.RawTransaction()); err != nil {
		return err
	}
	return nil
}
func (t TxWrapper) DecodeBinary(r io.Reader) error {
	return nil
}

func InvokeTx() TxWrapper {
	var (
		contractHashString = `ecf33479dadde66b721a0791ac03e3d06bb137ab`
		addrReceiverString = "ARUB61ysG7LLjuChqfjTBUNg6UB6gj4Mat"
	)
	contractHash, _ := hex.DecodeString(contractHashString)
	key, _ := neotransaction.DecodeFromWif("private-key")
	addr := key.CreateBasicAddress()
	addrReceiver, _ := neotransaction.ParseAddress(addrReceiverString)
	tx := neotransaction.CreateInvocationTransaction()
	extra := tx.ExtraData.(*neotransaction.InvocationExtraData)
	args := []interface{}{
		addr.ScripHash,
		addrReceiver.ScripHash,
		int16(1),
	}
	bytes, err := neotransaction.BuildCallMethodScript(contractHash, "transfer", args, true)
	if err != nil {
		logger.Println(err)
	}
	extra.Script = bytes
	tx.AppendAttribute(neotransaction.UsageScript, addr.ScripHash)
	tx.AppendBasicSignWitness(key)
	//logger.Printf("[transaction] make transaction. txid: %s\n", tx.TXID())
	return TxWrapper{
		NTx: tx,
	}
}
