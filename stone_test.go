package stone

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/ellcrys/crypto"
	"github.com/ellcrys/util"
)

func NewValidStone() *Stone {
	var meta = map[string]interface{}{
		"id": util.NewID(),
		"type": "some_stone",
		"created_at": time.Now().Unix(),
	}
	sh, _ := Create(meta, util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	return sh
} 

// TestCreateAStone create a valid, error free stone
func TestCreateAStone(t *testing.T) {
	stoneID := util.NewID()
	var meta = map[string]interface{}{
		"id": stoneID,
		"type": "currency",
		"created_at": time.Now().Unix(),
	}
	sh, err := Create(meta, util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, sh.Meta["id"], stoneID)
	assert.NotEmpty(t, sh.Signatures["meta"])
}

// TestMustProvideMetaWithContent test that a map describing the `meta` block is required
func TestMustProvideMetaWithContent(t *testing.T) {
	_, err := Create(make(map[string]interface{}), util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "`meta` block is missing `id` property")
}

// TestInvalidPrivateKey tests that an invalid private key returns an error
func TestInvalidPrivateKey(t *testing.T) {
	var issuerPrivateKey = util.ReadFromFixtures("tests/fixtures/rsa_invalid_1.txt")
	stoneID := util.NewID()
	var meta = map[string]interface{}{
		"id": stoneID,
		"type": "currency",
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
	txt := util.ReadFromFixtures("tests/fixtures/stone_1.json")
	stone, err := LoadJSON(txt)
	assert.Nil(t, err);
	assert.IsType(t, &Stone{}, stone)
}

// TestCorrectlySignMeta tests that a stone is correctly signed
func TestCorrectlySignMeta(t *testing.T) {
	txt := util.ReadFromFixtures("tests/fixtures/stone_1.json")
	stone, err := LoadJSON(txt)
	assert.Nil(t, err);
	expectedSignature := "eyJhbGciOiJSUzI1NiIsImp3ayI6eyJrdHkiOiJSU0EiLCJuIjoicTZHWW5qZ0tQYkxYSC1rZW5sbjZPZFZRcnl2SEMzVFV1ZS01dnh5QlRwaEhkUWc0djd1Mm9CczZYb1RRSVI2YS1UVlkwR2VFM3ZpakVaX1VwNlZDdG9YUEhWRk51VDBLSmJEaE1IajFVTmZJUnpTdUdOaWJ6bVAzX0NnanRvWWEwdXJyai1ubm5hWjBuYnBVdFRseDB5LW1jVUpnWGZSZDk0QzAtZ1JFUjBNIiwiZSI6IkFRQUIifX0.eyJjcmVhdGVkX2F0IjoxNDUzOTc1NTc1LCJpZCI6IjQ0MTc3ODE5MDZmYjBhODljMjk1OTU5YjBkZjAxNzgyZGJjNGRjOWYiLCJ0eXBlIjoiY3VycmVuY3kifQ.pEBlRBlIkmrMNJkBlvUWo5FK8N6-G83hirDNQLmYo6ojSkX0cXqak_mdHo7zUyLV0CxAvPuxb9fiYbz4S2tllIMpHm_RHQDDOXkl1ykiUrbcotrlfQmiOqvDzp91IL38m8Uy8-MBg-JB7K9nacCCLEph-BLn83AyyQeSVTQZGKo"
	signature, err := stone.Sign("meta", util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, expectedSignature, signature)
	assert.Equal(t, expectedSignature, stone.Signatures["meta"])
}

// TestCannotSignUnknownBlock tests that an error will occur when attempting to sign an unknown block
func TestCannotSignUnknownBlock(t *testing.T) {
	txt := util.ReadFromFixtures("tests/fixtures/stone_1.json")
	stone, err := LoadJSON(txt)
	assert.Nil(t, err)
	_, err = stone.Sign("unknown_block", util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	assert.NotNil(t, err)
	expectedMsg := "block unknown"
	assert.Equal(t, err.Error(), expectedMsg)
}

// TestCannotSignEmptyBlock tests that an error will occur when attempting to sign an empty block
func TestCannotSignEmptyBlock(t *testing.T) {
	txt := util.ReadFromFixtures("tests/fixtures/stone_1.json")
	stone, err := LoadJSON(txt)
	assert.Nil(t, err)
	_, err = stone.Sign("ownership", util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	assert.NotNil(t, err)
	expectedMsg := "failed to sign empty block"
	assert.Equal(t, err.Error(), expectedMsg)
}

// TestAddMeta tests that a `meta` block can be assigned and signed successful
func TestAddMeta(t *testing.T) {
	stoneID := util.NewID()
	var meta = map[string]interface{}{
		"id": stoneID,
		"type": "currency",
		"created_at": time.Now().Unix(),
	}
	sh := Empty()
	err := sh.AddMeta(meta, util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, sh.Meta["id"], meta["id"])
	assert.NotNil(t, sh.Signatures["meta"])
}

// TestAddOwnershipWithUnsetMetaID tests that an error will occur when attempting 
// to set ownership to a stone with no meta id
func TestAddOwnershipWithUnsetMetaID(t *testing.T) {
	var ownership = map[string]interface{}{}
	sh := Empty()
	err := sh.AddOwnership(ownership, util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "meta.id is not set")
}

// TestAddOwnership tests that the `ownership` block is assigned and signed successfully
func TestAddOwnership(t *testing.T) {
	var ownership = map[string]interface{}{
		"ref_id": "xxx",
		"type": "sole",
   		"sole": map[string]interface{}{
   			"address_id": "abcde",
   		},
   		"status": "transferred",
	}
	sh := Empty()
	sh.Meta["id"] = "xxx"
	err := sh.AddOwnership(ownership, util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, sh.Ownership["type"], ownership["type"])
	assert.NotNil(t, sh.Signatures["ownership"])
}


// TestAddOwnership tests that the `attributes` block is assigned and signed successfully
func TestAddAttributes(t *testing.T) {
	var attrs = map[string]interface{}{
		"ref_id": "xxx",
		"data": "abc",
	}
	sh := Empty()
	sh.Meta["id"] = "xxx"
	err := sh.AddAttributes(attrs, util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, sh.Attributes["some_data"], attrs["some_data"])
	assert.NotNil(t, sh.Signatures["attributes"])
}

// TestAddEmbed tests that the `embeds` block is assigned and signed successfully
func TestAddEmbed(t *testing.T) {

	sh := NewValidStone()

	embeds := map[string]interface{}{
		"ref_id": sh.Meta["id"],
		"data": []interface{}{
			map[string]interface{}{
				"meta": map[string]interface{}{
					"id": util.NewID(),
					"type": "coupon",
					"created_at": time.Now().Unix(),
				},
			},
		},
	}

	err := sh.AddEmbed(embeds, util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, sh.Embeds["ref_id"], embeds["ref_id"])
	assert.NotNil(t, sh.Signatures["embeds"])
}

// TestHasEmbedsTrue tests that a stone has it's embeds
// block set
func TestHasEmbedsTrue(t *testing.T) {

	sh := NewValidStone()

	embeds := map[string]interface{}{
		"ref_id": sh.Meta["id"],
		"data": []interface{}{
			map[string]interface{}{
				"meta": map[string]interface{}{
					"id": util.NewID(),
					"type": "coupon",
					"created_at": time.Now().Unix(),
				},
			},
		},
	}

	assert.Equal(t, sh.HasEmbeds(), false)
	err := sh.AddEmbed(embeds, util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, sh.HasEmbeds(), true)
}

// TestHashSignature tests that an attribute does not or has a signature
func TestHashSignature(t *testing.T) {
	var attrs = map[string]interface{}{
		"ref_id": "xxx",
		"data": "abc",
	}
	sh := Empty()
	sh.Meta["id"] = "xxx"
	assert.Equal(t, sh.HasSignature("attributes"), false)
	err := sh.AddAttributes(attrs, util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, sh.HasSignature("attributes"), true)
}

// TestCallVerifyWithUnknownBlockName tests that an error will occur when verifying an unknown block
func TestCallVerifyWithUnknownBlockName(t *testing.T) {
	stone, err := LoadJSON(util.ReadFromFixtures("tests/fixtures/stone_2.json"));
	assert.Nil(t, err)
	err = stone.Verify("some_block", util.ReadFromFixtures("tests/fixtures/rsa_pub_1.txt"))
	assert.NotNil(t, err)
	expectedMsg := "block unknown"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestCallVerifyWithInvalidPublicKey tests that an error will occur when verifying with an invalid public key
func TestCallVerifyWithInvalidPublicKey(t *testing.T) {
	stone, err := LoadJSON(util.ReadFromFixtures("tests/fixtures/stone_2.json"));
	assert.Nil(t, err)
	stone.Sign("meta", util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	err = stone.Verify("attributes", util.ReadFromFixtures("tests/fixtures/rsa_invalid_1.txt"))
	assert.NotNil(t, err)
	expectedMsg := `Public Key Error: unsupported key type "KEY"`
	assert.Equal(t, expectedMsg, err.Error())
}


// TestCallVerifyOnBlockWithNoSignature tests that an error will occur when verifying a block with no signature
// in the signatures block
func TestCallVerifyOnBlockWithNoSignature(t *testing.T) {
	sh := Empty()
	err := sh.Verify("attributes", util.ReadFromFixtures("tests/fixtures/rsa_pub_1.txt"))
	assert.NotNil(t, err)
	expectedMsg := "`attributes` block has no signature"
	assert.Equal(t, expectedMsg, err.Error())
}


// TestCallVerifyWhenBlockSignatureIsMalformed tests that an error will occur when verifying a block that has
// an invalid JWS signature
func TestCallVerifyWhenBlockSignatureIsMalformed(t *testing.T) {
	sh := Empty()
	sh.Signatures["attributes"] = "abcdefaa9*"
	err := sh.Verify("attributes", util.ReadFromFixtures("tests/fixtures/rsa_pub_1.txt"))
	assert.NotNil(t, err)
	expectedMsg := "`attributes` block signature could not be verified"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestCallVerifyWhenBlockSignatureInvalid tests that an error will occur when verifying a block 
// that has an invalid signature
func TestCallVerifyWhenBlockSignatureInvalid(t *testing.T) {
	sh := Empty()
	tamperedSig := "enWGZSZIifX0.eyJjVycmVuY3kifQ.pEBlIL38m8Uy8-Ko"
	sh.Signatures["attributes"] = tamperedSig
	err := sh.Verify("attributes", util.ReadFromFixtures("tests/fixtures/rsa_pub_1.txt"))
	assert.NotNil(t, err)
	expectedMsg := "`attributes` block signature could not be verified"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestVerifyMeta tests that a meta block signed with a private key is 
// successfully verified using the corresponding public key
func TestVerifyMeta(t *testing.T) {
	stone, err := LoadJSON(util.ReadFromFixtures("tests/fixtures/stone_2.json"));
	assert.Nil(t, err)
	stone.Sign("meta", util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	err = stone.Verify("meta", util.ReadFromFixtures("tests/fixtures/rsa_pub_1.txt"))
	assert.Nil(t, err)
}

// TestVerifyOwnership tests that an ownership block signed with a private key is 
// successfully verified using the corresponding public key
func TestVerifyOwnership(t *testing.T) {
	stone, err := LoadJSON(util.ReadFromFixtures("tests/fixtures/stone_2.json"));
	assert.Nil(t, err)
	stone.Sign("ownership", util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	err = stone.Verify("ownership", util.ReadFromFixtures("tests/fixtures/rsa_pub_1.txt"))
	assert.Nil(t, err)
}

// TestVerifyAttributes tests that an `attributes` block signed with a private key is 
// successfully verified using the corresponding public key
func TestVerifyAttributes(t *testing.T) {
	stone, err := LoadJSON(util.ReadFromFixtures("tests/fixtures/stone_2.json"));
	assert.Nil(t, err)
	stone.Sign("attributes", util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	err = stone.Verify("attributes", util.ReadFromFixtures("tests/fixtures/rsa_pub_1.txt"))
	assert.Nil(t, err)
}

// TestVerifyEmbeds tests that an `embeds` block signed with a private key is 
// successfully verified using the corresponding public key
func TestVerifyEmbeds(t *testing.T) {
	stone, err := LoadJSON(util.ReadFromFixtures("tests/fixtures/stone_4.json"));
	assert.Nil(t, err)
	stone.Sign("embeds", util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	err = stone.Verify("embeds", util.ReadFromFixtures("tests/fixtures/rsa_pub_1.txt"))
	assert.Nil(t, err)
}

// TestCloneStone
func TestCloneStone(t *testing.T) {
	stone, err := LoadJSON(util.ReadFromFixtures("tests/fixtures/stone_2.json"));
	assert.Nil(t, err)
	clone := stone.Clone()
	assert.Exactly(t, stone, clone) 
	stone.Signatures["meta"] = "blah_blah"
	assert.NotEmpty(t, stone.Signatures["meta"], clone.Signatures["meta"])
}

// TestHasOwnershipFalse tests that a stone does not have any ownership information
func TestHasOwnershipFalse(t *testing.T) {
	stone, err := LoadJSON(util.ReadFromFixtures("tests/fixtures/stone_1.json"));
	assert.Nil(t, err)
	assert.Equal(t, stone.HasOwnership(), false)
}

// TestHasOwnershipTrue tests that a stone has ownership information
func TestHasOwnershipTrue(t *testing.T) {
	stone, err := LoadJSON(util.ReadFromFixtures("tests/fixtures/stone_1.json"));
	assert.Nil(t, err)
	var ownership = map[string]interface{}{
		"ref_id": "4417781906fb0a89c295959b0df01782dbc4dc9f",
		"type": "sole",
   		"sole": map[string]interface{}{
			"address_id": "abcde",
   		},
	}
	err = stone.AddOwnership(ownership, util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, stone.HasOwnership(), true)
}

// TestHasAttributesReturnsTrue tests that a stone has attributes information
func TestHasAttributesReturnsTrue(t *testing.T) {
	stone, err := LoadJSON(util.ReadFromFixtures("tests/fixtures/stone_1.json"))
	assert.Nil(t, err)
	var attrs = map[string]interface{}{
		"ref_id": stone.Meta["id"],
		"data": "some_value",
	}
	err = stone.AddAttributes(attrs, util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, stone.HasAttributes(), true)
}

// TestHasAttributesReturnsFalse tests that a stone does not have attributes information
func TestHasAttributesReturnsFalse(t *testing.T) {
	stone, err := LoadJSON(util.ReadFromFixtures("tests/fixtures/stone_1.json"))
	assert.Nil(t, err)
	assert.Equal(t, stone.HasAttributes(), false)
}

// TestEncodeSuccessfully tests that a stone was encoded successfully
func TestEncodeSuccessfully(t *testing.T) {
	var meta = map[string]interface{}{
		"id": util.NewID(),
		"type": "currency",
		"created_at": time.Now().Unix(),
	}
	sh, err := Create(meta, util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	assert.Nil(t, err)
	enc, _ := util.MapToJSON(sh.Signatures)
	expectedEncodeVal := crypto.ToBase64([]byte(enc))
	assert.Equal(t, sh.Encode(), expectedEncodeVal)
}

// TestTokenToBlockSuccessfully tests that a JWS token is successfully decoded to a block
func TestTokenToBlockSuccessfully(t *testing.T) {
	var meta = map[string]interface{}{
		"id": util.NewID(),
		"type": "currency",
		"created_at": time.Now().Unix(),
	}
	sh, err := Create(meta, util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	assert.Nil(t, err)
	block, err := TokenToBlock(sh.Signatures["meta"].(string), "meta")
	assert.Nil(t, err)
	assert.Equal(t, meta["id"], block["id"])
}

// TestTokenToBlockWithInvalidToken test that an error occurs if token is invalid
func TestTokenToBlockWithInvalidToken(t *testing.T) {
	_, err := TokenToBlock("abcde", "meta")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "parameter is not a valid token")
}

// TestTokenToBlockWithInvalidPayload tests that an error occurs if payload part of the
// token is invalid
func TestTokenToBlockWithInvalidPayload(t *testing.T) {
	_, err := TokenToBlock("abcde.invalid_**payload.xyz", "meta")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "invalid meta token")
}

// TestTokenToBlockWithMalformedJSONInPayload test that an error occurs if payload data is malformed
func TestTokenToBlockWithMalformedJSONInPayload(t *testing.T) {
	malformedJSONPayload := crypto.ToBase64([]byte(`{ "field": "value" `))
	_, err := TokenToBlock("abcde."+malformedJSONPayload+".xyz", "meta")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "malformed meta block")
}

// TestDecodeWithUnSignedBlocks test that an empty block is derived after decoding 
// an encoded stone that had not signed it's blocks prior to encoding.
func TestDecodeWithUnSignedBlocks(t *testing.T) {
	sh, err := Load(util.ReadFromFixtures("tests/fixtures/stone_5.json"))
	assert.Nil(t, err)
	encStone := sh.Encode()
	decStone, err := Decode(encStone)
	assert.Nil(t, err)
	assert.Equal(t, len(decStone.Meta), 0)
	assert.Equal(t, len(decStone.Ownership), 0)
	assert.Equal(t, len(decStone.Attributes), 0)
	assert.Equal(t, len(decStone.Embeds), 0)
	assert.Equal(t, len(decStone.Signatures), 0)
}

// TestDecodeWithSignedBlock tests that a encoded stone blocks will be decoded
// correctly as long as blocks where signed before the encoding process was run.
func TestDecodeWithSignedBlock(t *testing.T) {
	sh, err := Load(util.ReadFromFixtures("tests/fixtures/stone_5.json"))
	assert.Nil(t, err)
	sh.Sign("meta", util.ReadFromFixtures("tests/fixtures/rsa_priv_1.txt"))
	encStone := sh.Encode()
	decStone, err := Decode(encStone)
	assert.Nil(t, err)
	assert.NotEqual(t, len(decStone.Meta), 0)
	assert.Exactly(t, sh.Meta, decStone.Meta)
}