package util

// Pop and return last element of a list
func Pop[T any](alist *[]T) T {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return rv
}