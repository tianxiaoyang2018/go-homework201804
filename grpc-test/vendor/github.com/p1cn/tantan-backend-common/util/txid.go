package util

import (
	"bytes"
	"encoding/base32"
	"encoding/binary"
	"math/rand"
	"strings"
	"time"
)

var txidRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// NewTxid generates a txid in the same way as the Nginx module txid.
//
// https://github.com/streadway/ngx_txid
//
// It uses a 96 bit binary formatted encoded as a base32 string.
// The 96 bits are used in this way:
// +------------- 64 bits------------+--- 32 bits ----+
// +------ 42 bits ------+--22 bits--|----------------+
// | msec since 1970-1-1 | random    | random         |
// +---------------------+-----------+----------------+
//
// In case there is an error in the function (highly unlikely) parts or all
// of the generated id might be set to 0.
func NewTxid() string {
	bs := make([]byte, 12)

	now := time.Now().UnixNano() / int64(time.Millisecond)
	now = now << 22

	buf := &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, now)

	copy(bs, buf.Bytes())

	randBs := make([]byte, 7)
	txidRand.Read(randBs)

	// Fill "half" of byte 5 and bytes 6-11 in buf with random bytes.
	bs[5] |= randBs[5] >> 2
	copy(bs[6:12], randBs[1:7])

	// Encode to base32
	s := base32.HexEncoding.EncodeToString(bs)
	s = strings.ToLower(s)
	s = strings.TrimRight(s, "=")
	return s
}
