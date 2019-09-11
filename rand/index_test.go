package grand

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
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
	{
		result := StringLetter(200)
		assert.Equal(t, 200, len(result))
		letterList := strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLOMNOPQRSTUVWXYZ", "")
		for _, letter := range letterList {
			assert.Regexp(t, letter, result)
		}
		log.Print(result)
	}

}
