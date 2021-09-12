package plugin_gif

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type context struct {
	folder string
	user   string
	imgs   []string
}

func downloadBy(name string) string {
	_, err := os.Stat(`data/image/sucai/` + name)
	if err != nil {
		download(`https://ghproxy.com/https://raw.githubusercontent.com/tdf1939/sucai/main/`+name, `data/image/sucai/`+name)
	}
	return `data/image/sucai/` + name
}

// 新的上下文
func newContext(user int64) context {
	var c context
	c.user = `data/image/user/` + strconv.FormatInt(user, 10) + `/`
	os.MkdirAll(c.user, 0755)
	c.imgs = make([]string, 2)
	c.imgs[0] = c.user + "yuan0.gif"
	c.imgs[1] = c.user + "yuan1.gif"
	return c
}

// 下载图片
func download(url, dlpath string) {
	// 创建目录
	var List = strings.Split(dlpath, `/`)
	err := os.MkdirAll(strings.TrimSuffix(dlpath, List[len(List)-1]), 0755)
	if err != nil {
		panic(err)
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)
	//创建文件
	file, err := os.Create(dlpath)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
}
