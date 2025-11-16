package main

import "fmt"

// Queston : 1 = Create a program that prints your:Name,Age,Favourite language Your goal (ex: become a backend engineer)

// func main() {
// 	name := "Shivam"
// 	Age := 10
// 	Favourite_lang := "Go"
// 	Future_Goal := "Backend engineer"

// 	fmt.Println("My name is :", name)
// 	fmt.Println("My age is :", Age)
// 	fmt.Println("My Favourite_lang is :", Favourite_lang)
// 	fmt.Println("My Future goal is :", Future_Goal)

// }

// Input Fields

// func main () {
// 	var name string
// 	var age int
// 	var Favourite_lang string
// 	var Future_Goal string

// 	// fmt.Println("Please Enter your details...")
// 	fmt.Println("Please Enter your name...")
// 	fmt.Println("Please Enter your age...")
// 	fmt.Println("Please Enter your Favourite_lang...")
// 	fmt.Println("Please Enter your Future_Goal...")

// 	fmt.Scan(&name)
// 	fmt.Scan(&age)
// 	fmt.Scan(&Favourite_lang)
// 	fmt.Scan(&Future_Goal)

// 		fmt.Println("My name is :", name)
// 		fmt.Println("My age is :", age)
// 		fmt.Println("My Favourite_lang is :", Favourite_lang)
// 		fmt.Println("My Future goal is :", Future_Goal)
// }

// Take 2 numbers from user and print: Sum, Difference, Multiplication

// func main () {

// 	var num1 int
// 	var num2 int

// 	fmt.Println("Please Enter a numbers one by one")

// 	fmt.Scan(&num1)
// 	fmt.Scan(&num2)

// 	sum := num1+num2
// 	Diff := num1-num2
// 	Mul := num1*num2

// 	fmt.Println("sum of two number is", sum)
// 	fmt.Println("Diff of two number is", Diff)
// 	fmt.Println("Multiplication of two number is", Mul)
// }



func greet (name string) {

	fmt.Println("this is my name", name)

}

func main () {
	greet("Shivm")

}