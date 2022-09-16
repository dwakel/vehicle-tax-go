package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"vehicle-tax/models"

	"gorm.io/gorm"
)

type TaxRepository struct {
	logger *log.Logger
	repo   *gorm.DB
}

func NewTaxRepo(logger *log.Logger, repo *gorm.DB) IVehicleRepository {
	return &TaxRepository{logger, repo}
}

func (this *TaxRepository) ListVehicleCategories(endingBefore *int64, startingAfter *int64, limit int64) ([]models.VehicleCategory, error) {
	var query string
	var result []models.VehicleCategory
	if endingBefore != nil {
		query = `SELECT
					a.*,
					LEAD(a.id) OVER(ORDER BY a.id DESC) prev,
					FIRST_VALUE(a.id) OVER(ORDER BY a.id DESC) ""next""
					FROM public.vehicle_category a
					WHERE a.id < @endingBefore
					ORDER BY a.id desc ` + appendLimit(limit) + ";"
	} else if startingAfter != nil {
		query = `SELECT * FROM (
						SELECT
						a.*,
						LEAD(a.id) OVER(ORDER BY a.id DESC) prev,
						FIRST_VALUE(a.id) OVER(ORDER BY a.id DESC) ""next""
					FROM public.vehicle_category a
					WHERE a.id > @StartingAfter
					ORDER BY a.id desc` + appendLimit(limit) + `) list ORDER BY id desc;`
	} else {
		query = `SELECT
					a.*,
					LEAD(a.id) OVER(ORDER BY a.id DESC) prev,
					null "next"
					FROM public.vehicle_category a
					ORDER BY a.id desc ` + appendLimit(limit) + ";"
	}

	err := this.repo.Raw(query, map[string]interface{}{"startingAfter": startingAfter, "endingBefore": endingBefore, "limit": limit}).Scan(&result).Error
	//err := this.repo.Select("description", "is_active").Create(&model).Error
	if err != nil {
		this.logger.Println("failed to fetch vehicle category")
		this.logger.Println(err)
		return nil, errors.New("failed to fetch vehicle category")
	}
	this.logger.Println("Successfully fetched vehicle category")
	return result, nil
}

func (this *TaxRepository) ListVehicleType(vehicleCategoryId *int64, endingBefore *int64, startingAfter *int64, limit int64) ([]models.VehicleType, error) {

	var query strings.Builder
	var result []models.VehicleType
	if endingBefore != nil {
		query.WriteString(`SELECT
							a.id, a.vehicle_category_id, a.short_name, a.description,
							LEAD(a.id) OVER(ORDER BY a.id DESC) prev,
							FIRST_VALUE(a.id) OVER(ORDER BY a.id DESC) ""next""
							FROM public.vehicle_type a
							WHERE a.id < @endingBefore`)
		if vehicleCategoryId != nil {
			query.WriteString(" AND a.vehicle_category_id = @vehicleCategoryId ")
		}
		query.WriteString(`ORDER BY a.id desc ` + appendLimit(limit))

	} else if startingAfter != nil {
		query.WriteString(`SELECT * FROM (
								SELECT
								a.*,
								LEAD(a.id) OVER(ORDER BY a.id DESC) prev,
								FIRST_VALUE(a.id) OVER(ORDER BY a.id DESC) "next"
							FROM public.vehicle_type a
							WHERE a.id > @StartingAfter `)
		if vehicleCategoryId != nil {
			query.WriteString("AND a.vehicle_category_id = @vehicleCategoryId ")
		}
		query.WriteString(`ORDER BY a.id desc ` + appendLimit(limit) + `) list ORDER BY id desc`)
	} else {
		query.WriteString(`SELECT
							a.*,
							LEAD(a.id) OVER(ORDER BY a.id DESC) prev,
							null "next"
							FROM public.vehicle_type a `)
		if vehicleCategoryId != nil {
			query.WriteString("WHERE a.vehicle_category_id = @vehicleCategoryId ")
		}
		query.WriteString(`ORDER BY a.id desc ` + appendLimit(limit) + ";")
	}

	err := this.repo.Raw(query.String(), map[string]interface{}{"vehicleCategoryId": vehicleCategoryId, "startingAfter": startingAfter, "endingBefore": endingBefore, "limit": limit}).Scan(&result).Error
	//err := this.repo.Select("description", "is_active").Create(&model).Error
	if err != nil {
		this.logger.Println("failed to fetch vehicle type")
		this.logger.Println(err)
		return nil, errors.New("failed to fetch vehicle type")
	}
	this.logger.Println("Successfully fetched vehicle type")
	return result, nil
}

