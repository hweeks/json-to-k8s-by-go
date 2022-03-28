package main

import (
	"fmt"
	"os"

	"github.com/hoisie/mustache"
)

func handle_err(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	path_to_templates := "../services/base/service.stache"
	_, err := os.ReadFile(path_to_templates)
	handle_err(err)
	services_conf := get_config()
	fmt.Println(services_conf)
	for service_name, service_conf := range services_conf {
		parsed := mustache.RenderFile(path_to_templates, service_conf)
		fmt.Println(service_name)
		fmt.Println(service_conf.APP_PORT)
		fmt.Println(parsed)
	}
}
