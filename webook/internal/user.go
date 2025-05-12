package internal

import (
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	emailPattern    = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
	passwordPattern = "^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[!@#$%^&*()_+])[A-Za-z\\d!@#$%^&*()_+]{8,}$"
)

type UserHandler struct {
	EmailRegexp    *regexp.Regexp
	PasswordRegexp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	emailRegexp := regexp.MustCompile(emailPattern, regexp.Compiled)
	r := regexp.MustCompile(passwordPattern, regexp.Compiled)
	return &UserHandler{
		EmailRegexp:    emailRegexp,
		PasswordRegexp: r,
	}
}

func (u *UserHandler) RegistryServer(e *gin.Engine) {
	group := e.Group("/users")
	group.POST("/signup", u.SignUp)
	group.POST("/login", u.Login)
	group.POST("/exit", u.Exit)
}

func (u *UserHandler) SignUp(c *gin.Context) {
	type SignUpReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req SignUpReq
	if err := c.Bind(&req); err != nil {
		c.JSON(400, "")
		return
	}
	fmt.Println(req)
	ok, err := u.EmailRegexp.MatchString(req.Email)
	if err != nil {
		c.String(http.StatusInternalServerError, "系统内部错误")
		return
	}
	if !ok {
		c.String(http.StatusBadRequest, "邮箱格式错误")
		return
		return
	}

	ok, err = u.PasswordRegexp.MatchString(req.Password)
	if err != nil {
		c.String(http.StatusInternalServerError, "系统内部错误")
		return
	}
	if !ok {
		c.String(http.StatusBadRequest, "密码必须大于8位，包含大小写和特殊字符")
		return
	}

	c.String(http.StatusOK, "注册成功")

}

func (u *UserHandler) Login(c *gin.Context) {

}

func (u *UserHandler) Exit(c *gin.Context) {

}