func (this *TaxRepository) ListVehicleTax(vehicleCategoryId *int64, endingBefore *int64, startingAfter *int64, limit int64) ([]models.VehicleTax, error) {
	var query strings.Builder
	var result []models.VehicleTax
	if endingBefore != nil {
		query.WriteString(`SELECT
							a.*,
							LEAD(a.id) OVER(ORDER BY a.id DESC) prev,
							FIRST_VALUE(a.id) OVER(ORDER BY a.id DESC) "next"
							FROM public.vehicle_tax a `)
		if vehicleCategoryId != nil {
			query.WriteString("LEFT JOIN public.vehicle_type b ON b.id = a.vehicle_type_id ")
		}
		query.WriteString(`WHERE a.id < @endingBefore `)
		if vehicleCategoryId != nil {
			query.WriteString("AND b.vehicle_category_id = @vehicleCategoryId ")
		}
		query.WriteString(`ORDER BY a.id desc ` + appendLimit(limit))

	} else if startingAfter != nil {
		query.WriteString(`SELECT * FROM (
								SELECT
								a.*,
								LEAD(a.id) OVER(ORDER BY a.id DESC) prev,
								FIRST_VALUE(a.id) OVER(ORDER BY a.id DESC) "next"
							FROM public.vehicle_tax a `)
		if vehicleCategoryId != nil {
			query.WriteString("LEFT JOIN public.vehicle_type b ON b.id = a.vehicle_type_id ")
		}
		query.WriteString(`WHERE a.id > @StartingAfter `)
		if vehicleCategoryId != nil {
			query.WriteString("AND b.vehicle_category_id = @vehicleCategoryId ")
		}
		query.WriteString(`ORDER BY a.id desc` + appendLimit(limit) + `) list ORDER BY id desc `)
	} else {
		query.WriteString(`SELECT
							a.*,
							LEAD(a.id) OVER(ORDER BY a.id DESC) prev,
							null "next"
							FROM public.vehicle_tax a `)
		if vehicleCategoryId != nil {
			query.WriteString("LEFT JOIN public.vehicle_type b ON b.id = a.vehicle_type_id ")
			query.WriteString("AND b.vehicle_category_id = @vehicleCategoryId ")
		}
		query.WriteString(`ORDER BY a.id desc ` + appendLimit(limit) + ";")
	}

	err := this.repo.Raw(query.String(), map[string]interface{}{"vehicleCategoryId": vehicleCategoryId, "startingAfter": startingAfter, "endingBefore": endingBefore, "limit": limit}).Scan(&result).Error
	//err := this.repo.Select("description", "is_active").Create(&model).Error
	if err != nil {
		this.logger.Println("failed to fetch vehicle type")
		this.logger.Println(err)
		return nil, errors.New("failed to fetch vehicle type")
	}
	this.logger.Println("Successfully fetched vehicle type")
	return result, nil
}

