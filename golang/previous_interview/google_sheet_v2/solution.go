/*
Solution
- First implement NewSpreadsheet
- Then Set func
- type Row map[int]int

Spreadsheet
- sheet map[string][]Row
*/
package main

import "fmt"

type Row map[int]int

type Spreadsheet struct {
	sheet map[string][]Row
}

func NewSpreadsheet(columns []string) *Spreadsheet {
	// Initialize to empty rows
	// Loop over columns
	// Set each column
	s := &Spreadsheet{sheet: make(map[string][]Row)}
	for _, v := range columns {
		s.sheet[v] = make([]Row, 0)
	}

	return s
}

func (s *Spreadsheet) Get(row int, column string) int {
	// First get the column
	col := s.sheet[column]
	// Iterate over the array until i find the row
	for _, v := range col {
		if value, ok := v[row]; ok {
			return value
		}
	}
	return 0
}

func (s *Spreadsheet) Set(row int, column string, value int) {
	// Find the column
	// Create a new Row
	// Set the column to the row value
	prevCol := s.sheet[column]
	r := Row{row: value}
	// arr := []Row{r}
	s.sheet[column] = append(prevCol, r)
}

/*
Print Solution
- Start at each column at index i (0)
- Iterate over the rows and print where row == i
- Iterate while row < n

+------+--------+--------+--------+
| | "fips" | "pop" | "area" |
+======+========+========+========+
| 0 | 1001 | 200000 | 5000 |
+------+--------+--------+--------+
| 1 | 0 | 0 | 0 |
+------+--------+--------+--------+
| 2 | 0 | 0 | 0 |
+------+--------+--------+--------+
| 3 | 0 | 0 | 0 |
+------+--------+--------+--------+
| 4 | 0 | 0 | 0 |
+------+--------+--------+--------+
| 5 | 1002 | 0 | 0 |
+------+--------+--------+--------+
| ... | ... | ... | ... |


Should Print Like:
1001, 200000, 5000
0
0
0
10002
...

	{
		"fips": {
			0: 1001,
			5: 1002
		}
	}
*/

// TODO: This needs to be printing across all columns
func (s *Spreadsheet) PrintFirstN(rowCount int) {
	cols := s.sheet
	row := 0
	for row < rowCount {
		for key, _ := range cols {
			fmt.Printf("Printing col %s %d\n", key, s.Get(row, key))
		}
		row++

	}
}

