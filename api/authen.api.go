package api

import ("github.com/gin-gonic/gin"
	"main/model"
	_"time"


)

func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}

func login(c *gin.Context) {
	c.JSON(200, gin.H{"result": "login"})
}

func register(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil {	
		c.JSON(200, gin.H{"result": "register", "data": user})
	}	
}
