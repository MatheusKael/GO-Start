package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Michael Jackson"
	want := regexp.MustCompile(`\b` + name + `\b`)
    msg, err := Hello("Michael Jackson") 
    if !want.MatchString(msg) || err != nil { 
        t.Fatalf(`Hello("Michael Jackson") = %q,  %v, want match for %#q, nil`, msg, err, want)
    }
}


func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")

    if msg != "" || nil == err { 
        t.Fatalf(`hello("") = %q, %v want "", error`, msg, err)
    }
}
