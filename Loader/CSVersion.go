package Loader

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// DefaultCSVersion is the Cobalt Strike release profiles target when
// -CSVersion is not supplied.
const DefaultCSVersion = "4.13"

// CSVersion is the team server release a profile is being generated for.
//
// Cobalt Strike removes Malleable C2 options between releases, so the generator
// has to know what it is writing for. 4.13 rejects stage.rdll_loader and
// stage.name, both of which earlier releases accept, and a profile carrying
// either one fails to load with "invalid option for <.stage>".
type CSVersion struct {
	Major int
	Minor int
}

// ParseCSVersion accepts "4.13", "4.13+" or "4", and fails loudly on anything
// else rather than silently targeting the wrong release.
func ParseCSVersion(value string) CSVersion {
	if value == "" {
		value = DefaultCSVersion
	}
	parts := strings.SplitN(strings.TrimSuffix(strings.TrimSpace(value), "+"), ".", 3)

	major, err := strconv.Atoi(parts[0])
	if err != nil || major < 0 {
		log.Fatalf("Error: -CSVersion must look like 4.13, got %q", value)
	}
	minor := 0
	if len(parts) > 1 {
		minor, err = strconv.Atoi(parts[1])
		if err != nil || minor < 0 {
			log.Fatalf("Error: -CSVersion must look like 4.13, got %q", value)
		}
	}
	return CSVersion{Major: major, Minor: minor}
}

// AtLeast reports whether the target release is major.minor or newer.
func (v CSVersion) AtLeast(major, minor int) bool {
	if v.Major != major {
		return v.Major > major
	}
	return v.Minor >= minor
}

func (v CSVersion) String() string {
	return strconv.Itoa(v.Major) + "." + strconv.Itoa(v.Minor)
}

// setNameLine matches the "set name" directive inside a PE clone block. The
// \s+name guard keeps it away from set pipename and set ssh_pipename.
var setNameLine = regexp.MustCompile(`(?m)^.*\bset\s+name\s+"([^"]*)".*$\n?`)

// PECloneName returns the module name a PE clone block masquerades as.
//
// This used to be recovered by splitting the block on ";" and indexing len-3,
// which breaks the moment the block gains or loses a directive, as it does when
// the name is stripped for 4.13.
func PECloneName(pe string) string {
	if m := setNameLine.FindStringSubmatch(pe); m != nil {
		return m[1]
	}
	return "unknown"
}

// StripPECloneName removes the "set name" directive from a PE clone block.
// Cobalt Strike 4.13 rejects stage.name while still accepting the rest of the
// clone, so the checksum, compile time, entry point and rich header all
// survive; only the spoofed module name is lost.
func StripPECloneName(pe string) string {
	return setNameLine.ReplaceAllString(pe, "")
}

// imageSizeLine matches the image_size_x86 and image_size_x64 directives inside
// a PE clone block.
var imageSizeLine = regexp.MustCompile(`(?m)^.*\bset\s+image_size_x(?:86|64)\s+"[^"]*".*$\n?`)

// StripPECloneImageSize removes the image_size directives from a PE clone
// block.
//
// Cobalt Strike requires each to be at least the size of the beacon DLL it
// stomps into the image, and rejects the profile otherwise:
//
//	[-] .stage.image_size_x86 must be larger than 372736 bytes
//	[-] .stage.image_size_x64 must be larger than 462848 bytes
//
// The values in Peclone_list are the real sizes of the modules being mimicked,
// and the beacon has outgrown most of them. Raising them to a fixed floor would
// only age out again as the beacon grows with each release, so they are dropped
// and Cobalt Strike sizes the image itself. Four of the thirty entries already
// carry no image_size directives and load fine, which is what this relies on.
func StripPECloneImageSize(pe string) string {
	return imageSizeLine.ReplaceAllString(pe, "")
}
