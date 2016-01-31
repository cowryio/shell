package shell

import "testing"
import "time"
import "github.com/cowryio/shell/Godeps/_workspace/src/github.com/stretchr/testify/assert"


var TEST_ISSUER_PRIVATE_KEYS = []string{
	"-----BEGIN RSA PRIVATE KEY-----\nMIICWwIBAAKBgQCroZieOAo9stcf6R6eWfo51VCvK8cLdNS577m/HIFOmEd1CDi/\nu7agGzpehNAhHpr5NVjQZ4Te+KMRn9SnpUK2hc8dUU25PQolsOEwePVQ18hHNK4Y\n2JvOY/f8KCO2hhrS6uuP6eedpnSdulS1OXHTL6ZxQmBd9F33gLT6BERHQwIDAQAB\nAoGAEZ/0ljrXAmL9KG++DzDaO1omgPaT6B9FQRrXDkMVHEcS/3eqrDXQmTxykAY/\ngUctTu4lgrE+uc76n/Kz2ctkwEKIKet56ylqp+wlEUt1G+udoi07tgd7XyxzoUJm\nZwSm89gKh+mEPxni0FrBNg6dR0n2gvKRecnXqyoGVOHZITECQQDXgRJyrzgc/JhB\nSOBznEjtXAZXRRu3o9UznztjU9Xz7NWXTVuHu8WqYmGWCOqnysMhXJ3xBddJyDTF\njuOJ0123AkEAy+H+3POcT2FDOuluqPmAZQAUU6Nxtbj02/JJtOy7jq5jnN27HVC3\nuQzmfsS5J2XeQQodOUwOy2Ub57/OMrMi1QJAGZsZgQz2wuL0iFVLbhE0zRcxHa91\ncqWB0Kdr3Ap7EoeifV7QsFkMTIlyBOy8TQGXm+AwWBIUmYyzUIIA4UB/EwJAO+Bo\nSB2nZ0yqQO/zVt7HjWIDljinGXZzOvEiImdwAcxHZvdbj5V4D3mxa8N8mQx6xGEj\nCgPDSIquMlaLSSqA7QJAAbQPa0frCkm1rkWWZ7QwGm7ptzOACwFEGefm/1mhmw3a\nvoWRTHhrDuEbeVH3iF8MWhLJLPFtuSShiQMsrVbXPA==\n-----END RSA PRIVATE KEY-----",
	"-----BEGIN KEY-----\nMIICWwIBAAKBgQCroZieOAo9stcf6R6eWfo51VCvK8cLdNS577m/HIFOmEd1CDi/\nu7agGzpehNAhHpr5NVjQZ4Te+KMRn9SnpUK2hc8dUU25PQolsOEwePVQ18hHNK4Y\n2JvOY/f8KCO2hhrS6uuP6eedpnSdulS1OXHTL6ZxQmBd9F33gLT6BERHQwIDAQAB\nAoGAEZ/0ljrXAmL9KG++DzDaO1omgPaT6B9FQRrXDkMVHEcS/3eqrDXQmTxykAY/\ngUctTu4lgrE+uc76n/Kz2ctkwEKIKet56ylqp+wlEUt1G+udoi07tgd7XyxzoUJm\nZwSm89gKh+mEPxni0FrBNg6dR0n2gvKRecnXqyoGVOHZITECQQDXgRJyrzgc/JhB\nSOBznEjtXAZXRRu3o9UznztjU9Xz7NWXTVuHu8WqYmGWCOqnysMhXJ3xBddJyDTF\njuOJ0123AkEAy+H+3POcT2FDOuluqPmAZQAUU6Nxtbj02/JJtOy7jq5jnN27HVC3\nuQzmfsS5J2XeQQodOUwOy2Ub57/OMrMi1QJAGZsZgQz2wuL0iFVLbhE0zRcxHa91\ncqWB0Kdr3Ap7EoeifV7QsFkMTIlyBOy8TQGXm+AwWBIUmYyzUIIA4UB/EwJAO+Bo\nSB2nZ0yqQO/zVt7HjWIDljinGXZzOvEiImdwAcxHZvdbj5V4D3mxa8N8mQx6xGEj\nCgPDSIquMlaLSSqA7QJAAbQPa0frCkm1rkWWZ7QwGm7ptzOACwFEGefm/1mhmw3a\nvoWRTHhrDuEbeVH3iF8MWhLJLPFtuSShiQMsrVbXPA==\n-----END RSA PRIVATE KEY-----",
}

