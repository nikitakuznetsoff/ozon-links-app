package transfomer

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const base = 62

func Encode(id int) string {
	var arr []string
	value := id
	for value > 0 {
		symbIndex := value % base
		if symbIndex < 10 {
			arr = append(arr, strconv.Itoa(symbIndex))
		} else if symbIndex < 36 {
			arr = append(arr, string(rune(97 + symbIndex - 10)))
		} else {
			arr = append(arr, string(rune(65 + symbIndex - 10 - 26)))
		}
		value /= base
	}
	var b strings.Builder
	for i := len(arr)-1; i >= 0; i-- {
		fmt.Fprintf(&b, arr[i])
	}
	return b.String()
}


func Decode(link string) (int, error) {
	id := 0
	for i, v := range link {
		symbIndex := 0
		if v >= 'A' && v <= 'Z' {
			symbIndex = int(v) - 65 + 10 + 26
		} else if v >= 'a' && v <= 'z' {
			symbIndex = int(v) - 97 + 10
		} else if v >= '0' && v <= '9' {
			symbIndex = int(v)
		} else {
			return 0, fmt.Errorf("unexpected character %c", v)
		}
		id += int(math.Pow(62, float64(len(link) - 1 - i))) * symbIndex
	}
	return id, nil
}
