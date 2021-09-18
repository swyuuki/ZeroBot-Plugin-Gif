package plugin_gif

import (
	"strconv"
	"strings"

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

var a1 = []string{"搓", "冲", "摸", "拍", "抬", "次", "丢", "吃", "敲", "透", "啃", "亲", "蹭",
	"爬", "撕", "一直",
	"灰度", "上翻", "下翻", "左翻", "右翻", "反色", "浮雕", "打码", "负片"}

// var a2 = []string{"旋转", "变形"}

func init() { // 插件主体

	zero.OnRegex(`^(` + strings.Join(a1, "|") + `)\D*?(\[CQ:(image,file=([0-9a-zA-Z]+).*?|at.+?(\d{5,11}))\].*|(\d+))$`).
		SetBlock(true).SetPriority(20).Handle(func(ctx *zero.Ctx) {
		NewPath(ctx.Event.UserID)
		List := ctx.State["regex_matched"].([]string)
		YuanTu(List[4]+List[5]+List[6], strconv.FormatInt(ctx.Event.UserID, 10))
		var picurl string
		switch {
		case List[1] == "爬":
			picurl = Ypath.Pa()
		case List[1] == "摸":
			picurl = Ypath.Mo()
		case List[1] == "吃":
			picurl = Ypath.Chi()
		case List[1] == "啃":
			picurl = Ypath.Ken()
		case List[1] == "蹭":
			picurl = Ypath.Ceng()
		case List[1] == "敲":
			picurl = Ypath.Qiao()
		case List[1] == "搓":
			picurl = Ypath.Cuo()
		case List[1] == "拍":
			picurl = Ypath.Pai()
		case List[1] == "丢":
			picurl = Ypath.Diu()
		case List[1] == "撕":
			picurl = Ypath.Si()
		case List[1] == "冲":
			picurl = Ypath.Chong()
		case List[1] == "一直":
			picurl = Ypath.YiZhi()
		default:
			picurl = Ypath.Other(List[1]) //"灰度", "上翻", "下翻", "左翻", "右翻", "反色", "倒放", "浮雕", "打码", "负片"
		}
		ctx.SendChain(
			//发送图片
			message.Image(picurl),
		)
	})
}


func YuanTu(s ...string) {
	for i, v := range s {
		_, err := strconv.Atoi(v)
		if err != nil {
			Download("https://gchat.qpic.cn/gchatpic_new//--"+strings.ToUpper(v)+"/0", Ypath.User+"yuan"+strconv.Itoa(i)+".gif")
		} else {
			Download("http://q4.qlogo.cn/g?b=qq&nk="+v+"&s=640", Ypath.User+"yuan"+strconv.Itoa(i)+".gif")
		}
	}
}
