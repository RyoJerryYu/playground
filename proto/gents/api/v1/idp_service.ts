// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.2.0
//   protoc               unknown
// source: api/v1/idp_service.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { Empty } from "../../google/protobuf/empty";
import { FieldMask } from "../../google/protobuf/field_mask";

export const protobufPackage = "memos.api.v1";

export interface IdentityProvider {
  /**
   * The name of the identityProvider.
   * Format: identityProviders/{id}
   */
  name: string;
  type: IdentityProvider_Type;
  title: string;
  identifierFilter: string;
  config?: IdentityProviderConfig | undefined;
}

export enum IdentityProvider_Type {
  TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED",
  OAUTH2 = "OAUTH2",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function identityProvider_TypeFromJSON(object: any): IdentityProvider_Type {
  switch (object) {
    case 0:
    case "TYPE_UNSPECIFIED":
      return IdentityProvider_Type.TYPE_UNSPECIFIED;
    case 1:
    case "OAUTH2":
      return IdentityProvider_Type.OAUTH2;
    case -1:
    case "UNRECOGNIZED":
    default:
      return IdentityProvider_Type.UNRECOGNIZED;
  }
}

export function identityProvider_TypeToNumber(object: IdentityProvider_Type): number {
  switch (object) {
    case IdentityProvider_Type.TYPE_UNSPECIFIED:
      return 0;
    case IdentityProvider_Type.OAUTH2:
      return 1;
    case IdentityProvider_Type.UNRECOGNIZED:
    default:
      return -1;
  }
}

export interface IdentityProviderConfig {
  oauth2Config?: OAuth2Config | undefined;
}

export interface FieldMapping {
  identifier: string;
  displayName: string;
  email: string;
}

export interface OAuth2Config {
  clientId: string;
  clientSecret: string;
  authUrl: string;
  tokenUrl: string;
  userInfoUrl: string;
  scopes: string[];
  fieldMapping?: FieldMapping | undefined;
}

export interface ListIdentityProvidersRequest {
}

export interface ListIdentityProvidersResponse {
  identityProviders: IdentityProvider[];
}

export interface GetIdentityProviderRequest {
  /**
   * The name of the identityProvider to get.
   * Format: identityProviders/{id}
   */
  name: string;
}

export interface CreateIdentityProviderRequest {
  /** The identityProvider to create. */
  identityProvider?: IdentityProvider | undefined;
}

export interface UpdateIdentityProviderRequest {
  /** The identityProvider to update. */
  identityProvider?:
    | IdentityProvider
    | undefined;
  /**
   * The update mask applies to the resource. Only the top level fields of
   * IdentityProvider are supported.
   */
  updateMask?: string[] | undefined;
}

export interface DeleteIdentityProviderRequest {
  /**
   * The name of the identityProvider to delete.
   * Format: identityProviders/{id}
   */
  name: string;
}

function createBaseIdentityProvider(): IdentityProvider {
  return { name: "", type: IdentityProvider_Type.TYPE_UNSPECIFIED, title: "", identifierFilter: "", config: undefined };
}

export const IdentityProvider: MessageFns<IdentityProvider> = {
  encode(message: IdentityProvider, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.type !== IdentityProvider_Type.TYPE_UNSPECIFIED) {
      writer.uint32(16).int32(identityProvider_TypeToNumber(message.type));
    }
    if (message.title !== "") {
      writer.uint32(26).string(message.title);
    }
    if (message.identifierFilter !== "") {
      writer.uint32(34).string(message.identifierFilter);
    }
    if (message.config !== undefined) {
      IdentityProviderConfig.encode(message.config, writer.uint32(42).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): IdentityProvider {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIdentityProvider();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.type = identityProvider_TypeFromJSON(reader.int32());
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.title = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.identifierFilter = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.config = IdentityProviderConfig.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<IdentityProvider>): IdentityProvider {
    return IdentityProvider.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<IdentityProvider>): IdentityProvider {
    const message = createBaseIdentityProvider();
    message.name = object.name ?? "";
    message.type = object.type ?? IdentityProvider_Type.TYPE_UNSPECIFIED;
    message.title = object.title ?? "";
    message.identifierFilter = object.identifierFilter ?? "";
    message.config = (object.config !== undefined && object.config !== null)
      ? IdentityProviderConfig.fromPartial(object.config)
      : undefined;
    return message;
  },
};

function createBaseIdentityProviderConfig(): IdentityProviderConfig {
  return { oauth2Config: undefined };
}

export const IdentityProviderConfig: MessageFns<IdentityProviderConfig> = {
  encode(message: IdentityProviderConfig, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.oauth2Config !== undefined) {
      OAuth2Config.encode(message.oauth2Config, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): IdentityProviderConfig {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIdentityProviderConfig();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.oauth2Config = OAuth2Config.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<IdentityProviderConfig>): IdentityProviderConfig {
    return IdentityProviderConfig.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<IdentityProviderConfig>): IdentityProviderConfig {
    const message = createBaseIdentityProviderConfig();
    message.oauth2Config = (object.oauth2Config !== undefined && object.oauth2Config !== null)
      ? OAuth2Config.fromPartial(object.oauth2Config)
      : undefined;
    return message;
  },
};

function createBaseFieldMapping(): FieldMapping {
  return { identifier: "", displayName: "", email: "" };
}

export const FieldMapping: MessageFns<FieldMapping> = {
  encode(message: FieldMapping, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.identifier !== "") {
      writer.uint32(10).string(message.identifier);
    }
    if (message.displayName !== "") {
      writer.uint32(18).string(message.displayName);
    }
    if (message.email !== "") {
      writer.uint32(26).string(message.email);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): FieldMapping {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFieldMapping();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.identifier = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.displayName = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.email = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<FieldMapping>): FieldMapping {
    return FieldMapping.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<FieldMapping>): FieldMapping {
    const message = createBaseFieldMapping();
    message.identifier = object.identifier ?? "";
    message.displayName = object.displayName ?? "";
    message.email = object.email ?? "";
    return message;
  },
};

function createBaseOAuth2Config(): OAuth2Config {
  return {
    clientId: "",
    clientSecret: "",
    authUrl: "",
    tokenUrl: "",
    userInfoUrl: "",
    scopes: [],
    fieldMapping: undefined,
  };
}

export const OAuth2Config: MessageFns<OAuth2Config> = {
  encode(message: OAuth2Config, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.clientId !== "") {
      writer.uint32(10).string(message.clientId);
    }
    if (message.clientSecret !== "") {
      writer.uint32(18).string(message.clientSecret);
    }
    if (message.authUrl !== "") {
      writer.uint32(26).string(message.authUrl);
    }
    if (message.tokenUrl !== "") {
      writer.uint32(34).string(message.tokenUrl);
    }
    if (message.userInfoUrl !== "") {
      writer.uint32(42).string(message.userInfoUrl);
    }
    for (const v of message.scopes) {
      writer.uint32(50).string(v!);
    }
    if (message.fieldMapping !== undefined) {
      FieldMapping.encode(message.fieldMapping, writer.uint32(58).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): OAuth2Config {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOAuth2Config();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.clientId = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.clientSecret = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.authUrl = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.tokenUrl = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.userInfoUrl = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.scopes.push(reader.string());
          continue;
        case 7:
          if (tag !== 58) {
            break;
          }

          message.fieldMapping = FieldMapping.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<OAuth2Config>): OAuth2Config {
    return OAuth2Config.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<OAuth2Config>): OAuth2Config {
    const message = createBaseOAuth2Config();
    message.clientId = object.clientId ?? "";
    message.clientSecret = object.clientSecret ?? "";
    message.authUrl = object.authUrl ?? "";
    message.tokenUrl = object.tokenUrl ?? "";
    message.userInfoUrl = object.userInfoUrl ?? "";
    message.scopes = object.scopes?.map((e) => e) || [];
    message.fieldMapping = (object.fieldMapping !== undefined && object.fieldMapping !== null)
      ? FieldMapping.fromPartial(object.fieldMapping)
      : undefined;
    return message;
  },
};

function createBaseListIdentityProvidersRequest(): ListIdentityProvidersRequest {
  return {};
}

export const ListIdentityProvidersRequest: MessageFns<ListIdentityProvidersRequest> = {
  encode(_: ListIdentityProvidersRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ListIdentityProvidersRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListIdentityProvidersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<ListIdentityProvidersRequest>): ListIdentityProvidersRequest {
    return ListIdentityProvidersRequest.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<ListIdentityProvidersRequest>): ListIdentityProvidersRequest {
    const message = createBaseListIdentityProvidersRequest();
    return message;
  },
};

function createBaseListIdentityProvidersResponse(): ListIdentityProvidersResponse {
  return { identityProviders: [] };
}

export const ListIdentityProvidersResponse: MessageFns<ListIdentityProvidersResponse> = {
  encode(message: ListIdentityProvidersResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    for (const v of message.identityProviders) {
      IdentityProvider.encode(v!, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ListIdentityProvidersResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListIdentityProvidersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.identityProviders.push(IdentityProvider.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<ListIdentityProvidersResponse>): ListIdentityProvidersResponse {
    return ListIdentityProvidersResponse.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<ListIdentityProvidersResponse>): ListIdentityProvidersResponse {
    const message = createBaseListIdentityProvidersResponse();
    message.identityProviders = object.identityProviders?.map((e) => IdentityProvider.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGetIdentityProviderRequest(): GetIdentityProviderRequest {
  return { name: "" };
}

export const GetIdentityProviderRequest: MessageFns<GetIdentityProviderRequest> = {
  encode(message: GetIdentityProviderRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetIdentityProviderRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetIdentityProviderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<GetIdentityProviderRequest>): GetIdentityProviderRequest {
    return GetIdentityProviderRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<GetIdentityProviderRequest>): GetIdentityProviderRequest {
    const message = createBaseGetIdentityProviderRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseCreateIdentityProviderRequest(): CreateIdentityProviderRequest {
  return { identityProvider: undefined };
}

export const CreateIdentityProviderRequest: MessageFns<CreateIdentityProviderRequest> = {
  encode(message: CreateIdentityProviderRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.identityProvider !== undefined) {
      IdentityProvider.encode(message.identityProvider, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateIdentityProviderRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateIdentityProviderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.identityProvider = IdentityProvider.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<CreateIdentityProviderRequest>): CreateIdentityProviderRequest {
    return CreateIdentityProviderRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<CreateIdentityProviderRequest>): CreateIdentityProviderRequest {
    const message = createBaseCreateIdentityProviderRequest();
    message.identityProvider = (object.identityProvider !== undefined && object.identityProvider !== null)
      ? IdentityProvider.fromPartial(object.identityProvider)
      : undefined;
    return message;
  },
};

function createBaseUpdateIdentityProviderRequest(): UpdateIdentityProviderRequest {
  return { identityProvider: undefined, updateMask: undefined };
}

export const UpdateIdentityProviderRequest: MessageFns<UpdateIdentityProviderRequest> = {
  encode(message: UpdateIdentityProviderRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.identityProvider !== undefined) {
      IdentityProvider.encode(message.identityProvider, writer.uint32(10).fork()).join();
    }
    if (message.updateMask !== undefined) {
      FieldMask.encode(FieldMask.wrap(message.updateMask), writer.uint32(18).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): UpdateIdentityProviderRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateIdentityProviderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.identityProvider = IdentityProvider.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.updateMask = FieldMask.unwrap(FieldMask.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<UpdateIdentityProviderRequest>): UpdateIdentityProviderRequest {
    return UpdateIdentityProviderRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<UpdateIdentityProviderRequest>): UpdateIdentityProviderRequest {
    const message = createBaseUpdateIdentityProviderRequest();
    message.identityProvider = (object.identityProvider !== undefined && object.identityProvider !== null)
      ? IdentityProvider.fromPartial(object.identityProvider)
      : undefined;
    message.updateMask = object.updateMask ?? undefined;
    return message;
  },
};

function createBaseDeleteIdentityProviderRequest(): DeleteIdentityProviderRequest {
  return { name: "" };
}

export const DeleteIdentityProviderRequest: MessageFns<DeleteIdentityProviderRequest> = {
  encode(message: DeleteIdentityProviderRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): DeleteIdentityProviderRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteIdentityProviderRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<DeleteIdentityProviderRequest>): DeleteIdentityProviderRequest {
    return DeleteIdentityProviderRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<DeleteIdentityProviderRequest>): DeleteIdentityProviderRequest {
    const message = createBaseDeleteIdentityProviderRequest();
    message.name = object.name ?? "";
    return message;
  },
};

export type IdentityProviderServiceDefinition = typeof IdentityProviderServiceDefinition;
export const IdentityProviderServiceDefinition = {
  name: "IdentityProviderService",
  fullName: "memos.api.v1.IdentityProviderService",
  methods: {
    /** ListIdentityProviders lists identity providers. */
    listIdentityProviders: {
      name: "ListIdentityProviders",
      requestType: ListIdentityProvidersRequest,
      requestStream: false,
      responseType: ListIdentityProvidersResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              27,
              18,
              25,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              105,
              100,
              101,
              110,
              116,
              105,
              116,
              121,
              80,
              114,
              111,
              118,
              105,
              100,
              101,
              114,
              115,
            ]),
          ],
        },
      },
    },
    /** GetIdentityProvider gets an identity provider. */
    getIdentityProvider: {
      name: "GetIdentityProvider",
      requestType: GetIdentityProviderRequest,
      requestStream: false,
      responseType: IdentityProvider,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([4, 110, 97, 109, 101])],
          578365826: [
            new Uint8Array([
              36,
              18,
              34,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              123,
              110,
              97,
              109,
              101,
              61,
              105,
              100,
              101,
              110,
              116,
              105,
              116,
              121,
              80,
              114,
              111,
              118,
              105,
              100,
              101,
              114,
              115,
              47,
              42,
              125,
            ]),
          ],
        },
      },
    },
    /** CreateIdentityProvider creates an identity provider. */
    createIdentityProvider: {
      name: "CreateIdentityProvider",
      requestType: CreateIdentityProviderRequest,
      requestStream: false,
      responseType: IdentityProvider,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              46,
              58,
              17,
              105,
              100,
              101,
              110,
              116,
              105,
              116,
              121,
              95,
              112,
              114,
              111,
              118,
              105,
              100,
              101,
              114,
              34,
              25,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              105,
              100,
              101,
              110,
              116,
              105,
              116,
              121,
              80,
              114,
              111,
              118,
              105,
              100,
              101,
              114,
              115,
            ]),
          ],
        },
      },
    },
    /** UpdateIdentityProvider updates an identity provider. */
    updateIdentityProvider: {
      name: "UpdateIdentityProvider",
      requestType: UpdateIdentityProviderRequest,
      requestStream: false,
      responseType: IdentityProvider,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [
            new Uint8Array([
              29,
              105,
              100,
              101,
              110,
              116,
              105,
              116,
              121,
              95,
              112,
              114,
              111,
              118,
              105,
              100,
              101,
              114,
              44,
              117,
              112,
              100,
              97,
              116,
              101,
              95,
              109,
              97,
              115,
              107,
            ]),
          ],
          578365826: [
            new Uint8Array([
              73,
              58,
              17,
              105,
              100,
              101,
              110,
              116,
              105,
              116,
              121,
              95,
              112,
              114,
              111,
              118,
              105,
              100,
              101,
              114,
              50,
              52,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              123,
              105,
              100,
              101,
              110,
              116,
              105,
              116,
              121,
              95,
              112,
              114,
              111,
              118,
              105,
              100,
              101,
              114,
              46,
              110,
              97,
              109,
              101,
              61,
              105,
              100,
              101,
              110,
              116,
              105,
              116,
              121,
              80,
              114,
              111,
              118,
              105,
              100,
              101,
              114,
              115,
              47,
              42,
              125,
            ]),
          ],
        },
      },
    },
    /** DeleteIdentityProvider deletes an identity provider. */
    deleteIdentityProvider: {
      name: "DeleteIdentityProvider",
      requestType: DeleteIdentityProviderRequest,
      requestStream: false,
      responseType: Empty,
      responseStream: false,
      options: {
        _unknownFields: {
          8410: [new Uint8Array([4, 110, 97, 109, 101])],
          578365826: [
            new Uint8Array([
              36,
              42,
              34,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              123,
              110,
              97,
              109,
              101,
              61,
              105,
              100,
              101,
              110,
              116,
              105,
              116,
              121,
              80,
              114,
              111,
              118,
              105,
              100,
              101,
              114,
              115,
              47,
              42,
              125,
            ]),
          ],
        },
      },
    },
  },
} as const;

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  create(base?: DeepPartial<T>): T;
  fromPartial(object: DeepPartial<T>): T;
}