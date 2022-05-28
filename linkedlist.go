type LLNode struct {
	val  int
	next *LLNode
}

type MyLinkedList struct {
	head   *LLNode
	tail   *LLNode
	length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (mll *MyLinkedList) Get(index int) int {
	fmt.Println("Get", index)
	fmt.Println("length: ", mll.length)
	if index == 4 {
		fmt.Println("Index 4", index)
		fmt.Println("length", mll.length)
	}
	if index >= mll.length || index < 0 {
		return -1
	}

	cur := mll.head
	i := 0
	for cur != nil {
		fmt.Println("i", i)
		fmt.Println(cur.val)
		fmt.Println(cur.next)
		if i == index {
			return cur.val
		}
		cur = cur.next
		i++
	}
	fmt.Println("makde it to -1")
	return -1
}

func (mll *MyLinkedList) AddAtHead(val int) {
	fmt.Println("Add head")
	fmt.Println("length: ", mll.length)
	cur := LLNode{
		val:  val,
		next: mll.head,
	}

	mll.head = &cur

	if mll.tail == nil {
		mll.tail = &cur
	}

	mll.length++
}

func (mll *MyLinkedList) AddAtTail(val int) {
	fmt.Println("Add Tail")
	fmt.Println("length: ", mll.length)
	cur := LLNode{
		val:  val,
		next: nil,
	}

	if mll.tail == nil {
		mll.tail = &cur
	} else {
		mll.tail.next = &cur
		mll.tail = &cur
	}

	mll.length++
}

func (mll *MyLinkedList) AddAtIndex(index int, val int) {
	fmt.Println("Add Index", index)
	fmt.Println("length: ", mll.length)
	if index > mll.length {
		return
	}
	if index == mll.length {
		mll.AddAtTail(val)
		return
	}
	if index == 0 {
		mll.AddAtHead(val)
		return
	}

	prev := mll.head
	cur := mll.head.next
	i := 1
	for cur != nil {
		if i == index {
			noo := LLNode{
				val:  val,
				next: cur,
			}
			prev.next = &noo

			mll.length++
			return
		}
		prev.next = cur
		cur = cur.next
		i++
	}

	mll.length++
	return

}

func (mll *MyLinkedList) DeleteAtIndex(index int) {
	fmt.Println("Delete")
	fmt.Println("length: ", mll.length)
	if index < 0 || index >= mll.length {
		return
	}

	if index == 0 {
		// Reset mll when deleting only node
		if mll.length == 1 {
			mll.head = nil
			mll.tail = nil
			mll.length = 0

			return
		}

		// Reset head when it is not the only node
		mll.head = mll.head.next
		mll.length--

		return
	}

	prev := mll.head
	cur := mll.head.next
	i := 1
	for cur != nil {
		if index == i {
			prev.next = cur.next
			mll.length--
			if cur.next == nil {
				mll.tail = cur
			}
			return
		}
		prev = cur
		cur = cur.next
		i++
	}

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */