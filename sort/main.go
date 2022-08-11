package main

import (
	"fmt"
	"sort"
)

type fruits []string

func (f fruits) Len() int {
	return len(f)
}

func (f fruits) Less(i, j int) bool {
	return len(f[i]) < len(f[j])
}

func (f fruits) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func main() {
	strs := []string{"a", "b", "z", "d", "e", "f", "g", "h", "i", "j", "k", "t", "m", "n", "o", "p", "q", "r", "s", "l", "u", "v", "w", "x", "y", "c"}

	fmt.Println("strings sorted: ", sort.StringsAreSorted(strs))

	sort.Strings(strs)
	fmt.Println(strs, "strings sorted: ", sort.StringsAreSorted(strs))

	ints := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println("ints sorted: ", sort.IntsAreSorted(ints))
	sort.Ints(ints)
	fmt.Println(ints, "ints sorted: ", sort.IntsAreSorted(ints))

	fruits := fruits{"orange", "apple", "pear", "grape", "banana", "kiwi", "mango", "lemon", "lime", "watermelon", "papaya", "cherry", "peach", "pineapple", "coconut", "pomegranate", "plum", "raspberry", "strawberry", "blueberry", "blackberry", "gooseberry", "cranberry", "clementine", "damson", "date", "elderberry", "fig", "ginkgo", "huckleberry", "jackfruit", "jambul", "japanese plum", "kiwano", "kumquat", "longan", "lychee", "mangosteen", "nectarine", "orange", "passionfruit", "peach", "pear", "persimmon", "pineapple", "plum", "pomegranate", "quince", "raspberry", "satsuma", "star fruit", "strawberry", "tangerine", "tomato", "ugli fruit", "watermelon", "xigua", "yuzu", "zuchinni"}

	fmt.Println("fruits sorted: ", sort.StringsAreSorted(fruits))
	sort.Sort(fruits)
	fmt.Println(fruits, "fruits sorted: ", sort.StringsAreSorted(fruits))
}
