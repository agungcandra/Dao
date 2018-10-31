package handler

import (
	"net/http"

	"github.com/agungcandra/dao"
	"github.com/julienschmidt/httprouter"
)

// Handler global logic
type Handler struct {
	dao *dao.Dao
}

func (handler *Handler) Healtz(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
