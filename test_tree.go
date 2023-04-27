/**
 * @author Co2
 * @file_name test_tree.go
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestMakeTree(t *testing.T) {
	TraversalTrees := make([]TraversalTree, 0)
	var a TraversalTree
	a.NowId = "1"
	a.FatherId = ""
	TraversalTrees = append(TraversalTrees, a)
	a.NowId = "3"
	a.FatherId = "1"
	TraversalTrees = append(TraversalTrees, a)
	a.NowId = "4"
	a.FatherId = "1"
	TraversalTrees = append(TraversalTrees, a)
	a.NowId = "5"
	a.FatherId = "31"
	TraversalTrees = append(TraversalTrees, a)
	a.NowId = "6"
	a.FatherId = "31"
	TraversalTrees = append(TraversalTrees, a)
	fmt.Println(TraversalTrees)

	pTraversalTrees := make([]*TraversalTree, 0)
	for i, _ := range TraversalTrees {
		var a *TraversalTree
		a = &TraversalTrees[i]
		pTraversalTrees = append(pTraversalTrees, a)
	}

	var node *TraversalTree
	node = &TraversalTrees[0]
	MakeTree(pTraversalTrees, node)
	fmt.Println("the result we got is", pTraversalTrees)

	data, _ := json.Marshal(node)
	var out bytes.Buffer
	json.Indent(&out, data, "", "\t")
	fmt.Printf("student=%v\n", out.String())
}
