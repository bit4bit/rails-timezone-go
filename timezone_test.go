package rails_timezone_go

import (
	"testing"
	"time"
)

func TestMapping(t *testing.T) {
	for railsZone, goZone := range mapping {
		loc, locErr := LoadLocation(railsZone)
		if locErr != nil {
			t.Fatal(locErr)
		}

		if loc == time.UTC && railsZone != "UTC" {
			t.Fatalf("error zone rails %s mapping to %s", railsZone, goZone)
		}
	}
}
