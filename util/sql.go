package util

import "database/sql"

func ValidNullString(s string) sql.NullString {
	return sql.NullString{
		Valid:  true,
		String: s,
	}
}
