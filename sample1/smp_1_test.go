package sample1

import(
	"testing"
)

func TestRetName(t *testing.T){
	actual := RetName("ddddddO")
	expect := "ddddddOです。宜しくお願いいたします。"
	if actual != expect {
		t.Errorf("actual: %v\necept: %v", actual, expect)
	}	
}