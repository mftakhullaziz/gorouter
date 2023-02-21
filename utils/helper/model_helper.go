package helper

import (
	"gohttp/model/domain"
	"gohttp/model/payload"
)

func ToCategoryResponse(category domain.Category) payload.CategoryResponse {
	return payload.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
		Desc: category.Desc,
	}
}

func ToCategoryResponses(categories []domain.Category) []payload.CategoryResponse {
	var categoryResponses []payload.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses,
			ToCategoryResponse(category))
	}
	return categoryResponses
}
