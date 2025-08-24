// gigasecond determine the date and time one gigasecond after a certain date
package gigasecond

import "time"

// AddGigasecond determine and return the date and time one billion(giga) second after a certain date
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
