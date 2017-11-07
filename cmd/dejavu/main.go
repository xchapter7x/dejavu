package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/oxtoacart/bpool"
)

var port string
var viewpath string
var bufpool *bpool.BufferPool
var templates = make(map[string]*template.Template)

func init() {
	bufpool = bpool.NewBufferPool(64)
	if port = fmt.Sprintf(":%s", os.Getenv("PORT")); port == ":" {
		port = ":8080"
	}

	if viewpath = os.Getenv("VIEW_PATH"); viewpath == "" {
		viewpath = "views"
	}
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob(viewpath + "**/*")
	v1 := router.Group("/v1")
	{
		v1.GET("/info", endpoint)
		v1.GET("/list/views", endpoint)
		v1.GET("/list/js", endpoint)
		v1.GET("/list/css", endpoint)
		layouts, err := filepath.Glob(viewpath + "**/*.tmpl")
		if err != nil {
			log.Fatal(err)
		}
		for _, layout := range layouts {
			fmt.Printf("adding route %s", layout)
			v1.POST(fmt.Sprintf("/render/view/%s", layout), renderTemplate(layout))
		}
	}
	router.Run(port)
}

func endpoint(c *gin.Context) {
	data, err := c.GetRawData()
	c.JSON(200, gin.H{
		"message": string(data),
		"error":   err,
	})
}

func renderTemplate(templatePath string) func(*gin.Context) {
	return func(c *gin.Context) {
		bdata, _ := c.GetRawData()
		data := gin.H{}
		json.Unmarshal(bdata, &data)
		c.HTML(http.StatusOK, templatePath, data)
	}
}
