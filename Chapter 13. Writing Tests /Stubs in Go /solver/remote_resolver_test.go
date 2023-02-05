package solver

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoteSolver_Resolve(t *testing.T) {
	type info struct {
		expression string
		code       int
		body       string
	}
	var io info
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expression := r.URL.Query().Get("expression")
		if expression != io.expression {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid expression: " + expression))
			return
		}
		w.WriteHeader(io.code)
		w.Write([]byte(io.body))
	}))
	defer server.Close()
	solver := RemoteSolver{server.URL, server.Client()}
	data := []struct {
		name     string
		io       info
		expected float64
		errMsg   string
	}{
		{"2 + 2 * 10 = 22", info{"2 + 2 * 10", http.StatusOK, "22"}, 22, ""},
		{"( 2 + 2 ) * 10 = 40", info{"( 2 + 2 ) * 10", http.StatusOK, "40"}, 40, ""},
		{"( 2 + 2 * 10", info{"( 2 + 2 * 10", http.StatusBadRequest, "invalid expression: ( 2 + 2 * 10"}, 0, "invalid expression: ( 2 + 2 * 10"},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			io = d.io
			result, err := solver.Resolve(context.Background(), d.io.expression)
			if result != d.expected {
				t.Errorf("expected `%f`, got `%f`,", d.expected, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("error expected %s, got %s", d.errMsg, errMsg)
			}
		})

	}
}
