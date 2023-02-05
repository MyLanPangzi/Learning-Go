package solver

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type RemoteSolver struct {
	MathServerURL string
	client        *http.Client
}

func (r RemoteSolver) Resolve(ctx context.Context, expression string) (float64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.MathServerURL+"?expression="+url.QueryEscape(expression), nil)
	if err != nil {
		return 0, err
	}
	response, err := r.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}
	if response.StatusCode != http.StatusOK {
		return 0, errors.New(string(data))
	}
	result, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
