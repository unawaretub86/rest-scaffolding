package token_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
	tokenP "github.com/unawaretub86/rest-scaffolding/token"
)

func TestJWTMaker(t *testing.T) {
	maker, err := tokenP.NewJWTMaker("asdcasdsadasdscasdascascascasdasd")
	require.NoError(t, err)

	usernName := "phil"
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(usernName, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, usernName, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := tokenP.NewJWTMaker("asdcasdsadasdscasdascascascasdasd")
	require.NoError(t, err)

	usernName := "phil"

	token, err := maker.CreateToken(usernName, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, tokenP.ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTToken(t *testing.T) {

	usernName := "phil"
	duration := time.Minute

	payload, err := tokenP.NewPayload(usernName, duration)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	makerToken, err := tokenP.NewJWTMaker("asdcasdsadasdscasdascascascasdasd")
	require.NoError(t, err)

	payload, err = makerToken.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, tokenP.ErrInvalidToken.Error())
	require.Nil(t, payload)
}
