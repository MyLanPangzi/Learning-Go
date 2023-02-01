package controller

import (
	"Values/identity"
	"Values/logic"
	"net/http"
)

type UserController struct {
	logic.UserLogic
}

func (c UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	//c.Log(ctx, "world")
	user, ok := identity.UserFromContext(ctx)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := r.URL.Query().Get("data")
	result, err := c.UserLogic.Process(ctx, user, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(result))
}
