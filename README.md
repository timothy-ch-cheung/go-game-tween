# go-game-tween
Demo combining Ebitengine, EbitenUI, gween, ebiten-camera and ebitengine-resource to create a game world map that pans to each level (including a minimap on the top right).

![](docs/map-tween.gif?raw=true)

Build WASM:
```shell
env GOOS=js GOARCH=wasm go build -o mapTween-1-0-1.wasm github.com/timothy-ch-cheung/go-game-tween

```
