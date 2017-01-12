package main

/* to import a single library we use : import "lib name"*/
/* to import a list of libs we can do it in this way :
import (
      "lib1"
      "lib2"
      "lib3"
)
*/

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello world")

	//declaring a variable
	var age int = 20

	//printing float numbers
	const pi float32 = 3.145558556688
	fmt.Println("%.3f \n", pi)

	//declaring list of variables
	var (
		varA = 2
		varB = 3.1444444
		varC = 2222222
	)

	fmt.Println("I am aymen hamzaoui i've", age, "\n", varA, "+", varB, " =", varA+varC)

	//numbers formatting (binary,decimal,char,hexadecimal,exponential)
	fmt.Printf("%d \n", 100)
	fmt.Printf("%b \n", 100)
	fmt.Printf("%c \n", 100)
	fmt.Printf("%x \n", 100)
	fmt.Printf("%e \n", 100)

	//declaring some boolean variables
	//var isBig bool =false

	//boolean operations
	// AND LOGIC OPERATOR
	fmt.Println("true and true =", true && true)
	fmt.Println("true and false =", true && false)
	fmt.Println("false and true =", false && true)
	fmt.Println("false and false =", false && false)

	fmt.Println("*************************************")
	//OR LOGIC OPERATOR
	fmt.Println("true or false =", true || false)
	fmt.Println("false or true =", false || true)
	fmt.Println("false or false =", false || false)
	fmt.Println("ture or ture =", true || true)

	fmt.Println("*************************************")

	//PLAYING WITH SOME LOOPS
	//1- FOR LOOP
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	fmt.Println("*************************************")
	/*2- WHILE LOOP : for is go's while we haven't :
	  while(condition) ==> do something*/
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//3- DECLARING FOREVER LOOP WE USE FOR ALSO BUT IN THIS WAY
	/*for{
	  fmt.Println("coding until you die !")
	}*/

	//FLOW CONTROL STATEMENT IF,ELSE,SWITCH,DEFER ...
	/*1- IF STATEMENT*/
	fmt.Println("*******************simple if statement***************************")
	x := 6
	if x > 0 {
		fmt.Println("sqrt(", x, ")=", math.Sqrt(float64(x)))
	}

	fmt.Println("*******************if else***************************")
	//2- if(condition) else if (condition)
	myAge := 5
	if myAge > 18 {
		fmt.Println("you can pass the drive test")
	} else {
		fmt.Println("you can't pass this test .. sorry for that !")
	}

	//3- switch STATEMENT
	fmt.Println("********************switch statement************************")
	switch myAge {
	case 18:
		fmt.Println("you are 18 years old !")
	case 20:
		fmt.Println("you are 20 years old !")
	case 25:
		fmt.Println("you are 25 years old !")
	case 35:
		fmt.Println("you are 35 years old !")
	default:
		fmt.Println("go have fun ! ")
	}

	//discovering a new stuff called defer
	fmt.Println("******************DEFER FUNCTION*****************")
	//defer fmt.Println("hamzaoui")
	//defer  fmt.Println("lamine")
	/* defer function assure functions scheduling. we use it to schedule the execution of functions. its principle is based on LIFO scheduling algorithm*/
	fmt.Println("aline")
	fmt.Println("lina")
	fmt.Println("kenza")
	fmt.Println("aymen")

	/*playing with tables*/
	fmt.Println("*************arrays declarations, operations**************")

	var table [6]int

	table[0] = 25
	table[1] = 33
	table[3] = 65
	table[4] = 89
	table[5] = 277

	for i := 0; i < len(table); i++ {
		fmt.Println("T[", i, "] =", table[i])

	}

	//there is another way to declare an array
	//for example
	array1 := [6]float64{2, 5, 6, 8, 9, 7}

	fmt.Println("array[", 2, "]", array1[2])

	//there is another method to show all the array's content
	fmt.Println("array1 with indexes :")
	for i, value := range array1 {
		fmt.Println(value, i)
	}
	//to avoid showing the indexes you can do this
	fmt.Println("array1 without indexes :")
	for _, value := range array1 {
		fmt.Println(value)
	}

	fmt.Println("slicing arrays")
	//learning slicing arrays
	numSlice := [5]int{5, 8, 9, 6}
	numSlice2 := numSlice[2:4]
	fmt.Println("numSlice2[2]=", numSlice2[1])
	fmt.Println(numSlice[:3]) //this takes a range of the array

	//you can declare an empty array : make([]typecontent,minsize,maxsize)
	numSlice4 := make([]int, 5, 10)
	copy(numSlice4, numSlice2)

	fmt.Println(numSlice4)

	//appending values to a slice
	numSlice2 = append(numSlice2, 0, 6, 9)
	fmt.Println("numslice2 after appending :\n", numSlice2)

	/*in this part , we gonna learn how
	to make maps or dictionnaries in golang
	PS: a map is a collection of data identified with
	what we call key*/
	fmt.Println("**************************************************")
	//defining a map
	myMap := make(map[string]int)
	myMap["key1"] = 42
	myMap["key2"] = 5
	myMap["key3"] = 423
	myMap["key4"] = 43
	myMap["key5"] = 422
	myMap["key6"] = 421
	fmt.Println(myMap)
	fmt.Println("myMap's length is", len(myMap))

	//to delete an item from a map
	delete(myMap, "key3")
	fmt.Println(myMap)

	//familizarizing with functionss
	fmt.Println("double of 2 =", double(2))

	//factorial
	fmt.Println("factorial of 5 = ", factorial(5))

	//calling panic
	panicit()

	//pointer notion
	a := 3
	changeXValnow(&a)
	fmt.Println("the new value of a is", a)

}

func double(number int) int {
	return number * 2
}

//recrusivity
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//exceptions and catching
func panicit() {

	//we define exception error text
	defer func() {
		fmt.Println(recover())
	}()
	panic("something is wrong ")
}

//lets have a look how pointer works in golang
func changeXValnow(x *int) {
	*x = 2
}
