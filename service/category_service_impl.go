package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"gohttp/model/domain"
	"gohttp/model/payload"
	"gohttp/repository"
	helper2 "gohttp/utils/helper"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryServiceImpl(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository, DB: DB, Validate: validate}
}

func (service CategoryServiceImpl) Create(ctx context.Context, request payload.CategoryRequest) payload.CategoryResponse {
	//TODO implement me
	err := service.Validate.Struct(request)
	helper2.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper2.PanicIfError(err)
	defer helper2.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
		Desc: request.Desc,
	}

	category = service.CategoryRepository.Create(ctx, tx, category)
	return helper2.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Update(ctx context.Context, request payload.CategoryUpdateRequest) payload.CategoryResponse {
	//TODO implement me
	err := service.Validate.Struct(request)
	helper2.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper2.PanicIfError(err)
	defer helper2.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper2.PanicIfError(err)

	category.Name = request.Name
	category.Desc = request.Desc

	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper2.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Delete(ctx context.Context, catId int) {
	//TODO implement me
	tx, err := service.DB.Begin()
	helper2.PanicIfError(err)
	defer helper2.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, catId)
	helper2.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service CategoryServiceImpl) FindById(ctx context.Context, catId int) payload.CategoryResponse {
	//TODO implement me
	tx, err := service.DB.Begin()
	helper2.PanicIfError(err)
	defer helper2.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, catId)
	helper2.PanicIfError(err)

	return helper2.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) FindAll(ctx context.Context) []payload.CategoryResponse {
	//TODO implement me
	tx, err := service.DB.Begin()
	helper2.PanicIfError(err)
	defer helper2.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper2.ToCategoryResponses(categories)
}
