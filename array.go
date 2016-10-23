//This code was shamelessly stolen
//I never would have figured this out on my own
//http://marcelom.github.io/2013/06/07/goshuffle.html

package main

import (
	"fmt"
	"math/rand"
	"time"//I get this part as a way to seed the random shuffle function
)

var card string = "\n+-----+\n|A   S|\n|     |\n|     |\n|     |\n|S   A|\n+-----+"

func Shuffle(a []string) {
	for i := range a {
		j := rand.Intn(i+1)
		a[i], a[j] = a[j], a[i]//not sure what this is all doing
	}
}


func main() {
	a := []string{//Hard-coded deck, may try to map suit types and values
		"AS","KS","QS","JS","XS","9S","8S","7S","6S","5S","4S","3S","2S",
		"AH","KH","QH","JH","XH","9H","8H","7H","6H","5H","4H","3H","2H",
		"AC","KC","QC","JC","XC","9C","8C","7C","6C","5C","4C","3C","2C",
		"AD","KD","QD","JD","XD","9D","8D","7D","6D","5D","4D","3D","2D",
	}
	rand.Seed(time.Now() .UnixNano())//Never would have reasoned out the UnixNano thing
	Shuffle(a)
	fmt.Println(a[0], a[1], a[2], a[3], a[4], card, card)
}
