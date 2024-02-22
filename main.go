package main

import (
	"fmt"
	"os"
	"time"
)

func visualizer(D [][]int, word1, word2 string) {
	for y := range len(word2) + 1 {
		for x := range len(word1) + 1 {
			if y == 0 && x == 0 {
				fmt.Printf("  |")
			} else if y == 0 {
				fmt.Printf("%s ", string(word1[x-1]))
			} else if x == 0 {
				fmt.Printf("%s |", string(word2[y-1]))
			} else {
				fmt.Printf("%d ", D[x][y])
			}
			if x == len(word1) {
				fmt.Println()
			}
		}
	}
}

func calculate(calc [][]bool, D [][]int, x int, y int, word1 string, word2 string) {

	if x == 0 {
		D[x][y] = y
		calc[x][y] = true
		return
	}

	if y == 0 {
		D[x][y] = x
		calc[x][y] = true
		return
	}

	if word1[x-1] == word2[y-1] {
		lastCalc := calc[x-1][y-1]
		if !lastCalc {
			calculate(calc, D, x-1, y-1, word1, word2)
		}
		D[x][y] = D[x-1][y-1]
		calc[x][y] = true
		return
	}

	insertCalc := calc[x][y-1]
	if !insertCalc {
		calculate(calc, D, x, y-1, word1, word2)
	}
	insert := D[x][y-1]

	replaceCalc := calc[x-1][y-1]
	if !replaceCalc {
		calculate(calc, D, x-1, y-1, word1, word2)
	}
	replace := D[x-1][y-1]

	deleteCalc := calc[x-1][y]
	if !deleteCalc {
		calculate(calc, D, x-1, y, word1, word2)
	}
	delete := D[x-1][y]

	D[x][y] = min(insert+1, replace+1, delete+1)
	calc[x][y] = true
}

func EditDistance(word1 string, word2 string) int {
	x := len(word1)
	y := len(word2)

	if x <= 0 && y <= 0 {
		return 0
	} else if x <= 0 || y <= 0 {
		return x + y
	}

	D := make([][]int, x+1)
	calc := make([][]bool, x+1)

	for i := range calc {
		calc[i] = make([]bool, y+1)
		D[i] = make([]int, y+1)
	}

	// visualizer(D, word1, word2)
	calculate(calc, D, x, y, word1, word2)

	if len(word1) <= 50 && len(word2) <= 50 {
		visualizer(D, word1, word2)
	}
	return D[x][y]
}

func tests() {
	dna1000_1 := "ggtggttggtcagaaccgtcccgtatgttcataactaggcactagtaccggggccaggacgggagtgcaatagcaagcccttatcaaaaccgtcgcgctaaccacgcaaagatacggtatcacatatgccaagaattggggatgggtattagaatgacctaggtcaacactccttgttagagcgagtggcgtgtgacgtaccacgtcgtacttaactagatcgcttaaagccccgatgtggccacttggaggattcaaaggccctaatgatcctcacacgctaccgaggttgacggcgcttcttgaaaacacaaatttcttggtgacatacgcctacgactcattgtcgtacttttcgtctatcaccaagcgaaacctcccccacttaaccatctatgcgaattgttattcggcaccgccaccgtggaaacccgtcataaaaggaccatgccaaattggtttcatcgacaaagtccattaagttcgatataaacttatttgcagctcgcaagataaaaggctatgtccatgccatgttcggcgcacctctcctcgcgctgtaggacgcaacgttcgttcataatcgagtagtcctgctgcactgatggagccatccattgcagcgtcagcgcttcgactccggcccgctcatcgctagttagctatccgtacagtatcagaacatcttggggcttagtaaagtggtcggatccggtgttttttgcagtagcaaatggtttctaaaaacctgtcggcttttagattttacgatccctcgagtcttcgacttcttcgatcgtcacggtcctaagtgtcttgcgaccaggtatcagtgggcgcgtgcactttttgagttcgaagttagcgagcgtccctagaagtatccaattgcacctgttgaaaggaggaatatcctcaaattttaggaccttttagccttacccatactcgtggtagaagcattcggtcgtcggttagagttccattagtaataaatcgc"
	dna1000_2 := "aagtggggcagtggctcacacccatcatttggtgctaggcaatatatggtgaaaattcggtgcgggaagccaatcttgatgcagtcaactaaggtaaggctggcatgactagaaagcgttgacggcactacgtccatacatgcagccagtcgagataagtacttatacggttaccatctatgaaccagaccggatgtaatccagattaaacgggattgggtctttgctttcacccgggcttggttagagacagcaccctttcctgattacacctcgcataaaaccctagattttaggacactggacggtcttttcgcgatgcttttggtgtgcgccggacaaaggttataaatggtgtctctagtgaaggacggtttagtcgatgccaacgtgtatcaatgtagggcacggccggaggtctcgctggtattgcatttcgggatccgatgaatatcgtacgatagtagtgtccacagaacctttgtgtagttatacgcgctgtggtaccgatggccatagccgtagtggtccgctttgtgtgctgcgctacctgccggccctttaagggaacacgtgtaagccagttaactgagttcctaacccccaagagcatcgctccgatgtgttacgtactctcgtcactccagagatgcacgctcgactagtggtctggcagttatcggcttcgtgaagtatcgcaagtcttgacgttggactttgggtattataaccaatgtcgtgacgatatcgtgtcctagcgggctacctacatgcgggcggtaatatcgcgaatggccgcccacaagagtagaatcagttttcgtgtcctccttggttttcctgcatcgaatgttagctaggctgggacatcaatatatgtttcgcgcgtctttggtagcttccactcatctaaacattatcctggcctactgaaagtaatttccagggaccaccaacgggtccctggccgtattacccagcatcgtttctcccaggtcaa"
	dna100_1 := dna1000_1[:100]
	dna100_2 := dna1000_2[:100]

	start := time.Now()
	fmt.Printf("EditDistance(): n = %d | %d", 100, EditDistance(dna100_1, dna100_2))
	end := time.Now()
	fmt.Printf(" in %v\n", end.Sub(start))

	start = time.Now()
	fmt.Printf("LogiDistance(): n = %d | %d", 100, LogiDistance(dna100_1, dna100_2))
	end = time.Now()
	fmt.Printf(" in %v\n", end.Sub(start))

	start = time.Now()
	fmt.Printf("EditDistance(): n = %d | %d", 1000, EditDistance(dna1000_1, dna1000_2))
	end = time.Now()
	fmt.Printf(" in %v\n", end.Sub(start))

	start = time.Now()
	fmt.Printf("LogiDistance(): n = %d | %d", 1000, LogiDistance(dna1000_1, dna1000_2))
	end = time.Now()
	fmt.Printf(" in %v\n", end.Sub(start))

}

func main() {
	if len(os.Args) == 3 {
		fmt.Printf("Program arguments %s and %s accepted\n", os.Args[1], os.Args[2])
		word1 := os.Args[1]
		word2 := os.Args[2]

		start := time.Now()
		fmt.Printf("EditDistance(%s, %s) | %d", word1, word2, EditDistance(word1, word2))
		end := time.Now()
		fmt.Printf(" in %v\n", end.Sub(start))

		start = time.Now()
		fmt.Printf("LogiDistance(%s, %s) | %d", word1, word2, LogiDistance(word1, word2))
		end = time.Now()
		fmt.Printf(" in %v\n", end.Sub(start))
	} else {
		tests()
	}

}
