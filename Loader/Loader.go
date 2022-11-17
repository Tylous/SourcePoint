package Loader

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"text/template"
	"math/rand"
	"time"
	"net"
	"encoding/binary"

	"github.com/Tylous/SourcePoint/Struct"
	"github.com/Tylous/SourcePoint/Utils"
)

type FlagOptions struct {
	sleeptime                string
	jitter                   string
	useragent                string
	uri                      string
	customuri                string
	beacon_PE                string
	processinject_min_alloc  string
	Post_EX_Process_Name     string
	metadata                 string
	injector                 string
	Host                     string
	outFile                  string
	Profile                  string
	ProfilePath              string
	cert_password            string
	custom_cert              string
	CDN                      string
	Yaml                     string
	tasks_max_size           string
	tasks_proxy_max_size     string
	tasks_dns_proxy_max_size string
	maxdns					 string
}

type Beacon_Com struct {
	Variables map[string]string
}
type Beacon_Stage_p1 struct {
	Variables map[string]string
}
type Beacon_Stage_p2 struct {
	Variables map[string]string
}
type Beacon_Stage_p3 struct {
	Variables map[string]string
}
type Process_Inject struct {
	Variables map[string]string
}
type Beacon_PostEX struct {
	Variables map[string]string
}
type Beacon_GETPOST_Profile struct {
	Variables map[string]string
}
type Beacon_GETPOST struct {
	Variables map[string]string
}
type Beacon_SSL struct {
	Variables map[string]string
}

var num_Profile int
var Post bool

func GenerateOptions(stage string, sleeptime string, jitter string, useragent string, uri string, customuri string, customuriGET string, customuriPOST string, beacon_PE string, processinject_min_alloc string, Post_EX_Process_Name string, metadata string, injector string, Host string, Profile string, ProfilePath string, outFile string, custom_cert string, cert_password string, CDN string, CDN_Value string, datajitter string, Keylogger string, Forwarder bool, tasks_max_size string, tasks_proxy_max_size string, tasks_dns_proxy_max_size string, maxdns string) {
	Beacon_Com := &Beacon_Com{}
	Beacon_Stage_p1 := &Beacon_Stage_p1{}
	Beacon_Stage_p2 := &Beacon_Stage_p2{}
	Beacon_Stage_p3 := &Beacon_Stage_p3{}
	Process_Inject := &Process_Inject{}
	Beacon_PostEX := &Beacon_PostEX{}
	Beacon_GETPOST := &Beacon_GETPOST{}
	Beacon_GETPOST_Profile := &Beacon_GETPOST_Profile{}
	Beacon_SSL := &Beacon_SSL{}
	var HostStageMessage string

	fmt.Println("[*] Preparing Varibles...")
	HostStageMessage, Beacon_Com.Variables = GenerateComunication(stage, sleeptime, jitter, useragent, datajitter, tasks_max_size, tasks_proxy_max_size, tasks_dns_proxy_max_size, maxdns)
	Beacon_PostEX.Variables = GeneratePostProcessName(Post_EX_Process_Name, Keylogger)
	Beacon_GETPOST.Variables = GenerateHTTPVaribles(Host, metadata, uri, customuri, customuriGET, customuriPOST, CDN, CDN_Value, Profile, Forwarder)
	Beacon_Stage_p2.Variables = GeneratePE(beacon_PE)
	Process_Inject.Variables = GenerateProcessInject(processinject_min_alloc, injector)
	Beacon_GETPOST_Profile.Variables, Beacon_SSL.Variables = GenerateProfile(Profile, CDN, CDN_Value, cert_password, custom_cert, ProfilePath, Host)
	fmt.Println("[*] Building Profile...")
	Build(custom_cert, cert_password, outFile, Beacon_Com, Beacon_Stage_p1, Beacon_Stage_p2, Beacon_Stage_p3, Process_Inject, Beacon_PostEX, Beacon_GETPOST, Beacon_GETPOST_Profile, Beacon_SSL)
	fmt.Println(HostStageMessage)
	PE := strings.Split(Beacon_Stage_p2.Variables["pe"], `;`)
	PE_Name := strings.Split(PE[len(PE)-3], `"`)
	fmt.Println("[*] Beacon DLL Spoofed To: " + PE_Name[1])
	PEX := strings.Split(Beacon_PostEX.Variables["Post_EX_Process_Name"], `sysnative\\`)
	PEX_Name := PEX[1]
	fmt.Println("[*] Post-Ex Process Name: " + PEX_Name[:(len(PEX_Name)-3)])
	fmt.Println("[!] Beacon Shellcode Will Obfuscate Beacon in Memory Prior to Sleeping")
	Name, _ := strconv.Atoi(Profile)
	fmt.Println("[*] Seleted Profile: " + Struct.Profile_Names[Name])
	fmt.Println("[+] Profile Generated: " + outFile)
	fmt.Println("[+] Happy Hacking")
}

