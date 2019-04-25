package api

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/g/os/glog"

	"github.com/gogf/gf/g/container/gmap"
	"github.com/gogf/gf/g/crypto/gaes"
	"github.com/gogf/gf/g/encoding/gbase64"
	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/gogf/gf/g/util/grand"
	"github.com/hailaz/gadmin/app/model"
)

type LoginController struct {
	gmvc.Controller
}

type CryptoKey struct {
	CryptoType string `json:"cryptotype"`
	Key        string `json:"key"`
	TimeStamp  int64  `json:"timestamp"` //sec
}

var CryptoKeyList = gmap.New()

// GetLoginCryptoKey description
//
// creatTime:2019年04月24日 13:57:34
// author:hailaz
func (c *LoginController) GetLoginCryptoKey() {
	sid := c.Request.Session.Id()
	ck := CryptoKey{CryptoType: "AES", Key: grand.RandLetters(16), TimeStamp: gtime.Second()}
	CryptoKeyList.Set(sid, ck)
	//c.Response.Writeln("sid:" + fmt.Sprintf("%v",sid))
	glog.Debug("sid:" + sid)
	Success(c.Response, ck)

}

// Login 登录
//
// creatTime:2019年04月24日 15:49:56
// author:hailaz
func (c *LoginController) Login() {
	sid := c.Request.Session.Id()
	var tk struct {
		Token string `json:"token"`
	}
	var ck CryptoKey
	if k := CryptoKeyList.Get(sid); k != nil {
		ck = k.(CryptoKey)
		//glog.Debugfln("%v", ck)
		data := c.Request.GetJson()
		name := data.GetString("name")
		pwd := data.GetString("pwd")
		//glog.Debugfln("%v %v", name, pwd)
		password := DecryptPassword(pwd, ck.Key)
		//glog.Debugfln("%v %v", name, password)
		if password != "" {
			u, err := model.GetUserByName(name)
			if err != nil {
				Fail(c.Response, 1)
				return
			}
			if u.Password == model.EncryptPassword(password) {
				token, _ := getJWT(u.UserName)
				pareseJWT(token)
				tk.Token = token
				Success(c.Response, tk)
				return
			}

		}
	}
	Fail(c.Response, 1)
}

// DecryptPassword 解密密码
// 输入base64编码的加密数据
// 算法模式：“CBC”
// 补码方式：“PKCS5Padding”
// 密钥偏移量：“I Love Go Frame!”
// creatTime:2019年04月24日 16:50:59
// author:hailaz
func DecryptPassword(data, key string) string {
	temp, err := gbase64.Decode(data)
	if err != nil {
		return ""
	}
	b, err := gaes.Decrypt([]byte(temp), []byte(key))
	if err != nil {
		return ""
	}
	return string(b)
}

type JsonWebToken struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const (
	TOKEN_SIGNING_KEY = "gadmin"
)

// getJWT description
//
// creatTime:2019年04月25日 10:28:51
// author:hailaz
func getJWT(username string) (string, error) {

	// Create the Claims
	claims := JsonWebToken{
		username,
		jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(gtime.Second() + 60)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(TOKEN_SIGNING_KEY))

}

// pareseJWT description
//
// creatTime:2019年04月25日 15:48:51
// author:hailaz
func pareseJWT(tokenString string) *JsonWebToken {
	// sample token is expired.  override time so it parses as valid
	token, _ := jwt.ParseWithClaims(tokenString, &JsonWebToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(TOKEN_SIGNING_KEY), nil
	})

	if claims, ok := token.Claims.(*JsonWebToken); ok && token.Valid {
		//fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
		return claims
	}
	return nil
}
