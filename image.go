package gomini

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"io"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	imgBase64Reg = regexp.MustCompile(`src=\"data:image\/([a-zA-Z]*);base64,(.*?)\"`)
	// 支持的图片后缀名
	supportImageExtNames = []string{".jpg", ".jpeg", ".png", ".ico", ".svg", ".bmp", ".gif"}
)

// 检查文件扩展名是否是图片
func IsImage(extName string) bool {
	for i := 0; i < len(supportImageExtNames); i++ {
		if supportImageExtNames[i] == extName {
			return true
		}
	}
	return false
}

// 转换base64为图片链接
func ConvertBase64ToImage(content string, imagePath, prefixURI string) (string, error) {
	if !IsExist(imagePath) {
		err := Mkdir(imagePath)
		if err != nil {
			return "", err
		}
	}
	// 0: src="data:image/jpeg;base64,2/9j/4AAQSkZJRgABAQAAAQABAAD/4Q+">
	// 1: jpeg
	// 2: 2/9j/4AAQSkZJRgABAQAAAQABAAD/4Q+
	all := imgBase64Reg.FindAllStringSubmatch(content, -1)
	for _, v := range all {
		if len(v) == 3 {
			hash := md5.New()
			io.Copy(hash, bytes.NewBufferString(v[2]))
			filename := hex.EncodeToString(hash.Sum(nil)) + "." + v[1]
			toFilename := filepath.Join(imagePath, filename)
			if !IsExist(toFilename) {
				dst, err := base64.StdEncoding.DecodeString(v[2])
				if err != nil {
					return "", err
				}
				err = ioutil.WriteFile(toFilename, dst, 0644)
				if err != nil {
					return "", err
				}
			}
			content = strings.Replace(content, v[0], `src="`+prefixURI+`/`+filename+`"`, -1)
		}
	}
	return content, nil
}
