package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/skhanal5/txs/internal/api/service"
)

type contextKey string

const userIDContextKey = contextKey("userID")

type AuthMiddleware struct {
	handler http.Handler
	secret  []byte
	skipPaths  map[string]struct{}
}

func NewAuthMiddleware(handler http.Handler, secret []byte, skipPaths []string) *AuthMiddleware {
	pathsMap := make(map[string]struct{}, len(skipPaths))
	for _, p := range skipPaths {
		pathsMap[p] = struct{}{}
	}

	return &AuthMiddleware{
		handler:   handler,
		secret:    secret,
		skipPaths: pathsMap,
	}
}

func (am *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, skip := am.skipPaths[r.URL.Path]; skip {
		am.handler.ServeHTTP(w, r)
		return
	}
	
	tokenStr, err := am.getTokenFromHeader(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, claims, err := am.validateToken(tokenStr)
	if err != nil || !token.Valid {
		http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	if claims.UserID == "" {
		http.Error(w, "user_id claim missing", http.StatusUnauthorized)
		return
	}

	r = am.attachUserToContext(r, claims.UserID)
	am.handler.ServeHTTP(w, r)
}

func (am *AuthMiddleware) getTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing Authorization header")
	}
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("invalid Authorization header format")
	}
	return strings.TrimPrefix(authHeader, "Bearer "), nil
}

func (am *AuthMiddleware) validateToken(tokenStr string) (*jwt.Token, *service.Claims, error) {
	claims := &service.Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return am.secret, nil
	})

	if err != nil {
		return nil, nil, err
	}

	return token, claims, nil
}

func (am *AuthMiddleware) attachUserToContext(r *http.Request, userID string) *http.Request {
	ctx := context.WithValue(r.Context(), userIDContextKey, userID)
	return r.WithContext(ctx)
}


func GetUserID(r *http.Request) (string, bool) {
	userID, ok := r.Context().Value(userIDContextKey).(string)
	return userID, ok
}
