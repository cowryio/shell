package stone

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/ellcrys/crypto"
)

func NewValidStone() *Stone {
	var meta = map[string]interface{}{
		"id": NewID(),
		"type": "some_stone",
		"created_at": time.Now().Unix(),
	}
	sh, _ := Create(meta, ReadFromFixtures("rsa_priv_1.txt"))
	return sh
} 

// TestCreateAStone create a valid, error free stone
func TestCreateAStone(t *testing.T) {
	stoneID := NewID()
	var meta = map[string]interface{}{
		"id": stoneID,
		"type": "currency",
		"created_at": time.Now().Unix(),
	}
	sh, err := Create(meta, ReadFromFixtures("rsa_priv_1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, sh.Meta["id"], stoneID)
	assert.NotEmpty(t, sh.Signatures["meta"])
}

// TestMustProvideMetaWithContent test that a map describing the `meta` block is required
func TestMustProvideMetaWithContent(t *testing.T) {
	_, err := Create(make(map[string]interface{}), ReadFromFixtures("rsa_priv_1.txt"))
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "`meta` block is missing `id` property")
}

// TestInvalidPrivateKey tests that an invalid private key returns an error
func TestInvalidPrivateKey(t *testing.T) {
	var issuerPrivateKey = ReadFromFixtures("rsa_invalid_1.txt")
	stoneID := NewID()
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
	txt := ReadFromFixtures("stone_1.json")
	stone, err := LoadJSON(txt)
	assert.Nil(t, err);
	assert.IsType(t, &Stone{}, stone)
}

// TestCorrectlySignMeta tests that a stone is correctly signed
func TestCorrectlySignMeta(t *testing.T) {
	txt := ReadFromFixtures("stone_1.json")
	stone, err := LoadJSON(txt)
	assert.Nil(t, err);
	expectedSignature := "eyJhbGciOiJSUzI1NiIsImp3ayI6eyJrdHkiOiJSU0EiLCJuIjoicTZHWW5qZ0tQYkxYSC1rZW5sbjZPZFZRcnl2SEMzVFV1ZS01dnh5QlRwaEhkUWc0djd1Mm9CczZYb1RRSVI2YS1UVlkwR2VFM3ZpakVaX1VwNlZDdG9YUEhWRk51VDBLSmJEaE1IajFVTmZJUnpTdUdOaWJ6bVAzX0NnanRvWWEwdXJyai1ubm5hWjBuYnBVdFRseDB5LW1jVUpnWGZSZDk0QzAtZ1JFUjBNIiwiZSI6IkFRQUIifX0.eyJjcmVhdGVkX2F0IjoxNDUzOTc1NTc1LCJpZCI6IjQ0MTc3ODE5MDZmYjBhODljMjk1OTU5YjBkZjAxNzgyZGJjNGRjOWYiLCJ0eXBlIjoiY3VycmVuY3kifQ.pEBlRBlIkmrMNJkBlvUWo5FK8N6-G83hirDNQLmYo6ojSkX0cXqak_mdHo7zUyLV0CxAvPuxb9fiYbz4S2tllIMpHm_RHQDDOXkl1ykiUrbcotrlfQmiOqvDzp91IL38m8Uy8-MBg-JB7K9nacCCLEph-BLn83AyyQeSVTQZGKo"
	signature, err := stone.Sign("meta", ReadFromFixtures("rsa_priv_1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, expectedSignature, signature)
	assert.Equal(t, expectedSignature, stone.Signatures["meta"])
}

// TestCannotSignUnknownBlock tests that an error will occur when attempting to sign an unknown block
func TestCannotSignUnknownBlock(t *testing.T) {
	txt := ReadFromFixtures("stone_1.json")
	stone, err := LoadJSON(txt)
	assert.Nil(t, err)
	_, err = stone.Sign("unknown_block", ReadFromFixtures("rsa_priv_1.txt"))
	assert.NotNil(t, err)
	expectedMsg := "block unknown"
	assert.Equal(t, err.Error(), expectedMsg)
}

// TestCannotSignEmptyBlock tests that an error will occur when attempting to sign an empty block
func TestCannotSignEmptyBlock(t *testing.T) {
	txt := ReadFromFixtures("stone_1.json")
	stone, err := LoadJSON(txt)
	assert.Nil(t, err)
	_, err = stone.Sign("ownership", ReadFromFixtures("rsa_priv_1.txt"))
	assert.NotNil(t, err)
	expectedMsg := "failed to sign empty block"
	assert.Equal(t, err.Error(), expectedMsg)
}

// TestAddMeta tests that a `meta` block can be assigned and signed successful
func TestAddMeta(t *testing.T) {
	stoneID := NewID()
	var meta = map[string]interface{}{
		"id": stoneID,
		"type": "currency",
		"created_at": time.Now().Unix(),
	}
	sh := Empty()
	err := sh.AddMeta(meta, ReadFromFixtures("rsa_priv_1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, sh.Meta["id"], meta["id"])
	assert.NotNil(t, sh.Signatures["meta"])
}

// TestAddOwnershipWithUnsetMetaID tests that an error will occur when attempting 
// to set ownership to a stone with no meta id
func TestAddOwnershipWithUnsetMetaID(t *testing.T) {
	var ownership = map[string]interface{}{}
	sh := Empty()
	err := sh.AddOwnership(ownership, ReadFromFixtures("rsa_priv_1.txt"))
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
	err := sh.AddOwnership(ownership, ReadFromFixtures("rsa_priv_1.txt"))
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
	err := sh.AddAttributes(attrs, ReadFromFixtures("rsa_priv_1.txt"))
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
					"id": NewID(),
					"type": "coupon",
					"created_at": time.Now().Unix(),
				},
			},
		},
	}

	err := sh.AddEmbed(embeds, ReadFromFixtures("rsa_priv_1.txt"))
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
					"id": NewID(),
					"type": "coupon",
					"created_at": time.Now().Unix(),
				},
			},
		},
	}

	assert.Equal(t, sh.HasEmbeds(), false)
	err := sh.AddEmbed(embeds, ReadFromFixtures("rsa_priv_1.txt"))
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
	err := sh.AddAttributes(attrs, ReadFromFixtures("rsa_priv_1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, sh.HasSignature("attributes"), true)
}

// TestCallVerifyWithUnknownBlockName tests that an error will occur when verifying an unknown block
// func TestCallVerifyWithUnknownBlockName(t *testing.T) {
// 	var attrs = map[string]interface{}{
// 		"some_data": "some_value",
// 	}
// 	sh := Empty()
// 	err := sh.AddAttributes(attrs, ReadFromFixtures("rsa_priv_1.txt"))
// 	assert.Nil(t, err)
// 	err = sh.Verify("some_block", ReadFromFixtures("rsa_pub_1.txt"))
// 	assert.NotNil(t, err)
// 	expectedMsg := "block unknown"
// 	assert.Equal(t, expectedMsg, err.Error())
// }

// TestCallVerifyWithInvalidPublicKey tests that an error will occur when verifying with an invalid public key
// func TestCallVerifyWithInvalidPublicKey(t *testing.T) {
// 	var attrs = map[string]interface{}{
// 		"some_data": "some_value",
// 	}
// 	sh := Empty()
// 	err := sh.AddAttributes(attrs, ReadFromFixtures("rsa_priv_1.txt"))
// 	assert.Nil(t, err)
// 	err = sh.Verify("attributes", ReadFromFixtures("rsa_invalid_1.txt"))
// 	assert.NotNil(t, err)
// 	expectedMsg := `Public Key Error: unsupported key type "KEY"`
// 	assert.Equal(t, expectedMsg, err.Error())
// }


// TestCallVerifyOnBlockWithNoSignature tests that an error will occur when verifying a block with no signature
// in the signatures block
func TestCallVerifyOnBlockWithNoSignature(t *testing.T) {
	sh := Empty()
	err := sh.Verify("attributes", ReadFromFixtures("rsa_pub_1.txt"))
	assert.NotNil(t, err)
	expectedMsg := "block `attributes` has no signature"
	assert.Equal(t, expectedMsg, err.Error())
}


// TestCallVerifyWhenBlockSignatureIsMalformed tests that an error will occur when verifying a block that has
// an invalid JWS signature
func TestCallVerifyWhenBlockSignatureIsMalformed(t *testing.T) {
	sh := Empty()
	sh.Signatures["attributes"] = "abcdefaa9*"
	err := sh.Verify("attributes", ReadFromFixtures("rsa_pub_1.txt"))
	assert.NotNil(t, err)
	expectedMsg := "invalid signature"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestCallVerifyWhenBlockSignatureInvalid tests that an error will occur when verifying a block 
// that has an invalid signature
func TestCallVerifyWhenBlockSignatureInvalid(t *testing.T) {
	sh := Empty()
	tamperedSig := "enWGZSZIifX0.eyJjVycmVuY3kifQ.pEBlIL38m8Uy8-Ko"
	sh.Signatures["attributes"] = tamperedSig
	err := sh.Verify("attributes", ReadFromFixtures("rsa_pub_1.txt"))
	assert.NotNil(t, err)
	expectedMsg := "invalid signature"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestVerifyMeta tests that a meta block signed with a private key is 
// successfully verified using the corresponding public key
func TestVerifyMeta(t *testing.T) {
	stone, err := LoadJSON(ReadFromFixtures("stone_2.json"));
	assert.Nil(t, err)
	stone.Sign("meta", ReadFromFixtures("rsa_priv_1.txt"))
	err = stone.Verify("meta", ReadFromFixtures("rsa_pub_1.txt"))
	assert.Nil(t, err)
}

// TestVerifyOwnership tests that an ownership block signed with a private key is 
// successfully verified using the corresponding public key
func TestVerifyOwnership(t *testing.T) {
	stone, err := LoadJSON(ReadFromFixtures("stone_2.json"));
	assert.Nil(t, err)
	stone.Sign("ownership", ReadFromFixtures("rsa_priv_1.txt"))
	err = stone.Verify("ownership", ReadFromFixtures("rsa_pub_1.txt"))
	assert.Nil(t, err)
}

// TestVerifyAttributes tests that an `attributes` block signed with a private key is 
// successfully verified using the corresponding public key
func TestVerifyAttributes(t *testing.T) {
	stone, err := LoadJSON(ReadFromFixtures("stone_2.json"));
	assert.Nil(t, err)
	stone.Sign("attributes", ReadFromFixtures("rsa_priv_1.txt"))
	err = stone.Verify("attributes", ReadFromFixtures("rsa_pub_1.txt"))
	assert.Nil(t, err)
}

// TestVerifyEmbeds tests that an `embeds` block signed with a private key is 
// successfully verified using the corresponding public key
func TestVerifyEmbeds(t *testing.T) {
	stone, err := LoadJSON(ReadFromFixtures("stone_4.json"));
	assert.Nil(t, err)
	stone.Sign("embeds", ReadFromFixtures("rsa_priv_1.txt"))
	err = stone.Verify("embeds", ReadFromFixtures("rsa_pub_1.txt"))
	assert.Nil(t, err)
}

// TestCloneStone
func TestCloneStone(t *testing.T) {
	stone, err := LoadJSON(ReadFromFixtures("stone_2.json"));
	assert.Nil(t, err)
	clone := stone.Clone()
	assert.Exactly(t, stone, clone) 
	stone.Signatures["meta"] = "blah_blah"
	assert.NotEmpty(t, stone.Signatures["meta"], clone.Signatures["meta"])
}

// TestHasOwnershipFalse tests that a stone does not have any ownership information
func TestHasOwnershipFalse(t *testing.T) {
	stone, err := LoadJSON(ReadFromFixtures("stone_1.json"));
	assert.Nil(t, err)
	assert.Equal(t, stone.HasOwnership(), false)
}

// TestHasOwnershipTrue tests that a stone has ownership information
func TestHasOwnershipTrue(t *testing.T) {
	stone, err := LoadJSON(ReadFromFixtures("stone_1.json"));
	assert.Nil(t, err)
	var ownership = map[string]interface{}{
		"ref_id": "4417781906fb0a89c295959b0df01782dbc4dc9f",
		"type": "sole",
   		"sole": map[string]interface{}{
			"address_id": "abcde",
   		},
	}
	err = stone.AddOwnership(ownership, ReadFromFixtures("rsa_priv_1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, stone.HasOwnership(), true)
}

// TestHasAttributesReturnsTrue tests that a stone has attributes information
func TestHasAttributesReturnsTrue(t *testing.T) {
	stone, err := LoadJSON(ReadFromFixtures("stone_1.json"))
	assert.Nil(t, err)
	var attrs = map[string]interface{}{
		"ref_id": stone.Meta["id"],
		"data": "some_value",
	}
	err = stone.AddAttributes(attrs, ReadFromFixtures("rsa_priv_1.txt"))
	assert.Nil(t, err)
	assert.Equal(t, stone.HasAttributes(), true)
}

// TestHasAttributesReturnsFalse tests that a stone does not have attributes information
func TestHasAttributesReturnsFalse(t *testing.T) {
	stone, err := LoadJSON(ReadFromFixtures("stone_1.json"))
	assert.Nil(t, err)
	assert.Equal(t, stone.HasAttributes(), false)
}

func TestEncodeSuccessfully(t *testing.T) {
	stoneID := NewID()
	var meta = map[string]interface{}{
		"id": stoneID,
		"type": "currency",
		"created_at": time.Now().Unix(),
	}
	sh, err := Create(meta, ReadFromFixtures("rsa_priv_1.txt"))
	assert.Nil(t, err)
	enc, _ := MapToJSON(sh.Signatures)
	expectedEncodeVal := crypto.ToBase64([]byte(enc))
	assert.Equal(t, sh.Encode(), expectedEncodeVal)
}