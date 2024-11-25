package jwt

// 自定义

import (
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

// 设置超时时间
//var TokenExpireDuration = time.Second * 30

//这个不可以在这里声明，因为配置文件还未加载完，就执行获取操作，使得获取到值为零值
//var TokenExpireDuration = viper.GetInt("auth.jwt_expire")

// 定义 secret 秘钥
var Mysecret = []byte("i love golang")

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中

// 自定义 token 类型
type MyClaims struct {
	UserID int64 `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
	// jwt.StandardClaims也可以，更适合简单的jwt声明中，与早期go代码的兼容性较好 
	// jwt.RegisteredClaims 结构体中包含了官方的声明字段: issuer(签发者), audience（受众）, expiration（过期时间）, 
	//notBefore（某时间之前无效）, subject（特定用户的标识）, jwtid（jwt标识【提高安全性】）
}

// 使用指定的secret生成返回token
func GenToken(userID int64, username string) (string, error) {
	// 在这里 用viper获取 ，确保配置文件已经被读取完
	var TokenExpireDuration = viper.GetInt("auth.jwt_expire")
	// 创建一个我们自己的声明的数据
	c := MyClaims{
		UserID: userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "bluebell",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(TokenExpireDuration)*time.Hour)),
		},
	}
	// 使用指定的签名算法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 返回token
	return token.SignedString(Mysecret)
}

// 解析token
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token：mc接收解析token后其中包含的信息
	var mc = new(MyClaims)
	// 解析token：通过将tokenString解码到mc结构体中
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return Mysecret, nil
	})
	if err != nil {
		// token解析失败
		return nil, err
	}
	if token.Valid { 
		// 校验token，token正确则返回
		return mc, nil
	}
	// 校验token失败
	return nil, errors.New("invalid token")
}



