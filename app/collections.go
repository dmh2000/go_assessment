package goassessment

func MakeMap(keys []interface{}, values []interface{}) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})

	for i, key := range keys {
		m[key] = values[i]
	}

	return m
}

func PushFront(list *LinkedList, value interface{}) {
	node := &LinkedListElement{list.head, value}
	list.head = node
}

func PopFront(list *LinkedList) interface{} {
	if list.head == nil {
		return nil
	}

	node := list.head
	list.head = node.next
	return node.value
}

func PushBack(list *LinkedList, value interface{}) {
	node := &LinkedListElement{nil, value}

	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}

	list.tail.next = node
	list.tail = node
}

func PopBack(list *LinkedList) interface{} {
	if list.head == nil {
		return nil
	}

	node := list.head
	if node.next == nil {
		list.head = nil
		list.tail = nil
		return node.value
	}

	for node.next.next != nil {
		node = node.next
	}

	value := node.next.value
	node.next = nil
	list.tail = node
	return value
}
