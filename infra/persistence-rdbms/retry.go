package main

import (
	"fmt"
	"time"
)

type ParamsRetry struct {
	FNRetry   retryInterval
	NoRetries uint
}

func Retry[T any](params *ParamsRetry, f func() (*T, error)) (*T, error) {
	var errResult error

	for i := 0; i < int(params.NoRetries); i++ {
		result, errRun := f()
		if errRun == nil {
			return result,
				nil
		}

		if i == int(params.NoRetries) {
			errResult = fmt.Errorf(
				"not ready after %d retries:%w",
				params.NoRetries,
				errRun,
			)

			break
		}

		time.Sleep(
			params.FNRetry(uint(i)),
		)
	}

	return nil,
		errResult
}
