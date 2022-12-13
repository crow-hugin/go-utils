package secret

import "testing"

func TestEncryptToken(t *testing.T) {
	t.Log(EncryptToken("123456"))
}
