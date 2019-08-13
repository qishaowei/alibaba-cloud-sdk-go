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

package eci

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeContainerGroups invokes the eci.DescribeContainerGroups API synchronously
// api document: https://help.aliyun.com/api/eci/describecontainergroups.html
func (client *Client) DescribeContainerGroups(request *DescribeContainerGroupsRequest) (response *DescribeContainerGroupsResponse, err error) {
	response = CreateDescribeContainerGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeContainerGroupsWithChan invokes the eci.DescribeContainerGroups API asynchronously
// api document: https://help.aliyun.com/api/eci/describecontainergroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeContainerGroupsWithChan(request *DescribeContainerGroupsRequest) (<-chan *DescribeContainerGroupsResponse, <-chan error) {
	responseChan := make(chan *DescribeContainerGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeContainerGroups(request)
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

// DescribeContainerGroupsWithCallback invokes the eci.DescribeContainerGroups API asynchronously
// api document: https://help.aliyun.com/api/eci/describecontainergroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeContainerGroupsWithCallback(request *DescribeContainerGroupsRequest, callback func(response *DescribeContainerGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeContainerGroupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeContainerGroups(request)
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

// DescribeContainerGroupsRequest is the request struct for api DescribeContainerGroups
type DescribeContainerGroupsRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer              `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string                        `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer              `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string                        `position:"Query" name:"OwnerAccount"`
	RegionId             string                        `position:"Query" name:"RegionId"`
	ZoneId               string                        `position:"Query" name:"ZoneId"`
	VSwitchId            string                        `position:"Query" name:"VSwitchId"`
	NextToken            string                        `position:"Query" name:"NextToken"`
	Limit                requests.Integer              `position:"Query" name:"Limit"`
	Tag                  *[]DescribeContainerGroupsTag `position:"Query" name:"Tag" type:"Repeated"`
	ContainerGroupIds    string                        `position:"Query" name:"ContainerGroupIds"`
	ContainerGroupName   string                        `position:"Query" name:"ContainerGroupName"`
	Status               string                        `position:"Query" name:"Status"`
}

type DescribeContainerGroupsTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// DescribeContainerGroupsResponse is the response struct for api DescribeContainerGroups
type DescribeContainerGroupsResponse struct {
	*responses.BaseResponse
	RequestId       string                                   `json:"RequestId" xml:"RequestId"`
	NextToken       string                                   `json:"NextToken" xml:"NextToken"`
	TotalCount      int                                      `json:"TotalCount" xml:"TotalCount"`
	ContainerGroups []DescribeContainerGroupsContainerGroup0 `json:"ContainerGroups" xml:"ContainerGroups"`
}

type DescribeContainerGroupsContainerGroup0 struct {
	ContainerGroupId   string                                     `json:"ContainerGroupId" xml:"ContainerGroupId"`
	ContainerGroupName string                                     `json:"ContainerGroupName" xml:"ContainerGroupName"`
	RegionId           string                                     `json:"RegionId" xml:"RegionId"`
	ZoneId             string                                     `json:"ZoneId" xml:"ZoneId"`
	Memory             float32                                    `json:"Memory" xml:"Memory"`
	Cpu                float32                                    `json:"Cpu" xml:"Cpu"`
	VSwitchId          string                                     `json:"VSwitchId" xml:"VSwitchId"`
	SecurityGroupId    string                                     `json:"SecurityGroupId" xml:"SecurityGroupId"`
	RestartPolicy      string                                     `json:"RestartPolicy" xml:"RestartPolicy"`
	IntranetIp         string                                     `json:"IntranetIp" xml:"IntranetIp"`
	Status             string                                     `json:"Status" xml:"Status"`
	InternetIp         string                                     `json:"InternetIp" xml:"InternetIp"`
	CreationTime       string                                     `json:"CreationTime" xml:"CreationTime"`
	SucceededTime      string                                     `json:"SucceededTime" xml:"SucceededTime"`
	EniInstanceId      string                                     `json:"EniInstanceId" xml:"EniInstanceId"`
	InstanceType       string                                     `json:"InstanceType" xml:"InstanceType"`
	ExpiredTime        string                                     `json:"ExpiredTime" xml:"ExpiredTime"`
	FailedTime         string                                     `json:"FailedTime" xml:"FailedTime"`
	Tags               []DescribeContainerGroupsLabel1            `json:"Tags" xml:"Tags"`
	Events             []DescribeContainerGroupsEvent1            `json:"Events" xml:"Events"`
	Containers         []DescribeContainerGroupsContainer1        `json:"Containers" xml:"Containers"`
	Volumes            []DescribeContainerGroupsVolume1           `json:"Volumes" xml:"Volumes"`
	InitContainers     []DescribeContainerGroupsContainer1        `json:"InitContainers" xml:"InitContainers"`
	HostAliases        []DescribeContainerGroupsHostAliase1       `json:"HostAliases" xml:"HostAliases"`
	DnsConfig          DescribeContainerGroupsDnsConfig1          `json:"DnsConfig" xml:"DnsConfig"`
	EciSecurityContext DescribeContainerGroupsEciSecurityContext1 `json:"EciSecurityContext" xml:"EciSecurityContext"`
}

type DescribeContainerGroupsLabel1 struct {
	Key   string `json:"Key" xml:"Key"`
	Value string `json:"Value" xml:"Value"`
}

type DescribeContainerGroupsEvent1 struct {
	Count          int    `json:"Count" xml:"Count"`
	Type           string `json:"Type" xml:"Type"`
	Name           string `json:"Name" xml:"Name"`
	Message        string `json:"Message" xml:"Message"`
	FirstTimestamp string `json:"FirstTimestamp" xml:"FirstTimestamp"`
	LastTimestamp  string `json:"LastTimestamp" xml:"LastTimestamp"`
	Reason         string `json:"Reason" xml:"Reason"`
}

type DescribeContainerGroupsContainer1 struct {
	Name            string                                   `json:"Name" xml:"Name"`
	Image           string                                   `json:"Image" xml:"Image"`
	Memory          float32                                  `json:"Memory" xml:"Memory"`
	Cpu             float32                                  `json:"Cpu" xml:"Cpu"`
	RestartCount    int                                      `json:"RestartCount" xml:"RestartCount"`
	WorkingDir      string                                   `json:"WorkingDir" xml:"WorkingDir"`
	ImagePullPolicy string                                   `json:"ImagePullPolicy" xml:"ImagePullPolicy"`
	Ready           bool                                     `json:"Ready" xml:"Ready"`
	Gpu             int                                      `json:"Gpu" xml:"Gpu"`
	VolumeMounts    []DescribeContainerGroupsVolumeMount2    `json:"VolumeMounts" xml:"VolumeMounts"`
	Ports           []DescribeContainerGroupsPort2           `json:"Ports" xml:"Ports"`
	EnvironmentVars []DescribeContainerGroupsEnvironmentVar2 `json:"EnvironmentVars" xml:"EnvironmentVars"`
	Commands        []string                                 `json:"Commands" xml:"Commands"`
	Args            []string                                 `json:"Args" xml:"Args"`
	PreviousState   DescribeContainerGroupsPreviousState2    `json:"PreviousState" xml:"PreviousState"`
	CurrentState    DescribeContainerGroupsCurrentState2     `json:"CurrentState" xml:"CurrentState"`
	ReadinessProbe  DescribeContainerGroupsReadinessProbe2   `json:"ReadinessProbe" xml:"ReadinessProbe"`
	LivenessProbe   DescribeContainerGroupsLivenessProbe2    `json:"LivenessProbe" xml:"LivenessProbe"`
	SecurityContext DescribeContainerGroupsSecurityContext2  `json:"SecurityContext" xml:"SecurityContext"`
}

type DescribeContainerGroupsVolumeMount2 struct {
	Name      string `json:"Name" xml:"Name"`
	MountPath string `json:"MountPath" xml:"MountPath"`
	ReadOnly  bool   `json:"ReadOnly" xml:"ReadOnly"`
}

type DescribeContainerGroupsPort2 struct {
	Port     int    `json:"Port" xml:"Port"`
	Protocol string `json:"Protocol" xml:"Protocol"`
}

type DescribeContainerGroupsEnvironmentVar2 struct {
	Key       string                            `json:"Key" xml:"Key"`
	Value     string                            `json:"Value" xml:"Value"`
	ValueFrom DescribeContainerGroupsValueFrom3 `json:"ValueFrom" xml:"ValueFrom"`
}

type DescribeContainerGroupsValueFrom3 struct {
	FieldRef DescribeContainerGroupsFieldRef4 `json:"FieldRef" xml:"FieldRef"`
}

type DescribeContainerGroupsFieldRef4 struct {
	FieldPath string `json:"FieldPath" xml:"FieldPath"`
}

type DescribeContainerGroupsPreviousState2 struct {
	State        string `json:"State" xml:"State"`
	DetailStatus string `json:"DetailStatus" xml:"DetailStatus"`
	ExitCode     int    `json:"ExitCode" xml:"ExitCode"`
	StartTime    string `json:"StartTime" xml:"StartTime"`
	FinishTime   string `json:"FinishTime" xml:"FinishTime"`
	Reason       string `json:"Reason" xml:"Reason"`
	Message      string `json:"Message" xml:"Message"`
	Signal       int    `json:"Signal" xml:"Signal"`
}

type DescribeContainerGroupsCurrentState2 struct {
	State        string `json:"State" xml:"State"`
	DetailStatus string `json:"DetailStatus" xml:"DetailStatus"`
	ExitCode     int    `json:"ExitCode" xml:"ExitCode"`
	StartTime    string `json:"StartTime" xml:"StartTime"`
	FinishTime   string `json:"FinishTime" xml:"FinishTime"`
	Reason       string `json:"Reason" xml:"Reason"`
	Message      string `json:"Message" xml:"Message"`
	Signal       int    `json:"Signal" xml:"Signal"`
}

type DescribeContainerGroupsReadinessProbe2 struct {
	InitialDelaySeconds int                               `json:"InitialDelaySeconds" xml:"InitialDelaySeconds"`
	PeriodSeconds       int                               `json:"PeriodSeconds" xml:"PeriodSeconds"`
	TimeoutSeconds      int                               `json:"TimeoutSeconds" xml:"TimeoutSeconds"`
	SuccessThreshold    int                               `json:"SuccessThreshold" xml:"SuccessThreshold"`
	FailureThreshold    int                               `json:"FailureThreshold" xml:"FailureThreshold"`
	Execs               []string                          `json:"Execs" xml:"Execs"`
	HttpGet             DescribeContainerGroupsHttpGet3   `json:"HttpGet" xml:"HttpGet"`
	TcpSocket           DescribeContainerGroupsTcpSocket3 `json:"TcpSocket" xml:"TcpSocket"`
}

type DescribeContainerGroupsHttpGet3 struct {
	Path   string `json:"Path" xml:"Path"`
	Port   int    `json:"Port" xml:"Port"`
	Scheme string `json:"Scheme" xml:"Scheme"`
}

type DescribeContainerGroupsTcpSocket3 struct {
	Host string `json:"Host" xml:"Host"`
	Port int    `json:"Port" xml:"Port"`
}

type DescribeContainerGroupsLivenessProbe2 struct {
	InitialDelaySeconds int                               `json:"InitialDelaySeconds" xml:"InitialDelaySeconds"`
	PeriodSeconds       int                               `json:"PeriodSeconds" xml:"PeriodSeconds"`
	TimeoutSeconds      int                               `json:"TimeoutSeconds" xml:"TimeoutSeconds"`
	SuccessThreshold    int                               `json:"SuccessThreshold" xml:"SuccessThreshold"`
	FailureThreshold    int                               `json:"FailureThreshold" xml:"FailureThreshold"`
	Execs               []string                          `json:"Execs" xml:"Execs"`
	HttpGet             DescribeContainerGroupsHttpGet3   `json:"HttpGet" xml:"HttpGet"`
	TcpSocket           DescribeContainerGroupsTcpSocket3 `json:"TcpSocket" xml:"TcpSocket"`
}

type DescribeContainerGroupsSecurityContext2 struct {
	ReadOnlyRootFilesystem bool                               `json:"ReadOnlyRootFilesystem" xml:"ReadOnlyRootFilesystem"`
	RunAsUser              int64                              `json:"RunAsUser" xml:"RunAsUser"`
	Capability             DescribeContainerGroupsCapability3 `json:"Capability" xml:"Capability"`
}

type DescribeContainerGroupsCapability3 struct {
	Adds []string `json:"Adds" xml:"Adds"`
}

type DescribeContainerGroupsVolume1 struct {
	Type                              string                                                     `json:"Type" xml:"Type"`
	Name                              string                                                     `json:"Name" xml:"Name"`
	NFSVolumePath                     string                                                     `json:"NFSVolumePath" xml:"NFSVolumePath"`
	NFSVolumeServer                   string                                                     `json:"NFSVolumeServer" xml:"NFSVolumeServer"`
	NFSVolumeReadOnly                 bool                                                       `json:"NFSVolumeReadOnly" xml:"NFSVolumeReadOnly"`
	ConfigFileVolumeConfigFileToPaths []DescribeContainerGroupsConfigFileVolumeConfigFileToPath2 `json:"ConfigFileVolumeConfigFileToPaths" xml:"ConfigFileVolumeConfigFileToPaths"`
}

type DescribeContainerGroupsConfigFileVolumeConfigFileToPath2 struct {
	Content string `json:"Content" xml:"Content"`
	Path    string `json:"Path" xml:"Path"`
}

type DescribeContainerGroupsHostAliase1 struct {
	Ip        string   `json:"Ip" xml:"Ip"`
	Hostnames []string `json:"Hostnames" xml:"Hostnames"`
}

type DescribeContainerGroupsDnsConfig1 struct {
	Options     []DescribeContainerGroupsOption2 `json:"Options" xml:"Options"`
	NameServers []string                         `json:"NameServers" xml:"NameServers"`
	Searches    []string                         `json:"Searches" xml:"Searches"`
}

type DescribeContainerGroupsOption2 struct {
	Name  string `json:"Name" xml:"Name"`
	Value string `json:"Value" xml:"Value"`
}

type DescribeContainerGroupsEciSecurityContext1 struct {
	Sysctls []DescribeContainerGroupsSysctl2 `json:"Sysctls" xml:"Sysctls"`
}

type DescribeContainerGroupsSysctl2 struct {
	Name  string `json:"Name" xml:"Name"`
	Value string `json:"Value" xml:"Value"`
}

// CreateDescribeContainerGroupsRequest creates a request to invoke DescribeContainerGroups API
func CreateDescribeContainerGroupsRequest() (request *DescribeContainerGroupsRequest) {
	request = &DescribeContainerGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Eci", "2018-08-08", "DescribeContainerGroups", "eci", "openAPI")
	return
}

// CreateDescribeContainerGroupsResponse creates a response to parse from DescribeContainerGroups response
func CreateDescribeContainerGroupsResponse() (response *DescribeContainerGroupsResponse) {
	response = &DescribeContainerGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
