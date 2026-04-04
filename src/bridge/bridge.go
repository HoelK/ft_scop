package main

//#cgo CFLAGS: -Wall -Werror -Wextra -I./src/bridge/abi.h
//#include "abi.h"
import "C"
import "log"
import "fmt"
import "bridge/src/parser"

func translateVertexs(vtxs []parser.Vertex) (*C.t_vertex, C.uint) {
	var v_count C.uint		= C.uint(len(vtxs))
	var cvtxs *C.t_vertex	= (*C.t_vertex)(C.malloc(C.sizeof_t_vertex * C.size_t(len(vtxs))))

	for i := 0; i < len(vtxs); i++ {
		cvtx := C.get_vertex(cvtxs, C.uint(i))
		cvtx.x = C.float(vtxs[i].X)
		cvtx.y = C.float(vtxs[i].Y)
		cvtx.z = C.float(vtxs[i].Z)
	}

	return cvtxs, v_count
}

func translateMaterial(mtl *parser.Material) (C.t_material) {
	var cmtl C.t_material

	if (mtl == nil) { return cmtl }
	cmtl.name = C.CString(mtl.Name)
	for j := 0; j < 3; j++ {
		cmtl.ka[j] = C.float(mtl.Ka[j])
		cmtl.ks[j] = C.float(mtl.Ks[j])
		cmtl.kd[j] = C.float(mtl.Kd[j])
	}
	cmtl.ns =	C.float(mtl.Ns)
	cmtl.ni =	C.float(mtl.Ni)
	cmtl.d =	C.float(mtl.D)
	cmtl.illum = C.int(mtl.Illum)

	return cmtl
}

func translateFaces(fcs []parser.Face, cvtxs *C.t_vertex) (*C.t_face, C.uint) {
	var f_count C.uint	= C.uint(len(fcs))
	var cfcs *C.t_face	= (*C.t_face)(C.malloc(C.sizeof_t_face * C.size_t(len(fcs))))

	for i := 0; i < len(fcs); i++ {
		cfc		:=	C.get_face(cfcs, C.uint(i))
		cfc.vtx	=	C.vtx_alloc(C.uint(len(fcs[i].Vids)))

		for y := 0; y < len(fcs[i].Vids); y++ {
			vtx		:=	C.get_vtx(cfc.vtx, C.uint(y))
			*vtx	=	C.get_vertex(cvtxs, C.uint(fcs[i].Vids[y]))
		}
	}

	return cfcs, f_count
}

func translate(data *parser.Data) *C.t_data {
	cdata := (*C.t_data)(C.malloc(C.sizeof_t_data))

	obj := &cdata.obj
	obj.name =	C.CString(data.Obj.Name)

	obj.vtxs, obj.v_count	= translateVertexs(data.Obj.Vtxs)
	obj.fcs, obj.f_count	= translateFaces(data.Obj.Fcs, obj.vtxs)
	obj.mtl					= translateMaterial(data.Obj.Mtl)
	obj.smooth				= C.bool(data.Obj.S)
	return cdata
}

//export parse
func parse(cpath *C.char) (*C.t_data) {
	var file parser.FILE
	var data parser.Data

	var path string = C.GoString(cpath)
	if (path == "") { log.Fatal("File required has argument") }

	file.Init(path)
	defer file.Fd.Close()
	data, err := parser.ParseObj(&file)
	if (err != nil) { log.Fatal(err) }
	fmt.Println(data.Obj.Name)
	cdata := translate(&data)
	return cdata
}

func main() {}
