package gerror

import "github.com/pkg/errors"

var (
	ErrorNotDFoundInDBByCondition = errors.New("Not Found In DB By Condition")
	ErrorFound = errors.New("查询失败")
)
