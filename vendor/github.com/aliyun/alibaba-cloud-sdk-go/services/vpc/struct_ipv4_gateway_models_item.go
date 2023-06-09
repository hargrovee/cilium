package vpc

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

// Ipv4GatewayModelsItem is a nested struct in vpc response
type Ipv4GatewayModelsItem struct {
	VpcId                   string `json:"VpcId" xml:"VpcId"`
	Status                  string `json:"Status" xml:"Status"`
	Ipv4GatewayId           string `json:"Ipv4GatewayId" xml:"Ipv4GatewayId"`
	Ipv4GatewayDescription  string `json:"Ipv4GatewayDescription" xml:"Ipv4GatewayDescription"`
	Enabled                 bool   `json:"Enabled" xml:"Enabled"`
	GmtCreate               string `json:"GmtCreate" xml:"GmtCreate"`
	Ipv4GatewayRouteTableId string `json:"Ipv4GatewayRouteTableId" xml:"Ipv4GatewayRouteTableId"`
	Ipv4GatewayName         string `json:"Ipv4GatewayName" xml:"Ipv4GatewayName"`
	ResourceGroupId         string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Tags                    []Tag  `json:"Tags" xml:"Tags"`
}
