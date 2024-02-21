package main

import (
	"fmt"
	"math"
	"time"
)

var D [][]int = [][]int{}

func Distance(word1 string, word2 string) int {
	distance := 0
	for i := 0; i < min(len(word1), len(word2)); i++ {
		if word1[i] != word2[i] {
			distance++
		}
	}

	return distance + int(math.Abs(float64(len(word1)-len(word2))))
}

func calculate(calc [][]bool, D [][]int, x int, y int, word1 string, word2 string) {
	if x == 0 && y == 0 {
		D[x][y] = 0
		calc[x][y] = true
		return
	}

	if x == 0 || y == 0 {
		D[x][y] = x + y
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
	D := make([][]int, x)
	calc := make([][]bool, x)

	for i := range calc {
		calc[i] = make([]bool, y)
		D[i] = make([]int, y)
	}

	x--
	y--

	if x == 0 && y == 0 {
		return 0
	} else if x == 0 || y == 0 {
		return x + y
	}

	if word1[x] == word2[y] {
		if !calc[x-1][y-1] {
			calculate(calc, D, x-1, y-1, word1, word2)
		}
		return D[x-1][y-1]
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

	return min(insert+1, replace+1, delete+1)

	// return min(EditDistance(word1[:x-1], word2[:y])+1, EditDistance(word1[:x-1], word2[:y-1])+1, EditDistance(word1[:x], word2[:y-1])+1)
}

func main() {
	// word1 := "test"
	// word2 := "trest"
	// fmt.Printf("Distance(%s, %s) = %d\n", word1, word2, Distance(word1, word2))
	// fmt.Printf("Distance(%s, %s) = %d\n", word1, word2, Distance(word1, word2))

	// fmt.Printf("EditDistance(%s, %s) = %d\n", word1, word2, EditDistance(word1, word2))

	// // provided by https://www.bioinformatics.org/sms2/random_dna.html
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
