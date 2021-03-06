package main

import (
	"fmt"
	"time"

	"github.com/tyler-technologies/exercism/go/acronym"
	"github.com/tyler-technologies/exercism/go/bob"
	diffsquares "github.com/tyler-technologies/exercism/go/difference-of-squares"
	"github.com/tyler-technologies/exercism/go/gigasecond"
	"github.com/tyler-technologies/exercism/go/grains"
	"github.com/tyler-technologies/exercism/go/hamming"
	greeting "github.com/tyler-technologies/exercism/go/hello-world"
	"github.com/tyler-technologies/exercism/go/isogram"
	"github.com/tyler-technologies/exercism/go/luhn"
	"github.com/tyler-technologies/exercism/go/raindrops"
	"github.com/tyler-technologies/exercism/go/reverse-string"
	"github.com/tyler-technologies/exercism/go/scrabble-score"
	space "github.com/tyler-technologies/exercism/go/space-age"
	twofer "github.com/tyler-technologies/exercism/go/two-fer"
)

func main() {
	// Gigasecond
	fmt.Println(gigasecond.AddGigasecond(time.Now()))

	// Bob
	fmt.Println(bob.Hey("You suck!"))

	// Acronym
	fmt.Println(acronym.Abbreviate("Complementary metal-oxide semiconductor"))

	// Hello World
	fmt.Println(greeting.HelloWorld())

	// SpaceAge
	fmt.Println(space.Age(1000000000, "Earth"))

	// Two Fer
	fmt.Println(twofer.ShareWith("McLovin"))

	// Hamming
	fmt.Println(hamming.Distance("ABCDEFG", "ABCDEFW"))

	// Raindrops
	fmt.Println(raindrops.Convert(256))

	// Scrabble-score
	fmt.Println(scrabble.Score("HonkyBallBearAtol"))

	// Isogram
	fmt.Println(isogram.IsIsogram("a-b-c d e f g"))

	// Reverse string
	fmt.Println(reverse.Reverse("Snodgrass"))

	// Difference of Squares
	fmt.Println(diffsquares.SquareOfSum(2))
	fmt.Println(diffsquares.SumOfSquares(3))
	fmt.Println(diffsquares.Difference(10))

	// Luhn
	fmt.Println(luhn.Valid("059"))

	// Grains
	// fmt.Println(grains.Square(0))
	fmt.Println(grains.Total())
}
