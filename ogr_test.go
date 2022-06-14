package gdal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvelope3D(t *testing.T) {

	env0 := Envelope3D{}
	assert.False(t, env0.IsInit())

	geom, _ := CreateFromWKT("LINESTRING (-1 -2 -3, 1 2 3)", SpatialReference{})
	env1 := geom.Envelope3D()
	assert.True(t, env1.IsInit())
	assert.Equal(t, env1.MinX(), -1.0)
	assert.Equal(t, env1.MinY(), -2.0)
	assert.Equal(t, env1.MinZ(), -3.0)
	assert.Equal(t, env1.MaxX(), 1.0)
	assert.Equal(t, env1.MaxY(), 2.0)
	assert.Equal(t, env1.MaxZ(), 3.0)

	env0.SetMinX(-4)
	assert.True(t, env0.IsInit())
	env0.SetMinY(-5)
	env0.SetMinZ(-6)
	env0.SetMaxX(4)
	env0.SetMaxY(5)
	env0.SetMaxZ(6)

	assert.Equal(t, env0.MinX(), -4.0)
	assert.Equal(t, env0.MinY(), -5.0)
	assert.Equal(t, env0.MinZ(), -6.0)
	assert.Equal(t, env0.MaxX(), 4.0)
	assert.Equal(t, env0.MaxY(), 5.0)
	assert.Equal(t, env0.MaxZ(), 6.0)

	assert.True(t, env0.Contains(env1))
	assert.False(t, env1.Contains(env0))
	assert.True(t, env0.Intersects(env1))
	assert.True(t, env1.Intersects(env0))

	env_union := env0.Union(env1)
	assert.Equal(t, env_union.MinX(), env0.MinX())
	assert.Equal(t, env_union.MinY(), env0.MinY())
	assert.Equal(t, env_union.MinZ(), env0.MinZ())
	assert.Equal(t, env_union.MaxX(), env0.MaxX())
	assert.Equal(t, env_union.MaxY(), env0.MaxY())
	assert.Equal(t, env_union.MaxZ(), env0.MaxZ())

	env_is := env1.Intersect(env0)
	assert.Equal(t, env_is.MinX(), env1.MinX())
	assert.Equal(t, env_is.MinY(), env1.MinY())
	assert.Equal(t, env_is.MinZ(), env1.MinZ())
	assert.Equal(t, env_is.MaxX(), env1.MaxX())
	assert.Equal(t, env_is.MaxY(), env1.MaxY())
	assert.Equal(t, env_is.MaxZ(), env1.MaxZ())

}
