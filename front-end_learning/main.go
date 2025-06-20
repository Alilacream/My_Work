package main

import(
	"github.com/gin-gonic/gin"
	"html/template"
)
	type People struct {
		Name string
		Age int
	}
func handler(ctx *gin.Context) {
	tmpl:= template.Must(template.ParseFiles("index.html"))
	people := map[string][]People {
		"People": {
			{Name: "Hassane", Age: 21},
		},
	}
	tmpl.Execute(ctx.Writer, people)
}
func main() {
	app := gin.Default()

	defer app.Run(":8000")

	app.GET("/", handler)
}