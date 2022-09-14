package viewModels

import (
	"vehicle-tax/interfaces"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ListBasic struct {
	EndingBefore      *int64
	StartingAfter     *int64
	Limit             int64
	VehicleCategoryId *int64
	Validation        error
}

func NewListBasic(endingBefore *int64,
	startingAfter *int64,
	limit int64,
	vehicleCategoryId *int64) interfaces.ServiceValidation {

	valid := (&ListBasic{
		endingBefore,
		startingAfter,
		limit,
		vehicleCategoryId,
		nil,
	}).Validator()

	return &ListBasic{
		endingBefore,
		startingAfter,
		limit,
		vehicleCategoryId,
		valid,
	}

}

func (v *ListBasic) Validator() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Limit, validation.Min(0)))
}
