package types

type SearchAjaxData struct {
	Error               string     `json:"error"`
	TotalRecords        int        `json:"iTotalRecords"`
	TotalDisplayRecords int        `json:"iTotalDisplayRecords"`
	Echo                int        `json:"sEcho"`
	Data                [][]string `json:"aaData"`
}

type RecordType int

const (
	FullLength RecordType = 1 << iota
	Demo
	EP
	Compilation
	Live
	BoxedSet
	Single
	Video
	Split
	Other
)

func SelectRecordType(record string) RecordType {
	var typeFound RecordType

	switch record {
	case "Full-length":
		typeFound = FullLength
	case "EP":
		typeFound = EP
	case "Compilation":
		typeFound = Compilation
	case "Demo":
		typeFound = Demo
	case "Video":
		typeFound = Video
	case "Single":
		typeFound = Single
	case "Live album":
		typeFound = Live
	case "Split":
		typeFound = Split
	case "Boxed set":
		typeFound = BoxedSet
	default:
		typeFound = Other
	}

	return typeFound
}
