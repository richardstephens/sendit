package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sendit/internal/context"
	"sendit/internal/routes"
	"strconv"
)

const defaultConfigFilename = "sendit-config.yaml"
const defaultPort = 7000

type findConfigError struct {
}
func (e findConfigError) Error() string {
	return "Could not find config file"
}
func findConfig() (string, error) {
	if len(os.Args) > 1 {
		return os.Args[1], nil
	}
	info, err := os.Stat(defaultConfigFilename)
	if err == nil && !info.IsDir() {
		return defaultConfigFilename, nil
	}
	info, err = os.Stat(path.Join("/etc", defaultConfigFilename))
	if err == nil && !info.IsDir() {
		return path.Join("/etc", defaultConfigFilename), nil
	}
	return "", findConfigError{}
}
func main() {
	configPath, err := findConfig()
	if err != nil{
		fmt.Println(err)
		return
	}

	config := context.Config{}
	serialisedConfig, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println("Could not read config file")
	}
	yaml.Unmarshal(serialisedConfig, &config)
	if config.Port == 0 {
		config.Port = defaultPort
	}

	indexTpl := template.Must(template.ParseFiles("web/index.html"))
	app := context.App{IndexTpl: indexTpl, Config: &config}

	fmt.Printf("Display name: %s\n", config.DisplayName)
	fmt.Printf("Display name: %s\n", config.DisplayDescription)
	fmt.Printf("Upload path: %s\n", config.UploadPath)
	fmt.Printf("Port: %d\n", config.Port)

	fs := http.FileServer(http.Dir("./web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			routes.HandleIndex(w, r, &app)
		} else if r.URL.Path == "/upload" {
			routes.HandleUpload(w, r, &app)
		} else {
			http.NotFound(w, r)
		}
	})
	err = http.ListenAndServe(":" + strconv.Itoa(config.Port), nil)
	if err != nil {
		panic(err)
	}
}

