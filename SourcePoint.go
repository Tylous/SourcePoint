package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	Loader "github.com/Tylous/SourcePoint/Loader"

	"gopkg.in/yaml.v2"
)

type FlagOptions struct {
	stage                    string
	sleeptime                string
	jitter                   string
	useragent                string
	uri                      string
	customuri                string
	customuriGET             string
	customuriPOST            string
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
	CDN_Value                string
	Datajitter               string
	Keylogger                string
	Forwarder                bool
	tasks_max_size           string
	tasks_proxy_max_size     string
	tasks_dns_proxy_max_size string
	syscall_method           string
	httplib                  string
	threadspoof              bool
	Yaml                     string
	beacongate               string
	eaf_bypass               bool
	rdll_use_syscalls        bool
	copy_pe_header           bool
	rdll_loader              string
	transform_obfuscate      string
	smartinject              bool
	sleep_mask               bool
}

type conf struct {
	Host                 string `yaml:"Host"`
	Stage                string `yaml:"Stage"`
	Keystore             string `yaml:"Keystore"`
	Password             string `yaml:"Password"`
	Metadata             string `yaml:"Metadata"`
	Injector             string `yaml:"Injector"`
	Outfile              string `yaml:"Outfile"`
	PE_Clone             string `yaml:"PE_Clone"`
	Profile              string `yaml:"Profile"`
	Post_EX_Process_Name string `yaml:"Post-EX Processname"`
	ProfilePath          string `yaml:"ProfilePath"`
	Allocation           string `yaml:"allocation"`
	Jitter               string `yaml:"Jitter"`
	Debug                bool   `yaml:"Debug"`
	Sleep                string `yaml:"Sleep"`
	Uri                  string `yaml:"Uri"`
	Customuri            string `yaml:"Customuri"`
	CustomuriGET         string `yaml:"CustomuriGET"`
	CustomuriPOST        string `yaml:"CustomuriPOST"`
	CDN                  string `yaml:"CDN"`
	CDN_Value            string `yaml:"CDN_Value"`
	Useragent            string `yaml:"Useragent"`
	Datajitter           string `yaml:"Datajitter"`
	Keylogger            string `yaml:"Keylogger"`
	Forwarder            bool   `yaml:"Forwarder"`
	TasksMaxSize         string `yaml:"TasksMaxSize"`
	TasksProxyMaxSize    string `yaml:"TasksProxyMaxSize"`
	TasksDnsProxyMaxSize string `yaml:"TasksDnsProxyMaxSize"`
	Syscall_method       string `yaml:"Syscall_method"`
	Httplib              string `yaml:"Httplib"`
	Threadspoof          bool   `yaml:"ThreadSpoof"`
	BeaconGate           string `yaml:"BeaconGate"`
	EafBypass            bool   `yaml:"EafBypass"`
	RdllUseSyscalls      bool   `yaml:"RdllUseSyscalls"`
	Copy_PE_Header       bool   `yaml:"CopyPEHeader"`
	RdllLoader           string `yaml:"RdllLoader"`
	TransformObfuscate   string `yaml:"TransformObfuscate"`
	SmartInject          bool   `yaml:"SmartInject"`
	SleepMask            bool   `yaml:"SleepMask"`
}

