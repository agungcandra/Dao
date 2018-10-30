package handler

import (
	"errors"
	"net/http"

	"github.com/agungcandra/dao/adapter"
	"github.com/julienschmidt/httprouter"
)

// CreateUser : handler for creating user
func CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	ctx := r.Context()
	select {
	case <-ctx.Done():
		return errors.New("Timeout")
	default:
	}

	adapter.UserAdapter.CreateUser()
}

func UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
