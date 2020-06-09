// Code generated by protoc-gen-goext. DO NOT EDIT.

package iam

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

type CreateIamTokenRequest_Identity = isCreateIamTokenRequest_Identity

func (m *CreateIamTokenRequest) SetIdentity(v CreateIamTokenRequest_Identity) {
	m.Identity = v
}

func (m *CreateIamTokenRequest) SetYandexPassportOauthToken(v string) {
	m.Identity = &CreateIamTokenRequest_YandexPassportOauthToken{
		YandexPassportOauthToken: v,
	}
}

func (m *CreateIamTokenRequest) SetJwt(v string) {
	m.Identity = &CreateIamTokenRequest_Jwt{
		Jwt: v,
	}
}

func (m *CreateIamTokenResponse) SetIamToken(v string) {
	m.IamToken = v
}

func (m *CreateIamTokenResponse) SetExpiresAt(v *timestamp.Timestamp) {
	m.ExpiresAt = v
}
