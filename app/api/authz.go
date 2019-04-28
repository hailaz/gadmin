package api

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/hailaz/gadmin/library/common"
)

// NewAuthorizer returns the authorizer.
// Use a casbin enforcer as input
func NewAuthorizer(e *casbin.Enforcer) ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		a := &BasicAuthorizer{enforcer: e}
		if !a.CheckPermission(r) {
			a.RequirePermission(r.Response.Writer)
			r.ExitAll()
		}
	}
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// GetUserName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func (a *BasicAuthorizer) GetUserName(r *ghttp.Request) string {
	token := r.GetString("token", r.Header.Get("X-Token"))
	if token != "" {
		jwtobj, err := common.PareseJWT(token)
		if err == nil && jwtobj != nil {
			return jwtobj.Username
		}
	}
	return ""
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(r *ghttp.Request) bool {
	user := a.GetUserName(r)
	method := r.Method
	path := r.URL.Path
	glog.Debugfln("user:%v ,method:%v ,path:%v", user, method, path)
	return a.enforcer.Enforce(user, path, method)
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(w http.ResponseWriter) {
	w.WriteHeader(403)
	w.Write([]byte("403 Forbidden\n"))
}
