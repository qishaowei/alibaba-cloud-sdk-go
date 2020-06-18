package aliyuncvc

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

// CheckMeetingCode invokes the aliyuncvc.CheckMeetingCode API synchronously
// api document: https://help.aliyun.com/api/aliyuncvc/checkmeetingcode.html
func (client *Client) CheckMeetingCode(request *CheckMeetingCodeRequest) (response *CheckMeetingCodeResponse, err error) {
	response = CreateCheckMeetingCodeResponse()
	err = client.DoAction(request, response)
	return
}

// CheckMeetingCodeWithChan invokes the aliyuncvc.CheckMeetingCode API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/checkmeetingcode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckMeetingCodeWithChan(request *CheckMeetingCodeRequest) (<-chan *CheckMeetingCodeResponse, <-chan error) {
	responseChan := make(chan *CheckMeetingCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckMeetingCode(request)
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

// CheckMeetingCodeWithCallback invokes the aliyuncvc.CheckMeetingCode API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/checkmeetingcode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckMeetingCodeWithCallback(request *CheckMeetingCodeRequest, callback func(response *CheckMeetingCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckMeetingCodeResponse
		var err error
		defer close(result)
		response, err = client.CheckMeetingCode(request)
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

// CheckMeetingCodeRequest is the request struct for api CheckMeetingCode
type CheckMeetingCodeRequest struct {
	*requests.RpcRequest
	MeetingCode string `position:"Body" name:"MeetingCode"`
	UserId      string `position:"Body" name:"UserId"`
}

// CheckMeetingCodeResponse is the response struct for api CheckMeetingCode
type CheckMeetingCodeResponse struct {
	*responses.BaseResponse
	ErrorCode   int         `json:"ErrorCode" xml:"ErrorCode"`
	Success     bool        `json:"Success" xml:"Success"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	Message     string      `json:"Message" xml:"Message"`
	MeetingInfo MeetingInfo `json:"MeetingInfo" xml:"MeetingInfo"`
}

// CreateCheckMeetingCodeRequest creates a request to invoke CheckMeetingCode API
func CreateCheckMeetingCodeRequest() (request *CheckMeetingCodeRequest) {
	request = &CheckMeetingCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-09-19", "CheckMeetingCode", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckMeetingCodeResponse creates a response to parse from CheckMeetingCode response
func CreateCheckMeetingCodeResponse() (response *CheckMeetingCodeResponse) {
	response = &CheckMeetingCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
