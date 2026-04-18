package main

import (
	"net/http"
	"universitySignalTransformation/pkg/lab1"
)

func main() {
	http.HandleFunc("/lab1/zad1", lab1.DrawExercise1)
	http.HandleFunc("/lab1/zad2", lab1.DrawExercise2)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
