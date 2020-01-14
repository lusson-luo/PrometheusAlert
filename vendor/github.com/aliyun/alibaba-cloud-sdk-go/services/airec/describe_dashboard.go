package airec

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

// DescribeDashboard invokes the airec.DescribeDashboard API synchronously
// api document: https://help.aliyun.com/api/airec/describedashboard.html
func (client *Client) DescribeDashboard(request *DescribeDashboardRequest) (response *DescribeDashboardResponse, err error) {
	response = CreateDescribeDashboardResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDashboardWithChan invokes the airec.DescribeDashboard API asynchronously
// api document: https://help.aliyun.com/api/airec/describedashboard.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDashboardWithChan(request *DescribeDashboardRequest) (<-chan *DescribeDashboardResponse, <-chan error) {
	responseChan := make(chan *DescribeDashboardResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDashboard(request)
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

// DescribeDashboardWithCallback invokes the airec.DescribeDashboard API asynchronously
// api document: https://help.aliyun.com/api/airec/describedashboard.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDashboardWithCallback(request *DescribeDashboardRequest, callback func(response *DescribeDashboardResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDashboardResponse
		var err error
		defer close(result)
		response, err = client.DescribeDashboard(request)
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

// DescribeDashboardRequest is the request struct for api DescribeDashboard
type DescribeDashboardRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
	EndDate    string `position:"Query" name:"EndDate"`
	StartDate  string `position:"Query" name:"StartDate"`
}

// DescribeDashboardResponse is the response struct for api DescribeDashboard
type DescribeDashboardResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Code      string       `json:"Code" xml:"Code"`
	Message   string       `json:"Message" xml:"Message"`
	Result    []ResultItem `json:"Result" xml:"Result"`
}

// CreateDescribeDashboardRequest creates a request to invoke DescribeDashboard API
func CreateDescribeDashboardRequest() (request *DescribeDashboardRequest) {
	request = &DescribeDashboardRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2018-10-12", "DescribeDashboard", "/openapi/instances/[InstanceId]/dashboard", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeDashboardResponse creates a response to parse from DescribeDashboard response
func CreateDescribeDashboardResponse() (response *DescribeDashboardResponse) {
	response = &DescribeDashboardResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}