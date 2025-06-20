# The Umbra Game Engine

**Umbra** is a cross-platform, data-driven game engine written in Go, built for crafting 2D and 3D experiences with clean scene structure, low-poly rendering, and scriptable behaviour.

Currently in early development. The goal is to provide a flexible, minimalist, and developer-friendly engine with support for both 2D and 3D game creation.

---

## 🚧 Project Status

**Phase 1:** Core Framework & 2D MVP

- [ ] Game loop with Ebiten
- [x] YAML-based scene file support
- [ ] Sprite rendering system
- [ ] Input handling
- [x] Scene graph and node hierarchy

---

## ✨ Features (Planned)

- [ ] Modular architecture with reusable systems
- [ ] Data-driven scene files (`.scene`, `.material`)
- [ ] 2D rendering with Ebiten
- [ ] 3D rendering and shader pipeline
- [ ] Go + Lua scripting support
- [ ] Cross-platform: macOS, Linux, Windows

---

## 🔧 Getting Started

### Prerequisites

- Go 1.22+
- Git

### Setup

```bash
git clone https://github.com/Umbra-Engine/umbra.git
cd umbra
```

### Directory Structure

```
cmd/umbra_demo       # Demo entry point
engine/              # Core engine modules
  ├── components/    # Game object components
  ├── constants/     # Project constants
  ├── core/          # Game loop and timing
  ├── logger/        # Log systems
  ├── mathx/         # Math equation handling
  ├── runtime/       # Runtime handling
  ├── scene/         # Scene graph and loaders
  ├── scripting/     # Custom script handling

```

---

## 🗺 Roadmap

You can view the full roadmap in the [GitHub Project]("https://github.com/orgs/Umbra-Engine/projects/1") (WIP).

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE)

---

## 🤝 Contributing

This engine is in active early development. Contributions will be welcome once the core is stabilized.

---

## 👤 Author

Made by Gabriel Savard
[www.gabsavard.com](www.gabsavard.com)
