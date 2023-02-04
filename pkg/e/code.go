package e

const (
	SUCCESS = 0
	ERROR   = 500

	ErrorExistUser = 1

	ErrorFailEncryption = 10006

	InvalidParams = 400

	//成员错误
	ErrorNotExistUser = 10003

	ErrorNotCompare = 10007

	ErrorAuthCheckTokenFail    = 30001 //token 错误
	ErrorAuthCheckTokenTimeout = 30002 //token 过期
	ErrorAuthToken             = 30003
	ErrorAuth                  = 30004
	ErrorDatabase              = 40001
)
