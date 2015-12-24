package main

import (
	"common/dragonbone"
	imgutil "common/img"
	"flag"
	"image"
	"log"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

func main() {
	xml := flag.String("xml", "texture.xml", "input xml")
	pngName := flag.String("png", "texture.png", "input png")
	outputDir := flag.String("outDir", "out", "output directory")
	flag.Parse()
	img, err := imgutil.LoadImg(*pngName)
	if err != nil {
		log.Fatal(err)
	}

	texture, err := dragonbone.DecodeXml(*xml)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range texture.SubTexture {
		name := filepath.Join(*outputDir, s.Name+".png")
		parentDir := filepath.Dir(name)
		if _, err := os.Stat(parentDir); os.IsNotExist(err) {
			err := os.MkdirAll(parentDir, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
		err := Extract(img, name, s.X, s.Y, s.Width, s.Height, s.Rotated)
		if err != nil {
			log.Println(err)
		}
	}
}

func Extract(img image.Image, name string, X, Y, W, H float32, rotated bool) error {
	dst := imaging.Crop(img, image.Rect(int(X), int(Y), int(X+W), int(Y+H)))
	if rotated {
		log.Fatal("图片被旋转过")
	}
	return imgutil.SaveImg(name, dst)
}
