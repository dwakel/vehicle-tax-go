package api

import (
	"log"
	"net/http"
	service "vehicle-tax/services"
)

type VehicleController struct {
	Logger     *log.Logger
	TaxService *service.IVehicleService
}

func NewVehicleController(logger *log.Logger,
	taxService *service.IVehicleService) *VehicleController {
	return &VehicleController{
		logger,
		taxService,
	}
}

// swagger:route GET templateMethod templateMethod
// Submits token
// Sample request:
//
//       ///     GET /?token=123456789
//
//responses:
//	200: Success
//	500: If the request processing fails due to an exception

// Returns a redirect to TemplateMethod
func (v *VehicleController) ListCategories(rw http.ResponseWriter, r *http.Request) {

	// res, err := (*v.TaxService).ListVehicleCategory(viewModels.NewListBasic())
	// if err != nil {
	// 	rw.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprint(rw, &errorResponse{Error: err})
	// 	return
	// }

	// rw.WriteHeader(http.StatusBadRequest)
	// fmt.Fprint(rw, &listResponse[models.VehicleCategory]{
	// 	Data:           res,
	// 	NextCursor:     nil,
	// 	PreviousCursor: nil,
	// 	TotalRecords:   int64(len(res)),
	// })
	// return
}

// Returns a redirect to TemplateMethod
func (v *VehicleController) ListTypes(rw http.ResponseWriter, r *http.Request) {

}

// Returns a redirect to TemplateMethod
func (v *VehicleController) ListTax(rw http.ResponseWriter, r *http.Request) {

}

// Returns a redirect to TemplateMethod
func (v *VehicleController) ListTaxSearchAndSort(rw http.ResponseWriter, r *http.Request) {

}

// Returns a redirect to TemplateMethods
func (v *VehicleController) CalculateDuty(rw http.ResponseWriter, r *http.Request) {

}
