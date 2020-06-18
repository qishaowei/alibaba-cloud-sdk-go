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

// QueryMemberRecord invokes the aliyuncvc.QueryMemberRecord API synchronously
// api document: https://help.aliyun.com/api/aliyuncvc/querymemberrecord.html
func (client *Client) QueryMemberRecord(request *QueryMemberRecordRequest) (response *QueryMemberRecordResponse, err error) {
	response = CreateQueryMemberRecordResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMemberRecordWithChan invokes the aliyuncvc.QueryMemberRecord API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/querymemberrecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMemberRecordWithChan(request *QueryMemberRecordRequest) (<-chan *QueryMemberRecordResponse, <-chan error) {
	responseChan := make(chan *QueryMemberRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMemberRecord(request)
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

// QueryMemberRecordWithCallback invokes the aliyuncvc.QueryMemberRecord API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/querymemberrecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMemberRecordWithCallback(request *QueryMemberRecordRequest, callback func(response *QueryMemberRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMemberRecordResponse
		var err error
		defer close(result)
		response, err = client.QueryMemberRecord(request)
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

// QueryMemberRecordRequest is the request struct for api QueryMemberRecord
type QueryMemberRecordRequest struct {
	*requests.RpcRequest
	MeetingUUID string `position:"Query" name:"MeetingUUID"`
}

// QueryMemberRecordResponse is the response struct for api QueryMemberRecord
type QueryMemberRecordResponse struct {
	*responses.BaseResponse
	ErrorCode   int         `json:"ErrorCode" xml:"ErrorCode"`
	Message     string      `json:"Message" xml:"Message"`
	Success     bool        `json:"Success" xml:"Success"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	MeetingInfo MeetingInfo `json:"MeetingInfo" xml:"MeetingInfo"`
}

// CreateQueryMemberRecordRequest creates a request to invoke QueryMemberRecord API
func CreateQueryMemberRecordRequest() (request *QueryMemberRecordRequest) {
	request = &QueryMemberRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-09-19", "QueryMemberRecord", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryMemberRecordResponse creates a response to parse from QueryMemberRecord response
func CreateQueryMemberRecordResponse() (response *QueryMemberRecordResponse) {
	response = &QueryMemberRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
