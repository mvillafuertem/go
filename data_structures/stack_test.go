package data_structures

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func createStack() Stack {
	stack := Stack{}
	stack.New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	return stack
}

func TestPush(t *testing.T) {

	stack := createStack()

	assert.Equal(t, 6, *stack.Length())
}

func TestPop(t *testing.T) {

	stack := createStack()

	stack.Pop()
	stack.Pop()

	assert.Equal(t, 4, *stack.Length())
}

func TestShouldPrintAllElement(t *testing.T) {

	stack := createStack()

	stack.Print()

	assert.Equal(t, 6, *stack.Length())
}

func TestShouldReturnLengthOfStack(t *testing.T) {

	stack := createStack()

	assert.Equal(t, 6, *stack.Length())
}

func TestShouldCopyStack(t *testing.T) {

	stack := createStack()

	result := Stack{}
	result.New()

	stack.Copy(&result)

	stack.Print()
	fmt.Println()
	result.Print()

	assert.Equal(t, 6, *result.Length())
	assert.Equal(t, result, stack)
	assert.Equal(t, &result, &stack)
}

func TestShouldReverseStack(t *testing.T) {

	stack := createStack()

	result := Stack{}
	result.New()

	stack.Reverse(&result)

	stack.Print()
	fmt.Println()
	result.Print()

	assert.Equal(t, 6, *result.Length())
	assert.NotEqual(t, result, stack)
}


func TestShouldSetLastValueStack(t *testing.T) {

	stack := createStack()

	stack.SetLastElement(0)

	stack.Print()
	fmt.Println()

	assert.Equal(t, 7, *stack.Length())
}

func TestShouldGetLastValueStack(t *testing.T) {

	stack := createStack()


	stack.Print()
	fmt.Println()

	assert.Equal(t, 1, 	*stack.GetLastElement())
}

func TestShouldRemoveElementStack(t *testing.T) {

	stack := createStack()
	stack.Print()
	fmt.Println()
	element := stack.RemoveElement(2)

	assert.Equal(t, true, 	element)


	stack.Print()
	fmt.Println()
}

func TestShouldFindElementStack(t *testing.T) {

	stack := createStack()
	stack.Print()
	fmt.Println()
	element := stack.FindElement(2)

	assert.Equal(t, 2, 	element)


	stack.Print()
	fmt.Println()
}

func TestShouldFindElementAndCountStack(t *testing.T) {

	stack := createStack()
	stack.Print()
	fmt.Println()
	element, count := stack.FindElementAndCount(2, 0, 0)


	assert.Equal(t, 2, 	element)
	assert.Equal(t, 1, 	count)


	stack.Print()
	fmt.Println()
}


func TestShouldRemoveElementResultAfterSubtractStack(t *testing.T) {

	stack := Stack{}
	stack.New()
	stack.Push(10)
	stack.Push(4)
	stack.Push(1)
	stack.Push(2)
	stack.Push(5)
	stack.Print()
	fmt.Println()
	n:= 9
	r := stack.RemoveElementResultAfterSubtract2(&n)

	assert.Equal(t, true, 	r)

	stack.Print()
	fmt.Println()

	p := unsafe.Pointer(&n)
	fmt.Println(*(*int)(p))

}
