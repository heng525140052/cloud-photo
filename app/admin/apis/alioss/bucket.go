package alioss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/dto"
	"go-admin/app/admin/models/alioss"
	"go-admin/tools/app"
)

// @Summary 创建Bucket
// @Description 获取JSON
// @Tags 阿里云储存
// @Accept multipart/form-data
// @Param name query string true "name"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/alioss/CreateBucket [post]
func CreateBucket(c *gin.Context) {

	fmt.Println(c)
	return

	// 创建存储空间（默认为标准存储类型），并设置存储空间的权限为公共读（默认为私有）。
	err := alioss.Client().CreateBucket("<yourBucketName1>", oss.ACL(oss.ACLPublicRead))
	if err != nil {
		app.Error(c, 200, err, "")
		return
	}
	app.OK(c, "", "上传成功")

}

// @Summary Buckets列表
// @Description 获取JSON
// @Tags 阿里云储存
// @Accept multipart/form-data
// @Param prefix query string true "prefix"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/alioss/BucketsList [post]
func ListBuckets(c *gin.Context) {

	name := c.DefaultPostForm("prefix", "")
	marker := oss.Marker(name)
	if name != "" {
		marker = oss.Prefix(name) //查询前缀
	}

	var outList []dto.BucketsListItemOutput //定义返回值

	lsRes, err := alioss.Client().ListBuckets(marker)
	if err != nil {
		app.Error(c, 200, err, "未查询到列表")
		return
	}

	for _, bucket := range lsRes.Buckets {
		outputItem := dto.BucketsListItemOutput{
			Name:         bucket.Name,
			Location:     bucket.Location,
			CreationDate: bucket.CreationDate,
			StorageClass: bucket.StorageClass,
		}

		outList = append(outList, outputItem)
	}

	app.OK(c, outList, "success")

}
