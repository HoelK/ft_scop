package parser

import "errors"
import "strconv"

func newmtl(data *Data, args []string) (*Material, error) {
	var mtl Material
	
	if (len(args) < 2) { return nil, errors.New("[newmtl] Missing argument (mtl name needed)") }
	if (len(args) > 2) { return nil, errors.New("[newmtl] Too much arguments") }

	mtl.Name = args[1]
	data.Mtls = append(data.Mtls, mtl)
	return &data.Mtls[len(data.Mtls) - 1], nil
}

func Ns(mtl *Material, args []string) error {
	var err error

	if (len(args) < 2) { return errors.New("[Ns] Missing argument") }
	if (len(args) > 2) { return errors.New("[Ns] Too much arguments") }

	if mtl.Ns, err = strconv.ParseFloat(args[1], 64); err != nil { return errors.New("[Ns] Failed Convertion [\"" + args[1] + "\"] to [float64]") }
	return nil
}

func Ka(mtl *Material, args []string) error {
	var err error

	if (len(args) < 4) { return errors.New("[Ka] Missing argument") }
	if (len(args) > 4) { return errors.New("[Ka] Too much arguments") }

	for i := 0; i < 3; i++ {
		if mtl.Ka[i], err = strconv.ParseFloat(args[1], 64); err != nil { return errors.New("[Ka] Failed Convertion [\"" + args[1] + "\"] to [float64]") }
	}
	return nil
}
func Kd(mtl *Material, args []string) error {
	var err error

	if (len(args) < 4) { return errors.New("[Kd] Missing argument") }
	if (len(args) > 4) { return errors.New("[Kd] Too much arguments") }

	for i := 0; i < 3; i++ {
		if mtl.Kd[i], err = strconv.ParseFloat(args[1], 64); err != nil { return errors.New("[Kd] Failed Convertion [\"" + args[1] + "\"] to [float64]") }
	}
	return nil
}

func Ks(mtl *Material, args []string) error {
	var err error

	if (len(args) < 4) { return errors.New("[Ks] Missing argument") }
	if (len(args) > 4) { return errors.New("[Ks] Too much arguments") }

	for i := 0; i < 3; i++ {
		if mtl.Ks[i], err = strconv.ParseFloat(args[1], 64); err != nil { return errors.New("[Ks] Failed Convertion [\"" + args[1] + "\"] to [float64]") }
	}
	return nil
	
}

func Ni(mtl *Material, args []string) error {
	var err error

	if (len(args) < 2) { return errors.New("[Ni] Missing argument") }
	if (len(args) > 2) { return errors.New("[Ni] Too much arguments") }

	if mtl.Ni, err = strconv.ParseFloat(args[1], 64); err != nil { return errors.New("[Ni] Failed Convertion [\"" + args[1] + "\"] to [float64]") }
	return nil
}

func d(mtl *Material, args []string) error {
	var err error

	if (len(args) < 2) { return errors.New("[d] Missing argument") }
	if (len(args) > 2) { return errors.New("[d] Too much arguments") }

	if mtl.D, err = strconv.ParseFloat(args[1], 64); err != nil { return errors.New("[d] Failed Convertion [\"" + args[1] + "\"] to [float64]") }
	return nil
}

func illum(mtl *Material, args []string) error {
	var err error

	if (len(args) < 2) { return errors.New("[illum] Missing argument") }
	if (len(args) > 2) { return errors.New("[illum] Too much arguments") }

	if mtl.Illum, err = strconv.ParseInt(args[1], 10, 64); err != nil { return errors.New("[illum] Failed Convertion [\"" + args[1] + "\"] to [int64]") }
	return nil
}
