package service

import (
	"context"
	"gohttp/model/payload"
)

type CategoryService interface {
	Create(ctx context.Context, request payload.CategoryRequest) payload.CategoryResponse
	Update(ctx context.Context, request payload.CategoryUpdateRequest) payload.CategoryResponse
	Delete(ctx context.Context, catId int)
	FindById(ctx context.Context, catId int) payload.CategoryResponse
	FindAll(ctx context.Context) []payload.CategoryResponse
}
