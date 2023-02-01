package tracker

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type guidKey int

const key guidKey = 1

func contextWithGUID(ctx context.Context, guid string) context.Context {
	return context.WithValue(ctx, key, guid)
}
func GuidFromContext(ctx context.Context) (string, bool) {
	guid, ok := ctx.Value(key).(string)
	return guid, ok
}
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if guid := r.Header.Get("X-GUID"); guid != "" {
			ctx = contextWithGUID(ctx, guid)
		} else {
			ctx = contextWithGUID(ctx, uuid.New().String())
		}
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

type Logger struct {
}

func (l Logger) Log(ctx context.Context, message string) {
	if guid, ok := GuidFromContext(ctx); ok {
		message = fmt.Sprintf("GUID %s - %s", guid, message)
	}
	log.Println(message)
}

func Request(r *http.Request) *http.Request {
	ctx := r.Context()
	if guid, ok := GuidFromContext(ctx); ok {
		r.Header.Add("X-GUID", guid)
	}
	return r
}
