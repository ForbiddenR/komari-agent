package vnstat

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseVnstat(t *testing.T) {
	str := `{"vnstatversion":"2.10","jsonversion":"2","interfaces":[{"name":"eth0","alias":"","created":{"date":{"year":2025,"month":12,"day":10},"timestamp":1765333863},"updated":{"date":{"year":2026,"month":1,"day":16},"time":{"hour":9,"minute":40},"timestamp":1768527600},"traffic":{"total":{"rx":69751508516,"tx":56277332965},"month":[{"id":1,"date":{"year":2025,"month":12},"timestamp":1764518400,"rx":41142776983,"tx":35021551997},{"id":2,"date":{"year":2026,"month":1},"timestamp":1767196800,"rx":28608731533,"tx":21255780968}]}}]}`
	vs := &VnstatJson{}
	err := json.Unmarshal([]byte(str), vs)
	if err != nil {
		panic(err)
	}
	fmt.Println(vs)
}
