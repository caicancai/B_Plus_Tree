package B_Plus_Tree

import "strings"

type key string

func (k key) compare(target key) int {
	return strings.Compare(k.toString(), target.toString())
}

func (k key) toString() string {
	return string(k)
}
