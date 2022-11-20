package oid

type Id string

const Undefined Id = "undefined"

func ToId(id string) Id {
	return Id(id)
}

func IsUndefined(id Id) bool {
	return id == Undefined
}
