package main

import "fmt"

func main() {

	FirstName := 1

	yourName := "Ayesha"
	var sentence = "I love you"

	var secondName string = "Maria"

	var Lastname string
	Lastname = "Ulfa"

	var var1 string = fmt.Sprintf("Heloo %d %s %s", FirstName, secondName, Lastname)
	fmt.Println(var1)
	fmt.Println("Hello", FirstName, Lastname+"!")
	fmt.Println(yourName)
	fmt.Println(sentence)

	//Multivariable

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fmt.Println(second, third, first)

	var fourth, fifth, sixth = "Empat", "Lima", "Enam"
	fmt.Printf("List angkanya adalah %s %s %s \n", fourth, fifth, sixth)

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "Hello"
	fmt.Println(one, isFriday, twoPointTwo, say)

	//Variable Underscore _

	name, _ := "Rauf", "Abdul"

	fmt.Printf("%s \n", name)

	//Variable using keyword New (for pointer)

	a := new(string)
	
	fmt.Println(*a)

}
