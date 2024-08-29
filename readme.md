# AutoMapper Package

The `automapper` package provides a simple utility to map data from one struct to another using JSON serialization and deserialization. This can be useful for quickly transforming one data structure into another without manually copying fields.

## Installation

```bash
go get github.com/fiuyang/automapper
```

## Usage

### Map Function

The `Map` function is the core utility provided by this package. It allows you to map data from one struct to another by converting the origin struct to JSON and then unmarshalling it into the destination struct.

#### Example:

```go
package main

import (
    "fmt"
    "github.com/fiuyang/automapper"
)

type Source struct {
    Name  string
    Age   int
    Email string
}

type Destination struct {
    Name  string
    Email string
}

func main() {
	source := Source{
		Name:  "John Doe",
		Age:   30,
		Email: "john.doe@example.com",
	}

    var destination Destination

    automapper.Map(source, &destination)

    fmt.Printf("Mapped Destination: %+v\n", destination)
}
```

### StructToJson Function

The `StructToJson` function converts a struct to its JSON string representation.

#### Example:

```go
jsonStr := automapper.StructToJson(source)
fmt.Println(jsonStr)
```

### JsonToStruct Function

The `JsonToStruct` function converts a JSON string to a generic Go data structure (`interface{}`).

#### Example:

```go
jsonStr := `{"Name":"John Doe","Email":"john.doe@example.com"}`
data, err := automapper.JsonToStruct(jsonStr)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Printf("Data: %+v\n", data)
```

## Notes

- **Type Safety**: The `Map` function does not perform type checks, so ensure that the types of fields in the destination struct are compatible with those in the origin struct.
- **Error Handling**: The `Map` function does not handle errors during the unmarshalling process. In production code, consider adding error handling as needed.
- **Limitations**: The package works well for simple structs with compatible JSON tags. However, for more complex transformations, consider other mapping strategies.