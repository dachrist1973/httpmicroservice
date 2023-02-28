package helper

import (
	"fmt"
	"interviewMSrvHTTP/srverrors"
	"net/url"
	"strconv"
)

func BuildWhereClause(values url.Values) (string, int, int, error) {
	var whereClause string
	var first bool = true
	limit := 0
	page := 1

	for k, v := range values {
		if k == "" || v[0] == "" {
			return "", 0, 0, fmt.Errorf(srverrors.BadQueryParam)
		}
		if k == "limit" {
			limit, _ = strconv.Atoi(v[0])
		} else if k == "page" {
			page, _ = strconv.Atoi(v[0])
		} else if !first {
			whereClause = whereClause + fmt.Sprintf(" AND %s = %s", k, v[0])
		} else {
			whereClause = fmt.Sprintf("%s = %s", k, v[0])
			first = false
		}
	}

	return whereClause, limit, page, nil
}
