package tasks

import "fmt"

func SendMail(msg string) (string, error) {
	result := fmt.Sprintf("the mail is sent to %s from test@gmail.com\n", msg)
	fmt.Println(result)
	return result, nil
}
