package smartag

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

// SmartAccessGateway is a nested struct in smartag response
type SmartAccessGateway struct {
	Name                  string      `json:"Name" xml:"Name"`
	MaxBandwidth          string      `json:"MaxBandwidth" xml:"MaxBandwidth"`
	SoftwareVersion       string      `json:"SoftwareVersion" xml:"SoftwareVersion"`
	HardwareVersion       string      `json:"HardwareVersion" xml:"HardwareVersion"`
	CreateTime            int         `json:"CreateTime" xml:"CreateTime"`
	EndTime               int         `json:"EndTime" xml:"EndTime"`
	City                  string      `json:"City" xml:"City"`
	SerialNumber          string      `json:"SerialNumber" xml:"SerialNumber"`
	AssociatedCcnId       string      `json:"AssociatedCcnId" xml:"AssociatedCcnId"`
	State                 string      `json:"State" xml:"State"`
	AclIds                string      `json:"AclIds" xml:"AclIds"`
	InstanceId            string      `json:"InstanceId" xml:"InstanceId"`
	SecurityLockThreshold int         `json:"SecurityLockThreshold" xml:"SecurityLockThreshold"`
	UserCount             int         `json:"UserCount" xml:"UserCount"`
	Status                string      `json:"Status" xml:"Status"`
	CidrBlock             string      `json:"CidrBlock" xml:"CidrBlock"`
	AssociatedCcnName     string      `json:"AssociatedCcnName" xml:"AssociatedCcnName"`
	Description           string      `json:"Description" xml:"Description"`
	SmartAGId             string      `json:"SmartAGId" xml:"SmartAGId"`
	DataPlan              int         `json:"DataPlan" xml:"DataPlan"`
	SnatEntries           SnatEntries `json:"SnatEntries" xml:"SnatEntries"`
}
