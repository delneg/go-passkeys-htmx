// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import ()

type Counter struct {
	UserID string
	Count  int64
}

type Credential struct {
	ID              []byte
	PublicKey       []byte
	AttestationType string
	Transport       []byte
	Flags           []byte
	Authenticator   []byte
	UserID          []byte
}

type User struct {
	ID          []byte
	Name        string
	DisplayName string
}
