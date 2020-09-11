package vod

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

// DynamicImage is a nested struct in vod response
type DynamicImage struct {
	VideoId        string `json:"VideoId" xml:"VideoId"`
	DynamicImageId string `json:"DynamicImageId" xml:"DynamicImageId"`
	JobId          string `json:"JobId" xml:"JobId"`
	FileURL        string `json:"FileURL" xml:"FileURL"`
	Width          string `json:"Width" xml:"Width"`
	Height         string `json:"Height" xml:"Height"`
	Duration       string `json:"Duration" xml:"Duration"`
	Format         string `json:"Format" xml:"Format"`
	FileSize       string `json:"FileSize" xml:"FileSize"`
	Fps            string `json:"Fps" xml:"Fps"`
	CreationTime   string `json:"CreationTime" xml:"CreationTime"`
}