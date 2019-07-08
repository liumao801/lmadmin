package common

import (
	"github.com/astaxie/beego"
	"liumao801/lmadmin/functions"
	"mime/multipart"
	"path/filepath"
)

type UploadController struct {
	beego.Controller
}

type uploadError struct {
	Message		string	`json:"message"` // 图片上传的错误提示信息
}

// 图片上传结构返回结构体
type UploadResult struct {
	Url 			string	`json:"url"`
	ThumbnailUrl 	string 	`json:"thumbnailUrl,omitempty"`
	Uploaded 		bool	`json:"uploaded"`
	FileName		string	`json:"fileName"`
	Msg 			string	`json:"msg"`
	Error			uploadError	`json:"error"`
}
// 公共的上传方法，适应 ckeditor 上传
func (c *UploadController) CommonUpload() {
	var allfiles map[string][]*multipart.FileHeader = c.Ctx.Request.MultipartForm.File
	/*
	allfiles struct info:
	{
		"upload": [
			{
				"Filename": "xs1-1.png",
				"Header": {
					"Content-Disposition": [
						"form-data; name=\"upload\"; filename=\"xs1-1.png\""
					],
					"Content-Type": [
						"image/png"
					]
				},
				"Size": 53434
			}
		]
	}
	*/
	uped := &UploadResult{}
	var rel []*UploadResult

	for k, _ := range allfiles {
		file, err := functions.LmUpload2(&c.Controller, k)
		//file, err := common.LmUpload(&c.Controller, k)
		if err != nil {
			uped.Uploaded = false
			uped.Msg = err.Error()
			uped.Url = ""
		} else {
			uped.Uploaded = true
			uped.Url = file.Url
			uped.ThumbnailUrl = file.ThumbnailUrl
			uped.Msg = ""
			//uped.FileName = file.Name // 上传图片的名称
			uped.FileName = filepath.Base(file.Url) // 上传后图片的名称
		}
		if c.GetString("refer") == "CKEDITOR" {
			c.Data["json"] = uped
			c.ServeJSON()
			c.StopRun()
		}
		rel = append(rel, uped)
	}

	c.Data["json"] = rel
	c.ServeJSON()
}

