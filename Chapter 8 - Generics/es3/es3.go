package main

import (
	"errors"
	"fmt"
)

type Node[T comparable] struct {
	Element T
	Next    *Node[T]
}

type LinkedList[T comparable] struct {
	Head *Node[T]
}

func (l *LinkedList[T]) Add(newElement T) {
	if l.Head == nil {
		l.Head = &Node[T]{Element: newElement}
		return
	}

	// aggiungo in testa
	oldHead := l.Head
	l.Head = &Node[T]{Element: newElement}
	l.Head.Next = oldHead

	// aggiungo in coda
	// cursor := l.Head
	// for cursor.Next != nil {
	// 	cursor = cursor.Next
	// }
	// cursor.Next = &Node[T]{Element: newElement}
}

func (l *LinkedList[T]) Insert(newElement T, index int) error {
	if index < 0 {
		return errors.New("indice < 0")
	}

	contaPos := 0
	cursor := l.Head
	for contaPos < index-1 {
		if cursor != nil {
			cursor = cursor.Next
		} else {
			return errors.New(fmt.Sprint("posizione", index, "inesistente"))
		}

		contaPos++
	}

	old := cursor.Next
	cursor.Next = &Node[T]{
		Element: newElement,
		Next:    old,
	}

	return nil
}

func (l *LinkedList[T]) Index(element T) int {
	contaPos := 0
	cursor := l.Head
	if cursor == nil {
		return -1
	}

	for cursor.Next != nil {
		if cursor.Element == element {
			return contaPos
		}

		cursor = cursor.Next
		contaPos++
	}

	return -1
}

func (l LinkedList[T]) String() {
	cursor := l.Head
	for cursor != nil {
		fmt.Println(cursor.Element)
		cursor = cursor.Next
	}
}

func main() {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.String()

	fmt.Println()
	err := list.Insert(1, list.Index(3))
	if err != nil {
		fmt.Println(err)
	}
	list.String()

	err = list.Insert(999, 999)
	if err != nil {
		fmt.Println(err)
	}

	type kevin struct {
		cool  string
		lista LinkedList[int]
	}

	kevinList := LinkedList[kevin]{}
	kevinList.Add(kevin{
		cool:  "si",
		lista: list,
	})
	kevinList.Add(kevin{
		cool:  "no",
		lista: list,
	})
	kevinList.Add(kevin{
		cool:  "forse",
		lista: list,
	})
	kevinList.String()
}
