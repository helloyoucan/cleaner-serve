package logic

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"math/rand"
	"mime/multipart"
	"strconv"
	"time"
)

func UploadFileToQiNiu(file *multipart.FileHeader)(url string,err error) {
	accessKey := "m4_wAcs1jktM6FSWAUq3ZN2nzloL4uM8oNJbd4Jo"
	secretKey := "wZaix0Ba8Cs3EaopkISKgOt7J6NvvtjLsw2iO7Cc"
	bucket := "cleaner-serve"
	domain:="http://qurxcrn0l.hn-bkt.clouddn.com/" // 临时测试域名，有效期一个月（2021-06-16）
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{}
	newFileName := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000) + "-" + file.Filename
	data, _ := file.Open()
	err = formUploader.Put(context.Background(), &ret, upToken, newFileName, data, file.Size, &putExtra)
	if err != nil {
		return "" ,err
	}
	url = storage.MakePublicURL(domain, ret.Key)
	return
}
