package application

import (
	"github.com/sirupsen/logrus"
)

type Middleware interface {
	Print(string, func(*int,... int), *int, ... int) func(*int, ... int)
}

type HighOrderFunctionsCallBacks struct{}

func (HighOrderFunctionsCallBacks) Print(operationName string, operation func(*int,... int), r *int, n... int) func(*int, ... int) {

	logrus.Infof("[MIDDLEWARE] [%s] %T %d %d", operationName, operation, *r, n)
	return operation
}
