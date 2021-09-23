package plugin_gif

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//所有路径
type Paths struct {
	User string   //用户目录路径
	Pngs []string //待处理图片路径
}

//更新用户目录
func NewPath(user int64) *Paths {
	a := `data/image/user/` + strconv.FormatInt(user, 10) + `/`
	os.MkdirAll(a, 0777)
	return &Paths{
		User: a,
		Pngs: []string{},
	}
}

//添加图片
func (cc *Paths) AddIm(s ...string) {
	for _, v := range s {
		_, err := strconv.Atoi(v)
		n := cc.User + strings.ToUpper(v) + `.png`
		if err != nil {
			Download("https://gchat.qpic.cn/gchatpic_new//--"+strings.ToUpper(v)+"/0", n)
		} else {
			Download("http://q4.qlogo.cn/g?b=qq&nk="+v+"&s=640", n)
		}
		cc.Pngs = append(cc.Pngs, n)
	}
}

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

//图片转Base64
func Base64(p string) string {
	//读原图片
	ff, _ := os.Open(p)
	defer ff.Close()
	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Read(sourcebuffer)
	//base64压缩
	return base64.StdEncoding.EncodeToString(sourcebuffer[:n])
}
