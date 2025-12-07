package gformat

import "database/sql"

func GetFloat64Null(s float64) sql.NullFloat64 {
	r := sql.NullFloat64{
		Float64: 0,
		Valid:   false,
	}
	if s > 0 {
		r.Float64 = s
		r.Valid = true
	}
	return r
}

func GetNullFloat64(s sql.NullFloat64) float64 {
	if !s.Valid {
		return 0
	}
	return s.Float64
}
