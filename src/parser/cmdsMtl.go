package parser

import "errors"
import "strconv"

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

func parseInt(name string, args []string, num *int64) error {
	var err error

	if err := checkArgs(args, 2); err != nil { return err }
	if *num, err = strconv.ParseInt(args[1], 10, 64); err != nil {
		return errors.New("[illum] Failed Convertion [\"" + args[1] + "\"] to [int64]")
	}
	return nil
}

func newmtl(data *Data, args []string) (*Material, error) {
	var mtl *Material = new(Material)

	if err := checkArgs(args, 2); err != nil { return nil, err }

	mtl.Name = args[1]
	data.Mtls[mtl.Name] = mtl
	return data.Mtls[mtl.Name], nil
}

func Ns(mtl *Material, args []string)		error	{ return parseFloat(args[0], args, &mtl.Ns) }
func Ni(mtl *Material, args []string)		error	{ return parseFloat(args[0], args, &mtl.Ni) }
func D(mtl *Material, args []string)		error	{ return parseFloat(args[0], args, &mtl.D) }
func Ka(mtl *Material, args []string)		error	{ return parseRGB(args[0], args, &mtl.Ka) }
func Kd(mtl *Material, args []string)		error	{ return parseRGB(args[0], args, &mtl.Kd) }
func Ks(mtl *Material, args []string)		error	{ return parseRGB(args[0], args, &mtl.Ks) }
func Illum(mtl *Material, args []string)	error	{ return parseInt(args[0], args, &mtl.Illum)}
