package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	Head *ListItem
	Tail *ListItem
	len  int
}

func (l list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.Head
}

func (l *list) Back() *ListItem {
	return l.Tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newListItem := new(ListItem)
	newListItem.Value = v

	if l.Len() == 0 {
		l.Head = newListItem
		l.Tail = newListItem
	} else {
		l.Head.Prev = newListItem
		newListItem.Next = l.Head
		l.Head = newListItem
	}

	l.len++
	return newListItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newListItem := new(ListItem)
	newListItem.Value = v

	if l.Len() == 0 {
		l.Head = newListItem
		l.Tail = newListItem
	} else {
		l.Tail.Next = newListItem
		newListItem.Prev = l.Tail
		l.Tail = newListItem
	}
	l.len++
	return newListItem
}

func (l *list) Remove(i *ListItem) {
	l.len--
	switch {
	case i.Prev == nil && i.Next == nil:
		l.Head = nil
		l.Tail = nil
	case i.Prev == nil:
		i.Next.Prev = nil
		l.Head = i.Next
	case i.Next == nil:
		i.Prev.Next = nil
		l.Tail = i.Prev
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if l.Head == i {
		return
	}
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	l := new(list)
	l.len = 0
	l.Head, l.Tail = nil, nil
	return l
}
