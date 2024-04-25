package vowel

import (
	"strings"
	"testing"
)
func TestVow(t *testing.T){

	input := []string{"an","dog","is","a","animal"}
	want := []string{"a","dog","is","an","animal"}
	got :=Vow(input)

	toJoinWant := strings.Join(want, " ")
	toJoinGot := strings.Join(got, " ")

	if toJoinGot != toJoinWant{
		t.Errorf("Got %v, Want, %v", got, want)
	}
}