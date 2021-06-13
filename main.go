package main

func main() {

	// cipher key
	key := "thisis32bitlongpassphraseimusing"

	// plaintext
	pt := "This is a secret"

	encryptWithAES(pt, key)

	encryptionWithBcrypt(pt)

	encryptWithCrypto(pt, key)

}

/*
https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64
*/
