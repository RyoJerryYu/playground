//Code generated by protoc-gen-go-fieldmask. DO NOT EDIT.
//versions:
//- protoc-gen-go-fieldmask v1.0.9
//- protoc (unknown)
//source: store/idp.proto

package store

// IIdentityProviderFieldPath is the interface for the field path of IdentityProvider
type IIdentityProviderFieldPath interface {
	String() string
	Id() string
	Name() string
	Type() string
	IdentifierFilter() string
	Config() IIdentityProviderConfigFieldPath
}

// identityProviderFieldPath is the implementation for the field path of IdentityProvider
type identityProviderFieldPath struct {
	fieldPath string // the field path to the current field, empty if it's root
	prefix    string // e.g. "fieldPath." or empty if it's root
}

// NewIdentityProviderFieldPath creates a new identityProviderFieldPath
func NewIdentityProviderFieldPath(fieldPath string) IIdentityProviderFieldPath {
	prefix := ""
	if fieldPath != "" {
		prefix = fieldPath + "."
	}
	return identityProviderFieldPath{fieldPath: fieldPath, prefix: prefix}
}

// String returns the field path
func (x identityProviderFieldPath) String() string { return x.fieldPath }

func (x identityProviderFieldPath) Id() string               { return x.prefix + "id" }
func (x identityProviderFieldPath) Name() string             { return x.prefix + "name" }
func (x identityProviderFieldPath) Type() string             { return x.prefix + "type" }
func (x identityProviderFieldPath) IdentifierFilter() string { return x.prefix + "identifier_filter" }
func (x identityProviderFieldPath) Config() IIdentityProviderConfigFieldPath {
	return NewIdentityProviderConfigFieldPath(x.prefix + "config")
}

// FieldPath returns the field path for IdentityProvider
func (x *IdentityProvider) FieldPath() IIdentityProviderFieldPath {
	return NewIdentityProviderFieldPath("")
}

// IIdentityProviderConfigFieldPath is the interface for the field path of IdentityProviderConfig
type IIdentityProviderConfigFieldPath interface {
	String() string
	Oauth2Config() IOAuth2ConfigFieldPath
}

// identityProviderConfigFieldPath is the implementation for the field path of IdentityProviderConfig
type identityProviderConfigFieldPath struct {
	fieldPath string // the field path to the current field, empty if it's root
	prefix    string // e.g. "fieldPath." or empty if it's root
}

// NewIdentityProviderConfigFieldPath creates a new identityProviderConfigFieldPath
func NewIdentityProviderConfigFieldPath(fieldPath string) IIdentityProviderConfigFieldPath {
	prefix := ""
	if fieldPath != "" {
		prefix = fieldPath + "."
	}
	return identityProviderConfigFieldPath{fieldPath: fieldPath, prefix: prefix}
}

// String returns the field path
func (x identityProviderConfigFieldPath) String() string { return x.fieldPath }

func (x identityProviderConfigFieldPath) Oauth2Config() IOAuth2ConfigFieldPath {
	return NewOAuth2ConfigFieldPath(x.prefix + "oauth2_config")
}

// FieldPath returns the field path for IdentityProviderConfig
func (x *IdentityProviderConfig) FieldPath() IIdentityProviderConfigFieldPath {
	return NewIdentityProviderConfigFieldPath("")
}

// IFieldMappingFieldPath is the interface for the field path of FieldMapping
type IFieldMappingFieldPath interface {
	String() string
	Identifier() string
	DisplayName() string
	Email() string
}

// fieldMappingFieldPath is the implementation for the field path of FieldMapping
type fieldMappingFieldPath struct {
	fieldPath string // the field path to the current field, empty if it's root
	prefix    string // e.g. "fieldPath." or empty if it's root
}

// NewFieldMappingFieldPath creates a new fieldMappingFieldPath
func NewFieldMappingFieldPath(fieldPath string) IFieldMappingFieldPath {
	prefix := ""
	if fieldPath != "" {
		prefix = fieldPath + "."
	}
	return fieldMappingFieldPath{fieldPath: fieldPath, prefix: prefix}
}

// String returns the field path
func (x fieldMappingFieldPath) String() string { return x.fieldPath }

func (x fieldMappingFieldPath) Identifier() string  { return x.prefix + "identifier" }
func (x fieldMappingFieldPath) DisplayName() string { return x.prefix + "display_name" }
func (x fieldMappingFieldPath) Email() string       { return x.prefix + "email" }

// FieldPath returns the field path for FieldMapping
func (x *FieldMapping) FieldPath() IFieldMappingFieldPath {
	return NewFieldMappingFieldPath("")
}

// IOAuth2ConfigFieldPath is the interface for the field path of OAuth2Config
type IOAuth2ConfigFieldPath interface {
	String() string
	ClientId() string
	ClientSecret() string
	AuthUrl() string
	TokenUrl() string
	UserInfoUrl() string
	Scopes() string
	FieldMapping() IFieldMappingFieldPath
}

// oAuth2ConfigFieldPath is the implementation for the field path of OAuth2Config
type oAuth2ConfigFieldPath struct {
	fieldPath string // the field path to the current field, empty if it's root
	prefix    string // e.g. "fieldPath." or empty if it's root
}

// NewOAuth2ConfigFieldPath creates a new oAuth2ConfigFieldPath
func NewOAuth2ConfigFieldPath(fieldPath string) IOAuth2ConfigFieldPath {
	prefix := ""
	if fieldPath != "" {
		prefix = fieldPath + "."
	}
	return oAuth2ConfigFieldPath{fieldPath: fieldPath, prefix: prefix}
}

// String returns the field path
func (x oAuth2ConfigFieldPath) String() string { return x.fieldPath }

func (x oAuth2ConfigFieldPath) ClientId() string     { return x.prefix + "client_id" }
func (x oAuth2ConfigFieldPath) ClientSecret() string { return x.prefix + "client_secret" }
func (x oAuth2ConfigFieldPath) AuthUrl() string      { return x.prefix + "auth_url" }
func (x oAuth2ConfigFieldPath) TokenUrl() string     { return x.prefix + "token_url" }
func (x oAuth2ConfigFieldPath) UserInfoUrl() string  { return x.prefix + "user_info_url" }
func (x oAuth2ConfigFieldPath) Scopes() string       { return x.prefix + "scopes" }
func (x oAuth2ConfigFieldPath) FieldMapping() IFieldMappingFieldPath {
	return NewFieldMappingFieldPath(x.prefix + "field_mapping")
}

// FieldPath returns the field path for OAuth2Config
func (x *OAuth2Config) FieldPath() IOAuth2ConfigFieldPath {
	return NewOAuth2ConfigFieldPath("")
}
