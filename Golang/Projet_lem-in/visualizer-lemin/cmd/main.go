package main

import (
	"fmt"
	"time"

	internal "visualizer-lemin/internal"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
)

// Command : ./lem-in example/example00.txt | go run ./cmd/main.go

func main() {
	input := internal.GetInput()
	// internal.GetInput()
	nb_ants, rooms, links, ants := internal.ParseFile(input)
	fmt.Println("\nnb_ants  : ", nb_ants)
	fmt.Println("\nrooms : ", rooms)
	fmt.Println("\nlinks : ", links)
	fmt.Println("\nants : ", ants)

	// Crée une application G3N
	Ga := new(internal.G3nAnt)
	a := app.App()
	Ga.Application = a
	Ga.Scene = core.NewNode()
	Ga.Scene.SetName("main scene")
	// Ajoute lumière ambiante blanche
	ambLight := light.NewAmbient(math32.NewColor("white"), 0.5)
	Ga.Scene.Add(ambLight)
	// joute une lumière directionele blanche
	dirLight := light.NewDirectional(math32.NewColor("white"), 1.0)
	dirLight.SetPosition(1, 0, 0)
	Ga.Scene.Add(dirLight)
	// Adds a Grid helper to the Scene initially not visible
	Ga.Grid = helper.NewGrid(50, 1, &math32.Color{0.4, 0.4, 0.4})
	Ga.ViewGrid = true
	Ga.Grid.SetVisible(Ga.ViewGrid)
	Ga.Scene.Add(Ga.Grid)
	// Sets the initial camera position
	Ga.CamPos = math32.Vector3{0, 15, 15}
	Ga.Cam = camera.New(1)
	Ga.Cam.SetPositionVec(&Ga.CamPos)
	Ga.Cam.LookAt(&math32.Vector3{0, 0, 0}, &math32.Vector3{0, 0.5, 0})
	Ga.Orbit = camera.NewOrbitControl(Ga.Cam)
	// Appel de la fonction pour générer les salles
	Ga.GenerateRooms(rooms)
	// Appel la fonction pour générer les couloirs
	Ga.CreateCorridors(links, rooms)
	// Appel la fonction pour créer les fourmis et les initialises dans la salle start
	// Ga.CreateAnts(&ants)
	// Set up callback to update viewport and camera aspect ratio when the window is resized
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := a.GetSize()
		a.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		Ga.Cam.SetAspect(float32(width) / float32(height))
	}
	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)
	Ga.BuildGui()
	// Set background color to gray
	a.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)
	// Ga.MoveAnts(&ants)
	// Run application main render loop
	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(Ga.Scene, Ga.Cam)
	})

}
