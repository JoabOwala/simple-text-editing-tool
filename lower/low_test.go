package lower

import (
	"strings"
	"testing"
)

func TestLow(t *testing.T){
	input := []string{"Welcome","TO","(low)","Kisumu","the","LAKE","SIDE","CITY","(low,","3)"}
	want := []string{"Welcome","to","Kisumu","the","lake","side","city"}
	got := Low(input)

	toJoinWant := strings.Join(want, " ")
	toJoinGot := strings.Join(got, " ")

	if toJoinGot != toJoinWant{
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}