package main

import (
	"encoding/json"
	"fmt"
	"ordering-system-backend/utils"
)

type MyStruct struct {
	Field1 int    `json:"field1"`
	Field2 string `json:"field2"`
	Field3 bool   `json:"field3"`
}

func (m MyStruct) MyStructFilter() (string, error) {
	filter := []string{"field1", "field3"}
	b, err := json.MarshalIndent(utils.SelectFields(m, filter...), "", "  ")
	if err != nil {
		panic(err.Error())
	}
	return string(b), nil
}

func main() {
	// config.InitConfig()

	// db.Connect()

	// router := gin.Default()
	// routes.SetUpRoutes(router)

	// router.Run("localhost:" + config.Env.Port)
	myInstance := MyStruct{
		Field1: 42,
		Field2: "Hello",
		Field3: true,
	}

	result, err := myInstance.MyStructFilter()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
