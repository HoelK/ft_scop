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
	Ka		float64		
	Kd		float64
	Ks		float64
	Ni		float64
	D		float64
	Illum	int8
}

type Object struct {
	Name	string
	Vtxs	[]Vertex
	Fcs		[]Face
	Mtls	[]Material
}

type Data struct {
	Objs []Object
}
