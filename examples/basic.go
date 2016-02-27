package main 

import (
	seed "../"
	"time"
)

var TXT = `
{
   "signatures": {
      "meta": "abcde",
      "ownership": "aa",
      "attributes": "abcde"
   },
   "meta":{
      "created_at":`+seed.IntToString(time.Now().Unix())+`,
      "seed_id": "`+seed.NewID()+`",
      "seed_type":"currency"
   },
   "ownership": {
      "type": "sole",
      "sole": {
      	  "address_id": "56c1f3d1f68daa4584000001"
      }
   },
   "embeds": [],
   "attributes": {
      "stuff": "stuff"
   }
}
`

func main() {

	// var sh = "eyJzaWduYXR1cmVzIjp7Im1ldGEiOiIzZjdhMTEyYjc5N2E1NzI0NzgyYTc1N2FkOTM5OWE4MDBhMTY4YTM4ZDA4ZjY2MGQxZTdjMzVhMWY3NTFkZDZiNzlmMzAzNTdhZDY1Y2NhN2JmMWQ3YjEwNmNiNzI0MGE3ZWIwMzZkYTM4ZTNkZjRkZWJlOGI4ZTk3YWQ4MmNjOTVhZjhkNmU5MTliNzk5NWY4NGE2NWYyMGI2YmQ1NzYyZDQyMjFhNzM0ZjY3MzA3YWI0ZWE2YTg0OTExZGFiNGM1ZGY0NmY2YWEzZTI1MzFiNmFiZjViMWRlZTU3MjIzM2FjNTFkOWUxYmVmY2VkM2MxNmQ1YTczYzM2NmMxZmJkIn0sIm1ldGEiOnsiY3JlYXRlZF9hdCI6MTQ1NTc5OTY3OSwic2hlbGxfaWQiOiJkMjVmMGQzZGE5ZDYwY2I1YWVjZDVkNTg4YzM5YjM0NzYxNzJhMWZiIiwic2hlbGxfdHlwZSI6ImN1cnJlbmN5In0sIm93bmVyc2hpcCI6e30sImVtYmVkcyI6W10sImF0dHJpYnV0ZXMiOnt9fQ=="
	// var issuerPublicKey = "-----BEGIN RSA PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC1CCgxMS5rro13G4frgt88bq00\ncdWj8r/VK3uUi7CK7Dq8dmbUU5SfqeuCzEg0EUIcaiSabndev5CfUVTPLKakmelA\nv7f8F0c3hTTaInZcZ9F9xuWdfklKVlU63fSEl14+qGddRyu/vs3wYHUgesZZG16R\nZcNmRfDqMd3XYkEz8QIDAQAB\n-----END RSA PUBLIC KEY-----\n"
	var issuerPrivateKey = "-----BEGIN RSA PRIVATE KEY-----\nMIICXgIBAAKBgQC1CCgxMS5rro13G4frgt88bq00cdWj8r/VK3uUi7CK7Dq8dmbU\nU5SfqeuCzEg0EUIcaiSabndev5CfUVTPLKakmelAv7f8F0c3hTTaInZcZ9F9xuWd\nfklKVlU63fSEl14+qGddRyu/vs3wYHUgesZZG16RZcNmRfDqMd3XYkEz8QIDAQAB\nAoGBAKSWhfQzgsDkMjnYDftRQSTwewjzdglY6pAkuHaViizEGaK/Az6Hvthq5HRG\nxl6QUksDNcQyKtU51YMDrtetANOdmL2w2SXrK2sRFLqj2zB5eqn+GW87+HN5tIYi\n4FAr4/2k6oSI9yFxQ6R+pusYTMzaZdeMYvMwu0P6W6Lbb4sBAkEA57QYjqQ0nU2L\n8/5sq7GtjzFLJ8HoD7F2keLaLWh0EH0xNPTkjykdD+cFuJbSgdzbQv8q8zh3EU5i\nWUfljoJ6/wJBAMgD0c8WtMJ+Yv1HJo1/TWq55I9zh/gvIjG22MCa+AB31ZfYlV6p\now9CHT2uAUIW+j1YT68PeShWIvm4vZ7fAQ8CQQCoN5xMkvKP8ajV77U9wbVb7FG/\n/4tXOWP37lav+NGq1vlOlS0KsrKixPrmVLloBsw5C8BG7IulSN8mKoiCukBJAkBv\nXLkPbVv9QjNJQ7kyZSOsfY3FVRTqWQvX1C9ApcfZMt9oqP0ZdKfGEhCHy/8FVhfD\n2gybsqjJjZPxqCtjblR/AkEA3jD4MNRulqbu+U28Nt0Vxx8kbpEyfSm68WqkoEqn\nu/ILWzlYzTyqoZ/b5e0igHwzHP2Ex9GV49+ROI/+z2uA/A==\n-----END RSA PRIVATE KEY-----\n"
	// var hostPrivateKey = "-----BEGIN RSA PRIVATE KEY-----\nMIICXwIBAAKBgQC8RC0TxSEfuFAsdmV36J/izIUrYhbVl62GGDoBMWsiCuGdsAlb\nTG0hh3I/7blk2q1Kc2P08hTtK7dEkaQIgHZvljPJ7kZ92SHElHsIIqFFEV3S+kyH\nGvbv0Qn6Nu8ghDjEi2gK4ZwDyvMzLzgdfP+u3OCEg7KihOzKRLRwR1trJwIDAQAB\nAoGBAKHHA+VFNB2JyHsskizj3OCLVrPc6jpIyHe+QbncuW7bYtyZ9LBbkDuLpIWC\nxhkAQIEFfxNxIsJbGnT1obsciKgH5gGAR02fR1B4j77SDlWebXqEPCv2JHB640OB\nORf0BgbQqUj+46jS24uPdmofgcyXp3BA8o054ccH+xLEdlohAkEA792PENfejkEM\nyX+wWQhKhTNPCU9/o53HAYR5CVjdBsa/OvVVcXLuhmxGlrYhCBjA3Axh/ahS3BGi\nkJLDBRM1UwJBAMjuFsLYQekVtGSqGoU5TPfVFjwggOaSkduGBWKf1z2rEQMPcIqh\n0fqOdY4wX4cV5W6w3CvgK8hH8MO0dOANRF0CQQCZ2v6qahzaGEWQdfPyl8vc9pVK\nvpB7rXd5tLRCV5qmfxMoSTc+Jt9yn78Dat1zKRWDz/mGz9IeUL16iHjJJ5H7AkEA\norwqlR5/q18X3pvipNn225aqzoHoxFYbafeO7wTUWC7vtVHQ7YcIQO0WitXk2MzE\nKLV3bNW/wBN8DZVfP4OfRQJBAKWzrB32yKFUvOdCjQNfn7xidxSFadmyLIYntEkf\n3Z4ZA1usDwIPWhmfnJBlHnw3QmFoHg1p2D7ve4DPXDrKnoM=\n-----END RSA PRIVATE KEY-----\n"
	var meta = map[string]interface{}{
		"seed_id": seed.NewID(),
		"seed_type": "currency",
		"created_at": time.Now().Unix(),
	}

	sh, e := seed.Create(meta, issuerPrivateKey)
	seed.Println("CREATE: ", e)
	seed.Println(sh.Encode())
	// sh2, _ := seed.Load(sh.Encode())
	// seed.Println(sh2.Encode())
	// seed.Println("VERIFY: ", sh2.Verify("meta", pubKey))

	// sh, e := seed.Load(TXT)
	// seed.Println("LOAD:", e)
	// sig, err := sh.Sign("meta", issuerPrivateKey)
	// if err != nil {
	// 	seed.Println("SIGN META:", err)
	// 	return
	// }

	// sh.Signatures["meta"] = sig

	// sig, err = sh.Sign("ownership", hostPrivateKey)
	// if err != nil {
	// 	seed.Println("SIGN OWNERSHIP:", err)
	// 	return
	// }
	// sh.Signatures["ownership"] = sig
	
	// sig, err = sh.Sign("attributes", issuerPrivateKey)
	// if err != nil {
	// 	seed.Println("SIGN ATTRIBUTES:", err)
	// 	return
	// }
	// sh.Signatures["attributes"] = sig

	// sh2, _ := seed.Load(TXT)
	// sig, _ = sh.Sign("meta", issuerPrivateKey)
	// sh2.Signatures["meta"] = sig
	// sh.AddEmbed(sh2, issuerPrivateKey)

	// seed.Println("VERIFY: ", sh.Verify("meta", issuerPublicKey))
	// seed.Println("\n", sh.Encode())

	// manager := seed.Manager{}
	// aSeed := manager.NewSeed("currency");
	// aSeed.AddTransaction(seed.GenesisTxn(1000));
	// fmt.Printf("%v\n", currencySeed)
	// seed.Println(aSeed.JSON())
	// anoSeed, err := seed.LoadJSON(TXT)
	// if err != nil {
	// 	seed.Println(err)
	// 	return
	// }
	// seed.Println(anoSeed.Embeds[1].Embeds[0].Txns)
}

