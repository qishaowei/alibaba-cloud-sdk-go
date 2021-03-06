package waf_openapi

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

// CreateOutputDomain invokes the waf_openapi.CreateOutputDomain API synchronously
func (client *Client) CreateOutputDomain(request *CreateOutputDomainRequest) (response *CreateOutputDomainResponse, err error) {
	response = CreateCreateOutputDomainResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOutputDomainWithChan invokes the waf_openapi.CreateOutputDomain API asynchronously
func (client *Client) CreateOutputDomainWithChan(request *CreateOutputDomainRequest) (<-chan *CreateOutputDomainResponse, <-chan error) {
	responseChan := make(chan *CreateOutputDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOutputDomain(request)
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

// CreateOutputDomainWithCallback invokes the waf_openapi.CreateOutputDomain API asynchronously
func (client *Client) CreateOutputDomainWithCallback(request *CreateOutputDomainRequest, callback func(response *CreateOutputDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOutputDomainResponse
		var err error
		defer close(result)
		response, err = client.CreateOutputDomain(request)
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

// CreateOutputDomainRequest is the request struct for api CreateOutputDomain
type CreateOutputDomainRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Domain          string `position:"Query" name:"Domain"`
	Lang            string `position:"Query" name:"Lang"`
	Region          string `position:"Query" name:"Region"`
}

// CreateOutputDomainResponse is the response struct for api CreateOutputDomain
type CreateOutputDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateCreateOutputDomainRequest creates a request to invoke CreateOutputDomain API
func CreateCreateOutputDomainRequest() (request *CreateOutputDomainRequest) {
	request = &CreateOutputDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "CreateOutputDomain", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateOutputDomainResponse creates a response to parse from CreateOutputDomain response
func CreateCreateOutputDomainResponse() (response *CreateOutputDomainResponse) {
	response = &CreateOutputDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
