package sqldata

import (
	"database/sql"
)

// DataSet ...
type DataSet = []map[string]string

// GetDataSet ...
func GetDataSet(rows *sql.Rows) DataSet {
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
			if *v == nil {
				colmap[m] = ""
			} else {
				colmap[m] = string((*v).([]byte))
			}
			// colmap[m] = fmt.Sprintf("%s", *v)
		}
		dataset = append(dataset, colmap)
	}

	return dataset
}
