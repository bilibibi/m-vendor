package alipay

func (this *Client)MiniVersionUpload(param MiniVersionUploadQuery) (results *MiniVersionUploadRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

func (this *Client)MiniVersionBuild(param MiniVersionBuildQuery) (results *MiniVersionBuildRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

func (this *Client)MiniVersionDelete(param MiniVersionDeleteQuery) (results *MiniVersionDeleteRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

func (this *Client)MiniVersionAuditApply(param MiniVersionAuditApplyQuery) (results *MiniVersionAuditApplyRsp, err error) {
	files := make(map[string]string)
	if param.FirstLicensePic != "" {
		files["first_license_pic"] = param.FirstLicensePic
	}
	if param.OutDoorPic != "" {
		files["out_door_pic"] = param.OutDoorPic
	}
	if len(files) > 0 {
		err = this.doRequestWithFile("POST", param, files, &results)
	} else {
		err = this.doRequest("POST", param, &results)
	}
	return results, err
}

func (this *Client)MiniVersionDetail(param MiniVersionDetailQuery) (results *MiniVersionDetailRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

func (this *Client)MiniVersionAuditedCancel(param MiniVersionAuditCancelQuery) (results *MiniVersionAuditCancelRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

func (this *Client)MiniExpirenceCreate(param MiniExperienceCreateQuery) (results *MiniExperienceCreateRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

func (this *Client)MiniExpirenceQrcode(param MiniExperienceQrcodeQuery) (results *MiniExperienceQrcodeRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

func (this *Client)MiniQrcodeCreate(param MiniQrcodeCreateQuery) (results *MiniQrcodeCreateRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

func (this *Client)MiniVersionOnline(param MiniVersionOnlineQuery) (results *MiniVersionOnlineRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

func (this *Client)MiniVersionOffline(param MiniVersionOfflineQuery) (results *MiniVersionOfflineRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

func (this *Client)MiniCategoryList(param MiniCategoryQuery) (results *MiniCategoryRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}