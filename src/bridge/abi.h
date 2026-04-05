# pragma once

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

static inline t_vertex	**vtx_alloc(ui32 size) { return (malloc(sizeof(t_vertex *) * size)); }

static inline t_face		*get_face(t_face *lst, ui32 i)			{ return (&(lst[i])); }
static inline t_object	*get_obj(t_object *lst, ui32 i)			{ return (&(lst[i])); }
static inline t_vertex	*get_vertex(t_vertex *lst, ui32 i)		{ return (&(lst[i])); }
static inline t_material	*get_material(t_material *lst, ui32 i)	{ return (&(lst[i])); }
static inline t_vertex	**get_vtx(t_vertex **lst, ui32 i)		{ return (&(lst[i])); }

static inline void		free_data(t_data *data)
{
	free(data->obj.vtxs);
	free(data->obj.fcs);
	free(data->obj.mtl.name);
}

static inline void		print_vertex(t_vertex vtx) { printf("v %f %f %f\n", vtx.x, vtx.y, vtx.z); }
static inline void		print_material(t_material mtl)
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

static inline void		print_data(t_data *data)
{
	printf("o %s\n", data->obj.name);
	for (ui32 y = 0; y < data->obj.v_count; y++)
	print_vertex(data->obj.vtxs[y]);
	print_material(data->obj.mtl);
	printf("s %d\n", data->obj.smooth);
}
