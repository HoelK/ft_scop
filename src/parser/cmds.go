package parser

import "errors"
import "strconv"

func mlib(args []string) error {
	if (len(args) < 1) { return errors.New("[mtllib] Missing file name") }
	return nil
}

func o(args []string) error {
	if (len(args) < 2) { return errors.New("[o] Missing Object name") }
	return nil
}

func v(args []string) error {
	var err error
	var conv [3]float64
	if (len(args) != 4) { return nil }

	for i := 0; i < 3; i++ {
		if conv[i], err = strconv.ParseFloat(args[i + 1], 64); err != nil {
			return errors.New("[v] Failed Convertion : [\"" + args[i + 1] + "\"] to [float64]")
		}
	}
	return nil
}

func f(args []string) error {
	var err error
	var conv []int64
	for i := 0; i < 3; i++ {
		if conv[i], err = strconv.ParseInt(args[i + 1], 10, 64); err != nil {
			return errors.New("[f] Failed Convertion : [\"" + args[i + 1] + "\"] to [int64]")
		}
	}
	if (len(args) < 2) { return errors.New("[f] Failed Convertion") }
	return nil
}
