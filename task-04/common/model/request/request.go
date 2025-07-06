package request

import (
	"github.com/astaxie/beego/validation"
	"github.com/endymion/go-base/task-04/common/logger"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logger.Info(err.Key, err.Message)
	}
	return
}
