package routes

import (
    "gin-api/controllers"
	"gin-api/service"
	// "gin-api/middleware"
	"net/http"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {

	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controllers.LoginController = controllers.LoginHandler(loginService, jwtService)


    r := gin.Default()

    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })

	r.POST("/login", func(c *gin.Context) {
		token := loginController.Login(c)
		if token != "" {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"code_message" : 401, "message": "Failed", "error": "Login Failed"})
		}
	})

	r.Static("/css", "./templates/css")
	r.LoadHTMLGlob("templates/*.html")
	
	
	v1 := r.Group("api/v1")
  	{
		// PROFILE
		v1.GET("/allprofiles", controllers.FindProfiles)
		v1.GET("/profiles", controllers.GetPaginateProfiles)
    	v1.POST("/profiles", controllers.CreateProfile)
		v1.GET("/profiles/:id", controllers.FindProfile)
		v1.PATCH("/profiles/:id", controllers.UpdateProfile)
		v1.DELETE("/profiles/:id", controllers.DeleteProfile)

		// TASK
    	v1.GET("/alltasks", controllers.FindTasks)
		v1.GET("/tasks", controllers.GetPaginateTasks)
    	v1.POST("/tasks", controllers.CreateTask)
    	v1.GET("/tasks/:id", controllers.FindTask)
    	v1.PATCH("/tasks/:id", controllers.UpdateTask)
    	v1.DELETE("tasks/:id", controllers.DeleteTask)
  	}

	  view := r.Group("view")
	  {
		  view.GET("/index", controllers.ShowAll)
	  }
    return r
}