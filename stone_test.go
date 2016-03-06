package stone

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/ellcrys/crypto"
)

var sampleKeys = []string{
	"-----BEGIN RSA PRIVATE KEY-----\nMIICWwIBAAKBgQCroZieOAo9stcf6R6eWfo51VCvK8cLdNS577m/HIFOmEd1CDi/\nu7agGzpehNAhHpr5NVjQZ4Te+KMRn9SnpUK2hc8dUU25PQolsOEwePVQ18hHNK4Y\n2JvOY/f8KCO2hhrS6uuP6eedpnSdulS1OXHTL6ZxQmBd9F33gLT6BERHQwIDAQAB\nAoGAEZ/0ljrXAmL9KG++DzDaO1omgPaT6B9FQRrXDkMVHEcS/3eqrDXQmTxykAY/\ngUctTu4lgrE+uc76n/Kz2ctkwEKIKet56ylqp+wlEUt1G+udoi07tgd7XyxzoUJm\nZwSm89gKh+mEPxni0FrBNg6dR0n2gvKRecnXqyoGVOHZITECQQDXgRJyrzgc/JhB\nSOBznEjtXAZXRRu3o9UznztjU9Xz7NWXTVuHu8WqYmGWCOqnysMhXJ3xBddJyDTF\njuOJ0123AkEAy+H+3POcT2FDOuluqPmAZQAUU6Nxtbj02/JJtOy7jq5jnN27HVC3\nuQzmfsS5J2XeQQodOUwOy2Ub57/OMrMi1QJAGZsZgQz2wuL0iFVLbhE0zRcxHa91\ncqWB0Kdr3Ap7EoeifV7QsFkMTIlyBOy8TQGXm+AwWBIUmYyzUIIA4UB/EwJAO+Bo\nSB2nZ0yqQO/zVt7HjWIDljinGXZzOvEiImdwAcxHZvdbj5V4D3mxa8N8mQx6xGEj\nCgPDSIquMlaLSSqA7QJAAbQPa0frCkm1rkWWZ7QwGm7ptzOACwFEGefm/1mhmw3a\nvoWRTHhrDuEbeVH3iF8MWhLJLPFtuSShiQMsrVbXPA==\n-----END RSA PRIVATE KEY-----",
	"-----BEGIN KEY-----\nMIICWwIBAAKBgQCroZieOAo9stcf6R6eWfo51VCvK8cLdNS577m/HIFOmEd1CDi/\nu7agGzpehNAhHpr5NVjQZ4Te+KMRn9SnpUK2hc8dUU25PQolsOEwePVQ18hHNK4Y\n2JvOY/f8KCO2hhrS6uuP6eedpnSdulS1OXHTL6ZxQmBd9F33gLT6BERHQwIDAQAB\nAoGAEZ/0ljrXAmL9KG++DzDaO1omgPaT6B9FQRrXDkMVHEcS/3eqrDXQmTxykAY/\ngUctTu4lgrE+uc76n/Kz2ctkwEKIKet56ylqp+wlEUt1G+udoi07tgd7XyxzoUJm\nZwSm89gKh+mEPxni0FrBNg6dR0n2gvKRecnXqyoGVOHZITECQQDXgRJyrzgc/JhB\nSOBznEjtXAZXRRu3o9UznztjU9Xz7NWXTVuHu8WqYmGWCOqnysMhXJ3xBddJyDTF\njuOJ0123AkEAy+H+3POcT2FDOuluqPmAZQAUU6Nxtbj02/JJtOy7jq5jnN27HVC3\nuQzmfsS5J2XeQQodOUwOy2Ub57/OMrMi1QJAGZsZgQz2wuL0iFVLbhE0zRcxHa91\ncqWB0Kdr3Ap7EoeifV7QsFkMTIlyBOy8TQGXm+AwWBIUmYyzUIIA4UB/EwJAO+Bo\nSB2nZ0yqQO/zVt7HjWIDljinGXZzOvEiImdwAcxHZvdbj5V4D3mxa8N8mQx6xGEj\nCgPDSIquMlaLSSqA7QJAAbQPa0frCkm1rkWWZ7QwGm7ptzOACwFEGefm/1mhmw3a\nvoWRTHhrDuEbeVH3iF8MWhLJLPFtuSShiQMsrVbXPA==\n-----END RSA PRIVATE KEY-----",
	"-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCroZieOAo9stcf6R6eWfo51VCv\nK8cLdNS577m/HIFOmEd1CDi/u7agGzpehNAhHpr5NVjQZ4Te+KMRn9SnpUK2hc8d\nUU25PQolsOEwePVQ18hHNK4Y2JvOY/f8KCO2hhrS6uuP6eedpnSdulS1OXHTL6Zx\nQmBd9F33gLT6BERHQwIDAQAB\n-----END PUBLIC KEY-----",
}

