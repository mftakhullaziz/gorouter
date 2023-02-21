package apirest

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	model "gohttp/model"
	"gohttp/model/payload"
	"gohttp/service"
	helper2 "gohttp/utils/helper"
	"net/http"
	"strconv"
)

type CategoryRestImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryRestImpl(categoryService service.CategoryService) CategoryRest {
	return &CategoryRestImpl{CategoryService: categoryService}
}

func (rest CategoryRestImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	categoryCreateRequest := payload.CategoryRequest{}
	helper2.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := rest.CategoryService.Create(request.Context(), categoryCreateRequest)
	response := model.DefaultResponse{
		StatusCode: 201,
		Message:    "Create Category Success",
		IsSuccess:  true,
		Data:       categoryResponse,
	}

	helper2.WriteToResponseBody(writer, &response)
}

func (rest CategoryRestImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	categoryUpdateRequest := payload.CategoryUpdateRequest{}
	helper2.ReadFromRequestBody(request, &categoryUpdateRequest)

	catId := params.ByName("catId")
	id, err := strconv.Atoi(catId)
	helper2.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := rest.CategoryService.Update(request.Context(), categoryUpdateRequest)
	response := model.DefaultResponse{
		StatusCode: 201,
		Message:    "Update Category Success",
		IsSuccess:  true,
		Data:       categoryResponse,
	}

	helper2.WriteToResponseBody(writer, &response)
}

func (rest CategoryRestImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	catId := params.ByName("catId")
	id, err := strconv.Atoi(catId)
	helper2.PanicIfError(err)

	rest.CategoryService.Delete(request.Context(), id)

	response := model.DefaultResponse{
		StatusCode: 201,
		Message:    "Delete Category Success",
		IsSuccess:  true,
	}

	helper2.WriteToResponseBody(writer, &response)
}

func (rest CategoryRestImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	catId := params.ByName("catId")
	id, err := strconv.Atoi(catId)
	helper2.PanicIfError(err)

	categoryResponse := rest.CategoryService.FindById(request.Context(), id)
	fmt.Println(categoryResponse)

	response := model.DefaultResponse{
		StatusCode: 201,
		Message:    "Find Category By " + strconv.Itoa(id) + "Success",
		IsSuccess:  true,
		Data:       categoryResponse,
	}

	helper2.WriteToResponseBody(writer, &response)
}

func (rest CategoryRestImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	categoryResponses := rest.CategoryService.FindAll(request.Context())

	response := model.DefaultResponse{
		StatusCode: 201,
		Message:    "Find Category Success",
		IsSuccess:  true,
		Data:       categoryResponses,
	}

	helper2.WriteToResponseBody(writer, &response)
}
