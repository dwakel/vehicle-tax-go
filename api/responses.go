package api

type errorResponse struct {
	Error any
}

type listResponse[T any] struct {
	Data           []T
	NextCursor     *int64
	PreviousCursor *int64
	TotalRecords   int64
}
