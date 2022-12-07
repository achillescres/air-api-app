package oid

import "strings"

type Id string

const Undefined Id = "oid_undefined"

func ToId(id string) Id {
	return Id(id)
}

func AssertId(id any) (Id, bool) {
	nId, ok := id.(Id)
	if !ok {
		nId = Undefined
	}

	return nId, ok
}

func IsUndefined(id Id) bool {
	return id == Undefined ||
		len(strings.TrimSpace(string(id))) == 0
}
