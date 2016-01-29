package main 

import (
	shell "../"
	"time"
)

const TXT = `
{
   "signatures": {
      "meta": "abcde"
   },
   "meta":{
      "created_at":1453975575,
      "genesis":true,
      "shell_id":"4417781906fb0a89c295959b0df01782dbc4dc9f",
      "shell_type":"currency"
   },
   "ownership":null,
   "embeds": [],
   "attributes":null
}
`

func main() {

	var issuerPrivateKey = "-----BEGIN RSA PRIVATE KEY-----\nMIICWwIBAAKBgQCroZieOAo9stcf6R6eWfo51VCvK8cLdNS577m/HIFOmEd1CDi/\nu7agGzpehNAhHpr5NVjQZ4Te+KMRn9SnpUK2hc8dUU25PQolsOEwePVQ18hHNK4Y\n2JvOY/f8KCO2hhrS6uuP6eedpnSdulS1OXHTL6ZxQmBd9F33gLT6BERHQwIDAQAB\nAoGAEZ/0ljrXAmL9KG++DzDaO1omgPaT6B9FQRrXDkMVHEcS/3eqrDXQmTxykAY/\ngUctTu4lgrE+uc76n/Kz2ctkwEKIKet56ylqp+wlEUt1G+udoi07tgd7XyxzoUJm\nZwSm89gKh+mEPxni0FrBNg6dR0n2gvKRecnXqyoGVOHZITECQQDXgRJyrzgc/JhB\nSOBznEjtXAZXRRu3o9UznztjU9Xz7NWXTVuHu8WqYmGWCOqnysMhXJ3xBddJyDTF\njuOJ0123AkEAy+H+3POcT2FDOuluqPmAZQAUU6Nxtbj02/JJtOy7jq5jnN27HVC3\nuQzmfsS5J2XeQQodOUwOy2Ub57/OMrMi1QJAGZsZgQz2wuL0iFVLbhE0zRcxHa91\ncqWB0Kdr3Ap7EoeifV7QsFkMTIlyBOy8TQGXm+AwWBIUmYyzUIIA4UB/EwJAO+Bo\nSB2nZ0yqQO/zVt7HjWIDljinGXZzOvEiImdwAcxHZvdbj5V4D3mxa8N8mQx6xGEj\nCgPDSIquMlaLSSqA7QJAAbQPa0frCkm1rkWWZ7QwGm7ptzOACwFEGefm/1mhmw3a\nvoWRTHhrDuEbeVH3iF8MWhLJLPFtuSShiQMsrVbXPA==\n-----END RSA PRIVATE KEY-----"
	var meta = map[string]interface{}{
		"shell_id": shell.NewID(),
		"shell_type": "currency",
		"created_at": time.Now().Unix(),
	}

	aShell, err := shell.Create(meta, issuerPrivateKey)
	// anoShell, err := shell.LoadJSON(TXT)
	// shell.Println(anoShell, err)

	shell.Println(aShell.JSON(), err)



	// manager := shell.Manager{}
	// aShell := manager.NewShell("currency");
	// aShell.AddTransaction(shell.GenesisTxn(1000));
	// fmt.Printf("%v\n", currencyShell)
	// shell.Println(aShell.JSON())
	// anoShell, err := shell.LoadJSON(TXT)
	// if err != nil {
	// 	shell.Println(err)
	// 	return
	// }
	// shell.Println(anoShell.Embeds[1].Embeds[0].Txns)
}

