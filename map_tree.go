/**
 * @author Co2
 * @file_name map_tree.go
 */

package main

type TraversalTree struct {
	NowId    string           `json:"NowId"`
	FatherId string           `json:"FatherId"`
	Rule     string           `json:"rule"`
	Child    []*TraversalTree `json:"childDict"`
}

func has(v1 TraversalTree, vs []*TraversalTree) bool {
	var has bool
	has = false
	for _, v2 := range vs {
		v3 := *v2
		if v1.NowId+v1.FatherId == v3.FatherId {
			has = true
			break
		}
	}
	return has

}

func findChild(v *TraversalTree, vs []*TraversalTree) (ret []*TraversalTree) {
	for _, v2 := range vs {
		if v.NowId+v.FatherId == v2.FatherId {
			ret = append(ret, v2)
			//logging.Println("发现子节点, 当前节点: ", v.FatherId, "父节点: ", v.NowId, "生成节点: ", v2.FatherId)
		} else {
			//logging.Printf("未发现子节点, 当前节点:%v 父节点:%v, 不能生成节点:%v\n", v.FatherId, v.NowId, v2.FatherId)
		}
	}
	return
}

func MakeTree(allNodes []*TraversalTree, rootNode *TraversalTree) {
	// allNodes 切片里面对象都是同一级的, rootNode是根
	childs := findChild(rootNode, allNodes)
	//logging.Println("获取个", len(childs), "子节点:", childs)
	for _, child := range childs {
		//logging.Println("同级节点加入, 节点指针", *child)
		rootNode.Child = append(rootNode.Child, child)
		if has(*child, allNodes) {
			//logging.Println("该节点指针", *child, "有子节点", child)
			MakeTree(allNodes, child)
		}
	}

}
