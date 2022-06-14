package go_checknumber

func CekGanjilGenap(number ...int) []string {
	var data []string
	for _, v := range number {
		if v%2 == 0 {
			data = append(data, "genap")
		} else {
			data = append(data, "ganjil")
		}
	}
	return data
}