package parser

import "errors"
import "strconv"

func checkArgs(args []string, amount int) error {
	if len(args) > amount { return errors.New("[" + args[0] + "] Too much arguments (" + string(amount) + " needed)") }
	if len(args) < amount { return errors.New("[" + args[0] + "] Missing arguments (" + string(amount) + " needed)") }
	return nil
}

func parseRGB(name string, args []string, rgb *[3]float64) (error) {
	var err error

	if err := checkArgs(args, 4); err != nil { return err }
	for i := 0; i < 3; i++ {
		if rgb[i], err = strconv.ParseFloat(args[1], 64); err != nil {
			return errors.New("[" + name + "] Failed Convertion [\"" + args[i + 1] + "\"] to [float64]")
		}
	}
	return nil
}

func parseFloat(name string, args []string, num *float64) error {
	var err error

	if err := checkArgs(args, 2); err != nil { return err }
	if *num, err = strconv.ParseFloat(args[1], 64); err != nil {
		return errors.New("[" + name + "] Failed Convertion [\"" + args[1] + "\"] to [float64]")
	}
	return nil
}

func newmtl(data *Data, args []string) (*Material, error) {
	var mtl Material

	if err := checkArgs(args, 2); err != nil { return nil, err }

	mtl.Name = args[1]
	data.Mtls = append(data.Mtls, mtl)
	return &data.Mtls[len(data.Mtls) - 1], nil
}


func Ns(mtl *Material, args []string) error { return parseFloat(args[0], args, &mtl.Ns) }
func Ni(mtl *Material, args []string) error { return parseFloat(args[0], args, &mtl.Ni) }
func d(mtl *Material, args []string) error { return parseFloat(args[0], args, &mtl.D) }
func Ka(mtl *Material, args []string) error { return parseRGB(args[0], args, &mtl.Ka) }
func Kd(mtl *Material, args []string) error { return parseRGB(args[0], args, &mtl.Kd) }
func Ks(mtl *Material, args []string) error { return parseRGB(args[0], args, &mtl.Ks) }

func illum(mtl *Material, args []string) error {
	var err error

	if err := checkArgs(args, 2); err != nil { return err }

	if mtl.Illum, err = strconv.ParseInt(args[1], 10, 64); err != nil { return errors.New("[illum] Failed Convertion [\"" + args[1] + "\"] to [int64]") }
	return nil
}
