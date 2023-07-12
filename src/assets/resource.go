package assets

import (
	"embed"
	"io"

	_ "image/png"

	resource "github.com/quasilyte/ebitengine-resource"
)

const (
	ImgNone resource.ImageID = iota
	ImgMap
	ImgMarkerIdle
	ImgMarkerSelected
	ImgMarkerLocked
)

func RegisterImageResources(loader *resource.Loader) {
	imageResources := map[resource.ImageID]resource.ImageInfo{
		ImgMap: {Path: "map.png"},

		ImgMarkerIdle:     {Path: "marker-idle.png"},
		ImgMarkerSelected: {Path: "marker-selected.png"},
		ImgMarkerLocked:   {Path: "marker-locked.png"},
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
