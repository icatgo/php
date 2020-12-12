package php

import "testing"

func TestExecAll(t *testing.T) {
	Exec("", nil, nil)
	System("", nil)
	Passthru("", nil)

}
