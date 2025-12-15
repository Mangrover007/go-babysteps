package models 
 
type Response[T ResponseData] struct {
	Data  T
	Error string
}

type ResponseData interface {
	int64 | float64 | string
}

