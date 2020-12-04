package php

import "time"

// Time time()
func Time() int64 {
	return time.Now().Unix()
}

// Date date()
// Date("01/04/2013 13:14:20 PM", 1520131400)
func Date(format string, timestamp int64) string {
	return time.Unix(timestamp, 0).Format(format)
}

// Sleep sleep()
func Sleep(seconds int64) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

// Usleep usleep()
func Usleep(microSeconds int64) {
	time.Sleep(time.Duration(microSeconds) * time.Microsecond)
}
