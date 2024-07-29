package data

type MediaType string

const (
	JSON MediaType = "application/json"
	XML  MediaType = "application/xml"
	CSV  MediaType = "text/csv"
	YAML MediaType = "text/yaml"
	PDF  MediaType = "application/pdf"
	XLS  MediaType = "application/vnd.ms-excel"
	XLSX MediaType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
)

type RoleType string

const (
	INVESTIGATION RoleType = "investigation"
	CASE          RoleType = "case"
	FACT          RoleType = "fact"
	LEAD          RoleType = "lead"
)
