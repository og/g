package grand

import (
	"math/rand"
	"time"
)

func StringBySeed(seed string, size int) string {
	seedB := []byte(seed)
	var result []byte
	seedLen := len(seedB)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i<size; i++ {
		result = append(result, seedB[r.Intn(seedLen)])
	}
	return string(result)
}
func StringLetter (size int) string {
	return StringBySeed("abcdefghijklmnopqrstuvwxyz", size)
}