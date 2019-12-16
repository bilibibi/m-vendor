package merchant

import (
	"io"
	"os"
	"path/filepath"

	"github.com/charsunny/wechat/mp/core"
	"fmt"
	"io/ioutil"
)

// UploadImage 上传图片到微信服务器, 返回的图片url给其他场景使用, 比如图文消息, 卡卷, POI.
func UploadImage(clt *core.Client, imgFilePath string) (url string, err error) {
	file, err := os.Open(imgFilePath)
	if err != nil {
		return
	}
	defer file.Close()

	return UploadImageFromReader(clt, filepath.Base(imgFilePath), file)
}

// UploadImageFromReader 上传图片到微信服务器, 返回的图片url给其他场景使用, 比如图文消息, 卡卷, POI.
//  NOTE: 参数 filename 不是文件路径, 是 multipart/form-data 里面 filename 的值.
func UploadImageFromReader(clt *core.Client, filename string, reader io.Reader) (url string, err error) {
	incompleteURL := fmt.Sprintf("https://api.weixin.qq.com/merchant/common/upload_img?filename=%s&?access_token=",filename)

	var result struct {
		core.Error
		URL string `json:"image_url"`
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	if err = clt.PostData(incompleteURL, data, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	url = result.URL
	return
}