var TEST_SHELL_DATA = []string {
	`{"signatures":{"meta":""},"meta":{"created_at":1453975575, "stone_id":"4417781906fb0a89c295959b0df01782dbc4dc9f","stone_type":"currency"},"ownership":null,"embeds":[],"attributes":null}`,
	`{"signatures":{"meta":"abcde","ownership":"abcde","attributes":"abcde","embeds":"abcde"},"meta":{"created_at": `+IntToString(time.Now().Unix())+`,"stone_id":"4417781906fb0a89c295959b0df01782dbc4dc9f","stone_type":"currency"},"ownership":{"type":"sole","sole":{"address_id":"abcde"},"status":"transferred"},"embeds":[{"signatures":{"meta":"abcde","ownership":"abcde"},"meta":{"created_at":1454443443,"stone_id":"9417781906fb0a89c295959b0df01782dbc4dc9f","stone_type":"currency"},"ownership":{"type":"sole","sole":{"address_id":"abcde"},"status":"transferred"},"embeds":[{"signatures":{"meta":"abcde","ownership":"abcde"},"meta":{"created_at":1454443443,"stone_id":"514417781906fb0a89c295959b0df01782dbc4dc9f","stone_type":"currency"},"ownership":{"type":"sole","sole":{"address_id":"abcde"},"status":"transferred"},"embeds":[],"attributes":{}}],"attributes":{}}],"attributes":{"some_data":"some_value"}}`,
}

func NewValidStone() *Stone {
	var meta = map[string]interface{}{
		"stone_id": NewID(),
		"stone_type": "some_stone",
		"created_at": time.Now().Unix(),
	}
	sh, _ := Create(meta, sampleKeys[0])
	return sh
} 

// TestCreateAStone create a valid, error free stone
func TestCreateAStone(t *testing.T) {
	stoneID := NewID()
	var meta = map[string]interface{}{
		"stone_id": stoneID,
		"stone_type": "currency",
		"created_at": time.Now().Unix(),
	}
	sh, err := Create(meta, sampleKeys[0])
	assert.Nil(t, err)
	assert.Equal(t, sh.Meta["stone_id"], stoneID)
	assert.NotEmpty(t, sh.Signatures["meta"])
}

// TestMustProvideMetaWithContent test that a map describing the `meta` block is required
func TestMustProvideMetaWithContent(t *testing.T) {
	_, err := Create(make(map[string]interface{}), sampleKeys[0])
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "`meta` block is missing `stone_id` property")
}

// TestInvalidPrivateKey tests that an invalid private key returns an error
func TestInvalidPrivateKey(t *testing.T) {
	var issuerPrivateKey = sampleKeys[1]
	stoneID := NewID()
	var meta = map[string]interface{}{
		"stone_id": stoneID,
		"stone_type": "currency",
		"created_at": time.Now().Unix(),
	}
	_, err := Create(meta, issuerPrivateKey)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), `Private Key Error: unsupported key type "KEY"`)
}

// TestCantLoadMalformedJSON tests that a malformed JSON string will produce an error
func TestCantLoadMalformedJSON(t *testing.T) {
	var txt = `{ "signatures": [ }`
	_, err := LoadJSON(txt);
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), `unable to parse json string`)
}

