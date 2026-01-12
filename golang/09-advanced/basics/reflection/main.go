package main

import (
	"fmt"
	"reflect"
)

type Config struct {
	Endpoint string `env:"API_URL"`
	Port     int    `env:"PORT"`
	Debug    bool
}

func main() {
	c := Config{
		Endpoint: "https://api.example.com",
		Port:     8080,
		Debug:    true,
	}

	inspectStruct(c)
}

func inspectStruct(i interface{}) {
	// 1. Get Value and Type
	val := reflect.ValueOf(i)
	typ := reflect.TypeOf(i)

	fmt.Printf("Type: %s\n", typ.Name())
	fmt.Printf("Kind: %s\n", typ.Kind())

	if typ.Kind() != reflect.Struct {
		fmt.Println("Error: Expected a struct")
		return
	}

	// 2. Iterate Fields
	fmt.Println("--- Fields ---")
	for x := 0; x < val.NumField(); x++ {
		fieldVal := val.Field(x)
		fieldType := typ.Field(x)

		name := fieldType.Name
		tag := fieldType.Tag.Get("env")
		value := fieldVal.Interface() // Convert back to interface{} to print

		fmt.Printf("Name: %-10s | Type: %-10s | Tag: %-10s | Value: %v\n",
			name, fieldType.Type, tag, value)
	}
}
