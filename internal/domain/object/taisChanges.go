package object

import (
	"fmt"
	"strings"
)

type TaisChange struct {
	Date            string `json:"date" binding:"required"`
	Aviacompany     string `json:"aviacompany" binding:"required"`
	FltNum          string `json:"fltNum" binding:"required"`
	OriginIATA      string `json:"originIATA" binding:"required"`
	DestinationIATA string `json:"destinationIATA" binding:"required"`
	RBD             string `json:"RBD" binding:"required"`
	Capacity        int    `json:"capacity" binding:"required"`
}

func (tC *TaisChange) String() string {
	if tC.Capacity <= 99 {
		return fmt.Sprintf("%s %s %s %s%s F %s 00%d 00%d",
			tC.Date,
			tC.Aviacompany,
			tC.FltNum,
			tC.OriginIATA,
			tC.DestinationIATA,
			tC.RBD,
			tC.Capacity,
			tC.Capacity,
		)
	}
	return fmt.Sprintf("%s %s %s %s%s F %s 0%d 0%d",
		tC.Date,
		tC.Aviacompany,
		tC.FltNum,
		tC.OriginIATA,
		tC.DestinationIATA,
		tC.RBD,
		tC.Capacity,
		tC.Capacity,
	)
}

type TaisChanges []TaisChange

func (tCs TaisChanges) String() string {
	res := ""
	for _, v := range tCs {
		res += v.String() + "\n"
	}
	return strings.TrimSpace(res)
}
