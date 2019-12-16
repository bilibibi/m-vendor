package vod

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// AddEditingProject invokes the vod.AddEditingProject API synchronously
// api document: https://help.aliyun.com/api/vod/addeditingproject.html
func (client *Client) AddEditingProject(request *AddEditingProjectRequest) (response *AddEditingProjectResponse, err error) {
	response = CreateAddEditingProjectResponse()
	err = client.DoAction(request, response)
	return
}

// AddEditingProjectWithChan invokes the vod.AddEditingProject API asynchronously
// api document: https://help.aliyun.com/api/vod/addeditingproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddEditingProjectWithChan(request *AddEditingProjectRequest) (<-chan *AddEditingProjectResponse, <-chan error) {
	responseChan := make(chan *AddEditingProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddEditingProject(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// AddEditingProjectWithCallback invokes the vod.AddEditingProject API asynchronously
// api document: https://help.aliyun.com/api/vod/addeditingproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddEditingProjectWithCallback(request *AddEditingProjectRequest, callback func(response *AddEditingProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddEditingProjectResponse
		var err error
		defer close(result)
		response, err = client.AddEditingProject(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// AddEditingProjectRequest is the request struct for api AddEditingProject
type AddEditingProjectRequest struct {
	*requests.RpcRequest
	CoverURL             string `position:"Query" name:"CoverURL"`
	Division             string `position:"Query" name:"Division"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	Timeline             string `position:"Query" name:"Timeline"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
}

// AddEditingProjectResponse is the response struct for api AddEditingProject
type AddEditingProjectResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Project   Project `json:"Project" xml:"Project"`
}

// CreateAddEditingProjectRequest creates a request to invoke AddEditingProject API
func CreateAddEditingProjectRequest() (request *AddEditingProjectRequest) {
	request = &AddEditingProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "AddEditingProject", "vod", "openAPI")
	return
}

// CreateAddEditingProjectResponse creates a response to parse from AddEditingProject response
func CreateAddEditingProjectResponse() (response *AddEditingProjectResponse) {
	response = &AddEditingProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}