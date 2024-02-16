package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	name  string
	left  string
	right string
}

func main() {
	f, err := os.Open("input2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	nodes := make(map[string]node)

	var actions []string
	ln := 1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if ln == 1 {
			actions = strings.Split(line, "")
		} else if line != "" {
			inputs := strings.Split(line, "=")
			directions := strings.Split(strings.Replace(strings.Replace(inputs[1], "(", "", 1), ")", "", 1), ",")
			nodes[strings.TrimSpace(inputs[0])] = node{
				name:  strings.TrimSpace(inputs[0]),
				left:  strings.TrimSpace(directions[0]),
				right: strings.TrimSpace(directions[1]),
			}
		}
		ln++
	}

	sum := 0
	start := "AAA"
	end := "ZZZ"
	position := start

	found := false
	for true {
		for _, a := range actions {
			n := nodes[position]
			if n.name == end {
				found = true
				break
			}
			switch a {
			case "R":
				position = n.right
			case "L":
				position = n.left
			}
			sum++
		}
		if found {
			break
		}
	}
	fmt.Println(fmt.Sprintf("Sum =  %d", sum))
}
