package common

import (
	"strings"

	"github.com/hailaz/gadmin/app/model"
)

func GetAction(act string) string {
	acts := strings.Split(strings.Split(strings.Split(act, ";")[0], ":")[0], ",")
	action := ""
	for _, v := range acts {
		if v == "All" || v == "REST" {
			return model.ACTION_ALL
		}
		if action == "" {
			action += strings.ToUpper("(" + v + ")")
		} else {
			action += strings.ToUpper("|(" + v + ")")
		}

	}
	return action
}
