package Struct

import (
	"regexp"
	"sort"
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

// The Slack profile appends __ar_v4 to the Cookie header in both directions.
// The http-post side was missing the ";" separator, so the request emitted
// _ga=GA1.2.875__ar_v4=... as one mangled cookie value while http-get emitted
// it correctly, leaving GET and POST from the same host visibly inconsistent.
// A profile impersonates one origin, so its responses must not disagree about
// what is serving them. The Outlook.Live http-stager response declared
// Server: nginx while its beacon traffic declared Microsoft-IIS/10.0, which
// hands a defender a free correlation between two responses from the same host.
// Real Outlook Web Access is IIS.
func TestServerHeaderIsConsistentWithinEachProfile(t *testing.T) {
	serverHeader := regexp.MustCompile(`header "Server" "([^"]*)"`)
	for i, profile := range HTTP_GET_POST_list {
		seen := make(map[string]bool)
		for _, m := range serverHeader.FindAllStringSubmatch(profile, -1) {
			seen[m[1]] = true
		}
		if len(seen) > 1 {
			values := make([]string, 0, len(seen))
			for v := range seen {
				values = append(values, v)
			}
			sort.Strings(values)
			t.Errorf("HTTP_GET_POST_list[%d] declares more than one Server value: %v", i, values)
		}
	}
}

func TestCookieFragmentsKeepTheirSeparator(t *testing.T) {
	for i, profile := range HTTP_GET_POST_list {
		if strings.Contains(profile, `append "__ar_v4=`) {
			t.Errorf("HTTP_GET_POST_list[%d]: __ar_v4 is appended with no leading separator, which mangles the Cookie value", i)
		}
	}
}
