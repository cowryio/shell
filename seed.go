// A seed represents a container of information
// that is considered valueable.
package seed

import (
	"encoding/json"
	"errors"
	"strings"
	"fmt"
	"github.com/ellcrys/crypto"
)

type Seed struct {
	Signatures map[string]interface{} 		`json:"signatures"`
	Meta map[string]interface{}				`json:"meta"`
	Ownership map[string]interface{} 		`json:"ownership"`
	Embeds []map[string]interface{} 		`json:"embeds"`
	Attributes map[string]interface{}		`json:"attributes"`
}

// Initialize a seed
func initialize(seed *Seed) *Seed {
	seed.Signatures = make(map[string]interface{})
	seed.Meta = make(map[string]interface{})
	seed.Ownership = make(map[string]interface{})
	seed.Embeds = []map[string]interface{}{}
	seed.Attributes = make(map[string]interface{})
	return seed
}

// creates an seed instances and initializes it
func Empty() *Seed {
	sh := &Seed{}
	return initialize(sh)
}

// Create a seed.The new seed is immediately signed using the issuer's private key
func Create(meta map[string]interface{}, issuerPrivateKey string) (*Seed, error) {

	seed := initialize(&Seed{})

	// validate meta
	if err := ValidateMetaBlock(meta); err != nil {
    	return &Seed{}, err
    } else {
		meta["created_at"] = IntToFloat64(meta["created_at"])
    }

    // set seed Meta field and create a meta signature
	seed.Meta = meta
	_, err := seed.Sign("meta", issuerPrivateKey)
	if err != nil {
		return &Seed{}, err
	}

	return seed, nil
}

// Creates a seed from a map. Validation of expected field is 
// not performed. Use Validate() before calling this method if validation
// is necessary
func loadMap(data map[string]interface{}) (*Seed, error) {

	var seed = &Seed{}

	// add signatures
    if signatures := data["signatures"]; signatures != nil {
    	seed.Signatures = data["signatures"].(map[string]interface{})
    }

    // add meta
    if meta := data["meta"]; meta != nil {
    	seed.Meta = data["meta"].(map[string]interface{})
    }

    // add ownership
    if ownership := data["ownership"]; ownership != nil {
    	seed.Ownership = data["ownership"].(map[string]interface{})
    }

    // add attributes
    if attributes := data["attributes"]; attributes != nil {
    	seed.Attributes = data["attributes"].(map[string]interface{})
    }

    // add embeds
    if embeds := data["embeds"]; embeds != nil {
    	for _, m := range embeds.([]interface{}) {
    		seed.Embeds = append(seed.Embeds, m.(map[string]interface{}))
    	}
    }

    return seed, nil
}

// Creates a new seed from a raw json or base 64 encoded json string. It does not 
// attempt to sign the blocks. 
// If the string passed in starts with "{", it is considered a JSON string, otherwise, it assumes string is base 64 encoded and
// will attempt to decoded it. 
func Load(seedStr string) (*Seed, error) {
	seedStr = strings.TrimSpace(seedStr)
	if seedStr == "" {
		return &Seed{}, errors.New("Cannot load empty seed string")
	} else {
		if fmt.Sprintf("%c", seedStr[0]) == "{" {					// json string
			return LoadJSON(seedStr)
		} else {
			decodedSeedStr, err := crypto.FromBase64(seedStr)
			if err != nil {
				return &Seed{}, errors.New("unable to decode encoded seed string")
			}
			return LoadJSON(decodedSeedStr)
		}
	}
}


// Create a seed from a json string by converting
// it to a map and then used to load a new seed instance
func LoadJSON(jsonStr string) (*Seed, error) {
	data, err := JSONToMap(jsonStr)
	if err != nil{
        return &Seed{}, err;
    }
    if err := Validate(data); err != nil {
    	return &Seed{}, err
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


// Sign any seed block by creating a canonical string representation
// of the block value and signing with the issuer's private key. The computed signature
// is store the `signatures` block
func(self *Seed) Sign(blockName string, privateKey string) (string, error) {
	
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
		for _, seed := range self.Embeds {
			canonicalString += ":" + GetCanonicalMapString(seed["meta"].(map[string]interface{}))
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
func(self *Seed) AddMeta(meta map[string]interface{}, issuerPrivateKey string) error {

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
func (self *Seed) AddOwnership(ownership map[string]interface{}, issuerPrivateKey string) error {

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
func (self *Seed) AddAttributes(attributes map[string]interface{}, issuerPrivateKey string) error {
	
	self.Attributes = attributes

	// sign block
    _, err := self.Sign("attributes", issuerPrivateKey)
	if err != nil {
		return err
	}
	
	return nil
}

// add a seed to the `embeds` block
func (self *Seed) AddEmbed(seed *Seed, issuerPrivateKey string) error {

	self.Embeds = append(self.Embeds, seed.ToMap())

	// sign block
	_, err := self.Sign("embeds", issuerPrivateKey)
	if err != nil {
		return err
	}

	return nil
}


// checks if a block has a signature
func(self *Seed) HasSignature(blockName string) bool {
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
func(self *Seed) Verify(blockName, issuerPublicKey string) error {

	var canonicalString string

	switch blockName {
	case "meta":
		canonicalString = GetCanonicalMapString(self.Meta)
	case "ownership":
		canonicalString = GetCanonicalMapString(self.Ownership)
	case "attributes":
		canonicalString = GetCanonicalMapString(self.Attributes)
	case "embeds": 
		for _, seed := range self.Embeds {
			canonicalString += ":" + GetCanonicalMapString(seed["meta"].(map[string]interface{}))
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

// checks if a seed object current state can
// pass as a valid seed
func(self *Seed) IsValid() error {
	return Validate(self.JSON())
}

// return seed as raw JSON string
func(self *Seed) JSON() string {
	bs, _ := json.Marshal(&self)
	return string(bs)
}

// returns a map representation of the seed
func(self *Seed) ToMap() map[string]interface{} {
	jsonStr := self.JSON()
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dat); err != nil {
        panic(err)
    }
    return dat
}

// return seed as a base64 json encoded string
func(self *Seed) Encode() string {
	jsonStr := self.JSON()
	return crypto.ToBase64([]byte(jsonStr))
}

// clone a seed
func(self *Seed) Clone() *Seed {
	jsonStr := self.JSON()
	seed, err := LoadJSON(jsonStr)
	if err != nil {
		panic(err)
	}
	return seed
}

// checks if the ownership block contains any property
func(self *Seed) HasOwnership() bool {
	return len(self.Ownership) > 0
}

// checks if the attributes block contains any property
func(self *Seed) HasAttributes() bool {
	return len(self.Attributes) > 0
}

// checks if the embeds block contains any property
func(self *Seed) HasEmbeds() bool {
	return len(self.Embeds) > 0
}