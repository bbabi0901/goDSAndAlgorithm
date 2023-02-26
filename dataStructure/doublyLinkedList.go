package dataStructure

import "fmt"

type DllNode struct {
	Data interface{}
	Prev *DllNode
	Next *DllNode
}

type DoublyLinkedList struct {
	NodeCount int
	Head      *DllNode
	Tail      *DllNode
}

func (dll *DoublyLinkedList) InsertAfter(prevNode, newNode *DllNode) error {
	next := prevNode.Next
	newNode.Next = next
	newNode.Prev = prevNode
	prevNode.Next = newNode
	next.Prev = newNode
	dll.NodeCount += 1
	return nil
}

func (dll *DoublyLinkedList) InsertBefore(nextNode, newNode *DllNode) error {
	prev := nextNode.Prev
	newNode.Next = nextNode
	newNode.Prev = prev
	prev.Next = newNode
	nextNode.Prev = newNode
	dll.NodeCount += 1
	return nil
}

func (dll *DoublyLinkedList) InsertAt(index int, newNode *DllNode) error {
	if index < 1 || index > dll.NodeCount+1 {
		return fmt.Errorf("Check Index")
	}
	prevNode := GetAt(dll, index-1)
	return dll.InsertAfter(prevNode, newNode)
}

func (dll *DoublyLinkedList) PopAfter(prevNode *DllNode) (*DllNode, error) {
	if prevNode == dll.Tail {
		return nil, fmt.Errorf("Check Index")
	}
	target := prevNode.Next
	prevNode.Next = target.Next
	target.Next.Prev = prevNode
	dll.NodeCount -= 1
	return target, nil
}

func (dll *DoublyLinkedList) PopBefore(nextNode *DllNode) (*DllNode, error) {
	if nextNode == dll.Head {
		return nil, fmt.Errorf("Check Index")
	}
	target := nextNode.Prev
	nextNode.Prev = target.Prev
	target.Prev.Next = nextNode
	dll.NodeCount -= 1
	return target, nil
}

func (dll *DoublyLinkedList) PopAt(index int) (*DllNode, error) {
	if index < 1 || index > dll.NodeCount+1 {
		return nil, fmt.Errorf("Check Index")
	}
	prevNode := GetAt(dll, index-1)
	return dll.PopAfter(prevNode)
}

func (dll *DoublyLinkedList) Concat(followingDll *DoublyLinkedList) {
	dll.Tail.Prev.Next = followingDll.Head.Next
	followingDll.Head.Next.Prev = dll.Tail.Prev
	dll.Tail = followingDll.Tail
	dll.NodeCount += followingDll.NodeCount
}

func Traverse(dll *DoublyLinkedList) []*DllNode {
	var result []*DllNode
	curr := dll.Head.Next
	for curr.Next != nil {
		result = append(result, curr)
		curr = curr.Next
	}
	return result
}

func Reverse(dll *DoublyLinkedList) []*DllNode {
	var result []*DllNode
	curr := dll.Tail.Prev
	for curr.Prev != nil {
		result = append(result, curr)
		curr = curr.Prev
	}
	return result
}

func GetAt(dll *DoublyLinkedList, pos int) *DllNode {
	if pos < 0 || pos > dll.NodeCount {
		return nil
	}
	var target *DllNode
	i := 0
	if pos < dll.NodeCount {
		target = dll.Head
		for i < pos {
			target = target.Next
			i++
		}
	} else {
		target = dll.Tail
		for i < dll.NodeCount-pos+1 {
			target = target.Prev
			i++
		}
	}
	return target
}

func InitDoublyLinkedList() *DoublyLinkedList {
	var dll *DoublyLinkedList
	dll = &DoublyLinkedList{NodeCount: 0, Head: InitNode(nil), Tail: InitNode(nil)}
	dll.Head.Prev = nil
	dll.Tail.Next = nil
	dll.Head.Next = dll.Tail
	dll.Tail.Prev = dll.Head
	return dll
}

func InitNode(data interface{}) *DllNode {
	return &DllNode{Data: data}
}
