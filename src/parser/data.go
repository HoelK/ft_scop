package parser

type Vertex struct {
	X float64
	Y float64
	Z float64
}

type Face struct {
	Vids	[]int64
}

type Material struct {
	Name	string
	Ns		float64
	Ka		[3]float64
	Kd		[3]float64
	Ks		[3]float64
	Ni		float64
	D		float64
	Illum	int64
}

type MtlFile struct {
	Name	string
	Mtls	[]Material
}

type Object struct {
	Name	string
	Vtxs	[]Vertex
	Fcs		[]Face
	Mtls	[]Material
}

type Data struct {
	Path string
	Objs []Object
	Mtls []Material
}
