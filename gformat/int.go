package gformat

import "database/sql"

func GetInt64Null(s int64) sql.NullInt64 {
	r := sql.NullInt64{
		Int64: 0,
		Valid:  false,
	}
	if s > 0 {
		r.Int64 = s
		r.Valid = true
	}
	return r
}

func GetNullInt64(s sql.NullInt64) int64 {
	if !s.Valid {
		return 0
	}
	return s.Int64
}

func GetInt32Null(s int32) sql.NullInt32 {
	r := sql.NullInt32{
		Int32: 0,
		Valid:  false,
	}
	if s > 0 {
		r.Int32 = s
		r.Valid = true
	}
	return r
}

func GetNullInt32(s sql.NullInt32) int32 {
	if !s.Valid {
		return 0
	}
	return s.Int32
}


func GetInt16Null(s int16) sql.NullInt16 {
	r := sql.NullInt16{
		Int16: 0,
		Valid:  false,
	}
	if s > 0 {
		r.Int16 = s
		r.Valid = true
	}
	return r
}

func GetNullInt16(s sql.NullInt16) int16 {
	if !s.Valid {
		return 0
	}
	return s.Int16
}