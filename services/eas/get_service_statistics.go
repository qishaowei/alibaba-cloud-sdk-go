package eas

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

// GetServiceStatistics invokes the eas.GetServiceStatistics API synchronously
// api document: https://help.aliyun.com/api/eas/getservicestatistics.html
func (client *Client) GetServiceStatistics(request *GetServiceStatisticsRequest) (response *GetServiceStatisticsResponse, err error) {
	response = CreateGetServiceStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// GetServiceStatisticsWithChan invokes the eas.GetServiceStatistics API asynchronously
// api document: https://help.aliyun.com/api/eas/getservicestatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetServiceStatisticsWithChan(request *GetServiceStatisticsRequest) (<-chan *GetServiceStatisticsResponse, <-chan error) {
	responseChan := make(chan *GetServiceStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetServiceStatistics(request)
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

// GetServiceStatisticsWithCallback invokes the eas.GetServiceStatistics API asynchronously
// api document: https://help.aliyun.com/api/eas/getservicestatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetServiceStatisticsWithCallback(request *GetServiceStatisticsRequest, callback func(response *GetServiceStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetServiceStatisticsResponse
		var err error
		defer close(result)
		response, err = client.GetServiceStatistics(request)
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

// GetServiceStatisticsRequest is the request struct for api GetServiceStatistics
type GetServiceStatisticsRequest struct {
	*requests.RoaRequest
	Metric      string           `position:"Query" name:"metric"`
	ServiceName string           `position:"Path" name:"service_name"`
	Count       requests.Integer `position:"Query" name:"count"`
	From        string           `position:"Query" name:"from"`
	To          string           `position:"Query" name:"to"`
	Region      string           `position:"Path" name:"region"`
}

// GetServiceStatisticsResponse is the response struct for api GetServiceStatistics
type GetServiceStatisticsResponse struct {
	*responses.BaseResponse
}

// CreateGetServiceStatisticsRequest creates a request to invoke GetServiceStatistics API
func CreateGetServiceStatisticsRequest() (request *GetServiceStatisticsRequest) {
	request = &GetServiceStatisticsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2018-05-22", "GetServiceStatistics", "/api/services/[region]/[service_name]/statistics", "", "")
	request.Method = requests.GET
	return
}

// CreateGetServiceStatisticsResponse creates a response to parse from GetServiceStatistics response
func CreateGetServiceStatisticsResponse() (response *GetServiceStatisticsResponse) {
	response = &GetServiceStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}