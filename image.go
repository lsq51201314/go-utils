package utils

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"image/png"

	"github.com/nfnt/resize"
)

// 图片实例
type Image struct{}

// 缩放图片
func (i Image) Zoom(src []byte, width, height uint) (data []byte, err error) {
	var img image.Image
	if img, _, err = image.Decode(bytes.NewReader(src)); err != nil {
		return
	}
	//缩放图片
	zoom := resize.Resize(width, height, img, resize.NearestNeighbor)
	buff := new(bytes.Buffer)
	if err = png.Encode(buff, zoom); err != nil {
		return
	}
	data = buff.Bytes()
	return
}
