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
	Id          int64  `json:"id"`
	ShortName   string `json:"shortName"`
	Description string `json:"description"`
}

type VehicleTaxDto struct {
	VehicleTypeId              int64   `json:"vehicleTypeId"`
	VehicleTypeName            string  `json:"vehicleTypeName"`
	VehicleTypeDescription     string  `json:"vehicleTypeDescription"`
	VehicleCategoryId          int64   `json:"vehicleCategoryId"`
	VehicleCategoryName        string  `json:"vehicleCategoryName"`
	VehicleCategoryDescription string  `json:"vehicleCategoryDescription"`
	ImportDuty                 float64 `json:"importDuty"`
	Vat                        float64 `json:"vat"`
	Nhil                       float64 `json:"nhil"`
	GetfundLevy                float64 `json:"getfundLevy"`
	AuLevy                     float64 `json:"auLevy"`
	EcowasLevy                 float64 `json:"ecowasLevy"`
	EximLevy                   float64 `json:"eximLevy"`
	ExamLevy                   float64 `json:"examLevy"`
	ProcessingFee              float64 `json:"processingFee"`
	SpecialImportLevy          float64 `json:"specialImportLevy"`
}

type VehicleTax struct {
	Id                int64   `json:"id"`
	VehicleTypeId     int64   `json:"vehicleTypeId"`
	ImportDuty        float64 `json:"importDuty"`
	Vat               float64 `json:"vat"`
	Nhil              float64 `json:"nhil"`
	GetfundLevy       float64 `json:"getfundLevy"`
	AuLevy            float64 `json:"auLevy"`
	EcowasLevy        float64 `json:"ecowasLevy"`
	EximLevy          float64 `json:"eximLevy"`
	ExamLevy          float64 `json:"examLevy"`
	ProcessingFee     float64 `json:"processingFee"`
	SpecialImportLevy float64 `json:"specialImportLevy"`
}

type VehicleType struct {
	Id                int64  `json:"id"`
	VehicleCategoryId int64  `json:"vehicleCategoryd"`
	ShortName         string `json:"shortName"`
	Description       string `json:"description"`
}

func (tax VehicleTaxDto) CalculateDuty(costImportFreight float64) float64 {
	return costImportFreight * (tax.ImportDuty + tax.Nhil + tax.GetfundLevy + tax.AuLevy + tax.EcowasLevy +
		 tax.EximLevy + tax.ExamLevy + tax.ProcessingFee + tax.SpecialImportLevy) + tax.calculateVat(costImportFreight)

}

func (tax VehicleTaxDto) calculateVat(CIF float64) float64 {
	return (CIF + CIF * tax.ImportDuty + CIF * tax.Nhil + CIF * tax.GetfundLevy) * tax.Vat
}
