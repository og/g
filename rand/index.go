package grand

import (
	"crypto/rand"
	"math/big"
)
func RunesBySpeed(seed []rune, size int ) []rune {
	var result  []rune
	for i:=0; i<size; i++ {
		randIndex, err :=rand.Int(rand.Reader, big.NewInt(int64(len(seed)))) ; if err !=nil {panic(err)}
		result = append(result, seed[randIndex.Int64()])
	}
	return result
}
func BytesBySeed(seed []byte, size int) []byte {
	result := []byte("")
	for i:=0; i<size; i++ {
		randIndex, err :=rand.Int(rand.Reader, big.NewInt(int64(len(seed)))) ; if err !=nil {panic(err)}
		result = append(result, seed[randIndex.Int64()])
	}
	return result
}
func StringLetter (size int) string {
	return string(BytesLetter(size))
}
func letterBytes() []byte {
	return []byte{0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7a}
}


func BytesLetter (size int) []byte {
	return BytesBySeed(letterBytes(), size)
}
