package bq

import (
	"errors"
	"fmt"
	"reflect"

	log "github.com/sirupsen/logrus"
)

func ToInterfaceSlice(slice interface{}) ([]interface{}, error) {

	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		msg := fmt.Sprintf("failed to cast interface into slice with 'given a non-slice type %s'", reflect.TypeOf(slice).String())
		log.Error(msg)
		return nil, errors.New(msg)
	}

	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret, nil
}
