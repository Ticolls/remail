package main

import (
	"fmt"

	"github.com/Ticolls/remail/client"
	"github.com/Ticolls/remail/email"
	"github.com/Ticolls/remail/handler"
)

func main() {
	email.Init()
	client.Init()

	err, tasks := client.GetTasks()

	if err != nil {
		err = email.SendEmail("Erro recuperando as tarefas", err.Error())
		fmt.Println(err.Error())
	}

	fmt.Println("Everything is ok, handling email...")

	handler.Init(tasks)

	// for _, task := range tasks {
	// 	fmt.Printf("tarefa: %s ; due: %v\n", task.Content, task.Due)
	// 	fmt.Printf("date: %v\n", task.Due.Date)
	// 	fmt.Printf("isRecurring: %v\n", task.Due.IsRecurring)
	// 	fmt.Printf("datetime: %v\n", task.Due.Datetime)
	// 	fmt.Printf("String: %v\n", task.Due.String)
	// 	fmt.Printf("timezone: %v\n", task.Due.Timezone)
	// }
}
