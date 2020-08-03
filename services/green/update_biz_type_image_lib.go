package green

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

// UpdateBizTypeImageLib invokes the green.UpdateBizTypeImageLib API synchronously
// api document: https://help.aliyun.com/api/green/updatebiztypeimagelib.html
func (client *Client) UpdateBizTypeImageLib(request *UpdateBizTypeImageLibRequest) (response *UpdateBizTypeImageLibResponse, err error) {
	response = CreateUpdateBizTypeImageLibResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateBizTypeImageLibWithChan invokes the green.UpdateBizTypeImageLib API asynchronously
// api document: https://help.aliyun.com/api/green/updatebiztypeimagelib.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateBizTypeImageLibWithChan(request *UpdateBizTypeImageLibRequest) (<-chan *UpdateBizTypeImageLibResponse, <-chan error) {
	responseChan := make(chan *UpdateBizTypeImageLibResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateBizTypeImageLib(request)
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

// UpdateBizTypeImageLibWithCallback invokes the green.UpdateBizTypeImageLib API asynchronously
// api document: https://help.aliyun.com/api/green/updatebiztypeimagelib.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateBizTypeImageLibWithCallback(request *UpdateBizTypeImageLibRequest, callback func(response *UpdateBizTypeImageLibResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateBizTypeImageLibResponse
		var err error
		defer close(result)
		response, err = client.UpdateBizTypeImageLib(request)
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

// UpdateBizTypeImageLibRequest is the request struct for api UpdateBizTypeImageLib
type UpdateBizTypeImageLibRequest struct {
	*requests.RpcRequest
	Scene        string `position:"Query" name:"Scene"`
	SourceIp     string `position:"Query" name:"SourceIp"`
	White        string `position:"Query" name:"White"`
	Review       string `position:"Query" name:"Review"`
	BizTypeName  string `position:"Query" name:"BizTypeName"`
	Black        string `position:"Query" name:"Black"`
	ResourceType string `position:"Query" name:"ResourceType"`
}

// UpdateBizTypeImageLibResponse is the response struct for api UpdateBizTypeImageLib
type UpdateBizTypeImageLibResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateBizTypeImageLibRequest creates a request to invoke UpdateBizTypeImageLib API
func CreateUpdateBizTypeImageLibRequest() (request *UpdateBizTypeImageLibRequest) {
	request = &UpdateBizTypeImageLibRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "UpdateBizTypeImageLib", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateBizTypeImageLibResponse creates a response to parse from UpdateBizTypeImageLib response
func CreateUpdateBizTypeImageLibResponse() (response *UpdateBizTypeImageLibResponse) {
	response = &UpdateBizTypeImageLibResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}