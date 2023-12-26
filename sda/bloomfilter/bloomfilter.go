package bloomfilter

import (
	"hash"
)

type BloomFilter struct {
	Bytes  []byte
	Hashes []hash.Hash64
}

func (bf *BloomFilter) Push(b []byte) {
	byteLen := len(bf.Bytes)
	for _, f := range bf.Hashes {
		f.Reset()
		f.Write(b)
		res := f.Sum64()
		row := res % uint64(byteLen)
		col := res & 7
		now := bf.Bytes[row] | 1<<col
		if now != bf.Bytes[row] {
			bf.Bytes[row] = now
		}
	}
}

func (bf *BloomFilter) Existed(b []byte) bool {
	byteLen := len(bf.Bytes)
	for _, f := range bf.Hashes {
		f.Reset()
		f.Write(b)
		res := f.Sum64()
		row := res % uint64(byteLen)
		col := res & 7
		if bf.Bytes[row]|1<<col != bf.Bytes[row] {
			return false
		}
	}
	return true
}
