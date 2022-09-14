package viewModels

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CalculateDuty struct {
	VehicleTypeId     int64
	CostImportFreight float64
	Validation        error
}

func NewCalculateDuty(
	VehicleTypeId int64,
	CostImportFreight float64) *CalculateDuty {

	valid := (&CalculateDuty{
		VehicleTypeId,
		CostImportFreight,
		nil,
	}).Validator()

	return &CalculateDuty{
		VehicleTypeId,
		CostImportFreight,
		valid,
	}

}

func (v *CalculateDuty) Validator() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.CostImportFreight, validation.Min(0)),
		validation.Field(&v.VehicleTypeId, validation.Min(1)))
}
