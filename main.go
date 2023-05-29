package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var words = []string{
	"ACT", "ADD", "AGE", "AIR", "ALL", "AND", "ANY", "ARE", "ART", "BAD", "BAG", "BAR",
	"BAT", "BED", "BET", "BIG", "BIT", "BOY", "BUT", "CAN", "CAR", "CAT", "CUP", "CUT",
	"DID", "DOG", "DRY", "EAT", "END", "FAN", "FAR", "FAT", "FIT", "FIX", "FLY", "FOR",
	"FUN", "GET", "GOT", "GUM", "GUN", "HAT", "HER", "HIM", "HIT", "HOT", "HOW", "ICE",
	"INK", "JET", "JOB", "KEY", "KID", "KIN", "LAY", "LET", "LIE", "LIT", "LOW", "MAN",
	"MAP", "MAY", "MEN", "MOM", "MUD", "NAG", "NEW", "NOT", "NUT", "OFF", "OIL", "OLD",
	"ONE", "OUR", "OUT", "PAN", "PEN", "PET", "PIE", "PIN", "PIT", "POT", "PUT", "RAN",
	"RED", "RUG", "RUN", "SAD", "SAT", "SAY", "SEE", "SET", "SHE", "SIN", "SIT", "SKY",
	"SON", "TAN", "TAP", "TEA", "TIE", "TOE", "TOP", "TOY", "TWO", "USE", "VAN", "WAR",
	"WAS", "WAY", "WET", "WHO", "WHY", "WIN", "YES", "YET", "YOU", "ZOO"}

var word string
var score int

func main() {
	r := gin.Default()
	word = words[newRand().Intn(len(words))]
	fmt.Println("Chosen word =========> :", word) // This line logs the chosen word

	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{"score": score})
	})

	r.POST("/", func(c *gin.Context) {
		guess := strings.ToUpper(c.PostForm("guess"))
		score = calculateScore(word, guess)

		if score == 30 {
			c.HTML(200, "index.tmpl", gin.H{"score": score, "message": "Correct, the word was " + word + "! Now, let's play again!"})
			word = words[newRand().Intn(len(words))]      // choose a new word
			fmt.Println("Chosen word =========> :", word) // This line logs the chosen word
			score = 0                                     // reset score
		} else {
			c.HTML(200, "index.tmpl", gin.H{"score": score, "message": ""})
		}
	})

	r.Run(":8080")
}

func calculateScore(word, guess string) int {
	score := 0
	for i := 0; i < 3; i++ {
		if strings.Contains(word, string(guess[i])) {
			score += 5
			if word[i] == guess[i] {
				// score += 10 // Bug in ChatGPT generated code
				score += 5
			}
		}
	}
	return score
}

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