var TEST_SHELL_DATA = []string {
	`{"signatures":{"meta":""},"meta":{"created_at":1453975575,"genesis":true,"shell_id":"4417781906fb0a89c295959b0df01782dbc4dc9f","shell_type":"currency"},"ownership":null,"embeds":[],"attributes":null}`,
}

// TestCreateAShell create a valid, error free shell
func TestCreateAShell(t *testing.T) {
	shellID := NewID()
	var meta = map[string]interface{}{
		"shell_id": shellID,
		"shell_type": "currency",
		"created_at": time.Now().Unix(),
	}
	sh, err := Create(meta, TEST_ISSUER_PRIVATE_KEYS[0])
	assert.Nil(t, err)
	assert.Equal(t, sh.Meta["shell_id"], shellID)
	assert.NotEmpty(t, sh.Signatures["meta"])
}

// TestMustProvideMetaWithContent test that a map describing the `meta` block is required
func TestMustProvideMetaWithContent(t *testing.T) {
	_, err := Create(make(map[string]interface{}), TEST_ISSUER_PRIVATE_KEYS[0])
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "`meta` block is missing `shell_id` property")
}

// TestInvalidPrivateKey tests that an invalid private key returns an error
func TestInvalidPrivateKey(t *testing.T) {
	var issuerPrivateKey = TEST_ISSUER_PRIVATE_KEYS[1]
	shellID := NewID()
	var meta = map[string]interface{}{
		"shell_id": shellID,
		"shell_type": "currency",
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

// TestLoadJSON tests that a valid shell json string can be loaded into a shell object
func TestLoadJSON(t *testing.T) {
	txt := TEST_SHELL_DATA[0]
	shell, err := LoadJSON(txt)
	assert.Nil(t, err);
	assert.IsType(t, &Shell{}, shell)
}

// TestLoadEncodedJSON tests that a base 64 encoded json string can be loaded into a shell object
func TestLoadEncodedJSON(t *testing.T) {
	encodedJSON := ToBase64([]byte(TEST_SHELL_DATA[0]))
	shell, err := Load(encodedJSON)
	assert.Nil(t, err);
	assert.IsType(t, &Shell{}, shell)
}

// TestCannotLoadInvalidEncodedJSON tests that an incorrect base64 encode will result in error when attempting to load into a shell
func TestCannotLoadInvalidEncodedJSON(t *testing.T) {
	encodedJSON := ToBase64([]byte("abcde"))
	_, err := Load(encodedJSON)
	assert.NotNil(t, err);
	assert.Equal(t, err.Error(), "unable to parse json string")
}

// TestCorrectlySignMeta tests that a shell is correctly signed
func TestCorrectlySignMeta(t *testing.T) {
	txt := TEST_SHELL_DATA[0]
	shell, err := LoadJSON(txt)
	assert.Nil(t, err);
	expectedCanonicalMapString := GetCanonicalMapString(shell.Meta)
	signer, err := ParsePrivateKey([]byte(TEST_ISSUER_PRIVATE_KEYS[0]))
	assert.Nil(t, err)
	expectedSignature, err := signer.Sign([]byte(expectedCanonicalMapString))
	assert.Nil(t, err)
	signature, err := shell.Sign("meta", TEST_ISSUER_PRIVATE_KEYS[0])
	assert.Nil(t, err)
	assert.Equal(t, expectedSignature, signature)
	assert.Equal(t, expectedSignature, shell.Signatures["meta"])
}

// TestCannotSignUnknownBlock tests that an error will occur when attempting to sign an unknown block
func TestCannotSignUnknownBlock(t *testing.T) {
	txt := TEST_SHELL_DATA[0]
	shell, err := LoadJSON(txt)
	assert.Nil(t, err)
	_, err = shell.Sign("unknown_block", TEST_ISSUER_PRIVATE_KEYS[0])
	assert.NotNil(t, err)
	expectedMsg := "block unknown"
	assert.Equal(t, err.Error(), expectedMsg)
}

// TestAddMeta tests that a meta block can be assigned and signed successful using the TestAddMeta method
func TestAddMeta(t *testing.T) {
	shellID := NewID()
	var meta = map[string]interface{}{
		"shell_id": shellID,
		"shell_type": "currency",
		"created_at": time.Now().Unix(),
	}
	sh := Empty()
	err := sh.AddMeta(meta, TEST_ISSUER_PRIVATE_KEYS[0])
	assert.Nil(t, err)
	assert.Equal(t, sh.Meta["shell_id"], meta["shell_id"])
	assert.NotNil(t, sh.Signatures["meta"])
}


