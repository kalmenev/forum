package pages

import (
	"html/template"
	"os"
	"path/filepath"
)

// var ALL = template.Must(template.ParseFiles("./ui/templates/index.html", "./ui/templates/sign-in.html", "./ui/templates/sign-up.html", "./ui/templates/createpost.html", "./ui/templates/allposts.html", "./ui/templates/error.html"))

var ALL = template.Must(template.ParseFiles(getTemplateFiles("./ui/templates")...))

func getTemplateFiles(dir string) []string {
	var fileList []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".html" {
			fileList = append(fileList, path)
		}

		return nil
	})
	if err != nil {
		panic(err.Error())
	}

	return fileList
}
