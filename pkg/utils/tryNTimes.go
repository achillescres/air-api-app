package utils

import (
	"reflect"
	"time"
)

func TryNTimes[returnType any](n int, f func() (returnType, error)) (returnType, error) {
	return TryNTimesWaiting[returnType](n, 0, f)
}

func TryNTimesWaiting[returnType any](n int, waitingDuration time.Duration, f func() (returnType, error)) (returnType, error) {
	reflect.ValueOf()
	var err error
	var res any
	for i := 0; i < n; i++ {
		res, err = f()
		if err == nil {
			return res, nil
		}
		res, _ = res.
			reflect.TypeOf(returnType)
		time.Sleep(waitingDuration)
	}

	return nil, err
}
