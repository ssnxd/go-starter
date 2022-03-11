package util

import (
	"reflect"
	"time"
)

// NOTE: Server accepts time in the format ISO-8601 (UTC)
// server will try to convert any date to this format so make sure you provide correct format
func ValidateDate(field reflect.Value) interface{} {

	time, err := time.Parse(time.RFC3339, field.String())
	
	if err != nil {
		return err
	}

	return time
}
