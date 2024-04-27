package binhexa

import (
	"strings"
	"testing"
)

func TestBin(t *testing.T){
	input := []string{"10","(bin)"}
	want := []string{"2"}
	got := Bin(input)

	toJoinWant := strings.Join(want, " ")
	toJoinGot := strings.Join(got, " ")

	if toJoinGot != toJoinWant{
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}

func TestHex(t *testing.T){
	input := []string{"42","(hex)"}
	want := []string{"66"}
	got := Hex(input)

	toJoinWant := strings.Join(want, " ")
	toJoinGot := strings.Join(got, " ")

	if toJoinGot != toJoinWant{
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}