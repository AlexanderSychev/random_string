# random_string

Utils for random strings generations for Go language

## Installation

```shell
go get github.com/AlexanderSychev/random_string
```

## Usage

Watch full API documentation [here](https://pkg.go.dev/github.com/AlexanderSychev/random_string#section-documentation)

### Base example

Generate a string of 50 characters which contains random digits and latin characters

```go
package main

import (
	"fmt"
	"github.com/AlexanderSychev/random_string"
	"math/rand"
	"time"
)

func main() {
	// Set seed value for default randomizer
	rand.Seed(time.Now().Unix())

	// Generate 50 characters length string
	str := random_string.Generate(50)

	// Print generated string
	fmt.Println(str)
}
```

### Custom generator

Create custom generator which can generate strings of 50 characters which contains characters from
custom set (`!!@#$%^&*()-=+_|`):

```go
package main

import (
	"fmt"
	"github.com/AlexanderSychev/random_string"
	"math/rand"
	"time"
)

func main() {
	// Set seed value for default randomizer
	rand.Seed(time.Now().Unix())

	// Create custom generator
	generator := random_string.NewGeneratorFromString(50, "!!@#$%^&*()-=+_|")

	// Generate 50 characters length string
	str := generator.Generate()

	// Print generated string
	fmt.Println(str)
}
```
