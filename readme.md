# Stone

A stone is a token that holds or represents digital asset for the purpose of transferring between persons and machines.

# Stone Specification

See [Stone Doc](http://stonedoc.org) for the stone specification. 


# Install
```
go get github.com/stonedoc/stone

```

# Example

```Go
import (
   Stone "github.com/stonedoc/stone"
)

func main(){

    	var privKey = "-----BEGIN RSA PRIVATE KEY-----\nMIICXgIBAAKBgQDUAwSH1WJcV7I/sU4w54BNYFHwgvpxiXkmPDDkEjFL6+LKX46p\nsEccT8ETR7enF42qQtV3iFrtLi3Rr5QtIPB2cjASoDvkQT7TlpsPG5SHJHqF+7iD\ndS25GMR9eoDtvB7TyBk0B1SjSOYIizzPfYgdFoIl/X82BCtQVL2xnsaaBwIDAQAB\nAoGBAMvLfs5nYp5rOg+ZixTdY2p9fSZZcQ40XH1RfJmvly1ouN9ZjZQ1u5VOYMT8\nul/m9ylEB1hYfTbine6i/SeIMzuXMP+fNktCEMKFEdqGhvodu8EqQtJMk3bHIqmO\ndrjXdn20emdqUHTNdZUPU2lK89Q4Z+m4jEFoOAtOhbe3AhlJAkEA+h/SFMbq5QRP\nrxwuhg3M55iGRdf21ch5x6X4zRKyUayYTDqGl2DWKOitK5LwI2EsdsTdGpeR49U2\n3rRTLYNcJQJBANj9/7ITENa6ciFipw6X95OGcccuLPUydkaZwT37nmDD4iCrCFS2\nx85R+h0iktf6xWKqbLSzajFGp8LLovHxr7sCQQC+A4x6Ij9yKdtLITKqvjMqwbFH\nv/ARqpHxPMINMKXs7Bxq1I9I0tT/EPv1PVRW3EyGEboSqJC5L1HWz9Dco41NAkAy\nj+UP8n9e+az0eI9iyChpWM+UUP8q12pWAyfTMJl0BNDhOdlEHB8sxU9ZkJ/U8dsi\npYGVDaV1+/fFXTwH0oBXAkEA3KlgV9nQHpSkQS1SrElVdBkOHPnO90orv7RtB2SM\nfiztPjnExA2AVEBIj4hDRE34sNFnBRTWyCHQqU2JPrkaeg==\n-----END RSA PRIVATE KEY-----"

	
   	var metaBlock = map[string]interface{}{
		"created_at": 1457441770,
    	"id": "4ba3443b6157753fb678a593e5d7684c84d1b207",
     	"type": "currency",
	} 
  	 
    // create new stone object. Pass a valid meta block. 
   	stn, err := Stone.Create(metaBlock, privKey)
   	if err != nil {
      	panic(err)
   	}
   
}
```

# Load a stone

Given a JSON representation of a stone. A stone object can be loaded using the load() method. The JSON object must be a valid stone object. An error will be returned if stone validation is not passed.

```Go
import (
   Stone "github.com/ellcrys/stone"
)

func main(){

	var metaBlock = `{
       	"meta": {
		  	"id": "61144c09f35f1fd6d75265ceaf5bc8757c3a46c3",
	  		"created_at": 1457441779,
	  		"type": "coupon"
        }
   	}`
  	
    // load JSON stone
   	stn, err := Stone.Load(metaBlock)
   	if err != nil {
      	panic(err)
   	}  	
}
```

This method will not sign or verifies the resulting stone object. Use Sign() and Verify() method.


# Sign a block

#### Stone.sign(blockName, privateKey)


All blocks (except the signatures block) must be signed before encoding to base64 string for transmission. This method signs a block and deposits the resulting signature in the signatures block.

```Go
import (
   Stone "github.com/ellcrys/stone"
   "fmt"
)

func main(){

   var privKey = "-----BEGIN RSA PRIVATE KEY-----\nMIICXgIBAAKBgQDUAwSH1WJcV7I/sU4w54BNYFHwgvpxiXkmPDDkEjFL6+LKX46p\nsEccT8ETR7enF42qQtV3iFrtLi3Rr5QtIPB2cjASoDvkQT7TlpsPG5SHJHqF+7iD\ndS25GMR9eoDtvB7TyBk0B1SjSOYIizzPfYgdFoIl/X82BCtQVL2xnsaaBwIDAQAB\nAoGBAMvLfs5nYp5rOg+ZixTdY2p9fSZZcQ40XH1RfJmvly1ouN9ZjZQ1u5VOYMT8\nul/m9ylEB1hYfTbine6i/SeIMzuXMP+fNktCEMKFEdqGhvodu8EqQtJMk3bHIqmO\ndrjXdn20emdqUHTNdZUPU2lK89Q4Z+m4jEFoOAtOhbe3AhlJAkEA+h/SFMbq5QRP\nrxwuhg3M55iGRdf21ch5x6X4zRKyUayYTDqGl2DWKOitK5LwI2EsdsTdGpeR49U2\n3rRTLYNcJQJBANj9/7ITENa6ciFipw6X95OGcccuLPUydkaZwT37nmDD4iCrCFS2\nx85R+h0iktf6xWKqbLSzajFGp8LLovHxr7sCQQC+A4x6Ij9yKdtLITKqvjMqwbFH\nv/ARqpHxPMINMKXs7Bxq1I9I0tT/EPv1PVRW3EyGEboSqJC5L1HWz9Dco41NAkAy\nj+UP8n9e+az0eI9iyChpWM+UUP8q12pWAyfTMJl0BNDhOdlEHB8sxU9ZkJ/U8dsi\npYGVDaV1+/fFXTwH0oBXAkEA3KlgV9nQHpSkQS1SrElVdBkOHPnO90orv7RtB2SM\nfiztPjnExA2AVEBIj4hDRE34sNFnBRTWyCHQqU2JPrkaeg==\n-----END RSA PRIVATE KEY-----"

	var metaBlock = `{
       	"meta": {
		  	"id": "61144c09f35f1fd6d75265ceaf5bc8757c3a46c3",
	  		"created_at": 1457441779,
	  		"type": "coupon"
        }
   	}`
  	
    // load JSON stone
   	stn, err := Stone.Load(metaBlock)
   	if err != nil {
      	panic(err)
   	}  	
    
    // sign meta block
    signature, err := stn.Sign("meta", privKey)
    if err != nil {
      	panic(err)
   	}
    
    fmt.Println(signature)     // eyJhbGciOiJS.eyJjcmVhdGVkX2F0Ij.NyD4WN1ymdpgY032
}
```


# Verify a block's signature

#### Stone.verify(blockName, publicKey)

```Go
import (
   Stone "github.com/ellcrys/stone"
   "fmt"
)

func main(){

	var privKey = "-----BEGIN RSA PRIVATE KEY-----\nMIICXgIBAAKBgQDUAwSH1WJcV7I/sU4w54BNYFHwgvpxiXkmPDDkEjFL6+LKX46p\nsEccT8ETR7enF42qQtV3iFrtLi3Rr5QtIPB2cjASoDvkQT7TlpsPG5SHJHqF+7iD\ndS25GMR9eoDtvB7TyBk0B1SjSOYIizzPfYgdFoIl/X82BCtQVL2xnsaaBwIDAQAB\nAoGBAMvLfs5nYp5rOg+ZixTdY2p9fSZZcQ40XH1RfJmvly1ouN9ZjZQ1u5VOYMT8\nul/m9ylEB1hYfTbine6i/SeIMzuXMP+fNktCEMKFEdqGhvodu8EqQtJMk3bHIqmO\ndrjXdn20emdqUHTNdZUPU2lK89Q4Z+m4jEFoOAtOhbe3AhlJAkEA+h/SFMbq5QRP\nrxwuhg3M55iGRdf21ch5x6X4zRKyUayYTDqGl2DWKOitK5LwI2EsdsTdGpeR49U2\n3rRTLYNcJQJBANj9/7ITENa6ciFipw6X95OGcccuLPUydkaZwT37nmDD4iCrCFS2\nx85R+h0iktf6xWKqbLSzajFGp8LLovHxr7sCQQC+A4x6Ij9yKdtLITKqvjMqwbFH\nv/ARqpHxPMINMKXs7Bxq1I9I0tT/EPv1PVRW3EyGEboSqJC5L1HWz9Dco41NAkAy\nj+UP8n9e+az0eI9iyChpWM+UUP8q12pWAyfTMJl0BNDhOdlEHB8sxU9ZkJ/U8dsi\npYGVDaV1+/fFXTwH0oBXAkEA3KlgV9nQHpSkQS1SrElVdBkOHPnO90orv7RtB2SM\nfiztPjnExA2AVEBIj4hDRE34sNFnBRTWyCHQqU2JPrkaeg==\n-----END RSA PRIVATE KEY-----"
    
    var pubKey = "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDUAwSH1WJcV7I/sU4w54BNYFHw\ngvpxiXkmPDDkEjFL6+LKX46psEccT8ETR7enF42qQtV3iFrtLi3Rr5QtIPB2cjAS\noDvkQT7TlpsPG5SHJHqF+7iDdS25GMR9eoDtvB7TyBk0B1SjSOYIizzPfYgdFoIl\n/X82BCtQVL2xnsaaBwIDAQAB\n-----END PUBLIC KEY-----"

	var metaBlock = `{
       	"meta": {
		  	"id": "61144c09f35f1fd6d75265ceaf5bc8757c3a46c3",
	  		"created_at": 1457441779,
	  		"type": "coupon"
        }
   	}`
  	
    // load JSON stone
   	stn, err := Stone.Load(metaBlock)
   	if err != nil {
      	panic(err)
   	}  	
    
    // sign meta block
    signature, err := stn.Sign("meta", privKey)
    if err != nil {
      	panic(err)
   	}
    
    fmt.Println(signature)     // eyJhbGciOiJS.eyJjcmVhdGVkX2F0Ij.NyD4WN1ymdpgY032
    
    // verify meta block signature
    err = stn.Verify("meta", pubKey)
    if err != nil {
      	panic(err)
   	}
    
}
```


# Encode a stone

#### stone.encode()

This method creates a base64 url string of the signatures block. The resulting string can be shared or transferred.

```Go
import (
   Stone "github.com/ellcrys/stone"
   "fmt"
)

func main(){

	var privKey = "-----BEGIN RSA PRIVATE KEY-----\nMIICXgIBAAKBgQDUAwSH1WJcV7I/sU4w54BNYFHwgvpxiXkmPDDkEjFL6+LKX46p\nsEccT8ETR7enF42qQtV3iFrtLi3Rr5QtIPB2cjASoDvkQT7TlpsPG5SHJHqF+7iD\ndS25GMR9eoDtvB7TyBk0B1SjSOYIizzPfYgdFoIl/X82BCtQVL2xnsaaBwIDAQAB\nAoGBAMvLfs5nYp5rOg+ZixTdY2p9fSZZcQ40XH1RfJmvly1ouN9ZjZQ1u5VOYMT8\nul/m9ylEB1hYfTbine6i/SeIMzuXMP+fNktCEMKFEdqGhvodu8EqQtJMk3bHIqmO\ndrjXdn20emdqUHTNdZUPU2lK89Q4Z+m4jEFoOAtOhbe3AhlJAkEA+h/SFMbq5QRP\nrxwuhg3M55iGRdf21ch5x6X4zRKyUayYTDqGl2DWKOitK5LwI2EsdsTdGpeR49U2\n3rRTLYNcJQJBANj9/7ITENa6ciFipw6X95OGcccuLPUydkaZwT37nmDD4iCrCFS2\nx85R+h0iktf6xWKqbLSzajFGp8LLovHxr7sCQQC+A4x6Ij9yKdtLITKqvjMqwbFH\nv/ARqpHxPMINMKXs7Bxq1I9I0tT/EPv1PVRW3EyGEboSqJC5L1HWz9Dco41NAkAy\nj+UP8n9e+az0eI9iyChpWM+UUP8q12pWAyfTMJl0BNDhOdlEHB8sxU9ZkJ/U8dsi\npYGVDaV1+/fFXTwH0oBXAkEA3KlgV9nQHpSkQS1SrElVdBkOHPnO90orv7RtB2SM\nfiztPjnExA2AVEBIj4hDRE34sNFnBRTWyCHQqU2JPrkaeg==\n-----END RSA PRIVATE KEY-----"
    
    
	var metaBlock = `{
       	"meta": {
		  	"id": "61144c09f35f1fd6d75265ceaf5bc8757c3a46c3",
	  		"created_at": 1457441779,
	  		"type": "coupon"
        }
   	}`
  	
    // load JSON stone
   	stn, err := Stone.Load(metaBlock)
   	if err != nil {
      	panic(err)
   	}
    
    // sign meta block
    _, err = stn.Sign("meta", privKey)
    if err != nil {
      	panic(err)
   	}
    
    encoding := stn.Encode()
    fmt.Println(encoding)          // eyJtZXRhIjoiZXlKaGJHY2lPaUpTVXpJMU5pS...
}
```

# Decode an encoded stone

#### stone.decode(enc string)

A base64 url encoded stone can be decoded to the original stone object it was derived from. 

```Go
import (
   Stone "github.com/ellcrys/stone"
   "fmt"
)

func main(){

	var privKey = "-----BEGIN RSA PRIVATE KEY-----\nMIICXgIBAAKBgQDUAwSH1WJcV7I/sU4w54BNYFHwgvpxiXkmPDDkEjFL6+LKX46p\nsEccT8ETR7enF42qQtV3iFrtLi3Rr5QtIPB2cjASoDvkQT7TlpsPG5SHJHqF+7iD\ndS25GMR9eoDtvB7TyBk0B1SjSOYIizzPfYgdFoIl/X82BCtQVL2xnsaaBwIDAQAB\nAoGBAMvLfs5nYp5rOg+ZixTdY2p9fSZZcQ40XH1RfJmvly1ouN9ZjZQ1u5VOYMT8\nul/m9ylEB1hYfTbine6i/SeIMzuXMP+fNktCEMKFEdqGhvodu8EqQtJMk3bHIqmO\ndrjXdn20emdqUHTNdZUPU2lK89Q4Z+m4jEFoOAtOhbe3AhlJAkEA+h/SFMbq5QRP\nrxwuhg3M55iGRdf21ch5x6X4zRKyUayYTDqGl2DWKOitK5LwI2EsdsTdGpeR49U2\n3rRTLYNcJQJBANj9/7ITENa6ciFipw6X95OGcccuLPUydkaZwT37nmDD4iCrCFS2\nx85R+h0iktf6xWKqbLSzajFGp8LLovHxr7sCQQC+A4x6Ij9yKdtLITKqvjMqwbFH\nv/ARqpHxPMINMKXs7Bxq1I9I0tT/EPv1PVRW3EyGEboSqJC5L1HWz9Dco41NAkAy\nj+UP8n9e+az0eI9iyChpWM+UUP8q12pWAyfTMJl0BNDhOdlEHB8sxU9ZkJ/U8dsi\npYGVDaV1+/fFXTwH0oBXAkEA3KlgV9nQHpSkQS1SrElVdBkOHPnO90orv7RtB2SM\nfiztPjnExA2AVEBIj4hDRE34sNFnBRTWyCHQqU2JPrkaeg==\n-----END RSA PRIVATE KEY-----"
    
    
	var metaBlock = `{
       	"meta": {
		  	"id": "61144c09f35f1fd6d75265ceaf5bc8757c3a46c3",
	  		"created_at": 1457441779,
	  		"type": "coupon"
        }
   	}`
  	
    // load JSON stone
   	stn, err := Stone.Load(metaBlock)
   	if err != nil {
      	panic(err)
   	}
    
    // sign meta block
    _, err = stn.Sign("meta", privKey)
    if err != nil {
      	panic(err)
   	}
    
    // encode stone object
    enc := stn.Encode()
    fmt.Println(enc)          // eyJtZXRhIjoiZXlKaGJHY2lPaUpTVXpJMU5pS...
    
    // decode
    decodedStone, err := Stone.Decode(enc)
    if err != nil {
      	panic(err)
   	}
    
    fmt.Println(decodedStone)
    
}
```

# Other Methods

See [GoDoc](https://godoc.org/github.com/ellcrys/stone) for full documentation
