package util

import "database/sql"

// NewNullString insert나 update시 string필드들이 ""로 되어있기때문에 null값넣고싶으면 sql.Nullstring으로 만들어서 insert
func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
