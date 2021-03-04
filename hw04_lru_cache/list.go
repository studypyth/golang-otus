package hw04

type List interface {
	Len() int                          // длина списка
	Front() *ListItem                  // первый элемент списка
	Back() *ListItem                   // последний элемент списка
	PushFront(v interface{}) *ListItem // добавить значение в начало
	PushBack(v interface{}) *ListItem  // добавить значение в конец
	Remove(i *ListItem)                // удалить элемент
	MoveToFront(i *ListItem)           // переместить элемент в начало
}

type list struct {
	len   int
	first *ListItem
	last  *ListItem
}

func (list *list) Len() int {
	return list.len
}

func (list *list) Front() *ListItem {
	return list.first
}

func (list *list) Back() *ListItem {
	return list.last
}

func (list *list) PushFront(v interface{}) *ListItem {
	var el ListItem
	el.Value = v
	if list.len == 0 {
		list.first = &el
		list.last = &el
	} else {
		el.Next = list.first
		list.first.Prev = &el
		list.first = &el
	}
	list.len++
	return &el
}

func (list *list) PushBack(v interface{}) *ListItem {
	var el ListItem
	el.Value = v
	if list.len == 0 {
		list.first = &el
		list.last = &el
	} else {
		el.Prev = list.last
		list.last.Next = &el
		list.last = &el
	}
	list.len++
	return &el
}

func (list *list) Remove(i *ListItem) {
	// 70 9 3
	prev := i.Prev
	next := i.Next
	switch {
	case prev == nil:
		list.first = next
		list.first.Prev = nil
	case next == nil:
		list.last = prev
		list.last.Next = nil
	default:
		prev.Next = next
		next.Prev = prev
	}
	list.len--
}

func (list *list) MoveToFront(i *ListItem) {
	if prev := i.Prev; prev != nil {
		tmpfirst := list.first
		list.first = i
		list.first.Prev = nil
		list.first.Next = tmpfirst
		list.last = prev
		list.last.Prev = prev.Prev
		list.last.Next = nil
	}
}

func NewList() List {
	return new(list)
}
