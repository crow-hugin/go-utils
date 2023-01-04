package secret

import "testing"

func TestEncryptToken(t *testing.T) {
	t.Log(EncryptToken("123123123123123"))
	//t.Log(DecryptToken("aHUYX6P6sbE5YEFfAtOjrd8DU5LecGOjPP0RR9GcozOQGE5R1vzHMty6CpPBOyH90lOF4buYixeHd17StJIeET5X1C7vwbYGZP1M5y8c0WqQJgBs_btRSwb5txg3yLEl6"))
	t.Log(DecryptToken("mBnAiw3i0gMJB5jjmS9e7g_c_c"))
}
