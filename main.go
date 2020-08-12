package main

import (
	linkedList "DoubleLinkedListInGo/src"
	"fmt"
)

func main() {

	l := linkedList.LinkedList{}

	fmt.Println("\n************* Insert *************")
	l.Insert(12)
	l.Insert(13)
	l.Insert(14)
	l.Insert(15)
	l.Insert(-12)
	fmt.Println("************* Print *************")
	l.Print()

	// fmt.Println("\n************* InsertAtStart *************")
	// l.InsertAtStart(12)
	// l.InsertAtStart(13)
	// l.InsertAtStart(14)
	// l.InsertAtStart(15)
	// l.InsertAtStart(-12)
	// fmt.Println("************* Print *************")
	// l.Print()

	// fmt.Println("\n************* InsertAt *************")
	// l.InsertAt(0, 12)
	// l.InsertAt(-1, 13)
	// l.InsertAt(1, 14)
	// l.InsertAt(2, 15)
	// l.InsertAt(3, 4444)
	// l.InsertAt(4, 153)
	// l.InsertAt(1, 150)
	// fmt.Println("************* Print *************")
	// l.Print()

	// fmt.Println("Head: ", l.GetHead())
	// fmt.Println("Tail: ", l.GetTail())

	// fmt.Println("\n************* Search *************")
	// fmt.Println("Position of 12 value: ", l.Search(12))
	// fmt.Println("Position of 14 value: ", l.Search(14))
	// fmt.Println("Position of 15 value: ", l.Search(15))
	// fmt.Println("Position of 100 value: ", l.Search(100))

	// fmt.Println("\n************* GetAt *************")
	// fmt.Println("Get At 1st position: ", l.GetAt(0))
	// fmt.Println("Get At 3rd position: ", l.GetAt(2))
	// fmt.Println("Get At 4th position: ", l.GetAt(3))
	// fmt.Println("Get At -5 position ", l.GetAt(-5))
	// fmt.Println("Get At 100 position ", l.GetAt(100))

	// fmt.Println("\n************* DeleteAt *************")
	// l.DeleteAt(1)
	// fmt.Println("************* Print *************")
	// l.Print()

	// fmt.Println("\n************* DeleteAtEnd *************")
	// l.DeleteAtEnd()
	// l.DeleteAtEnd()
	// l.DeleteAtEnd()
	// l.DeleteAtEnd()
	// l.DeleteAtEnd()
	// fmt.Println("************* Print *************")
	// l.Print()

	// fmt.Println("\n************* DeleteAtFirst *************")
	// l.DeleteAtFirst()
	// l.DeleteAtFirst()
	// l.DeleteAtFirst()
	// l.DeleteAtFirst()
	// l.DeleteAtFirst()
	// fmt.Println("************* Print *************")
	// l.Print()

	// fmt.Println("************* Reverse *************")
	// l.Reverse()
}
