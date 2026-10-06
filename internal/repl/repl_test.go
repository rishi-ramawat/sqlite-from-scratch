package repl

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunExitsOnExitCommand(t *testing.T) {
	var out bytes.Buffer
	if err := Run(strings.NewReader(".exit\n"), &out); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if got := out.String(); got != "db > " {
		t.Fatalf("unexpected output %q", got)
	}
}

func TestRunReportsUnrecognizedInput(t *testing.T) {
	var out bytes.Buffer
	if err := Run(strings.NewReader(".foo\nselect 1\n.exit\n"), &out); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	got := out.String()
	for _, want := range []string{`Unrecognized command: ".foo"`, `Unrecognized statement: "select 1"`} {
		if !strings.Contains(got, want) {
			t.Errorf("output missing %q:\n%s", want, got)
		}
	}
}
