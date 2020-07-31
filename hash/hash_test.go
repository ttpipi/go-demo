package hash

import (
	"crypto/md5"
	"hash/fnv"
	"testing"
)

var testData string = "YysUAY1DxCPajnnsZj2R"

func TestMd5(t *testing.T) {
	f := md5.New()
	f.Reset()
	f.Write([]byte(testData))
	res := f.Sum(nil)
	t.Log(string(res))
}

func TestFnv(t *testing.T) {
	f := fnv.New64()
	f.Reset()
	f.Write([]byte(testData))
	res := f.Sum64()
	t.Log(res)
}
