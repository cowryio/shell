package shell

import "testing"
import "github.com/cowryio/shell-go/shell/Godeps/_workspace/src/github.com/stretchr/testify/assert"

// TestParsePublicKey tests that public key is valid
func TestParseGoodPublicKey(t *testing.T) {
	pubKey := "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCroZieOAo9stcf6R6eWfo51VCv\nK8cLdNS577m/HIFOmEd1CDi/u7agGzpehNAhHpr5NVjQZ4Te+KMRn9SnpUK2hc8d\nUU25PQolsOEwePVQ18hHNK4Y2JvOY/f8KCO2hhrS6uuP6eedpnSdulS1OXHTL6Zx\nQmBd9F33gLT6BERHQwIDAQAB\n-----END PUBLIC KEY-----"
	_, err := ParsePublicKey([]byte(pubKey))
	assert.Nil(t, err)
}

// TestUnsupportedPublicKeyType tests that a public key having an unsupported key type will not be parsed
func TestUnsupportedPublicKeyType(t *testing.T) {
	pubKey := "-----BEGIN KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCroZieOAo9stcf6R6eWfo51VCv\nK8cLdNS577m/HIFOmEd1CDi/u7agGzpehNAhHpr5NVjQZ4Te+KMRn9SnpUK2hc8d\nUU25PQolsOEwePVQ18hHNK4Y2JvOY/f8KCO2hhrS6uuP6eedpnSdulS1OXHTL6Zx\nQmBd9F33gLT6BERHQwIDAQAB\n-----END PUBLIC KEY-----"
	_, err := ParsePublicKey([]byte(pubKey))
	expectedMsg := `unsupported key type "KEY"`
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), expectedMsg)
}

// TestParseGoodPrivateKey tests that a private key is valid
func TestParseGoodPrivateKey(t *testing.T) {
	key := "-----BEGIN RSA PRIVATE KEY-----\nMIICWwIBAAKBgQCroZieOAo9stcf6R6eWfo51VCvK8cLdNS577m/HIFOmEd1CDi/\nu7agGzpehNAhHpr5NVjQZ4Te+KMRn9SnpUK2hc8dUU25PQolsOEwePVQ18hHNK4Y\n2JvOY/f8KCO2hhrS6uuP6eedpnSdulS1OXHTL6ZxQmBd9F33gLT6BERHQwIDAQAB\nAoGAEZ/0ljrXAmL9KG++DzDaO1omgPaT6B9FQRrXDkMVHEcS/3eqrDXQmTxykAY/\ngUctTu4lgrE+uc76n/Kz2ctkwEKIKet56ylqp+wlEUt1G+udoi07tgd7XyxzoUJm\nZwSm89gKh+mEPxni0FrBNg6dR0n2gvKRecnXqyoGVOHZITECQQDXgRJyrzgc/JhB\nSOBznEjtXAZXRRu3o9UznztjU9Xz7NWXTVuHu8WqYmGWCOqnysMhXJ3xBddJyDTF\njuOJ0123AkEAy+H+3POcT2FDOuluqPmAZQAUU6Nxtbj02/JJtOy7jq5jnN27HVC3\nuQzmfsS5J2XeQQodOUwOy2Ub57/OMrMi1QJAGZsZgQz2wuL0iFVLbhE0zRcxHa91\ncqWB0Kdr3Ap7EoeifV7QsFkMTIlyBOy8TQGXm+AwWBIUmYyzUIIA4UB/EwJAO+Bo\nSB2nZ0yqQO/zVt7HjWIDljinGXZzOvEiImdwAcxHZvdbj5V4D3mxa8N8mQx6xGEj\nCgPDSIquMlaLSSqA7QJAAbQPa0frCkm1rkWWZ7QwGm7ptzOACwFEGefm/1mhmw3a\nvoWRTHhrDuEbeVH3iF8MWhLJLPFtuSShiQMsrVbXPA==\n-----END RSA PRIVATE KEY-----"
	_, err := ParsePrivateKey([]byte(key))
	assert.Nil(t, err)
}

// TestUnsupportedPrivateKeyType tests that a private key having an unsupported key type will not be parsed
func TestUnsupportedPrivateKeyType(t *testing.T) {
	key := "-----BEGIN PRIVATE KEY-----\nMIICWwIBAAKBgQCroZieOAo9stcf6R6eWfo51VCvK8cLdNS577m/HIFOmEd1CDi/\nu7agGzpehNAhHpr5NVjQZ4Te+KMRn9SnpUK2hc8dUU25PQolsOEwePVQ18hHNK4Y\n2JvOY/f8KCO2hhrS6uuP6eedpnSdulS1OXHTL6ZxQmBd9F33gLT6BERHQwIDAQAB\nAoGAEZ/0ljrXAmL9KG++DzDaO1omgPaT6B9FQRrXDkMVHEcS/3eqrDXQmTxykAY/\ngUctTu4lgrE+uc76n/Kz2ctkwEKIKet56ylqp+wlEUt1G+udoi07tgd7XyxzoUJm\nZwSm89gKh+mEPxni0FrBNg6dR0n2gvKRecnXqyoGVOHZITECQQDXgRJyrzgc/JhB\nSOBznEjtXAZXRRu3o9UznztjU9Xz7NWXTVuHu8WqYmGWCOqnysMhXJ3xBddJyDTF\njuOJ0123AkEAy+H+3POcT2FDOuluqPmAZQAUU6Nxtbj02/JJtOy7jq5jnN27HVC3\nuQzmfsS5J2XeQQodOUwOy2Ub57/OMrMi1QJAGZsZgQz2wuL0iFVLbhE0zRcxHa91\ncqWB0Kdr3Ap7EoeifV7QsFkMTIlyBOy8TQGXm+AwWBIUmYyzUIIA4UB/EwJAO+Bo\nSB2nZ0yqQO/zVt7HjWIDljinGXZzOvEiImdwAcxHZvdbj5V4D3mxa8N8mQx6xGEj\nCgPDSIquMlaLSSqA7QJAAbQPa0frCkm1rkWWZ7QwGm7ptzOACwFEGefm/1mhmw3a\nvoWRTHhrDuEbeVH3iF8MWhLJLPFtuSShiQMsrVbXPA==\n-----END RSA PRIVATE KEY-----"
	_, err := ParsePublicKey([]byte(key))
	expectedMsg := `unsupported key type "PRIVATE KEY"`
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), expectedMsg)
}

