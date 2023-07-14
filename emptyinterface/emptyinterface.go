package main

func main() {
	var randomValues interface{}

	_ = randomValues

	randomValues = "Jalan Bogorejo"

	randomValues = 20

	randomValues = true

	randomValues = []string{"Dandi", "Dian"}

	// type assertion pada empty interface
	var v interface{}

	v = 20

	// v = v * 9 (ini error)

	// dengen type assertion untuk menanggulangi error
	if value, ok := v.(int); ok == true {
		v = value * 9
	}

	// Empty interface (Map & Slice with Empty Interface)
	// deklarasi slice dengan empty interface
	rs := []interface{}{1, "Dandi", true, 2, "Irwanto", true}

	rm := map[string]interface{}{
		"Name": "Dandi",
		"Status": true,
		"Age": 23,
	}

	_, _ = rs, rm
}