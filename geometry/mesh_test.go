package geometry

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/YaacovCa/vector"
	"github.com/stretchr/testify/assert"
)

func TestOBJ(t *testing.T) {

	//arrange
	vertices := []vector.Vector3{
		vector.Vector3Zero(),
		vector.Vector3Right(),
		vector.Vector3Up(),
	}
	triangles := []Tri{
		Tri{0, 1, 2},
	}
	m := Mesh{triangles, vertices}
	out := bytes.Buffer{}

	//act
	err := m.ToOBJ(&out)

	//assert
	assert.NoError(t, err)

	scanner := bufio.NewScanner(&out)

	scanner.Scan()
	assert.Equal(t, "v 0.00 0.00 0.00", scanner.Text())
	scanner.Scan()
	assert.Equal(t, "v 1.00 0.00 0.00", scanner.Text())
	scanner.Scan()
	assert.Equal(t, "v 0.00 1.00 0.00", scanner.Text())
	scanner.Scan()
	assert.Equal(t, "f 1 2 3", scanner.Text())

}