package secure

import (
	"fmt"
)

var (
	CurrentChromiumVersions = []string{"134", "135", "136", "137", "138", "139", "140"}

	AppleSiliconCPUs = []string{
		"Apple M1",
		"Apple M1 Pro",
		"Apple M1 Max",
		"Apple M1 Ultra",

		"Apple M2",
		"Apple M2 Pro",
		"Apple M2 Max",
		"Apple M2 Ultra",

		"Apple M3",
		"Apple M3 Pro",
		"Apple M3 Max",
		"Apple M3 Ultra",

		"Apple M4",
		"Apple M4 Pro",
		"Apple M4 Max",
	}

	OlderMacOSVersions = []string{
		"10_14",
		"10_14_1",
		"10_14_2",
		"10_14_3",
		"10_14_4",
		"10_14_5",
		"10_14_6",

		"10_15",
		"10_15_1",
		"10_15_2",
		"10_15_3",
		"10_15_4",
		"10_15_5",
		"10_15_6",
		"10_15_7",
	}

	NewerMacOSVersions = []string{
		"11_0",
		"11_0_1",
		"11_1",
		"11_2",
		"11_2_1",
		"11_2_2",
		"11_2_3",
		"11_3",
		"11_3_1",
		"11_4",
		"11_5",
		"11_5_1",
		"11_5_2",
		"11_6",
		"11_6_1",
		"11_6_2",
		"11_6_3",
		"11_6_4",
		"11_6_5",
		"11_6_6",
		"11_6_7",
		"11_6_8",
		"11_7",
		"11_7_1",
		"11_7_2",
		"11_7_3",
		"11_7_4",
		"11_7_5",
		"11_7_6",
		"11_7_7",
		"11_7_8",
		"11_7_9",
		"11_7_10",

		"12_0",
		"12_0_1",
		"12_1",
		"12_2",
		"12_2_1",
		"12_3",
		"12_3_1",
		"12_4",
		"12_5",
		"12_5_1",
		"12_6",
		"12_6_1",
		"12_6_2",
		"12_6_3",
		"12_6_4",
		"12_6_5",
		"12_6_6",
		"12_6_7",
		"12_6_8",
		"12_6_9",
		"12_7",
		"12_7_1",
		"12_7_2",
		"12_7_3",
		"12_7_4",
		"12_7_5",
		"12_7_6",

		"13_0",
		"13_0_1",
		"13_1",
		"13_2",
		"13_2_1",
		"13_3",
		"13_3_1",
		"13_4",
		"13_4_1",
		"13_5",
		"13_5_1",
		"13_5_2",
		"13_6",
		"13_6_1",
		"13_6_2",
		"13_6_3",
		"13_6_4",
		"13_6_5",
		"13_6_6",
		"13_6_7",
		"13_6_8",
		"13_6_9",
		"13_7",
		"13_7_1",
		"13_7_2",
		"13_7_3",
		"13_7_4",
		"13_7_5",
		"13_7_6",
		"13_7_7",
		"13_7_8",

		"14_0",
		"14_1",
		"14_1_1",
		"14_1_2",
		"14_2",
		"14_2_1",
		"14_3",
		"14_3_1",
		"14_4",
		"14_4_1",
		"14_5",
		"14_6",
		"14_6_1",
		"14_7",
		"14_7_1",
		"14_7_2",
		"14_7_3",
		"14_7_4",
		"14_7_5",
		"14_7_6",
		"14_7_7",
		"14_7_8",
		"14_8",
		"14_8_1",

		"15_0",
		"15_0_1",
		"15_1",
		"15_1_1",
		"15_2",
		"15_3",
		"15_3_1",
		"15_3_2",
		"15_4",
		"15_4_1",
		"15_5",
		"15_6",
		"15_6_1",
		"15_7",
		"15_7_1",

		"26_0",
		"26_0_1",
	}
)

func NewUserAgent() string {
	var u string

	if Number(0, 1) == 0 { // Apple Silicon
		cpu := AppleSiliconCPUs[Number(0, len(AppleSiliconCPUs)-1)]
		macOs := NewerMacOSVersions[Number(0, len(NewerMacOSVersions)-1)]
		chrome := CurrentChromiumVersions[Number(0, len(CurrentChromiumVersions)-1)]

		u = fmt.Sprintf("Mozilla/5.0 (Macintosh; %s; macOS %s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s.0.0.0 Safari/537.36", cpu, macOs, chrome)
	} else { // Intel
		macOs := OlderMacOSVersions[Number(0, len(OlderMacOSVersions)-1)]
		chrome := CurrentChromiumVersions[Number(0, len(CurrentChromiumVersions)-1)]

		u = fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X %s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s.0.0.0 Safari/537.36", macOs, chrome)
	}

	return u
}

func NewHeaders() map[string]string {
	return map[string]string{
		"User-Agent":      NewUserAgent(),
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"Accept-Language": "en-US;q=1.0",
	}
}
