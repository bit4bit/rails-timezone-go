rails-timezone-go
=================

Convert between Rails ActiveSupport TimeZone names and the IANA Time Zone Database format, for GO

Inspired and copy from https://github.com/davidwood/rails-timezone-js.

Thanks.

# INSTALL

~~~
 go get github.com/bit4bit/rails-timezone-go
~~~

# HOW USE

~~~go
package main

import (
	"github.com/bit4bit/rails-timezone-go"
	"fmt"
)

func main() {
	loc, err := rails_timezone_go.LoadLocation)
	if err != nil {
		...
	}
	fmt.Print(loc)
}
~~~