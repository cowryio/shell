# Stone

A stone is a token that holds or represents digital asset for the purpose of transferring between persons and machines.

# Stone Specification

See [Stone Doc](http://stonedoc.org) for the stone specification. 


# Install
```
go get github.com/ellcrys/stone
```

# Example

```golang
import (
   Stone "github.com/ellcrys/stone"
)

func main(){

    var privKey = "-----BEGIN RSA PRIVATE KEY-----\nMIICXgIBAAKBgQ...kaeg==\n-----END RSA PRIVATE KEY-----"
	
   	var metaBlock = `{
	  	id: "61144c09f35f1fd6d75265ceaf5bc8757c3a46c3",
	  	created_at: 1457441779,
	  	type: "coupon"
   	}`
  	 
   	stn, err := Stone.Create(json1, privKey)
   	if err != nil {
      	panic(err)
   	}
   
}
```