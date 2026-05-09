package app

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"io/fs"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"

	"paper_quarters/internal/assets"
)

type imageCache struct {
	images map[string]*ebiten.Image
}

func newImageCache() *imageCache {
	return &imageCache{images: make(map[string]*ebiten.Image)}
}

func (c *imageCache) load(name string) (*ebiten.Image, error) {
	if img, ok := c.images[name]; ok {
		return img, nil
	}

	file, err := assets.FS.Open(name)
	if err != nil {
		return nil, fmt.Errorf("open image %s: %w", name, err)
	}
	defer file.Close()

	src, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("decode image %s: %w", name, err)
	}

	img := ebiten.NewImageFromImage(src)
	c.images[name] = img
	return img, nil
}

func readDirSorted(dir string) ([]fs.DirEntry, error) {
	files, err := assets.FS.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read embedded dir %s: %w", dir, err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	return files, nil
}

func cardValueFromFileName(fileName string) (int, error) {
	base := strings.TrimSuffix(fileName, path.Ext(fileName))
	parts := strings.Split(base, ".")
	value, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("parse card value from %s: %w", fileName, err)
	}
	return value, nil
}
