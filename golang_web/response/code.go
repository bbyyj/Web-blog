package response

const (
	QueryFailed uint32 = iota
	QuerySuccess

	OperateFailed = iota + 98
	OperateSuccess

	LoginFailed = iota + 196
	LoginSuccess

	Unauthorized = iota + 294

	DeleteFailed = iota + 393
	DeleteSuccess

	GetResourceSuccess = iota + 554
	GetResourceFailed  = iota + 555
)

var MessageForCode = map[uint32]string{
	QuerySuccess:       "查询成功",
	QueryFailed:        "查询失败",
	OperateSuccess:     "操作成功",
	OperateFailed:      "操作失败",
	LoginSuccess:       "登录成功",
	LoginFailed:        "登录失败",
	Unauthorized:       "未获得授权",
	DeleteSuccess:      "删除成功",
	DeleteFailed:       "删除失败",
	GetResourceSuccess: "获取资源成功",
	GetResourceFailed:  "获取资源失败",
}
