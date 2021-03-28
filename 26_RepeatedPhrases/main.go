package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Given a string, split on whitespace and return a slice of words, retaining sentence order
func splitSentence(sentence string) []string {
	return regexp.MustCompile(`\s+`).Split(sentence, -1)
}

// Remove any characters, like commas, that may be in the word
// Will retain hyphens
func removeNonAlphanumeric(s string) string {
	re := regexp.MustCompile(`[^a-zA-z-]`)
	return strings.ToLower(string(re.ReplaceAll([]byte(s), []byte(""))))
}

type phraseCount struct {
	phrase string
	count  int
}

func main() {
	// From the "sorting algorithms" wikipedia article: https://en.wikipedia.org/wiki/Sorting_algorithm
	var corpus = `Stable sort algorithms sort repeated elements in the same order that they appear in the input. When sorting some kinds of data, only part of the data is examined when determining the sort order. For example, in the card sorting example to the right, the cards are being sorted by their rank, and their suit is being ignored. This allows the possibility of multiple different correctly sorted versions of the original list. Stable sorting algorithms choose one of these, according to the following rule: if two items compare as equal, like the two 5 cards, then their relative order will be preserved, so that if one came before the other in the input, it will also come before the other in the output.

    Stability is important for the following reason: say that student records consisting of name and class section are sorted dynamically on a web page, first by name, then by class section in a second operation. If a stable sorting algorithm is used in both cases, the sort-by-class-section operation will not change the name order; with an unstable sort, it could be that sorting by section shuffles the name order. Using a stable sort, users can choose to sort by section and then by name, by first sorting using name and then sort again using section, resulting in the name order being preserved. (Some spreadsheet programs obey this behavior: sorting by name, then by section yields an alphabetical list of students by section.)
    Another reason: unstable sort may yield different output for the same input from run to run. Such behavior is unsuitable for some applications, for example for client-server applications where the server uses pagination for output and performs a new search-and-sort for every new page requested by the client.
    
    More formally, the data being sorted can be represented as a record or tuple of values, and the part of the data that is used for sorting is called the key. In the card example, cards are represented as a record (rank, suit), and the key is the rank. A sorting algorithm is stable if whenever there are two records R and S with the same key, and R appears before S in the original list, then R will always appear before S in the sorted list.
    
    When equal elements are indistinguishable, such as with integers, or more generally, any data where the entire element is the key, stability is not an issue. Stability is also not an issue if all keys are different.
    
    Unstable sorting algorithms can be specially implemented to be stable. One way of doing this is to artificially extend the key comparison, so that comparisons between two objects with otherwise equal keys are decided using the order of the entries in the original input list as a tie-breaker. Remembering this order, however, may require additional time and space.
    
    While there are a large number of sorting algorithms, in practical implementations a few algorithms predominate. Insertion sort is widely used for small data sets, while for large data sets an asymptotically efficient sort is used, primarily heap sort, merge sort, or quicksort. Efficient implementations generally use a hybrid algorithm, combining an asymptotically efficient algorithm for the overall sort with insertion sort for small lists at the bottom of a recursion. Highly tuned implementations use more sophisticated variants, such as Timsort (merge sort, insertion sort, and additional logic), used in Android, Java, and Python, and introsort (quicksort and heap sort), used (in variant forms) in some C++ sort implementations and in .NET.

    For more restricted data, such as numbers in a fixed interval, distribution sorts such as counting sort or radix sort are widely used. Bubble sort and variants are rarely used in practice, but are commonly found in teaching and theoretical discussions.

    When physically sorting objects (such as alphabetizing papers, tests or books) people intuitively generally use insertion sorts for small sets. For larger sets, people often first bucket, such as by initial letter, and multiple bucketing allows practical sorting of very large sets. Often space is relatively cheap, such as by spreading objects out on the floor or over a large area, but operations are expensive, particularly moving an object a large distance – locality of reference is important. Merge sorts are also practical for physical objects, particularly as two hands can be used, one for each list to merge, while other algorithms, such as heap sort or quick sort, are poorly suited for human use. Other algorithms, such as library sort, a variant of insertion sort that leaves spaces, are also practical for physical use.`

	// Regex to find all matches separated by "." or "!"
	reFindAllSentences := regexp.MustCompile(`([^.!]*?)[.!]\s*`)
	matches := reFindAllSentences.FindAllSubmatch([]byte(corpus), -1)

	// Map of all phrases with a count of how many times that phrase has been observed
	phraseMap := make(map[string]int)

	for _, match := range matches {
		sentence := string(match[1])
		words := splitSentence(sentence)

		for j, word := range words {
			// Clean up the words to remove any commas, colons, or other characters that we want to ignore
			words[j] = removeNonAlphanumeric(word)
		}

		// Find all word cominations of at least 2 words, not more than 10 words
		const minPhraseSize = 2
		const maxPhraseSize = 10
		for front := 0; front < len(words)-minPhraseSize; front++ {
			for back := front + (minPhraseSize - 1); back < len(words) && back-front < maxPhraseSize; back++ {
				phrase := strings.Join(words[front:back+1], " ")

				if count, ok := phraseMap[phrase]; ok {
					phraseMap[phrase] = count + 1
				} else {
					phraseMap[phrase] = 1
				}
			}
		}
	}

	// Convert the map to a list for sorting
	var list []phraseCount

	for k, v := range phraseMap {
		list = append(list, phraseCount{k, v})
	}

	sort.SliceStable(list, func(i, j int) bool {
		// Sort by highest count
		if list[i].count > list[j].count {
			return true
		}
		if list[i].count < list[j].count {
			return false
		}
		// Subsort by longest phrase.
		// This implies that "additionally observed" is more important than "in the"
		return len(list[i].phrase) >= len(list[j].phrase)
	})

	fmt.Println("Top 5 phrases:")
	for _, phrase := range list[:5] {
		fmt.Printf(" - %s (%d)\n", phrase.phrase, phrase.count)
	}
	fmt.Println("Number of phrases: ", len(list))
}
