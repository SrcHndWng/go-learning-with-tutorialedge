package main

import (
	"fmt"

	"github.com/SrcHndWng/go-learning-with-tutorialedge/mutex/models"
)

func main() {
	fmt.Println("Go Mutex Example")

	done := make(chan bool)
	account := models.NewAccount(1000)
	go account.Withdraw(700, done)
	go account.Deposite(500, done)
	<-done
	<-done

	fmt.Printf("New Balance %d\n", account.Balance())
}