// TestSignWithPrivateKey test that a valid private key will successfully sign a string
func TestSignWithPrivateKey(t *testing.T) {
	key := "-----BEGIN RSA PRIVATE KEY-----\nMIICWwIBAAKBgQCroZieOAo9stcf6R6eWfo51VCvK8cLdNS577m/HIFOmEd1CDi/\nu7agGzpehNAhHpr5NVjQZ4Te+KMRn9SnpUK2hc8dUU25PQolsOEwePVQ18hHNK4Y\n2JvOY/f8KCO2hhrS6uuP6eedpnSdulS1OXHTL6ZxQmBd9F33gLT6BERHQwIDAQAB\nAoGAEZ/0ljrXAmL9KG++DzDaO1omgPaT6B9FQRrXDkMVHEcS/3eqrDXQmTxykAY/\ngUctTu4lgrE+uc76n/Kz2ctkwEKIKet56ylqp+wlEUt1G+udoi07tgd7XyxzoUJm\nZwSm89gKh+mEPxni0FrBNg6dR0n2gvKRecnXqyoGVOHZITECQQDXgRJyrzgc/JhB\nSOBznEjtXAZXRRu3o9UznztjU9Xz7NWXTVuHu8WqYmGWCOqnysMhXJ3xBddJyDTF\njuOJ0123AkEAy+H+3POcT2FDOuluqPmAZQAUU6Nxtbj02/JJtOy7jq5jnN27HVC3\nuQzmfsS5J2XeQQodOUwOy2Ub57/OMrMi1QJAGZsZgQz2wuL0iFVLbhE0zRcxHa91\ncqWB0Kdr3Ap7EoeifV7QsFkMTIlyBOy8TQGXm+AwWBIUmYyzUIIA4UB/EwJAO+Bo\nSB2nZ0yqQO/zVt7HjWIDljinGXZzOvEiImdwAcxHZvdbj5V4D3mxa8N8mQx6xGEj\nCgPDSIquMlaLSSqA7QJAAbQPa0frCkm1rkWWZ7QwGm7ptzOACwFEGefm/1mhmw3a\nvoWRTHhrDuEbeVH3iF8MWhLJLPFtuSShiQMsrVbXPA==\n-----END RSA PRIVATE KEY-----"
	signer, err := ParsePrivateKey([]byte(key))
	assert.Nil(t, err)
	signature, err := signer.Sign([]byte("hello"))
	if !assert.Nil(t, err) {
		t.Error("unable to sign:", err)
	} else {
		expectedSignature := "a527f1e81f65a00e06ed09434d58ae54ec3acb6f097d1fce8d60781b157c0186da0a0dbefc8cceea5df77b2f95658d94f22fb2641eeb33674ebd85e472f65f2bb1243f2ea1d2d4b6cb20b60c77371eee3fe01227e2ccae1f7bb957d54814d1d9ceefd5b789b57fd10da69961d78e5e60a55326de185f51edcb5bf05bfa6c828b"
		if !assert.Equal(t, signature, expectedSignature) {
			t.Errorf("should match expected hex string")
		}
	}
}

// TestVerifyWithPublicKey tests that a valid public key will verify a signature
func TestVerifyWithPublicKey(t *testing.T) {
	pubKey := "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCroZieOAo9stcf6R6eWfo51VCv\nK8cLdNS577m/HIFOmEd1CDi/u7agGzpehNAhHpr5NVjQZ4Te+KMRn9SnpUK2hc8d\nUU25PQolsOEwePVQ18hHNK4Y2JvOY/f8KCO2hhrS6uuP6eedpnSdulS1OXHTL6Zx\nQmBd9F33gLT6BERHQwIDAQAB\n-----END PUBLIC KEY-----"
	signer, err := ParsePublicKey([]byte(pubKey))
	assert.Nil(t, err)
	signature := "a527f1e81f65a00e06ed09434d58ae54ec3acb6f097d1fce8d60781b157c0186da0a0dbefc8cceea5df77b2f95658d94f22fb2641eeb33674ebd85e472f65f2bb1243f2ea1d2d4b6cb20b60c77371eee3fe01227e2ccae1f7bb957d54814d1d9ceefd5b789b57fd10da69961d78e5e60a55326de185f51edcb5bf05bfa6c828b"
	if !assert.Nil(t, err) {
		t.Errorf("could not decode hex signature")
	}
	verified := signer.Verify([]byte("hello"), signature)
	if !assert.Nil(t, verified) {
		t.Errorf("could not verify signature")
	}
}
