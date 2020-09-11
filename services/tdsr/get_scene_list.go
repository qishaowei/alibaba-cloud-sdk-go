package tdsr

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

// GetSceneList invokes the tdsr.GetSceneList API synchronously
// api document: https://help.aliyun.com/api/tdsr/getscenelist.html
func (client *Client) GetSceneList(request *GetSceneListRequest) (response *GetSceneListResponse, err error) {
	response = CreateGetSceneListResponse()
	err = client.DoAction(request, response)
	return
}

// GetSceneListWithChan invokes the tdsr.GetSceneList API asynchronously
// api document: https://help.aliyun.com/api/tdsr/getscenelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSceneListWithChan(request *GetSceneListRequest) (<-chan *GetSceneListResponse, <-chan error) {
	responseChan := make(chan *GetSceneListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSceneList(request)
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

// GetSceneListWithCallback invokes the tdsr.GetSceneList API asynchronously
// api document: https://help.aliyun.com/api/tdsr/getscenelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSceneListWithCallback(request *GetSceneListRequest, callback func(response *GetSceneListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSceneListResponse
		var err error
		defer close(result)
		response, err = client.GetSceneList(request)
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

// GetSceneListRequest is the request struct for api GetSceneList
type GetSceneListRequest struct {
	*requests.RpcRequest
	AccountId string `position:"Query" name:"AccountId"`
}

// GetSceneListResponse is the response struct for api GetSceneList
type GetSceneListResponse struct {
	*responses.BaseResponse
	RequestId  string                 `json:"RequestId" xml:"RequestId"`
	Success    bool                   `json:"Success" xml:"Success"`
	ErrMessage string                 `json:"ErrMessage" xml:"ErrMessage"`
	Data       map[string]interface{} `json:"Data" xml:"Data"`
}

// CreateGetSceneListRequest creates a request to invoke GetSceneList API
func CreateGetSceneListRequest() (request *GetSceneListRequest) {
	request = &GetSceneListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("tdsr", "2020-01-01", "GetSceneList", "", "")
	request.Method = requests.POST
	return
}

// CreateGetSceneListResponse creates a response to parse from GetSceneList response
func CreateGetSceneListResponse() (response *GetSceneListResponse) {
	response = &GetSceneListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}