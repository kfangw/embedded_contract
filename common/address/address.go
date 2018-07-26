package address

import (
	"errors"
	"io"
)

const ADDR_LEN = 20

type Address [ADDR_LEN]byte

// Serialize serialize Address into io.Writer
func (addr *Address) Serialize(w io.Writer) error {
	_, err := w.Write(addr[:])
	return err
}

// Deserialize deserialize Address from io.Reader
func (addr *Address) Deserialize(r io.Reader) error {
	_, err := io.ReadFull(r, addr[:])
	if err != nil {
		return errors.New("deserialize Address error")
	}
	return nil
}

func AddressFromString(addrStr string) Address {
	copylen := len(addrStr)
	// Temporary code needed to fix
	if copylen > ADDR_LEN {
		copylen = ADDR_LEN
	}
	var addr Address
	copy(addr[:], addrStr[:copylen])
	return addr
}
