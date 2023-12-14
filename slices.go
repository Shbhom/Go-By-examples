package main

import (
	"fmt"
	"slices"
)

func Main7() {
	var s []string // slices are intialized with len 0 and contains nothing => nil
	fmt.Println(s, len(s) == 0, s == nil)

	s = make([]string, 10) //we can pass a value to which a slice cannot be extended
	fmt.Println("\n", s, s == nil, "\nlen of slice: ", len(s), "\ncapacity of slice: ", cap(s))
	s[9] = "alkdfjlk"
	fmt.Println("\nlength and capacity of s respectively", len(s), cap(s))
	s = append(s, "asqweoprq")
	fmt.Println("\nnew length and capcity of s respectively", len(s), cap(s))
	/*
		difference between len and capacity is if we know that our slice is going to extend
		to a certain value which's greater than len so we set a value for the capacity of the
		the slice to which we can extend the length to
		10,12 -> our len initlly was 10 and cap was 12 and then when we appended another element
		the lenght got increased by one and it will keep on increasing till cap
		if no cap value is entered during creation of an slice then after we reached the last element
		i.e. in our case it was 10, after which if we appended we'll get 20 as our new capacity i.e. 10+10
	*/

	s = append(s, "asldfkja", "asldfkjaskj", "asdkfjal;sk")
	fmt.Println()
	for i := 0; i < len(s); i++ {
		if s[i] == "" {
			continue
		}
		fmt.Println(i, s[i])
	}
	c := make([]string, len(s))
	copy(c, s) //copy the whole slice

	fmt.Println(c, len(c), cap(c))
	//c's cap will go to 28 i.e. currently c contains 0-13 elements=> len(c)=14 => new capacity = 14+14=28
	c = append(c, "adlkfjas", "asdlfkkjas")
	fmt.Println(c, len(c), cap(c))

	sliceFrom9To13 := s[9:13] //[include:exclude] or [) in sense of math
	fmt.Println(sliceFrom9To13)

	sliceFrom9 := s[9:]
	fmt.Println(sliceFrom9)

	copiedSlice := s[:]
	fmt.Println(copiedSlice)
	//we can also check whether two slices are equal or not
	fmt.Println(slices.Equal(c, s))
}
