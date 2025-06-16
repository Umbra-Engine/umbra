package runtime

import (
	"github.com/Umbra-Engine/umbra/engine/logger"
	"github.com/Umbra-Engine/umbra/engine/scene"
	"time"
)

type SceneRunner struct {
	scene      *scene.Scene
	updatables []Updatable
	running    bool
}

type Updatable interface {
	Update(dt float64)
}

func (sr *SceneRunner) InitScene(scene *scene.Scene) {
	if scene == nil {
		logger.Fatal("InitScene called with nil scene")
	}
	sr.scene = scene
	sr.updatables = nil

	for i := range scene.Objects {
		collectUpdatables(&scene.Objects[i], &sr.updatables)
	}
}

func (sr *SceneRunner) Run(targetFPS int) {
	// Default to 60 if there is no specified FPS
	if targetFPS <= 0 {
		targetFPS = 60
	}

	ticker := time.NewTicker(time.Second / time.Duration(targetFPS))
	defer ticker.Stop()

	sr.running = true
	lastTime := time.Now()

	for sr.running {
		now := time.Now()
		dt := now.Sub(lastTime).Seconds()
		lastTime = now

		for _, u := range sr.updatables {
			u.Update(dt)
		}
		<-ticker.C
	}
}

type Stoppable interface {
	OnStop()
}

func (sr *SceneRunner) Stop() {
	sr.running = false
	for _, u := range sr.updatables {
		if s, ok := u.(Stoppable); ok {
			s.OnStop()
		}
	}
}

func collectUpdatables(obj *scene.GameObject, updatables *[]Updatable) {
	for _, comp := range obj.RuntimeComponents {
		if s, ok := comp.(Updatable); ok {
			*updatables = append(*updatables, s)
			if hasOnStart, ok := comp.(interface{ OnStart() }); ok {
				hasOnStart.OnStart()
			}
		}
	}

	for i := range obj.Children {
		collectUpdatables(&obj.Children[i], updatables)
	}
}
