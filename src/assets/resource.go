package assets

import (
	"embed"
	"io"

	resource "github.com/quasilyte/ebitengine-resource"
	_ "image/png"
)

const (
	ImgNone resource.ImageID = iota
	ImgMap
)

func RegisterImageResources(loader *resource.Loader) {
	imageResources := map[resource.ImageID]resource.ImageInfo{
		ImgMap: {Path: "map.png"},
	}
	for id, res := range imageResources {
		loader.ImageRegistry.Set(id, res)
		loader.LoadImage(id)
	}
}

func OpenAssetFunc(path string) io.ReadCloser {
	f, err := gameAssets.Open("resources/" + path)
	if err != nil {
		panic(err)
	}
	return f
}

//go:embed all:resources
var gameAssets embed.FS
