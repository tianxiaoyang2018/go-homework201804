package util

import "testing"

func TestConvertSliceTo2DSlice(t *testing.T) {
	// test1
	list := [][]string{
		[]string{},
		[]string{"1", "0", "-0", "-2", "test1", "1", "0", "-0", "-2", "test2", "1", "0", "-0", "-2", "test3"},
		[]string{"a", "b", "c"},
		[]string{"a", "b", "c", "d", "e", "f"},
	}
	expectedList := [][][]string{
		// test0
		[][]string{},

		// test1
		[][]string{
			{"1", "0", "-0", "-2", "test1"},
			{"1", "0", "-0", "-2", "test2"},
			{"1", "0", "-0", "-2", "test3"},
		},

		// test2
		[][]string{
			{"a", "b", "c"},
		},

		// test3
		[][]string{
			{"a", "b", "c", "d", "e"},
			{"f"},
		},
	}
	for i := 0; i < len(list); i++ {
		if !isEqual2DSlice(expectedList[i], ConvertSliceToSlice2D(list[i], 5)) {
			t.Errorf("TestConvertSliceTo2DSlice test%v failed!", i)
		}
	}

}

func isEqual2DSlice(listA, listB [][]string) bool {
	if len(listA) != len(listB) {
		return false
	}

	for i := 0; i < len(listA); i++ {
		for k := 0; k < len(listA[i]); k++ {
			if listA[i][k] != listB[i][k] {
				return false
			}
		}
	}

	return true
}

func TestSliceEqualRegardlessOfOrder(t *testing.T) {
	inputs := []struct {
		a     []string
		b     []string
		equal bool
	}{
		{[]string{"a", "a", "b"}, []string{"a", "b", "b"}, false},
		{[]string{"a", "a", "b"}, []string{"a", "b"}, false},

		{[]string{"a", "a", "b"}, []string{"a", "a", "b"}, true},
		{[]string{}, []string{}, true},
	}
	for _, input := range inputs {
		if res := SliceEqualRegardlessOfOrder(input.a, input.b); res != input.equal {
			t.Errorf("compare %v and %v, wanna %v, got %v", input.a, input.b, input.equal, res)
		}
	}
}
