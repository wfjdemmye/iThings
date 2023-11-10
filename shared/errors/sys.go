package errors

const SYS_ERROR = 100000

var (
	OK               = NewCodeError(200, "成功")
	Default          = NewCodeError(SYS_ERROR+1, "其他错误")
	TokenExpired     = NewCodeError(SYS_ERROR+2, "token已经过期")
	TokenNotValidYet = NewCodeError(SYS_ERROR+3, "token还未生效")
	TokenMalformed   = NewCodeError(SYS_ERROR+4, "token格式错误")
	TokenInvalid     = NewCodeError(SYS_ERROR+5, "违法的token")
	Parameter        = NewCodeError(SYS_ERROR+6, "参数错误")
	System           = NewCodeError(SYS_ERROR+7, "系统错误")
	Database         = NewCodeError(SYS_ERROR+8, "数据库错误")
	NotFind          = NewCodeError(SYS_ERROR+9, "未查询到")
	Duplicate        = NewCodeError(SYS_ERROR+10, "参数重复")
	SignatureExpired = NewCodeError(SYS_ERROR+11, "签名已经过期")
	Permissions      = NewCodeError(SYS_ERROR+12, "权限不足")
	Method           = NewCodeError(SYS_ERROR+13, "method不支持")
	Type             = NewCodeError(SYS_ERROR+14, "参数的类型不对")
	OutRange         = NewCodeError(SYS_ERROR+15, "参数的值超出范围")
	TimeOut          = NewCodeError(SYS_ERROR+16, "等待超时")
	Server           = NewCodeError(SYS_ERROR+17, "本实例处理不了该信息")
	NotRealize       = NewCodeError(SYS_ERROR+18, "尚未实现")
	NotEmpty         = NewCodeError(SYS_ERROR+19, "不为空")
	Panic            = NewCodeError(SYS_ERROR+20, "系统异常，请联系开发者")
	NotEnable        = NewCodeError(SYS_ERROR+21, "未启用")
	Company          = NewCodeError(SYS_ERROR+22, "该功能是企业版功能")
	Script           = NewCodeError(SYS_ERROR+23, "脚本执行失败")
	OnGoing          = NewCodeError(SYS_ERROR+24, "正在执行中")     //事务分布式事务中如果返回该错误码,分布式事务会定时重试
	Failure          = NewCodeError(SYS_ERROR+25, "执行失败,需要回滚") //事务分布式事务中如果返回该错误码,分布式事务会进行回滚

)
