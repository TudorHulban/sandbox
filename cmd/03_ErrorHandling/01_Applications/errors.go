package main

import (
	"fmt"
	"time"
)

type ErrService struct {
	Issue error

	Caller     string
	MethodName string

	NanosecondsDuration int64
	WithTime            bool
}

const areaErrService = "Service"

func (e ErrService) Error() string {
	var res [4]string

	if e.WithTime {
		res[0] = fmt.Sprintf("\nArea: %s [%d] - duration nanoseconds: %d",
			areaErrService, time.Now().Unix(), e.NanosecondsDuration)
	} else {
		res[0] = fmt.Sprintf("\nArea: %s", areaErrService)
	}

	res[1] = fmt.Sprintf("Caller: %s", e.Caller)
	res[2] = fmt.Sprintf("Method Name: %s", e.MethodName)
	res[3] = fmt.Sprintf("Issue: %s", e.Issue.Error())

	return res[0] + "\n" + res[1] + "\n" + res[2] + "\n" + res[3]
}

type ErrDomain struct {
	Issue error

	Caller     string
	MethodName string

	WithTime bool
}

const areaErrDomain = "Domain"

func (e ErrDomain) Error() string {
	var res [4]string

	if e.WithTime {
		res[0] = fmt.Sprintf("\nArea: %s [%d]", areaErrDomain, time.Now().Unix())
	} else {
		res[0] = fmt.Sprintf("\nArea: %s", areaErrDomain)
	}

	res[1] = fmt.Sprintf("Caller: %s", e.Caller)
	res[2] = fmt.Sprintf("Method Name: %s", e.MethodName)
	res[3] = fmt.Sprintf("Issue: %s", e.Issue.Error())

	return res[0] + "\n" + res[1] + "\n" + res[2] + "\n" + res[3]
}
