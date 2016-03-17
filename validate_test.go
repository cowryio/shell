package stone

import "testing"
import "time"
import "github.com/stretchr/testify/assert"

// TestInvalidJSON tests that an invalid json string returns an error
func TestInvalidJSON(t *testing.T) {
	var str = `{ "meta : "" }`
	err := Validate(str)
	assert.NotNil(t, err)
	expectedMsg := `unable to parse json string`
	assert.Equal(t, expectedMsg, err.Error())
}

// TestValidateWithMap tests that a map representing a stone can be passed to Validate()
func TestValidateWithMap(t *testing.T) {
	s := map[string]interface{}{
		"meta": map[string]interface{}{
			"id": NewID(),
			"type": "coupon",
			"created_at": time.Now().Unix(),
		},
	}
	err := Validate(s)
	assert.Nil(t, err)
}

// TestValidateWithUnsupportedType tests that an unsupported type will not be allowed
func TestValidateWithUnsupportedType(t *testing.T) {
	err := Validate(123)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "unsupported parameter type")
}

// TestMustHaveMetaBlock tests that a json string must have a `meta` property
func TestMustHaveMetaBlock(t *testing.T) {
	var str = `{}`
	err := Validate(str)
	assert.NotNil(t, err)
	expectedMsg := "missing `meta` block"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestInvalidMetaValueType tests that the `meta` block value type must be a JSON object
func TestInvalidMetaValueType(t *testing.T) {
	var str = `{ "meta": "some stuff" }`
	err := Validate(str)
	assert.NotNil(t, err)
	expectedMsg := "`meta` block value type is invalid. Expects a JSON object"
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

// TestMetaMustHaveStoneIDProperty tests that a meta block data must have a `id` property
func TestMetaMustHaveStoneIDProperty(t *testing.T) {
	var d = make(map[string]interface{})
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta` block is missing `id` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestMetaMustHaveStoneTypeProperty tests that type property is set
func TestMetaMustHaveStoneTypeProperty(t *testing.T) {
	d := map[string]interface{}{
		"id": 1234,
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta` block is missing `type` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestMetaMustHaveCreatedAtProperty tests that created_at property is set
func TestMetaMustHaveCreatedAtProperty(t *testing.T) {
	d := map[string]interface{}{
		"id":   1234,
		"type": "coupon",
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta` block is missing `created_at` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestStoneIDNotString tests that id value type must be string
func TestStoneIDMustString(t *testing.T) {
	d := map[string]interface{}{
		"id":   1234,
		"type": "coupon",
		"created_at": 1000,
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta.id` value type is invalid. Expects a string"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestStoneIDLengthInvalid tests that a id must have 40 characters
func TestStoneIDLengthInvalid(t *testing.T) {
	d := map[string]interface{}{
		"id":   "abcd",
		"type": "coupon",
		"created_at": 1000,
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta.id` must have 40 characters. Preferrable a UUIDv4 SHA1 hashed string"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestStoneTypeMustBeString tests that type value type must be string
func TestStoneTypeMustBeString(t *testing.T) {
	d := map[string]interface{}{
		"id":   Sha1("abcd"),
		"type": 123,
		"created_at": 1000,
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta.type` value type is invalid. Expects a string"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestCreatedAtMustBeNumber tests that created_at value type must be a number
func TestCreatedAtMustBeNumber(t *testing.T) {
	d := map[string]interface{}{
		"id":   Sha1("abcd"),
		"type": "coupon",
		"created_at": "111",
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta.created_at` value type is invalid. Expects a number"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestCreatedAtBeforeStartTime test that a created_at time before the start/launch time is invalid
func TestCreatedAtBeforeStartTime(t *testing.T) {
	d := map[string]interface{}{
		"id":   Sha1("abcd"),
		"type": "coupon",
		"created_at": 100000,
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta.created_at` value is too far in the past. Expects unix time on or after 2016-01-28T11:06:15+01:00"
	assert.Equal(t, expectedMsg, err.Error())
}


// TestMetaBlockHasUnexpectedProperty tests that an error occurs when the ownership block
// contains an unexpected property
func TestOwnershipBlockHasUnexpectedProperty(t *testing.T) {
	d := map[string]interface{}{
		"some_property": "abcde",
	}
	err := ValidateOwnershipBlock(d, "xxx")
	assert.NotNil(t, err)
	expectedMsg := "`some_property` property is unexpected in `ownership` block"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestOwnershipRefIDPropertyMissing test that an error will occur when ownership.ref_id property is missing
func TestOwnershipRefIDPropertyMissing(t *testing.T) {
	d := map[string]interface{}{}
	err := ValidateOwnershipBlock(d, "xxx")
	assert.NotNil(t, err)
	expectedMsg := "`ownership` block is missing `ref_id` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestOwnershipRefIDWithInvalidValueType tests that an error will occur when ownership.ref_id value type is invalid
func TestOwnershipRefIDWithInvalidValueType(t *testing.T) {
	d := map[string]interface{}{
		"ref_id": 123,
	}
	err := ValidateOwnershipBlock(d, "xxx")
	assert.NotNil(t, err)
	expectedMsg := "`ownership.ref_id` value type is invalid. Expects string value"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestRefIDWithInvalidMetaID tests that an error will occur when ref_id property is not equal to 
// meta id parameter
func TestRefIDWithInvalidMetaID(t *testing.T) {
	d := map[string]interface{}{
		"ref_id": "xxx",
	}
	err := ValidateOwnershipBlock(d, "yyy")
	assert.NotNil(t, err)
	expectedMsg := "`ownership.ref_id` not equal to `meta.id`"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestOwnershipTypePropertyMissing tests that an error occurs when ownership.type property is missing
func TestOwnershipTypePropertyMissing(t *testing.T) {
	d := map[string]interface{}{
		"ref_id": "xxx",
	}
	err := ValidateOwnershipBlock(d, "xxx")
	assert.NotNil(t, err)
	expectedMsg := "`ownership` block is missing `type` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestInvalidOwnershipTypeValue tests that an error will occur when
// `ownership.type` is set to an unacceptable value
func TestInvalidOwnershipTypeValue(t *testing.T) {
	d := map[string]interface{}{
		"ref_id": "xxx",
		"type": "some_value",
	}
	err := ValidateOwnershipBlock(d, "xxx")
	assert.NotNil(t, err)
	expectedMsg := "`ownership.type` property has unexpected value"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestMissingSoleProperty tests that an error will occur when `ownership.type` is 
// `sole` and `ownership.sole` property is missing
func TestMissingSoleProperty(t *testing.T) {
	d := map[string]interface{}{
		"ref_id": "xxx",
		"type": "sole",
	}
	err := ValidateOwnershipBlock(d, "xxx")
	assert.NotNil(t, err)
	expectedMsg := "`ownership` block is missing `sole` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestInvalidSolePropertyType tests that an error will occur when `ownership.sole` value
// type is not a map of interface{} value
func TestInvalidSolePropertyType(t *testing.T) {
	d := map[string]interface{}{
		"ref_id": "xxx",
		"type": "sole",
		"sole": "abcde",
	}
	err := ValidateOwnershipBlock(d, "xxx")
	assert.NotNil(t, err)
	expectedMsg := "`ownership.sole` value type is invalid. Expects a JSON object"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestSolePropertyMissingAddressIDProperty tests that an error will occur when
// `ownership.sole` is missing `address_id` property
func TestSolePropertyMissingAddressIDProperty(t *testing.T) {
	d := map[string]interface{}{
		"ref_id": "xxx",
		"type": "sole",
		"sole": map[string]interface{}{},
	}
	err := ValidateOwnershipBlock(d, "xxx")
	assert.NotNil(t, err)
	expectedMsg := "`ownership.sole` property is missing `address_id` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestInvalidSolePropertyAddressID tests that an error will occur when `ownership.sole.address_id` 
// value type is not string
func TestInvalidSolePropertyAddressID(t *testing.T) {
	d := map[string]interface{}{
		"ref_id": "xxx",
		"type": "sole",
		"sole": map[string]interface{}{
			"address_id": 123,	
		},
	}
	err := ValidateOwnershipBlock(d, "xxx")
	assert.NotNil(t, err)
	expectedMsg := "`ownership.sole.address_id` value type is invalid. Expects a string"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestInvalidOwnershipStatusPropertyValueType tests that an error will occur when 
// `ownership.status` is set with an invalid value type
func TestInvalidOwnershipStatusPropertyValueType(t *testing.T) {
	d := map[string]interface{}{
		"ref_id": "xxx",
		"type": "sole",
		"sole": map[string]interface{}{
			"address_id": "abcde",	
		},
		"status": 123,
	}
	err := ValidateOwnershipBlock(d, "xxx")
	assert.NotNil(t, err)
	expectedMsg := "`ownership.status` value type is invalid. Expects a string"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestUnexpectedOwnershipStatusValue tests that an error will occur when 
// `ownership.status` is set with an unexpected value
func TestUnexpectedOwnershipStatusValue(t *testing.T) {
	d := map[string]interface{}{
		"ref_id": "xxx",
		"type": "sole",
		"sole": map[string]interface{}{
			"address_id": "abcde",	
		},
		"status": "unexpected_value",
	}
	err := ValidateOwnershipBlock(d, "xxx")
	assert.NotNil(t, err)
	expectedMsg := "`ownership.status` property has unexpected value"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestAddAttributesWithUnexpectedProp tests that an error will occur when an attribute 
// block contain unexpected property
func TestAttributesWithUnexpectedProp(t *testing.T) {
	d := map[string]interface{}{
		"unexpected_key": "some_value",
	}
	err := ValidateAttributesBlock(d, "")
	assert.NotNil(t, err)
	expectedMsg := "`unexpected_key` property is unexpected in `attributes` block"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestAttributesWithMissingRefID tests that an error will occur if
// attribute block is missing ref id
func TestAttributesWithMissingRefID(t *testing.T) {
	d := map[string]interface{}{}
	err := ValidateAttributesBlock(d, "")
	assert.NotNil(t, err)
	expectedMsg := "`attributes` block is missing `ref_id` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestAttributesWithWrongRefID tests that an error will occur if
// attribute.ref_id is not equal to meta id
func TestAttributesWithWrongRefID(t *testing.T) {
	d := map[string]interface{}{
		"ref_id": "xxx",
	}
	err := ValidateAttributesBlock(d, "")
	assert.NotNil(t, err)
	expectedMsg := "`attributes.ref_id` not equal to `meta.id`"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestAttributesWithMissingData tests that an error will occur if
// attribute.data property is missing
func TestAttributesWithMissingData(t *testing.T) {
	d := map[string]interface{}{
		"ref_id": "xxx",
	}
	err := ValidateAttributesBlock(d, "xxx")
	assert.NotNil(t, err)
	expectedMsg := "`attributes` block is missing `data` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestEmbedsWithUnexpectedProp tests that an error will occur if 
// embeds block contains unexpected property
func TestEmbedsWithUnexpectedProp(t *testing.T) {
	d := map[string]interface{}{
		"unexpected_key": "some_value",
	}
	err := ValidateEmbedsBlock(d, "xxx")
	assert.NotNil(t, err)
	expectedMsg := "`unexpected_key` property is unexpected in `embeds` block"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestEmbedsWithRefIDMissing tests that an error will occur if embeds
// block is missing ref_id property
func TestEmbedsWithRefIDMissing(t *testing.T) {
	d := map[string]interface{}{}
	err := ValidateEmbedsBlock(d, "xxx")
	assert.NotNil(t, err)
	expectedMsg := "`embeds` block is missing `ref_id` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestEmbedsWithWrongRefID tests that an error will occur if
// ref_id is not equal to meta id
func TestEmbedsWithWrongRefID(t *testing.T) {
	d := map[string]interface{}{
		"ref_id": "xxx",
	}
	err := ValidateEmbedsBlock(d, "yyy")
	assert.NotNil(t, err)
	expectedMsg := "`embeds.ref_id` not equal to `meta.id`"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestEmbedsWithDataMissing tests that an error will occur if data property
// is missing
func TestEmbedsWithDataMissing(t *testing.T) {
	d := map[string]interface{}{
		"ref_id": "xxx",
	}
	err := ValidateEmbedsBlock(d, "xxx")
	assert.NotNil(t, err)
	expectedMsg := "`embeds` block is missing `data` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestEmbedsWithInvalidDataValueType that an error will occur if data 
// property has invalid value type
func TestEmbedsWithInvalidDataValueType(t *testing.T) {
	d := map[string]interface{}{
		"ref_id": "xxx",
		"data": 123,
	}
	err := ValidateEmbedsBlock(d, "xxx")
	assert.NotNil(t, err)
	expectedMsg := "`embeds.data` value type is invalid. Expects a slice of JSON objects"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestIgnoreChildEmbedsOfEmbeds tests that child embeds are not validated
func TestIgnoreChildEmbedsOfEmbeds(t *testing.T) {
	d := map[string]interface{}{
		"ref_id": "xxx",
		"data": []interface{}{
			map[string]interface{}{
				"meta": map[string]interface{}{
					"id": NewID(),
					"type": "coupon",
					"created_at": time.Now().Unix(),
				},
				"embeds": map[string]interface{}{
					"ref_id": "xxx",
					"data": "*invalid_type*",
				},
			},
		},
	}
	err := ValidateEmbedsBlock(d, "xxx")
	assert.Nil(t, err)
	childEmbed := d["data"].([]interface{})[0].(map[string]interface{})
	assert.NotNil(t, childEmbed["embeds"])
	assert.NotNil(t, childEmbed["embeds"].(map[string]interface{})["data"])
}