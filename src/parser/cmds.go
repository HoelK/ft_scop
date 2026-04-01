package parser

import "fmt"
import "errors"
import "strconv"

func mlib(data *Data, args []string) error {
	if (len(args) < 1) { return errors.New("[mtllib] Missing file name") }

	for i := 1; i < len(args); i++ {
		path := data.Path + args[i]
		parseMtl(data, path)
		var mat *Material
		mat = &data.Mtls[len(data.Mtls) - 1]
		fmt.Println("[Material]", mat.Name)
		fmt.Println("[Ns]", mat.Ns)
		fmt.Println("[Ka]", mat.Ka)
		fmt.Println("[Kd]", mat.Kd)
		fmt.Println("[Ks]", mat.Ks)
		fmt.Println("[Ni]", mat.Ni)
		fmt.Println("[D]", mat.D)
		fmt.Println("[Illum]", mat.Illum)
	}
	return nil
}

func o(data *Data, args []string) error {
	if (len(args) < 2) { return errors.New("[o] Missing Object name") }

	var newObj Object
	newObj.Name = args[1]
	data.Objs = append(data.Objs, newObj)
	fmt.Println("[LOG][o]", args[1], " Object added")
	return nil
}

func v(data *Data, args []string) error {
	var err		error
	var vtx		Vertex

	err = checkArgs(args, 4)

	if vtx.X, err = strconv.ParseFloat(args[1], 64); err != nil { return errors.New("[v] Failed Convertion : [\"" + args[1] + "\"] to [float64]") }
	if vtx.Y, err = strconv.ParseFloat(args[2], 64); err != nil { return errors.New("[v] Failed Convertion : [\"" + args[2] + "\"] to [float64]") }
	if vtx.Z, err = strconv.ParseFloat(args[3], 64); err != nil { return errors.New("[v] Failed Convertion : [\"" + args[3] + "\"] to [float64]") }

	data.Objs[0].Vtxs = append(data.Objs[0].Vtxs, vtx)
	fmt.Println("[LOG][v] Vertex added")
	return nil
}

func f(data *Data, args []string) error {
	var fc		Face
	var err		error
	var buf		int64

	if (len(args) < 4) { return errors.New("[f] Missing parameters (at least 3 required)") }

	for i := 0; i < (len(args) - 1); i++ {
		if buf, err = strconv.ParseInt(args[i + 1], 10, 32); err != nil {
			return errors.New("[f] Failed Convertion : [\"" + args[i + 1] + "\"] to [int64]")
		}
		for buf < 0 { fc.Vids[i] = int64(len(data.Objs[i].Vtxs)) - fc.Vids[i] }
		if (buf > int64(len(data.Objs[0].Vtxs))) { return errors.New("[f] " + args[i] + " Index Out of Range") }
		fc.Vids = append(fc.Vids, buf)
	}

	data.Objs[0].Fcs = append(data.Objs[0].Fcs, fc)
	fmt.Println("[LOG][f] Face added")
	return nil
}
