package notification

import "fmt"

type ConsoleNotifier struct {
}

func (c *ConsoleNotifier) Update(msg string) {
	fmt.Println(msg)
}