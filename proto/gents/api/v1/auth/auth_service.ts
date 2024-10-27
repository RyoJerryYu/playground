// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.2.0
//   protoc               unknown
// source: api/v1/auth/auth_service.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { Empty } from "../../../google/protobuf/empty";
import { User } from "../users/user_service";

export const protobufPackage = "memos.api.v1.auth";

export interface GetAuthStatusRequest {
}

export interface GetAuthStatusResponse {
  user?: User | undefined;
}

export interface SignInRequest {
  /** The username to sign in with. */
  username: string;
  /** The password to sign in with. */
  password: string;
  /** Whether the session should never expire. */
  neverExpire: boolean;
}

export interface SignInWithSSORequest {
  /** The ID of the SSO provider. */
  idpId: number;
  /** The code to sign in with. */
  code: string;
  /** The redirect URI. */
  redirectUri: string;
}

export interface SignUpRequest {
  /** The username to sign up with. */
  username: string;
  /** The password to sign up with. */
  password: string;
}

export interface SignOutRequest {
}

function createBaseGetAuthStatusRequest(): GetAuthStatusRequest {
  return {};
}

export const GetAuthStatusRequest: MessageFns<GetAuthStatusRequest> = {
  encode(_: GetAuthStatusRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetAuthStatusRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetAuthStatusRequest();
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

  create(base?: DeepPartial<GetAuthStatusRequest>): GetAuthStatusRequest {
    return GetAuthStatusRequest.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<GetAuthStatusRequest>): GetAuthStatusRequest {
    const message = createBaseGetAuthStatusRequest();
    return message;
  },
};

function createBaseGetAuthStatusResponse(): GetAuthStatusResponse {
  return { user: undefined };
}

export const GetAuthStatusResponse: MessageFns<GetAuthStatusResponse> = {
  encode(message: GetAuthStatusResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetAuthStatusResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetAuthStatusResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.user = User.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<GetAuthStatusResponse>): GetAuthStatusResponse {
    return GetAuthStatusResponse.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<GetAuthStatusResponse>): GetAuthStatusResponse {
    const message = createBaseGetAuthStatusResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseSignInRequest(): SignInRequest {
  return { username: "", password: "", neverExpire: false };
}

export const SignInRequest: MessageFns<SignInRequest> = {
  encode(message: SignInRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    if (message.password !== "") {
      writer.uint32(18).string(message.password);
    }
    if (message.neverExpire !== false) {
      writer.uint32(24).bool(message.neverExpire);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): SignInRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSignInRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.username = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.password = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.neverExpire = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<SignInRequest>): SignInRequest {
    return SignInRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<SignInRequest>): SignInRequest {
    const message = createBaseSignInRequest();
    message.username = object.username ?? "";
    message.password = object.password ?? "";
    message.neverExpire = object.neverExpire ?? false;
    return message;
  },
};

function createBaseSignInWithSSORequest(): SignInWithSSORequest {
  return { idpId: 0, code: "", redirectUri: "" };
}

export const SignInWithSSORequest: MessageFns<SignInWithSSORequest> = {
  encode(message: SignInWithSSORequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.idpId !== 0) {
      writer.uint32(8).int32(message.idpId);
    }
    if (message.code !== "") {
      writer.uint32(18).string(message.code);
    }
    if (message.redirectUri !== "") {
      writer.uint32(26).string(message.redirectUri);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): SignInWithSSORequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSignInWithSSORequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.idpId = reader.int32();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.code = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.redirectUri = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<SignInWithSSORequest>): SignInWithSSORequest {
    return SignInWithSSORequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<SignInWithSSORequest>): SignInWithSSORequest {
    const message = createBaseSignInWithSSORequest();
    message.idpId = object.idpId ?? 0;
    message.code = object.code ?? "";
    message.redirectUri = object.redirectUri ?? "";
    return message;
  },
};

function createBaseSignUpRequest(): SignUpRequest {
  return { username: "", password: "" };
}

export const SignUpRequest: MessageFns<SignUpRequest> = {
  encode(message: SignUpRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.username !== "") {
      writer.uint32(10).string(message.username);
    }
    if (message.password !== "") {
      writer.uint32(18).string(message.password);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): SignUpRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSignUpRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.username = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.password = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<SignUpRequest>): SignUpRequest {
    return SignUpRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<SignUpRequest>): SignUpRequest {
    const message = createBaseSignUpRequest();
    message.username = object.username ?? "";
    message.password = object.password ?? "";
    return message;
  },
};

function createBaseSignOutRequest(): SignOutRequest {
  return {};
}

export const SignOutRequest: MessageFns<SignOutRequest> = {
  encode(_: SignOutRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): SignOutRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSignOutRequest();
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

  create(base?: DeepPartial<SignOutRequest>): SignOutRequest {
    return SignOutRequest.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<SignOutRequest>): SignOutRequest {
    const message = createBaseSignOutRequest();
    return message;
  },
};

export type AuthServiceDefinition = typeof AuthServiceDefinition;
export const AuthServiceDefinition = {
  name: "AuthService",
  fullName: "memos.api.v1.auth.AuthService",
  methods: {
    /** GetAuthStatus returns the current auth status of the user. */
    getAuthStatus: {
      name: "GetAuthStatus",
      requestType: GetAuthStatusRequest,
      requestStream: false,
      responseType: User,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              21,
              34,
              19,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              97,
              117,
              116,
              104,
              47,
              115,
              116,
              97,
              116,
              117,
              115,
            ]),
          ],
        },
      },
    },
    /** SignIn signs in the user with the given username and password. */
    signIn: {
      name: "SignIn",
      requestType: SignInRequest,
      requestStream: false,
      responseType: User,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              21,
              34,
              19,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              97,
              117,
              116,
              104,
              47,
              115,
              105,
              103,
              110,
              105,
              110,
            ]),
          ],
        },
      },
    },
    /** SignInWithSSO signs in the user with the given SSO code. */
    signInWithSSO: {
      name: "SignInWithSSO",
      requestType: SignInWithSSORequest,
      requestStream: false,
      responseType: User,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              25,
              34,
              23,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              97,
              117,
              116,
              104,
              47,
              115,
              105,
              103,
              110,
              105,
              110,
              47,
              115,
              115,
              111,
            ]),
          ],
        },
      },
    },
    /** SignUp signs up the user with the given username and password. */
    signUp: {
      name: "SignUp",
      requestType: SignUpRequest,
      requestStream: false,
      responseType: User,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              21,
              34,
              19,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              97,
              117,
              116,
              104,
              47,
              115,
              105,
              103,
              110,
              117,
              112,
            ]),
          ],
        },
      },
    },
    /** SignOut signs out the user. */
    signOut: {
      name: "SignOut",
      requestType: SignOutRequest,
      requestStream: false,
      responseType: Empty,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              22,
              34,
              20,
              47,
              97,
              112,
              105,
              47,
              118,
              49,
              47,
              97,
              117,
              116,
              104,
              47,
              115,
              105,
              103,
              110,
              111,
              117,
              116,
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