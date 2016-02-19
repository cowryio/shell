// A shell represents a container of information
// that is considered valueable.
package shell

import (
	"encoding/json"
	"errors"
	"strings"
	"fmt"
)

type Shell struct {
	Signatures map[string]interface{} 		`json:"signatures"`
	Meta map[string]interface{}				`json:"meta"`
	Ownership map[string]interface{} 		`json:"ownership"`
	Embeds []map[string]interface{} 		`json:"embeds"`
	Attributes map[string]interface{}		`json:"attributes"`
}

// Initialize a shell
func initialize(shell *Shell) *Shell {
	shell.Signatures = make(map[string]interface{})
	shell.Meta = make(map[string]interface{})
	shell.Ownership = make(map[string]interface{})
	shell.Embeds = []map[string]interface{}{}
	shell.Attributes = make(map[string]interface{})
	return shell
}

// creates an shell instances and initializes it
func Empty() *Shell {
	sh := &Shell{}
	return initialize(sh)
}

// Create a shell.The new shell is immediately signed using the issuer's private key
func Create(meta map[string]interface{}, issuerPrivateKey string) (*Shell, error) {

	shell := initialize(&Shell{})

	// validate meta
	if err := ValidateMetaBlock(meta); err != nil {
    	return &Shell{}, err
    } else {
		meta["created_at"] = IntToFloat64(meta["created_at"])
    }

    // set shell Meta field and create a meta signature
	shell.Meta = meta
	_, err := shell.Sign("meta", issuerPrivateKey)
	if err != nil {
		return &Shell{}, err
	}

	return shell, nil
}

// Creates a shell from a map. Validation of expected field is 
// not performed. Use Validate() before calling this method if validation
// is necessary
func loadMap(data map[string]interface{}) (*Shell, error) {

	var shell = &Shell{}

	// add signatures
    if signatures := data["signatures"]; signatures != nil {
    	shell.Signatures = data["signatures"].(map[string]interface{})
    }

    // add meta
    if meta := data["meta"]; meta != nil {
    	shell.Meta = data["meta"].(map[string]interface{})
    }

    // add ownership
    if ownership := data["ownership"]; ownership != nil {
    	shell.Ownership = data["ownership"].(map[string]interface{})
    }

    // add attributes
    if attributes := data["attributes"]; attributes != nil {
    	shell.Attributes = data["attributes"].(map[string]interface{})
    }

    // add embeds
    if embeds := data["embeds"]; embeds != nil {
    	for _, m := range embeds.([]interface{}) {
    		shell.Embeds = append(shell.Embeds, m.(map[string]interface{}))
    	}
    }

    return shell, nil
}

// Creates a new shell from a raw json or base 64 encoded json string. It does not 
// attempt to sign the blocks. 
// If the string passed in starts with "{", it is considered a JSON string, otherwise, it assumes string is base 64 encoded and
// will attempt to decoded it. 
func Load(shellStr string) (*Shell, error) {
	shellStr = strings.TrimSpace(shellStr)
	if shellStr == "" {
		return &Shell{}, errors.New("Cannot load empty shell string")
	} else {
		if fmt.Sprintf("%c", shellStr[0]) == "{" {					// json string
			return LoadJSON(shellStr)
		} else {
			decodedShellStr, err := FromBase64(shellStr)
			if err != nil {
				return &Shell{}, errors.New("unable to decode encoded shell string")
			}
			return LoadJSON(decodedShellStr)
		}
	}
}


// Create a shell from a json string by converting
// it to a map and then used to load a new shell instance
func LoadJSON(jsonStr string) (*Shell, error) {
	data, err := JSONToMap(jsonStr)
	if err != nil{
        return &Shell{}, err;
    }
    if err := Validate(data); err != nil {
    	return &Shell{}, err
    }
	return loadMap(data)  
}

// converts a json string to map 
func JSONToMap(jsonStr string) (map[string]interface{}, error) {
	var data map[string]interface{}
	d := json.NewDecoder(strings.NewReader(jsonStr))
	if err := d.Decode(&data); err != nil {
        return make(map[string]interface{}), errors.New("unable to parse json string");
    }
	return data, nil
}


