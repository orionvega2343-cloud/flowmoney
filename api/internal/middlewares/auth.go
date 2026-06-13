package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// Структура параметров
type Claims struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func AuthMiddleware(next http.Handler, secret string) http.Handler {
	//Замыкание
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenStr := r.Header.Get("Authorization") //Получаем токен по тегу
		//Валидируем токен
		if tokenStr == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		trimmed := strings.TrimPrefix(tokenStr, "Bearer ") //Обрезаем префикс
		//Генерируем токен
		token, err := jwt.ParseWithClaims(trimmed, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		//Проверяем на валидность
		if err != nil || token == nil || !token.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		//Вызываем хендлер
		ctx := context.WithValue(r.Context(), "user", token.Claims.(*Claims))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
