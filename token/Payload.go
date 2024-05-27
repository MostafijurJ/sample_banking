package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var (
	ErrExpiredToken = errors.New("token is expired")
	ErrInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		Id:        id,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func (p *Payload) GetIssuer() (string, error) {
	// Assuming the issuer is the username in this case
	return p.Username, nil
}

func (p *Payload) GetSubject() (string, error) {
	// Assuming the subject is the ID in this case
	return p.Id.String(), nil
}

func (p *Payload) GetAudience() (jwt.ClaimStrings, error) {
	// Assuming the audience is a single entry with the username in this case
	return jwt.ClaimStrings{p.Username}, nil
}

func (p *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	n := jwt.NumericDate{
		Time: p.ExpiredAt,
	}
	return &n, nil
}

func (p *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	n := jwt.NumericDate{
		Time: p.IssuedAt,
	}
	return &n, nil
}
func (p *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	n := jwt.NumericDate{
		Time: p.IssuedAt,
	}
	return &n, nil
}
