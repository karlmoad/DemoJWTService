package authentication

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/gbrlsnchs/jwt/v2"
	"github.com/gorilla/context"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	BEARER_SCHEMA string = "Bearer "
	PEMSTART = "-----BEGIN CERTIFICATE-----\n"
	PEMEND = "\n-----END CERTIFICATE-----\n"
	ENV_KEY ="AZURE_APP_KEY"
)

var msKeySource *microsoftKeySource
var appKey string

func AuthenticationHandler(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check status of the microsoft key source info
		if msKeySource == nil || msKeySource.IsExpired() {
			ms, err := NewMicrosoftKeySource()
			if err != nil {
				w.Write([]byte("Unable to verify token"))
				w.WriteHeader(500)
				return
			}
			msKeySource = ms
		}

		authToken, err := getTokenFromRequest(r)
		if err != nil{
			w.Write([]byte(err.Error()))
			w.WriteHeader(400)
			return
		}

		token, err := validateToken(authToken)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(400)
			return
		}

		context.Set(r, "TOKEN", *token)
		next(w, r)
	})
}

func validateToken(authToken string) (*Token,error) {

	var token Token


	payload, sig, err := jwt.Parse(authToken)
	if err != nil {
		return nil, err
	}

	var jot jwt.JWT
	err = jwt.Unmarshal(payload, &jot)
	if err != nil {
		return nil, err
	}

	err = jwt.Unmarshal(payload, &token)
	if err != nil {
		return nil, err
	}

	keys, err := msKeySource.GetKeys(jot.KeyID())
	if len(keys) <= 0 || err != nil {
		return nil, errors.New("error retrieving token verification key")
	}

	sCert := fmt.Sprintf("%s%s%s",PEMSTART, keys[0], PEMEND)

	cpem, _ := pem.Decode([]byte(sCert))

	if cpem == nil {
		return nil, errors.New("error retrieving token verification key, certificate pem is nil")
	}

	cert, err := x509.ParseCertificate(cpem.Bytes)
	if err != nil {
		return nil, errors.New("error retrieving token verification key, error parsing certificate")
	}

	pubkey := cert.PublicKey.(*rsa.PublicKey)

	rs256 := jwt.NewRS256(nil, pubkey)

	err = rs256.Verify(payload,sig)
	if err != nil {
		return nil, errors.New("unable to verify signature")
	}

	now := time.Now()

	iatValidator := jwt.IssuedAtValidator(now)
	expValidator := jwt.ExpirationTimeValidator(now)
	audValidator := jwt.AudienceValidator(appKey)
	if err = jot.Validate(iatValidator, expValidator, audValidator); err != nil {
		switch err {
		case jwt.ErrIatValidation:
			return nil, errors.New("invalid issued time on token")
		case jwt.ErrExpValidation:
			return nil, errors.New("token has expired")
		case jwt.ErrAudValidation:
			return nil, errors.New("invalid token audience")
		}
	}

	return &token, nil
}

func getTokenFromRequest(req *http.Request) (string, error) {
	// Grab the raw Authoirzation header
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization header required")
	}

	// Confirm the request is sending Bearer Authentication credentials.
	if !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
		return "", errors.New("Authorization requires Bearer scheme")
	}

	token := strings.TrimSpace(authHeader[len(BEARER_SCHEMA):])

	return token, nil
}

func init() {
	msKeySource,_ = NewMicrosoftKeySource()
	appKey = os.Getenv(ENV_KEY)
}