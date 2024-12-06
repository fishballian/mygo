package mlib

type Trie struct {
	child [26]*Trie
	end   bool
}

func (t *Trie) Insert(word string) {
	cur := t
	for i := range word {
		b := word[i] - 'a'
		if cur.child[b] == nil {
			cur.child[b] = &Trie{}
		}
		cur = cur.child[b]
	}
	cur.end = true
}
