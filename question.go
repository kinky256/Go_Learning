/*Write a program that prints the numbers from 1 to 100, but for multiples of
three, print “Fizz” instead of the number, and for the multiples of five, print
“Buzz.” For numbers that are multiples of both three and five, print “FizzBuzz.”*/

package main

import "fmt"

func main(){
	var i int
	for i=0;i<=100;i++{
		if i%3==0{
			fmt.Println("Fizz")
		}else if i%5==0{
			fmt.Println("buzz")
		}
	}
}

