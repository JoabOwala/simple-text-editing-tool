package upper

import (
	"strings"
	"testing"
)

func TestUp(t *testing.T){
	input := []string{"Welcome","to","kisumu","(up)","the","lake","side","city","(up,","3)"}
	want := []string{"Welcome","to","KISUMU","the","LAKE","SIDE","CITY"}
	got := Up(input)

	toJoinWant := strings.Join(want, " ")
	toJointGot := strings.Join(got, " ")

	if toJointGot != toJoinWant{
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}