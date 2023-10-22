package token

import "time"

//Maker is an interface to manage tokens
type Maker interface {
	//this creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	//this verify token to provide access
	VerifyToken(token string) (*Payload, error)
}
