package auth

const TokenName = "Authorization"

type Token struct {
	UserId    string // 用户 id
	LoginName string // 用户名
	Expires   int64  // 有效至
}

// Signed 序列化得到 token
func Signed(token Token) (string, error) {
	return "", nil
}

// Parse 反序列化 token 得到必须的数据
func Parse(tokenStr string) (*Token, error) {
	return nil, nil
}
