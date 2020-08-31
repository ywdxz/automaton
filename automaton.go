package automaton

import (
	"container/list"
	// "fmt"
)

type CheckResult struct {
	StartIndex int
	EndIndex   int
	TokenID    int
}

type Automaton interface {
	Check(src []byte) (result []CheckResult)
	// Print()
}

func NewAutomaton(worlds []string) Automaton {

	wordMap := make(map[int][]byte)
	for i, w := range worlds {
		wordMap[i] = []byte(w)
	}

	instance := &engine{
		rootNode: newNode(),
		wordMap:  wordMap,
	}

	instance.buildPrefixTree()
	instance.buildMismatchPointer()

	return instance
}

var (
	//default TokenID = -1
	defaultTokenID = -1
	nodeID         = 0
)

type node struct {
	id          int
	fail        *node
	nextNodeMap map[byte]*node
	tokenID     int
	worldLen    int
}

func newNode() (n *node) {
	n = &node{
		nextNodeMap: make(map[byte]*node),
		tokenID:     defaultTokenID,
		id:          nodeID + 1,
	}
	nodeID++

	return
}

type engine struct {
	rootNode *node
	wordMap  map[int][]byte //tokenID => world
}

func (e *engine) buildPrefixTree() {

	for tokenID, world := range e.wordMap {
		curNode := e.rootNode
		for _, c := range world {
			n, ok := curNode.nextNodeMap[c]
			if !ok {
				n = newNode()
				curNode.nextNodeMap[c] = n
			}
			curNode = n
		}
		curNode.tokenID = tokenID
		curNode.worldLen = len(world)
	}
}

func (e *engine) buildMismatchPointer() {

	ll := list.New()
	ll.PushFront(e.rootNode)

	for ll.Len() > 0 {
		curNode := ll.Remove(ll.Back()).(*node)
		for c, n := range curNode.nextNodeMap {
			n.fail = e.rootNode
			if curNode != e.rootNode {
				for p := curNode.fail; p != nil; p = p.fail {
					if n1, ok := p.nextNodeMap[c]; ok {
						n.fail = n1
						break
					}
				}
			}
			ll.PushFront(n)
		}
	}
}

func (e *engine) Check(src []byte) (result []CheckResult) {
	result = e.check(src)
	return
}

func (e *engine) check(src []byte) (result []CheckResult) {

	curNode := e.rootNode

	for k, v := range src {

		for curNode.nextNodeMap[v] == nil && curNode != e.rootNode {
			curNode = curNode.fail
		}
		curNode = curNode.nextNodeMap[v]
		if curNode == nil {
			curNode = e.rootNode
		}

		if curNode != e.rootNode {
			if curNode.tokenID != defaultTokenID {
				//hit world
				result = append(result, CheckResult{
					StartIndex: k - curNode.worldLen + 1,
					EndIndex:   k + 1,
					TokenID:    curNode.tokenID,
				})
			}
		}
	}
	return
}

// func (e *engine) Print() {

// 	ll := list.New()
// 	ll.PushFront(e.rootNode)

// 	for ll.Len() > 0 {
// 		curNode := ll.Remove(ll.Back()).(*node)
// 		fmt.Printf("[%d],%d,%+v,%d\n", curNode.id, curNode.tokenID, curNode.nextNodeMap, len(curNode.nextNodeMap))
// 		for c, v := range curNode.nextNodeMap {
// 			ll.PushFront(v)
// 			fmt.Printf("{%c}\n", c)
// 		}
// 	}

// }
