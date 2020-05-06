package gcrypto

import (
	gtest "github.com/og/x/test"
	"testing"
)

func TestSHA512(t *testing.T) {
	as := gtest.NewAS(t)
	as.Equal(SHA512("nimo"), "166a0152f808e1d939d952cab92b3cd783916e10b6e136ad22022e8fa4013355e1a38479e934da33982edcb59d70ef5ec198b9cf4d14d9ea6da3dff0a4ab51c1")
	warningMessage := ""
	warningLog = func(msg string) {
		warningMessage = msg
	}
	as.Equal(SHA512(""), "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e")
	as.Equal(warningMessage, "SHA521(s string) s is empty string, There may have been some errors")
	{
		hash := SaltSHA512("12345", "salt_sdfhqwf")
		as.Equal(hash, "70c44ce2ff284274be29974f32dd27184e21ceb4b985cd866b9e714de1ca4dabe4ed2f9b172b0efe247194b075388fe44755a995b0f8dcda10e7a7327370afa7")
		as.True(CheckSaltSHA512("12345", "salt_sdfhqwf", "70c44ce2ff284274be29974f32dd27184e21ceb4b985cd866b9e714de1ca4dabe4ed2f9b172b0efe247194b075388fe44755a995b0f8dcda10e7a7327370afa7"))
	}
}
