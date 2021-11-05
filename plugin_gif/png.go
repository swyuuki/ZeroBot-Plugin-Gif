package plugin_gif

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"math/rand"
	"strconv"
	"time"

	"github.com/tdf1939/img"
)

// PNG 转 base64
func encodePNG(im image.Image) ([]byte, error) {
	buffer := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, buffer)
	err := png.Encode(encoder, im)
	if err != nil {
		return nil, err
	}
	encoder.Close()
	return buffer.Bytes(), nil
}

// 爬
func pa(pngs ...string) ([]byte, error) {
	tou := img.ImDc(pngs[0], 0, 0).Circle(0).Im
	// 随机爬图序号
	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(60) + 1
	dc := img.ImDc(DLSc("pa/"+strconv.Itoa(rand)+".png"), 0, 0).
		DstOver(tou, 100, 100, 0, 400).Im
	return encodePNG(dc)
}

// 撕
func si(pngs ...string) ([]byte, error) {
	tou := img.ImDc(pngs[0], 0, 0).Im
	im1 := img.Rotate(tou, 20, 380, 380)
	im2 := img.Rotate(tou, -12, 380, 380)
	dc := img.ImDc(DLSc("si/0.png"), 0, 0).
		DstOver(im1.Im, im1.W, im1.H, -3, 370).
		DstOver(im2.Im, im2.W, im2.H, 653, 310).Im
	return encodePNG(dc)
}

// 一直
func yizhi(pngs ...string) ([]byte, error) {
	tou := img.ImsDc(pngs[0], 0, 0)
	var dc []*image.NRGBA
	for _, v := range tou {
		dc = append(dc, img.ImDc(DLSc("xiao/0.png"), 0, 0).
			Over(v, 249, 249, 0, 0).
			Over(v, 47, 47, 140, 250).Clone().Im,
		)
	}
	return encodeGIF(img.AndGif(1, dc))
}

// TODO 关注 必须下载素材到指定位置才可用
func guanZhu(txt1, txt2 string, pngs ...string) ([]byte, error) {
	// 加载图片
	im := img.ImDc(pngs[0], 0, 0)
	//叠加图片
	dst := img.ImDc(DLSc("guanzhu/gg.png"), 0, 0).
		DstOver(im.Im, 155, 155, 45, 45)
	dc := dst.Text(DLSc("font/3.ttf"), 38, img.Black, 210, 110, txt1).
		Text(DLSc("font/3.ttf"), 36, img.Grey, 210, 170, txt2).Im
	return encodePNG(dc)
}

// 简单操作
func other(handle func(*img.Dc) *img.Dc) func(pngs ...string) ([]byte, error) {
	return func(pngs ...string) ([]byte, error) {
		im := img.ImDc(pngs[0], 0, 0)
		return encodePNG(handle(im).Im)
	}
}
