/**
 * Copyright (c) 2018 Adrian K <adrian.git@kuguar.dev>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package boot

import (
	"context"
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/adrianpk/kamien/api"
	"{{.Package}}/common"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	goerrors "github.com/go-errors/errors"
)

var (
	// UserCtxKey - Package standard User context key.
	UserCtxKey = ContextKey("user")
)

// ContextKey - Package standard context key.
type ContextKey string

func (c ContextKey) String() string {
	return "{{.Package}}#" + string(c)
}

// AppClaims provides custom claim for JWT
type AppClaims struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// using asymmetric crypto/RSA keys
// location of private/public key files
const (
	// openssl genrsa -out {{.AppNameLowercase}}.rsa 1024
	privKeyPath = "{{.AppNameLowercase}}.rsa"
	// openssl rsa -in {{.AppNameLowercase}}.rsa -pubout > {{.AppNameLowercase}}.rsa.pub
	pubKeyPath = "{{.AppNameLowercase}}.rsa.pub"
)

// Private key for signing and public key for verification
var (
	//verifyKey, signKey []byte
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// Read the key files before starting http handlers
func initKeys() {
	fullPrivKeyPath := path.Join(*AssetsDir, "keys", privKeyPath)
	log.Infof("Reading private key from '%s'", fullPrivKeyPath)
	signBytes, err := ioutil.ReadFile(fullPrivKeyPath)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	fullPubKeyPath := path.Join(*AssetsDir, "keys", pubKeyPath)
	log.Infof("Reading public key from '%s'", fullPubKeyPath)
	verifyBytes, err := ioutil.ReadFile(fullPubKeyPath)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
}

// GenerateJWT generates a new JWT token
func GenerateJWT(userID string, username, role string) (string, error) {
	// Create the Claims
	claims := AppClaims{
		userID,
		username,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 20).Unix(),
			Issuer:    "admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

// Authorize Middleware for validating JWT tokens
func Authorize(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Get token from request
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})

	if err != nil {
		werr := wrapError(err)
		switch err.(type) {
		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired: //JWT expired
				api.RenderErr(rw, r, common.ErrTokenExpired, werr, "Auth token expired", 401)
				return
			default:
				api.RenderErr(rw, r, common.ErrTokenParsing, werr, "Cannot parse auth token", 401)
				return
			}
		default:
			api.RenderErr(rw, r, common.ErrTokenParsing, werr, "Cannot parse auth token", 401)
			return
		}
	}
	if token.Valid {
		// Set user name to HTTP context
		claims := token.Claims.(*AppClaims)
		//claims := AppClaims{UserID: "5958b185-8150-4aae-b53f-0c44771ddec5", Username: "admin", Role: "admin"}
		ctx := context.WithValue(r.Context(), UserCtxKey, *claims)
		r = r.WithContext(ctx)
		next(rw, r)
	} else {
		api.RenderErr(rw, r, common.ErrTokenInvalid, wrapError(common.ErrTokenInvalid), "Non valid auth token", 401)
	}
}

// TokenFromAuthHeader is a "TokenExtractor" that takes a given request and extracts
// the JWT token from the Authorization header.
func TokenFromAuthHeader(r *http.Request) (string, error) {
	// Look for an Authorization header
	if ah := r.Header.Get("Authorization"); ah != "" {
		// Should be a bearer token
		if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
			return ah[7:], nil
		}
	}
	return "", errors.New("No token in the HTTP request")
}

// Private
// wrapError - Wrap error to preserve stacktrace
func wrapError(err error) *goerrors.Error {
	return goerrors.Wrap(err, 1)
}
