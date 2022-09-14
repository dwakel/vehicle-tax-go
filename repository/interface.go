package repository

import "vehicle-tax/models"

type IVehicleRepository interface {
	ListVehicleCategories(endingBefore *int64, startingAfter *int64, limit int64) ([]models.VehicleCategory, error)
	ListVehicleType(vehicleCategoryId *int64, endingBefore *int64, startingAfter *int64, limit int64) ([]models.VehicleType, error)
	ListVehicleTax(vehicleCategoryId *int64, endingBefore *int64, startingAfter *int64, limit int64) ([]models.VehicleTax, error)
	ListVehicleTaxSearchAndSort(searchBy map[string]any, sortKey string, sortOrder models.SortOrder) ([]models.VehicleTaxDto, error)
	FetchVehicleTaxByTypeId(vehicleTypeId int64) (*models.VehicleTaxDto, error)
}
