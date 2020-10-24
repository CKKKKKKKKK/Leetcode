package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	//Split the list
	if head == nil {
		return
	}
	pilot, split := head, head
	for pilot.Next != nil && pilot.Next.Next != nil {
		split = split.Next
		pilot = pilot.Next.Next
	}
	//Reverse the second one
	ptr := split.Next
	split.Next = nil
	var head2 *ListNode = nil
	for ptr != nil {
		tmp := ptr.Next
		ptr.Next = head2
		head2 = ptr
		ptr = tmp
	}
	//Merge two lists
	ptr1 := head
	ptr2 := head2
	for ptr2 != nil {
		tmp := ptr2.Next
		ptr2.Next = ptr1.Next
		ptr1.Next = ptr2
		ptr1 = ptr2.Next
		ptr2 = tmp
	}
}

func main() {

}
