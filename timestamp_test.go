package timestamp

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

var verifyTimestampTests = []struct {
	it       string
	increase time.Duration
	expected string
	err      error
}{
	{
		it:       "Should increase by 1 hour",
		increase: 1 * time.Hour,
		expected: "1984/04/04 01:00:00.000",
	},
	{
		it:       "Should increase by 2 hours",
		increase: 2 * time.Hour,
		expected: "1984/04/04 02:00:00.000",
	},
	{
		it:       "Should increase by 1 second",
		increase: 1 * time.Second,
		expected: "1984/04/04 00:00:01.000",
	},
	{
		it:       "Should increase by 1 millisecond",
		increase: 1 * time.Millisecond,
		expected: "1984/04/04 00:00:00.001",
	},
}

func TestFrom(t *testing.T) {
	g := Goblin(t)

	g.Describe("timestamp.From() using fake clock", func() {
		for _, tt := range verifyTimestampTests {

			testdata := tt

			g.It(testdata.it, func(done Done) {
				Stubs()

				_stubClockAdvance(testdata.increase)

				result := From(DepGetTime())

				g.Assert(result).Equal(testdata.expected)

				StubsRestore()
				done()
			})
		}
	})
}

func TestNow(t *testing.T) {
	g := Goblin(t)

	g.Describe("timestamp.Now() using real clock", func() {
		g.It("Should get current timestamp in correct format", func(done Done) {
			result := Now()

			re := regexp.MustCompile(`(?m)^[\d]{4}\/[\d]{2}\/[\d]{2} [\d]{2}:[\d]{2}:[\d]{2}\.[\d]{3}$`)

			g.Assert(true).Equal(re.MatchString(result))

			done()
		})
	})
}

func TestNow_Fake(t *testing.T) {
	g := Goblin(t)

	g.Describe("timestamp.Now() using fake clock", func() {
		for _, tt := range verifyTimestampTests {

			testdata := tt

			g.It(testdata.it, func(done Done) {
				Stubs()

				_stubClockAdvance(testdata.increase)

				result := Now()

				g.Assert(result).Equal(testdata.expected)

				StubsRestore()
				done()
			})
		}
	})
}

func ExampleNow() {
	result := Now()
	fmt.Println(result)
	// result would be similar to "1984/04/04 01:00:00.000"
}

func ExampleFrom() {
	result := From(time.Now())
	fmt.Println(result)
	// result would be similar to "1984/04/04 01:00:00.000"
}
