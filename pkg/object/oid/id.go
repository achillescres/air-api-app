package oid

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type Id string

func ToId(id string) (Id, error) {
	parse, err := uuid.Parse(strings.TrimSpace(id))
	if err != nil {
		return "", errors.New(fmt.Sprintf("error parsing uuid froms string: %s", err.Error()))
	}

	return Id(parse.String()), nil
}

func NewId() Id {
	return Id(uuid.New().String())
}
