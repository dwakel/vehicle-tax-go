package api

import "vehicle-tax/models"

type SearchSortRequest struct {
	SearchBy  map[string]any
	SortKey   string
	SortOrder models.SortOrder
	Page      int32
	PerPage   int32
}
