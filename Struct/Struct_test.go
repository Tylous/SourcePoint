package Struct

import (
	"strings"
	"testing"
)

// smartinject is a post-ex option. Emitting it inside the stage block made
// Cobalt Strike reject every generated profile with
// "invalid option for <.stage>".
func TestStageBlockDoesNotSetSmartinject(t *testing.T) {
	if strings.Contains(Beacon_Stage_Struct_p1(), "smartinject") {
		t.Error("the stage block sets smartinject, which is a post-ex option")
	}
}

// Every value in the stage block has to be quoted. An unquoted boolean made
// Cobalt Strike reject the profile with "Unknown statement in <.stage>".
func TestStageSleepMaskIsQuoted(t *testing.T) {
	want := `set sleep_mask "{{.Variables.sleep_mask}}";`
	if !strings.Contains(Beacon_Stage_Struct_p1(), want) {
		t.Errorf("stage block does not contain %s", want)
	}
}

// post-ex is where smartinject belongs, and it has to be driven by the
// -SmartInject flag rather than hardcoded to "true".
func TestPostExSmartinjectIsTemplated(t *testing.T) {
	want := `set smartinject "{{.Variables.smartinject}}";`
	if !strings.Contains(Beacon_PostEX_Struct(), want) {
		t.Errorf("post-ex block does not contain %s", want)
	}
}
