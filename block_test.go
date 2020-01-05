package sdcgo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BlockSuite struct {
	suite.Suite
	b Block
}

func (s *BlockSuite) SetupTest() {
	s.b = Block{Nonce: 1}
}

func (s *BlockSuite) TestDeserializeBlock() {
	b, err := s.b.Serialize()

	s.Nil(err)

	block, err := DeserializeBlock(b)

	s.Nil(err)

	s.Equal(1, block.Nonce)
}

func TestBlock(t *testing.T) {
	suite.Run(t, new(BlockSuite))
}
