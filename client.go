package main

import (
	"fmt"
	"net/rpc"
)

type Grade struct {
	Student string
	Subject string
	Grade   float64
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	var op int64
	for {
		fmt.Println("1) Add grade")
		fmt.Println("2) Student prom")
		fmt.Println("3) General prom")
		fmt.Println("4) Subject prom")
		fmt.Println("0) Exit")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var name string
			fmt.Print("Student: ")
			fmt.Scanln(&name)
			var subject string
			fmt.Print("Subject: ")
			fmt.Scanln(&subject)
			var grade float64
			fmt.Print("Grade: ")
			fmt.Scanln(&grade)
			grade_args := &Grade{name,subject,grade}
			var result string
			err = c.Call("Server.AddGrade", grade_args, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.AddGrade: ", result)
			}
		case 2:
			var name string
			fmt.Print("Student: ")
			fmt.Scanln(&name)

			var result float64
			err = c.Call("Server.StudentProm", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.StudentProm", name, " : ", result)
			}
		case 3:
			var result float64
			err = c.Call("Server.GralProm", result, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.GralProm: ", result)
			}
		case 4:
			var subject string
			fmt.Print("Subject: ")
			fmt.Scanln(&subject)

			var result float64
			err = c.Call("Server.SubjectProm", subject, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.SubjectProm", subject, " : ", result)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}
