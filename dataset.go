package sqldata

import (
	"database/sql"
)

// DataSet ...
type DataSet = []map[string]string

// GetDataSet ...
unc GetDataSet(rows *sql.Rows) DataSet {
	cols, _ := rows.Columns()
	colsize := len(cols)
	dataset := DataSet{}

	for rows.Next() {
		colmap := make(map[string]string)
		// colmap := make(map[string]string)
		coldata := make([]interface{}, colsize)

		for i := 0; i < colsize; i++ {
			coldata[i] = new(interface{})
		}
		rows.Scan(coldata...)

		for i, m := range cols {
			v := coldata[i].(*interface{})

			switch (*v).(type) {
			case nil:
				colmap[m] = ""
			case int64:
				colmap[m] = fmt.Sprintf("%v", *v)
			default:
				colmap[m] = fmt.Sprintf("%s", *v)
			}
		}
		dataset = append(dataset, colmap)
	}

	return dataset
}

