package main

type Numbers []int

type byInc struct {
	Numbers
}

type byDec struct {
	Numbers
}

func (n Numbers) Len() int              { return len(n) }
func (n Numbers) Swap(i, j int)         { n[i], n[j] = n[j], n[i] }
func (n Numbers) Greater(i, j int) bool { return n[i] > n[j] }

func (n byInc) Len() int           { return len(n.Numbers) }
func (n byInc) Swap(i, j int)      { n.Numbers[i], n.Numbers[j] = n.Numbers[j], n.Numbers[i] }
func (n byInc) Less(i, j int) bool { return n.Numbers[i] < n.Numbers[j] }

func (n byDec) Len() int           { return len(n.Numbers) }
func (n byDec) Swap(i, j int)      { n.Numbers[i], n.Numbers[j] = n.Numbers[j], n.Numbers[i] }
func (n byDec) Less(i, j int) bool { return n.Numbers[i] > n.Numbers[j] }
