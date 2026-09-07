package main

import (
	"errors"
	"fmt"
	"os"

	userSdk "github.com/dasilro/go_course_sdk/user"
)

func main() {
	userTrans := userSdk.NewHttpClient("http://localhost:8081", "")

	user, err := userTrans.Get("35a43361-e1fe-4357-823b-e7f35aa7c720")
	if err != nil {
		if errors.As(err, &userSdk.ErrNotFound{}) {
			fmt.Println("Not found: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal Server Error:", err.Error())
		os.Exit(1)
	}

	fmt.Println(user)
}
