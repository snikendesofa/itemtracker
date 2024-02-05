package Credentials

import "github.com/gin-gonic/gin"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoadLogin(context *gin.Context) {

}

func Login(context *gin.Context) {

}