// Sign any shell block by creating a canonical string representation
// of the block value and signing with the issuer's private key. The computed signature
// is store the `signatures` block
func(self *Shell) Sign(blockName string, privateKey string) (string, error) {
	
	var canonicalString string

	switch blockName {
	case "meta":
		canonicalString = GetCanonicalMapString(self.Meta)
		break
	case "ownership":
		canonicalString = GetCanonicalMapString(self.Ownership)
		break
	case "attributes":
		canonicalString = GetCanonicalMapString(self.Attributes)
		break
	case "embeds": 
		for _, shell := range self.Embeds {
			canonicalString += ":" + GetCanonicalMapString(shell["meta"].(map[string]interface{}))
		}
		canonicalString = strings.Trim(canonicalString, ":")
		break
	default:
		return "", errors.New("block unknown")
	}

	signer, err := ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return "", errors.New(fmt.Sprintf("Private Key Error: %v", err))
	}

	signature, err := signer.Sign([]byte(canonicalString))
	if err != nil {
		return "", errors.New(fmt.Sprintf("Signature Error: %v", err))
	}

	self.Signatures[blockName] = signature
	return signature, nil
}

// Assign and sign a valid meta value to the meta block
func(self *Shell) AddMeta(meta map[string]interface{}, issuerPrivateKey string) error {

	// validate meta
	if err := ValidateMetaBlock(meta); err != nil {
    	return err
    }

    self.Meta = meta

    // sign meta block
    _, err := self.Sign("meta", issuerPrivateKey)
	if err != nil {
		return err
	}

    return nil
}

// Assign and sign a valid ownership data to the ownership block
func (self *Shell) AddOwnership(ownership map[string]interface{}, issuerPrivateKey string) error {

	// validate 
	if err := ValidateOwnershipBlock(ownership); err != nil {
    	return err
    }

	self.Ownership = ownership

	// sign block
    _, err := self.Sign("ownership", issuerPrivateKey)
	if err != nil {
		return err
	}

	return nil
}

// Assign a attribute data to the attributes bloc
func (self *Shell) AddAttributes(attributes map[string]interface{}, issuerPrivateKey string) error {
	
	self.Attributes = attributes

	// sign block
    _, err := self.Sign("attributes", issuerPrivateKey)
	if err != nil {
		return err
	}
	
	return nil
}

// add a shell to the `embeds` block
func (self *Shell) AddEmbed(shell *Shell, issuerPrivateKey string) error {

	self.Embeds = append(self.Embeds, shell.ToMap())

	// sign block
	_, err := self.Sign("embeds", issuerPrivateKey)
	if err != nil {
		return err
	}

	return nil
}


// checks if a block has a signature
func(self *Shell) HasSignature(blockName string) bool {
	switch blockName {
	case "meta", "ownership", "attributes", "embeds":
		return self.Signatures[blockName] != nil && strings.TrimSpace(self.Signatures[blockName].(string)) != ""
		break
	default:
		return false
	}
	return false
}

// Verify one or all block. If blockName is set to an empty string,
// all blocks are verified.
func(self *Shell) Verify(blockName, issuerPublicKey string) error {

	var canonicalString string

	switch blockName {
	case "meta":
		canonicalString = GetCanonicalMapString(self.Meta)
	case "ownership":
		canonicalString = GetCanonicalMapString(self.Ownership)
	case "attributes":
		canonicalString = GetCanonicalMapString(self.Attributes)
	case "embeds": 
		for _, shell := range self.Embeds {
			canonicalString += ":" + GetCanonicalMapString(shell["meta"].(map[string]interface{}))
		}
		canonicalString = strings.Trim(canonicalString, ":")
		break
	default:
		return errors.New("block name "+blockName+" is unknown")
	}

	signer, err := ParsePublicKey([]byte(issuerPublicKey))
	if err != nil {
		return errors.New(fmt.Sprintf("Public Key Error: %v", err))
	}

	// block has no signature
	if !self.HasSignature(blockName) {
		return errors.New("block `"+blockName+"` has no signature")
	}
	
	return signer.Verify([]byte(canonicalString), self.Signatures[blockName].(string))
}  

// checks if a shell object current state can
// pass as a valid shell
func(self *Shell) IsValid() error {
	return Validate(self.JSON())
}

// return shell as raw JSON string
func(self *Shell) JSON() string {
	bs, _ := json.Marshal(&self)
	return string(bs)
}

// returns a map representation of the shell
func(self *Shell) ToMap() map[string]interface{} {
	jsonStr := self.JSON()
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dat); err != nil {
        panic(err)
    }
    return dat
}

// return shell as a base64 json encoded string
func(self *Shell) Encode() string {
	jsonStr := self.JSON()
	return ToBase64([]byte(jsonStr))
}

// clone a shell
func(self *Shell) Clone() *Shell {
	jsonStr := self.JSON()
	shell, err := LoadJSON(jsonStr)
	if err != nil {
		panic(err)
	}
	return shell
}

// checks if the ownership block is field
func(self *Shell) HasOwnership() bool {
	return len(self.Ownership) > 0
}

