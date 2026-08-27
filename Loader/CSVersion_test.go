package Loader

import (
	"strings"
	"testing"

	"github.com/Tylous/SourcePoint/Struct"
)

func TestParseCSVersion(t *testing.T) {
	cases := []struct {
		in    string
		major int
		minor int
	}{
		{"", 4, 13},
		{"4.13", 4, 13},
		{"4.13+", 4, 13},
		{"4.12", 4, 12},
		{" 4.9 ", 4, 9},
		{"5", 5, 0},
	}
	for _, c := range cases {
		got := ParseCSVersion(c.in)
		if got.Major != c.major || got.Minor != c.minor {
			t.Errorf("ParseCSVersion(%q) = %d.%d, want %d.%d", c.in, got.Major, got.Minor, c.major, c.minor)
		}
	}
}

func TestCSVersionAtLeast(t *testing.T) {
	cases := []struct {
		version string
		want    bool
	}{
		{"4.13", true},
		{"4.14", true},
		{"5.0", true},
		{"4.12", false},
		{"4.9", false},
		{"3.14", false},
	}
	for _, c := range cases {
		if got := ParseCSVersion(c.version).AtLeast(4, 13); got != c.want {
			t.Errorf("ParseCSVersion(%q).AtLeast(4, 13) = %v, want %v", c.version, got, c.want)
		}
	}
}

// Every PE clone entry must yield a name, since the summary line reports it and
// the 4.13 path strips the directive that carries it.
func TestPECloneNameReadsEveryEntry(t *testing.T) {
	for i, pe := range Struct.Peclone_list {
		name := PECloneName(pe)
		if name == "" || name == "unknown" {
			t.Errorf("Peclone_list[%d]: could not read the module name", i)
		}
		if !strings.HasSuffix(strings.ToLower(name), ".dll") {
			t.Errorf("Peclone_list[%d]: name %q does not look like a module", i, name)
		}
	}
}

// 4.13 rejects stage.name, so the directive has to go and nothing else may.
// The entries are not uniform (four of the thirty carry no image_size
// directives), so this compares against each entry rather than against a fixed
// list of directives.
func TestStripPECloneNameRemovesOnlyTheNameDirective(t *testing.T) {
	for i, pe := range Struct.Peclone_list {
		stripped := StripPECloneName(pe)
		if strings.Contains(stripped, "set name") {
			t.Errorf("Peclone_list[%d]: set name survived stripping", i)
		}
		for _, line := range strings.Split(pe, "\n") {
			if strings.TrimSpace(line) == "" || strings.Contains(line, "set name") {
				continue
			}
			if !strings.Contains(stripped, line) {
				t.Errorf("Peclone_list[%d]: stripping also removed %q", i, strings.TrimSpace(line))
			}
		}
	}
}

// The beacon has outgrown the image_size values baked into most clone entries,
// so 4.13 rejects them with "must be larger than N bytes". They have to go, and
// nothing else may go with them.
func TestStripPECloneImageSizeRemovesOnlyThoseDirectives(t *testing.T) {
	for i, pe := range Struct.Peclone_list {
		stripped := StripPECloneImageSize(pe)
		if strings.Contains(stripped, "image_size_x86") || strings.Contains(stripped, "image_size_x64") {
			t.Errorf("Peclone_list[%d]: an image_size directive survived stripping", i)
		}
		for _, line := range strings.Split(pe, "\n") {
			if strings.TrimSpace(line) == "" || strings.Contains(line, "image_size_x") {
				continue
			}
			if !strings.Contains(stripped, line) {
				t.Errorf("Peclone_list[%d]: stripping also removed %q", i, strings.TrimSpace(line))
			}
		}
	}
}

// Together, the two strips have to leave a clone block 4.13 accepts: no name,
// no image_size, but the rest of the masquerade intact.
func TestStrippedCloneKeepsTheRemainingMasquerade(t *testing.T) {
	for i, pe := range Struct.Peclone_list {
		stripped := StripPECloneImageSize(StripPECloneName(pe))
		for _, gone := range []string{"set name", "image_size_x86", "image_size_x64"} {
			if strings.Contains(stripped, gone) {
				t.Errorf("Peclone_list[%d]: %q survived", i, gone)
			}
		}
		for _, keep := range []string{"set checksum", "set compile_time", "set entry_point", "set rich_header"} {
			if !strings.Contains(stripped, keep) {
				t.Errorf("Peclone_list[%d]: %q did not survive", i, keep)
			}
		}
	}
}

// set pipename and set ssh_pipename must not be mistaken for set name.
func TestStripPECloneNameLeavesPipenamesAlone(t *testing.T) {
	in := "set pipename \"foo\";\nset ssh_pipename \"bar\";\nset name \"baz.dll\";\n"
	got := StripPECloneName(in)
	if strings.Contains(got, `set name "baz.dll"`) {
		t.Error("set name was not stripped")
	}
	if !strings.Contains(got, "set pipename") || !strings.Contains(got, "set ssh_pipename") {
		t.Errorf("a pipename directive was stripped: %q", got)
	}
}
