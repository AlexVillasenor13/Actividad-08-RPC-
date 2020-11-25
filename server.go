package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

type Grade struct {
	Student string
	Subject string
	Grade   float64
}

type Server struct {
	Subjects map[string]map[string]float64
	Students map[string]map[string]float64
}

func (this *Server) AddGrade(grade *Grade, reply *string) error {
	for student, _ := range this.Subjects[grade.Subject] {
		if student == grade.Student {
			return errors.New("Existing grade")
		}
	}
	if this.findSubject(grade.Subject) {
		this.Subjects[grade.Subject][grade.Student] = grade.Grade
	} else {

		student := make(map[string]float64)
		student[grade.Student] = grade.Grade
		this.Subjects[grade.Subject] = student

	}
	if this.findStudent(grade.Student) {
		this.Students[grade.Student][grade.Subject] = grade.Grade
	} else {
		subject := make(map[string]float64)
		subject[grade.Subject] = grade.Grade
		this.Students[grade.Student] = subject
	}
	s := fmt.Sprintf("%.2f", grade.Grade)
	*reply = "Added Grade: " + grade.Subject + "-" + grade.Student + ": " + s
	return nil
}

func (this *Server) findStudent(name string) bool {
	for student, _ := range this.Students {
		if student == name {
			return true
		}
	}
	return false
}
func (this *Server) findSubject(name string) bool {
	for subject, _ := range this.Subjects {
		if subject == name {
			return true
		}
	}
	return false
}

func (this *Server) StudentProm(name string, reply *float64) error {
	total_grades := 0.0
	total_subjects := 0.0
	for _, grade := range this.Students[name] {
		total_grades += grade
		total_subjects += 1
	}
	if total_subjects > 0 {
		*reply = total_grades / total_subjects
		return nil
	} else {
		return errors.New("Wrong student")
	}
}

func (this *Server) SubjectProm(name string, reply *float64) error {
	total_grades := 0.0
	total_students := 0.0
	for _, grade := range this.Subjects[name] {
		total_grades += grade
		total_students += 1
	}
	if total_students > 0 {
		*reply = total_grades / total_students
		return nil
	} else {
		return errors.New("Wrong subject")
	}
}

func (this *Server) GralProm(reply_value float64, reply *float64) error {
	total_grades_students := 0.0
	total_students := 0.0
	for student, _ := range this.Students {
		prom := 0.0
		err := this.StudentProm(student, &prom)
		if err != nil {
			return err
		}
		total_grades_students += prom
		total_students += 1
	}
	if total_students > 0 {
		*reply = total_grades_students / total_students
		return nil
	} else {
		return errors.New("Empty students")
	}
}

func server(Subjects, Students map[string]map[string]float64) {
	new_server := new(Server)
	new_server.Students = Students
	new_server.Subjects = Subjects
	rpc.Register(new_server)
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	Subjects := make(map[string]map[string]float64)
	Students := make(map[string]map[string]float64)
	go server(Subjects, Students)

	var input string
	fmt.Scanln(&input)
}
