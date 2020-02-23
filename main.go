package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/karrick/godirwalk"
)

var prjPath string

func init() {
	flag.StringVar(&prjPath, "p", ".", "项目路径")
}

func main() {
	flag.Parse()

	allowExt := []string{
		".js",
		".py",
		".html",
		".css",
		".jsx",
		".yml",
	}

	excludeDir := []string{
		"/docs/",
		"/node_modules/",
		"/gui/build/",
		"/vendor/",
	}

	output := "all_code.txt"

	os.Remove(output)

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	p := path.Join(pwd, prjPath)

	all := ""

	if err := godirwalk.Walk(p, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			name := osPathname[len(p):]
			if strings.HasPrefix(name, "/.") {
				return nil
			}

			for _, v := range excludeDir {
				if strings.Contains(name, v) {
					return nil
				}
			}

			allow := false

			ext := path.Ext(osPathname)
			for _, v := range allowExt {
				if v == ext {
					allow = true
					break
				}
			}
			if !allow {
				return nil
			}

			b, err := ioutil.ReadFile(osPathname)
			if err != nil {
				fmt.Print(err)
			}

			content := "文件路径: " + name + "\n\n" + string(b) + "\n\n"
			all += content

			return nil
		},
		Unsorted: true,
	}); err != nil {
		fmt.Println(err)
	}

	if err := ioutil.WriteFile("all_code.txt", []byte(all), 0644); err != nil {
		fmt.Println(err)
	}
}
