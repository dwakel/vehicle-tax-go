package models

type SortOrder string

const (
	ASC  = "asc"
	DESC = "desc"
)

func (s SortOrder) String() string {
	switch s {
	case ASC:
		return ASC
	case DESC:
		return DESC
	default:
		return ASC
	}
}

type VehicleCategory struct {
	Id          int64  `json:"id" pg:"id"`
	ShortName   string `json:"shortName" pg:"short_name"`
	Description string `json:"description" pg:"description"`
}

type VehicleTaxDto struct {
	VehicleTypeId              int64   `json:"vehicleTypeId" pg:"vehicle_type_id"`
	VehicleTypeName            string  `json:"vehicleTypeName" pg:"vehicle_type_name"`
	VehicleTypeDescription     string  `json:"vehicleTypeDescription" pg:"vehicle_type_description"`
	VehicleCategoryId          int64   `json:"vehicleCategoryId" pg:"vehicle_category_id"`
	VehicleCategoryName        string  `json:"vehicleCategoryName" pg:"vehicle_category_name"`
	VehicleCategoryDescription string  `json:"vehicleCategoryDescription" pg:"vehicle_category_description"`
	ImportDuty                 float64 `json:"importDuty" pg:"import_duty"`
	Vat                        float64 `json:"vat" pg:"vat"`
	Nhil                       float64 `json:"nhil" pg:"nhil"`
	GetfundLevy                float64 `json:"getfundLevy" pg:"getfund_levy"`
	AuLevy                     float64 `json:"auLevy" pg:"au_levy"`
	EcowasLevy                 float64 `json:"ecowasLevy" pg:"ecowas_levy"`
	EximLevy                   float64 `json:"eximLevy" pg:"exim_levy"`
	ExamLevy                   float64 `json:"examLevy" pg:"exam_levy"`
	ProcessingFee              float64 `json:"processingFee" pg:"processing_fee"`
	SpecialImportLevy          float64 `json:"specialImportLevy" pg:"special_import_levy"`
}

type VehicleTax struct {
	Id                int64   `json:"id" pg:"id"`
	VehicleTypeId     int64   `json:"vehicleTypeId" pg:"vehicle_type_id"`
	ImportDuty        float64 `json:"importDuty" pg:"import_duty"`
	Vat               float64 `json:"vat" pg:"vat"`
	Nhil              float64 `json:"nhil" pg:"nhil"`
	GetfundLevy       float64 `json:"getfundLevy" pg:"getfund_levy"`
	AuLevy            float64 `json:"auLevy" pg:"au_levy"`
	EcowasLevy        float64 `json:"ecowasLevy" pg:"ecowas_levy"`
	EximLevy          float64 `json:"eximLevy" pg:"exim_levy"`
	ExamLevy          float64 `json:"examLevy" pg:"exam_levy"`
	ProcessingFee     float64 `json:"processingFee" pg:"processing_fee"`
	SpecialImportLevy float64 `json:"specialImportLevy" pg:"special_import_levy"`
}

type VehicleType struct {
	Id                int64  `json:"id" pg:"id"`
	VehicleCategoryId int64  `json:"vehicleCategoryid" pg:"vehicle_category_id"`
	ShortName         string `json:"shortName" pg:"short_name"`
	Description       string `json:"description" pg:"description"`
}

func (tax VehicleTaxDto) CalculateDuty(costImportFreight float64) float64 {
	return costImportFreight*(tax.ImportDuty+tax.Nhil+tax.GetfundLevy+tax.AuLevy+tax.EcowasLevy+
		tax.EximLevy+tax.ExamLevy+tax.ProcessingFee+tax.SpecialImportLevy) + tax.calculateVat(costImportFreight)

}

func (tax VehicleTaxDto) calculateVat(CIF float64) float64 {
	return (CIF + CIF*tax.ImportDuty + CIF*tax.Nhil + CIF*tax.GetfundLevy) * tax.Vat
}
