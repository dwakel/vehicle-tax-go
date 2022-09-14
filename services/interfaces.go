package service

import (
	"vehicle-tax/models"
	"vehicle-tax/viewModels"
)

type IVehicleService interface {
	ListVehicleCategory(query viewModels.ListBasic) ([]models.VehicleCategory, error)
	ListVehicleType(query viewModels.ListBasic) ([]models.VehicleType, error)
	ListVehicleTax(query viewModels.ListBasic) ([]models.VehicleTax, error)
	ListVehicleTaxSearchSort(query viewModels.BasicSearchSort) ([]models.VehicleTaxDto, error)
	CalculateDuty(query viewModels.CalculateDuty) (float64, error)
}
