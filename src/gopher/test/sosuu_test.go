package main

import (
	"reflect"
	"testing"
)

func TestSosuchk_minus(t *testing.T) {
	e := len([]int{})
	a := len(sosuchk(-1))
	if e != a {
		t.Errorf(`expect="%v" actual="%v"`, e, a)
	}
}

func TestSosuchk0(t *testing.T) {
	e := len([]int{})
	a := len(sosuchk(0))
	if e != a {
		t.Errorf(`expect="%v" actual="%v"`, e, a)
	}
}
func TestSosuchk1(t *testing.T) {
	e := len([]int{})
	a := len(sosuchk(1))
	if e != a {
		t.Errorf(`expect="%v" actual="%v"`, e, a)
	}
}

func TestSosuchk2(t *testing.T) {
	e := []int{2}
	a := sosuchk(2)
	if !reflect.DeepEqual(e, a) {
		t.Errorf(`expect="%v" actual="%v"`, e, a)
	}
}

func TestSosuchk3(t *testing.T) {
	e := []int{2, 3}
	a := sosuchk(3)
	if !reflect.DeepEqual(e, a) {
		t.Errorf(`expect="%v" actual="%v"`, e, a)
	}
}

func TestSosuchk4(t *testing.T) {
	e := []int{2, 3}
	a := sosuchk(4)
	if !reflect.DeepEqual(e, a) {
		t.Errorf(`expect="%v" actual="%v"`, e, a)
	}
}

func TestSosuchk10(t *testing.T) {
	e := []int{2, 3, 5, 7}
	a := sosuchk(10)
	if !reflect.DeepEqual(e, a) {
		t.Errorf(`expect="%v" actual="%v"`, e, a)
	}
}
