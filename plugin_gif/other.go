package plugin_gif

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

//获取素材
func DLSc(nm string) string {
	_, err := os.Stat(`data/image/sucai/` + nm)
	if err != nil {
		Download(`https://ghproxy.com/https://raw.githubusercontent.com/tdf1939/sucai/main/`+nm, `data/image/sucai/`+nm)
	}
	return `data/image/sucai/` + nm
}

//下载图片
func Download(url, dlpath string) {
	//创建目录
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
	io.Copy(writer, reader)
}
