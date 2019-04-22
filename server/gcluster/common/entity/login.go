package entity

type LoginTokenCtx struct {
	Account string
	Pw      string
	Id      int64
	Token   string
}

func NewLoginTokenCtx(account string, pw string, id int64, token string) *LoginTokenCtx {
	return &LoginTokenCtx{Account: account, Pw: pw, Id: id, Token: token}
}

type LoginResult struct {
	Code  string
	Token string
}

func NewLoginResult(code string, token string) *LoginResult {
	return &LoginResult{Code: code, Token: token}
}

type LoginRegisterResult struct {
	Success bool
}

func NewLoginRegisterResult(success bool) *LoginRegisterResult {
	return &LoginRegisterResult{Success: success}
}

type LoginCheckResult struct {
	Id int64
}

func NewLoginCheckResult(id int64) *LoginCheckResult {
	return &LoginCheckResult{Id: id}
}
