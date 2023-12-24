package back

import (
	"fmt"
	"net/http"
	"os"
	// "path/filepath"
	"strings"

	"io"

	"github.com/gin-gonic/gin"
)

// var dir = "C:\\Users\\Sergio\\Pictures\\Ahri"

// var dir = "D:\\"

var dir = ""

var abs_dir = dir

func Run(tls bool, port int) {

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.Discard

	// createPreviews(dir)

	// set up cert

	r := gin.Default()

	// r.Use(func(c *gin.Context) {
	// 	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	// 	// set cache control to 1 minute
	// 	// c.Header("Cache-Control", "max-age=60")
	// })

	r.POST("/gohome", func(c *gin.Context) {
		dir = ""
		abs_dir = dir
		c.Redirect(http.StatusFound, "/")
	})

	r.POST("/setdir/*root", func(c *gin.Context) {
		dir = c.Param("root")
		dir = strings.ReplaceAll(c.Param("root"), "/", "\\")
		dir = strings.TrimPrefix(dir, "\\")
		abs_dir = dir
		fmt.Println("Set dir to: ", dir)
		grid(c)
	})

	r.StaticFS("/static", http.Dir("static"))

	r.GET("/files/*path", func(c *gin.Context) {
		file_path := c.Param("path")
		abs_path := abs_dir + "\\" + file_path
		c.File(abs_path)
	})

	r.StaticFS("/previews", http.Dir(dir+"\\previews"))

	r.GET("/", folder)

	r.GET("/:folder/*path", folder)

	r.NoRoute(func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "/previews") {
			previews(c)
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	r.GET("/grid/:dir/*deer", grid)
	r.GET("/grid/:dir", grid)
	r.GET("/grid", grid)

	// middleware for cache control

	fmt.Println("Starting server on port ", port)
	var err error
	if tls {
		certpath := "C:\\Certbot\\live\\soncore.sytes.net"
		cert := certpath + "\\fullchain.pem"
		key := certpath + "\\privkey.pem"
		err = r.RunTLS(":"+fmt.Sprint(port), cert, key)
	} else {
		err = r.Run(":" + fmt.Sprint(port))
	}
	if err != nil {
		fmt.Println("Error starting server: ", err)
		os.Exit(1)
	}
}
