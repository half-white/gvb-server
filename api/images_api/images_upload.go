package images_api

import (
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/service/image_ser"
	"io/fs"
	"os"

	"github.com/gin-gonic/gin"
)

// 图片白名单
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

func (ImagesApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在文件", c)
		return
	}

	//判断路径是否存在
	_, err = os.ReadDir(global.Config.Upload.Path)
	if err != nil {
		err = os.MkdirAll(global.Config.Upload.Path, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	//不存在就创建
	var resList []image_ser.FileUploadResponse

	for _, file := range fileList {
		//上传文件
		serviceRes := service.ServiceApp.ImageService.ImaggUploadService(file)
		if serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}

		if !global.Config.Qiniu.Enable {
			//本地存储
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}
		resList = append(resList, serviceRes)
	}
	res.OkWithData(resList, c)
}
