package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	res := ""
	rows := make([]string, numRows)
	r := 0
	nextRow := -1

	for _, v := range s {
		if r == 0 || r == numRows-1 {
			nextRow *= -1
		}

		rows[r] += string(v)

		r += nextRow
	}

	for _, v := range rows {
		res += v
	}

	return res
}
