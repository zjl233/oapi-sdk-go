/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

// Package human_authentication code generated by oapi sdk gen
package larkhuman_authentication

import (
	"fmt"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 生成枚举值

const (
	UserIdTypeOpenId  = "open_id"
	UserIdTypeUserId  = "user_id"
	UserIdTypeUnionId = "union_id"
)

// 生成数据类型

type Identity struct {
	IdentityName *string `json:"identity_name,omitempty"`
	IdentityCode *string `json:"identity_code,omitempty"`
	Mobile       *string `json:"mobile,omitempty"`
}

// builder开始
type IdentityBuilder struct {
	identityName     string
	identityNameFlag bool
	identityCode     string
	identityCodeFlag bool
	mobile           string
	mobileFlag       bool
}

func NewIdentityBuilder() *IdentityBuilder {
	builder := &IdentityBuilder{}
	return builder
}

func (builder *IdentityBuilder) IdentityName(identityName string) *IdentityBuilder {
	builder.identityName = identityName
	builder.identityNameFlag = true
	return builder
}
func (builder *IdentityBuilder) IdentityCode(identityCode string) *IdentityBuilder {
	builder.identityCode = identityCode
	builder.identityCodeFlag = true
	return builder
}
func (builder *IdentityBuilder) Mobile(mobile string) *IdentityBuilder {
	builder.mobile = mobile
	builder.mobileFlag = true
	return builder
}

func (builder *IdentityBuilder) Build() *Identity {
	req := &Identity{}
	if builder.identityNameFlag {
		req.IdentityName = &builder.identityName

	}
	if builder.identityCodeFlag {
		req.IdentityCode = &builder.identityCode

	}
	if builder.mobileFlag {
		req.Mobile = &builder.mobile

	}
	return req
}

// builder结束

// 生成请求和响应结果类型，以及请求对象的Builder构造器

type CreateIdentityReqBodyBuilder struct {
	identityName     string
	identityNameFlag bool
	identityCode     string
	identityCodeFlag bool
	mobile           string
	mobileFlag       bool
}

// 生成body的New构造器
func NewCreateIdentityReqBodyBuilder() *CreateIdentityReqBodyBuilder {
	builder := &CreateIdentityReqBodyBuilder{}
	return builder
}

// 1.2 生成body的builder属性方法
func (builder *CreateIdentityReqBodyBuilder) IdentityName(identityName string) *CreateIdentityReqBodyBuilder {
	builder.identityName = identityName
	builder.identityNameFlag = true
	return builder
}
func (builder *CreateIdentityReqBodyBuilder) IdentityCode(identityCode string) *CreateIdentityReqBodyBuilder {
	builder.identityCode = identityCode
	builder.identityCodeFlag = true
	return builder
}
func (builder *CreateIdentityReqBodyBuilder) Mobile(mobile string) *CreateIdentityReqBodyBuilder {
	builder.mobile = mobile
	builder.mobileFlag = true
	return builder
}

// 1.3 生成body的build方法
func (builder *CreateIdentityReqBodyBuilder) Build() *CreateIdentityReqBody {
	req := &CreateIdentityReqBody{}
	if builder.identityNameFlag {
		req.IdentityName = &builder.identityName
	}
	if builder.identityCodeFlag {
		req.IdentityCode = &builder.identityCode
	}
	if builder.mobileFlag {
		req.Mobile = &builder.mobile
	}
	return req
}

// 上传文件path开始
type CreateIdentityPathReqBodyBuilder struct {
	identityName     string
	identityNameFlag bool
	identityCode     string
	identityCodeFlag bool
	mobile           string
	mobileFlag       bool
}

func NewCreateIdentityPathReqBodyBuilder() *CreateIdentityPathReqBodyBuilder {
	builder := &CreateIdentityPathReqBodyBuilder{}
	return builder
}
func (builder *CreateIdentityPathReqBodyBuilder) IdentityName(identityName string) *CreateIdentityPathReqBodyBuilder {
	builder.identityName = identityName
	builder.identityNameFlag = true
	return builder
}
func (builder *CreateIdentityPathReqBodyBuilder) IdentityCode(identityCode string) *CreateIdentityPathReqBodyBuilder {
	builder.identityCode = identityCode
	builder.identityCodeFlag = true
	return builder
}
func (builder *CreateIdentityPathReqBodyBuilder) Mobile(mobile string) *CreateIdentityPathReqBodyBuilder {
	builder.mobile = mobile
	builder.mobileFlag = true
	return builder
}

func (builder *CreateIdentityPathReqBodyBuilder) Build() (*CreateIdentityReqBody, error) {
	req := &CreateIdentityReqBody{}
	if builder.identityNameFlag {
		req.IdentityName = &builder.identityName
	}
	if builder.identityCodeFlag {
		req.IdentityCode = &builder.identityCode
	}
	if builder.mobileFlag {
		req.Mobile = &builder.mobile
	}
	return req, nil
}

// 上传文件path结束

// 1.4 生成请求的builder结构体
type CreateIdentityReqBuilder struct {
	apiReq *larkcore.ApiReq
	body   *CreateIdentityReqBody
}

// 生成请求的New构造器
func NewCreateIdentityReqBuilder() *CreateIdentityReqBuilder {
	builder := &CreateIdentityReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 1.5 生成请求的builder属性方法
func (builder *CreateIdentityReqBuilder) UserId(userId string) *CreateIdentityReqBuilder {
	builder.apiReq.QueryParams.Set("user_id", fmt.Sprint(userId))
	return builder
}
func (builder *CreateIdentityReqBuilder) UserIdType(userIdType string) *CreateIdentityReqBuilder {
	builder.apiReq.QueryParams.Set("user_id_type", fmt.Sprint(userIdType))
	return builder
}
func (builder *CreateIdentityReqBuilder) Body(body *CreateIdentityReqBody) *CreateIdentityReqBuilder {
	builder.body = body
	return builder
}

// 1.5 生成请求的builder的build方法
func (builder *CreateIdentityReqBuilder) Build() *CreateIdentityReq {
	req := &CreateIdentityReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.QueryParams = builder.apiReq.QueryParams
	req.apiReq.Body = builder.body
	return req
}

type CreateIdentityReqBody struct {
	IdentityName *string `json:"identity_name,omitempty"`
	IdentityCode *string `json:"identity_code,omitempty"`
	Mobile       *string `json:"mobile,omitempty"`
}

type CreateIdentityReq struct {
	apiReq *larkcore.ApiReq
	Body   *CreateIdentityReqBody `body:""`
}

type CreateIdentityRespData struct {
	VerifyUid *string `json:"verify_uid,omitempty"`
}

type CreateIdentityResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *CreateIdentityRespData `json:"data"`
}

func (resp *CreateIdentityResp) Success() bool {
	return resp.Code == 0
}

// 生成消息事件结构体

// 生成请求的builder构造器
// 1.1 生成body的builder结构体
