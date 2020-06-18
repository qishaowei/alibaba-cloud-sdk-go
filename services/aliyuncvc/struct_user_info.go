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

// UserInfo is a nested struct in aliyuncvc response
type UserInfo struct {
	UserMobile          string `json:"UserMobile" xml:"UserMobile"`
	CurNum              int    `json:"CurNum" xml:"CurNum"`
	DepartName          string `json:"DepartName" xml:"DepartName"`
	GroupName           string `json:"GroupName" xml:"GroupName"`
	CreateTime          int64  `json:"CreateTime" xml:"CreateTime"`
	UserTel             string `json:"UserTel" xml:"UserTel"`
	UserAvatarUrl       string `json:"UserAvatarUrl" xml:"UserAvatarUrl"`
	UserId              string `json:"UserId" xml:"UserId"`
	UserName            string `json:"UserName" xml:"UserName"`
	GroupId             string `json:"GroupId" xml:"GroupId"`
	MaxNum              int    `json:"MaxNum" xml:"MaxNum"`
	DepartId            string `json:"DepartId" xml:"DepartId"`
	UserEmail           string `json:"UserEmail" xml:"UserEmail"`
	MemberConcurrentMax int    `json:"MemberConcurrentMax" xml:"MemberConcurrentMax"`
}
