# ascii-graph

`ascii-graph` is a lightweight Go package/application for rendering terminal-friendly charts.

`
┌─────────────────────────────────────────┐
│         ChartRenderer (interface)       │
│  Render(g *Graph, canvas Canvas) error  │
│  Name() string                          │
└─────────────────────────────────────────┘
              │
    ┌─────────┴───────────┐
    │                     │
┌───▼──────┐          ┌───▼─────────┐
│Cartesian │          │  PieRenderer │
│Renderer  │          │              │
│(abstract)│          └──────────────┘
└───┬──────┘
    │
    ├── BarRenderer
    ├── LineRenderer  
    └── ScatterRenderer`
