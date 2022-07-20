package main

import (
	"fmt"
	"main/block"
)

func main() {
	/*
		// ipfs in çalışması gerekiyor ki ona bağlansın.
		sh := shell.NewShell("localhost:5001")

		//cid, err := sh.Add(strings.NewReader("hello world selam")) // yazı için

		cid, err := sh.AddDir("./ornek.png") //dosya için

		//Güvenliği arttırmak için anahtarı rastgele döngü halinde değiştirmek.

		bytes := make([]byte, 32) //AES-256 için rastgele bir 32 bayt anahtar oluşturun.
		if _, err := rand.Read(bytes); err != nil {
			panic(err.Error())
		}
		fmt.Printf("\n")
		key := hex.EncodeToString(bytes) //anahtarı bayt cinsinden kodlayın ve gizli olarak saklayın, bir kasaya koyun
		fmt.Printf("key(anahtar) : %s\n", key)

		encrypted := block.Encrypt(cid, key)
		fmt.Printf("encrypted(şifreli) : %s\n", encrypted)
	*/
	decrypted := block.Decrypt("62bf98b4c55696043a8948157b9377c00dd8f401620e535a3df369b20a75f4023118abfa40d9f6533799e7e6ab59603a14697d15c13f011582a234175388a03741ba9bb7feb872b82a20", "e163b77bf4f5df523a90eada25a3f85b36066b1f66e43bbc114cd0450ec3f410")
	fmt.Printf("decrypted(şifre çözüm) : %s\n", decrypted)

	fmt.Printf("https://ipfs.io/ipfs/%s", decrypted)
}
