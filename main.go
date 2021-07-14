package main

import (
	"fmt"
	"time"

	"github.com/tyler-technologies/exercism/go/acronym"
	"github.com/tyler-technologies/exercism/go/bob"
	"github.com/tyler-technologies/exercism/go/gigasecond"
	greeting "github.com/tyler-technologies/exercism/go/hello-world"
)

func main() {
	result := gigasecond.AddGigasecond(time.Now())

	fmt.Println(result)
	fmt.Println(bob.Hey("You suck!"))
	fmt.Println(acronym.Abbreviate("Complementary metal-oxide semiconductor"))

	fmt.Println(greeting.HelloWorld())
}
