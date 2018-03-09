package cms

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

// ListMyGroupInstancesDetails invokes the cms.ListMyGroupInstancesDetails API synchronously
// api document: https://help.aliyun.com/api/cms/listmygroupinstancesdetails.html
func (client *Client) ListMyGroupInstancesDetails(request *ListMyGroupInstancesDetailsRequest) (response *ListMyGroupInstancesDetailsResponse, err error) {
	response = CreateListMyGroupInstancesDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// ListMyGroupInstancesDetailsWithChan invokes the cms.ListMyGroupInstancesDetails API asynchronously
// api document: https://help.aliyun.com/api/cms/listmygroupinstancesdetails.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListMyGroupInstancesDetailsWithChan(request *ListMyGroupInstancesDetailsRequest) (<-chan *ListMyGroupInstancesDetailsResponse, <-chan error) {
	responseChan := make(chan *ListMyGroupInstancesDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMyGroupInstancesDetails(request)
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

// ListMyGroupInstancesDetailsWithCallback invokes the cms.ListMyGroupInstancesDetails API asynchronously
// api document: https://help.aliyun.com/api/cms/listmygroupinstancesdetails.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListMyGroupInstancesDetailsWithCallback(request *ListMyGroupInstancesDetailsRequest, callback func(response *ListMyGroupInstancesDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMyGroupInstancesDetailsResponse
		var err error
		defer close(result)
		response, err = client.ListMyGroupInstancesDetails(request)
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

// ListMyGroupInstancesDetailsRequest is the request struct for api ListMyGroupInstancesDetails
type ListMyGroupInstancesDetailsRequest struct {
	*requests.RpcRequest
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	Total      requests.Boolean `position:"Query" name:"Total"`
	GroupId    requests.Integer `position:"Query" name:"GroupId"`
	Category   string           `position:"Query" name:"Category"`
}

// ListMyGroupInstancesDetailsResponse is the response struct for api ListMyGroupInstancesDetails
type ListMyGroupInstancesDetailsResponse struct {
	*responses.BaseResponse
	RequestId    string                                 `json:"RequestId" xml:"RequestId"`
	Success      bool                                   `json:"Success" xml:"Success"`
	ErrorCode    int                                    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string                                 `json:"ErrorMessage" xml:"ErrorMessage"`
	PageNumber   int                                    `json:"PageNumber" xml:"PageNumber"`
	PageSize     int                                    `json:"PageSize" xml:"PageSize"`
	Total        int                                    `json:"Total" xml:"Total"`
	Resources    ResourcesInListMyGroupInstancesDetails `json:"Resources" xml:"Resources"`
}

// CreateListMyGroupInstancesDetailsRequest creates a request to invoke ListMyGroupInstancesDetails API
func CreateListMyGroupInstancesDetailsRequest(request *ListMyGroupInstancesDetailsRequest) {
	request = &ListMyGroupInstancesDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "ListMyGroupInstancesDetails", "cms", "openAPI")
	return
}

// CreateListMyGroupInstancesDetailsResponse creates a response to parse from ListMyGroupInstancesDetails response
func CreateListMyGroupInstancesDetailsResponse() (response *ListMyGroupInstancesDetailsResponse) {
	response = &ListMyGroupInstancesDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}