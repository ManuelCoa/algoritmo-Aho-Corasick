package algoritmo

type Node struct {
	Children map[rune]*Node
	Fail     *Node
	Output   []string
}

type AhoCorasick struct {
	Root *Node
}

func NewAhoCorasick(patterns []string) *AhoCorasick {
	ac := &AhoCorasick{Root: &Node{Children: make(map[rune]*Node)}}
	ac.buildTrie(patterns)
	ac.buildFailLinks()
	return ac
}

func (ac *AhoCorasick) buildTrie(patterns []string) {
	for _, pattern := range patterns {
		current := ac.Root
		for _, char := range pattern {
			if _, exists := current.Children[char]; !exists {
				current.Children[char] = &Node{Children: make(map[rune]*Node)}
			}
			current = current.Children[char]
		}
		current.Output = append(current.Output, pattern)
	}
}

func (ac *AhoCorasick) buildFailLinks() {
	queue := []*Node{}
	for _, child := range ac.Root.Children {
		child.Fail = ac.Root
		queue = append(queue, child)
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for char, child := range current.Children {
			fail := current.Fail
			for fail != nil && fail.Children[char] == nil {
				fail = fail.Fail
			}

			if fail == nil {
				child.Fail = ac.Root
			} else {
				child.Fail = fail.Children[char]
				child.Output = append(child.Output, child.Fail.Output...)
			}

			queue = append(queue, child)
		}
	}
}

func (ac *AhoCorasick) Search(text string) map[string][]int {
	results := make(map[string][]int)
	current := ac.Root

	for i, char := range text {
		for current != nil && current.Children[char] == nil {
			current = current.Fail
		}

		if current == nil {
			current = ac.Root
			continue
		}

		current = current.Children[char]
		for _, match := range current.Output {
			results[match] = append(results[match], i-len(match)+1)
		}
	}

	return results
}

func AhoCorasickSearch(text string, patterns []string) map[string][]int {
	ac := NewAhoCorasick(patterns)
	return ac.Search(text)
}
