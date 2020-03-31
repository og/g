package grand

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandBySeed(t *testing.T) {
	{
		seed := "ab"
		result := StringBySeed(seed, 10)
		assert.Equal(t, 10, len(result))
		assert.Regexp(t, "a", result)
		assert.Regexp(t, "b", result)
	}
	{
		seed := "abcd"
		result := StringBySeed(seed, 10)
		assert.Equal(t, 10, len(result))
		assert.Regexp(t, "a", result)
		assert.Regexp(t, "b", result)
		assert.Regexp(t, "c", result)
		assert.Regexp(t, "d", result)
	}
	// {
	// 	result := StringLetter(200)
	// 	assert.Equal(t, 200, len(result))
	// 	letterList := strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLOMNOPQRSTUVWXYZ", "")
	// 	for _, letter := range letterList {
	// 		assert.Regexp(t, letter, result)
	// 	}
	// }
	{
		sList := []string{}
		countHash := map[string]int{}
		for i:=0; i<100000; i++ {
			s := StringBySeed("1234567890", 10)
			sList = append(sList, s)
			count, has := countHash[s]
			if !has {
				countHash[s] = 0
				count = 0
			}
			if count > 1 {
				t.Fatalf("%s  count is  %v", s, count)
			}
			countHash[s] = count+1
		}
	}
}
