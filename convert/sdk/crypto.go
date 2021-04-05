package sdk

import (
	"github.com/onflow/flow-go/crypto"
	"github.com/onflow/flow-go/crypto/hash"
)

// StringToSigningAlgorithm converts a string to a SigningAlgorithm.
func stringToSigningAlgorithm(s string) crypto.SigningAlgorithm {
	switch s {
	case crypto.ECDSAP256.String():
		return crypto.ECDSAP256
	case crypto.ECDSASecp256k1.String():
		return crypto.ECDSASecp256k1
	default:
		return crypto.UnknownSigningAlgorithm
	}
}

// StringToHashingAlgorithm converts a string to a HashingAlgorithm.
func stringToHashingAlgorithm(s string) hash.HashingAlgorithm {
	switch s {
	case hash.SHA2_256.String():
		return hash.SHA2_256
	case hash.SHA3_256.String():
		return hash.SHA3_256
	default:
		return hash.UnknownHashingAlgorithm
	}
}
