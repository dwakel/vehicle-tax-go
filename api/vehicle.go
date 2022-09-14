package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"vehicle-tax/models"
	service "vehicle-tax/services"
	"vehicle-tax/viewModels"

	"github.com/gorilla/mux"
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

	var endingBefore, startingAfter *int64
	var limit int64 = 10

	end, err := strconv.Atoi(r.URL.Query().Get("endingBefore"))
	if err == nil {
		*endingBefore = int64(end)
	}
	start, err := strconv.Atoi(r.URL.Query().Get("startingAfter"))
	if err == nil {
		*startingAfter = int64(start)
	}
	lim, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err == nil {
		limit = int64(lim)
	}

	res, err := (*v.TaxService).ListVehicleCategory(*viewModels.NewListBasic(endingBefore, startingAfter, limit, nil))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, &errorResponse{Error: err})
		return
	}

	rw.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(rw, &listResponse[models.VehicleCategory]{
		Data:           res,
		NextCursor:     nil,
		PreviousCursor: nil,
		TotalRecords:   int64(len(res)),
	})
}

func (v *VehicleController) ListTypes(rw http.ResponseWriter, r *http.Request) {
	var endingBefore, startingAfter, categoryId *int64
	var limit int64 = 10

	end, err := strconv.Atoi(r.URL.Query().Get("endingBefore"))
	if err == nil {
		*endingBefore = int64(end)
	}
	start, err := strconv.Atoi(r.URL.Query().Get("startingAfter"))
	if err == nil {
		*startingAfter = int64(start)
	}
	cat, err := strconv.Atoi(r.URL.Query().Get("categoryId"))
	if err == nil {
		*categoryId = int64(cat)
	}
	lim, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err == nil {
		limit = int64(lim)
	}

	res, err := (*v.TaxService).ListVehicleType(*viewModels.NewListBasic(endingBefore, startingAfter, limit, categoryId))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, &errorResponse{Error: err})
		return
	}

	rw.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(rw, &listResponse[models.VehicleType]{
		Data:           res,
		NextCursor:     nil,
		PreviousCursor: nil,
		TotalRecords:   int64(len(res)),
	})
}

func (v *VehicleController) ListTax(rw http.ResponseWriter, r *http.Request) {
	var endingBefore, startingAfter *int64
	var limit int64 = 10

	end, err := strconv.Atoi(r.URL.Query().Get("endingBefore"))
	if err == nil {
		*endingBefore = int64(end)
	}
	start, err := strconv.Atoi(r.URL.Query().Get("startingAfter"))
	if err == nil {
		*startingAfter = int64(start)
	}
	lim, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err == nil {
		limit = int64(lim)
	}

	res, err := (*v.TaxService).ListVehicleTax(*viewModels.NewListBasic(endingBefore, startingAfter, limit, nil))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, &errorResponse{Error: err})
		return
	}

	rw.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(rw, &listResponse[models.VehicleTax]{
		Data:           res,
		NextCursor:     nil,
		PreviousCursor: nil,
		TotalRecords:   int64(len(res)),
	})

}

func (v *VehicleController) ListTaxSearchAndSort(rw http.ResponseWriter, r *http.Request) {
	var req SearchSortRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, &errorResponse{Error: err})
		return
	}

	res, err := (*v.TaxService).ListVehicleTaxSearchSort(*viewModels.NewBasicSearchSort(req.SearchBy, req.SortKey, req.SortOrder, req.Page, req.PerPage))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, &errorResponse{Error: err})
		return
	}

	rw.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(rw, &listResponse[models.VehicleTaxDto]{
		Data:           res,
		NextCursor:     nil,
		PreviousCursor: nil,
		TotalRecords:   int64(len(res)),
	})

}

func (v *VehicleController) CalculateDuty(rw http.ResponseWriter, r *http.Request) {
	var costImportFreight float64
	vehicleTypeId, err := strconv.Atoi(mux.Vars(r)["vehicleTypeId"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, &errorResponse{Error: errors.New("invalid request id")})
		return
	}
	cif, err := strconv.ParseFloat(r.URL.Query().Get("costImportFreight"), 64)
	if err == nil {
		costImportFreight = float64(cif)
	}
	res, err := (*v.TaxService).CalculateDuty(*viewModels.NewCalculateDuty(int64(vehicleTypeId), costImportFreight))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, &errorResponse{Error: err})
		return
	}

	rw.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(rw, res)

}
