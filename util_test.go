package sbc

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UtilSuite struct {
	suite.Suite
}

func (s *UtilSuite) TestSerialize() {
	b, err := Serialize("test_Serialize()")

	s.Nil(err)
	s.Len(b, 20)
}

func TestUtil(t *testing.T) {
	suite.Run(t, new(UtilSuite))
}
