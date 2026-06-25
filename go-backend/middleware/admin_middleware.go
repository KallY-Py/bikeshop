package middleware

import (
	"net/http"
	"vue-go-backend/utils"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRole, ok := r.Context().Value(utils.UserRoleKey).(string)
		if !ok || userRole != "admin" {
			utils.RespondWithError(w, http.StatusForbidden, "Admin access required")
			return
		}
		next.ServeHTTP(w, r)
	})
}
