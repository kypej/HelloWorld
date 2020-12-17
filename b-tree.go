package main

var sourceSlice = []int{15, 78, 57, 42, 27, 44, 14, 34, 69, 92, 64, 86, 26, 21, 7, 6, 62, 93, 81, 19, 56, 39, 71, 33, 29, 30, 91, 70, 77, 47, 12, 2, 83, 63, 5, 89, 50, 55, 59, 65, 53, 36, 84, 46, 67, 32, 11, 23, 88, 74}

type node struct {
	index int
	Left  *node
	Right *node
}

var rootNode node

func (n *node) doBTree(sourceSlice []int) {
	for i := range sourceSlice {
		if i == 0 {
			n.index = sourceSlice[i]
			continue
		}
		setNode(sourceSlice[i], n)
	}
}

func setNode(index int, nod *node) {
	if nod.index < index {
		if nod.Right == nil {
			nod.Right = &node{
				index: index,
			}
		} else if nod.Right.index == index {
			return
		} else {
			setNode(index, nod.Right)
		}
	} else {
		if nod.Left == nil {
			nod.Left = &node{
				index: index,
			}
		} else if nod.Left.index == index {
			return
		} else {
			setNode(index, nod.Left)
		}
	}
}

func (n *node) getNodeByIndex(index int) *node {
	switch {
	case n.index == index:
		return n
	case n.index < index:
		return getNode(index, *n.Right)
	case n.index > index:
		return getNode(index, *n.Left)
	}
	return nil
}

func getNode(index int, nod node) *node {
	switch {
	case nod.index == index:
		return &nod
	case nod.index < index:
		return getNode(index, *nod.Right)
	case nod.index > index:
		return getNode(index, *nod.Left)
	}
	return nil
}

func (n *node) setValue(index int) {
	switch {
	case n.index == index:
		return
	case n.index < index:
		setNode(index, n.Right)
	case n.index > index:
		setNode(index, n.Left)
	}
}

func (n *node) delValue(index int) *node {
	//n = n.getNodeByIndex(index)
	switch {
	case n == nil:
		return nil
	case n.index > index:
		n.Left = n.Left.delValue(index)
		return n
	case n.index < index:
		n.Right = n.Right.delValue(index)
		return n
	case n.Left == nil && n.Right == nil:
		n = nil
		//return nil
		return n
	case n.Left == nil && n.Right != nil:
		n = n.Right
		return n
	case n.Left != nil && n.Right == nil:
		n = n.Left
		return n
	case n.Left != nil && n.Right != nil:
		n = findMin(n.Right)
		return n
	}
	return n
}

func (n *node) del(index int) *node {
	nod := n.getNodeByIndex(index)
	switch {
	case nod.Left == nil && nod.Right == nil:
		//TODO удаление не робит
		nod = nil
	case nod.Left == nil && nod.Right != nil:
		*nod = *nod.Right
	case nod.Left != nil && nod.Right == nil:
		*nod = *nod.Left
	case nod.Left != nil && nod.Right != nil:
		*nod = *findMin(nod.Right)
	}
	return n
}

func findMin(nod *node) *node {
	if nod.Left == nil {
		return nod
	} else {
		findMin(nod.Left)
	}
	return nil
}

func findMax(nod *node) *node {
	if nod.Right == nil {
		return nod
	} else {
		findMax(nod.Right)
	}
	return nil
}
