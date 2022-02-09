package app

import (
	"github.com/gocraft/dbr"
	"strings"
)

type Pagination struct {
	Limit     uint64 `query:"limit"`
	Offset    uint64 `query:"offset"`
	SortKey   string `query:"sortKey"`
	SortOrder string `query:"sortOrder"`
}

func (p Pagination) Asc() bool {
	return strings.ToLower(p.SortOrder) == "asc"
}

func (p Pagination) WithLimit(smtm *dbr.SelectStmt) {
	if p.Limit != 0 {
		smtm.Limit(p.Limit)
	}
}
