package main

import (
	"net/http"
	"universitySignalTransformation/pkg/lab1"
	"universitySignalTransformation/pkg/lab2"
	"universitySignalTransformation/pkg/lab3"
)

func main() {
	http.HandleFunc("/lab1/zad1", lab1.DrawExercise1)
	http.HandleFunc("/lab1/zad2", lab1.DrawExercise2)
	http.HandleFunc("/lab2/zad1", lab2.DrawExercise1)
	http.HandleFunc("/lab2/zad2", lab2.DrawExercise2)
	http.HandleFunc("/lab3/zad1", lab3.DrawExercise1)
	http.HandleFunc("/lab3/zad2", lab3.DrawExercise2)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
