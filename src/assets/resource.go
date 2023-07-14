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
	ImgMiniMap
	ImgMarkerIdle
	ImgMarkerSelected
	ImgMarkerLocked
	ImgBtnIdle
	ImgBtnHover
	ImgBtnPressed
	ImgBtnDisabled
)

func RegisterImageResources(loader *resource.Loader) {
	imageResources := map[resource.ImageID]resource.ImageInfo{
		ImgMap:     {Path: "map.png"},
		ImgMiniMap: {Path: "mini-map.png"},

		ImgMarkerIdle:     {Path: "marker-idle.png"},
		ImgMarkerSelected: {Path: "marker-selected.png"},
		ImgMarkerLocked:   {Path: "marker-locked.png"},

		ImgBtnIdle:     {Path: "btn-idle.png"},
		ImgBtnHover:    {Path: "btn-hover.png"},
		ImgBtnPressed:  {Path: "btn-pressed.png"},
		ImgBtnDisabled: {Path: "btn-disabled.png"},
	}

	for id, res := range imageResources {
		loader.ImageRegistry.Set(id, res)
		loader.LoadImage(id)
	}
}

const (
	FontNone resource.FontID = iota
	FontDefault
	FontSmall
)

func RegisterFontResources(loader *resource.Loader) {
	fontResources := map[resource.FontID]resource.FontInfo{
		FontDefault: {Path: "fibberish.ttf", Size: 12},
		FontSmall:   {Path: "PrintChar21.ttf", Size: 6},
	}
	for id, res := range fontResources {
		loader.FontRegistry.Set(id, res)
		loader.LoadFont(id)
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
