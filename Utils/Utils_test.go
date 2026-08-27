package Utils

import (
	"strconv"
	"strings"
	"testing"
)

// RandIndex must be able to return every index of a table, including the last
// one. The hand-written GenerateNumer(0, len(list)-1) calls it replaces were
// exclusive of their upper bound, so the final entry of every lookup table in
// Struct was unreachable.
func TestRandIndexCoversEveryEntry(t *testing.T) {
	for _, n := range []int{1, 5, 8, 9, 15, 30, 65} {
		seen := make(map[int]bool, n)
		for i := 0; i < n*400; i++ {
			idx := RandIndex(n)
			if idx < 0 || idx >= n {
				t.Fatalf("RandIndex(%d) returned out-of-range index %d", n, idx)
			}
			seen[idx] = true
		}
		if len(seen) != n {
			t.Errorf("RandIndex(%d) only ever produced %d distinct indices; last entry unreachable", n, len(seen))
		}
	}
}

func TestRandIndexHandlesEmptyTable(t *testing.T) {
	if got := RandIndex(0); got != 0 {
		t.Errorf("RandIndex(0) = %d, want 0", got)
	}
	if got := RandIndex(-3); got != 0 {
		t.Errorf("RandIndex(-3) = %d, want 0", got)
	}
}

// An inverted or empty range used to panic inside rand.Intn.
func TestRandRangeDoesNotPanicOnEmptyRange(t *testing.T) {
	if got := randRange(10, 10); got != 10 {
		t.Errorf("randRange(10, 10) = %d, want 10", got)
	}
	if got := randRange(10, 4); got != 10 {
		t.Errorf("randRange(10, 4) = %d, want 10", got)
	}
}

func TestGenerateNumerStaysInRange(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n, err := strconv.Atoi(GenerateNumer(30, 75))
		if err != nil {
			t.Fatalf("GenerateNumer returned a non-number: %v", err)
		}
		if n < 30 || n >= 75 {
			t.Fatalf("GenerateNumer(30, 75) = %d, outside [30, 75)", n)
		}
	}
}

func parseURIs(t *testing.T, set string) []string {
	t.Helper()
	if !strings.HasPrefix(set, `set uri "`) || !strings.HasSuffix(set, "\";\n") {
		t.Fatalf("malformed uri statement: %q", set)
	}
	body := strings.TrimSuffix(strings.TrimPrefix(set, `set uri "`), "\";\n")
	return strings.Fields(body)
}

// GenerateURIValues must emit exactly as many URIs as the operator asked for.
// The original loop dropped any candidate whose random segment started with
// "-" instead of retrying, so `-Uri 8` regularly produced fewer than 8.
func TestGenerateURIValuesReturnsRequestedCount(t *testing.T) {
	for _, profile := range []int{1, 2, 3, 4, 5, 6, 7} {
		for _, want := range []int{1, 3, 8, 20} {
			for _, post := range []bool{false, true} {
				got := parseURIs(t, GenerateURIValues(want, profile, post, ""))
				if len(got) != want {
					t.Errorf("profile %d post=%v: asked for %d URIs, got %d", profile, post, want, len(got))
				}
			}
		}
	}
}

// A profile that repeats the same URI is a free clustering signal, and the
// per-call time-based reseeding made repeats likely on coarse clocks.
func TestGenerateURIValuesAreUnique(t *testing.T) {
	uris := parseURIs(t, GenerateURIValues(50, 2, false, ""))
	seen := make(map[string]bool, len(uris))
	for _, u := range uris {
		if seen[u] {
			t.Fatalf("duplicate URI generated: %s", u)
		}
		seen[u] = true
	}
}

func TestGenerateURIValuesUsesProfileBasePath(t *testing.T) {
	cases := []struct {
		profile int
		post    bool
		prefix  string
	}{
		{1, false, "/c/msdownload/update/others/2021/10/"},
		{2, false, "/messages/"},
		{3, false, "/functionalStatus/"},
		{3, true, "/rest/2/meetings"},
		{4, false, "/owa/"},
		{6, false, "/chat/"},
		{7, false, "/s/"},
		{7, true, "/n"},
	}
	for _, c := range cases {
		for _, u := range parseURIs(t, GenerateURIValues(5, c.profile, c.post, "")) {
			if !strings.HasPrefix(u, c.prefix) {
				t.Errorf("profile %d post=%v: %q does not start with %q", c.profile, c.post, u, c.prefix)
			}
		}
	}
	for _, u := range parseURIs(t, GenerateURIValues(5, 8, false, "/api/v2/")) {
		if !strings.HasPrefix(u, "/api/v2/") {
			t.Errorf("custom profile: %q does not use the supplied base URI", u)
		}
	}
}

// Segments beginning with "-" stand out in traffic and the generator has always
// meant to reject them.
func TestGenerateURIValuesNeverStartASegmentWithDash(t *testing.T) {
	for _, u := range parseURIs(t, GenerateURIValues(100, 8, false, "/x/")) {
		if strings.HasPrefix(strings.TrimPrefix(u, "/x/"), "-") {
			t.Errorf("URI segment starts with '-': %s", u)
		}
	}
}

// Two profiles generated back to back must not be identical. Re-seeding the
// global source from time.Now() on every call meant that calls landing in the
// same clock tick returned the same "random" value.
func TestGeneratorsDoNotRepeatWithinAClockTick(t *testing.T) {
	const draws = 200
	values := make(map[string]bool, draws)
	for i := 0; i < draws; i++ {
		values[GenerateValue(6, 15)] = true
	}
	if len(values) < draws*9/10 {
		t.Errorf("GenerateValue produced only %d distinct values out of %d draws", len(values), draws)
	}
}
