package api

import (
	"encoding/base64"
	"time"

	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/hailaz/gadmin/app/model"
	"github.com/hailaz/gadmin/library/common"
)

var (
	// The underlying JWT middleware.
	GfJWTMiddleware *jwt.GfJWTMiddleware
)

// Initialization function,
// rewrite this function to customized your own JWT settings.
func init() {
	//TokenLookup:     "header: Authorization, query: token, cookie: jwt",
	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Minute * 5,
		MaxRefresh:      time.Minute * 5,
		IdentityKey:     "username",              // 用户关机字
		TokenLookup:     "header: Authorization", // 捕抓请求的指定数据
		TokenHeadName:   "gadmin",                // token 头名称
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,   //登录验证
		LoginResponse:   LoginResponse,   //登录返回token
		RefreshResponse: RefreshResponse, //刷新token
		Unauthorized:    Unauthorized,    //未登录返回
		IdentityHandler: IdentityHandler, //返回数据给Authorizator
		PayloadFunc:     PayloadFunc,     //将Authenticator返回的内容记录到jwt
		Authorizator:    Authorizator,    //接收IdentityHandler数据并判断权限
	})
	if err != nil {
		glog.Fatal("JWT Error:" + err.Error())
	}
	GfJWTMiddleware = authMiddleware
}

// GetLoginCryptoKey 获取登录的加密key
//
// createTime:2019年04月24日 13:57:34
// author:hailaz
func GetLoginCryptoKey(r *ghttp.Request) {
	kid := r.Session.Id()
	ck := common.GenCryptoKey(kid)
	//glog.Debug("kid:" + kid)
	Success(r, ck)
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

func Authorizator(data interface{}, r *ghttp.Request) bool {
	method := r.Method
	path := r.URL.Path
	glog.Debugfln("user:%v ,method:%v ,path:%v", data, method, path)
	return model.Enforcer.Enforce(data, path, method)

}

// IdentityHandler sets the identity for JWT.
func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	return claims["username"]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(r *ghttp.Request, code int, message string) {
	Fail(r, code, message)
}

// LoginResponse is used to define customized login-successful callback function.
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	var tk struct {
		Token  string `json:"token"`
		Expire string `json:"expire"`
	}
	tk.Token = GfJWTMiddleware.TokenHeadName + " " + token
	tk.Expire = expire.Format(time.RFC3339)
	Success(r, tk)
}

// RefreshResponse is used to get a new token no matter current token is expired or not.
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	var tk struct {
		Token  string `json:"token"`
		Expire string `json:"expire"`
	}
	tk.Token = GfJWTMiddleware.TokenHeadName + " " + token
	tk.Expire = expire.Format(time.RFC3339)
	Success(r, tk)
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// Check error (e) to determine the appropriate error message.
func Authenticator(r *ghttp.Request) (interface{}, error) {
	data := r.GetJson()
	name := data.GetString("username")
	pwd := data.GetString("password")
	kid := data.GetString("kid")

	if ck := common.GetCryptoKey(kid); ck != nil {
		if gtime.Second()-ck.TimeStamp >= 5 {
			return nil, jwt.ErrFailedAuthentication
		}
		//glog.Debugfln("%v", ck.Id)
		//glog.Debugfln("%v", ck.Key)
		//glog.Debugfln("%v %v", name, pwd)
		decodePwd, err := base64.StdEncoding.DecodeString(pwd)
		if err != nil {
			return nil, err
		}
		decryptPwd, _ := common.RsaDecrypt(decodePwd, []byte(ck.PrivateKey))
		//glog.Debug(string(decryptPwd))
		password := string(decryptPwd)
		//glog.Debugfln("%v %v", name, password)
		if password != "" {
			u, err := model.GetUserByName(name)
			if err != nil {
				return nil, err
			}
			if u.Password == model.EncryptPassword(password) {
				return g.Map{
					"username": u.UserName,
					"id":       u.Id,
				}, nil
			}

		}
	}

	return nil, jwt.ErrFailedAuthentication
}
