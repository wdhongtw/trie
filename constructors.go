package trie

// BuildFromMap may be useful for var declaration
func BuildFromMap[T any](inputs map[string]T) *TrieGen[T] {
	t := &TrieGen[T]{}
	for key, value := range inputs {
		t.Put([]byte(key), value)
	}
	return t
}

// BuildFromList can be used to create Trie with arbitrary bytes slice as key (not valid strings, etc)
func BuildFromList[T any](inputs []struct {
	Key   []byte
	Value T
}) *TrieGen[T] {
	t := &TrieGen[T]{}
	for i := range inputs {
		t.Put(inputs[i].Key, inputs[i].Value)
	}
	return t
}

// BuildPrefixesOnly used to create just searching prefixes without any data
func BuildPrefixesOnly[T any](strs ...string) *TrieGen[struct{}] {

	t := &TrieGen[struct{}]{}

	for i := range strs {
		t.Put([]byte(strs[i]), struct{}{})
	}

	return t
}
