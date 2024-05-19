package discordlib

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"log"
)

type DiscordCommandValidator struct {
	PublicKey string
}

func (d DiscordCommandValidator) VerifyRequest(signature string, timestamp string, body string) error {
	// Decode the public key
	pubKey, err := hex.DecodeString(d.PublicKey)
	if err != nil {
		return fmt.Errorf("failed to decode public key: %v", err)
	}
	log.Println(pubKey)

	// Decode the public key
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("failed to decode sig: %v", err)
	}
	log.Println(sig)

	// Verify the signature
	if !ed25519.Verify(pubKey, []byte(timestamp+body), sig) {
		return fmt.Errorf("invalid request signature")
	}

	return nil
}

func (d DiscordCommandValidator) IsPing(cmdType int) bool {
	return cmdType == 1
}