package main

import (
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	// set up database connection
	db := Init(os.Getenv("DATABASE_URL"))
	defer db.Close()

	// set up our router
	r := gin.Default()
	r.Use(CORSMiddleware())

	// This will ensure that the angular files are served correctly
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./web/frontend/dist/frontend/index.html")
		} else {
			c.File("./web/frontend/dist/frontend/" + path.Join(dir, file))
		}
	})

	r.GET("/health", HealthCheckHandler)
	// set up our routes and handlers
	authorized := r.Group("/")
	// authorized.Use(Middleware())

	users := authorized.Group("/users")
	users.GET("/:id", GetUserHandler)
	users.POST("/", AddUserHandler)
	users.DELETE("/:id", DeleteUserHandler)
	users.PUT("/:id", UpdateUserHandler)

	err := r.Run(os.Getenv("ADDR"))
	if err != nil {
		panic(err)
	}
}
