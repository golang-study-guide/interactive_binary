package main

// need to first run: go get github.com/fatih/color
import (
	"bufio" //https://golang.org/pkg/bufio/
	"fmt"
	"os"
	"reflect"
)

func main() {
	var f *os.File // "f" is a pointer to memory location that will hold a variable of the type os.File.
	// os.File is a path to a file on the fileystem
	f = os.Stdin // here we chose the Stdin as the the filepath to use. In Linux, stdin, stdout, and stderr are all essentially files behind the scenese

	defer fmt.Println("\nYou did ctrl+d, so exitiing now")
	defer f.Close() // https://tour.golang.org/flowcontrol/12
	// defer means that this line is triggered once the main function finishes running.
	// we are applying the Close method for the f object.

	// This is a bit like opening the file to start writing to it
	scanner := bufio.NewScanner(f)

	fmt.Print("Please enter an input: ")

	// https://golang.org/pkg/bufio/#Scanner.Scan
	// xxx..Scan() is effectively the bash equivalent of of the 'read' command
	for scanner.Scan() {
		fmt.Println(reflect.TypeOf(scanner.Text()))
		fmt.Println(">>>", scanner.Text())
		fmt.Print("Please enter another input: ")
	}

}
