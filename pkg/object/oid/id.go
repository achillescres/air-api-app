package oid

import "strings"

type Id string

const Undefined Id = "oid_undefined"

func ToId(id string) Id {
	return Id(id)
}

func IsUndefined(id Id) bool {
	return id == Undefined ||
		len(strings.TrimSpace(string(id))) == 0
}
