package cloudcallcenter

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

// DialogueLite invokes the cloudcallcenter.DialogueLite API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/dialoguelite.html
func (client *Client) DialogueLite(request *DialogueLiteRequest) (response *DialogueLiteResponse, err error) {
	response = CreateDialogueLiteResponse()
	err = client.DoAction(request, response)
	return
}

// DialogueLiteWithChan invokes the cloudcallcenter.DialogueLite API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/dialoguelite.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DialogueLiteWithChan(request *DialogueLiteRequest) (<-chan *DialogueLiteResponse, <-chan error) {
	responseChan := make(chan *DialogueLiteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DialogueLite(request)
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

// DialogueLiteWithCallback invokes the cloudcallcenter.DialogueLite API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/dialoguelite.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DialogueLiteWithCallback(request *DialogueLiteRequest, callback func(response *DialogueLiteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DialogueLiteResponse
		var err error
		defer close(result)
		response, err = client.DialogueLite(request)
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

// DialogueLiteRequest is the request struct for api DialogueLite
type DialogueLiteRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	VendorParams string `position:"Query" name:"VendorParams"`
	CallType     string `position:"Query" name:"CallType"`
	Utterance    string `position:"Query" name:"Utterance"`
}

// DialogueLiteResponse is the response struct for api DialogueLite
type DialogueLiteResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Feedback       Feedback `json:"Feedback" xml:"Feedback"`
}

// CreateDialogueLiteRequest creates a request to invoke DialogueLite API
func CreateDialogueLiteRequest() (request *DialogueLiteRequest) {
	request = &DialogueLiteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DialogueLite", "", "")
	request.Method = requests.POST
	return
}

// CreateDialogueLiteResponse creates a response to parse from DialogueLite response
func CreateDialogueLiteResponse() (response *DialogueLiteResponse) {
	response = &DialogueLiteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}