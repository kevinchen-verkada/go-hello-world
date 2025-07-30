package formatting_test

import (
	"testing"

	"golang.kevinchen-verkada/helloworld/formatting"
)

func TestFormattingUppercase(t *testing.T) {
	expected := "HELLO THERE"
	actual := formatting.FormatOutput("Hello there", false /* useBox */)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
