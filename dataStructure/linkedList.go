package dataStructure

import "fmt"

type LlNode struct {
	Data interface{}
	Next *LlNode
}

type LinkedList struct {
	NodeCount int
	Head      *LlNode
	Tail      *LlNode
}

func GetAtFromLl(ll *LinkedList, pos int) *LlNode {
	if pos < 1 || pos > ll.NodeCount {
		return nil
	}
	curr := ll.Head
	for i := 1; i < pos; i++ {
		curr = curr.Next
	}
	return curr
}

func Traverse(ll *LinkedList) []*LlNode {
	var answer []*LlNode
	curr := ll.Head
	for curr != nil {
		answer = append(answer, curr)
		curr = curr.Next
	}
	return answer
}

func (ll *LinkedList) InsertAt(pos int, newNode *LlNode) error {
	if pos < 1 || pos > ll.NodeCount+1 {
		return fmt.Errorf("Check Index")
	}
	if pos == 1 {
		newNode.Next = ll.Head
		ll.Head = newNode
	} else {
		if pos == ll.NodeCount+1 {
			newNode.Next = ll.Tail.Next
			ll.Tail = newNode
		} else {
			newNode.Next = GetAtFromLl(ll, pos-1).Next
		}
		GetAtFromLl(ll, pos-1).Next = newNode
	}
	ll.NodeCount += 1
	return nil
}

func (ll *LinkedList) PopAt(pos int) *LlNode {
	if pos < 1 || pos > ll.NodeCount {
		return nil
	}
	n := GetAtFromLl(ll, pos)
	if pos == 1 {
		if ll.NodeCount == 1 {
			ll.Head = nil
			ll.Tail = nil
		} else {
			ll.Head = ll.Head.Next
		}
	} else {
		prev := GetAtFromLl(ll, pos-1)
		if pos == ll.NodeCount {
			prev.Next = nil
			ll.Tail = prev
		} else {
			prev.Next = n.Next
		}
	}
	ll.NodeCount -= 1
	return n
}

// 변수 선언시 *를 붙이면 pointer형 변수로 메모리 주소를 값으로 가진다.
// 이미 선언된 포인터 변수를 사용시 &를 붙이면 해당 변수의 메모리 주소.
// 이미 선언된 포인터 변수를 사용시 *를 붙이면 주소에 할당된 값.
func MakeNode(data interface{}) *LlNode {
	return &LlNode{data, nil}
}

func MakeLinkedList() *LinkedList {
	return &LinkedList{0, nil, nil}
}
