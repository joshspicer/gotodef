package main

func main() {
    localOne()
    localTwo()
    localThree()
    FunctionA()
    FunctionB()
    FunctionC()
    FunctionD()
}
 
func ignored() {
    println("Nobody calls me!")
}

func localOne() {
    println("local one is called")
}

func localTwo() {
    println("local two is called")
}

func localThree() {
    println("local three is called")
}