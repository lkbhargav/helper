package helper

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

const StartAscii = 97

type Trie[T constraints.Ordered] struct {
	Tree       *TrieNode[T]
	WordsCount int
}

type TrieNode[T constraints.Ordered] struct {
	Value      string
	Children   [26]*TrieNode[T]
	Occurences int
	End        bool
}

func NewTrie[T constraints.Ordered]() Trie[T] {
	rootNode := TrieNode[T]{
		Value:    "root",
		Children: [26]*TrieNode[T]{},
	}

	return Trie[T]{
		Tree:       &rootNode,
		WordsCount: 0,
	}
}

func (t *Trie[T]) AddAWord(word string) error {
	word = strings.ToLower(strings.TrimSpace(word))

	node := t.Tree

	for idx, c := range word {
		lastChar := false

		if idx == len(word)-1 {
			lastChar = true
		}

		if c-StartAscii < 0 {
			return fmt.Errorf("invalid character found: %c | word: %s", c, word)
		}

		nextNode := node.Children[c-StartAscii]

		if nextNode == nil {
			node.Children[c-StartAscii] = &TrieNode[T]{
				Value:      string(c),
				Children:   [26]*TrieNode[T]{},
				End:        lastChar,
				Occurences: 1,
			}
		} else {
			nextNode.Occurences += 1

			if lastChar && !nextNode.End {
				nextNode.End = true
			}
		}

		node = node.Children[c-StartAscii]
	}

	t.WordsCount += 1

	return nil
}

func (t *Trie[T]) AddWords(words ...string) error {
	for _, word := range words {
		err := t.AddAWord(word)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Trie[T]) Lookup(word string) bool {
	word = strings.ToLower(word)

	node := t.Tree

	for idx, c := range word {
		lastChar := false

		if idx == len(word)-1 {
			lastChar = true
		}

		nextNode := node.Children[c-StartAscii]

		if nextNode == nil {
			return false
		} else {
			if lastChar && nextNode.End {
				return true
			}
		}

		node = node.Children[c-StartAscii]
	}

	return false
}

type trieContainer[T constraints.Ordered] struct {
	node *TrieNode[T]
	str  string
}

func (t *Trie[T]) PossibilitiesForPrefix(prefix string) []string {
	words := []string{}

	node := t.Tree

	prefix = strings.ToLower(prefix)

	for _, c := range prefix {
		nextNode := node.Children[c-StartAscii]

		if nextNode == nil {
			return []string{}
		}

		node = node.Children[c-StartAscii]
	}

	if node.End {
		words = append(words, prefix)
	}

	stack := []trieContainer[T]{}

	for _, child := range node.Children {
		if child != nil {
			stack = append(stack, trieContainer[T]{
				node: child,
				str:  child.Value,
			})
		}
	}

outer:
	for len(stack) > 0 {
		stackLen := len(stack)

		for _, val := range stack {
			if val.node.End {
				words = append(words, prefix+val.str)
				if len(words) >= 10 {
					break outer
				}
			}

			for _, c := range val.node.Children {
				if c != nil {
					stack = append(stack, trieContainer[T]{
						node: c,
						str:  val.str + c.Value,
					})
				}
			}
		}

		stack = stack[stackLen:]
	}

	return words
}
