package main

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/wekex35/go-gen/utils"
)

type TemplateArgs struct {
	Name          string
	NameLowerCase string
}

func renderWriteToFile(tmpl string, func_name string, file_name string) {

	td := TemplateArgs{strings.Title(func_name), func_name}
	t, err := template.New("name").Parse(tmpl)
	if err != nil {
		fmt.Println("errrror", err)
	}
	f, err := os.Create(file_name)
	if err != nil {
		fmt.Println("create file: ", err)
		return
	}
	err = t.Execute(f, td)
	if err != nil {
		panic(err)
	}
}

func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}

func getFileDirectory(filetype string, file_path string, filename string, isBase bool) string {
	mdir := filepath.Join("./src", file_path, filename)

	dir := filepath.Join(mdir, filetype)

	if isBase {
		dir = mdir
	}
	fmt.Println(dir)
	if err := ensureDir(dir); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		os.Exit(1)
	}
	return fmt.Sprintf("%s/%s.%s.go", dir, filename, filetype)
}

func controller(file_name string, file_path string) {
	renderWriteToFile(utils.ControllerTmpl, file_name, getFileDirectory("controllers", file_path, file_name, true))
}

func interfaces(file_name string, file_path string) {
	renderWriteToFile(utils.InterfaceTmpl, file_name, getFileDirectory("interfaces", file_path, file_name, false))
}

func model(file_name string, file_path string) {
	renderWriteToFile(utils.ModelTmpl, file_name, getFileDirectory("models", file_path, file_name, false))
}

func router(file_name string, file_path string) {
	renderWriteToFile(utils.RouterTmpl, file_name, getFileDirectory("router", file_path, file_name, true))
}

func service(file_name string, file_path string) {
	renderWriteToFile(utils.ServiceTmpl, file_name, getFileDirectory("service", file_path, file_name, true))
}

func dto(file_name string, file_path string) {
	renderWriteToFile(utils.DtoTmpl, file_name, getFileDirectory("dto", file_path, file_name, false))
}

func res(file_name string, file_path string) {
	controller(file_name, file_path)
	interfaces(file_name, file_path)
	model(file_name, file_path)
	router(file_name, file_path)
	service(file_name, file_path)
	dto(file_name, file_path)
}

func generate(cmd string, name string) {
	df := strings.Split(name, "/")
	file_name := df[len(df)-1]
	file_path := strings.Join(df[0:len(df)-1], "/")

	switch cmd {
	case "res":
		{
			res(file_name, file_path)
			break
		}
	case "co":
		{
			controller(file_name, file_path)
			break
		}
	case "se":
		{
			service(file_name, file_path)
			break
		}
	case "ro":
		{
			router(file_name, file_path)
			break
		}
	case "mo":
		{
			model(file_name, file_path)
			break
		}
	case "in":
		{
			interfaces(file_name, file_path)
			break
		}

	case "dto":
		{
			dto(file_name, file_path)
			break
		}

	}

}

func Contains(collection []string, element string) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}

	return false
}

func help() {
	fmt.Println("<cmd>	<args>	description")
	fmt.Println("res	<name>	to generate resouces like router,dto,controller, and service")
	fmt.Println("co	<name>	to generate controller")
	fmt.Println("se	<name>	to generate service")
	fmt.Println("ro	<name>	to generate router")
	fmt.Println("mo	<name>	to generate model")
	fmt.Println("in	<name>	to generate interface")
	fmt.Println("dto	<name>	to generate dto")

	os.Exit(1)
}

func main() {
	//accepted cmd
	args := os.Args
	acceptedCmd := []string{"res",
		"co",
		"se",
		"ro",
		"mo",
		"in",
		"dto"}
	if len(args) < 2 {
		fmt.Println("Please use one of the <cmd>")
		help()
	}

	cmd := args[1]
	present := Contains(acceptedCmd, cmd)

	if !present {
		fmt.Println("Invalid <cmd>")
		help()
	}

	if len(args) >= 3 {
		name := args[2]
		generate(cmd, name)
	} else {
		fmt.Println("name is required")
		help()
	}
}
