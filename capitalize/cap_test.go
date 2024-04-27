package capitalize

import (
	"strings"
	"testing"
)

func TestLow(t *testing.T){
	input := []string{"Welcome","to","kisumu","(cap)","the","lake","side","city","(cap,","3)"}
	want := []string{"Welcome","to","Kisumu","the","Lake","Side","City"}
	got := Cap(input)

	toJoinWant := strings.Join(want, " ")
	toJoinGot := strings.Join(got, " ")

	if toJoinGot != toJoinWant{
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}