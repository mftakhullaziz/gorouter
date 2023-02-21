package exception

import (
	"gohttp/model"
	"gohttp/utils/helper"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	InternalServerError(w, r, err)
}

func InternalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	response := model.DefaultResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    "INTERNAL SERVER ERROR",
		IsSuccess:  false,
		Data:       err,
	}

	helper.WriteToResponseBody(w, response)
}
