package types

import (
	commontypes "github.com/a-castellano/music-manager-common-types/types"
)

type SearchAjaxData struct {
	Error               string     `json:"error"`
	TotalRecords        int        `json:"iTotalRecords"`
	TotalDisplayRecords int        `json:"iTotalDisplayRecords"`
	Echo                int        `json:"sEcho"`
	Data                [][]string `json:"aaData"`
}

func SelectRecordType(record string) commontypes.RecordType {
	var typeFound commontypes.RecordType

	switch record {
	case "Full-length":
		typeFound = commontypes.FullLength
	case "EP":
		typeFound = commontypes.EP
	case "Compilation":
		typeFound = commontypes.Compilation
	case "Demo":
		typeFound = commontypes.Demo
	case "Video":
		typeFound = commontypes.Video
	case "Single":
		typeFound = commontypes.Single
	case "Live album":
		typeFound = commontypes.Live
	case "Split":
		typeFound = commontypes.Split
	case "Boxed set":
		typeFound = commontypes.BoxedSet
	default:
		typeFound = commontypes.Other
	}

	return typeFound
}
