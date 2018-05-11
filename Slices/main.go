package main

import "fmt"

func main() {
	// create a new slice of size 5
	s := []int{1, 2, 3, 4, 5}

	// create a new variable
	// t now holds a reference to the entire underlying array of slice, s
	t := s[:]

	// create a new variable, r
	// r now holds a reference to the underlying array of slice t (and s)
	// r will have a length of 2, as it can only see the first 2 elements in the underlying array
	r := t[:2]

	// this increases how many elements that r can see of the underlying array
	// increasing its length
	r = r[:3]

	// since r references the same underlying array as s and t, this append impacts both
	// since r only has access to the first 3 elements, this append adds 99 to the 4th slot in ALL arrays
	// this means after this append, 99 can/will be in the middle of the other arrays
	r = append(r, 99)

	// this appends 6 to the end of the initial s slice, and the underlying array that r and t sit on top of
	// this call doubles the capacity (as typical with arrays) in order to fit this value
	// increasing capacity from 5 (initial) to 10
	// NOTE: We will not visually see '6' in the other slices as the elements they can see do not go up to
	// the max length of the array. They would need to be resized with s[:len(s)]
	s = append(s, 6)

	printSlice(t)
	printSlice(s)
	printSlice(r)
}

func printSlice(slice []int) {
	fmt.Print("len:", len(slice), " cap:", cap(slice), " data:", slice, "\n")
}
