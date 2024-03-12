package main

import "fmt"

type notifier interface {
	Notify()
}

type user struct {
	name  string
	email string
}

func (u *user) Notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func sendNotification(n notifier) {
	n.Notify()
}

func main() {
	u := user{"Bill", "bill@email.com"}
	sendNotification(&u)
}
