// h 20181015
//
// DrvLot.IES

package ies

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/ethereum/go-ethereum/crypto/ecies"
)

func Enc(pub string, tkn, msg []byte) string {
	ret := ""
	for {
		_pub, err := hex.DecodeString(pub)
		if err != nil {
			break
		}
		ctt, err := Encrypt(rand.Reader, ImportECDSAPublic(crypto.ToECDSAPub(_pub)), msg, nil, tkn)
		if err != nil {
			break
		} else {
			ret = hex.EncodeToString(ctt)
		}
		//
		// Finally
		if true {
			break
		}
	}
	return ret
}
