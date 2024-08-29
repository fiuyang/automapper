package automapper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestMap(t *testing.T) {
	person := Person{Name: "Alice", Age: 30}
	var employee Employee

	Map(person, &employee)

	assert.Equal(t, Employee{Name: "Alice", Age: 30}, employee)
}

func TestStructToJson(t *testing.T) {
	person := Person{Name: "Bob", Age: 25}
	expectedJson := `{"name":"Bob","age":25}`

	jsonResult := StructToJson(person)

	assert.Equal(t, expectedJson, jsonResult)
}

func TestJsonToStruct(t *testing.T) {
	jsonStr := `{"name":"Charlie","age":35}`
	expected := map[string]interface{}{
		"name": "Charlie",
		"age":  float64(35),
	}

	result, err := JsonToStruct(jsonStr)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestJsonToStructInvalidJson(t *testing.T) {
	invalidJsonStr := `{"Name":"Charlie", "Age":}`

	_, err := JsonToStruct(invalidJsonStr)
	assert.NotNil(t, err)
}