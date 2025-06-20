package main

import(
	"fmt"
	"bufio"	
	"os"
	"strings"
	"Sort"
)

type Pair struct{
	count int
	word string
}


func main() {
	file := os.Create("Scanned_in.txt")

	scanner := bufio.NewScanner(file)	

	scanner.Split(bufio.ScanWords)// split the phrase inputed into words

	var word string
	var countword map[string]int
	defer file.Close()	

	 for scanner.Scan() {
		word = scanner.Text()
		word = strings.ToLower(word)
		word = strings.Trim(word, ",!.?;^$\'""")
		countword[word]++
	 }
	 var pairs []Pair
	 for w, c := range countword{
		contain := Pair{w, c}
		pairs = append(pairs,contain)
	 }

	Sort.Merge_IT(pairs, 0, len(pairs)-1) // sort the pairs by count
	 fmt.Println("Word Frequency Count:")
	 for i:=0 ; i<5 && i<len(pairs); i++ {
		fmt.Printf("%s-> %d\n", pairs[i].word, pairs[i].count)
	 }
	 fmt.Println("voila the programme is done")
}