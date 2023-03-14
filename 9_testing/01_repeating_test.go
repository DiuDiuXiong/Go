package main

import (
	"fmt"
	"testing"
)

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// edge cases
		{"", 0},
		{"b", 1},
		{"bbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"这里是丢丢", 4},
		{"一二三二一", 3},

		// Should get error
		//{"123ff", 1},
	}

	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubStr(tt.s); actual != tt.ans {
			t.Errorf("Got %d for input %s: "+"expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "化肥会挥发黑化肥发灰灰化肥发黑黑化肥发灰会挥发灰化肥挥发会发黑黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花黑灰化肥会挥发发灰黑讳为花飞灰黑化肥会挥发发黑灰为讳飞花黑灰化肥灰会挥发发灰黑讳为黑灰花会飞灰黑化肥会会挥发发黑灰为讳飞花化为灰黑化黑灰化肥灰会挥发发灰黑讳为黑灰花会回飞灰化灰黑化肥会会挥发发黑灰为讳飞花回化为灰"
	ans := 11
	for i := 0; i < b.N; i++ { // b.N have algorithm to determine how many terms to get benchmark
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("Got %d for input %s: "+"expected %d", actual, s, ans)
		}
	}

}

func BenchmarkSubStr2(b *testing.B) {
	s := "化肥会挥发黑化肥发灰灰化肥发黑黑化肥发灰会挥发灰化肥挥发会发黑黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花黑灰化肥会挥发发灰黑讳为花飞灰黑化肥会挥发发黑灰为讳飞花黑灰化肥灰会挥发发灰黑讳为黑灰花会飞灰黑化肥会会挥发发黑灰为讳飞花化为灰黑化黑灰化肥灰会挥发发灰黑讳为黑灰花会回飞灰化灰黑化肥会会挥发发黑灰为讳飞花回化为灰"
	ans := 11
	for i := 0; i < 13; i++ {
		s = s + s + s
	}
	fmt.Println(len(s))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("Got %d for input %s: "+"expected %d", actual, s, ans)
		}
	}
}
