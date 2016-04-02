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

// known block names
var KnownBlockNames = []string{ "meta", "ownership", "attributes", "embeds" }

type Stone struct {
	Meta 		map[string]interface{}		`json:"meta"`
	Ownership 	map[string]interface{} 		`json:"ownership"`
	Embeds 		map[string]interface{} 		`json:"embeds"`
	Attributes 	map[string]interface{}		`json:"attributes"`
	Signatures 	map[string]interface{} 		`json:"signatures"`
}

// Initialize a stone
func initialize(stone *Stone) *Stone {
	stone.Meta 			= make(map[string]interface{})
	stone.Ownership 	= make(map[string]interface{})
	stone.Embeds 		= make(map[string]interface{})
	stone.Attributes 	= make(map[string]interface{})
	stone.Signatures 	= make(map[string]interface{})
	return stone
}

// creates an stone instances and initializes it
func Empty() *Stone {
	sh := &Stone{}
	return initialize(sh)
}

// Create a stone with a inital meta block
// The new stone is immediately signed using the issuer's private key
func Create(meta map[string]interface{}, issuerPrivateKey string) (*Stone, error) {

	stone := initialize(&Stone{})

	// validate meta
	if err := ValidateMetaBlock(meta); err != nil {
    	return &Stone{}, err
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

	var stone = initialize(&Stone{})

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
    	stone.Embeds = data["embeds"].(map[string]interface{})
    }

    return stone, nil
}

// Creates a new stone from a json string. It does not attempt to sign the blocks. 
func Load(stoneStr string) (*Stone, error) {

	// empty string not allowed
	stoneStr = strings.TrimSpace(stoneStr)
	if stoneStr == "" {
		return &Stone{}, errors.New("Cannot load empty string")
	} 

	return LoadJSON(stoneStr)
}


// Create a stone from a json string by converting
// it to a map and then used to load a new stone instance
func LoadJSON(jsonStr string) (*Stone, error) {

	data, err := JSONToMap(jsonStr)
	if err != nil{
        return &Stone{}, err;
    }
    
    // validate...
    if err := Validate(data); err != nil {
    	return &Stone{}, err
    }

	return loadMap(data)  
}

// Decode a base64 encoded stone token
func Decode(encStone string) (*Stone, error) {

	var stone = Empty()

	// decode encoded token to get the signatures
	sigJSON, err := crypto.FromBase64Raw(encStone)
	if err != nil {
		return &Stone{}, errors.New("failed to decode token")
	}

	// convert json to map
	stoneMap, err := JSONToMap(sigJSON)
	if err != nil {
		return &Stone{}, errors.New("failed to parse token")
	}

	// parse and load meta block
	if IsStringValue(stoneMap["meta"]) && stoneMap["meta"].(string) != "" {

		var token = stoneMap["meta"].(string);

		block, err := TokenToBlock(token, "meta")
		if err != nil {
			return stone, err
		}

		stone.Meta = block
		stone.Signatures["meta"] = token
	}

	// parse and load ownership block
	if IsStringValue(stoneMap["ownership"]) && stoneMap["ownership"].(string) != "" {

		var token = stoneMap["ownership"].(string);

		block, err := TokenToBlock(token, "ownership")
		if err != nil {
			return stone, err
		}

		stone.Ownership = block
		stone.Signatures["ownership"] = token
	}

	// parse and load ownership block
	if IsStringValue(stoneMap["attributes"]) && stoneMap["attributes"].(string) != "" {

		var token = stoneMap["attributes"].(string);

		block, err := TokenToBlock(token, "attributes")
		if err != nil {
			return stone, err
		}

		stone.Attributes = block
		stone.Signatures["attributes"] = token
	}

	// parse and load embeds block
	if IsStringValue(stoneMap["embeds"]) && stoneMap["embeds"].(string) != "" {

		var token = stoneMap["embeds"].(string);

		block, err := TokenToBlock(token, "embeds")
		if err != nil {
			return stone, err
		}

		stone.Embeds = block
		stone.Signatures["embeds"] = token
	}	

	return stone, nil
}

// Decodes a payload of a JWS token to a block
func TokenToBlock(token string, blockName string) (map[string]interface{}, error) {

	var block = map[string]interface{}{}

	var payload, err = GetJWSPayload(token)
	if err != nil {
		return block, err
	}
	
	blockJSON, err := crypto.FromBase64Raw(payload)
	if err != nil {
		return block, errors.New("invalid "+blockName+" token")
	}

	block, err = JSONToMap(blockJSON)
	if err != nil {
		return block, errors.New("malformed "+blockName+" block")
	}

	return block, nil
}