func GetIpFromCidr(netw string) string {
	_, ipv4Net, err := net.ParseCIDR(netw)
	if err != nil {
		log.Fatal(err)
	}
	mask := binary.BigEndian.Uint32(ipv4Net.Mask)
	start := binary.BigEndian.Uint32(ipv4Net.IP)
	finish := (start & mask) | (mask ^ 0xffffffff)
	var hosts []string
	for i := start + 1; i <= finish-1; i++ {
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		hosts = append(hosts, ip.String())
	}
	rand.Seed(time.Now().Unix())
	return hosts[rand.Intn(len(hosts))]
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func GenerateRandomString(length int) string {
	var charset = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
	  b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
  }

func GenerateComunication(stage, sleeptime, jitter, useragent, datajitter string, tasks_max_size string, tasks_proxy_max_size string, tasks_dns_proxy_max_size string, maxdns string) (string, map[string]string) {
	Beacon_Com := &Beacon_Com{}
	Beacon_Com.Variables = make(map[string]string)
	var HostStageMessage string
	if stage == "False" || stage == "false" {
		Beacon_Com.Variables["stage"] = "false"
		HostStageMessage = "[!] Host Staging Is Disabled - Staged Payloads Are Not Available But Your Beacon Payload Is Not Available To Anyone That Connects"
	} else {
		Beacon_Com.Variables["stage"] = "true"
		HostStageMessage = "[!] Host Staging Is Enabled - Staged Payloads Are Available But Your Beacon Payload Is Available To Anyone That Connects To Your Server To Request It"
	}
	if sleeptime != "" {
		Beacon_Com.Variables["sleep"] = sleeptime + "000"
	} else if sleeptime == "" {
		Beacon_Com.Variables["sleep"] = Utils.GenerateNumer(30, 75) + "000"
	}
	if jitter != "" {
		Beacon_Com.Variables["jitter"] = jitter
	}
	if jitter == "" {
		Beacon_Com.Variables["jitter"] = Utils.GenerateNumer(10, 40)
	}
	if datajitter != "" {
		Beacon_Com.Variables["datajitter"] = datajitter
	}
	if datajitter == "" {
		Beacon_Com.Variables["datajitter"] = Utils.GenerateNumer(10, 60)
	}
	if tasks_max_size != "" {
		Beacon_Com.Variables["tasks_max_size"] = tasks_max_size
	} else {
		Beacon_Com.Variables["tasks_max_size"] = "1048576"
	}
	if tasks_proxy_max_size != "" {
		Beacon_Com.Variables["tasks_proxy_max_size"] = tasks_proxy_max_size
	} else {
		Beacon_Com.Variables["tasks_proxy_max_size"] = "921600"
	}
	if tasks_dns_proxy_max_size != "" {
		Beacon_Com.Variables["tasks_dns_proxy_max_size"] = tasks_dns_proxy_max_size
	} else {
		Beacon_Com.Variables["tasks_dns_proxy_max_size"] = "71680"
	}
	//DNS Configs
	Beacon_Com.Variables["dns_idle"] = GetIpFromCidr("73.140.245.0/24")
	if(maxdns) != "" {
		Beacon_Com.Variables["maxdns"] = maxdns
	} else {
		Beacon_Com.Variables["maxdns"] = "200"
	}
	Beacon_Com.Variables["dns_sleep"] = "300"
	Beacon_Com.Variables["dns_ttl"] = "10"
	Beacon_Com.Variables["dns_stager_prepend"] = "v=spf1 include:spf.protection.outlook.com -all"
	Beacon_Com.Variables["dns_stager_subhost"] = GenerateRandomString(3) + "."
	Beacon_Com.Variables["beacon"] = GenerateRandomString(3) + "."
	Beacon_Com.Variables["get_A"] = GenerateRandomString(3) + "."
	Beacon_Com.Variables["get_AAAA"] = GenerateRandomString(3) + "."
	Beacon_Com.Variables["get_TXT"] = GenerateRandomString(3) + "."
	Beacon_Com.Variables["put_metadata"] = GenerateRandomString(3) + "."
	Beacon_Com.Variables["put_output"] = GenerateRandomString(3) + "."
	Beacon_Com.Variables["ns_response"] = "idle"

	SSH_Numb, _ := strconv.Atoi(Utils.GenerateNumer(0, 4))
	Beacon_Com.Variables["SSH_Banner"] = Struct.SSH_Banner[SSH_Numb]

	pipe_number, _ := strconv.Atoi(Utils.GenerateNumer(0, 7))
	Beacon_Com.Variables["pipename"] = Struct.Pipename_list[pipe_number] + Utils.GenerateNumer(3000, 9000)
	Beacon_Com.Variables["pipename_stager"] = Struct.Pipename_list[pipe_number] + Utils.GenerateNumer(1000, 9000)
	Beacon_Com.Variables["SSH_pipename"] = Struct.Pipename_list[pipe_number]
	if useragent != "" {
		if useragent == "Win10Chrome" {
			num_agent, _ := strconv.Atoi(Utils.GenerateNumer(0, 9))
			Beacon_Com.Variables["useragent"] = Struct.Useragent_list[num_agent]
		} else if useragent == "Win10Edge" {
			num_agent, _ := strconv.Atoi(Utils.GenerateNumer(9, 16))
			Beacon_Com.Variables["useragent"] = Struct.Useragent_list[num_agent]
		} else if useragent == "Win10IE" {
			num_agent, _ := strconv.Atoi(Utils.GenerateNumer(16, 22))
			Beacon_Com.Variables["useragent"] = Struct.Useragent_list[num_agent]

		} else if useragent == "Win10Firefox" {
			num_agent, _ := strconv.Atoi(Utils.GenerateNumer(22, 27))
			Beacon_Com.Variables["useragent"] = Struct.Useragent_list[num_agent]

		} else if useragent == "Win10" {
			num_agent, _ := strconv.Atoi(Utils.GenerateNumer(0, 27))
			Beacon_Com.Variables["useragent"] = Struct.Useragent_list[num_agent]

		} else if useragent == "Win6.3" {
			num_agent, _ := strconv.Atoi(Utils.GenerateNumer(27, 37))
			Beacon_Com.Variables["useragent"] = Struct.Useragent_list[num_agent]

		} else if useragent == "Linux" {
			num_agent, _ := strconv.Atoi(Utils.GenerateNumer(37, 51))
			Beacon_Com.Variables["useragent"] = Struct.Useragent_list[num_agent]

		} else if useragent == "Mac" {
			num_agent, _ := strconv.Atoi(Utils.GenerateNumer(51, 65))
			Beacon_Com.Variables["useragent"] = Struct.Useragent_list[num_agent]
		} else {
			Beacon_Com.Variables["useragent"] = useragent
		}
	}
	if useragent == "" {
		num_agent, _ := strconv.Atoi(Utils.GenerateNumer(0, 64))
		Beacon_Com.Variables["useragent"] = Struct.Useragent_list[num_agent]

	}

	return HostStageMessage, Beacon_Com.Variables
}

func GeneratePostProcessName(Post_EX_Process_Name, Keylogger string) map[string]string {
	Beacon_PostEX := &Beacon_PostEX{}
	Beacon_PostEX.Variables = make(map[string]string)
	if Post_EX_Process_Name != "" {
		num_PSPN, _ := strconv.Atoi(Post_EX_Process_Name)
		Beacon_PostEX.Variables["Post_EX_Process_Name"] = Struct.Post_EX_Process_Name[(num_PSPN - 1)]
	}
	if Post_EX_Process_Name == "" {
		num_Post_EX_Process_Name, _ := strconv.Atoi(Utils.GenerateNumer(0, 17))
		Beacon_PostEX.Variables["Post_EX_Process_Name"] = Struct.Post_EX_Process_Name[num_Post_EX_Process_Name]
	}
	if Keylogger == "GetAsyncKeyState" || Keylogger == "SetWindowsHookEx" {
		Beacon_PostEX.Variables["Keylogger"] = Keylogger
	} else if Keylogger == "" {
		Beacon_PostEX.Variables["Keylogger"] = "SetWindowsHookEx"
	} else {
	}
	pipe_number, _ := strconv.Atoi(Utils.GenerateNumer(0, 7))
	Beacon_PostEX.Variables["pipename"] = Struct.Pipename_list[pipe_number]

	return Beacon_PostEX.Variables
}

func GenerateHTTPVaribles(Host, metadata, uri, customuri, customuriGET, customuriPOST, CDN, CDN_Value, Profile string, Forwarder bool) map[string]string {
	Beacon_GETPOST := &Beacon_GETPOST{}
	Beacon_GETPOST.Variables = make(map[string]string)
	Beacon_GETPOST.Variables["Host"] = Host
	if Profile == "" {
		num_Profile, _ = strconv.Atoi(Utils.GenerateNumer(1, 5))
	} else {
		num_Profile, _ = strconv.Atoi(Profile)
	}
	if metadata == "base64" {
		Beacon_GETPOST.Variables["metadata_mode"] = metadata
	} else if metadata == "base64url" {
		Beacon_GETPOST.Variables["metadata_mode"] = metadata
	} else if metadata == "netbios" {
		Beacon_GETPOST.Variables["metadata_mode"] = metadata
	} else if metadata == "netbiosu" {
		Beacon_GETPOST.Variables["metadata_mode"] = metadata
	} else if metadata == "" {
		Beacon_GETPOST.Variables["metadata_mode"] = "netbios"
	} else {
		log.Fatal("Error: Please provide a valid metadata option")
	}
	if uri == "" {
		Post = false
		uri := customuri
		if customuriGET != "" && customuriPOST != "" {
			uri = customuriGET
			fmt.Println("[*] GET URI base: " + uri)
		}

		Beacon_GETPOST.Variables["HTTP_GET_URI"] = Utils.GenerateURIValues(1, num_Profile, Post, uri)
		Post = true
		if customuriGET != "" && customuriPOST != "" {
			uri = customuriPOST
			fmt.Println("[*] POST URI base: " + uri)
		}

		Beacon_GETPOST.Variables["HTTP_POST_URI"] = Utils.GenerateURIValues(1, num_Profile, Post, uri)

	}
	if uri != "" {
		num_uri, _ := strconv.Atoi(uri)
		Post = false
		uri := customuri
		if customuriGET != "" && customuriPOST != "" {
			uri = customuriGET
			fmt.Println("[*] GET URI base: " + uri)
		}
		Beacon_GETPOST.Variables["HTTP_GET_URI"] = Utils.GenerateURIValues(num_uri, num_Profile, Post, uri)
		Post = true
		if customuriGET != "" && customuriPOST != "" {
			uri = customuriPOST
			fmt.Println("[*] POST URI base: " + uri)
		}
		Beacon_GETPOST.Variables["HTTP_POST_URI"] = Utils.GenerateURIValues(num_uri, num_Profile, Post, uri)
	}
	if CDN != "" {
		Beacon_GETPOST.Variables["CDN"] = "header \"Cookie\" \"" + CDN + "=" + CDN_Value + "\";"
	}
	if CDN == "" {
		Beacon_GETPOST.Variables["CDN"] = ""
	}

	Beacon_GETPOST.Variables["number64"] = Utils.GenerateNumer(19340, 15370000)
	Beacon_GETPOST.Variables["number86"] = Utils.GenerateNumer(19340, 15370000)

	Beacon_GETPOST.Variables["namprdnumber"] = Utils.GenerateNumer(2, 8)
	Beacon_GETPOST.Variables["maxage"] = Utils.GenerateNumer(172800, 31536001)
	Beacon_GETPOST.Variables["Age"] = Utils.GenerateNumer(1222, 2500)

	Beacon_GETPOST.Variables["UValue"] = Utils.GenerateValue(6, 15)
	Beacon_GETPOST.Variables["CSMValue"] = Utils.GenerateValue(6, 15)

	if Forwarder == true {
		Beacon_GETPOST.Variables["forward"] = "true"
	} else {
		Beacon_GETPOST.Variables["forward"] = "false"
	}

	return Beacon_GETPOST.Variables
}

func GeneratePE(beacon_PE string) map[string]string {
	Beacon_Stage_p2 := &Beacon_Stage_p2{}
	Beacon_Stage_p2.Variables = make(map[string]string)
	if beacon_PE == "" {
		PE_Num, _ := strconv.Atoi(Utils.GenerateNumer(0, 30))
		Beacon_Stage_p2.Variables["pe"] = Struct.Peclone_list[PE_Num]
	}
	if beacon_PE != "" {
		PE_Num, _ := strconv.Atoi(beacon_PE)
		if PE_Num > 30 {
			log.Fatal("Error: Please provide a valid PE number less the 31 option")
		}
		Beacon_Stage_p2.Variables["pe"] = Struct.Peclone_list[(PE_Num - 1)]
	}
	return Beacon_Stage_p2.Variables
}

func GenerateProcessInject(processinject_min_alloc, injector string) map[string]string {
	Process_Inject := &Process_Inject{}
	Process_Inject.Variables = make(map[string]string)
	if processinject_min_alloc == "" {
		Process_Inject.Variables["processinject_min_alloc"] = Utils.GenerateNumer(4096, 57841)
	}
	if processinject_min_alloc != "" {
		processinject_min_alloc_int, _ := strconv.Atoi(processinject_min_alloc)
		if processinject_min_alloc_int < 4096 {
			log.Fatal("Error: Minimum amount of memory to request for injected content needs to be greater than 4096")
		} else {
			Process_Inject.Variables["processinject_min_alloc"] = processinject_min_alloc
		}
	}
	Process_Inject.Variables["ThreadStartNum"] = Utils.GenerateNumer(500, 2500)
	if injector == "NtMapViewOfSection" {
		Process_Inject.Variables["injector"] = injector
	} else if injector == "VirtualAllocEx" {
		Process_Inject.Variables["injector"] = injector
	} else if injector == "HeapAlloc" {
		Process_Inject.Variables["injector"] = injector
	} else {
		log.Fatal("Error: Please provide a valid Process Injector option")
	}
	return Process_Inject.Variables
}

func GenerateProfile(Profile, CDN, CDN_Value, cert_password, custom_cert, ProfilePath, hostname string) (map[string]string, map[string]string) {
	Beacon_GETPOST_Profile := &Beacon_GETPOST_Profile{}
	Beacon_SSL := &Beacon_SSL{}
	Beacon_GETPOST_Profile.Variables = make(map[string]string)
	Beacon_SSL.Variables = make(map[string]string)
	if Profile == "" {
		CNAME := "\nhttps-certificate {\rset CN       \"" + hostname + "\"; #Common Name"
		Beacon_SSL.Variables["Cert"] = CNAME + Struct.Cert[num_Profile-1]
		Beacon_GETPOST_Profile.Variables["Profile"] = Struct.HTTP_GET_POST_list[num_Profile-1]
	}
	if Profile != "" {
		num_Profile, _ = strconv.Atoi(Profile)
		if num_Profile <= 4 {
			CNAME := "\nhttps-certificate {\rset CN       \"" + hostname + "\"; #Common Name"
			Beacon_SSL.Variables["Cert"] = CNAME + Struct.Cert[num_Profile-1]
			Beacon_GETPOST_Profile.Variables["Profile"] = Struct.HTTP_GET_POST_list[(num_Profile - 1)]
			fmt.Println("[!] Self Signed SSL Cerificate Used")
		} else if num_Profile == 6 {
			if CDN == "" {
				log.Fatal("Error: Please provide a CDN value in order to use AzureEdge profiles")
			}
			if cert_password == "" {
				log.Fatal("Error: Please provide a Password value to use this profile")
			}
			if custom_cert == "" {
				log.Fatal("Error: Please provide a Keystore value to use this profile")
			}
			Beacon_SSL.Variables["Cert"] = Struct.Cert[4]
			Beacon_GETPOST_Profile.Variables["Profile"] = Struct.HTTP_GET_POST_list[(num_Profile - 1)]
		} else if num_Profile == 5 || num_Profile == 7 {
			if cert_password == "" {
				log.Fatal("Error: Please provide a Password value to use this profile")
			}
			if custom_cert == "" {
				log.Fatal("Error: Please provide a Keystore value to use this profile")
			}
			Beacon_SSL.Variables["Cert"] = Struct.Cert[4]
			Beacon_GETPOST_Profile.Variables["Profile"] = Struct.HTTP_GET_POST_list[(num_Profile - 1)]
		} else if num_Profile == 8 {
			if cert_password == "" && custom_cert == "" {
				CNAME := "\rhttps-certificate {\rset CN       \"" + hostname + "\"; #Common Name"
				Beacon_SSL.Variables["Cert"] = CNAME + Struct.Cert[0]
				fmt.Println("[!] Self Signed SSL Cerificate Used")
			}
			if cert_password == "" && custom_cert != "" {
				log.Fatal("Error: Please provide a Password value to use this profile")
			}
			if custom_cert == "" && cert_password != "" {
				log.Fatal("Error: Please provide a Keystore value to use this profile")
			}
			if cert_password != "" && custom_cert != "" {
				Beacon_SSL.Variables["Cert"] = Struct.Cert[4]
			}
			Beacon_GETPOST_Profile.Variables["Profile"] = Utils.Readfile(ProfilePath)
		} else {
			log.Fatal("Error: Please provide a Profile number of 8 or less")
		}
	}
	if custom_cert != "" && cert_password != "" {
		fmt.Println("[*] Valid SSL Cerificate Used")
		Beacon_SSL.Variables["Cert"] = Struct.Cert[4]
		if strings.HasSuffix(custom_cert, ".store") {
			Beacon_SSL.Variables["CertName"] = custom_cert
		} else {
			Beacon_SSL.Variables["CertName"] = custom_cert + ".store"
		}
		Beacon_SSL.Variables["Password"] = cert_password

	}
	return Beacon_GETPOST_Profile.Variables, Beacon_SSL.Variables
}

func Build(custom_cert, cert_password, outFile string, Beacon_Com *Beacon_Com, Beacon_Stage_p1 *Beacon_Stage_p1, Beacon_Stage_p2 *Beacon_Stage_p2, Beacon_Stage_p3 *Beacon_Stage_p3, Process_Inject *Process_Inject, Beacon_PostEX *Beacon_PostEX, Beacon_GETPOST *Beacon_GETPOST, Beacon_GETPOST_Profile *Beacon_GETPOST_Profile, Beacon_SSL *Beacon_SSL) {
	var buffer bytes.Buffer

	Beacon_Com_Struct_Template, err := template.New("Beacon_Com").Parse(Struct.Beacon_Com_Struct())
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err := Beacon_Com_Struct_Template.Execute(&buffer, Beacon_Com); err != nil {
		log.Fatal(err)
	}
	first := buffer.String()
	buffer.Reset()
	Beacon_Stage_Struct_p1_Template, err := template.New("Beacon_Stage_p1").Parse(Struct.Beacon_Stage_Struct_p1())
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err := Beacon_Stage_Struct_p1_Template.Execute(&buffer, Beacon_Stage_p1); err != nil {
		log.Fatal(err)
	}
	second := buffer.String()
	buffer.Reset()
	Beacon_Stage_Struct_p2_Template, err := template.New("Beacon_Stage_p2").Parse(Struct.Beacon_Stage_p2_Stuct())
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err := Beacon_Stage_Struct_p2_Template.Execute(&buffer, Beacon_Stage_p2); err != nil {
		log.Fatal(err)
	}
	third := buffer.String()
	buffer.Reset()

	Beacon_Stage_Struct_p3_Template, err := template.New("Beacon_Stage_p3").Parse(Struct.Beacon_Stage_Struct_p3())
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err := Beacon_Stage_Struct_p3_Template.Execute(&buffer, Beacon_Stage_p3); err != nil {
		log.Fatal(err)
	}
	fourth := buffer.String()
	buffer.Reset()

	Process_Inject_Struct_Template, err := template.New("Process_Inject").Parse(Struct.Process_Inject_Struct())
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err := Process_Inject_Struct_Template.Execute(&buffer, Process_Inject); err != nil {
		log.Fatal(err)
	}
	fifth := buffer.String()
	buffer.Reset()

	Beacon_PostEX_Struct_Template, err := template.New("Beacon_PostEX").Parse(Struct.Beacon_PostEX_Struct())
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err := Beacon_PostEX_Struct_Template.Execute(&buffer, Beacon_PostEX); err != nil {
		log.Fatal(err)
	}
	sixth := buffer.String()
	buffer.Reset()

	Beacon_GETPOST_Profile_Struct_Template, err := template.New("Beacon_GETPOST_Profile").Parse(Struct.Beacon_GETPOST_Profile_Struct())
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err := Beacon_GETPOST_Profile_Struct_Template.Execute(&buffer, Beacon_GETPOST_Profile); err != nil {
		log.Fatal(err)
	}
	seventh_profile := buffer.String()
	buffer.Reset()

	Beacon_GETPOST_Struct_Template, err := template.New("Beacon_GETPOST").Parse(seventh_profile)
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err := Beacon_GETPOST_Struct_Template.Execute(&buffer, Beacon_GETPOST); err != nil {
		log.Fatal(err)
	}
	seventh := buffer.String()
	buffer.Reset()

	Beacon_SSL_Template, err := template.New("Beacon_SSL").Parse(Struct.Beacon_SSL_Struct())
	if err != nil {
		log.Fatal(err)

	}
	buffer.Reset()
	if err := Beacon_SSL_Template.Execute(&buffer, Beacon_SSL); err != nil {
		log.Fatal(err)
	}
	eight := buffer.String()
	buffer.Reset()
	if custom_cert != "" && cert_password != "" {
		Beacon_SSL_Template, err := template.New("Beacon_SSL").Parse(eight)
		if err != nil {
			log.Fatal(err)

		}
		buffer.Reset()
		if err := Beacon_SSL_Template.Execute(&buffer, Beacon_SSL); err != nil {
			log.Fatal(err)
		}
		eight = buffer.String()
		buffer.Reset()
	}

	compiled := first + second + third + fourth + fifth + sixth + seventh + eight
	Utils.Writefile(outFile, compiled)
}