// TestLoadJSON tests that a valid stone json string can be loaded into a stone object
func TestLoadJSON(t *testing.T) {
	txt := TEST_SHELL_DATA[0]
	stone, err := LoadJSON(txt)
	assert.Nil(t, err);
	assert.IsType(t, &Stone{}, stone)
}

// TestLoadEncodedJSON tests that a base 64 encoded json string can be loaded into a stone object
func TestLoadEncodedJSON(t *testing.T) {
	encodedJSON := crypto.ToBase64([]byte(TEST_SHELL_DATA[0]))
	stone, err := Load(encodedJSON)
	assert.Nil(t, err);
	assert.IsType(t, &Stone{}, stone)
}

// TestCannotLoadInvalidEncodedJSON tests that an incorrect base64 encode will result in error when attempting to load into a stone
func TestCannotLoadInvalidEncodedJSON(t *testing.T) {
	encodedJSON := crypto.ToBase64([]byte("abcde"))
	_, err := Load(encodedJSON)
	assert.NotNil(t, err);
	assert.Equal(t, err.Error(), "unable to parse json string")
}

// TestCorrectlySignMeta tests that a stone is correctly signed
func TestCorrectlySignMeta(t *testing.T) {
	txt := TEST_SHELL_DATA[0]
	stone, err := LoadJSON(txt)
	assert.Nil(t, err);
	expectedCanonicalMapString := GetCanonicalMapString(stone.Meta)
	signer, err := crypto.ParsePrivateKey([]byte(sampleKeys[0]))
	assert.Nil(t, err)
	expectedSignature, err := signer.Sign([]byte(expectedCanonicalMapString))
	assert.Nil(t, err)
	signature, err := stone.Sign("meta", sampleKeys[0])
	assert.Nil(t, err)
	assert.Equal(t, expectedSignature, signature)
	assert.Equal(t, expectedSignature, stone.Signatures["meta"])
}

// TestCannotSignUnknownBlock tests that an error will occur when attempting to sign an unknown block
func TestCannotSignUnknownBlock(t *testing.T) {
	txt := TEST_SHELL_DATA[0]
	stone, err := LoadJSON(txt)
	assert.Nil(t, err)
	_, err = stone.Sign("unknown_block", sampleKeys[0])
	assert.NotNil(t, err)
	expectedMsg := "block unknown"
	assert.Equal(t, err.Error(), expectedMsg)
}

// TestAddMeta tests that a `meta` block can be assigned and signed successful
func TestAddMeta(t *testing.T) {
	stoneID := NewID()
	var meta = map[string]interface{}{
		"stone_id": stoneID,
		"stone_type": "currency",
		"created_at": time.Now().Unix(),
	}
	sh := Empty()
	err := sh.AddMeta(meta, sampleKeys[0])
	assert.Nil(t, err)
	assert.Equal(t, sh.Meta["stone_id"], meta["stone_id"])
	assert.NotNil(t, sh.Signatures["meta"])
}

// TestAddOwnership tests that the `ownership` block is assigned and signed successfully
func TestAddOwnership(t *testing.T) {
	var ownership = map[string]interface{}{
		"type": "sole",
   		"sole": map[string]interface{}{
   			"address_id": "abcde",
   		},
   		"status": "transferred",
	}
	sh := Empty()
	err := sh.AddOwnership(ownership, sampleKeys[0])
	assert.Nil(t, err)
	assert.Equal(t, sh.Ownership["type"], ownership["type"])
	assert.NotNil(t, sh.Signatures["ownership"])
}

// TestAddOwnership tests that the `attributes` block is assigned and signed successfully
func TestAddAttributes(t *testing.T) {
	var attrs = map[string]interface{}{
		"some_data": "some_value",
	}
	sh := Empty()
	err := sh.AddAttributes(attrs, sampleKeys[0])
	assert.Nil(t, err)
	assert.Equal(t, sh.Attributes["some_data"], attrs["some_data"])
	assert.NotNil(t, sh.Signatures["attributes"])
}

