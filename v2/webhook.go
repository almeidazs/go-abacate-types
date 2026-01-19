package v2

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
)

const AbacatePayPublicKey = "t9dXRhHHo3yDEj5pVDYz0frf7q6bMKyMRmxxCPIPp3RCplBfXRxqlC6ZpiWmOqj4L63qEaeUOtrCI8P0VMUgo6iIga2ri9ogaHFs0WIIywSMg0q7RmBfybe1E5XJcfC4IW3alNqym0tXoAKkzvfEjZxV6bE0oG2zJrNNYmUCKZyV0KZ3JS8Votf9EAWWYdiDkMkpbMdPggfh1EqHlVkMiTady6jOR3hyzGEHrIz2Ret0xHKMbiqkr9HS1JhNHDX9"

// VerifyWebhookSignature verifies whether the webhook signature is valid.
func VerifyWebhookSignature(rawBody string, signature string) bool {
	mac := hmac.New(sha256.New, []byte(AbacatePayPublicKey))
	mac.Write([]byte(rawBody))

	expectedSig := mac.Sum(nil)

	receivedSig, err := base64.StdEncoding.DecodeString(signature)

	if err != nil {
		return false
	}

	return subtle.ConstantTimeCompare(expectedSig, receivedSig) == 1
}
