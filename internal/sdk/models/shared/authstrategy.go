// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type AuthStrategyType string

const (
	AuthStrategyTypeKeyAuth                      AuthStrategyType = "key_auth"
	AuthStrategyTypeClientCredentials            AuthStrategyType = "client_credentials"
	AuthStrategyTypeSelfManagedClientCredentials AuthStrategyType = "self_managed_client_credentials"
)

type AuthStrategy struct {
	AuthStrategyKeyAuth           *AuthStrategyKeyAuth
	AuthStrategyClientCredentials *AuthStrategyClientCredentials

	Type AuthStrategyType
}

func CreateAuthStrategyKeyAuth(keyAuth AuthStrategyKeyAuth) AuthStrategy {
	typ := AuthStrategyTypeKeyAuth

	typStr := CredentialType(typ)
	keyAuth.CredentialType = typStr

	return AuthStrategy{
		AuthStrategyKeyAuth: &keyAuth,
		Type:                typ,
	}
}

func CreateAuthStrategyClientCredentials(clientCredentials AuthStrategyClientCredentials) AuthStrategy {
	typ := AuthStrategyTypeClientCredentials

	typStr := AuthStrategyClientCredentialsCredentialType(typ)
	clientCredentials.CredentialType = typStr

	return AuthStrategy{
		AuthStrategyClientCredentials: &clientCredentials,
		Type:                          typ,
	}
}

func CreateAuthStrategySelfManagedClientCredentials(selfManagedClientCredentials AuthStrategyClientCredentials) AuthStrategy {
	typ := AuthStrategyTypeSelfManagedClientCredentials

	typStr := AuthStrategyClientCredentialsCredentialType(typ)
	selfManagedClientCredentials.CredentialType = typStr

	return AuthStrategy{
		AuthStrategyClientCredentials: &selfManagedClientCredentials,
		Type:                          typ,
	}
}

func (u *AuthStrategy) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		CredentialType string `json:"credential_type"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.CredentialType {
	case "key_auth":
		authStrategyKeyAuth := new(AuthStrategyKeyAuth)
		if err := utils.UnmarshalJSON(data, &authStrategyKeyAuth, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.AuthStrategyKeyAuth = authStrategyKeyAuth
		u.Type = AuthStrategyTypeKeyAuth
		return nil
	case "client_credentials":
		authStrategyClientCredentials := new(AuthStrategyClientCredentials)
		if err := utils.UnmarshalJSON(data, &authStrategyClientCredentials, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.AuthStrategyClientCredentials = authStrategyClientCredentials
		u.Type = AuthStrategyTypeClientCredentials
		return nil
	case "self_managed_client_credentials":
		authStrategyClientCredentials := new(AuthStrategyClientCredentials)
		if err := utils.UnmarshalJSON(data, &authStrategyClientCredentials, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.AuthStrategyClientCredentials = authStrategyClientCredentials
		u.Type = AuthStrategyTypeSelfManagedClientCredentials
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AuthStrategy) MarshalJSON() ([]byte, error) {
	if u.AuthStrategyKeyAuth != nil {
		return utils.MarshalJSON(u.AuthStrategyKeyAuth, "", true)
	}

	if u.AuthStrategyClientCredentials != nil {
		return utils.MarshalJSON(u.AuthStrategyClientCredentials, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}