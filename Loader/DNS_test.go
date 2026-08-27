package Loader

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var dnsDirective = regexp.MustCompile(`set\s+(\S+)\s+"([^"]*)";`)

func dnsValues(t *testing.T, block string) map[string]string {
	t.Helper()
	out := make(map[string]string)
	for _, m := range dnsDirective.FindAllStringSubmatch(block, -1) {
		out[m[1]] = m[2]
	}
	return out
}

// DNS is opt-in, so a profile generated without it must be unchanged.
func TestDNSBeaconOmittedWhenDisabled(t *testing.T) {
	if got := GenerateDNSBeacon(false, ""); got != "" {
		t.Errorf("expected no dns-beacon block when DNS is off, got %q", got)
	}
}

// The block shipped carrying Cobalt Strike's documentation example values,
// which are the most heavily signatured DNS C2 labels in existence. None of
// them may survive into a generated profile.
func TestDNSBeaconDropsDocumentationValues(t *testing.T) {
	block := GenerateDNSBeacon(true, "")
	docValues := []string{
		"doc.bc.", "doc.1a.", "doc.4a.", "doc.tx.", "doc.md.", "doc.po.",
		"doc-stg-prepend", "doc-stg-sh.",
	}
	for _, v := range docValues {
		if strings.Contains(block, v) {
			t.Errorf("block still carries the documentation value %q", v)
		}
	}
}

// The get_*/put_* prefixes are how the teamserver distinguishes request types.
// Two sharing a value would break the channel, not just look wrong.
// All seven prefixes have to share a length. Redrawing a collision at a wider
// length made every label after the first collision one character longer.
func TestDNSBeaconLabelsShareALength(t *testing.T) {
	keys := []string{
		"beacon", "get_A", "get_AAAA", "get_TXT",
		"put_metadata", "put_output", "dns_stager_subhost",
	}
	for i := 0; i < 200; i++ {
		v := dnsValues(t, GenerateDNSBeacon(true, ""))
		want := len(v["beacon"])
		for _, k := range keys {
			if len(v[k]) != want {
				t.Fatalf("%s = %q is %d characters, but beacon is %d", k, v[k], len(v[k]), want)
			}
		}
	}
}

func TestDNSBeaconLabelsAreDistinct(t *testing.T) {
	keys := []string{
		"beacon", "get_A", "get_AAAA", "get_TXT",
		"put_metadata", "put_output", "dns_stager_subhost",
	}
	for i := 0; i < 50; i++ {
		v := dnsValues(t, GenerateDNSBeacon(true, ""))
		owner := make(map[string]string, len(keys))
		for _, k := range keys {
			got := v[k]
			if got == "" {
				t.Fatalf("%s was not set", k)
			}
			if prev, dup := owner[got]; dup {
				t.Fatalf("%s and %s share the prefix %q", prev, k, got)
			}
			owner[got] = k
		}
	}
}

func TestDNSBeaconVariesBetweenProfiles(t *testing.T) {
	seen := make(map[string]bool)
	for i := 0; i < 25; i++ {
		v := dnsValues(t, GenerateDNSBeacon(true, ""))
		if seen[v["beacon"]] {
			t.Fatalf("beacon prefix repeated across profiles: %q", v["beacon"])
		}
		seen[v["beacon"]] = true
	}
}

// dns_idle is the "no tasks" sentinel, so it stays an operator decision rather
// than something the generator randomizes.
func TestDNSIdleDefaultsAndHonoursOperatorValue(t *testing.T) {
	if got := dnsValues(t, GenerateDNSBeacon(true, ""))["dns_idle"]; got != "0.0.0.0" {
		t.Errorf("dns_idle default = %q, want 0.0.0.0", got)
	}
	if got := dnsValues(t, GenerateDNSBeacon(true, "8.8.4.4"))["dns_idle"]; got != "8.8.4.4" {
		t.Errorf("dns_idle = %q, want 8.8.4.4", got)
	}
}

// Cobalt Strike rejects the profile outright unless dns_max_txt is divisible
// by four. The commented-out block this feature replaced carried 199, which is
// not, so the value has to be checked rather than inherited.
func TestDNSMaxTXTIsDivisibleByFour(t *testing.T) {
	got := dnsValues(t, GenerateDNSBeacon(true, ""))["dns_max_txt"]
	n, err := strconv.Atoi(got)
	if err != nil {
		t.Fatalf("dns_max_txt = %q, which is not a number", got)
	}
	if n%4 != 0 {
		t.Errorf("dns_max_txt = %d, which Cobalt Strike rejects because it is not divisible by four", n)
	}
}

// c2lint warns when a prefix runs past eight characters, because every
// indicator character is data space lost from each query.
func TestDNSLabelsStayWithinTheRecommendedLength(t *testing.T) {
	keys := []string{
		"beacon", "get_A", "get_AAAA", "get_TXT",
		"put_metadata", "put_output", "dns_stager_subhost",
	}
	for i := 0; i < 100; i++ {
		v := dnsValues(t, GenerateDNSBeacon(true, ""))
		for _, k := range keys {
			if len(v[k]) > 8 {
				t.Fatalf("%s = %q is %d characters, over the 8 character maximum", k, v[k], len(v[k]))
			}
		}
	}
}

// Every label must be a syntactically valid lowercase DNS label sequence.
func TestDNSBeaconLabelsAreValid(t *testing.T) {
	valid := regexp.MustCompile(`^([a-z][a-z0-9]*\.)+$`)
	keys := []string{
		"beacon", "get_A", "get_AAAA", "get_TXT",
		"put_metadata", "put_output", "dns_stager_subhost",
	}
	for i := 0; i < 50; i++ {
		v := dnsValues(t, GenerateDNSBeacon(true, ""))
		for _, k := range keys {
			if !valid.MatchString(v[k]) {
				t.Errorf("%s = %q is not a valid dotted label sequence", k, v[k])
			}
		}
	}
}
