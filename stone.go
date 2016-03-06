// A stone represents a container of information
// that is considered valueable.
package stone

import (
	"encoding/json"
	"errors"
	"strings"
	"fmt"
	"github.com/ellcrys/crypto"
)

type Stone struct {
	Signatures map[string]interface{} 		`json:"signatures"`
	Meta map[string]interface{}				`json:"meta"`
	Ownership map[string]interface{} 		`json:"ownership"`
	Embeds []map[string]interface{} 		`json:"embeds"`
	Attributes map[string]interface{}		`json:"attributes"`
}

// Initialize a stone
func initialize(stone *Stone) *Stone {
	stone.Signatures = make(map[string]interface{})
	stone.Meta = make(map[string]interface{})
	stone.Ownership = make(map[string]interface{})
	stone.Embeds = []map[string]interface{}{}
	stone.Attributes = make(map[string]interface{})
	return stone
}

// creates an stone instances and initializes it
func Empty() *Stone {
	sh := &Stone{}
	return initialize(sh)
}

// Create a stone.The new stone is immediately signed using the issuer's private key
func Create(meta map[string]interface{}, issuerPrivateKey string) (*Stone, error) {

	stone := initialize(&Stone{})

	// validate meta
	if err := ValidateMetaBlock(meta); err != nil {
    	return &Stone{}, err
    } else {
		meta["created_at"] = IntToFloat64(meta["created_at"])
    }

    // set stone Meta field and create a meta signature
	stone.Meta = meta
	_, err := stone.Sign("meta", issuerPrivateKey)
	if err != nil {
		return &Stone{}, err
	}

	return stone, nil
}

// Creates a stone from a map. Validation of expected field is 
// not performed. Use Validate() before calling this method if validation
// is necessary
func loadMap(data map[string]interface{}) (*Stone, error) {

	var stone = &Stone{}

	// add signatures
    if signatures := data["signatures"]; signatures != nil {
    	stone.Signatures = data["signatures"].(map[string]interface{})
    }

    // add meta
    if meta := data["meta"]; meta != nil {
    	stone.Meta = data["meta"].(map[string]interface{})
    }

    // add ownership
    if ownership := data["ownership"]; ownership != nil {
    	stone.Ownership = data["ownership"].(map[string]interface{})
    }

    // add attributes
    if attributes := data["attributes"]; attributes != nil {
    	stone.Attributes = data["attributes"].(map[string]interface{})
    }

    // add embeds
    if embeds := data["embeds"]; embeds != nil {
    	for _, m := range embeds.([]interface{}) {
    		stone.Embeds = append(stone.Embeds, m.(map[string]interface{}))
    	}
    }

    return stone, nil
}

// Creates a new stone from a raw json or base 64 encoded json string. It does not 
// attempt to sign the blocks. 
// If the string passed in starts with "{", it is considered a JSON string, otherwise, it assumes string is base 64 encoded and
// will attempt to decoded it. 
func Load(stoneStr string) (*Stone, error) {
	stoneStr = strings.TrimSpace(stoneStr)
	if stoneStr == "" {
		return &Stone{}, errors.New("Cannot load empty stone string")
	} else {
		if fmt.Sprintf("%c", stoneStr[0]) == "{" {					// json string
			return LoadJSON(stoneStr)
		} else {
			decodedStoneStr, err := crypto.FromBase64(stoneStr)
			if err != nil {
				return &Stone{}, errors.New("unable to decode encoded stone string")
			}
			return LoadJSON(decodedStoneStr)
		}
	}
}


// Create a stone from a json string by converting
// it to a map and then used to load a new stone instance
func LoadJSON(jsonStr string) (*Stone, error) {
	data, err := JSONToMap(jsonStr)
	if err != nil{
        return &Stone{}, err;
    }
    if err := Validate(data); err != nil {
    	return &Stone{}, err
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


// Sign any stone block by creating a canonical string representation
// of the block value and signing with the issuer's private key. The computed signature
// is store the `signatures` block
func(self *Stone) Sign(blockName string, privateKey string) (string, error) {
	
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
		for _, stone := range self.Embeds {
			canonicalString += ":" + GetCanonicalMapString(stone["meta"].(map[string]interface{}))
		}
		canonicalString = strings.Trim(canonicalString, ":")
		break
	default:
		return "", errors.New("block unknown")
	}

	signer, err := crypto.ParsePrivateKey([]byte(privateKey))
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
func(self *Stone) AddMeta(meta map[string]interface{}, issuerPrivateKey string) error {

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
func (self *Stone) AddOwnership(ownership map[string]interface{}, issuerPrivateKey string) error {

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
func (self *Stone) AddAttributes(attributes map[string]interface{}, issuerPrivateKey string) error {
	
	self.Attributes = attributes

	// sign block
    _, err := self.Sign("attributes", issuerPrivateKey)
	if err != nil {
		return err
	}
	
	return nil
}

// add a stone to the `embeds` block
func (self *Stone) AddEmbed(stone *Stone, issuerPrivateKey string) error {

	self.Embeds = append(self.Embeds, stone.ToMap())

	// sign block
	_, err := self.Sign("embeds", issuerPrivateKey)
	if err != nil {
		return err
	}

	return nil
}


// checks if a block has a signature
func(self *Stone) HasSignature(blockName string) bool {
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
func(self *Stone) Verify(blockName, issuerPublicKey string) error {

	var canonicalString string

	switch blockName {
	case "meta":
		canonicalString = GetCanonicalMapString(self.Meta)
	case "ownership":
		canonicalString = GetCanonicalMapString(self.Ownership)
	case "attributes":
		canonicalString = GetCanonicalMapString(self.Attributes)
	case "embeds": 
		for _, stone := range self.Embeds {
			canonicalString += ":" + GetCanonicalMapString(stone["meta"].(map[string]interface{}))
		}
		canonicalString = strings.Trim(canonicalString, ":")
		break
	default:
		return errors.New("block name "+blockName+" is unknown")
	}

	signer, err := crypto.ParsePublicKey([]byte(issuerPublicKey))
	if err != nil {
		return errors.New(fmt.Sprintf("Public Key Error: %v", err))
	}

	// block has no signature
	if !self.HasSignature(blockName) {
		return errors.New("block `"+blockName+"` has no signature")
	}
	
	return signer.Verify([]byte(canonicalString), self.Signatures[blockName].(string))
}  

// checks if a stone object current state can
// pass as a valid stone
func(self *Stone) IsValid() error {
	return Validate(self.JSON())
}

// return stone as raw JSON string
func(self *Stone) JSON() string {
	bs, _ := json.Marshal(&self)
	return string(bs)
}

// returns a map representation of the stone
func(self *Stone) ToMap() map[string]interface{} {
	jsonStr := self.JSON()
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dat); err != nil {
        panic(err)
    }
    return dat
}

// return stone as a base64 json encoded string
func(self *Stone) Encode() string {
	jsonStr := self.JSON()
	return crypto.ToBase64([]byte(jsonStr))
}

// clone a stone
func(self *Stone) Clone() *Stone {
	jsonStr := self.JSON()
	stone, err := LoadJSON(jsonStr)
	if err != nil {
		panic(err)
	}
	return stone
}

// checks if the ownership block contains any property
func(self *Stone) HasOwnership() bool {
	return len(self.Ownership) > 0
}

// checks if the attributes block contains any property
func(self *Stone) HasAttributes() bool {
	return len(self.Attributes) > 0
}

// checks if the embeds block contains any property
func(self *Stone) HasEmbeds() bool {
	return len(self.Embeds) > 0
}