package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	name    string
	left    string
	right   string
	starter bool
	ender   bool
}

func main() {
	f, err := os.Open("input3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	nodes := make(map[string]node)

	var actions []string
	var positions []string
	ln := 1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if ln == 1 {
			actions = strings.Split(line, "")
		} else if line != "" {
			starter := false
			ender := false
			inputs := strings.Split(line, "=")
			directions := strings.Split(strings.Replace(strings.Replace(inputs[1], "(", "", 1), ")", "", 1), ",")

			name := strings.TrimSpace(inputs[0])
			data := strings.Split(name, "")
			if data[2] == "A" {
				positions = append(positions, name)
				starter = true
			}
			if data[2] == "Z" {
				ender = true
			}
			nodes[name] = node{
				name:    name,
				left:    strings.TrimSpace(directions[0]),
				right:   strings.TrimSpace(directions[1]),
				starter: starter,
				ender:   ender,
			}
		}
		ln++
	}

	sum := 0
	toFound := len(positions)
	found := 0
	for true {
		for i, p := range positions {
			for _, a := range actions {
				fmt.Println("---------")
				fmt.Println("Action = ", a)
				found := 0
				n := nodes[p]
				if n.ender {
					found++
					break
				}
				switch a {
				case "R":
					positions[i] = n.right
				case "L":
					positions[i] = n.left
				}
				fmt.Println(positions[i])
			}
			sum++
		}
		if found == toFound {
			break
		}
	}
	fmt.Println(fmt.Sprintf("Sum =  %d", sum))
}
