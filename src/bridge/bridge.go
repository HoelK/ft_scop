package main

/*
#cgo CFLAGS: -Wall -Werror -Wextra

#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>

typedef uint32_t ui32;

typedef struct s_vertex
{
	float x;
	float y;
	float z;
} t_vertex;

typedef struct s_face
{
	t_vertex	**vtx;
	ui32		count;
} t_face;

typedef struct s_material
{
	char	*name;
	float	ns;
	float	ka[3];
	float	kd[3];
	float	ks[3];
	float	ni;
	float	d;
	int		illum;
} t_material;

typedef struct s_object
{
	char		*name;
	t_vertex	*vtxs;
	t_face		*fcs;
	t_material	mtl;
	ui32		v_count;
	ui32		f_count;
	bool		smooth;
} t_object;

typedef struct s_data
{
	t_object	obj;
} t_data;

t_vertex	**vtx_alloc(ui32 size) { return (malloc(sizeof(t_vertex *) * size)); }

t_face		*get_face(t_face *lst, ui32 i)			{ return (&(lst[i])); }
t_object	*get_obj(t_object *lst, ui32 i)			{ return (&(lst[i])); }
t_vertex	*get_vertex(t_vertex *lst, ui32 i)		{ return (&(lst[i])); }
t_material	*get_material(t_material *lst, ui32 i)	{ return (&(lst[i])); }
t_vertex	**get_vtx(t_vertex **lst, ui32 i)		{ return (&(lst[i])); }

void		free_data(t_data *data)
{
	free(data->obj.vtxs);
	free(data->obj.fcs);
	free(data->obj.mtl.name);
}

void		print_vertex(t_vertex vtx) { printf("v %f %f %f\n", vtx.x, vtx.y, vtx.z); }
void		print_material(t_material mtl)
{
	printf("name %s\n", mtl.name);
	printf("Ns %f\n", mtl.ns);
	printf("Ka %f %f %f\n", mtl.ka[0], mtl.ka[1], mtl.ka[2]);
	printf("Kd %f %f %f\n", mtl.kd[0], mtl.kd[1], mtl.kd[2]); 
	printf("Ks %f %f %f\n", mtl.ks[0], mtl.ks[1], mtl.ks[2]);
	printf("Ni %f\n", mtl.ni);
	printf("d %f\n", mtl.d);
	printf("illum %d\n", mtl.illum);
}

void		print_data(t_data *data)
{
	printf("o %s\n", data->obj.name);
	for (ui32 y = 0; y < data->obj.v_count; y++)
	print_vertex(data->obj.vtxs[y]);
	print_material(data->obj.mtl);
	printf("s %d\n", data->obj.smooth);
}

*/
import "C"
import "os"
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
	C.print_data(cdata)
	// redoFile(&data)
	return cdata
}

func main() {
	if len(os.Args) != 2 { log.Fatal("File required has argument") }

	parse(C.CString(os.Args[1]))
}
