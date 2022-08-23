package dtos

import (
	"errors"
	"time"
)

type Payload struct {
	AccountId string    `json:"account_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(accountId string) (*Payload, error) {
	payload := &Payload{
		AccountId: accountId,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().AddDate(1, 0, 0),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return errors.New("token has expired")
	}
	return nil
}
