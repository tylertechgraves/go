package main

import (
	"fmt"
	"time"

	"github.com/tyler-technologies/exercism/go/bob"
	"github.com/tyler-technologies/exercism/go/gigasecond"
)

func main() {
	result := gigasecond.AddGigasecond(time.Now())

	fmt.Println(result)
	fmt.Println(bob.Hey("You suck!"))
}
