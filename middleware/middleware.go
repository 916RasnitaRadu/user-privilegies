package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"users-privi/model"
)

func extractUser(r *http.Request) (*model.User, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}
	defer r.Body.Close()

	var req model.UserRequest
	if err := json.Unmarshal(body, &req); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}

	user, ok := model.Users[req.UserID]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

// Permission middleware
func RequirePermission(permissions []string, next func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := extractUser(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		for _, p := range permissions {
			if !user.HasPermission(p) {
				http.Error(w, "Forbidden: missing permission "+p, http.StatusForbidden)
				return
			}
		}

		next(w, r)
	}
}

func IsUser(u model.User) bool {
	return u.HasPermission(model.ViewExperience) && u.HasPermission(model.Ialamuie)
}

func IsSupplier(u model.User) bool {
	return u.HasPermission(model.ViewExperience) && u.HasPermission(model.Ialamuie) && u.HasPermission(model.Dalamuie)
}

func OnlyUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := extractUser(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !IsUser(*user) {
			http.Error(w, "not enough permissions", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func OnlySupplier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := extractUser(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !IsSupplier(*user) {
			http.Error(w, "not enough permissions", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
