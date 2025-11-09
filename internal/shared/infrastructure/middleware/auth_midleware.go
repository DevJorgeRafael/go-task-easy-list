package middleware

import (
	"context"
	sharedhttp "go-task-easy-list/internal/shared/http"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	jwtSecret string
}

func NewAuthMiddleware(jwtSecret string) *AuthMiddleware {
	return &AuthMiddleware{jwtSecret: jwtSecret}
}

// RequireAuth valida el JWT y extrae el userId
func (m *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Extraer token del header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			sharedhttp.ErrorResponse(w, http.StatusUnauthorized, "Token no proporcionado")
			return
		}

		// 2. Verificar formato "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			sharedhttp.ErrorResponse(w, http.StatusUnauthorized, "Formato de token inválido")
			return
		}

		tokenString := parts[1]

		// 3. Parsear y validar JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verificar algoritmo
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(m.jwtSecret), nil
		})

		if err != nil || !token.Valid {
			sharedhttp.ErrorResponse(w, http.StatusUnauthorized, "Token inválido o expirado")
			return
		}

		// 4. Extraer claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			sharedhttp.ErrorResponse(w, http.StatusUnauthorized, "Claims inválidos")
			return
		}

		// 5. Extraer userId
		userID, ok := claims["userId"].(string)
		if !ok {
			sharedhttp.ErrorResponse(w, http.StatusUnauthorized, "userId no encontrado en token")
			return
		}

		// 6. Guardar userId en el contexto
		ctx := context.WithValue(r.Context(), "userId", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}