func (c *conf) getConf(yamlfile string) *conf {

	yamlFile, err := ioutil.ReadFile(yamlfile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func options() *FlagOptions {
	sleeptime := flag.String("Sleep", "", "Initial beacon sleep time")
	stage := flag.String("Stage", "false", "Disable host staging (Default: False)")
	jitter := flag.String("Jitter", "", "Jitter percentage for beacon call home")
	useragent := flag.String("Useragent", "", `UserAgent string for the beacon to use (Leave blank to randomly select one):
[*] Win10Chrome
[*] Win10Edge
[*] Win10IE
[*] Win10
[*] Win6.3
[*] Linux
[*] Mac`)
	uri := flag.String("Uri", "", "The number URIs a profile for beacons to choose from")
	customuri := flag.String("Customuri", "", "The base URI for custom HTTP GET/POST profile - Cannot be used with CustomuriGET or CustomuriPOST")
	customuriGET := flag.String("CustomuriGET", "", "The base URI for custom HTTP GET profile - Must be used with CustomuriPOST")
	customuriPOST := flag.String("CustomuriPOST", "", "The base URI for custom HTTP POST profile - Must be used with CustomuriGET")
	beacon_PE := flag.String("PE_Clone", "", `PE file beacon will mimic (Use the number):
[1] ActivationManager.dll
[2] audioeng.dll
[3] AzureSettingSyncProvider.dll
[4] BingMaps.dll
[5] DIAGCPL.dll
[6] EDGEHTML.dll
[7] FILEMGMT.dll
[8] FIREWALLCONTROLPANEL.dll
[9] GPSVC.dll
[10] gpupvdev.dll
[11] libcrypto.dll
[12] srvcli.dll
[13] srvsvc.dll
[14] Windows.Storage.Search.dll
[15] Windows.System.Diagnostics.dll
[16] Windows.System.Launcher.dll
[17] Windows.System.SystemManagement.dll
[18] Windows.UI.BioFeedback.dll
[19] Windows.UI.BlockedShutdown.dll
[20] Windows.UI.Core.TextInput.DLL
[21] winsqlite3.dll
[22] WMNetMgr.DLL
[23] wwanapi.dll
[24] WWANSVC.DLL
[25] wow64win.dll
[26] wow64.dll
[27] ctiuser.dll (Carbon Black's DLL)
[28] InProcessClient.dll (SentinelOne's DLL)
[29] umppc.dll (CrowdStrike's DLL)
[30] CyMemDef64.dll (Cylance's DLL)`)
	processinject_min_alloc := flag.String("Allocation", "", "Minimum amount of memory to request for injected content (must be higher than 4096)")
	Post_EX_Process_Name := flag.String("PostEX_Name", "", `File Post-Ex activities will spawn and inject into (Use the number):
[1] WerFault.exe
[2] WWAHost.exe
[3] choice.exe
[4] bootcfg.exe
[5] dtdump.exe
[6] expand.exe
[7] fsutil.exe
[8] gpupdate.exe
[9] gpresult.exe
[10] logman.exe
[11] mcbuilder.exe
[12] mtstocom.exe
[13] pcaui.exe
[14] powercfg.exe
[15] svchost.exe`)
	Profile := flag.String("Profile", "", `HTTP GET/POST profile (Use the number):
[1] Windowsupdate
[2] Slack
[3] Gotomeeting
[4] Outlook.Live
[5] Safebrowsing [Cloudfront Compatible]
[6] AzureEdge [AzureEdge Compatible]
[7] Field-Keyword [Cloudfront Compatible]
[8] Custom (Used with ProfilePath)`)
	ProfilePath := flag.String("ProfilePath", "", "Path of custom HTTP GET/POST profile...")
	metadata := flag.String("Metadata", "base64url", `Specifies how to transform and embed metadata into the HTTP request:
[*] base64
[*] base64url
[*] netbios
[*] netbiosu`)
	injector := flag.String("Injector", "", `Select the preferred method to allocate memory in the remote process:
[*] VirtualAllocEx (Great for cross architecture i.e x86 -> x64 and x64->x86)
[*] NtMapViewOfSection (A more stealthly option, however fails over to VirtualAllocEx, generating more events when it does)`)
	Keylogger := flag.String("Keylogger", "", `Select the preferred method the beacon will use to log keystrokes: 
[*] GetAsyncKeyState (Uses GetAsyncKeyState API (Separate DLL for x86/x64 process))
[*] SetWindowsHookEx (Uses SetWindowsHookEx API)`)
	Datajitter := flag.String("Datajitter", "50", "Appends a value to HTTP-Get and HTTP-Post server output")
	Host := flag.String("Host", "", "Team server domain name")
	outFile := flag.String("Outfile", "", "Name of output file")
	custom_cert := flag.String("Keystore", "", "SSL keystore name")
	cert_password := flag.String("Password", "", "SSL certificate password")
	CDN_Value := flag.String("CDN-Value", "", "CDN cookie value (typically used for AzureEdge profiles)")
	CDN := flag.String("CDN", "", "CDN cookie name (typically used for AzureEdge profiles)")
	Forwarder := flag.Bool("Forwarder", false, "Enabled the X-forwarded-For header (Good for when your C2 is behind a redirector)")
	tasks_max_size := flag.String("TasksMaxSize", "", "The maximum size (in bytes) of task(s) and proxy data that can be transferred through a communication channel at a check in")
	tasks_proxy_max_size := flag.String("TasksProxyMaxSize", "", "The maximum size (in bytes) of proxy data to transfer via the communication channel at a check in")
	tasks_dns_proxy_max_size := flag.String("TasksDnsProxyMaxSize", "", "The maximum size (in bytes) of proxy data to transfer via the DNS communication channel at a check in")
	syscall_method := flag.String("Syscall", "None", `Defines the ability to use direct/indirect system calls instead of the standard Windows API functions calls:
[*] None
[*] Direct
[*] Indirect`)
	httplib := flag.String("Httplib", "winhttp", `Select the default HTTP Beacon library:
[*] wininet
[*] winhttp'`)
	threadspoof := flag.Bool("ThreadSpoof", true, "Sets post-ex DLLs to spawn threads with a spoofed start address. These are generated randomly")
	Yaml := flag.String("Yaml", "", "Path to the Yaml config file")
	beacongate := flag.String("BeaconGate", "", "Specify beacon gate options (All, Comms, Core, Cleanup) or specific APIs")
	eaf_bypass := flag.Bool("EafBypass", false, "Enable EAF Bypass")
	rdll_use_syscalls := flag.Bool("RdllUseSyscalls", false, "Use Syscalls for Rdll")
	copy_pe_header := flag.Bool("CopyPEHeader", false, "Copy PE Header")
	rdll_loader := flag.String("RdllLoader", "PrependLoader", "Rdll Loader Options PrependLoader or StompLoader (Older method)")
	transform_obfuscate := flag.String("TransformObfuscate", "", `Transform obfuscate options (comma-separated list):
[*] lznt1
[*] rc4 "64"
[*] xor "32"
[*] base64
Example: "lznt1,rc4 \"64\",xor \"32\",base64"`)
	smartinject := flag.Bool("SmartInject", false, "Enable Smart Inject")
	sleep_mask := flag.Bool("SleepMask", true, "Enable Sleep Mask")
	flag.Parse()
	return &FlagOptions{stage: *stage, sleeptime: *sleeptime, jitter: *jitter, useragent: *useragent, uri: *uri, customuri: *customuri, customuriGET: *customuriGET, customuriPOST: *customuriPOST, beacon_PE: *beacon_PE, processinject_min_alloc: *processinject_min_alloc, Post_EX_Process_Name: *Post_EX_Process_Name, metadata: *metadata, injector: *injector, Host: *Host, Profile: *Profile, ProfilePath: *ProfilePath, outFile: *outFile, custom_cert: *custom_cert, cert_password: *cert_password, CDN: *CDN, CDN_Value: *CDN_Value, Yaml: *Yaml, Datajitter: *Datajitter, Keylogger: *Keylogger, Forwarder: *Forwarder, tasks_max_size: *tasks_max_size, tasks_proxy_max_size: *tasks_proxy_max_size, tasks_dns_proxy_max_size: *tasks_dns_proxy_max_size, syscall_method: *syscall_method, httplib: *httplib, threadspoof: *threadspoof, beacongate: *beacongate, eaf_bypass: *eaf_bypass, rdll_use_syscalls: *rdll_use_syscalls, copy_pe_header: *copy_pe_header, rdll_loader: *rdll_loader, transform_obfuscate: *transform_obfuscate, smartinject: *smartinject, sleep_mask: *sleep_mask}

}

func main() {
	fmt.Println(`
	   _____                            ____        _       __ 
	  / ___/____  __  _______________  / __ \____  (_)___  / /_
	  \__ \/ __ \/ / / / ___/ ___/ _ \/ /_/ / __ \/ / __ \/ __/
	 ___/ / /_/ / /_/ / /  / /__/  __/ ____/ /_/ / / / / / /_  
	/____/\____/\__,_/_/   \___/\___/_/    \____/_/_/ /_/\__/  
  							(@Tyl0us)
															 `)

	opt := options()
	var c conf
	if opt.Yaml != "" {
		c.getConf(opt.Yaml)
	}

	if opt.Yaml != "" {
		opt.stage = c.Stage
		opt.Post_EX_Process_Name = c.Post_EX_Process_Name
		opt.Host = c.Host
		opt.custom_cert = c.Keystore
		opt.cert_password = c.Password
		opt.metadata = c.Metadata
		opt.outFile = c.Outfile
		opt.beacon_PE = c.PE_Clone
		opt.Profile = c.Profile
		opt.processinject_min_alloc = c.Allocation
		opt.jitter = c.Jitter
		opt.sleeptime = c.Sleep
		opt.uri = c.Uri
		opt.customuri = c.Customuri
		opt.customuriGET = c.CustomuriGET
		opt.customuriPOST = c.CustomuriPOST
		opt.CDN = c.CDN
		opt.useragent = c.Useragent
		opt.ProfilePath = c.ProfilePath
		opt.injector = c.Injector
		opt.Datajitter = c.Datajitter
		opt.Keylogger = c.Keylogger
		opt.Forwarder = c.Forwarder
		opt.tasks_max_size = c.TasksMaxSize
		opt.tasks_proxy_max_size = c.TasksProxyMaxSize
		opt.tasks_dns_proxy_max_size = c.TasksDnsProxyMaxSize
		opt.syscall_method = c.Syscall_method
		opt.httplib = c.Httplib
		opt.threadspoof = c.Threadspoof
		opt.beacongate = c.BeaconGate
		opt.eaf_bypass = c.EafBypass
		opt.rdll_use_syscalls = c.RdllUseSyscalls
		opt.copy_pe_header = c.Copy_PE_Header
		opt.rdll_loader = c.RdllLoader
		opt.transform_obfuscate = c.TransformObfuscate
		opt.smartinject = c.SmartInject
	}

	if opt.outFile == "" {
		log.Fatal("Error: Please provide a file name to save the profile into")
	}
	if opt.Host == "" {
		log.Fatal("Error: Please provide the hostname or IP")
	}
	if opt.customuri != "" && (opt.customuriGET != "" || opt.customuriPOST != "") {
		log.Fatal("Error: Using Customuri with either of CustomuriGET or CustomuriPOST is not supported")
	}
	if (opt.customuriGET != "" && opt.customuriPOST == "") || (opt.customuriGET == "" && opt.customuriPOST != "") {
		log.Fatal("Error: When using CustomuriGET/CustomuriPOST, both must be sepecified")
	}
	fmt.Println(c.TasksMaxSize)
	Loader.GenerateOptions(opt.stage, opt.sleeptime, opt.jitter, opt.useragent, opt.uri, opt.customuri, opt.customuriGET, opt.customuriPOST, opt.beacon_PE, opt.processinject_min_alloc, opt.Post_EX_Process_Name, opt.metadata, opt.injector, opt.Host, opt.Profile, opt.ProfilePath, opt.outFile, opt.custom_cert, opt.cert_password, opt.CDN, opt.CDN_Value, opt.Datajitter, opt.Keylogger, opt.Forwarder, opt.tasks_max_size, opt.tasks_proxy_max_size, opt.tasks_dns_proxy_max_size, opt.syscall_method, opt.httplib, opt.threadspoof, opt.beacongate, opt.eaf_bypass, opt.rdll_use_syscalls, opt.copy_pe_header, opt.rdll_loader, opt.transform_obfuscate, opt.smartinject, opt.sleep_mask)
}
