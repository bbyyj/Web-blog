package response

import (
	"blog_web/utils"
	"fmt"
)

//专门用来处理错误信息
func CheckError(err error, msg string, args ...interface{}) bool {
	if err != nil {
		if len(args) > 0 {
			utils.Logger().Warning("%s:%v", fmt.Sprintf(msg, args...), err)
		} else {
			utils.Logger().Warning("%s:%v", msg, err)
		}
		return true
	}

	return false
}
