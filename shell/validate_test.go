package shell

import "testing"
import "github.com/cowryio/shell-go/shell/Godeps/_workspace/src/github.com/stretchr/testify/assert"

// TestInvalidJSON tests that an invalid json string returns an error
func TestInvalidJSON(t *testing.T) {
	var str = `{ "meta : "" }`
	err := Validate(str)
	assert.NotNil(t, err)
	expectedMsg := `unable to parse json string`
	assert.Equal(t, expectedMsg, err.Error())
}

// TestMetaMustHaveShellIDProperty tests that a meta block data must have a `shell_id` property
func TestMetaMustHaveShellIDProperty(t *testing.T) {
	var d = make(map[string]interface{})
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta` block is missing `shell_id` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestMetaMustHaveShellTypeProperty tests that shell_type property is set
func TestMetaMustHaveShellTypeProperty(t *testing.T) {
	d := map[string]interface{}{
		"shell_id": 1234,
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta` block is missing `shell_type` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestMetaMustHaveCreatedAtProperty tests that created_at property is set
func TestMetaMustHaveCreatedAtProperty(t *testing.T) {
	d := map[string]interface{}{
		"shell_id":   1234,
		"shell_type": "coupon",
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta` block is missing `created_at` property"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestShellIDNotString tests that shell_id value type must be string
func TestShellIDMustString(t *testing.T) {
	d := map[string]interface{}{
		"shell_id":   1234,
		"shell_type": "coupon",
		"created_at": 1000,
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta.shell_id` value type is invalid. Expects string value"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestShellIDLengthInvalid tests that a shell_id must have 40 characters
func TestShellIDLengthInvalid(t *testing.T) {
	d := map[string]interface{}{
		"shell_id":   "abcd",
		"shell_type": "coupon",
		"created_at": 1000,
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta.shell_id` must have 40 characters. Preferrable a UUIDv4 SHA1 hashed string"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestShellTypeMustBeString tests that shell_type value type must be string
func TestShellTypeMustBeString(t *testing.T) {
	d := map[string]interface{}{
		"shell_id":   Sha1("abcd"),
		"shell_type": 123,
		"created_at": 1000,
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`meta.shell_type` value type is invalid. Expects string value"
	assert.Equal(t, expectedMsg, err.Error())
}

// TestCreatedAtMustBeNumber tests that created_at value type must be a number
func TestCreatedAtMustBeNumber(t *testing.T) {
	d := map[string]interface{}{
		"shell_id":   Sha1("abcd"),
		"shell_type": "coupon",
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
		"shell_id":   Sha1("abcd"),
		"shell_type": "coupon",
		"created_at": 100000,
	}
	err := ValidateMetaBlock(d)
	assert.NotNil(t, err)
	expectedMsg := "`created_at` value is too far in the past. Expects unix time on or after 2016-01-28T11:06:15+01:00"
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

// TestAttributesSignatureTypeMustBeString tests that when `attributes` property is set, it's value type must be string
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
			"shell_id": "` + Sha1("stuff") + `", 
			"shell_type": "cur", 
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
			"shell_id": "` + Sha1("stuff") + `", 
			"shell_type": "cur", 
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
			"shell_id": "` + Sha1("stuff") + `", 
			"shell_type": "cur", 
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
			"shell_id": "` + Sha1("stuff") + `", 
			"shell_type": "cur", 
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
			"shell_id": "` + Sha1("stuff") + `", 
			"shell_type": "cur", 
			"created_at": 1453975575 
		},
		"ownership": { "stuff": "stuff" },
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
			"shell_id": "` + Sha1("stuff") + `", 
			"shell_type": "cur", 
			"created_at": 1453975575 
		},
		"ownership": { "stuff": "stuff" },
		"attributes": { "stuff": "stuff" }
	}`
	err := Validate(str)
	assert.NotNil(t, err)
	expectedMsg := "missing `attributes` property in `signatures` block"
	assert.Equal(t, expectedMsg, err.Error())
}
