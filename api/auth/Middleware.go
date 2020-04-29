package auth

import (
	"api/jwtoken"
	"context"
	"database/sql"
	"net/http"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var clientCtxKey = &contextKey{"client"}

type contextKey struct {
	name string
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenstring := ""
			if r != nil {
				tokenstring = r.Header.Get("Authorization")
			}

			// Allow unauthenticated users in
			if tokenstring == "" {
				next.ServeHTTP(w, r)
				return
			}

			username, err := jwtoken.VerifyToken(tokenstring)
			if err != nil {
				//http.Error(w, "Invalid cookie", http.StatusForbidden)
				next.ServeHTTP(w, r)
				return
			}

			// get the user from the database
			//user := getUserByID(db, userId)

			// put it in context
			ctx := context.WithValue(r.Context(), clientCtxKey, username)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) string {
	raw, _ := ctx.Value(clientCtxKey).(string)
	return raw
}
