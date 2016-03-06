package stone

import "testing"
import "github.com/stretchr/testify/assert"

// TestInvalidJSON tests that an invalid json string returns an error
func TestInvalidJSON(t *testing.T) {
	var str = `{ "meta : "" }`
	err := Validate(str)
	assert.NotNil(t, err)
	expectedMsg := `unable to parse json string`
	assert.Equal(t, expectedMsg, err.Error())
}

// TestMetaBlockHasUnexpectedProperty tests that an error occurs when the meta block
// contains an unexpected property
func TestMetaBlockHasUnexpectedProperty(t *testing.T) {
	d := map[string]interface{}{
		"some_property": "abcde",
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`some_property` property is unexpected in `meta` block"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestMetaMustHaveStoneIDProperty tests that a meta block data must have a `stone_id` property
func TestMetaMustHaveStoneIDProperty(t *testing.T) {
	var d = make(map[string]interface{})
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta` block is missing `stone_id` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestMetaMustHaveStoneTypeProperty tests that stone_type property is set
func TestMetaMustHaveStoneTypeProperty(t *testing.T) {
	d := map[string]interface{}{
		"stone_id": 1234,
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta` block is missing `stone_type` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestMetaMustHaveCreatedAtProperty tests that created_at property is set
func TestMetaMustHaveCreatedAtProperty(t *testing.T) {
	d := map[string]interface{}{
		"stone_id":   1234,
		"stone_type": "coupon",
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta` block is missing `created_at` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestStoneIDNotString tests that stone_id value type must be string
func TestStoneIDMustString(t *testing.T) {
	d := map[string]interface{}{
		"stone_id":   1234,
		"stone_type": "coupon",
		"created_at": 1000,
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta.stone_id` value type is invalid. Expects string value"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestStoneIDLengthInvalid tests that a stone_id must have 40 characters
func TestStoneIDLengthInvalid(t *testing.T) {
	d := map[string]interface{}{
		"stone_id":   "abcd",
		"stone_type": "coupon",
		"created_at": 1000,
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta.stone_id` must have 40 characters. Preferrable a UUIDv4 SHA1 hashed string"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestStoneTypeMustBeString tests that stone_type value type must be string
func TestStoneTypeMustBeString(t *testing.T) {
	d := map[string]interface{}{
		"stone_id":   Sha1("abcd"),
		"stone_type": 123,
		"created_at": 1000,
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta.stone_type` value type is invalid. Expects string value"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestCreatedAtMustBeNumber tests that created_at value type must be a number
func TestCreatedAtMustBeNumber(t *testing.T) {
	d := map[string]interface{}{
		"stone_id":   Sha1("abcd"),
		"stone_type": "coupon",
		"created_at": "111",
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`created_at` value type is invalid. Expects a number"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestCreatedAtBeforeStartTime test that a created_at time before the start/launch time is invalid
func TestCreatedAtBeforeStartTime(t *testing.T) {
	d := map[string]interface{}{
		"stone_id":   Sha1("abcd"),
		"stone_type": "coupon",
		"created_at": 100000,
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`created_at` value is too far in the past. Expects unix time on or after 2016-01-28T11:06:15+01:00"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestSignaturesBlockHasUnexpectedProperty tests that an error will occur when 
// `signatures` block contains unexpected property
func TestSignaturesBlockHasUnexpectedProperty(t *testing.T) {
	d := map[string]interface{}{
		"some_property": "abcde",
	}
	err := ValidateSignaturesBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`some_property` property is unexpected in `signatures` block"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestSignaturesBlockMustHaveMetaProperty test that signatures block must have `meta` property
func TestSignaturesBlockMustHaveMetaProperty(t *testing.T) {
	d := map[string]interface{}{}
	err := ValidateSignaturesBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "missing `signatures.meta` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestMetaSignatureTypeMustBeString test that the signatures.meta property value type is string
func TestMetaSignatureTypeMustBeString(t *testing.T) {
	d := map[string]interface{}{
		"meta": 1234,
	}
	err := ValidateSignaturesBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`signatures.meta` value type is invalid. Expects string value"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestOwnershipSignatureTypeMustBeString tests that when `ownership` property is set, it's value type must be string
func TestOwnershipSignatureTypeMustBeString(t *testing.T) {
	d := map[string]interface{}{
		"meta":      "abcde",
		"ownership": 100,
	}
	err := ValidateSignaturesBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`signatures.ownership` value type is invalid. Expects string value"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestAttributesSignatureTypeMustBeString tests that when `attributes` 
// property is set, it's value type must be string
func TestAttributesSignatureTypeMustBeString(t *testing.T) {
	d := map[string]interface{}{
		"meta":       "abcde",
		"ownership":  "abcde",
		"attributes": 100,
	}
	err := ValidateSignaturesBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`signatures.attributes` value type is invalid. Expects string value"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestEmbedsSignatureTypeMustBeString test that the signatures.embeds property value type is string
func TestEmbedsSignatureTypeMustBeString(t *testing.T) {
	d := map[string]interface{}{
		"meta": "abcde",
		"embeds": 1234,
	}
	err := ValidateSignaturesBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`signatures.embeds` value type is invalid. Expects string value"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestMetaBlockHasUnexpectedProperty tests that an error occurs when the ownership block
// contains an unexpected property
func TestOwnershipBlockHasUnexpectedProperty(t *testing.T) {
	d := map[string]interface{}{
		"some_property": "abcde",
	}
	err := ValidateOwnershipBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`some_property` property is unexpected in `ownership` block"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestOwnershipTypePropertyMissing tests that an error occurs when 
// `ownership` block is set with missing `type` property
func TestOwnershipTypePropertyMissing(t *testing.T) {
	d := map[string]interface{}{}
	err := ValidateOwnershipBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`ownership` block is missing `type` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestInvalidOwnershipTypeValue tests that an error will occur when
// `ownership.type` is set to an unacceptable value
func TestInvalidOwnershipTypeValue(t *testing.T) {
	d := map[string]interface{}{
		"type": "some_value",
	}
	err := ValidateOwnershipBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`ownership.type` property has unexpected value"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestMissingSoleProperty tests that an error will occur when `ownership.type` is 
// `sole` and `ownership.sole` property is missing
func TestMissingSoleProperty(t *testing.T) {
	d := map[string]interface{}{
		"type": "sole",
	}
	err := ValidateOwnershipBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`ownership` block is missing `sole` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestInvalidSolePropertyType tests that an error will occur when `ownership.sole` value
// type is not a map of interface{} value
func TestInvalidSolePropertyType(t *testing.T) {
	d := map[string]interface{}{
		"type": "sole",
		"sole": "abcde",
	}
	err := ValidateOwnershipBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`ownership.sole` value type is invalid. Expects a JSON object"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestSolePropertyMissingAddressIDProperty tests that an error will occur when
// `ownership.sole` is missing `address_id` property
func TestSolePropertyMissingAddressIDProperty(t *testing.T) {
	d := map[string]interface{}{
		"type": "sole",
		"sole": map[string]interface{}{},
	}
	err := ValidateOwnershipBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`ownership.sole` property is missing `address_id` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestInvalidSolePropertyAddressID tests that an error will occur when `ownership.sole.address_id` 
// value type is not string
func TestInvalidSolePropertyAddressID(t *testing.T) {
	d := map[string]interface{}{
		"type": "sole",
		"sole": map[string]interface{}{
			"address_id": 123,	
		},
	}
	err := ValidateOwnershipBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`ownership.sole.address_id` value type is invalid. Expects string value"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestInvalidSoleStatusPropertyValueType tests that an error will occur when 
// `ownership.status` is set with an invalid value type
func TestInvalidSoleStatusPropertyValueType(t *testing.T) {
	d := map[string]interface{}{
		"type": "sole",
		"sole": map[string]interface{}{
			"address_id": "abcde",	
		},
		"status": 123,
	}
	err := ValidateOwnershipBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`ownership.status` value type is invalid. Expects string value"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestUnexpectedOwnershipStatusValue tests that an error will occur when 
// `ownership.status` is set with an unexpected value
func TestUnexpectedOwnershipStatusValue(t *testing.T) {
	d := map[string]interface{}{
		"type": "sole",
		"sole": map[string]interface{}{
			"address_id": "abcde",	
		},
		"status": "unexpected_value",
	}
	err := ValidateOwnershipBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`ownership.status` property has unexpected value"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestMustHaveMetaBlock tests that a json string must have a `meta` property
func TestMustHaveMetaBlock(t *testing.T) {
	var str = `{ "signatures": {}}`
	err := Validate(str)
	assert.NotNil(t, err)
	expectedMsg := "missing `meta` block"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestInvalidMetaValueType tests that the `meta` block/property value type must be a JSON object
func TestInvalidMetaValueType(t *testing.T) {
	var str = `{ "meta": "some stuff" }`
	err := Validate(str)
	assert.NotNil(t, err)
	expectedMsg := "`meta` block value type is invalid. Expects a JSON object"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestMustHaveSignatureBlock tests that the json string must have a `signatures` property
func TestMustHaveSignatureBlock(t *testing.T) {
	var str = `{ 
		"meta": { 
			"stone_id": "` + Sha1("stuff") + `", 
			"stone_type": "cur", 
			"created_at": 1453975575 
		} 
	}`
	err := Validate(str)
	assert.NotNil(t, err)
	expectedMsg := "missing `signatures` block"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestInvalidSignaturesValueType tests that the `signatures` block/property value type must be a JSON object
func TestInvalidSignaturesBlockValueType(t *testing.T) {
	var str = `{ 
		"signatures": "a_string", 
		"meta": { 
			"stone_id": "` + Sha1("stuff") + `", 
			"stone_type": "cur", 
			"created_at": 1453975575 
		} 
	}`
	err := Validate(str)
	assert.NotNil(t, err)
	expectedMsg := "`signature` block value type is invalid. Expects a JSON object"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestInvalidOwnershipBlockValueType tests that the value type of `ownership block must be a map
func TestInvalidOwnershipBlockValueType(t *testing.T) {
	var str = `{ 
		"signatures": { 
			"meta": "abcde"
		}, 
		"meta": { 
			"stone_id": "` + Sha1("stuff") + `", 
			"stone_type": "cur", 
			"created_at": 1453975575 
		},
		"ownership": "abcde"
	}`
	err := Validate(str)
	assert.NotNil(t, err)
	expectedMsg := "`ownership` block value type is invalid. Expects a JSON object"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestSignaturesBlockMustHaveOwnershipProperty tests that whenever `ownership` block is set,
// `signatures` block must have `ownership` property
func TestSignaturesBlockMustHaveOwnershipProperty(t *testing.T) {
	var str = `{ 
		"signatures": { 
			"meta": "abcde"
		}, 
		"meta": { 
			"stone_id": "` + Sha1("stuff") + `", 
			"stone_type": "cur", 
			"created_at": 1453975575 
		},
		"ownership": { "stuff": "stuff" }
	}`
	err := Validate(str)
	assert.NotNil(t, err)
	expectedMsg := "missing `ownership` property in `signatures` block"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestInvalidAttributesBlockValueType tests that the value type of `attributes` block must be a map
func TestInvalidAttributesBlockValueType(t *testing.T) {
	var str = `{ 
		"signatures": { 
			"meta": "abcde",
			"ownership": "abcde"
		}, 
		"meta": { 
			"stone_id": "` + Sha1("stuff") + `", 
			"stone_type": "cur", 
			"created_at": 1453975575 
		},
		"ownership": { 
			"type": "sole", 
		"sole": {
				"address_id": "1234"
			} 
		},
		"attributes": "abcde"
	}`
	err := Validate(str)
	assert.NotNil(t, err)
	expectedMsg := "`attributes` block value type is invalid. Expects a JSON object"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestSignaturesBlockMustHaveAttributesProperty tests that whenever `attributes` block is set,
// `signatures` block must have `attributes` property
func TestSignaturesBlockMustHaveAttributesProperty(t *testing.T) {
	var str = `{ 
		"signatures": { 
			"meta": "abcde",
			"ownership": "abcde"
		}, 
		"meta": { 
			"stone_id": "` + Sha1("stuff") + `", 
			"stone_type": "cur", 
			"created_at": 1453975575 
		},
		"ownership": { 
			"type": "sole", 
			"sole": {
				"address_id": "1234"
			}
		},
		"attributes": { "stuff": "stuff" }
	}`
	err := Validate(str)
	assert.NotNil(t, err)
	expectedMsg := "missing `attributes` property in `signatures` block"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestInvalidEmbedsValueType tests that an error will occur if `embeds` block value type is not a
// slice of objects
func TestInvalidEmbedsValueType(t *testing.T) {
	var str = `{ 
		"signatures": { 
			"meta": "abcde",
			"ownership": "abcde",
			"attributes": "abcde"
		}, 
		"meta": { 
			"stone_id": "` + Sha1("stuff") + `", 
			"stone_type": "cur", 
			"created_at": 1453975575 
		},
		"ownership": { 
			"type": "sole", 
			"sole": {
				"address_id": "1234"
			}
		},
		"attributes": { "stuff": "stuff" },
		"embeds": "abcdef"
	}`
	err := Validate(str)
	assert.NotNil(t, err)
	expectedMsg := "`embeds` block value type is invalid. Expects a list of only JSON objects"
	assert.Equal(t, expectedMsg, err.Error())
}


// TestNoErrorEmptyEmbedsBlock tests that no error will occur if `embeds` block 
// is set to an empty slice
func TestNoErrorEmptyEmbedsBlock(t *testing.T) {
	var str = `{ 
		"signatures": { 
			"meta": "abcde",
			"ownership": "abcde",
			"attributes": "abcde"
		}, 
		"meta": { 
			"stone_id": "` + Sha1("stuff") + `", 
			"stone_type": "cur", 
			"created_at": 1453975575 
		},
		"ownership": { 
			"type": "sole", 
			"sole": {
				"address_id": "1234"
			}
		},
		"attributes": { "stuff": "stuff" },
		"embeds": []
	}`
	err := Validate(str)
	assert.Nil(t, err)
}

// TestEmbedsBlockWithInvalidStone tests that an error will occur when the `embeds` block contains
// an invalid stone 
func TestEmbedsBlockWithInvalidStone(t *testing.T) {
	var str = `{ 
		"signatures": { 
			"meta": "abcde",
			"ownership": "abcde",
			"attributes": "abcde",
			"embeds": "abcde"
		}, 
		"meta": { 
			"stone_id": "` + Sha1("stuff") + `", 
			"stone_type": "cur", 
			"created_at": 1453975575 
		},
		"ownership": { 
			"type": "sole", 
			"sole": {
				"address_id": "1234"
			}
		},
		"attributes": { "stuff": "stuff" },
		"embeds": [{ }]
	}`
	err := Validate(str)
	assert.NotNil(t, err)
	expectedMsg := "unable to validate embed at index 0. Reason: missing `meta` block"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestIgnoreDeeperEmbedsLevel test that the validator will not validate nested embeds
// other that the ones in the stone been validated
func TestIgnoreDeeperEmbedsLevel(t *testing.T) {
	var str = `{ 
		"signatures": { 
			"meta": "abcde",
			"ownership": "abcde",
			"attributes": "abcde",
			"embeds": "abcde"
		}, 
		"meta": { 
			"stone_id": "` + Sha1("stuff") + `", 
			"stone_type": "cur", 
			"created_at": 1453975575 
		},
		"ownership": { 
			"type": "sole", 
			"sole": {
				"address_id": "1234"
			}
		},
		"attributes": { "stuff": "stuff" },
		"embeds": [{"signatures":{"meta":"abcde","ownership":"abcde"},"meta":{"created_at":1454443443,"stone_id":"4417781906fb0a89c295959b0df01782dbc4dc9f","stone_type":"currency"},"ownership":{"type":"sole","sole":{"address_id":"abcde"},"status":"transferred"},"embeds":[{ "invalid": "embed" }],"attributes":{}}]
	}`
	err := Validate(str)
	assert.Nil(t, err)
}
