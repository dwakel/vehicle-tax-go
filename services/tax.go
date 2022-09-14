package service

import (
	"fmt"
	"log"
	"vehicle-tax/models"
	"vehicle-tax/repository"
	"vehicle-tax/viewModels"
)

type TaxService struct {
	logger  *log.Logger
	taxRepo repository.IVehicleRepository
}

func NewTaxService(logger *log.Logger, taxRepo repository.IVehicleRepository) IVehicleService {
	return &TaxService{logger, taxRepo}
}

func (t *TaxService) ListVehicleCategory(query viewModels.ListBasic) ([]models.VehicleCategory, error) {
	if query.Validation != nil {
		fmt.Fprint(t.logger.Writer(), "Validation failed: ", query.Validation)
	}

	taxes, err := t.taxRepo.ListVehicleCategories(query.EndingBefore, query.StartingAfter, query.Limit)

	if err != nil {
		return nil, err
	}

	t.logger.Println("Vehicle Categories Listed Succesfully")
	return taxes, nil
}

func (t *TaxService) ListVehicleType(query viewModels.ListBasic) ([]models.VehicleType, error) {
	if query.Validation != nil {
		fmt.Fprint(t.logger.Writer(), "Validation failed: ", query.Validation)
	}

	taxes, err := t.taxRepo.ListVehicleType(query.EndingBefore, query.StartingAfter, query.VehicleCategoryId, query.Limit)

	if err != nil {
		return nil, err
	}

	t.logger.Println("Vehicle Type Listed Succesfully")
	return taxes, nil
}

func (t *TaxService) ListVehicleTax(query viewModels.ListBasic) ([]models.VehicleTax, error) {
	if query.Validation != nil {
		fmt.Fprint(t.logger.Writer(), "Validation failed: ", query.Validation)
	}

	taxes, err := t.taxRepo.ListVehicleTax(query.EndingBefore, query.StartingAfter, query.VehicleCategoryId, query.Limit)

	if err != nil {
		return nil, err
	}

	t.logger.Println("Vehicle Tax Listed Succesfully")
	return taxes, nil
}

func (t *TaxService) ListVehicleTaxSearchSort(query viewModels.BasicSearchSort) ([]models.VehicleTaxDto, error) {
	if query.Validation != nil {
		fmt.Fprint(t.logger.Writer(), "Validation failed: ", query.Validation)
	}

	duty, err := t.taxRepo.ListVehicleTaxSearchAndSort(query.SearchBy, query.SortKey, query.SortOrder)

	if err != nil {
		return nil, err
	}

	var skip, take int32 = 0, query.PerPage
	if query.PerPage >= 0 {
		skip = (query.Page - 1) * query.PerPage
	} else {
		take = int32(len(duty))
	}

	t.logger.Println("Vehicle Tax Listed Succesfully")
	return duty[skip : skip+take], nil
}

func (t *TaxService) CalculateDuty(query viewModels.CalculateDuty) (float64, error) {
	if query.Validation != nil {
		fmt.Fprint(t.logger.Writer(), "Validation failed: ", query.Validation)
	}

	taxInfo, err := (t.taxRepo.FetchVehicleTaxByTypeId(query.VehicleTypeId))

	if err != nil {
		return *new(float64), err
	}

	return taxInfo.CalculateDuty(query.CostImportFreight), nil
}
