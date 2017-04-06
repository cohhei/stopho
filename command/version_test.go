package command

import (
	"fmt"
	"strings"
	"testing"
)

func TestCmdVersion(t *testing.T) {
	app, outStream := InitAppForTests()

	err := app.Run([]string{"stopho", "version"})
	if err != nil {
		t.Fatalf("expected no error from Run, got %s", err)
	}

	expected := fmt.Sprintf("%s version %s", Name, Version)
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("\nexpected: \n%s\n got: \n%s", expected, outStream.String())
	}
}
