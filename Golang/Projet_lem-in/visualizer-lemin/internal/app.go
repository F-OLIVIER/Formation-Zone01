package visualizer_lemin

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

func (Ga *G3nAnt) GenerateRooms(rooms []Room) {
	for _, value := range rooms {
		x, y := value.X, value.Y
		// Créez une géométrie pour la salle
		//	- cube
		// geom := geometry.NewCube(0.5)
		//	- sphére
		geom := geometry.NewSphere(0.35, 50, 10)

		// Créez un matériau pour la salle
		if value.r_type == "start" {
			mat := material.NewStandard(math32.NewColor("green"))
			mat.SetTransparent(true)
			mat.SetOpacity(0.5)
			// Créez un nœud pour représenter la salle
			room := graphic.NewMesh(geom, mat)
			// Positionnez la salle aux coordonnées x et y
			room.SetPosition(float32(x), 0, float32(y))
			// Ajoutez la salle à la scène
			Ga.Scene.Add(room)
		} else if value.r_type == "end" {
			mat := material.NewStandard(math32.NewColor("red"))
			mat.SetTransparent(true)
			mat.SetOpacity(0.5)
			// Créez un nœud pour représenter la salle
			room := graphic.NewMesh(geom, mat)
			// Positionnez la salle aux coordonnées x et y
			room.SetPosition(float32(x), 0, float32(y))
			// Ajoutez la salle à la scène
			Ga.Scene.Add(room)
		} else {
			mat := material.NewStandard(math32.NewColor("blue"))
			mat.SetTransparent(true)
			mat.SetOpacity(0.5)
			// Créez un nœud pour représenter la salle
			room := graphic.NewMesh(geom, mat)
			// Positionnez la salle aux coordonnées x et y
			room.SetPosition(float32(x), 0, float32(y))
			// Ajoutez la salle à la scène
			Ga.Scene.Add(room)
		}
	}
}

func (ga *G3nAnt) CreateCorridors(links []Link, rooms []Room) {
	for _, value := range links {
		room1Name := value.Current
		var room1_X, room1_Y int
		room2Name := value.Target
		var room2_X, room2_Y int
		for _, room := range rooms {
			if room.Name == room1Name {
				room1_X = room.X
				room1_Y = room.Y
			}
			if room.Name == room2Name {
				room2_X = room.X
				room2_Y = room.Y
			}
		}

		// Calcule les coordonnées du milieu entre les deux salles
		centerX := (float32(room1_X) + float32(room2_X)) / 2
		centerY := (float32(room1_Y) + float32(room2_Y)) / 2
		// Calculez la direction entre les centres des deux salles
		direction := math32.NewVector3(float32(room2_X-room1_X), 0.0, float32(room2_Y-room1_Y))
		// Calculez la longueur du tube en fonction de la distance entre les deux centres des salles
		distance := direction.Length()
		// Normalizez la direction pour obtenir un vecteur unitaire
		direction.Normalize()
		// Créez un box entre les deux salles
		box := geometry.NewBox(0.2, 0.4, float32(distance))
		// Créez un maillage à partir du box
		mat := material.NewStandard(math32.NewColor("gray"))
		mat.SetTransparent(true)
		mat.SetOpacity(0.5)
		boxMesh := graphic.NewMesh(box, mat)
		// Positionnez le couloir au centre entre les deux salles
		boxMesh.SetPosition(centerX, 0, centerY)
		// Calculez l'angle de rotation
		rotationAngle := math32.Atan2(direction.X, direction.Z)
		// Faites pivoter le couloir autour de l'axe Z
		boxMesh.RotateOnAxis(&math32.Vector3{0, 1, 0}, rotationAngle)
		// Ajoutez le couloir à la scène
		ga.Scene.Add(boxMesh)
	}
}

func (Ga *G3nAnt) BuildGui() error {
	gui.Manager().Set(Ga.Scene)
	// Adds menu bar
	mb := gui.NewMenuBar()
	mb.SetLayoutParams(&gui.VBoxLayoutParams{Expand: 0, AlignH: gui.AlignWidth})
	Ga.Scene.Add(mb)
	// Create "View" menu and adds it to the menu bar
	m2 := gui.NewMenu()
	vGrid := m2.AddOption("View grid helper").SetIcon(checkOFF)
	vGrid.SetIcon(getIcon(Ga.ViewGrid))
	vGrid.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		Ga.ViewGrid = !Ga.ViewGrid
		vGrid.SetIcon(getIcon(Ga.ViewGrid))
		Ga.Grid.SetVisible(Ga.ViewGrid)
	})
	mb.AddMenu("View", m2)
	return nil
}
func getIcon(state bool) string {
	if state {
		return checkON
	} else {
		return checkOFF
	}
}
