package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/relvacode/iso8601"
	"log"
	"os"
	"strings"
	"time"
)

func promptYN(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt, " [Ynq]: ")
		text, _ := reader.ReadString('\n')
		ltext := strings.TrimSpace(strings.ToLower(text))
		if ltext == "q" {
			os.Exit(0)
		}
		if ltext == "y" || ltext == "yes" || ltext == "" {
			return true
		}
		if ltext == "n" || ltext == "no" {
			return false
		}
		fmt.Println("Please type 'y' [or press Enter] or 'n', or 'q' to quit")
	}
}

func parseDates(okDate, badDate string) (ok time.Time, bad time.Time) {
	var err error
	if badDate == "today" {
		bad = time.Now()
	} else {
		if bad, err = iso8601.Parse([]byte(badDate)); err != nil {
			log.Fatalf("date parse error for %s: %v", badDate, err)
		}
	}
	if ok, err = iso8601.Parse([]byte(okDate)); err != nil {
		log.Fatalf("date parse error for %s: %v", badDate, err)
	}
	return ok, bad
}

func bisect(a, b time.Time) time.Time {
	return a.Add(b.Sub(a) / 2)
}

func format(t time.Time) string {
	return strings.Split(t.Format(time.RFC3339), "T")[0]
}

func announce(ok, mid, bad time.Time) {
	fmt.Printf("\nok    : %s\nbisect: %s <=== \nbad   : %s\n\n",
		format(ok),
		format(mid),
		format(bad),
	)
}

func main() {
	okDate := flag.String("ok-date", "", "date that this thing was last OK (YYYY-MM-DD format)")
	badDate := flag.String("bad-date", "today", "date that this thing stopped being OK (YYYY-MM-DD format, or 'today')")
	flag.Parse()

	ok, bad := parseDates(*okDate, *badDate)

	for {
		midpoint := bisect(ok, bad)
		announce(ok, midpoint, bad)
		if promptYN("is this bisect point OK?") {
			ok = midpoint
		} else {
			bad = midpoint
		}
	}
}
