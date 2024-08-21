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
			mw.logger.Info("X-UserId not specified")
			resp = response.BadRequest
			resp.WriteJSON(w)
			return
		}
		digest := r.Header.Get("X-Digest")
		if digest == "" {
			mw.logger.Info("X-Digest not specified")
			resp = response.BadRequest
			resp.WriteJSON(w)
			return
		}
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			mw.logger.Error("could not convert userId into int: " + err.Error())
			resp = response.BadRequest
			resp.WriteJSON(w)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			mw.logger.Error("request body reading err: " + err.Error())
			resp = response.BadRequest
			resp.WriteJSON(w)
			return
		}
		data := userIdStr + string(body)
		sha.Write([]byte(data))
		expectedDigest := hex.EncodeToString(sha.Sum(nil))
		if !hmac.Equal([]byte(expectedDigest), []byte(digest)) {
			mw.logger.Info("digest hash and expected are not equal")
			resp = response.Unauthorized
			resp.WriteJSON(w)
			return
		}

		context.Set(r, "userId", userId)
		next.ServeHTTP(w, r)
	})
}
