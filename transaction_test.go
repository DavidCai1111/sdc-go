package sdcgo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TransactionSuite struct {
	suite.Suite
	txinput  TXInput
	txoutput TXOutput
}

func (s *TransactionSuite) SetupTest() {
	s.txinput = TXInput{ScriptSig: "test_ScriptSig"}
	s.txoutput = TXOutput{ScriptPubKey: "test_ScriptPubKey"}
}

func (s *TransactionSuite) TestCanUnlockOutputWith() {
	s.True(s.txinput.CanUnlockOutputWith("test_ScriptSig"))
}

func (s *TransactionSuite) TestCanBeUnlockedWith() {
	s.True(s.txoutput.CanBeUnlockedWith("test_ScriptPubKey"))
}

func TestTransaction(t *testing.T) {
	suite.Run(t, new(TransactionSuite))
}
