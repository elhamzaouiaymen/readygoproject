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
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
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

	//datatype
	rect := Rectangle{50, 100}
	//fmt.Println("rect attributes are :\n", rect.leftx, "\n", rect.topY, "\n", rect.height, "\n", rect.width, "area", rect.area())

	trian := Triangle{3, 9}
	circ := Circle{5}

	fmt.Println("triangle's area is : ", getArea(trian), "\n Circle's area :", getArea(circ), "\n Rectangle's area is :", getArea(rect))

	//advanced part : I/O files
	fmt.Println("we began the advanced part of golang programming")
	mystring := "here we go"
	fmt.Println(mystring, "contains : ", strings.Count(mystring, "e"), " e")
	fmt.Println(strings.Contains(mystring, "we"))
	fmt.Println(strings.Replace(mystring, "e", "x", 2))
	fmt.Println(strings.Index(mystring, "we"))

	//how to split a string
	csvString := "1,2,5,8,6,5"
	fmt.Println("splitted string : \n", strings.Split(csvString, ","))

	//how to sort a list of string and nums
	liststrings := []string{"a", "c", "b"}
	sort.Strings(liststrings)
	fmt.Println("list of strings sorted", liststrings)

	listnums := []string{"7", "1", "6"}
	sort.Strings(listnums)
	fmt.Println("list of nums sorted", listnums)

	//io/outil file

	file, err := os.Create("elbenz.txt")
	//if error is happened log the error
	//the equivalent of null in golang is nil
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("this is an e-class 2016")

	stream, err := ioutil.ReadFile("elbenz.txt")
	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)
	fmt.Println(readString)

	//casting and conversion from a type to another

	randval1 := 125.23
	randval2 := "566.02"
	randval3 := "5555"
	randval4 := 22

	fmt.Println("from int to float casting :", float64(randval4))
	fmt.Println("from float to int casting : ", int(randval1))

	//from string to int
	newInt, _ := strconv.ParseInt(randval2, 0, 64)
	fmt.Println(newInt)

	//from string to float
	newFloat, _ := strconv.ParseFloat(randval3, 64)
	fmt.Println(newFloat)

	//in this last part we gonna make an http server
	//we import from the package net the package http

	http.HandleFunc("/", handler1)
	http.HandleFunc("/mercedes", handler2)
	//we definea listener and the port
	http.ListenAndServe(":8080", nil)
}

//defining handlers of the server

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this the best or nothing")
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

//in this part we gonna learn some basics of classes and polymorphism
//defining datatype
/*type Rectangle struct {
	leftx  float64
	topY   float64
	height float64
	width  float64
}*/
//defining a method for a datatype
/* func (meth_receiver *type) method_name(attribute if has)outputtype{
	method body
}*/

//polymorphism principles
type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base    float64
	hauteur float64
}

func (rect Rectangle) area() float64 {
	return rect.height * rect.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (t Triangle) area() float64 {
	return (t.base * t.hauteur) / 2
}

func getArea(sha Shape) float64 {
	return sha.area()
}

/*to implement a method of an interface
func (typename type) method_name() outputtype{
	method body
}*/
