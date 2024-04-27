package punctuation

import (
	"strings"
	"testing"
)

func TestPunc(t *testing.T){
	input := []string{"This","is", "Kisumu"," ","...","the","lake","side","city"," ","'"}
	want := []string{"This","is", "Kisumu","...","the","lake","side","city","'"}
	got :=Punc(input)

	toJoinWant := strings.Join(want, " ")
	toJoinGot := strings.Join(got, " ")

	if toJoinGot != toJoinWant {
		t.Errorf("Got: %v, Want: %v", got,want)
	}
}