package pager

import (
	"reflect"
)

// PageModel page model
type PageModel struct {
	Count int64       `json:"count"`
	List  interface{} `json:"list"`
}

// PageCond page condition
type PageCond struct {
	Page     int
	PageSize int
}

// NewPageModel new page model
func NewPageModel(page, pageSize int, totalRecords int64, records interface{}) *PageModel {
	sliceValue := reflect.Indirect(reflect.ValueOf(records))
	if sliceValue.Kind() != reflect.Slice {
		panic("not a slice")
	}

	if totalRecords < 0 {
		totalRecords = 0
	}

	return &PageModel{
		Count: totalRecords,
		List:  records,
	}
}

// ValPageAndPageSize validate page pageSize
func ValPageAndPageSize(page, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return page, pageSize
}
