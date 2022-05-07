package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "--help")
	}
	if os.Args[1] == "--help" {
		fmt.Println("spammer <message> <repetitions> <delay>\n" +
			"eg: spammer hello 100 10 # spam hello 100 times with 10ms delay between each repetition")
		return
	}

	message := os.Args[1]
	reps, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	delay, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("starting in 10 seconds, go to the app to spam, hit q three times in a second to quit")
	go func() {
		accumulator := 0
		echan := robotgo.EventStart()
		for v := range echan {
			if v.Keychar == 'q' {
				accumulator++
				if accumulator >= 3 {
					fmt.Println("\nExiting!")
					robotgo.End()
					os.Exit(1)
				}
				time.AfterFunc(time.Second, func() {
					accumulator--
				})
			}
		}
	}()

	robotgo.Sleep(10)

	for i := 0; i < reps; i++ {
		robotgo.TypeStr(message, 0)
		robotgo.MilliSleep(delay)
		robotgo.KeyTap("enter")
	}
}
