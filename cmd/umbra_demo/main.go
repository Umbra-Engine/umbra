package main

import (
	"github.com/Umbra-Engine/umbra/engine/core"
	"github.com/Umbra-Engine/umbra/engine/logger"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"runtime"
)

type Game struct{}

func main() {
	runtime.LockOSThread() // Ensures all GLFW call runs on main OS thread (macOS requires this)

	if err := glfw.Init(); err != nil {
		logger.Fatal("failed to initialize GLFW: %v", err)
	}
	defer glfw.Terminate()

	config, err := core.LoadConfig("config.yaml")
	if err != nil {
		logger.Fatal("failed to log config: %v", err)
	}

	initWindow()
	window := createWindow(config.Window)
	initGL(config.Window)

	gameLoop(window)
	window.Destroy()
}

func initWindow() {
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}

func createWindow(windowConfig core.WindowConfig) *glfw.Window {
	window, err := glfw.CreateWindow(windowConfig.Width, windowConfig.Height, windowConfig.Title, nil, nil)
	if err != nil {
		logger.Fatal("failed to create the %s game window", windowConfig.Title)
	}
	window.MakeContextCurrent()
	return window
}

func initGL(windowConfig core.WindowConfig) {
	if err := gl.Init(); err != nil {
		logger.Fatal("failed to initialize OpenGL: %v", err)
	}
	gl.Viewport(0, 0, int32(windowConfig.Width), int32(windowConfig.Height))
	gl.ClearColor(
		float32(windowConfig.BackgroundColor.X),
		float32(windowConfig.BackgroundColor.Y),
		float32(windowConfig.BackgroundColor.Z),
		float32(windowConfig.BackgroundColor.W),
	)
}

func gameLoop(window *glfw.Window) {
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// TODO: Add Umbra-specific rendering logic here

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
