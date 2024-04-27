package apostro

import (
	"strings"
	"testing"
)

func TestApos(t *testing.T){
	input := []string{"'"," ","the","lake","side","city"}
	want := []string{"'the","lake","side","city"}
	got := Apos(input)

	toJoinWant := strings.Join(want, " ")
	toJoinGot := strings.Join(got, " ")

	if toJoinGot != toJoinWant{
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}