package Utils

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "1234567890"
const capital = "ABCDEF"
const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"
const alphanum = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const lowercasealpha = "abcdefghijklmnopqrstuvwxyz"

// rng is seeded once, from crypto/rand, and reused for the life of the process.
//
// Every generator in this file used to call rand.Seed(time.Now().UnixNano())
// on entry. That is deprecated as of go1.20, and it actively works against a
// polymorphic generator: consecutive calls that land inside the same clock tick
// reseed the global source to the same state and hand back byte-identical
// "random" values. That is how a single profile ends up with duplicate URIs.
var rng = rand.New(rand.NewSource(seed()))

func seed() int64 {
	var b [8]byte
	if _, err := crand.Read(b[:]); err != nil {
		return time.Now().UnixNano()
	}
	return int64(binary.LittleEndian.Uint64(b[:]))
}

// randRange returns a random int in [min, max). An empty or inverted range
// returns min rather than panicking inside rand.Intn.
func randRange(min, max int) int {
	if max <= min {
		return min
	}
	return rng.Intn(max-min) + min
}

// RandIndex returns a random index into a slice of length n, i.e. [0, n).
//
// Callers used to hand-write GenerateNumer(0, len(list)-1), which is exclusive
// of its upper bound and so made the last entry of every lookup table
// unreachable. Deriving the bound from len() keeps the tables and the picker in
// sync when entries are added.
func RandIndex(n int) int {
	if n <= 0 {
		return 0
	}
	return rng.Intn(n)
}

func randomString(n int, charset string) string {
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rng.Intn(len(charset))]
	}
	return string(b)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Readfile(inputFile string) string {
	output, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(output)
}

func Writefile(outFile, result string) {
	cf, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	check(err)
	defer cf.Close()
	_, err = cf.Write([]byte(result))
	check(err)
}

func RandStringBytes(n int) string {
	return randomString(n, letters)
}

func VarNumberLength(min, max int) string {
	return randomString(randRange(min, max), letters)
}

func GenerateNumer(min, max int) string {
	return strconv.Itoa(randRange(min, max))
}

func GenerateValue(min, max int) string {
	return randomString(randRange(min, max), alpha)
}

func GenerateSingleValue(num int) string {
	return randomString(num, alphanum)
}

func GenHex() string {
	// Up to 3 hex characters (16^3 == 4096).
	return fmt.Sprintf("%x", randRange(0, 4096))
}

// uriBase returns the path prefix and suffix used by a given profile. Splitting
// this out of GenerateURIValues keeps the retry loop below readable.
func uriBase(profileType int, post bool, customuri string) (prefix, suffix string) {
	switch profileType {
	case 1:
		return "/c/msdownload/update/others/2021/10/", ""
	case 2:
		return "/messages/", ""
	case 3:
		if post {
			return "/rest/2/meetings", ""
		}
		return "/functionalStatus/", ""
	case 4:
		return "/owa/", ""
	case 5:
		return "/safebrowsing/" + GenerateValue(4, 10) + "/", ""
	case 6:
		return "/chat/", ""
	case 7:
		if post {
			return "/n", "/avp/amznussraps/"
		}
		return "/s/", "/field-keywords/"
	default:
		// Profiles 8 and 9 are operator-supplied; anything else has already
		// been rejected by the caller.
		return customuri, ""
	}
}

func GenerateURIValues(numb int, profile_type int, Post bool, customuri string) string {
	baseuri, enduri := uriBase(profile_type, Post, customuri)

	var sb strings.Builder
	sb.WriteString("set uri \"")

	seen := make(map[string]bool, numb)
	// maxAttempts stops a pathological retry loop; with 14+ character segments
	// the rejection paths below are hit vanishingly rarely.
	maxAttempts := numb*100 + 100
	for generated, attempts := 0, 0; generated < numb && attempts < maxAttempts; attempts++ {
		// Segment length: 14-29 for Windows Update, 20-35 for everything else.
		// Preserved from the original rand.Intn(30-14)+14 / +20 expressions.
		min, max := 20, 36
		if profile_type == 1 {
			min, max = 14, 30
		}
		value := randomString(randRange(min, max), alpha)

		// A path segment starting with '-' stands out, so it is rejected. The
		// original loop dropped the URI outright instead of retrying, so
		// -Uri 8 could quietly emit as few as one URI.
		if strings.HasPrefix(value, "-") {
			continue
		}
		uri := baseuri + value + enduri
		if seen[uri] {
			continue
		}
		seen[uri] = true
		sb.WriteString(uri + " ")
		generated++
	}

	sb.WriteString("\";\n")
	return sb.String()
}
