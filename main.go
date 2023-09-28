package main

import (
	"github.com/Ticolls/remail/client"
	"github.com/Ticolls/remail/email"
)

func main() {
	email.Init()
	client.Init()

	// for _, task := range tasks {
	// 	fmt.Printf("tarefa: %s ; due: %v\n", task.Content, task.Due)
	// 	fmt.Printf("date: %v\n", task.Due.Date)
	// 	fmt.Printf("isRecurring: %v\n", task.Due.IsRecurring)
	// 	fmt.Printf("datetime: %v\n", task.Due.Datetime)
	// 	fmt.Printf("String: %v\n", task.Due.String)
	// 	fmt.Printf("timezone: %v\n", task.Due.Timezone)
	// }
}
