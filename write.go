package main

import (
	"log"
	"os"
	"strconv"
)

var c []int

func arr() {

	for i := 1; i < 100; i++ {
		if i%5 == 0 {
			loadarr(i)
		}
	}

}

func loadarr(a int) {

	c = append(c, a)
}

func writeFile() {

	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatalf("failed to creating file: %s", err)
	}

	defer file.Close()

	// for i := 0; i < len(c); i++
	for c := range c {
		//fmt.Println(c[i])
		file.WriteString(strconv.Itoa(c) + "\n")
	}

}

func main() {
	arr()
	// fmt.Println(c)
	writeFile()

}
