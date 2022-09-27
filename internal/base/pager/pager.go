package pager

import (
	"errors"
	"reflect"

	"xorm.io/xorm"
)

// Help xorm page helper
func Help(page, pageSize int, rowsSlicePtr interface{}, rowElement interface{}, session *xorm.Session) (total int64, err error) {
	page, pageSize = ValPageAndPageSize(page, pageSize)

	sliceValue := reflect.Indirect(reflect.ValueOf(rowsSlicePtr))
	if sliceValue.Kind() != reflect.Slice {
		return 0, errors.New("not a slice")
	}

	startNum := (page - 1) * pageSize
	return session.Limit(pageSize, startNum).FindAndCount(rowsSlicePtr, rowElement)
}
