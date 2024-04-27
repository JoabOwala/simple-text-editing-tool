package punctuation

import (
	"strings"
	"testing"
)
func TestPunc(t *testing.T) {
	input := []string{"the","lake","side","city"," ","'"}
	want := []string{"the","lake","side","city'"}
	got := Punc(input)
	toJoinWant := strings.Join(want, " ")
	toJoinGot := strings.Join(got, " ")
	if toJoinGot != toJoinWant {
		t.Errorf("got, %v, wanted %v", got, want)
	}
}