// TestAddEmbed tests that a stone object can be embeded into 
// another stone with no error and is also signed
func TestAddEmbed(t *testing.T) {
	sh := NewValidStone()
	embed := NewValidStone()
	sh.AddEmbed(embed, sampleKeys[0])
	expectedSignature, err := sh.Sign("embeds", sampleKeys[0])
	assert.Nil(t, err)
	assert.Equal(t, len(sh.Embeds), 1)
	assert.Exactly(t, sh.Embeds[0], embed.ToMap())
	assert.NotNil(t, sh.Signatures["embeds"])
	assert.Equal(t, sh.Signatures["embeds"], expectedSignature)
}

// TestHashSignature tests that an attribute does not or has a signature
func TestHashSignature(t *testing.T) {
	var attrs = map[string]interface{}{
		"some_data": "some_value",
	}
	sh := Empty()
	assert.Equal(t, sh.HasSignature("attributes"), false)
	assert.Equal(t, sh.HasSignature("ownership"), false)
	err := sh.AddAttributes(attrs, sampleKeys[0])
	assert.Nil(t, err)
	assert.Equal(t, sh.HasSignature("attributes"), true)
	assert.Equal(t, sh.HasSignature("ownership"), false)
}

// TestCallVerifyWithUnknownBlockName tests that an error will occur when verifying an unknown block
func TestCallVerifyWithUnknownBlockName(t *testing.T) {
	var attrs = map[string]interface{}{
		"some_data": "some_value",
	}
	sh := Empty()
	err := sh.AddAttributes(attrs, sampleKeys[0])
	assert.Nil(t, err)
	err = sh.Verify("some_block", sampleKeys[2])
	assert.NotNil(t, err)
	expectedMsg := "block name some_block is unknown"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestCallVerifyWithInvalidPublicKey tests that an error will occur when verifying with an invalid public key
func TestCallVerifyWithInvalidPublicKey(t *testing.T) {
	var attrs = map[string]interface{}{
		"some_data": "some_value",
	}
	sh := Empty()
	err := sh.AddAttributes(attrs, sampleKeys[0])
	assert.Nil(t, err)
	err = sh.Verify("attributes", sampleKeys[1])
	assert.NotNil(t, err)
	expectedMsg := `Public Key Error: unsupported key type "KEY"`
	assert.Equal(t, expectedMsg, err.Error())
}


// TestCallVerifyOnBlockWithNoSignature tests that an error will occur when verifying a block with no signature
// in the signatures block
func TestCallVerifyOnBlockWithNoSignature(t *testing.T) {
	sh := Empty()
	err := sh.Verify("attributes", sampleKeys[2])
	assert.NotNil(t, err)
	expectedMsg := "block `attributes` has no signature"
	assert.Equal(t, expectedMsg, err.Error())
}


// TestCallVerifyWhenBlockSignatureHexEncodeIsInvalid tests that an error will occur when verifying a block that has
// a signature that cannot be decoded from it hex encoded variation
func TestCallVerifyWhenBlockSignatureHexEncodeIsInvalid(t *testing.T) {
	sh := Empty()
	sh.Signatures["attributes"] = "abcdefaa9*"
	err := sh.Verify("attributes", sampleKeys[2])
	assert.NotNil(t, err)
	expectedMsg := "invalid signature: unable to decode from hex to string"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestCallVerifyWhenBlockSignatureInvalid tests that an error will occur when verifying a block 
// that has an invalid signature
func TestCallVerifyWhenBlockSignatureInvalid(t *testing.T) {
	sh := Empty()
	sh.Signatures["attributes"] = "abcdef"
	err := sh.Verify("attributes", sampleKeys[2])
	assert.NotNil(t, err)
	expectedMsg := "crypto/rsa: verification error"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestVerifyMeta tests that a meta block signed with a private key is 
// successfully verified using the corresponding public key
func TestVerifyMeta(t *testing.T) {
	stone, err := LoadJSON(TEST_SHELL_DATA[1]);
	assert.Nil(t, err)
	stone.Sign("meta", sampleKeys[0])
	err = stone.Verify("meta", sampleKeys[2])
	assert.Nil(t, err)
}

// TestVerifyOwnership tests that an ownership block signed with a private key is 
// successfully verified using the corresponding public key
func TestVerifyOwnership(t *testing.T) {
	stone, err := LoadJSON(TEST_SHELL_DATA[1]);
	assert.Nil(t, err)
	stone.Sign("ownership", sampleKeys[0])
	err = stone.Verify("ownership", sampleKeys[2])
	assert.Nil(t, err)
}

// TestVerifyAttributes tests that an `attributes` block signed with a private key is 
// successfully verified using the corresponding public key
func TestVerifyAttributes(t *testing.T) {
	stone, err := LoadJSON(TEST_SHELL_DATA[1]);
	assert.Nil(t, err)
	stone.Sign("attributes", sampleKeys[0])
	err = stone.Verify("attributes", sampleKeys[2])
	assert.Nil(t, err)
}

// TestVerifyEmbeds tests that an `attributes` block signed with a private key is 
// successfully verified using the corresponding public key
func TestVerifyEmbeds(t *testing.T) {
	stone, err := LoadJSON(TEST_SHELL_DATA[1]);
	assert.Nil(t, err)
	stone.Sign("embeds", sampleKeys[0])
	err = stone.Verify("embeds", sampleKeys[2])
	assert.Nil(t, err)
}

// TestCloneStone
func TestCloneStone(t *testing.T) {
	stone, err := LoadJSON(TEST_SHELL_DATA[1]);
	assert.Nil(t, err)
	clone := stone.Clone()
	assert.Exactly(t, stone, clone) 
	stone.Signatures["meta"] = "blah_blah"
	assert.NotEmpty(t, stone.Signatures["meta"], clone.Signatures["meta"])
}

// TestHasOwnershipFalse tests that a stone does not have any ownership information
func TestHasOwnershipFalse(t *testing.T) {
	stone, err := LoadJSON(TEST_SHELL_DATA[0]);
	assert.Nil(t, err)
	assert.Equal(t, stone.HasOwnership(), false)
}

// TestHasOwnershipTrue tests that a stone has ownership information
func TestHasOwnershipTrue(t *testing.T) {
	stone, err := LoadJSON(TEST_SHELL_DATA[0]);
	assert.Nil(t, err)
	var ownership = map[string]interface{}{
		"type": "sole",
   		"sole": map[string]interface{}{
			"address_id": "abcde",
   		},
	}
	err = stone.AddOwnership(ownership, sampleKeys[0])
	assert.Nil(t, err)
	assert.Equal(t, stone.HasOwnership(), true)
}

// TestHasAttributesTrue tests that a stone has attributes information
func TestHasAttributesTrue(t *testing.T) {
	stone, err := LoadJSON(TEST_SHELL_DATA[0]);
	assert.Nil(t, err)
	var attrs = map[string]interface{}{
		"some_data": "some_value",
	}
	err = stone.AddAttributes(attrs, sampleKeys[0])
	assert.Nil(t, err)
	assert.Equal(t, stone.HasAttributes(), true)
}

// TestHasAttributesFalse tests that a stone does not have attributes information
func TestHasAttributesFalse(t *testing.T) {
	stone, err := LoadJSON(TEST_SHELL_DATA[0]);
	assert.Nil(t, err)
	assert.Equal(t, stone.HasAttributes(), false)
}

// TestHasEmbedsTrue tests that a stone has embeds information
func TestHasEmbedsTrue(t *testing.T) {
	sh := NewValidStone()
	embed := NewValidStone()
	err := sh.AddEmbed(embed, sampleKeys[0])
	assert.Nil(t, err)
	assert.Equal(t, sh.HasEmbeds(), true)
}

// TestHasEmbedsFalse tests that a stone has no embeds information
func TestHasEmbedsFalse(t *testing.T) {
	sh := NewValidStone()
	assert.Equal(t, sh.HasEmbeds(), false)
}