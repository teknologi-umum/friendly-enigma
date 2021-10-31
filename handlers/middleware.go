package handlers

import (
	"context"
	"net/http"
	"refrigerator/business"
	"refrigerator/packages/jwt"
	"strings"
)

func (d *Deps) HasAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		jwtHeader := strings.Join(r.Header["Authorization"], " ")

		secret := []byte("askdij a bsdfiWBNEOIFUwbfnioWUBWE")
		u, err := jwt.VerifyJWT(secret, jwtHeader)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ngga gitu woi"))
			return
		}

		switch r.Method {
		case http.MethodGet:
			ok := business.HasAccess(u.Permission, business.READ)
			if !ok {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("gaboleh masuk, bayar jatah preman dulu"))
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, "user", u)))

		case http.MethodPost:
			ok := business.HasAccess(u.Permission, business.CREATE)
			if !ok {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("gaboleh masuk, bayar jatah preman dulu"))
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, "user", u)))
		case http.MethodPatch:
			ok := business.HasAccess(u.Permission, business.UPDATE)
			if !ok {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("gaboleh masuk, bayar jatah preman dulu"))
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, "user", u)))
		case http.MethodDelete:
			ok := business.HasAccess(u.Permission, business.DELETE)
			if !ok {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("gaboleh masuk, bayar jatah preman dulu"))
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, "user", u)))
		default:
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("gaboleh masuk, bayar jatah preman dulu"))
		}
	})
}
