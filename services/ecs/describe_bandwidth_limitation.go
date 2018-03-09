package ecs

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

// DescribeBandwidthLimitation invokes the ecs.DescribeBandwidthLimitation API synchronously
// api document: https://help.aliyun.com/api/ecs/describebandwidthlimitation.html
func (client *Client) DescribeBandwidthLimitation(request *DescribeBandwidthLimitationRequest) (response *DescribeBandwidthLimitationResponse, err error) {
	response = CreateDescribeBandwidthLimitationResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBandwidthLimitationWithChan invokes the ecs.DescribeBandwidthLimitation API asynchronously
// api document: https://help.aliyun.com/api/ecs/describebandwidthlimitation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBandwidthLimitationWithChan(request *DescribeBandwidthLimitationRequest) (<-chan *DescribeBandwidthLimitationResponse, <-chan error) {
	responseChan := make(chan *DescribeBandwidthLimitationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBandwidthLimitation(request)
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

// DescribeBandwidthLimitationWithCallback invokes the ecs.DescribeBandwidthLimitation API asynchronously
// api document: https://help.aliyun.com/api/ecs/describebandwidthlimitation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBandwidthLimitationWithCallback(request *DescribeBandwidthLimitationRequest, callback func(response *DescribeBandwidthLimitationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBandwidthLimitationResponse
		var err error
		defer close(result)
		response, err = client.DescribeBandwidthLimitation(request)
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

// DescribeBandwidthLimitationRequest is the request struct for api DescribeBandwidthLimitation
type DescribeBandwidthLimitationRequest struct {
	*requests.RpcRequest
}

// DescribeBandwidthLimitationResponse is the response struct for api DescribeBandwidthLimitation
type DescribeBandwidthLimitationResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Bandwidths Bandwidths `json:"Bandwidths" xml:"Bandwidths"`
}

// CreateDescribeBandwidthLimitationRequest creates a request to invoke DescribeBandwidthLimitation API
func CreateDescribeBandwidthLimitationRequest(request *DescribeBandwidthLimitationRequest) {
	request = &DescribeBandwidthLimitationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeBandwidthLimitation", "ecs", "openAPI")
	return
}

// CreateDescribeBandwidthLimitationResponse creates a response to parse from DescribeBandwidthLimitation response
func CreateDescribeBandwidthLimitationResponse() (response *DescribeBandwidthLimitationResponse) {
	response = &DescribeBandwidthLimitationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}