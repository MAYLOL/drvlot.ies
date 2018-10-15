// h 20181015
//
// DrvLot.IES

package ies

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println(enc("0471755c96a787417ef214e8af106c25e9de89f95a2b152b641e7d40531819ab5a53c9381042911fed71f0b67623d242f37c79d51a3b863d1628fc3e83ea249696", []byte("hello"), []byte("world")))
}
