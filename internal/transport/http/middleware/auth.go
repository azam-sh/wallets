package middleware

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"net/http"
	"strconv"
	"wallets/pkg/response"

	"github.com/gorilla/context"
)

func (mw middleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			resp response.Response
			sha  = sha1.New()
		)
		userIdStr := r.Header.Get("X-UserId")
		if userIdStr == "" {
			resp = response.BadRequest
			resp.WriteJSON(w)
			return
		}
		digest := r.Header.Get("X-Digest")
		if digest == "" {
			resp = response.BadRequest
			resp.WriteJSON(w)
			return
		}
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			resp = response.BadRequest
			resp.WriteJSON(w)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			resp = response.BadRequest
			resp.WriteJSON(w)
			return
		}
		data := userIdStr + string(body)
		sha.Write([]byte(data))
		expectedDigest := hex.EncodeToString(sha.Sum(nil))
		if !hmac.Equal([]byte(expectedDigest), []byte(digest)) {
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		context.Set(r, "userId", userId)
		next.ServeHTTP(w, r)
	})
}
