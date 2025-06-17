package assets

import (
	"fmt"
	"image"
	"image/draw"
	"os"
	"path/filepath"
	"sync"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// Manager handles loading and caching assets like textures.
type Manager struct {
	textures map[string]uint32
	mu       sync.RWMutex
	basePath string
}

// NewManager creates a new asset manager.
func NewManager(basePath string) *Manager {
	return &Manager{
		textures: make(map[string]uint32),
		basePath: basePath,
	}
}

// LoadTexture loads an image file (PNG, JPG) and caches it.
func (m *Manager) LoadTexture(name string) (uint32, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// If already loaded
	if tex, exists := m.textures[name]; exists {
		return tex, nil
	}

	// Construct full path
	fullPath := filepath.Join(m.basePath, name)
	file, err := os.Open(fullPath)
	if err != nil {
		return uint32(0), fmt.Errorf("failed to open texture %q: %w", name, err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return uint32(0), fmt.Errorf("failed to decode image %q: %w", name, err)
	}

	// Convert to RGBA format
	rgbaImg := image.NewRGBA(img.Bounds())
	draw.Draw(rgbaImg, rgbaImg.Bounds(), img, image.Point{}, draw.Src)

	// Generate OpenGL texture
	var texID uint32
	gl.GenTextures(1, &texID)
	gl.BindTexture(gl.TEXTURE_2D, texID)

	width := int32(rgbaImg.Rect.Size().X)
	height := int32(rgbaImg.Rect.Size().Y)
	data := gl.Ptr(rgbaImg.Pix)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		width,
		height,
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		data,
	)

	// Texture parameters
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	// Generate mipmaps
	gl.GenerateMipmap(gl.TEXTURE_2D)

	// Unbind texture
	gl.BindTexture(gl.TEXTURE_2D, 0)

	// Cache and return
	m.textures[name] = texID
	return texID, nil
}

// GetTexture returns a previously loaded texture, or nil if not loaded
func (m *Manager) GetTexture(name string) (uint32, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if tex, ok := m.textures[name]; ok {
		return tex, nil
	}
	return uint32(0), fmt.Errorf("missing or invalid texture %s", name)
}
