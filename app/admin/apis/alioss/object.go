package alioss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/dto"
	"go-admin/app/admin/models/alioss"
	"go-admin/tools/app"
)

// @Summary 拷贝object
// @Description 获取JSON
// @Tags 阿里云储存
// @Accept multipart/form-data
// @Param objectName query string true "objectName"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/alioss/isObjectExits [post]
func IsObjectExits(c *gin.Context){
	objectName := c.DefaultPostForm("objectName", "")

	// 判断文件是否存在。
	isExist, err := alioss.Bucket().IsObjectExist(objectName)
	if err != nil {
		app.Error(c, 2, err, "error")
		return
	}
	app.OK(c, isExist, "success")
}


// @Summary Object列表
// @Description 获取JSON
// @Tags 阿里云储存
// @Accept multipart/form-data
// @Param prefix query string true "prefix"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/alioss/ListObjects [post]
func ListObjects(c *gin.Context) {

	name := c.DefaultPostForm("prefix", "")

	MaxKeys := 20

	marker := oss.Marker(name)
	if name != "" {
		marker = oss.Prefix(name) //查询前缀
	}

	var outList []dto.ObjectListItemOutput //定义返回值

	lsRes, err := alioss.Bucket().ListObjects(oss.MaxKeys(MaxKeys), marker)
	if err != nil {
		app.Error(c, 2, err, "未查询到列表")
		return
	}

	// 打印列举文件，默认情况下一次返回20条记录。
	for _, object := range lsRes.Objects {

		outputItem := dto.ObjectListItemOutput{
			Key:          object.Key,
			Type:         object.Type,
			ETag:         object.ETag,
			StorageClass: object.StorageClass,
			LastModified: object.LastModified,
		}
		outList = append(outList, outputItem)
	}

	app.OK(c, outList, "success")

}


// @Summary 拷贝object
// @Description 获取JSON
// @Tags 阿里云储存
// @Accept multipart/form-data
// @Param objectName query string true "objectName"
// @Param destObjectName query string true "destObjectName"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/alioss/copyObject [post]
func CopyObject(c *gin.Context) {
	objectName := c.DefaultPostForm("objectName", "")
	destObjectName := c.DefaultPostForm("destObjectName","")

	// 拷贝文件到同一个存储空间的另一个文件。
	out, err := alioss.Bucket().CopyObject(objectName, destObjectName)
	if err != nil {
		app.Error(c, 2, err, "拷贝失败")
		return
	}
	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = alioss.Bucket().DeleteObject(objectName)
	if err != nil {
		app.Error(c, 2, err, "删除源文件失败")
		return
	}

	app.OK(c, out, "success")

}

// @Summary 拷贝object
// @Description 获取JSON
// @Tags 阿里云储存
// @Accept multipart/form-data
// @Param objectName query string true "objectName"
// @Param destObjectName query string true "destObjectName"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/alioss/objectDetailedMeta [post]
func ObjectDetailedMeta(c *gin.Context){
	objectName := c.DefaultPostForm("objectName", "")

	// 获取文件元信息。
	props, err := alioss.Bucket().GetObjectDetailedMeta(objectName)
	if err != nil {
		app.Error(c, 2, err, "查看文件信息失败")
		return
	}
	app.OK(c, props, "success")

}


// @Summary 生成下载链接
// @Description 获取JSON
// @Tags 阿里云储存
// @Accept multipart/form-data
// @Param objectName query string true "objectName"
// @Param destObjectName query string true "destObjectName"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/alioss/getObjectSignUrl [post]
func GetObjectSignUrl(c *gin.Context){
	objectName := c.DefaultPostForm("objectName", "")

	_, err := alioss.Bucket().IsObjectExist(objectName)
	if err != nil {
		app.Error(c, 2, err, "error")
		return
	}
	// 获取文件元信息。
	signedURL, err := alioss.Bucket().SignURL(objectName, oss.HTTPGet, 3600)
	if err != nil {
		app.Error(c, 2, err, "查看文件信息失败")
		return
	}
	app.OK(c, signedURL, "success")

}
