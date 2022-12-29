package safeObject

import "github.com/achillescres/saina-api/internal/domain/object"

type TaisChangesSafe object.TaisChanges

func ToTaisChangesSafe(tC object.TaisChanges) *TaisChangesSafe {
	tCS := TaisChangesSafe(tC)
	return &tCS
}

func (tCS *TaisChangesSafe) ToEntity() *object.TaisChanges {
	tS := object.TaisChanges(*tCS)
	return &tS
}
