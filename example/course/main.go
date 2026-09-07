package main

import (
	"errors"
	"fmt"
	"os"

	courseSdk "github.com/dasilro/go_course_sdk/course"
)

func main() {
	courseTrans := courseSdk.NewHttpClient("http://localhost:8082", "")

	course, err := courseTrans.Get("9833c349-4aa6-4bf4-a73a-ebef94103bdd")
	if err != nil {
		if errors.As(err, &courseSdk.ErrNotFound{}) {
			fmt.Println("Not found: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal Server Error:", err.Error())
		os.Exit(1)
	}

	fmt.Println(course)
}
