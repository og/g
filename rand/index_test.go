package grand

import (
	gtest "github.com/og/x/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandBySeed(t *testing.T) {
	{
		seed := "ab"
		result := RunesBySpeed(seed, 10)
		assert.Equal(t, 10, len(result))
		assert.Regexp(t, "a", string(result))
		assert.Regexp(t, "b", string(result))
	}
	{
		seed := "abcd"
		result := RunesBySpeed(seed, 20)
		assert.Equal(t, 20, len(result))
		assert.Regexp(t, "a", string(result))
		assert.Regexp(t, "b", string(result))
		assert.Regexp(t, "c", string(result))
		assert.Regexp(t, "d", string(result))
	}
	{
		seed := "a我"
		result := RunesBySpeed(seed, 20)
		assert.Equal(t, 20, len(result))
		assert.Regexp(t, "a", string(result))
		assert.Regexp(t, "我", string(result))
	}
	{
		sList := []string{}
		countHash := map[string]int{}
		for i:=0; i<100000; i++ {
			s := string(RunesBySpeed("1234567890", 10))
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

func TestLetterBytes(t  *testing.T) {
	as := gtest.NewAS(t)
	as.Equal(string(letterBytes()), "abcdefghijklmnopqrstuvwxyz")

}