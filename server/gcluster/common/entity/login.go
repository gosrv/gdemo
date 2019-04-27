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
	Code string
}

func NewLoginRegisterResult(success bool) *LoginRegisterResult {
	if success {
		return &LoginRegisterResult{Code: "ok"}
	} else {
		return &LoginRegisterResult{Code: "error"}
	}
}

type LoginCheckResult struct {
	Code string
	Id int64
}

func NewLoginCheckResult(code string, id int64) *LoginCheckResult {
	return &LoginCheckResult{Code:code, Id: id}
}
