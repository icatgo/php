package php

import "testing"

func TestUcwords(t *testing.T) {
	result := Ucwords("hello|world")
	println(result)
	if Ucwords("hello|world") != "Hello|World" {
		t.Error("hello world convert error")
	}
}
