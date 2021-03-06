package slb

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

// SetLoadBalancerHTTPSListenerAttribute invokes the slb.SetLoadBalancerHTTPSListenerAttribute API synchronously
// api document: https://help.aliyun.com/api/slb/setloadbalancerhttpslistenerattribute.html
func (client *Client) SetLoadBalancerHTTPSListenerAttribute(request *SetLoadBalancerHTTPSListenerAttributeRequest) (response *SetLoadBalancerHTTPSListenerAttributeResponse, err error) {
	response = CreateSetLoadBalancerHTTPSListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// SetLoadBalancerHTTPSListenerAttributeWithChan invokes the slb.SetLoadBalancerHTTPSListenerAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/setloadbalancerhttpslistenerattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetLoadBalancerHTTPSListenerAttributeWithChan(request *SetLoadBalancerHTTPSListenerAttributeRequest) (<-chan *SetLoadBalancerHTTPSListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *SetLoadBalancerHTTPSListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLoadBalancerHTTPSListenerAttribute(request)
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

// SetLoadBalancerHTTPSListenerAttributeWithCallback invokes the slb.SetLoadBalancerHTTPSListenerAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/setloadbalancerhttpslistenerattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetLoadBalancerHTTPSListenerAttributeWithCallback(request *SetLoadBalancerHTTPSListenerAttributeRequest, callback func(response *SetLoadBalancerHTTPSListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLoadBalancerHTTPSListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.SetLoadBalancerHTTPSListenerAttribute(request)
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

// SetLoadBalancerHTTPSListenerAttributeRequest is the request struct for api SetLoadBalancerHTTPSListenerAttribute
type SetLoadBalancerHTTPSListenerAttributeRequest struct {
	*requests.RpcRequest
	AccessKeyId            string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	HealthCheckTimeout     requests.Integer `position:"Query" name:"HealthCheckTimeout"`
	XForwardedFor          string           `position:"Query" name:"XForwardedFor"`
	HealthCheckURI         string           `position:"Query" name:"HealthCheckURI"`
	Description            string           `position:"Query" name:"Description"`
	UnhealthyThreshold     requests.Integer `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold       requests.Integer `position:"Query" name:"HealthyThreshold"`
	AclStatus              string           `position:"Query" name:"AclStatus"`
	Scheduler              string           `position:"Query" name:"Scheduler"`
	AclType                string           `position:"Query" name:"AclType"`
	HealthCheck            string           `position:"Query" name:"HealthCheck"`
	MaxConnection          requests.Integer `position:"Query" name:"MaxConnection"`
	EnableHttp2            string           `position:"Query" name:"EnableHttp2"`
	CookieTimeout          requests.Integer `position:"Query" name:"CookieTimeout"`
	StickySessionType      string           `position:"Query" name:"StickySessionType"`
	VpcIds                 string           `position:"Query" name:"VpcIds"`
	VServerGroupId         string           `position:"Query" name:"VServerGroupId"`
	AclId                  string           `position:"Query" name:"AclId"`
	ListenerPort           requests.Integer `position:"Query" name:"ListenerPort"`
	Cookie                 string           `position:"Query" name:"Cookie"`
	HealthCheckType        string           `position:"Query" name:"HealthCheckType"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth              requests.Integer `position:"Query" name:"Bandwidth"`
	StickySession          string           `position:"Query" name:"StickySession"`
	HealthCheckMethod      string           `position:"Query" name:"HealthCheckMethod"`
	HealthCheckDomain      string           `position:"Query" name:"HealthCheckDomain"`
	RequestTimeout         requests.Integer `position:"Query" name:"RequestTimeout"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	Gzip                   string           `position:"Query" name:"Gzip"`
	TLSCipherPolicy        string           `position:"Query" name:"TLSCipherPolicy"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	ServerCertificateId    string           `position:"Query" name:"ServerCertificateId"`
	CACertificateId        string           `position:"Query" name:"CACertificateId"`
	BackendProtocol        string           `position:"Query" name:"BackendProtocol"`
	Tags                   string           `position:"Query" name:"Tags"`
	IdleTimeout            requests.Integer `position:"Query" name:"IdleTimeout"`
	LoadBalancerId         string           `position:"Query" name:"LoadBalancerId"`
	XForwardedForSLBIP     string           `position:"Query" name:"XForwardedFor_SLBIP"`
	HealthCheckInterval    requests.Integer `position:"Query" name:"HealthCheckInterval"`
	XForwardedForProto     string           `position:"Query" name:"XForwardedFor_proto"`
	XForwardedForSLBID     string           `position:"Query" name:"XForwardedFor_SLBID"`
	HealthCheckConnectPort requests.Integer `position:"Query" name:"HealthCheckConnectPort"`
	HealthCheckHttpCode    string           `position:"Query" name:"HealthCheckHttpCode"`
	VServerGroup           string           `position:"Query" name:"VServerGroup"`
}

// SetLoadBalancerHTTPSListenerAttributeResponse is the response struct for api SetLoadBalancerHTTPSListenerAttribute
type SetLoadBalancerHTTPSListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetLoadBalancerHTTPSListenerAttributeRequest creates a request to invoke SetLoadBalancerHTTPSListenerAttribute API
func CreateSetLoadBalancerHTTPSListenerAttributeRequest() (request *SetLoadBalancerHTTPSListenerAttributeRequest) {
	request = &SetLoadBalancerHTTPSListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerHTTPSListenerAttribute", "slb", "openAPI")
	return
}

// CreateSetLoadBalancerHTTPSListenerAttributeResponse creates a response to parse from SetLoadBalancerHTTPSListenerAttribute response
func CreateSetLoadBalancerHTTPSListenerAttributeResponse() (response *SetLoadBalancerHTTPSListenerAttributeResponse) {
	response = &SetLoadBalancerHTTPSListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
