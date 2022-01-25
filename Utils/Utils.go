package Utils

import (
	"io/ioutil"
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
const lowercasealpha = "abcdefghijklmnopqrstuvwxyz"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Readfile(inputFile string) string {
	output, err := ioutil.ReadFile(inputFile)
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

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]

	}
	return string(b)
}

func VarNumberLength(min, max int) string {
	var r string
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(max-min) + min
	n := num
	r = RandStringBytes(n)
	return r
}

func GenerateNumer(min, max int) string {

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(max-min) + min
	number := strconv.Itoa(num)
	return number

}

func GenerateValue(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(max-min) + min
	n := num
	b := make([]byte, n)
	for i := range b {
		b[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(b)
}

func GenerateURIValues(numb int, profile_type int, Post bool, customuri string) string {
	var uri string
	var baseuri string
	var enduri string
	var num int
	if profile_type == 1 {
		baseuri = "/c/msdownload/update/others/2021/10/"
	}
	if profile_type == 2 {
		baseuri = "/messages/"
	}
	if profile_type == 3 {
		if Post == false {
			baseuri = "/functionalStatus/"
		} else if Post == true {
			baseuri = "/rest/2/meetings"
		}
	}
	if profile_type == 4 {
		baseuri = "/owa/"
	}
	if profile_type == 5 {
		baseuri = "/safebrowsing/" + GenerateValue(4, 10) + "/"
	}
	if profile_type == 6 {
		baseuri = "/chat/"
	}
	if profile_type == 7 {
		if Post == false {
			baseuri = "/s/"
			enduri = "/field-keywords/"
		} else if Post == true {
			baseuri = "/n"
			enduri = "/avp/amznussraps/"
		}
	}
	if profile_type == 8 {
		baseuri = "" + customuri + ""
	}
	if profile_type == 9 {
		baseuri = "" + customuri + ""
	}
	uri = "set uri \""
	for ii := 1; ii <= numb; ii++ {
		rand.Seed(time.Now().UnixNano())
		if profile_type == 1 {
			num = rand.Intn(30-14) + 14
		} else {
			num = rand.Intn(30-14) + 20
		}
		n := num
		b := make([]byte, n)
		for i := range b {
			b[i] = alpha[rand.Intn(len(alpha))]
		}
		value := string(b)
		if strings.HasPrefix(value, "-") {
			ii = ii
		} else {
			if enduri != "" {
				uri += baseuri + value + enduri + " "
			} else {
				uri += baseuri + value + " "
			}
		}
	}
	uri += "\";\n"
	return uri

}
