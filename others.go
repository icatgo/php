package php

import (
	"fmt"
	"time"
)

// Echo echo
func Echo(args ...interface{}) {
	fmt.Print(args...)
}

// Uniqid uniqid()
func Uniqid(prefix string) string {
	now := time.Now()
	return fmt.Sprintf("%s%08x%05x", prefix, now.Unix(), now.UnixNano()%0x100000)
}
