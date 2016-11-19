package crypto

import (
	"encoding/hex"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&CryptoSuite{})

type CryptoSuite struct{}

func (s *CryptoSuite) SetUpSuite(c *C) {
}

func (s *CryptoSuite) TestCreateKeyBlock(c *C) {
	identity := []byte("Identity")
	psk, _ := hex.DecodeString("0011223344")
	cr, _ := hex.DecodeString("00000001E68D63E65CDEF492AA9877330CA7EEB5C786487F31DAE89452104156")
	sr, _ := hex.DecodeString("5823185CF999576643D2E838C4FCAEAC6AE89C12C9E25517D95FE115C50BF080")

	kb, _ := CreateKeyBlock(identity, psk, cr, sr)

	c.Assert(hex.EncodeToString(kb.MasterSecret), Equals, "1558b8bfc82eb58fdc576021312662f208e68a9e7ddec50d5f47a1841672d4d268e46bebe5a2954d63f1cda811df4faf")
	c.Assert(hex.EncodeToString(kb.ClientWriteKey), Equals, "ce3db22cd0931c3e752176b43eb1939a")
	c.Assert(hex.EncodeToString(kb.ServerWriteKey), Equals, "3501b84cf54e2654090e82abefcdccbd")
	c.Assert(hex.EncodeToString(kb.ClientIV), Equals, "f21ce4e5")
	c.Assert(hex.EncodeToString(kb.ServerIV), Equals, "235a077a")
}

func (s *CryptoSuite) TestNonce(c *C) {
	iv, _ := hex.DecodeString("F21CE4E5")
	nonce := CreateNonce(iv, 1, 0)
	c.Assert(hex.EncodeToString(nonce), Equals, "f21ce4e50001000000000000")
}

func (s *CryptoSuite) TestGeneratePrf(c *C) {
	ms, _ := hex.DecodeString("20A8A0E9172B0F7A1F370CF082B2FAD79BBC5F0757452B176695124960074985ED9D444A5D188D3397C74B3277EB1B0F")
	hash, _ := hex.DecodeString("777DBBC320A905D5BA76AD9323A986256991DD99FCDD265F4202A39C12C87F8F")

	prf := GeneratePrf(ms, []byte(" finished"), hash, "client", 12)

	c.Assert(hex.EncodeToString(prf), Equals, "8c6ccb73e751bd05b36636af")
}