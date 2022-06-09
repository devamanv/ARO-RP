package resourcemanager

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

// ListHandshakesForAccount invokes the resourcemanager.ListHandshakesForAccount API synchronously
func (client *Client) ListHandshakesForAccount(request *ListHandshakesForAccountRequest) (response *ListHandshakesForAccountResponse, err error) {
	response = CreateListHandshakesForAccountResponse()
	err = client.DoAction(request, response)
	return
}

// ListHandshakesForAccountWithChan invokes the resourcemanager.ListHandshakesForAccount API asynchronously
func (client *Client) ListHandshakesForAccountWithChan(request *ListHandshakesForAccountRequest) (<-chan *ListHandshakesForAccountResponse, <-chan error) {
	responseChan := make(chan *ListHandshakesForAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListHandshakesForAccount(request)
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

// ListHandshakesForAccountWithCallback invokes the resourcemanager.ListHandshakesForAccount API asynchronously
func (client *Client) ListHandshakesForAccountWithCallback(request *ListHandshakesForAccountRequest, callback func(response *ListHandshakesForAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListHandshakesForAccountResponse
		var err error
		defer close(result)
		response, err = client.ListHandshakesForAccount(request)
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

// ListHandshakesForAccountRequest is the request struct for api ListHandshakesForAccount
type ListHandshakesForAccountRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListHandshakesForAccountResponse is the response struct for api ListHandshakesForAccount
type ListHandshakesForAccountResponse struct {
	*responses.BaseResponse
	TotalCount int                                  `json:"TotalCount" xml:"TotalCount"`
	RequestId  string                               `json:"RequestId" xml:"RequestId"`
	PageSize   int                                  `json:"PageSize" xml:"PageSize"`
	PageNumber int                                  `json:"PageNumber" xml:"PageNumber"`
	Handshakes HandshakesInListHandshakesForAccount `json:"Handshakes" xml:"Handshakes"`
}

// CreateListHandshakesForAccountRequest creates a request to invoke ListHandshakesForAccount API
func CreateListHandshakesForAccountRequest() (request *ListHandshakesForAccountRequest) {
	request = &ListHandshakesForAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "ListHandshakesForAccount", "", "")
	request.Method = requests.POST
	return
}

// CreateListHandshakesForAccountResponse creates a response to parse from ListHandshakesForAccount response
func CreateListHandshakesForAccountResponse() (response *ListHandshakesForAccountResponse) {
	response = &ListHandshakesForAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}