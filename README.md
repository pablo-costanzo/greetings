# Greetings in Go

This package provides a simple way to obtain personalized greetings in Go.

## Installation

Run the following command to install the package:

```bash
go get -u github.com/pablo-costanzo/greetings
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/pablo-costanzo/greetings"
)

func main() {
    message, err := greetings.Hello("Pablo")

    if err != nil {
        fmt.Println("An error occurred:", err)
        return
    }

    fmt.Println(message)
}

```

This example imports the package github.com/pablo-costanzo/greetings and calls the Hello function to get a personalized greeting. If an error occurs, an error message is printed.
