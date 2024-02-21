package main

import (
	"fmt"
	"math"
)

var D [][]int = [][]int{}
var calc [][]bool = [][]bool{}

func Distance(word1 string, word2 string) int {
	distance := 0
	for i := 0; i < min(len(word1), len(word2)); i++ {
		if word1[i] != word2[i] {
			distance++
		}
	}

	return distance + int(math.Abs(float64(len(word1)-len(word2))))
}

func EditDistance(word1 string, word2 string) int {
	x := len(word1)
	y := len(word2)

	if x == 0 && y == 0 {
		return 0
	} else if x == 0 || y == 0 {
		return x + y
	}

	if word1[x-1] == word2[y-1] {
		return EditDistance(word1[:x-1], word2[:y-1])
	}

	return min(EditDistance(word1[:x-1], word2[:y])+1, EditDistance(word1[:x-1], word2[:y-1])+1, EditDistance(word1[:x], word2[:y-1])+1)
}

func main() {
	// word1 := "test"
	// word2 := "trest"
	// fmt.Printf("Distance(%s, %s) = %d\n", word1, word2, Distance(word1, word2))
	// fmt.Printf("Distance(%s, %s) = %d\n", word1, word2, Distance(word1, word2))

	// fmt.Printf("EditDistance(%s, %s) = %d\n", word1, word2, EditDistance(word1, word2))

	// provided by https://www.bioinformatics.org/sms2/random_dna.html
	dna1000_1 := "ggtggttggtcagaaccgtcccgtatgttcataactaggcactagtaccggggccaggacgggagtgcaatagcaagcccttatcaaaaccgtcgcgctaaccacgcaaagatacggtatcacatatgccaagaattggggatgggtattagaatgacctaggtcaacactccttgttagagcgagtggcgtgtgacgtaccacgtcgtacttaactagatcgcttaaagccccgatgtggccacttggaggattcaaaggccctaatgatcctcacacgctaccgaggttgacggcgcttcttgaaaacacaaatttcttggtgacatacgcctacgactcattgtcgtacttttcgtctatcaccaagcgaaacctcccccacttaaccatctatgcgaattgttattcggcaccgccaccgtggaaacccgtcataaaaggaccatgccaaattggtttcatcgacaaagtccattaagttcgatataaacttatttgcagctcgcaagataaaaggctatgtccatgccatgttcggcgcacctctcctcgcgctgtaggacgcaacgttcgttcataatcgagtagtcctgctgcactgatggagccatccattgcagcgtcagcgcttcgactccggcccgctcatcgctagttagctatccgtacagtatcagaacatcttggggcttagtaaagtggtcggatccggtgttttttgcagtagcaaatggtttctaaaaacctgtcggcttttagattttacgatccctcgagtcttcgacttcttcgatcgtcacggtcctaagtgtcttgcgaccaggtatcagtgggcgcgtgcactttttgagttcgaagttagcgagcgtccctagaagtatccaattgcacctgttgaaaggaggaatatcctcaaattttaggaccttttagccttacccatactcgtggtagaagcattcggtcgtcggttagagttccattagtaataaatcgc"
	dna1000_2 := "aagtggggcagtggctcacacccatcatttggtgctaggcaatatatggtgaaaattcggtgcgggaagccaatcttgatgcagtcaactaaggtaaggctggcatgactagaaagcgttgacggcactacgtccatacatgcagccagtcgagataagtacttatacggttaccatctatgaaccagaccggatgtaatccagattaaacgggattgggtctttgctttcacccgggcttggttagagacagcaccctttcctgattacacctcgcataaaaccctagattttaggacactggacggtcttttcgcgatgcttttggtgtgcgccggacaaaggttataaatggtgtctctagtgaaggacggtttagtcgatgccaacgtgtatcaatgtagggcacggccggaggtctcgctggtattgcatttcgggatccgatgaatatcgtacgatagtagtgtccacagaacctttgtgtagttatacgcgctgtggtaccgatggccatagccgtagtggtccgctttgtgtgctgcgctacctgccggccctttaagggaacacgtgtaagccagttaactgagttcctaacccccaagagcatcgctccgatgtgttacgtactctcgtcactccagagatgcacgctcgactagtggtctggcagttatcggcttcgtgaagtatcgcaagtcttgacgttggactttgggtattataaccaatgtcgtgacgatatcgtgtcctagcgggctacctacatgcgggcggtaatatcgcgaatggccgcccacaagagtagaatcagttttcgtgtcctccttggttttcctgcatcgaatgttagctaggctgggacatcaatatatgtttcgcgcgtctttggtagcttccactcatctaaacattatcctggcctactgaaagtaatttccagggaccaccaacgggtccctggccgtattacccagcatcgtttctcccaggtcaa"
	dna100_1 := dna1000_1[:100]
	dna100_2 := dna1000_2[:100]
	fmt.Printf("EditDistance(%s, %s) = %d\n", dna100_1, dna100_2, EditDistance(dna100_1, dna100_2))
}
