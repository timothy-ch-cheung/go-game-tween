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
	ImgNextBtnIdle
	ImgNextBtnHover
	ImgNextBtnPressed
	ImgNextBtnDisabled
	ImgPrevBtnIdle
	ImgPrevBtnHover
	ImgPrevBtnPressed
	ImgPrevBtnDisabled
)

func RegisterImageResources(loader *resource.Loader) {
	imageResources := map[resource.ImageID]resource.ImageInfo{
		ImgMap: {Path: "map.png"},

		ImgMarkerIdle:     {Path: "marker-idle.png"},
		ImgMarkerSelected: {Path: "marker-selected.png"},
		ImgMarkerLocked:   {Path: "marker-locked.png"},

		ImgNextBtnIdle:     {Path: "next-btn-idle.png"},
		ImgNextBtnHover:    {Path: "next-btn-hover.png"},
		ImgNextBtnPressed:  {Path: "next-btn-pressed.png"},
		ImgNextBtnDisabled: {Path: "next-btn-disabled.png"},

		ImgPrevBtnIdle:     {Path: "prev-btn-idle.png"},
		ImgPrevBtnHover:    {Path: "prev-btn-hover.png"},
		ImgPrevBtnPressed:  {Path: "prev-btn-pressed.png"},
		ImgPrevBtnDisabled: {Path: "prev-btn-disabled.png"},
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
