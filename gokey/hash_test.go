package gokey

import "testing"

type hashTest struct {
	input  []byte
	output string
}

func TestGenerateMD5HashFromKey(t *testing.T) {
	tests := []hashTest{
		{[]byte("foo"), "acbd18db4cc2f85cedef654fccc4a4d8"},
		{[]byte("bar"), "37b51d194a7513e45b56f6524f2d51f2"},
		{[]byte("gopher"), "ca654591d4ac97414391907f882b3c05"},
		{[]byte("Irregardless"), "7dac0723ab11aec37ae53d26df848282"},
		{[]byte("earthquake"), "e9ad9c2394f7dc7b6a69fb43e52a7382"},
		{[]byte("endeavor"), "c5f5aaefa43684051f6fa380eef4b59e"},
	}
	generateMD5 := selectHash(MD5)
	for _, v := range tests {
		if out := generateMD5(v.input); out != v.output {
			t.Errorf("Find %s, expected %s", string(v.output), out)
		}
	}

}

func TestGenerateSha1HashFromKey(t *testing.T) {
	tests := []hashTest{
		{[]byte("foo"), "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33"},
		{[]byte("bar"), "62cdb7020ff920e5aa642c3d4066950dd1f01f4d"},
		{[]byte("gopher"), "4188736a00fbfb506aca06281acf338290455c21"},
		{[]byte("Irregardless"), "bd633cd299f2b91082944cd661e21923f106bff2"},
		{[]byte("earthquake"), "e8dea393aacba5f48ec3fdfd16114a89a2b59c63"},
		{[]byte("endeavor"), "2c84ef5af854a4eefb5832ffc01d8c0edd261645"},
	}
	generateSha1 := selectHash(SHA1)
	for _, v := range tests {
		if out := generateSha1(v.input); out != v.output {
			t.Errorf("Find %s, expected %s", string(v.output), out)
		}
	}

}

func TestGenerateSha256HashFromKey(t *testing.T) {
	tests := []hashTest{
		{[]byte("foo"), "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"},
		{[]byte("bar"), "fcde2b2edba56bf408601fb721fe9b5c338d10ee429ea04fae5511b68fbf8fb9"},
		{[]byte("gopher"), "9cc1ee455a3363ffc504f40006f70d0c8276648a5d3eb3f9524e94d1b7a83aef"},
		{[]byte("Irregardless"), "5bf4000a8a603e715286ad206e46a40e930c9e20998bf3ea36ddf687e6acced2"},
		{[]byte("earthquake"), "fa7835fc1c74a9e014c750a4f452b56f215702a9802d386a0560e5099da94fbb"},
		{[]byte("endeavor"), "1cd38b20bf937895efffc247bd9b85abbacfc5bfe21bfc12a44efa47a1da2343"},
	}
	generateSha256 := selectHash(SHA256)
	for _, v := range tests {
		if out := generateSha256(v.input); out != v.output {
			t.Errorf("Find %s, expected %s", string(v.output), out)
		}
	}

}
