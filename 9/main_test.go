package fortest

import (
	"testing"
)

func TestFuncSum(t *testing.T) {

	type test struct {
		data   []int
		expect int
	}

	tests := []test{
		{[]int{5, 15, 25}, 45},
		{[]int{2, 25, 30}, 57},
		{[]int{8, 12, 46}, 66},
	}

	for i, v := range tests {
		got := FuncSum(v.data...)
		want := v.expect
		if got != want {
			t.Errorf("Data %d - Sum failed! Got: %d but want: %d\n", i, got, want)
		} else {
			t.Logf("Data %d - Passed!", i)
		}
	}

}

func BenchmarkFuncSum(b *testing.B) {
	for i:=0; i<b.N; i++ {
		FuncSum(1,5,15,50,100)
	}
}
