package logic

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type Logger interface {
	Log(ctx context.Context, string2 string)
}
type RequestDecorator func(r *http.Request) *http.Request

type UserLogic struct {
	RequestDecorator RequestDecorator
	Logger           Logger
	Remote           string
}

func (l UserLogic) Process(ctx context.Context, user string, data string) (string, error) {
	l.Logger.Log(ctx, fmt.Sprintf("starting process for %s with %s", user, data))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, l.Remote, nil)
	req.AddCookie(&http.Cookie{
		Name:   "user",
		Value:  user,
		Domain: "localhost:8080",
	})
	if err != nil {
		l.Logger.Log(ctx, fmt.Sprintf("error building remote %v", data))
		return "", err
	}
	req = l.RequestDecorator(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
