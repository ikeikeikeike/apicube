package repo

import (
	"errors"

	"github.com/spf13/cast"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

const (
	dbFormat = "2006-01-02 15:04:05"
)

var (
	// AscOrder defines ORDER BY query as ascending
	AscOrder = qm.OrderBy("id ASC")
	// DescOrder defines ORDER BY query as descending
	DescOrder = qm.OrderBy("id DESC")
	// ProductPublic defines WHERE query
	ProductPublic = qm.Where("product_status_id = ?", 1)
	// ProductPrivate defines WHERE query
	ProductPrivate = qm.Where("product_status_id = ?", 2)
)

// XXX: this is so fuzzzy
func fuzzy(v interface{}, arr []interface{}) bool {
	for _, i := range arr {
		if cast.ToString(i) == cast.ToString(v) {
			return true
		}
	}

	return false
}

func preloadBy(where []qm.QueryMod, loads ...string) ([]qm.QueryMod, error) {
	if len(where) <= 0 {
		return nil, errors.New("no queries")
	}

	mods := preloads(loads...)
	for _, w := range where {
		mods = append(mods, w)
	}

	return mods, nil
}

func preload(id uint, loads ...string) []qm.QueryMod {
	mods := []qm.QueryMod{qm.Where("id = ?", id)}
	for _, load := range loads {
		mods = append(mods, qm.Load(load))
	}

	return mods
}

func preloads(loads ...string) []qm.QueryMod {
	mods := []qm.QueryMod{qm.OrderBy("id ASC")}
	for _, load := range loads {
		mods = append(mods, qm.Load(load))
	}

	return mods
}
