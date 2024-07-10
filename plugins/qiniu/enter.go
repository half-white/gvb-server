package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"gvb_server/config"
	"gvb_server/global"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// 获取上传的Token
func getToken(q config.Qiniu) string {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

// 获取上传的配置
func getCfg(q config.Qiniu) storage.Config {
	cfg := storage.Config{}
	//空间对应的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	//是否使用https域名
	cfg.UseHTTPS = false
	//上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	return cfg
}

// UploadImage 上传图片 文件数组，前缀
func UploadImage(data []byte, imageName string, prefix string) (filepath string, err error) {
	if !global.Config.Qiniu.Enable {
		return "", errors.New("请启用七牛云")
	}
	q := global.Config.Qiniu
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("请设定accessKey和secretKey")
	}
	if float64(len(data))/float64(1024*1024) > q.Size {
		return "", errors.New("图片大小超出限定值")
	}
	upToken := getToken(q)
	cfg := getCfg(q)

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))

	//获取当前时间
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf("%s/%s_%s", prefix, now, imageName)

	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", errors.New("未知错误")
	}
	return fmt.Sprintf("%s%s", q.CDN, ret.Key), nil
}
