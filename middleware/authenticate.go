package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"ecommerce/util"
	"net/http"
	"strings"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			util.ErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessTokenArr := strings.Split(accessToken, " ")
		if len(accessTokenArr) != 2 {
			util.ErrorResponse(w, "Token structure: `Bearer Token`", http.StatusUnauthorized)
			return
		}

		if accessTokenArr[0] != "Bearer" {
			util.ErrorResponse(w, "Token need to start with `Bearer`", http.StatusUnauthorized)
			return
		}

		requestTokenArr := strings.Split(accessTokenArr[1], ".")
		if len(requestTokenArr) != 3 {
			util.ErrorResponse(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		jwtSecret := config.GetConfig().JwtSecret

		tokenHeader := requestTokenArr[0]
		tokenPayload := requestTokenArr[1]
		tokenSignature := requestTokenArr[2]

		byteMessage := []byte(tokenHeader + "." + tokenPayload)
		byteSecret := []byte(jwtSecret)

		hash := hmac.New(sha256.New, byteSecret)
		hash.Write(byteMessage)
		byteHashSignature := hash.Sum(nil)
		hashSignature := util.ConvertToBase64(byteHashSignature)

		if tokenSignature != hashSignature {
			util.ErrorResponse(w, "Signature mis-matched", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
