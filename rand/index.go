package grand

import (
	"crypto/rand"
	"math/big"
	"time"
)

func StringBySeed(seed string, size int) string {
	seedB := []byte(seed)
	result := []byte()
	for i:=0; i<size; i++ {
		randIndex, err :=rand.Int(rand.Reader, big.NewInt(int64(len(seed)))) ; if err !=nil {panic(err)}
		result = append(result, seedB[randIndex.Int64()])
	}
	return string(result)
}
func StringLetter (size int) string {
	return StringBySeed("abcdefghijklmnopqrstuvwxyz", size)
}