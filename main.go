package main

import (
	"fmt"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter encoded text: ")
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Printf("error while reading string: %v\n", err)
	// 	return
	// }

	// processTxt(input)
	processTxt("LLRR=")
}

type LHandler struct {
	Index   int
	IsFound bool
}

func processTxt(str string) {
	result := []int{0}

	h := new(LHandler)
	h.IsFound = false
	for i := range str {

		val := 0
		switch str[i] {
		case 'L':
			if !h.IsFound {
				h.Index = i
				h.IsFound = true
			}

			if h.IsFound {
				for j := i; j > h.Index-1; j-- {
					result[j]++
				}
			}

			result = append(result, val)
		case 'R':
			h.IsFound = false
			val = result[i] + 1
			result = append(result, val)
		case '=':
			h.IsFound = false
			result = append(result, result[i])
		}
		fmt.Println(string(str[i]))
		fmt.Println(result, fmt.Sprintf("i: %d", i))
	}

	fmt.Println("===========================")
	fmt.Println(result)
}
