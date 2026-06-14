package main

import (
	"net/http"
	"universitySignalTransformation/pkg/lab1"
	"universitySignalTransformation/pkg/lab12"
	"universitySignalTransformation/pkg/lab13"
	"universitySignalTransformation/pkg/lab2"
	"universitySignalTransformation/pkg/lab3"
	"universitySignalTransformation/pkg/lab4"
	"universitySignalTransformation/pkg/lab5"
	"universitySignalTransformation/pkg/lab6"
	"universitySignalTransformation/pkg/lab8"
	"universitySignalTransformation/pkg/lab9"
)

func main() {
	//encoded := lab11.HammingCode("test")
	//lab11.HammingDecode(encoded)

	/*
		lab5.SaveAllExercise1Data("pkg/lab5/test_dla_3.txt", 3)
		lab5.SaveAllExercise1Data("pkg/lab5/test_dla_6.txt", 6)
		lab5.SaveAllExercise1Data("pkg/lab5/test_dla_10.txt", 10)


		lab7.SaveAllExercise1Data("pkg/lab7/test_dla_3.txt", 3)
		lab7.SaveAllExercise1Data("pkg/lab7/test_dla_6.txt", 6)
		lab7.SaveAllExercise1Data("pkg/lab7/test_dla_10.txt", 10)

	*/
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

	//Lab6 endpoints
	http.HandleFunc("/lab6/zad1", lab6.DrawExercise1)
	http.HandleFunc("/lab6/zad2", lab6.DrawExercise2)

	//Lab8 endpoints
	http.HandleFunc("/lab8/zad1", lab8.DrawDemodulator)

	//Lab9 endpoints
	http.HandleFunc("/lab9/zad1", lab9.DrawDemodulator)

	//Lab12 endpoints
	http.HandleFunc("/lab12/zad1/Hamming_7_4", lab12.DrawExercise1Hamming_7_4)
	http.HandleFunc("/lab12/zad1/Hamming_15_11", lab12.DrawExercise1Hamming_15_11)
	http.HandleFunc("/lab12/zad2/Hamming_7_4", lab12.DrawExercise2Hamming_7_4)
	http.HandleFunc("/lab12/zad2/Hamming_15_11", lab12.DrawExercise2Hamming_15_11)

	//Lab13 endpoints
	http.HandleFunc("/lab13/zad1/Hamming_7_4/config1", lab13.DrawExercise13Part1Hamming_7_4_Config1)
	http.HandleFunc("/lab13/zad1/Hamming_7_4/config2", lab13.DrawExercise13Part1Hamming_7_4_Config2)

	http.HandleFunc("/lab13/zad1/Hamming_15_11/config1", lab13.DrawExercise13Part1Hamming_15_11_Config1)
	http.HandleFunc("/lab13/zad1/Hamming_15_11/config2", lab13.DrawExercise13Part1Hamming_15_11_Config2)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}

}
