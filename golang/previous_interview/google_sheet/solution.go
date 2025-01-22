/*

Notes
- Rows/Cols? A-Z, 1 - 10
- Build the spreadsheet - method?
- Try to implement setCell first
- Then getCell

Data Structure
- Use a map[string]string
- map[string]string{ "A1": "45" }

Solution
- First handle setting a cell
- Then handle getting a cell
- Then handle expressions

setCell(target string, value string)
- map[target] = value

getCell(input string)
- lookup in map and return
- if there's an "=" then add operation
	- split the values
	- take operand1 + operand2

Todo
- Create a OpenSheet type map[string]string
- Create a setCell method
- Create a getCell method
- Add expression handling

*/

type OpenSheet map[string]string

func (os *OpenSheet) setCell(target, value string) {
	(*os)[target] = value
}

func (os *OpenSheet) getCell(target string) string {
	if val, ok := (*os)[target]; ok {

		if string(val[0]) == "=" {
			split := strings.Split(val[1:], "+")

			val1, _ := strconv.Atoi((*os)[split[0]])
			val2, _ := strconv.Atoi((*os)[split[1]])

			sum := val1 + val2

			return strconv.Itoa(sum)
		}

		return val
	}

	return ""
}