// get a block or panic of block is unknown
func(self *Stone) getBlock(name string) map[string]interface{} {
	if name == "meta" { return self.Meta }
	if name == "ownership" { return self.Ownership }
	if name == "attributes" { return self.Attributes }
	if name == "embeds" { return self.Embeds }
	panic("unknown block")
}

// Sign a block. The signing process takes the value of a block and signs
// it using JWS. The signature generated is included in the 
// `signatures` block. If a block is empty or unknown, an error is returned.
func(self *Stone) Sign(blockName string, privateKey string) (string, error) {
	
	var block map[string]interface{}

	signer, err := crypto.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return "", errors.New("Private Key Error: " + err.Error())
	}

	// block name must be known
	if !InStringSlice(KnownBlockNames, blockName) {
		return "", errors.New("block unknown")
	}

	block = self.getBlock(blockName)
	if IsMapEmpty(block) {
		return "", errors.New("failed to sign empty block")
	}

	// sign block
	payload, _ := MapToJSON(block)
	signature, err := signer.JWS_RSA_Sign(payload)
	if err != nil {
		return "", errors.New("failed to sign block")
	}
	
	self.Signatures[blockName] = signature
	return signature, nil
}


// Verify a block. Using the public of the signer
func(self *Stone) Verify(blockName, signerPublicKey string) error {

	signer, err := crypto.ParsePublicKey([]byte(signerPublicKey))
	if err != nil {
		return errors.New(fmt.Sprintf("Public Key Error: %v", err))
	}

	// block name must be known
	if !InStringSlice(KnownBlockNames, blockName) {
		return errors.New("block unknown")
	}

	// ensure block has signature
	if !self.HasSignature(blockName) {
		return errors.New("`"+blockName+"` block has no signature")
	}

	// verify
	_, err = signer.JWS_RSA_Verify(self.Signatures[blockName].(string))
	if err != nil {
		return errors.New(fmt.Sprintf("`%s` block signature could not be verified", blockName))
	}

	return nil
}  

// Encode a base64 url equivalent of the signatures.
func(self *Stone) Encode() string {
	var signaturesStr, _ = MapToJSON(self.Signatures)
	return crypto.ToBase64([]byte(signaturesStr))
}

// Add meta block. Validation and block signing are carried out
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

// Assign and sign a valid ownership data to the ownership block.
// `meta.id` property must be set. 
func (self *Stone) AddOwnership(ownership map[string]interface{}, issuerPrivateKey string) error {

	if self.Meta["id"] == nil || (self.Meta["id"] != nil && strings.TrimSpace(self.Meta["id"].(string)) == "") {
		return errors.New("meta.id is not set")
	}

	// validate 
	if err := ValidateOwnershipBlock(ownership, self.Meta["id"].(string)); err != nil {
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

// Assign and sign a valid attribute data to the attributes block.
// `meta.id` property must be set. 
func (self *Stone) AddAttributes(attributes map[string]interface{}, issuerPrivateKey string) error {
	
	if self.Meta["id"] == nil || (self.Meta["id"] != nil && strings.TrimSpace(self.Meta["id"].(string)) == "") {
		return errors.New("meta.id is not set")
	}

	// validate 
	if err := ValidateAttributesBlock(attributes, self.Meta["id"].(string)); err != nil {
    	return err
    }

	self.Attributes = attributes

	// sign block
    _, err := self.Sign("attributes", issuerPrivateKey)
	if err != nil {
		return err
	}
	
	return nil
}

// add a stone to the `embeds` block
func (self *Stone) AddEmbed(embeds map[string]interface{}, issuerPrivateKey string) error {

	if self.Meta["id"] == nil || (self.Meta["id"] != nil && strings.TrimSpace(self.Meta["id"].(string)) == "") {
		return errors.New("meta.id is not set")
	}

	// validate 
	if err := ValidateEmbedsBlock(embeds, self.Meta["id"].(string)); err != nil {
    	return err
    }

	self.Embeds = embeds

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
		return self.Signatures[blockName] != nil
		break
	default:
		return false
	}
	return false
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
	var dat = make(map[string]interface{})
	dat["signatures"] = self.Signatures
	dat["meta"] = self.Meta
	dat["ownership"] = self.Ownership
	dat["attributes"] = self.Attributes
	dat["embeds"] = self.Embeds
    return dat
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