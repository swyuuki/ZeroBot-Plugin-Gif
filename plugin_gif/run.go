package plugin_gif

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tdf1939/img"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

// FirstValueInList 判断正则匹配的第一个参数是否在列表中
func FirstValueInList(list []string) zero.Rule {
	return func(ctx *zero.Ctx) bool {
		first := ctx.State["regex_matched"].([]string)[1]
		for i := range list {
			if first == list[i] {
				return true
			}
		}
		return false
	}
}

var (
	table = map[string]func(...string) ([]byte, error){
		// 动图
		"摸": mo,
		"搓": cuo,
		"敲": qiao,
		"吃": chi,
		"蹭": ceng,
		"啃": ken,
		"拍": pai,
		"冲": chong,
		"丢": diu,
		// 静图
		"爬":  pa,
		"撕":  si,
		"一直": yizhi,
		// 简单操作
		"上翻": other(func(im *img.Dc) *img.Dc { return im.FlipV() }),
		"下翻": other(func(im *img.Dc) *img.Dc { return im.FlipV() }),
		"左翻": other(func(im *img.Dc) *img.Dc { return im.FlipH() }),
		"右翻": other(func(im *img.Dc) *img.Dc { return im.FlipH() }),
		"反色": other(func(im *img.Dc) *img.Dc { return im.Invert() }),
		"灰度": other(func(im *img.Dc) *img.Dc { return im.Grayscale() }),
		"浮雕": other(func(im *img.Dc) *img.Dc { return im.Convolve3x3() }),
		"打码": other(func(im *img.Dc) *img.Dc { return im.Blur(10) }),
		"负片": other(func(im *img.Dc) *img.Dc { return im.Invert().Grayscale() }),
	}
)

// HasImageFunc 判断开头是否存在对应函数
func HasImageFunc() zero.Rule {
	return func(ctx *zero.Ctx) bool {
		for k := range table {
			if strings.HasPrefix(ctx.Event.RawMessage, k) {
				ctx.State["image_type"] = k
				return true
			}
		}
		return false
	}
}

// 判断字符串为纯数字
func isNum(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// CatchAssets 获取素材
func CatchAssets() zero.Rule {
	return func(ctx *zero.Ctx) bool {
		var assets []string
		for _, v := range ctx.Event.Message {
			switch {
			case v.Type == "at":
				assets = append(assets, "http://q4.qlogo.cn/g?b=qq&nk="+v.Data["qq"]+"&s=640")
			case v.Type == "image":
				assets = append(assets, "https://gchat.qpic.cn/gchatpic_new//--"+strings.ToUpper(v.Data["file"])+"/0")
			case v.Type == "text":
				temp := strings.TrimLeft(v.Data["text"], ctx.State["image_type"].(string))
				if isNum(temp) {
					assets = append(assets, "http://q4.qlogo.cn/g?b=qq&nk="+temp+"&s=640")
				}
			}
		}
		ctx.State["image_assets"] = assets
		return len(assets) > 0
	}
}

func init() { // 插件主体
	zero.On("message", HasImageFunc(), CatchAssets()).
		SetBlock(true).SetPriority(20).Handle(func(ctx *zero.Ctx) {
		handle := table[ctx.State["image_type"].(string)]
		assets := ctx.State["image_assets"].([]string)
		assets = append(assets, fmt.Sprintf("http://q4.qlogo.cn/g?b=qq&nk=%d&s=640", ctx.Event.UserID))
		data, err := handle(assets...)
		if err != nil {
			ctx.SendChain(message.Text("ERROR: ", err))
		}
		ctx.SendChain(message.Image("base64://" + string(data)))
	})
}
