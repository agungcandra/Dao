package handler

import (
	"errors"
	"net/http"

	"github.com/agungcandra/dao"
	"github.com/agungcandra/dao/entity"
	us "github.com/agungcandra/dao/service/user"
	"github.com/julienschmidt/httprouter"
)

// UserHandler hold user
type UserHandler struct {
	dao *dao.Dao
}

func NewUserHandler(dao *dao.Dao) *UserHandler {
	return &UserHandler{
		dao: dao,
	}
}

// CreateUser : handler for creating user
func (uHandler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	ctx := r.Context()
	select {
	case <-ctx.Done():
		return errors.New("Timeout")
	default:
	}

	user := entity.User{
		Username: "ASD",
		Password: "ASD",
		Name:     "ASD",
	}

	_, err := us.CreateUser(uHandler.dao.Database(), user)
	return err
}
