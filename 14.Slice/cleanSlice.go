package main

type MySlice []int

func (s MySlice) Remove(index int) []int {
	return append(s[:index], s[index+1:]...)
}
