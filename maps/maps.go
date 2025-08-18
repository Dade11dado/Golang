package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	//First way to create map
	am := map[string]int{
		"Todd":    42,
		"Henry":   16,
		"Pageant": 14}

	fmt.Printf("the age of Henry was %v \n", am["Henry"])

	//Second way to create map
	an := make(map[string]int)
	an["Lucas"] = 28
	an["Stephany"] = 12
	an["Step"] = 37

	fmt.Println(an)

	//Range over a map
	for k, v := range an {
		fmt.Println(k, v)
	}

	//Delete an element
	delete(an, "Lucas")

	fmt.Println(an)

	//comma ok idiom
	if v, ok := an["Georgy"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key not found")
	}

	//Counting words in a book
	fmt.Println("-----------Count words in a book -----------------")
	f, err := os.Open("book.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	defer f.Close()

	words, err := freq(f)
	if err != nil {
		log.Fatalf("Error from freq in mai %s \n", err)
	}

	pairs := sortWordFreq(words)

	for _, pair := range pairs {
		fmt.Printf("%s \t %d \n", pair.key, pair.Value)
	}

	w, n, err := maxWord(words)

	if err != nil {
		log.Fatalf("Error with maxWords %d \n", err)
	}

	fmt.Printf("%#v has a frequency of %d \n", w, n)

}

func freq(r io.Reader) (map[string]int, error) {

	wordFreq := make(map[string]int)
	s := bufio.NewScanner(r)
	s.Split((bufio.ScanWords))

	for s.Scan() {
		word := strings.ToLower(s.Text())
		wordFreq[word]++
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return wordFreq, nil

}

//For sorting frequency of words

type Pair struct {
	key   string
	Value int
}

//implement len less and swap methods on PaairLIst
//to satisfy the sort interface

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value > p[j].Value //ascending order
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortWordFreq(m map[string]int) PairList {
	//convert map to pair list
	pairs := make(PairList, len(m))
	i := 0
	for key, value := range m {
		pairs[i] = Pair{key, value}
		i++
	}

	//sort the pair list
	sort.Sort(pairs)

	return pairs
}

func maxWord(m map[string]int) (string, int, error) {
	if len(m) == 0 {
		return "", 0, fmt.Errorf("Empty map")
	}

	maxW := ""
	maxF := 0

	for k, v := range m {
		if v > maxF {
			maxW = k
			maxF = v
		}
	}

	return maxW, maxF, nil
}
