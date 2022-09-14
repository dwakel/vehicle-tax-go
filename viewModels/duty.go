package viewModels

import (
	"vehicle-tax/interfaces"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CalculateDuty struct {
	VehicleTypeId     int64
	CostImportFreight float64
	Validation        error
}

func NewCalculateDuty(endingBefore *int64,
	VehicleTypeId int64,
	CostImportFreight float64) interfaces.ServiceValidation {

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
