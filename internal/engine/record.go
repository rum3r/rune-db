package engine

import (
	"encoding/binary"
	"io"
)

// opType marks what a record does: store a value or delete a key.
type opType byte

const (
	opPut    opType = 0
	opDelete opType = 1
)

// headerSize is the fixed prefix on every record on disk:
// 1 byte op + 4 bytes key length + 4 bytes value length.
const headerSize = 1 + 4 + 4

// record is a single entry in the append-only log.
type record struct {
	op    opType
	key   string
	value string
}

// encode serialises the record into its on-disk layout:
// [op][keyLen][valLen][key][value].
func (r record) encode() []byte {
	key := []byte(r.key)
	value := []byte(r.value)

	buf := make([]byte, headerSize+len(key)+len(value))
	buf[0] = byte(r.op)
	binary.BigEndian.PutUint32(buf[1:5], uint32(len(key)))
	binary.BigEndian.PutUint32(buf[5:9], uint32(len(value)))
	copy(buf[headerSize:], key)
	copy(buf[headerSize+len(key):], value)
	return buf
}

// decodeRecord reads one record from r. It also returns the total number of
// bytes the record occupied on disk, so callers can track file offsets while
// scanning the log.
func decodeRecord(r io.Reader) (rec record, size int, err error) {
	header := make([]byte, headerSize)
	if _, err = io.ReadFull(r, header); err != nil {
		return record{}, 0, err
	}

	keyLen := binary.BigEndian.Uint32(header[1:5])
	valLen := binary.BigEndian.Uint32(header[5:9])

	body := make([]byte, keyLen+valLen)
	if _, err = io.ReadFull(r, body); err != nil {
		return record{}, 0, err
	}

	rec = record{
		op:    opType(header[0]),
		key:   string(body[:keyLen]),
		value: string(body[keyLen:]),
	}
	return rec, headerSize + int(keyLen) + int(valLen), nil
}
