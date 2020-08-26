// +build go1.13

package main

import (
	"errors"

	"github.com/ssoor/gex/pkg/tool"
)

func asBuildErrors(err error) *tool.BuildErrors {
	var errs *tool.BuildErrors
	if errors.As(err, &errs) {
		return errs
	}
	return nil
}
