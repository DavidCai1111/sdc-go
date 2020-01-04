package sdcgo

// Transaction is a transaction in the chain.
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

// TXInput is the input of a transaction.
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

// CanUnlockOutputWith checks whether the input can unlock the output.
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

// TXOutput is the output of a transaction.
type TXOutput struct {
	Value        int
	ScriptPubKey string
}

// CanBeUnlockedWith checks whether the output can be unlocked.
func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}
