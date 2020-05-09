package ddoscoo

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

// DescribeDDoSEvents invokes the ddoscoo.DescribeDDoSEvents API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/describeddosevents.html
func (client *Client) DescribeDDoSEvents(request *DescribeDDoSEventsRequest) (response *DescribeDDoSEventsResponse, err error) {
	response = CreateDescribeDDoSEventsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDDoSEventsWithChan invokes the ddoscoo.DescribeDDoSEvents API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describeddosevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDDoSEventsWithChan(request *DescribeDDoSEventsRequest) (<-chan *DescribeDDoSEventsResponse, <-chan error) {
	responseChan := make(chan *DescribeDDoSEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDDoSEvents(request)
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

// DescribeDDoSEventsWithCallback invokes the ddoscoo.DescribeDDoSEvents API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describeddosevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDDoSEventsWithCallback(request *DescribeDDoSEventsRequest, callback func(response *DescribeDDoSEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDDoSEventsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDDoSEvents(request)
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

// DescribeDDoSEventsRequest is the request struct for api DescribeDDoSEvents
type DescribeDDoSEventsRequest struct {
	*requests.RpcRequest
	StartTime       requests.Integer `position:"Query" name:"StartTime"`
	Eip             string           `position:"Query" name:"Eip"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageSize        string           `position:"Query" name:"PageSize"`
	Offset          requests.Integer `position:"Query" name:"Offset"`
	EndTime         requests.Integer `position:"Query" name:"EndTime"`
}

// DescribeDDoSEventsResponse is the response struct for api DescribeDDoSEvents
type DescribeDDoSEventsResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Total     int64   `json:"Total" xml:"Total"`
	Events    []Event `json:"Events" xml:"Events"`
}

// CreateDescribeDDoSEventsRequest creates a request to invoke DescribeDDoSEvents API
func CreateDescribeDDoSEventsRequest() (request *DescribeDDoSEventsRequest) {
	request = &DescribeDDoSEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2017-12-28", "DescribeDDoSEvents", "ddoscoo", "openAPI")
	return
}

// CreateDescribeDDoSEventsResponse creates a response to parse from DescribeDDoSEvents response
func CreateDescribeDDoSEventsResponse() (response *DescribeDDoSEventsResponse) {
	response = &DescribeDDoSEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}