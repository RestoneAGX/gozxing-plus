package decoder

import (
	"reflect"
	"testing"
)

func testDataBlocks(t *testing.T, block *DataBlock, expectnum int, expectwords []byte) {
	if n := block.GetNumDataCodewords(); n != expectnum {
		t.Fatalf("NumDataCodeWords = %v, expect %v", n, expectnum)
	}
	if words := block.GetCodewords(); !reflect.DeepEqual(words, expectwords) {
		t.Fatalf("CodeWords = \n%v\nexpect:\n%v", words, expectwords)
	}
}

func TestDataBlock_GetDataBlocks(t *testing.T) {
	version, _ := Version_GetVersionForNumber(5)
	ecLeve := ErrorCorrectionLevel_H
	words := make([]byte, version.GetTotalCodewords())
	for i := range words {
		words[i] = byte(i % 256)
	}

	_, e := DataBlock_GetDataBlocks([]byte{}, version, ecLeve)
	if e == nil {
		t.Fatalf("Getdatablocks must be error")
	}

	num0 := 11
	words0 := []byte{
		0x00, 0x04, 0x08, 0x0c, 0x10, 0x14, 0x18, 0x1c, 0x20, 0x24, 0x28,
		0x2e, 0x32, 0x36, 0x3a, 0x3e, 0x42, 0x46, 0x4a, 0x4e, 0x52, 0x56,
		0x5a, 0x5e, 0x62, 0x66, 0x6a, 0x6e, 0x72, 0x76, 0x7a, 0x7e, 0x82,
	}
	num1 := 11
	words1 := []byte{
		0x01, 0x05, 0x09, 0x0d, 0x11, 0x15, 0x19, 0x1d, 0x21, 0x25, 0x29,
		0x2f, 0x33, 0x37, 0x3b, 0x3f, 0x43, 0x47, 0x4b, 0x4f, 0x53, 0x57,
		0x5b, 0x5f, 0x63, 0x67, 0x6b, 0x6f, 0x73, 0x77, 0x7b, 0x7f, 0x83,
	}
	num2 := 12
	words2 := []byte{
		0x02, 0x06, 0x0a, 0x0e, 0x12, 0x16, 0x1a, 0x1e, 0x22, 0x26, 0x2a,
		0x2c, 0x30, 0x34, 0x38, 0x3c, 0x40, 0x44, 0x48, 0x4c, 0x50, 0x54,
		0x58, 0x5c, 0x60, 0x64, 0x68, 0x6c, 0x70, 0x74, 0x78, 0x7c, 0x80, 0x84,
	}
	num3 := 12
	words3 := []byte{
		0x03, 0x07, 0x0b, 0x0f, 0x13, 0x17, 0x1b, 0x1f, 0x23, 0x27, 0x2b,
		0x2d, 0x31, 0x35, 0x39, 0x3d, 0x41, 0x45, 0x49, 0x4d, 0x51, 0x55,
		0x59, 0x5d, 0x61, 0x65, 0x69, 0x6d, 0x71, 0x75, 0x79, 0x7d, 0x81, 0x85,
	}

	blocks, e := DataBlock_GetDataBlocks(words, version, ecLeve)
	if e != nil {
		t.Fatalf("Getdatablocks returns error, %v", e)
	}
	if l := len(blocks); l != 4 {
		t.Fatalf("len(blocks) = %v, expect 4", l)
	}
	testDataBlocks(t, blocks[0], num0, words0)
	testDataBlocks(t, blocks[1], num1, words1)
	testDataBlocks(t, blocks[2], num2, words2)
	testDataBlocks(t, blocks[3], num3, words3)
}
