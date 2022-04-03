package main

import (
	"fmt"
	"os"
	"path"

	"github.com/hoisie/mustache"
)

func handle_err(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	path_curr, err := os.Getwd()
	handle_err(err)
	resolved_path := path.Join(path_curr, "./services/base/service.stache")
	services_conf := get_config()
	fmt.Println(services_conf)
	for service_name, service_conf := range services_conf {
		parsed := mustache.RenderFile(resolved_path, service_conf)
		fmt.Println(service_name)
		fmt.Println(service_conf.APP_PORT)
		fmt.Println(parsed)
	}
}
