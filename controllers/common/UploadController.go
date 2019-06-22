package common

import (
	"github.com/astaxie/beego"
	"liumao801/lmadmin/functions"
	"mime/multipart"
)

type UploadController struct {
	beego.Controller
}

type UploadResult struct {
	Url 		string	`json:"url"`
	Uploaded 	bool	`json:"uploaded"`
	Msg 		string	`json:"msg"`
}
// 公共的上传方法
func (c *UploadController) CommonUpload() {
	var allfiles map[string][]*multipart.FileHeader = c.Ctx.Request.MultipartForm.File

	uped := &UploadResult{}
	var rel []*UploadResult

	for k, _ := range allfiles {
		file, err := functions.LmUpload(&c.Controller, k)
		//file, err := common.LmUpload(&c.Controller, k)
		if err != "" {
			uped.Uploaded = false
			uped.Msg = err
			uped.Url = ""
		} else {
			uped.Uploaded = true
			uped.Url = file
			uped.Msg = ""
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
/*
type Sizer interface {
	Size() int64
}

const (
	LOCAL_FILE_DIR		= "static/upload/"
	MIN_FILE_SIZE 		= 1			// bytes
	MAX_FILE_SIZE 		= 5000000	// bytes
	IMAGE_TYPES 		= "(jpg|gif|p?jpeg|(x-)?png)"
	ACCEPT_FILE_TYPES = IMAGE_TYPES
	EXPIRATION_TIME 	= 300	// seconds
	THUMBNAIL_PARAM 	= "=s80"
)

var (
	imageTypes 		= regexp.MustCompile(IMAGE_TYPES)
	acceptFileTypes = regexp.MustCompile(ACCEPT_FILE_TYPES)
)

type FileInfo struct {
	Url 		 string `json:"url,omitempty"`
	ThumbnailUrl string `json:"thumbnailUrl,omitempty"`
	Name 		 string `json:"name"`
	Type 		 string `json:"type"`
	Size 		 int64  `json:"size"`
	Error 		 string `json:"error,omitempty"`
	DeleteUrl 	 string `json:"deleteUrl,omitempty"`
	DeleteType 	 string `json:"deleteType,omitempty"`
}
// 检测文件类型是否合法
func (fi *FileInfo) ValidateType() (valid bool) {
	if acceptFileTypes.MatchString(fi.Type) {
		return true
	}

	fi.Error = "Filetype not allowed"
	return false
}
// 检查文件大小是否合法
func (fi *FileInfo) ValidateSize() (valid bool) {
	if fi.Size < MIN_FILE_SIZE {
		fi.Error = "File is too small"
	} else if fi.Size > MAX_FILE_SIZE {
		fi.Error = "File is too large"
	} else {
		return true
	}
	return false
}
// 检测是否合法
func (fi *FileInfo) check(err error)  {
	if err != nil {
		panic(err)
	}
}
func (fi *FileInfo) escape(s string) string {
	return strings.Replace(url.QueryEscape(s), "+", "%20", -1)
}
func (fi *FileInfo) getFormValue(p *multipart.Part) string{
	var b bytes.Buffer
	io.CopyN(&b, p, int64((1<<20)))	// Copy max: 1 MiB
	return b.String()
}
// 截取字符串
func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
// 获取父级 目录
func getParentDirectory(dir string) string {
	return substr(dir, 0, strings.LastIndex(dir, "/"))
}
// 获取当前目录
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		utils.LogError(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
// 获取文件后缀
func (fi *FileInfo) fileExt() {
	ext := path.Ext(fi.Name)
	fi.Type = ext
}

// 文件上传的公共保存方法
// 返回文件路径和错误信息
// 返回第一个参数是文件；第二个参数是错误提示，""代表没有错误
func LmUpload(c *beego.Controller,fileName string) (string, string) {
	file, fileHeader, err := c.GetFile(fileName)
	if err != nil {
		return "", "文件获取失败"
	}
	defer file.Close()

	fi := &FileInfo{
		Name: fileHeader.Filename,
	}
	// 获取文件类型
	fi.fileExt()

	if !fi.ValidateType() {
		return "", "文件类型错误"
	}

	if sizeInterface, ok := file.(Sizer); ok {
		fi.Size = sizeInterface.Size()
		if !fi.ValidateSize() {
			return "", fi.Error
		}
	} else {
		return "", "文件大小获取失败"
	}
	now := time.Now()
	ctrlName, _ := c.GetControllerAndAction()
	dirPath := LOCAL_FILE_DIR + strings.ToLower(ctrlName[0 : len(ctrlName)-10]) + "/" + now.Format("2006-01") + "/" + now.Format("02")
	fileExt := strings.TrimLeft(fi.Type, ".")
	fileSaveName := fmt.Sprintf("%s_%d.%s", controllers.RandCode(5, 3), now.Unix(), fileExt)
	filePath := fmt.Sprintf("%s/%s", dirPath, fileSaveName)

	beego.Info("filePath===========",filePath)
	if !controllers.IsDir(dirPath) {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return "", "文件夹“" + dirPath + "”创建失败"
		}
	}
	// 保存位置在 static/upload， 没有文件夹要先创建
	c.SaveToFile(fileName, filePath)
	return "/" + filePath, ""
}*/