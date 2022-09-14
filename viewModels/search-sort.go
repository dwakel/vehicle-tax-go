package viewModels

import (
	"math"
	"vehicle-tax/interfaces"
	"vehicle-tax/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type BasicSearchSort struct {
	SearchBy   map[string]any
	SortKey    string
	SortOrder  models.SortOrder
	Page       int32
	PerPage    int32
	Validation error
}

func NewBasicSearchSort(searchBy map[string]any,
	sortKey string,
	sortOrder models.SortOrder,
	page int32,
	perPage int32) interfaces.ServiceValidation {

	valid := (&BasicSearchSort{
		searchBy,
		sortKey,
		sortOrder,
		page,
		perPage,
		nil,
	}).Validator()

	return &BasicSearchSort{
		searchBy,
		sortKey,
		sortOrder,
		page,
		perPage,
		valid,
	}

}

func (v *BasicSearchSort) Validator() error {
	err := validation.ValidateStruct(&v,
		validation.Field(&v.PerPage, validation.Min(0), validation.Max(50)),
		validation.Field(&v.Page, validation.Min(0), validation.Max(math.MaxInt64)),
	)
	return err
}