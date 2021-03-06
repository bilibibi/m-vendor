package baas

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

// GetTemplates invokes the baas.GetTemplates API synchronously
// api document: https://help.aliyun.com/api/baas/gettemplates.html
func (client *Client) GetTemplates(request *GetTemplatesRequest) (response *GetTemplatesResponse, err error) {
	response = CreateGetTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// GetTemplatesWithChan invokes the baas.GetTemplates API asynchronously
// api document: https://help.aliyun.com/api/baas/gettemplates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTemplatesWithChan(request *GetTemplatesRequest) (<-chan *GetTemplatesResponse, <-chan error) {
	responseChan := make(chan *GetTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTemplates(request)
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

// GetTemplatesWithCallback invokes the baas.GetTemplates API asynchronously
// api document: https://help.aliyun.com/api/baas/gettemplates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTemplatesWithCallback(request *GetTemplatesRequest, callback func(response *GetTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTemplatesResponse
		var err error
		defer close(result)
		response, err = client.GetTemplates(request)
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

// GetTemplatesRequest is the request struct for api GetTemplates
type GetTemplatesRequest struct {
	*requests.RpcRequest
}

// GetTemplatesResponse is the response struct for api GetTemplates
type GetTemplatesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateGetTemplatesRequest creates a request to invoke GetTemplates API
func CreateGetTemplatesRequest() (request *GetTemplatesRequest) {
	request = &GetTemplatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-07-31", "GetTemplates", "", "")
	return
}

// CreateGetTemplatesResponse creates a response to parse from GetTemplates response
func CreateGetTemplatesResponse() (response *GetTemplatesResponse) {
	response = &GetTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
