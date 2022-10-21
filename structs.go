package main

type List[T comparable] struct {
	head, tail *element[T]
}
type element[T comparable] struct {
	next *element[T]
	val  T
}
