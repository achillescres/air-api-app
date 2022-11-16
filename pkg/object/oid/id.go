package oid

type Id string

func ToId(id string) Id {
	return Id(id)
}
