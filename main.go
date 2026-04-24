package main

import (
	"net/http"
	"universitySignalTransformation/pkg/lab1"
	"universitySignalTransformation/pkg/lab2"
	"universitySignalTransformation/pkg/lab3"
	"universitySignalTransformation/pkg/lab4"
	"universitySignalTransformation/pkg/lab5"
)

func main() {
	//Test value for lab5 bandwidth function
	lab5.Bandwidth(lab4.SignalGenerationExerise1(1.5, 2000, 50, 10, 25.5, "Z_A"), 3)

	//Lab1 endpoints
	http.HandleFunc("/lab1/zad1", lab1.DrawExercise1)
	http.HandleFunc("/lab1/zad2", lab1.DrawExercise2)

	//Lab2 endpoints
	http.HandleFunc("/lab2/zad1", lab2.DrawExercise1)
	http.HandleFunc("/lab2/zad2", lab2.DrawExercise2)

	//Lab3 endpoints
	http.HandleFunc("/lab3/zad1", lab3.DrawExercise1)
	http.HandleFunc("/lab3/zad2", lab3.DrawExercise2)

	//Lab4 endpoints
	http.HandleFunc("/lab4/zad1/Za", lab4.DrawExercise_Za)
	http.HandleFunc("/lab4/zad1/Zf", lab4.DrawExercise_Zf)
	http.HandleFunc("/lab4/zad1/Zp", lab4.DrawExercise_Zp)

	//Lab5 endpoints
	http.HandleFunc("/lab5/zad1/Ma", lab5.DrawExercise_Ma)
	http.HandleFunc("/lab5/zad1/Mf", lab5.DrawExercise_Mf)
	http.HandleFunc("/lab5/zad1/Mp", lab5.DrawExercise_Mp)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}

}
