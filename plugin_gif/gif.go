package plugin_gif

import (
	"image"

	"github.com/tdf1939/img"
	// "bot/img" // 基础词库
)

//摸
func (cc Paths) Mo() string {
	tou := img.ImDc(cc.Pngs[0], 0, 0).Circle(0).Im
	mo := []*image.NRGBA{
		img.ImDc(DLSc(`mo/0.png`), 0, 0).DstOver(tou, 80, 80, 32, 32).Im,
		img.ImDc(DLSc(`mo/1.png`), 0, 0).DstOver(tou, 70, 90, 42, 22).Im,
		img.ImDc(DLSc(`mo/2.png`), 0, 0).DstOver(tou, 75, 85, 37, 27).Im,
		img.ImDc(DLSc(`mo/3.png`), 0, 0).DstOver(tou, 85, 75, 27, 37).Im,
		img.ImDc(DLSc(`mo/4.png`), 0, 0).DstOver(tou, 90, 70, 22, 42).Im,
	}
	img.SaveGif(img.AndGif(1, mo), cc.User+`摸.gif`)
	return img.SGpic(cc.User + `摸.gif`)
}

//摸
func (cc Paths) Cuo() string {
	tou := img.ImDc(cc.Pngs[0], 110, 110).Circle(0).Im
	m1 := img.Rotate(tou, 72, 0, 0)
	m2 := img.Rotate(tou, 144, 0, 0)
	m3 := img.Rotate(tou, 216, 0, 0)
	m4 := img.Rotate(tou, 288, 0, 0)
	cuo := []*image.NRGBA{
		img.ImDc(DLSc(`cuo/0.png`), 0, 0).DstOverC(tou, 0, 0, 75, 130).Im,
		img.ImDc(DLSc(`cuo/1.png`), 0, 0).DstOverC(m1.Im, 0, 0, 75, 130).Im,
		img.ImDc(DLSc(`cuo/2.png`), 0, 0).DstOverC(m2.Im, 0, 0, 75, 130).Im,
		img.ImDc(DLSc(`cuo/3.png`), 0, 0).DstOverC(m3.Im, 0, 0, 75, 130).Im,
		img.ImDc(DLSc(`cuo/4.png`), 0, 0).DstOverC(m4.Im, 0, 0, 75, 130).Im,
	}
	img.SaveGif(img.AndGif(5, cuo), cc.User+`搓.gif`)
	return img.SGpic(cc.User + `搓.gif`)
}
//敲
func (cc Paths) Qiao() string {
	tou := img.ImDc(cc.Pngs[0], 40, 40).Circle(0).Im
	qiao := []*image.NRGBA{
		img.ImDc(DLSc(`qiao/0.png`), 0, 0).Over(tou, 40, 33, 57, 52).Im,
		img.ImDc(DLSc(`qiao/1.png`), 0, 0).Over(tou, 38, 36, 58, 50).Im,
	}
	img.SaveGif(img.AndGif(1, qiao), cc.User+`敲.gif`)
	return img.SGpic(cc.User + `敲.gif`)

}

//吃
func (cc Paths) Chi() string {
	tou := img.ImDc(cc.Pngs[0], 32, 32).Im
	chi := []*image.NRGBA{
		img.ImDc(DLSc(`chi/0.png`), 0, 0).DstOver(tou, 0, 0, 1, 38).Im,
		img.ImDc(DLSc(`chi/1.png`), 0, 0).DstOver(tou, 0, 0, 1, 38).Im,
		img.ImDc(DLSc(`chi/2.png`), 0, 0).DstOver(tou, 0, 0, 1, 38).Im,
	}
	img.SaveGif(img.AndGif(1, chi), cc.User+`吃.gif`)
	return img.SGpic(cc.User + `吃.gif`)

}

//拍
func (cc Paths) Pai() string {
	tou := img.ImDc(cc.Pngs[0], 30, 30).Circle(0).Im
	pai := []*image.NRGBA{
		img.ImDc(DLSc(`pai/0.png`), 0, 0).Over(tou, 0, 0, 1, 47).Im,
		img.ImDc(DLSc(`pai/1.png`), 0, 0).Over(tou, 0, 0, 1, 67).Im,
	}
	img.SaveGif(img.AndGif(1, pai), cc.User+`拍.gif`)
	return img.SGpic(cc.User + `拍.gif`)

}

//冲
func (cc Paths) Chong() string {
	tou := img.ImDc(cc.Pngs[0], 0, 0).Circle(0).Im
	chong := []*image.NRGBA{
		img.ImDc(DLSc(`xqe/0.png`), 0, 0).Over(tou, 30, 30, 15, 53).Im,
		img.ImDc(DLSc(`xqe/1.png`), 0, 0).Over(tou, 30, 30, 40, 53).Im,
	}
	img.SaveGif(img.AndGif(1, chong), cc.User+`冲.gif`)
	return img.SGpic(cc.User + `冲.gif`)

}

//丢
func (cc Paths) Diu() string {
	tou := img.ImDc(cc.Pngs[0], 0, 0).Circle(0).Im
	diu := []*image.NRGBA{
		img.ImDc(DLSc(`diu/0.png`), 0, 0).Over(tou, 32, 32, 108, 36).Im,
		img.ImDc(DLSc(`diu/1.png`), 0, 0).Over(tou, 32, 32, 122, 36).Im,
		img.ImDc(DLSc(`diu/2.png`), 0, 0).Im,
		img.ImDc(DLSc(`diu/3.png`), 0, 0).Over(tou, 123, 123, 19, 129).Im,
		img.ImDc(DLSc(`diu/4.png`), 0, 0).Over(tou, 185, 185, -50, 200).Over(tou, 33, 33, 289, 70).Im,
		img.ImDc(DLSc(`diu/5.png`), 0, 0).Over(tou, 32, 32, 280, 73).Im,
		img.ImDc(DLSc(`diu/6.png`), 0, 0).Over(tou, 35, 35, 259, 31).Im,
		img.ImDc(DLSc(`diu/7.png`), 0, 0).Over(tou, 175, 175, -50, 220).Im,
	}
	img.SaveGif(img.AndGif(7, diu), cc.User+`丢.gif`)
	return img.SGpic(cc.User + `丢.gif`)
}
