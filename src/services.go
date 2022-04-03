package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"strconv"

	"gopkg.in/guregu/null.v4/zero"
)

type services_base struct {
	REPLICA_COUNT string `json:"replica_count"`
	CPU_BASE      string `json:"cpu_base"`
	MEM_BASE      string `json:"mem_base"`
	CPU_LIMIT     string `json:"cpu_limit"`
	MEM_LIMIT     string `json:"mem_limit"`
}

type services_extended struct {
	BASE_NAME     string      `json:"base_name"`
	IMAGE_URL     string      `json:"image_url"`
	APP_PORT      string      `json:"app_port"`
	HEALTH_CHECK  string      `json:"health_check"`
	REPLICA_COUNT zero.Int    `json:"replica_count,omitempty"`
	CPU_BASE      zero.String `json:"cpu_base,omitempty"`
	MEM_BASE      zero.String `json:"mem_base,omitempty"`
	CPU_LIMIT     zero.String `json:"cpu_limit,omitempty"`
	MEM_LIMIT     zero.String `json:"mem_limit,omitempty"`
}

type services_template struct {
	BASE_NAME     string `json:"base_name"`
	IMAGE_URL     string `json:"image_url"`
	APP_PORT      string `json:"app_port"`
	HEALTH_CHECK  string `json:"health_check"`
	REPLICA_COUNT int    `json:"replica_count"`
	CPU_BASE      string `json:"cpu_base"`
	MEM_BASE      string `json:"mem_base"`
	CPU_LIMIT     string `json:"cpu_limit"`
	MEM_LIMIT     string `json:"mem_limit"`
}

type services_json struct {
	Base     services_base                `json:"base"`
	Services map[string]services_extended `json:"services"`
}

func unmarshall_partial(values services_extended) services_template {
	final_values := services_template{}
	final_values.APP_PORT = values.APP_PORT
	final_values.BASE_NAME = values.BASE_NAME
	final_values.CPU_BASE = values.CPU_BASE.ValueOrZero()
	final_values.CPU_LIMIT = values.CPU_LIMIT.ValueOrZero()
	final_values.HEALTH_CHECK = values.HEALTH_CHECK
	final_values.IMAGE_URL = values.IMAGE_URL
	final_values.MEM_BASE = values.MEM_BASE.ValueOrZero()
	final_values.MEM_LIMIT = values.MEM_LIMIT.ValueOrZero()
	final_values.REPLICA_COUNT = int(values.REPLICA_COUNT.ValueOrZero())
	return final_values
}

func get_config() map[string]services_template {
	path_curr, err := os.Getwd()
	handle_err(err)
	resolved_path := path.Join(path_curr, "./src/services.json")
	service_json, err := os.Open(resolved_path)
	handle_err(err)
	defer service_json.Close()
	json_bytes, _ := ioutil.ReadAll(service_json)
	// create a struct, don't do a var decleration
	all_services := services_json{}
	err = json.Unmarshal(json_bytes, &all_services)
	if err != nil {
		panic(err)
	}
	// you must make a map to allow unmarshalling into it
	renderable_services := make(map[string]services_template)
	for sub_service, values := range all_services.Services {
		// this is a lot of overhead to achieve a nullable field
		// parse for null
		if values.REPLICA_COUNT == zero.IntFrom(0) {
			n, err := strconv.ParseInt(all_services.Base.REPLICA_COUNT, 10, 64)
			handle_err(err)
			values.REPLICA_COUNT = zero.IntFrom(int64(n))
		}
		if values.CPU_BASE == zero.StringFrom("") {
			values.CPU_BASE = zero.StringFrom(all_services.Base.CPU_BASE)
		}
		if values.MEM_BASE == zero.StringFrom("") {
			values.MEM_BASE = zero.StringFrom(all_services.Base.MEM_BASE)
		}
		if values.CPU_LIMIT == zero.StringFrom("") {
			values.CPU_LIMIT = zero.StringFrom(all_services.Base.CPU_LIMIT)
		}
		if values.MEM_LIMIT == zero.StringFrom("") {
			values.MEM_LIMIT = zero.StringFrom(all_services.Base.MEM_LIMIT)
		}
		// parse and fill
		final_values := unmarshall_partial(values)
		renderable_services[sub_service] = final_values
	}
	return renderable_services
}
