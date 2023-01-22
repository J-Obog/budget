package resource

import (
	"encoding/json"
	"net/http"

	"github.com/J-Obog/paidoff/db"
)

type CategoryResource struct {
	categoryStore db.CategoryStore
}

func NewCategoryStore(categoryStore db.CategoryStore) *CategoryResource {
	return &CategoryResource{
		categoryStore: categoryStore,
	}
}

func (this *CategoryResource) GetCategory(req Request) *Response {
	categoryId := this.mustGetCategoryId(req)

	category, err := this.categoryStore.Get(categoryId)

	if err != nil {
		//return 500
	}

	if category == nil {
		//return 404
	}

	categoryResponse := this.toCategoryResponse(*category)
	responseBody, err := json.Marshal(categoryResponse)

	if err != nil {
		//return 500
	}

	return &Response{
		Body:   responseBody,
		Status: http.StatusOK,
	}
}

func (this *CategoryResource) GetCategories(req Request) *Response {
	categories, err := this.categoryStore.GetAll()

	categoryListResponse := this.toCategoryListResponse(categories)
	responseBody, err := json.Marshal(categoryListResponse)

	if err != nil {
		//return 500
	}

	return &Response{
		Body:   responseBody,
		Status: http.StatusOK,
	}
}

func (this *CategoryResource) toCategoryResponse(category db.Category) *CategoryResponse {
	return nil
}

func (this *CategoryResource) toCategoryListResponse(categories []db.Category) *CategoryListResponse {
	return nil
}

func (this *CategoryResource) mustGetCategoryId(req Request) string {
	return req.UrlParams["id"].(string)
}
