package main

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	x, y, z := 1, 2, 3

	want := []int{1, 2, 3}

	lst := List[int]{}
	lst.Push(x)
	lst.Push(y)
	lst.Push(z)

	var got []int
	var got2 []int

	got2 = append(got2, lst.head.val)
	got2 = append(got2, lst.head.next.val)
	got2 = append(got2, lst.head.next.next.val)

	for e := lst.head; e != nil; e = e.next {
		got = append(got, e.val)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v got %v", want, got)
	}

	if !reflect.DeepEqual(got2, want) {
		t.Fatalf("want %v got %v", want, got2)
	}
}

func TestPushAndGetAll(t *testing.T) {
	x, y, z := 1, 2, 3

	want := []int{1, 2, 3}

	lst := List[int]{}
	lst.Push(x)
	lst.Push(y)
	lst.Push(z)

	got := lst.GetAll()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v got %v", want, got)
	}
}

func TestPull(t *testing.T) {
	x, y, z, w := 1, 2, 3, 4

	want := []int{2, 3, 4}

	lst := List[int]{}
	lst.Push(x)
	lst.Push(y)
	lst.Push(z)
	lst.Push(w)

	lst.Pull(1)
	got := lst.GetAll()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v got %v", want, got)
	}
}
