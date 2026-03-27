package main

/*
#cgo CFLAGS: -Wall -Werror -Wextra

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
	float	ka;
	float	kd;
	float	ks;
	float	ni;
	float	d;
	int		illum;
} t_material;

typedef struct s_object
{
	t_vertex	*vtxs;
	t_face		*fcs;
	t_material	*mtls;
	ui32		v_count;
	ui32		f_count;
	ui32		m_count;
} t_object;

typedef struct s_data
{
	t_object	*objs;
	ui32		o_count;
} t_data;

t_vertex	**vtx_alloc(ui32 size) { return (malloc(sizeof(t_vertex *) * size)); }

t_face		*get_face(t_face *lst, ui32 i)			{ return (&(lst[i])); }
t_object	*get_obj(t_object *lst, ui32 i)			{ return (&(lst[i])); }
t_vertex	*get_vertex(t_vertex *lst, ui32 i)		{ return (&(lst[i])); }
t_material	*get_material(t_material *lst, ui32 i)	{ return (&(lst[i])); }
t_vertex	**get_vtx(t_vertex **lst, ui32 i)		{ return (&(lst[i])); }

void		free_data(t_data *data)
{
	for (ui32 i = 0; i < data->o_count; i++)
	{
		free(data->objs[i].vtxs);
		free(data->objs[i].fcs);
		free(data->objs[i].mtls);
	}
}

*/
import "C"
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

func translateMaterials(mtls []parser.Material) (*C.t_material, C.uint) {
	var m_count C.uint		= C.uint(len(mtls))
	var cmtls *C.t_material	= (*C.t_material)(C.malloc(C.sizeof_t_material * C.size_t(len(mtls))))

	for i := 0; i < len(mtls); i++ {
		cmtl := C.get_material(cmtls, C.uint(i))
		cmtl.name = C.CString(mtls[i].Name)
		cmtl.ns = C.float(mtls[i].Ns)
		cmtl.ka = C.float(mtls[i].Ka)
		cmtl.kd = C.float(mtls[i].Kd)
		cmtl.ks = C.float(mtls[i].Ks)
		cmtl.ni = C.float(mtls[i].Ni)
		cmtl.d = C.float(mtls[i].D)
		cmtl.illum = C.int(mtls[i].Illum)
	}

	return cmtls, m_count
}

func translateFaces(fcs []parser.Face, cvtxs *C.t_vertex) (*C.t_face, C.uint) {
	var f_count C.uint	= C.uint(len(fcs))
	var cfcs *C.t_face	= (*C.t_face)(C.malloc(C.sizeof_t_face * C.size_t(len(fcs))))

	for i := 0; i < len(fcs); i++ {
		cfc := C.get_face(cfcs, C.uint(i))
		cfc.vtx = C.vtx_alloc(C.uint(len(fcs[i].Vids)))

		for y := 0; i < len(fcs[i].Vids); y++ {
			vtx := C.get_vtx(cfc.vtx, C.uint(y))
			*vtx = C.get_vertex(cvtxs, C.uint(fcs[i].Vids[y]))
		}
	}

	return cfcs, f_count
}

func translate(data *parser.Data) *C.t_data {
	cdata := (*C.t_data)(C.malloc(C.sizeof_t_data))

	cdata.o_count = C.uint(len(data.Objs))
	cdata.objs = (*C.t_object)(C.malloc(C.sizeof_t_object * C.size_t(len(data.Objs))))

	for i := 0; i < len(data.Objs); i++ {
		obj :=			C.get_obj(cdata.objs, C.uint(i))

		obj.vtxs, obj.v_count	= translateVertexs(data.Objs[i].Vtxs)
		obj.fcs, obj.f_count	= translateFaces(data.Objs[i].Fcs, obj.vtxs)
		obj.mtls, obj.m_count	= translateMaterials(data.Objs[i].Mtls)
	}
	return cdata
}

//export ParseObj
func ParseObj(path *C.char) *C.t_data {
	var data parser.Data
	var cdata *C.t_data = translate(&data)

	return cdata
}

func main() {}
