package main

func v(args []string) ([3]float32, error) {
	var a [3]float32
	a[0] = 0
	if (len(args) != 0) { return a, nil }
	return a, nil
}

func f(args []string) ([]int32, error) {
	a := make([]int32, 3)
	a[0] = 0
	if (len(args) != 0) { return a, nil }
	return a, nil
}

func o(args []string) ([]string, error) {
	a := make([]string, 1)
	if (len(args) < 2) { return a, nil }
	a[0] = args[1]
	return a, nil
}

func mlib(args []string) ([]string, error) {
	a := make([]string, 1)
	a[0] = ""
	return a, nil
}

