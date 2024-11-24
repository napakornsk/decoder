package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter encoded text: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error while reading string: %v\n", err)
		return
	}

	input = strings.Trim(input, "\n,.")

	fmt.Println(input, ":", processTxt(input))
	fmt.Println("==RLL: ", processTxt("==RLL"))
	fmt.Println("LLRR=: ", processTxt("LLRR="))
	fmt.Println("=LLRR: ", processTxt("=LLRR"))
	fmt.Println("RRL=R: ", processTxt("RRL=R"))
}
func processTxt(s string) string {
	n := len(s) + 1
	result := make([]int, n)
	// Loop forward
	for i, c := range s {
		if c == 'R' {
			result[i+1] = result[i] + 1
		} else if c == '=' {
			result[i+1] = result[i]
		}
	}

	// Loop backward to modify the L in one direction
	for i := len(s) - 1; i > -1; i-- {
		if s[i] == 'L' {
			// If the current result is less than the previous one that means the L rule hasn't applied yet so
			// we increment it to guarantee it has more value than the previous one
			if result[i] < result[i+1]+1 {
				result[i] = result[i+1] + 1
			}
			if i > 0 && s[i-1] == '=' {
				result[i-1] = result[i]
			}
		}
	}

	sb := ""
	for _, r := range result {
		sb += fmt.Sprint(r)
	}
	return sb
}

// type LHandler struct {
// 	Index     int
// 	LastIndex int
// 	lastChar  rune
// 	IsFound   bool
// 	val       int
// }

// func sort(result []int, txt string, h LHandler) []int {
// 	if h.IsFound {
// 		fmt.Println("try to sort L from index:", h.Index, " to ", h.LastIndex)
// 		for i := h.Index; i < h.LastIndex; i++ {
// 			for next := h.Index + 1; next < h.LastIndex+1; next++ {
// 				if result[h.Index] < result[next] {
// 					fmt.Println("swapping")
// 					temp := result[h.Index]
// 					result[h.Index] = result[next]
// 					result[next] = temp
// 				}
// 			}
// 		}
// 		if h.Index > 0 {
// 			if txt[h.Index-1] == '=' {
// 				result[h.Index-1] = result[h.Index]
// 				result[h.Index-2] = result[h.Index]
// 			}
// 		}
// 		return result
// 	} else {
// 		return result
// 	}
// }

// func processTxt(str string) {
// 	result := []int{0} // Start with a single 0
// 	h := new(LHandler)
// 	h.IsFound = false

// 	for i := 0; i < len(str); i++ {
// 		fmt.Println("result before logic: ", result)

// 		switch str[i] {
// 		case 'L':

// 			if h.lastChar != 'L' {
// 				h.lastChar = 'L'
// 				h.Index = i
// 				h.val = 0
// 			} else {
// 				fmt.Println("Found another L")
// 				h.LastIndex = i
// 				h.IsFound = true
// 				h.val = result[len(result)-1] + 1
// 			}
// 			result = append(result, h.val)

// 		case 'R':
// 			if h.lastChar != 'R' {
// 				result = sort(result, str, *h)
// 				h.lastChar = 'R'
// 			}
// 			h.val = result[len(result)-1] + 1

// 			// h.IsFound = false
// 			// h.val = result[len(result)-1] + 1
// 			result = append(result, h.val)

// 		case '=':
// 			if h.lastChar != '=' {
// 				result = sort(result, str, *h)
// 				h.lastChar = '='
// 			}
// 			h.IsFound = false
// 			h.val = result[len(result)-1]
// 			result = append(result, h.val)
// 		}
// 		if h.lastChar == 'L' {
// 			sort(result, str, *h)
// 		}
// 		fmt.Println(string(h.lastChar))

// 		// fmt.Println(h)
// 		// fmt.Println(string(str[i]))
// 		fmt.Println(result, fmt.Sprintf("i: %d", i))
// 	}

// 	fmt.Println("===========================")
// 	fmt.Println(result)
// }
