package goqs

import "testing"

func TestStackSize(t *testing.T) {
	stack := MakeStack()
	for i := 0; i < 10; i++ {
		stack.Push(HTMLElement{})
	}
	if stack.Size() != 10 {
		t.Error("Size calculations are wrong")
	}
}

func TestStackPop(t *testing.T) {
	stack := MakeStack()
	arr := []string{"6", "5", "4", "3", "2", "1"}
	for _, v := range arr {
		stack.Push(HTMLElement{Tag: v})
	}

	for i := len(arr) - 1; i >= 0; i-- {
		elem, err := stack.Top()
		if err != nil || elem.Tag != arr[i] {
			t.Error("Top failed.")
		}

		elem, err = stack.Pop()
		if err != nil || elem.Tag != arr[i] {
			t.Error("Pop failed.")
		}
	}
}

func TestStackErr(t *testing.T) {
	stack := MakeStack()
	dummyError := EmptyStackError{}

	_, err := stack.Top()
	if err == nil || err.Error() != dummyError.Error() {
		t.Error("Top Error catching fails")
	}

	_, err = stack.Pop()
	if err == nil || err.Error() != dummyError.Error() {
		t.Error("Pop Error catching fails")
	}
}
