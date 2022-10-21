package main

func Map[T any](input []T, f func(T) T) []T {
	var result []T
	for _, v := range input {
		result = append(result, f(v))
	}
	return result
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}
func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func (lst *List[T]) Pull(v T) {
	var elems List[T]
	for e := lst.head; e != nil; e = e.next {

		if v != e.val {
			elems.Push(e.val)
		}
	}

	*lst = elems
}
