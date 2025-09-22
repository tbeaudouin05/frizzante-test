package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"main/lib/config"
)

// Claims represents the standard claims plus any custom ones we might read.
// Kept minimal; access extras via MapClaims.

type Claims struct {
	Subject string
	Email   string
	Role    string
	Raw     jwt.MapClaims
}

// ParseAndVerify parses and verifies a JWT access token.
// If SUPABASE_JWKS_URL is set, it verifies using JWKS (RS256 recommended by Supabase).
// Otherwise, it returns an error.
func ParseAndVerify(tokenString string) (Claims, error) {
	var empty Claims
	if strings.TrimSpace(tokenString) == "" {
		return empty, errors.New("empty token")
	}

	jwksURL, jerr := config.SupabaseJWKSURL()
	if jerr != nil || strings.TrimSpace(jwksURL) == "" {
		return empty, errors.New("SUPABASE_JWKS_URL is required for JWT verification")
	}

	// JWKS path (asymmetric; Supabase may use RS256 or ES256)
	tok, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Expect RSA or ECDSA algorithms
		switch t.Method.(type) {
		case *jwt.SigningMethodRSA, *jwt.SigningMethodECDSA:
			// ok
		default:
			return nil, fmt.Errorf("unexpected signing method: %T", t.Method)
		}
		kid, _ := t.Header["kid"].(string)
		if strings.TrimSpace(kid) == "" {
			return nil, errors.New("missing kid header for JWKS")
		}
		pub, perr := jwksKey(jwksURL, kid)
		if perr != nil {
			return nil, perr
		}
		return pub, nil
	})
	if err != nil {
		return empty, err
	}
	if !tok.Valid {
		return empty, errors.New("invalid token")
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return empty, errors.New("invalid claims")
	}

	out := Claims{Raw: claims}
	if sub, ok := claims["sub"].(string); ok {
		out.Subject = sub
	}
	if eml, ok := claims["email"].(string); ok {
		out.Email = eml
	}
	if role, ok := claims["role"].(string); ok {
		out.Role = role
	}
	return out, nil
}

// jwksKey fetches the JWKS from jwksURL and returns a public key (*rsa.PublicKey or *ecdsa.PublicKey)
// that matches the given kid.
func jwksKey(jwksURL, kid string) (interface{}, error) {
	// Fetch JWKS
	resp, err := http.Get(jwksURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("jwks http status: %d", resp.StatusCode)
	}

	var jwks struct {
		Keys []struct {
			Kty string `json:"kty"`
			Kid string `json:"kid"`
			Use string `json:"use"`
			Alg string `json:"alg"`
			// RSA
			N string `json:"n"`
			E string `json:"e"`
			// EC
			Crv string `json:"crv"`
			X   string `json:"x"`
			Y   string `json:"y"`
		} `json:"keys"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return nil, err
	}
	for _, k := range jwks.Keys {
		if k.Kid != kid {
			continue
		}
		switch strings.ToUpper(k.Kty) {
		case "RSA":
			// n and e are base64url without padding
			nb, err := base64.RawURLEncoding.DecodeString(k.N)
			if err != nil {
				return nil, fmt.Errorf("invalid n: %w", err)
			}
			eb, err := base64.RawURLEncoding.DecodeString(k.E)
			if err != nil {
				return nil, fmt.Errorf("invalid e: %w", err)
			}
			// Convert e to int
			var eInt int
			if len(eb) == 0 {
				return nil, errors.New("empty exponent")
			}
			// Common exponents are small (e.g., 65537). Convert big-endian bytes to int.
			for _, b := range eb {
				eInt = eInt<<8 | int(b)
			}
			n := new(big.Int).SetBytes(nb)
			return &rsa.PublicKey{N: n, E: eInt}, nil
		case "EC":
			// ES256 typically uses P-256 curve
			if strings.EqualFold(k.Crv, "P-256") || strings.EqualFold(k.Crv, "secp256r1") {
				xb, err := base64.RawURLEncoding.DecodeString(k.X)
				if err != nil {
					return nil, fmt.Errorf("invalid x: %w", err)
				}
				yb, err := base64.RawURLEncoding.DecodeString(k.Y)
				if err != nil {
					return nil, fmt.Errorf("invalid y: %w", err)
				}
				x := new(big.Int).SetBytes(xb)
				y := new(big.Int).SetBytes(yb)
				return &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}, nil
			}
			return nil, fmt.Errorf("unsupported EC curve: %s", k.Crv)
		default:
			return nil, fmt.Errorf("unsupported kty: %s", k.Kty)
		}
	}
	return nil, errors.New("kid not found in JWKS")
}
