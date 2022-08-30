package utils

var InterfaceTmpl = `package {{ .NameLowerCase}}

type {{ .Name}}Repository interface {
}
`

var ModelTmpl = `package {{ .NameLowerCase}}

type {{ .Name}} struct {
}
`

var ControllerTmpl = `package {{ .NameLowerCase}}

import "github.com/gin-gonic/gin"

func Create{{ .Name}}(c *gin.Context) {
	create(c.Request.Body)
}

func RetrieveAll{{ .Name}}(c *gin.Context) {
	findAll()
}

func Retrieve{{ .Name}}(c *gin.Context) {
	id := c.Param("id")
	findOne(id)
}

func Update{{ .Name}}(c *gin.Context) {
	id := c.Param("id")
	update(id, c.Request.Body)
}

func Delete{{ .Name}}(c *gin.Context) {
	id := c.Param("id")
	remove(id)
}
`

var RouterTmpl = `package {{ .NameLowerCase}}

import "github.com/gin-gonic/gin"

func {{ .Name}}Router(router *gin.RouterGroup) {
	router.GET("/:id", Retrieve{{ .Name}})
	router.GET("/", RetrieveAll{{ .Name}})
	router.POST("/", Create{{ .Name}})
	router.PATCH("/", Update{{ .Name}})
	router.DELETE("/", Delete{{ .Name}})
}
`

var ServiceTmpl = `package {{ .NameLowerCase}}

import "fmt"

func create(data interface{}) string {
	return "This action adds a new {{ .NameLowerCase}}"
}

func findAll() string {
	return "This action adds all moto"
}

func findOne(id string) string {
	return fmt.Sprintf("This action returns a #%s {{ .NameLowerCase}}", id)
}

func update(id string, data interface{}) string {
	return fmt.Sprintf("This action updates a #%s {{ .NameLowerCase}}", id)
}

func remove(id string) string {
	return fmt.Sprintf("This action removes a #%s {{ .NameLowerCase}}", id)
}
`

var DtoTmpl = `package {{ .NameLowerCase}}

type Create{{ .Name}}Dto struct{}

type Update{{ .Name}}Dto struct{}

type Query{{ .Name}}Dto struct{}

`
