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
	maxdns					 string
	Yaml                     string
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
	TasksMaxSize         string `yaml:"TasksMaxSize"`
	TasksProxyMaxSize    string `yaml:"TasksProxyMaxSize"`
	TasksDnsProxyMaxSize string `yaml:"TasksDnsProxyMaxSize"`
	Maxdns				 string `yaml:"Maxdns"`
	Forwarder            bool   `yaml:"Forwarder"`
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
[1] srv.dll
[2] ActivationManager.dll
[3] audioeng.dll
[4] AzureSettingSyncProvider.dll
[5] BingMaps.dll
[6] BootMenuUX.dll
[7] DIAGCPL.dll
[8] FIREWALLCONTROLPANEL.dll
[9] WMNetMgr.dll
[10] wwanapi.dll
[11] Windows.Storage.Search.dll
[12] Windows.System.Diagnostics.dll
[13] Windows.System.Launcher.dll
[14] Windows.System.SystemManagement.dll
[15] Windows.UI.BioFeedback.dll
[16] Windows.UI.BlockedShutdown.dll
[17] Windows.UI.Core.TextInput.dll
[18] FILEMGMT.dll
[19] polprocl.dll
[20] GPSVC.dll
[21] libcrypto.dll
[22] rdpcomapi.dll
[23] winsqlite3.dll
[24] wow64.dll
[25] wow64win.dll
[26] WWANSVC.dll
[27] CyMemDef64.dll (Cylance's DLL)
[28] InProcessClient.dll (SentinelOne's DLL)
[29] ctiuser.dll (Carbon Black's DLL)
[30] umppc.dll (CrowdStrike's DLL)`)
	processinject_min_alloc := flag.String("Allocation", "", "Minimum amount of memory to request for injected content (must be higher than 4096)")
	Post_EX_Process_Name := flag.String("PostEX_Name", "", `File Post-Ex activities will spawn and inject into (Use the number):
[1] WerFault.exe
[2] WWAHost.exe
[3] wlanext.exe
[4] auditpol.exe
[5] bootcfg.exe
[6] choice.exe
[7] bootcfg.exe
[8] dtdump.exe
[9] expand.exe
[10] fsutil.exe
[11] gpupdate.exe
[12] gpresult.exe
[13] logman.exe
[14] mcbuilder.exe
[15] mtstocom.exe
[16] pcaui.exe
[17] powercfg.exe
[18] svchost.exe`)
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
	tasks_max_size := flag.String("TasksMaxSize", "", "The maximum size (in bytes) of task(s) and proxy data that can be transferred through a communication channel at a check in")
	tasks_proxy_max_size := flag.String("TasksProxyMaxSize", "", "The maximum size (in bytes) of proxy data to transfer via the communication channel at a check in")
	tasks_dns_proxy_max_size := flag.String("TasksDnsProxyMaxSize", "", "The maximum size (in bytes) of proxy data to transfer via the DNS communication channel at a check in")
	maxdns := flag.String("Maxdns", "", "Maximum length of hostname when uploading data over DNS (0-255) (default 200)")
	Forwarder := flag.Bool("Forwarder", false, "Enabled the X-forwarded-For header (Good for when your C2 is behind a redirector)")
	Yaml := flag.String("Yaml", "", "Path to the Yaml config file")
	flag.Parse()
	return &FlagOptions{stage: *stage, sleeptime: *sleeptime, jitter: *jitter, useragent: *useragent, uri: *uri, customuri: *customuri, customuriGET: *customuriGET, customuriPOST: *customuriPOST, beacon_PE: *beacon_PE, processinject_min_alloc: *processinject_min_alloc, Post_EX_Process_Name: *Post_EX_Process_Name, metadata: *metadata, injector: *injector, Host: *Host, Profile: *Profile, ProfilePath: *ProfilePath, outFile: *outFile, custom_cert: *custom_cert, cert_password: *cert_password, CDN: *CDN, CDN_Value: *CDN_Value, Yaml: *Yaml, Datajitter: *Datajitter, Keylogger: *Keylogger, Forwarder: *Forwarder, tasks_max_size: *tasks_max_size, tasks_proxy_max_size: *tasks_proxy_max_size, tasks_dns_proxy_max_size: *tasks_dns_proxy_max_size, maxdns: *maxdns}
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
		opt.maxdns = c.Maxdns
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
	Loader.GenerateOptions(opt.stage, opt.sleeptime, opt.jitter, opt.useragent, opt.uri, opt.customuri, opt.customuriGET, opt.customuriPOST, opt.beacon_PE, opt.processinject_min_alloc, opt.Post_EX_Process_Name, opt.metadata, opt.injector, opt.Host, opt.Profile, opt.ProfilePath, opt.outFile, opt.custom_cert, opt.cert_password, opt.CDN, opt.CDN_Value, opt.Datajitter, opt.Keylogger, opt.Forwarder, opt.tasks_max_size, opt.tasks_proxy_max_size, opt.tasks_dns_proxy_max_size, opt.maxdns)
}