func (this *TaxRepository) ListVehicleTaxSearchAndSort(searchBy map[string]any, sortKey string, sortOrder models.SortOrder) ([]models.VehicleTaxDto, error) {
	var query string
	var result []models.VehicleTaxDto

	var search strings.Builder
	var sort string

	switch sortKey {
	case "importduty":
		sort = fmt.Sprintf("a.import_duty %v", sortOrder.String())
	case "vat":
		sort = fmt.Sprintf("a.vat %v", sortOrder.String())
	case "nhil":
		sort = fmt.Sprintf("a.nhil %v", sortOrder.String())
	case "getfundlevy":
		sort = fmt.Sprintf("a.getfund_levy %v", sortOrder.String())
	case "aulevy":
		sort = fmt.Sprintf("a.au_levy %v", sortOrder.String())
	case "ecowaslevy":
		sort = fmt.Sprintf("a.ecowas_levy %v", sortOrder.String())
	case "eximlevy":
		sort = fmt.Sprintf("a.exim_levy %v", sortOrder.String())
	case "examlevy":
		sort = fmt.Sprintf("a.exam_levy %v", sortOrder.String())
	case "processingfee":
		sort = fmt.Sprintf("a.processing_fee %v", sortOrder.String())
	case "specialimportlevy":
		sort = fmt.Sprintf("a.special_import_levy %v", sortOrder.String())
	default:
		sort = fmt.Sprintf("a.vehicle_type_id %v", sortOrder.String())
	}

	for k, v := range searchBy {
		fmt.Println("HERE: ", k)
		if strings.ToLower(k) == "importduty" {
			search.WriteString(fmt.Sprint("AND a.import_duty = %v ", v))
		}
		if strings.ToLower(k) == "vat" {
			search.WriteString(fmt.Sprint("AND a.vat = %v ", v))
		}
		if strings.ToLower(k) == "nhil" {
			search.WriteString(fmt.Sprint("AND a.nhil = %v ", v))
		}
		if strings.ToLower(k) == "getfundlevy" {
			search.WriteString(fmt.Sprint("AND a.getfund_levy = %v ", v))
		}
		if strings.ToLower(k) == "aulevy" {
			search.WriteString(fmt.Sprint("AND a.au_levy = %v ", v))
		}
		if strings.ToLower(k) == "ecowaslevy" {
			search.WriteString(fmt.Sprint("AND a.ecowas_levy = %v ", v))
		}
		if strings.ToLower(k) == "eximlevy" {
			search.WriteString(fmt.Sprint("AND a.exim_levy = %v ", v))
		}
		if strings.ToLower(k) == "examlevy" {
			search.WriteString(fmt.Sprint("AND a.exam_levy = %v ", v))
		}
		if strings.ToLower(k) == "processingfee" {
			search.WriteString(fmt.Sprint("AND a.processing_fee = %v ", v))
		}
		if strings.ToLower(k) == "specialimportlevy" {
			search.WriteString(fmt.Sprint("AND a.special_import_levy = %v ", v))
		}
		if strings.ToLower(k) == "categoryname" {
			search.WriteString("AND LOWER(c.short_name) LIKE '%" + fmt.Sprint(v) + "%' ")
		}
		if strings.ToLower(k) == "typename" {
			search.WriteString("AND LOWER(b.short_name) LIKE '%" + fmt.Sprint(v) + "%' ")
		}
		if strings.ToLower(k) == "categoryname" {
			search.WriteString("AND LOWER(c.short_name) LIKE '%" + fmt.Sprint(v) + "%' ")
		}
		if strings.ToLower(k) == "typedescription" {
			search.WriteString("AND LOWER(b.description) LIKE '%" + fmt.Sprint(v) + "%' ")
		}
	}

	query = `SELECT
					a.*, b.short_name vehicle_type_name, b.description vehicle_type_description,
					c.short_name vehicle_category_name, c.description vehicle_category_description
				FROM public.vehicle_tax a
				LEFT JOIN public.vehicle_type b ON b.id = a.vehicle_type_id
				LEFT JOIN public.vehicle_category c ON c.id = b.vehicle_category_id
				WHERE a.id <> 0 ` +
		search.String() +
		`ORDER BY ` + sort

	err := this.repo.Raw(query).Scan(&result).Error
	//err := this.repo.Select("description", "is_active").Create(&model).Error
	if err != nil {
		this.logger.Println("failed to fetch vehicle tax dto")
		this.logger.Println(err)
		return nil, errors.New("failed to fetch vehicle tax dto")
	}
	this.logger.Println("Successfully fetched vehicle tax dto")
	return result, nil
}

func (this *TaxRepository) FetchVehicleTaxByTypeId(vehicleTypeId int64) (*models.VehicleTaxDto, error) {
	var app models.VehicleTaxDto
	command := `SELECT
					a.*, b.short_name vehicle_type_name, b.description vehicle_type_description,
					c.short_name vehicle_category_name, c.description vehicle_category_description
				FROM public.vehicle_tax a
				LEFT JOIN public.vehicle_type b ON b.id = a.vehicle_type_id
				LEFT JOIN public.vehicle_category c ON c.id = b.vehicle_category_id
				WHERE b.id = @vehicleTypeId ;`
	err := this.repo.Raw(command, sql.Named("vehicleTypeId", vehicleTypeId)).Scan(&app).Error
	if err != nil {
		this.logger.Println("failed to fetch tax information")
		this.logger.Println(err)
		return nil, errors.New("failed to fetch tax information")
	}
	this.logger.Println("Successfully fetched application")
	return &app, nil
}

func appendLimit(limit int64) string {
	if limit >= 0 {
		return "Limit @limit"
	}
	return ""
}
