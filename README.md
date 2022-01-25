
<p align="center"> <img src=Screenshots/logo.png  border="2px solid #555">

# SourcePoint
SourcePoint is a polymorphic C2 profile generator for Cobalt Strike C2s, written in Go. SourcePoint allows unique C2 profiles to be generated on the fly that helps reduce our Indicators of Compromise ("IoCs") and allows the operator to spin up complex profiles with minimal effort. This was done by extensively reviewing [Articles](https://www.cobaltstrike.com/help-malleable-c2) as well as [Patch Notes](https://www.cobaltstrike.com/releasenotes.txt) to identify key functions and modifiable features. SourcePoint was designed with the intent of addressing the issue of how to make our C2 activity harder to detect, focusing on moving away from malicious IoCs to suspicious ones. The goal here is that it is harder to detect our C2 if our IoCs are not malicious in-nature and require additional research to discover the suspicious nature. SourcePoint contains numerous different configurable options to choose from to modify your profile (in most cases if left blank SourcePoint will randomly choose them for you). The generated profiles modify all aspects of your C2. The goal of this project is to not only aid in circumventing detection-based controls but also help blend C2 traffic and activity into the environment, making said activity hard to detect. 


<p align="center"> <img src=Screenshots/C2int_p1.png width="900" height="710" border="2px solid #555">
<p align="center"> <img src=Screenshots/C2int_p2.png width="900" height="480" border="2px solid #555">

```
go install github.com/Tylous/SourcePoint
```

## Installation
```
$go get gopkg.in/yaml.v2

$go build SourcePoint.go
```

## Usage

```
#./SourcePoint -h

	   _____                            ____        _       __
	  / ___/____  __  _______________  / __ \____  (_)___  / /_
	  \__ \/ __ \/ / / / ___/ ___/ _ \/ /_/ / __ \/ / __ \/ __/
	 ___/ / /_/ / /_/ / /  / /__/  __/ ____/ /_/ / / / / / /_
	/____/\____/\__,_/_/   \___/\___/_/    \____/_/_/ /_/\__/
  							(@Tyl0us)

                                                                                                                        
Usage of ./SourcePoint:
  -Allocation string
        Minimum amount of memory to request for injected content (must be higher than 4096)
  -CDN string
        CDN cookie name (typically used for AzureEdge profiles)
  -CDN-Value string
        CDN cookie value (typically used for AzureEdge profiles)
  -Customuri string
        The base URI for custom HTTP GET/POST profile (default "0")
  -Datajitter string
        Appends a value to HTTP-Get and HTTP-Post server output (default "50")
  -Forwarder
        Enabled the X-forwarded-For header (Good for when your C2 is behind a redirector)
  -Host string
        Team server domain name
  -Injector string
        Select the preferred method to allocate memory in the remote process:
        [*] VirtualAllocEx (Great for cross architecture i.e x86 -> x64 and x64->x86)
        [*] NtMapViewOfSection (A more stealthly option, however fails over to VirtualAllocEx, generating more events when it does)
  -Jitter string
        Jitter percentage for beacon call home
  -Keylogger string
        Select the preferred method the beacon will use to log keystrokes: 
        [*] GetAsyncKeyState (Uses GetAsyncKeyState API (Separate DLL for x86/x64 process))
        [*] SetWindowsHookEx (Uses SetWindowsHookEx API)
  -Keystore string
        SSL keystore name
  -Metadata string
        Specifies how to transform and embed metadata into the HTTP request:
        [*] base64
        [*] base64url
        [*] netbios
        [*] netbiosu (default "base64url")
  -Outfile string
        Name of output file
  -PE_Clone string
        PE file beacon will mimic (Use the number):
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
        [30] umppc.dll (CrowdStrike's DLL)
  -Password string
        SSL certificate password
  -PostEX_Name string
        File Post-Ex activities will spawn and inject into (Use the number):
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
        [18] svchost.exe
  -Profile string
        HTTP GET/POST profile (Use the number):
        [1] Windowsupdate
        [2] Slack
        [3] Gotomeeting
        [4] Outlook.Live
        [5] Safebrowsing [Cloudfront Compatible]
        [6] AzureEdge [AzureEdge Compatible]
        [7] Field-Keyword [Cloudfront Compatible]
        [8] Custom (Used with ProfilePath)
  -ProfilePath string
        Path of custom HTTP GET/POST profile...
  -Sleep string
        Initial beacon sleep time
  -Stage string
        Disable host staging (Default: False) (default "False")
  -Uri string
        The number URIs a profile for beacons to choose from
  -Useragent string
        UserAgent string for the beacon to use (Leave blank to randomly select one):
        [*] Win10Chrome
        [*] Win10Edge
        [*] Win10IE
        [*] Win10
        [*] Win6.3
        [*] Linux
        [*] Mac
  -Yaml string
        Path to the Yaml config file
```


## Important

SourcePoint primarily automates the build process of a profile. It’s very important to know all the features modified in these profiles. Knowing these features can really help increase your success. 

## Options 

While there are a lot of settings and features described in the help function of SourcePoint, there are numerous important features baked into each profile that are important to be familiar with. These features are:

### Global Options
This part of your profile modifies how the beacon operators. Some of the features used to modify the behaviour are:

* Host Stage - Allows the team server to host staged shellcode for HTTP, HTTPS, DNS. If this is enabled, anyone sending a GET request with a specific value such has `/9ZXq` can pull the shellcode as well
* Sleep - The length of time that a beacon calls back home
* Jitter - Appends a percentage to the beacon call home time
* Useragent - The useragent string used when communicating HTTP and HTTPS traffic. Using the appropriate useragent string can help blend into the environment 
* Data Jitter - Adds a random-length string to all GET and POST requests to ensure incoming requests are not the same length
* SMB Frame Header - Adds a header value to the SMB beacon messages
* Pipename - Sets the name of the SMB pipe the beacons is going to use for communication
* Pipename Stager - Sets the name of the SMB stager for the beacons
* TCP Frame Header - Adds a header value to the TCP beacon messages
* SSH Banner - The SSH banner used
* SSH Pipename - The name used for the SSH banner

### Stage
This part of your profile controls how beacon is loaded into memory and edit the content of the beacon DLL. Some of the features used to modify the behaviour are:

* Obfuscate - Obfuscates the import table of the reflective DLL
* Stomppe - Asks the payload to stomp MZ, PE and, e_lfanen values after loading
* Clean up - Tells the beacon to free up memory assoicated with the refelctive DLL that initalized it
* UseRWX - Ensures shellcode does not use Read, Write Execute permissions
* Smart Inject - Uses embedded function pointer hints to bootstrap the beacon agent without walking kernel32 EAT
* Sleep Mask - TCP and SMB beacons will obfuscate themselves at rest while they wait for the connection to be established
* PE Header - Changes the characteristics of your beacon Reflective DLL to look like something else in memory
* Transformation - Transform beacon's Reflective DLL stage by removing or adding strings to the .rdata section

### Process-Inject
This part of your profile controls how the beacon shapes injected content and controls process injection behavior. Some of the features used to modify the behaviour are:

* Allocator - Determines how the beacon's Reflective loader allocates memory 
* Minimum Allocation - Minimum amount of memory to request for injected content
* Userwx - Ensures shellcode does not use Read, Write Execute permissions (The alternative is RW)
* Startrwx - Use Read, Write Execute as initial permissions for injected content (The alternative is RW)
* Transformer - Adds a block of padding content injected by the beacon
* Execute - This section determines how to execute the injected code

### Post-Exec
This part of your profile controls how the beacon handles post-exploitation modules and commands. Some of the features used to modify the behaviour are:

* Spawnto - Determines the default temporary process beacon will spawn for its post-exploitation command and options
* Obfuscate - Obfuscates the import table of the reflective DLL
* Smart Inject - Pass key function pointers from beacon to its child jobs
* AMSI disable - Disable AMSI for powerpick, execute-assembly, and psinject (Certain EDRs can detect this best avoid using these tools)
* Keylogger - Determines how the keystroker logging API use to capture keystrokes


### Profiles
Currently SourcePoint provides you with 6 baked in options for HTTP/HTTPS traffic profiles, based on existing profiles. Of these 6, 4 of them are influenced by and based on:
* Microsoft Window's Update Communication
* Slack's Message Communication
* Gotomeeting's Active Meeting Communication 
* Microsoft Outlook's Email Communication

3 of the profile options (5, 6 and 7) are designed specifically for:
* Cloudfront.net
* AzureEdge.net

The last option (8) is designed to input a custom profile. This option is designed to allow an operator to utilize a completely custom traffic profile. There are many cases where a completely unique traffic profile will yield high success rather than one of these. This also allows operators to still utilize SourcePoint's malleability features with their go-to or favorite traffic profile. As this allows for unique profiles it’s important to ensure you tweak and adjust the profile for SourcePoint to work. At a minimum:
* Replace - `header "Host" "acme.com";` with `header "Host" "{{.Variables.Host}}";`
* Replace - `/pathtolegitpage/` under the GET field with `{{.Variables.HTTP_GET_URI}}`
* Replace - `/pathtolegitpage/` under the POST field with `{{.Variables.HTTP_POST_URI}}`


To do so, use the following options `-CustomURI` and `-ProfilePath` along with `-Profile 8`. While developing a profile, it’s highly recommended to use the native ./c2lint to verify everything is working. 


## Sample Yaml Configs

```
Stage: "False"
Host: "acme-email.com"
Keystore: "acme-email.com.store"
Password: "Password"
Metadata: "netbios"
Injector: "VirtualAllocEx"
Outfile: "acme.profile"
PE_Clone: 20
Profile: 4
Allocation: 5312
Jitter: 30
Debug: true
Sleep: 35
Uri: 3
Useragent:  "Mac"
Post-EX Processname: 11
Datajitter: 40
Keylogger: "SetWindowsHookEx"
Customuri: 
CDN:
CDN_Value: 
ProfilePath: 
Forwarder: False
```



## SSL Certificate

Profiles mode 1-4 can be used without a validate SSL, SourcePoint will generate a self-signed certificate related to the profile type. However, valid SSL certificates are extremely important the success of any type of C2. For many reasons but obviously no certificate means the traffic is going to be unencrypted (i.e. HTTP WHICH SHOULD NEVER BE USED) but using a self-signed cert comes with its obvious limitations. There are many ways to obtain a valid SSL certificate to make a keystore my go to way is using a modified version of [HTTPsC2DoneRight.sh](https://github.com/killswitch-GUI/CobaltStrike-ToolKit/blob/master/HTTPsC2DoneRight.sh), created by [Cham423](https://github.com/cham423/cs-tools). 


## DNS

Currently DNS customization not offered directly through SourcePoint. To still allow dns-based beacons there is a commented out section for dns-beacon in every generated profile.


## To Do List
- [ ] Add More Profiles
- [ ] DNS Staging
