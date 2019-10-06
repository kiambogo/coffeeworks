package models

import "sort"

type Badge struct {
	Name     string
	Selected bool
}

var Badges = map[string]Badge{
	"beverages": beverages,
	"wifi":      wifi,
	"service":   service,
	"seating":   seating,
	"outlets":   outlets,
	"food":      food,
}

var BadgeNames []string

func init() {
	for name := range Badges {
		BadgeNames = append(BadgeNames, name)
	}
	sort.Slice(BadgeNames, func(i, j int) bool { return BadgeNames[i] < BadgeNames[j] })

}

var beverages Badge = Badge{Name: "beverages"}
var wifi Badge = Badge{Name: "wifi"}
var service Badge = Badge{Name: "service"}
var seating Badge = Badge{Name: "seating"}
var outlets Badge = Badge{Name: "outlets"}
var food Badge = Badge{Name: "food"}
