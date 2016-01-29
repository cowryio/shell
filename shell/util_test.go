package shell

import "testing"
import "github.com/cowryio/shell-go/shell/Godeps/_workspace/src/github.com/stretchr/testify/assert"

// TestByteArrToString tests that a byte array is properly converted to string
func TestByteArrToString(t *testing.T) {
	var s = ByteArrToString([]byte("hello"))
	assert.Equal(t, s, "hello")
}

// TestSha1 tests that a SHA1 hash will return an exact hash for a specific string everytime
func TestSha1(t *testing.T) {
	var h = Sha1("hello")
	assert.Equal(t, h, Sha1("hello"))
}

// TestNewID tests should create an id with 40 characters
func TestNewID(t *testing.T) {
	var id = NewID()
	assert.Equal(t, len(id), 40)
}

// TestValueNotInStringSlice tests that a string value is not contained in a string slice
func TestValueNotInStringSlice(t *testing.T) {
	var ss = []string{"john", "doe"}
	var r = InStringSlice(ss, "jane")
	assert.Equal(t, r, false)
}

// TestGetMapKeys tests that GetMapKeys will return a list of all keys it contains
func TestGetMapKeys(t *testing.T) {
	var keys = GetMapKeys(map[string]interface{}{
		"key1": "0",
		"key2": "1",
	})
	r := InStringSlice(keys, "key1")
	r2 := InStringSlice(keys, "key2")
	assert.Equal(t, len(keys), 2)
	assert.Equal(t, r, true)
	assert.Equal(t, r2, true)
}

// TestGetCanonicalMapString tests that a predictable, normalized string version of a map is returned
func TestGetCanonicalMapString(t *testing.T) {
	var m = map[string]interface{}{
		"pete":    2,
		"abraham": 30,
		"jamie":   "jedi",
		"zebra": map[string]interface{}{
			"xonia":  "fighter",
			"belami": "protector",
		},
	}
	expected := "abraham:30:jamie:jedi:pete:2:zebra:belami:protector:xonia:fighter"
	assert.Equal(t, expected, GetCanonicalMapString(m))
}

// TestHasKey tests that a key exist or doesn't exist in a map
func TestHasKey(t *testing.T) {
	var m = map[string]interface{}{
		"stuff_a": 2,
		"stuff_b": 3,
	}
	assert.Equal(t, HasKey(m, "stuff_b"), true)
	assert.Equal(t, HasKey(m, "stuff_a"), true)
	assert.Equal(t, HasKey(m, "stuff_c"), false)
}

// TestIsStringValue tests that a variable holds a string value or not
func TestIsStringValue(t *testing.T) {
	assert.Equal(t, IsStringValue("lorem"), true)
	assert.Equal(t, IsStringValue(20), false)
}

// TestValueInStringSlice tests that a string value is contained in a string slice or not
func TestValueInStringSlice(t *testing.T) {
	var ss = []string{"john", "doe"}
	assert.Equal(t, InStringSlice(ss, "john"), true)
	assert.Equal(t, InStringSlice(ss, "jane"), false)
}

// TestEnv tests that a default value is returned when environment value being fetched is not set
func TestEnv(t *testing.T) {
	assert.Equal(t, Env("SOME_VAR", "default_value"), "default_value")
}

// TestIsMapOfAny tests that a variable's value type is a map of any type as value
func TestIsMapOfAny(t *testing.T) {
	var m = make(map[string]interface{})
	assert.Equal(t, IsMapOfAny(m), true)
	assert.Equal(t, IsMapOfAny(10), false)
}
