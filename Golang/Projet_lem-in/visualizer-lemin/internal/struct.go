package visualizer_lemin

import (
	"github.com/g3n/engine/animation"
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gui/assets/icon"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/helper"
)

type Link struct {
	Current string
	Target  string
}

type Room struct {
	r_type string
	Name   string
	X      int
	Y      int
}

type Ant struct {
	Move int    // index du mouvement de la fourmis
	Name string // nom de la fourmis
	Room string // Salle de depart de la fourmis
}

// ------------- G3N -------------
type G3nAnt struct {
	*app.Application                // Embedded application object
	Grid             *helper.Grid   // Grid helper
	ViewGrid         bool           // Grid helper visible flag
	CamPos           math32.Vector3 // Initial camera position
	Models           []*core.Node   // Models being shown
	Scene            *core.Node
	Cam              *camera.Camera
	Orbit            *camera.OrbitControl
	Anims            []*animation.Animation
}

const (
	checkON  = icon.CheckBox
	checkOFF = icon.CheckBoxOutlineBlank
)
