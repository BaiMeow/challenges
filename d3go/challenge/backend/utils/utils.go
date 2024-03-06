package utils

import (
	"strings"
)

type Tree struct {
	Children []*Tree `json:"children,omitempty"`
	URL      string  `json:"url,omitempty"`
	Label    string  `json:"label,omitempty"`
}

func NewTree(label string) *Tree {
	return &Tree{
		Label: label,
	}
}

func FindSubTreeByLabel(tree *Tree, label string) *Tree {
	if tree == nil || tree.Children == nil {
		return nil
	}
	for i := range tree.Children {
		if tree.Children[i].Label == label {
			return tree.Children[i]
		}
	}
	return nil
}

func Paths2Tree(path string, link string, tree *Tree) {
	if path == "" {
		return
	}
	current, remain := pathSplit(path)
	if remain == "" {
		t := NewTree(current)
		t.URL = link
		tree.Children = append(tree.Children, t)
		return
	}
	var t *Tree
	if t = FindSubTreeByLabel(tree, current); t == nil || t.URL != "" {
		t = NewTree(current)
		tree.Children = append(tree.Children, t)
	}
	Paths2Tree(remain, link, t)
}

// recursive processing path
func pathSplit(path string) (string, string) {
	index := strings.Index(path, "/")
	if index == -1 {
		return path, ""
	} else if index == len(path)-1 {
		return path[:index], ""
	}
	return path[:index], path[index+1:]
}
