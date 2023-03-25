package gokey

import "testing"

type hashMD5Test struct {
	input  []byte
	output string
}

func TestGenerateMD5HashFromKey(t *testing.T) {
	tests := []hashMD5Test{
		{[]byte("foo"), "acbd18db4cc2f85cedef654fccc4a4d8"},
		{[]byte("bar"), "37b51d194a7513e45b56f6524f2d51f2"},
		{[]byte("gopher"), "ca654591d4ac97414391907f882b3c05"},
		{[]byte("Irregardless"), "7dac0723ab11aec37ae53d26df848282"},
		{[]byte("earthquake"), "e9ad9c2394f7dc7b6a69fb43e52a7382"},
		{[]byte("endeavor"), "c5f5aaefa43684051f6fa380eef4b59e"},
	}
	generateMD5 := selectHash("md5")
	for _, v := range tests {
		if out, err := generateMD5(v.input); err != nil || out != v.output {
			t.Errorf("Find %s, expected %s", string(v.output), out)
		}
	}

}
