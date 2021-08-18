package main

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/tdf1939/ZeroBot-Plugin-Gif/plugin_gif"

	// 以下为内置依赖，勿动
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
)

var (
	contents = []string{
		"* OneBot + ZeroBot + Golang ",
		"* Version 1.1.0 - 2021-08-06 23:36:29 +0800 CST",
		"* Copyright © 2020 - 2021  Kanri, DawnNights, Fumiama, Suika",
		"* Project: https://github.com/FloatTech/ZeroBot-Plugin",
	}
	banner = strings.Join(contents, "\n")
)

func init() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[zero][%time%][%lvl%]: %msg% \n",
	})
	log.SetLevel(log.DebugLevel)
}

func main() {
	fmt.Print(
		"======================[ZeroBot-Plugin]======================",
		"\n", banner, "\n",
		"============================================================\n",
	) // 启动打印
	zero.Run(zero.Config{
		NickName:      []string{"椛椛", "ATRI", "atri", "亚托莉", "アトリ"},
		CommandPrefix: "/",

		// SuperUsers 某些功能需要主人权限，可通过以下两种方式修改
		// []string{}：通过代码写死的方式添加主人账号
		// os.Args[1:]：通过命令行参数的方式添加主人账号
		SuperUsers: append([]string{"12345678", "87654321"}, os.Args[1:]...),

		Driver: []zero.Driver{
			&driver.WSClient{
				// OneBot 正向WS 默认使用 6700 端口
				Url:         "ws://yun.yuuuki.top:16700",
				AccessToken: "124816",
			},
		},
	})

	// 帮助
	zero.OnFullMatchGroup([]string{"help", "/help", ".help", "菜单", "帮助"}, zero.OnlyToMe).SetBlock(false).SetPriority(999).
		Handle(func(ctx *zero.Ctx) {
			ctx.Send(banner)
		})
	select {}
}
