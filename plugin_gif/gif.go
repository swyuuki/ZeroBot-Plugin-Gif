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
		img.ImDc(cc.Sc+`mo/0.png`, 0, 0).DstOver(tou, 80, 80, 32, 32).Im,
		img.ImDc(cc.Sc+`mo/1.png`, 0, 0).DstOver(tou, 70, 90, 42, 22).Im,
		img.ImDc(cc.Sc+`mo/2.png`, 0, 0).DstOver(tou, 75, 85, 37, 27).Im,
		img.ImDc(cc.Sc+`mo/3.png`, 0, 0).DstOver(tou, 85, 75, 27, 37).Im,
		img.ImDc(cc.Sc+`mo/4.png`, 0, 0).DstOver(tou, 90, 70, 22, 42).Im,
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
		img.ImDc(cc.Sc+`cuo/0.png`, 0, 0).DstOverC(tou, 0, 0, 75, 130).Im,
		img.ImDc(cc.Sc+`cuo/1.png`, 0, 0).DstOverC(m1.Im, 0, 0, 75, 130).Im,
		img.ImDc(cc.Sc+`cuo/2.png`, 0, 0).DstOverC(m2.Im, 0, 0, 75, 130).Im,
		img.ImDc(cc.Sc+`cuo/3.png`, 0, 0).DstOverC(m3.Im, 0, 0, 75, 130).Im,
		img.ImDc(cc.Sc+`cuo/4.png`, 0, 0).DstOverC(m4.Im, 0, 0, 75, 130).Im,
	}
	img.SaveGif(img.AndGif(5, cuo), cc.User+`搓.gif`)
	return img.SGpic(cc.User + `搓.gif`)
}

//拍
func (cc Paths) Pai() string {
	tou := img.ImDc(cc.Pngs[0], 30, 30).Circle(0).Im
	pai := []*image.NRGBA{
		img.ImDc(cc.Sc+`pai/0.png`, 0, 0).Over(tou, 0, 0, 1, 47).Im,
		img.ImDc(cc.Sc+`pai/1.png`, 0, 0).Over(tou, 0, 0, 1, 67).Im,
	}
	img.SaveGif(img.AndGif(1, pai), cc.User+`拍.gif`)
	return img.SGpic(cc.User + `拍.gif`)

}

//冲
func (cc Paths) Chong() string {
	tou := img.ImDc(cc.Pngs[0], 0, 0).Circle(0).Im
	chong := []*image.NRGBA{
		img.ImDc(cc.Sc+`xqe/0.png`, 0, 0).Over(tou, 30, 30, 15, 53).Im,
		img.ImDc(cc.Sc+`xqe/1.png`, 0, 0).Over(tou, 30, 30, 40, 53).Im,
	}
	img.SaveGif(img.AndGif(1, chong), cc.User+`冲.gif`)
	return img.SGpic(cc.User + `冲.gif`)

}

//丢
func (cc Paths) Diu() string {
	tou := img.ImDc(cc.Pngs[0], 0, 0).Circle(0).Im
	diu := []*image.NRGBA{
		img.ImDc(cc.Sc+`diu/0.png`, 0, 0).Over(tou, 63, 63, 216, 72).Im,
		img.ImDc(cc.Sc+`diu/1.png`, 0, 0).Over(tou, 65, 65, 244, 71).Im,
		img.ImDc(cc.Sc+`diu/2.png`, 0, 0).Im,
		img.ImDc(cc.Sc+`diu/3.png`, 0, 0).Over(tou, 245, 245, 38, 257).Im,
		img.ImDc(cc.Sc+`diu/4.png`, 0, 0).Over(tou, 370, 370, -100, 400).Over(tou, 66, 66, 578, 139).Im,
		img.ImDc(cc.Sc+`diu/5.png`, 0, 0).Over(tou, 64, 64, 560, 145).Im,
		img.ImDc(cc.Sc+`diu/6.png`, 0, 0).Over(tou, 70, 70, 517, 62).Im,
		img.ImDc(cc.Sc+`diu/7.png`, 0, 0).Over(tou, 350, 350, -100, 440).Im,
	}
	img.SaveGif(img.AndGif(1, diu), cc.User+`丢.gif`)
	return img.SGpic(cc.User + `丢.gif`)
}
