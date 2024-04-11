package cmd

import (
	"fmt"
	"time"
)

func Timer(loading []string, second int, codeIfPressedEnter func(), codeIfTimeLast func()) {

	enterPressed := make(chan bool)

	var duration time.Duration = time.Duration(second)

	go waitForEnter(enterPressed, second, loading)
	select {
	case <-enterPressed:
		codeIfPressedEnter()
		fmt.Println()
	case <-time.After(duration * time.Second):
		codeIfTimeLast()
		fmt.Println()
	}

}

func waitForEnter(enterPressed chan<- bool, second int, loading []string) {
	stop := make(chan struct{})
	defer close(stop)

	go showAnimation(stop, second, loading)

	var input string
	fmt.Scanln(&input)
	enterPressed <- true
}

func showAnimation(stop <-chan struct{}, seconds int, animation []string) {

	animationInterval := time.Duration(seconds) * time.Second / time.Duration(len(animation))

	timer := time.NewTimer(time.Duration(seconds) * time.Second)

	ticker := time.NewTicker(animationInterval)
	defer ticker.Stop()

	i := 0
	for {
		select {
		case <-stop:
			fmt.Printf("\r%s", animation[i])
			fmt.Print("\r\033[K")
			return
		case <-ticker.C:
			fmt.Printf("\r%s", animation[i])
			i = (i + 1) % len(animation)
		case <-timer.C:
			fmt.Printf("\r%s", animation[len(animation)-1])
			fmt.Print("\r\033[K")
			return
		}
	}
}
