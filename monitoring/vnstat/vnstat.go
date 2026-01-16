package vnstat

import (
	"encoding/json"
	"os/exec"
	"time"
)

var (
	networkIn      uint64 = 0
	networkOut     uint64 = 0
	lastUpdateTime        = time.Unix(0, 0)
)

type VnstatJson struct {
	VnstatVersion string  `json:"vnstatversion"`
	JsonVersion   string  `json:"jsonversion"`
	Interfaces    []Iface `json:"interfaces"`
}

type Iface struct {
	// v1
	Id   string `json:"id"`
	Nick string `json:"nick"`
	// v2
	Name  string `json:"name"`
	Alias string `json:"alias"`
	// common
	Traffic Traffic `json:"traffic"`
}

type Traffic struct {
	Total RT `json:"total"`
	// // v1 months
	// Months []DateRT `json:"months"`
	// // va day
	// Days []DateRT `json:"days"`
	// // v2 month
	// Month []DateRT `json:"month"`
	// // v2 day
	// Day []DateRT `json:"day"`
}

// type Date struct {
// 	Year  int32  `json:"year"`
// 	Month uint32 `json:"month"`
// 	Day   uint32 `json:"day"`
// }

type RT struct {
	Rx uint64 `json:"rx"`
	Tx uint64 `json:"tx"`
}

// type DateRT struct {
// 	Id   uint64 `json:"id"`
// 	Date Date   `json:"date"`
// 	RT
// }

func GetTotalTraffic() (uint64, uint64, error) {
	if time.Since(lastUpdateTime) < 5*time.Minute {
		return networkIn, networkOut, nil
	}
	defer func() {
		lastUpdateTime = time.Now()
	}()
	o, err := exec.Command("/usr/bin/vnstat", "--json", "m").Output()
	if err != nil {
		return 0, 0, err
	}
	var result VnstatJson
	err = json.Unmarshal(o, &result)
	if err != nil {
		return 0, 0, err
	}
	for _, iface := range result.Interfaces {
		networkIn += iface.Traffic.Total.Rx
		networkOut += iface.Traffic.Total.Tx
	}
	return networkIn, networkOut, nil
}
