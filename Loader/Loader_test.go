package Loader

import "testing"

func httpVars(t *testing.T, profile string) map[string]string {
	t.Helper()
	return GenerateHTTPVaribles("acme-email.com", "base64url", "", "", "", "", "", "", profile, false)
}

// The Slack profile hardcoded its stager URIs ("/messages/DALBNSf25" and
// "/messages/DALBNSF25"), so every profile SourcePoint ever produced for that
// template carried the same two paths. They must now vary per run.
func TestStagerURIsVaryBetweenProfiles(t *testing.T) {
	const runs = 50
	seen := make(map[string]bool, runs*2)
	for i := 0; i < runs; i++ {
		v := httpVars(t, "2")
		for _, key := range []string{"stager_x86", "stager_x64"} {
			got := v[key]
			if got == "" {
				t.Fatalf("%s was not generated", key)
			}
			if seen[got] {
				t.Errorf("%s repeated across generated profiles: %q", key, got)
			}
			seen[got] = true
		}
	}
}

// The GoToMeeting profile derived both stager URIs from a single value, so the
// x86 and x64 staging paths were byte-identical.
func TestStagerURIsDifferPerArchitecture(t *testing.T) {
	for i := 0; i < 50; i++ {
		v := httpVars(t, "3")
		if v["stager_x86"] == v["stager_x64"] {
			t.Fatalf("stager URIs identical across architectures: %q", v["stager_x86"])
		}
	}
}

// UValue appears in the beacon's own check-in traffic (the "U="/"REF=ID="
// prepends and the wla42 cookie). Deriving a stager URI from it linked the
// staging request and the check-ins by a shared unique token.
func TestStagerURIsAreIndependentOfCheckinToken(t *testing.T) {
	for i := 0; i < 50; i++ {
		v := httpVars(t, "3")
		for _, key := range []string{"stager_x86", "stager_x64"} {
			if v[key] == v["UValue"] {
				t.Fatalf("%s reuses UValue, which also appears in check-in traffic: %q", key, v[key])
			}
		}
	}
}
