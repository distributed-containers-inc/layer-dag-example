package main

import (
	"testing"
	"time"
)

func TestAdd1(t *testing.T) {
	time.Sleep(time.Second * 5)
	if Add(1, 1) != 2 {
		t.Error("1 + 1 != 2")
	}
}

func TestAdd2(t *testing.T) {
	time.Sleep(time.Second * 5)
	if Add(2, 2) != 4 {
		t.Error("2 + 2 != 4")
	}
}

func TestAdd3(t *testing.T) {
	time.Sleep(time.Second * 5)
	if Add(3, 3) != 6 {
		t.Error("3 + 3 != 6")
	}
}

func TestAdd4(t *testing.T) {
	time.Sleep(time.Second * 5)
	if Add(4, 4) != 8 {
		t.Error("4 + 4 != 8")
	}
}

func TestAdd5(t *testing.T) {
	time.Sleep(time.Second * 5)
	if Add(5, 5) != 10 {
		t.Error("5 + 5 != 10")
	}
}

func TestAdd6(t *testing.T) {
	time.Sleep(time.Second * 5)
	if Add(6, 6) != 12 {
		t.Error("6 + 6 != 12")
	}
}

func TestAdd7(t *testing.T) {
	time.Sleep(time.Second * 5)
	if Add(7, 7) != 14 {
		t.Error("7 + 7 != 14")
	}
}

func TestAdd8(t *testing.T) {
	time.Sleep(time.Second * 5)
	if Add(8, 8) != 16 {
		t.Error("8 + 8 != 16")
	}
}

func TestAdd9(t *testing.T) {
	time.Sleep(time.Second * 5)
	if Add(9, 9) != 18 {
		t.Error("9 + 9 != 18")
	}
}