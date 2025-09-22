package config

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Package-level cache for configuration values loaded from either .env or process env.
var (
	loaded bool
	once   sync.Once
	values map[string]string
)

// load reads a .env file if present from the project root (current working dir),
// otherwise falls back to process environment variables. Values from process env
// override .env entries.
func load() {
	values = map[string]string{}

	// Try to load .env from current working directory or parent directories (up to 3 levels)
	candidates := []string{".env"}
	// attempt parent directories as well to be robust when running from subdirs
	cwd, err := os.Getwd()
	if err == nil {
		candidates = append(candidates,
			filepath.Join(cwd, ".env"),
			filepath.Join(cwd, "..", ".env"),
			filepath.Join(cwd, "..", "..", ".env"),
		)
	}

	for _, p := range candidates {
		f, err := os.Open(p)
		if err != nil {
			continue
		}
		s := bufio.NewScanner(f)
		for s.Scan() {
			line := strings.TrimSpace(s.Text())
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			if i := strings.Index(line, "="); i >= 0 {
				k := strings.TrimSpace(line[:i])
				v := strings.TrimSpace(line[i+1:])
				v = strings.Trim(v, `"'`)
				if k != "" {
					values[k] = v
				}
			}
		}
		_ = f.Close()
		break
	}

	// Overlay with process env
	for _, kv := range os.Environ() {
		i := strings.Index(kv, "=")
		if i < 0 {
			continue
		}
		k := kv[:i]
		v := kv[i+1:]
		values[k] = v
	}
	loaded = true
}

func get(key string) string {
	once.Do(load)
	return values[key]
}

func require(key string) (string, error) {
	v := get(key)
	if strings.TrimSpace(v) == "" {
		return "", errors.New("missing required env: " + key)
	}
	return v, nil
}

// Public accessors for env configurations

func SupabaseURL() (string, error) {
	return require("SUPABASE_URL")
}

func SupabaseAnonKey() (string, error) {
	return require("SUPABASE_ANON_KEY")
}

// JWT secret used by Supabase for signing access tokens (GoTrue JWT secret)
func SupabaseJWTSecret() (string, error) {
	return require("SUPABASE_JWT_SECRET")
}

// JWKS URL for asymmetric JWT verification (e.g., https://<project>.supabase.co/auth/v1/keys)
func SupabaseJWKSURL() (string, error) {
	return require("SUPABASE_JWKS_URL")
}

// Name of the auth cookie storing the access token
func AuthCookieName() (string, error) {
	return require("AUTH_COOKIE_NAME")
}
