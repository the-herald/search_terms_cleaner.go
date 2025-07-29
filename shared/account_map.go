package shared

import (
	"strings"
)

// AccountMap maps readable aliases to actual account IDs
var AccountMap = map[string]string{
	"satilla":                                     "5616230554",
	"sfs":                                         "5616230554",
	"satilla family smiles":                      "5616230554",
	"fnbmd":                                       "3035218698",
	"first national":                              "3035218698",
	"first national bank of mount dora":          "3035218698",
	"gabaie":                                      "6666797635",
	"gabaie & associates":                         "6666797635",
	"godley":                                      "6655601976",
	"godley station":                              "6655601976",
	"godley station dental":                       "6655601976",
	"gsd":                                         "6655601976",
	"three ten":                                   "5692134970",
	"three ten timber":                            "5692134970",
	"four seasons":                                "1335938339",
	"four seasons residences":                     "1335938339",
	"fsrlv":                                       "1335938339",
	"four seasons las vegas":                      "1335938339",
	"first doctors":                               "1462306408",
	"fdwl":                                        "1462306408",
	"first doctors weight loss":                  "1462306408",
	"ballparkdj":                                  "5287833435",
	"andrew casey":                                "6807963143",
	"acec":                                        "6807963143",
	"andrew casey electrical":                    "6807963143",
	"precise":                                     "5392828629",
	"precise home":                                "5392828629",
	"dynamic":                                     "6309687513",
	"dynamic warehouse":                           "6309687513",
	"rosco":                                       "3962597664",
	"rosco generators":                            "3962597664",
	"h2h":                                         "2206893203",
	"heart to home":                               "2206893203",
	"heart to home meals":                         "2206893203",
	"hthm":                                        "2206893203",
	"uc":                                          "3020635710",
	"uc components":                               "3020635710",
	"capio":                                       "3030064078",
	"capiorn":                                     "3030064078",
	"iowa":                                        "9552845701",
	"iowa countertops":                            "9552845701",
	"action":                                      "4224597425",
	"action engraving":                            "4224597425",
	"scribble":                                    "6555309398",
	"scribblevet":                                 "6555309398",
	"woodlands":                                   "3466831668",
	"woodlands family dental":                     "3466831668",
	"sound concrete":                              "3168167882",
	"sound concrete solutions":                    "3168167882",
	"bpd":                                         "9162462492",
	"bullet proof":                                "9162462492",
	"bullet proof diesel":                         "9162462492",
	"tristate":                                    "3229754921",
	"tristate siding":                             "3229754921",
	"sign":                                        "8343552815",
	"sign systems":                                "8343552815",
}

// ResolveAccountID tries to match user input to a known account ID
func ResolveAccountID(input string) (string, bool) {
	normalized := Normalize(input)
	for key, val := range AccountMap {
		if strings.Contains(normalized, key) {
			return val, true
		}
	}
	return "", false
}
