package image_ser

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/plugins/qiniu"
	"gvb_server/utils"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

var (
	WhiteImageList = []string{
		"jpg",
		"jpeg",
		"png",
		"ico",
		"tiff",
		"gif",
		"svg",
		"webg",
	}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  //文件名
	IsSuccess bool   `json:"is_success"` //是否上传成功
	Msg       string `json:"msg"`        //消息
}

// 文件上传的方法
func (ImageService) ImaggUploadService(file *multipart.FileHeader) (res FileUploadResponse) {

	fileName := file.Filename
	basepath := global.Config.Upload.Path
	filePath := path.Join(basepath, file.Filename)
	res.FileName = filePath

	//文件白名单判断
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = "非法文件"
		return
	}

	//判断大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过设定大小，当前大小为:%.2f,设定大小为:%dMB", size, global.Config.Upload.Size)
		return
	}

	//读取文件内容 hash
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteData)
	//需要在数据库中查询这个图片是否存在
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil {
		//找到了
		res.Msg = "图片已经存在"
		res.FileName = bannerModel.Path
		return
	}

	fileType := ctype.Local
	res.Msg = "图片上传成功"
	res.IsSuccess = true
	//上传七牛云
	if global.Config.Qiniu.Enable {
		filePath, err = qiniu.UploadImage(byteData, fileName, global.Config.Qiniu.Prefix)
		if err != nil {
			global.Log.Error(err)
			res.Msg = err.Error()
			return
		}
		res.FileName = filePath
		res.Msg = "上传七牛成功"
		fileType = ctype.QiNiu
	}

	//图片入库
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: fileType,
	})
	return
}
