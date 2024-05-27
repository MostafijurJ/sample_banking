package token

import (
	"fmt"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	"time"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetrucKey []byte
}

func NewPasetoMaker(symmetrucKey string) (Maker, error) {
	if len(symmetrucKey) < chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size, must be at least %d characters", minSecretKeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetrucKey: []byte(symmetrucKey),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	return maker.paseto.Encrypt(maker.symmetrucKey, payload, nil)
}
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := maker.paseto.Decrypt(token, maker.symmetrucKey, payload, nil)
	if err != nil {
		return nil, err
	}

	if err := payload.Valid(); err != nil {
		return nil, err
	}

	return payload, nil
}
