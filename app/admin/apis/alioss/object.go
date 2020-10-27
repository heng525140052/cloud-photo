package alioss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/dto"
	"go-admin/app/admin/models/alioss"
	"go-admin/tools/app"
)

// @Summary Object列表
// @Description 获取JSON
// @Tags 阿里云储存
// @Accept multipart/form-data
// @Param prefix query string true "prefix"
// @Param page_size query string true "page_size"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/alioss/BucketsList [post]
func ListObjects(c *gin.Context) {

	name := c.DefaultPostForm("prefix", "")
	marker := oss.Marker("12345.jpg")

	if name != "" {
		marker = oss.Prefix(name) //查询前缀
	}

	var outList []dto.ObjectListItemOutput //定义返回值

	lsRes, err := alioss.Bucket().ListObjects(oss.MaxKeys(2),marker)
	if err != nil {
		app.Error(c, 2, err, "未查询到列表")
		return
	}

	// 打印列举文件，默认情况下一次返回100条记录。
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
