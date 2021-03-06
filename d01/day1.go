package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	r := bufio.NewReader(f)
	data, err := r.ReadString('\n')
	check(err)
	data = string(data[0 : len(data)-1])

	part1(data)
	part2(data)

}

func part1(input string) {
	total := 0
	for i, c := range input {
		x := int(c)
		y := int(input[(i+1)%len(input)])
		if x == y {
			total += x - 48
		}
	}
	fmt.Println(total)
}

func part2(input string) {
	total := 0
	for i, c := range input {
		x := int(c)
		y := int(input[(i+(len(input)/2))%len(input)])
		if x == y {
			total += x - 48
		}
	}
	fmt.Println(total)
}
