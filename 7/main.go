package main

import "fmt"

type Student struct {
    Name   string
    Mentor *string // наставник может быть не назначен
}

func printMentor(s Student) {
	if s.Mentor != nil{
		fmt.Println(*s.Mentor)
	}else{
		fmt.Println("наставник пока не назначен")
	}
}

func main(){
	m := "Назиров Расул"
	
// присвой адрес переменной m в поле Mentor
	students:=Student{
		Name:"Imran",
		Mentor: &m,
	}
	printMentor(students)
	students.Mentor = &m
	printMentor(students)
}