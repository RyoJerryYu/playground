// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.2.0
//   protoc               unknown
// source: store/resource.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { Timestamp } from "../google/protobuf/timestamp";
import { StorageS3Config } from "./workspace_setting";

export const protobufPackage = "memos.store";

export enum ResourceStorageType {
  RESOURCE_STORAGE_TYPE_UNSPECIFIED = "RESOURCE_STORAGE_TYPE_UNSPECIFIED",
  LOCAL = "LOCAL",
  S3 = "S3",
  EXTERNAL = "EXTERNAL",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function resourceStorageTypeFromJSON(object: any): ResourceStorageType {
  switch (object) {
    case 0:
    case "RESOURCE_STORAGE_TYPE_UNSPECIFIED":
      return ResourceStorageType.RESOURCE_STORAGE_TYPE_UNSPECIFIED;
    case 1:
    case "LOCAL":
      return ResourceStorageType.LOCAL;
    case 2:
    case "S3":
      return ResourceStorageType.S3;
    case 3:
    case "EXTERNAL":
      return ResourceStorageType.EXTERNAL;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ResourceStorageType.UNRECOGNIZED;
  }
}

export function resourceStorageTypeToNumber(object: ResourceStorageType): number {
  switch (object) {
    case ResourceStorageType.RESOURCE_STORAGE_TYPE_UNSPECIFIED:
      return 0;
    case ResourceStorageType.LOCAL:
      return 1;
    case ResourceStorageType.S3:
      return 2;
    case ResourceStorageType.EXTERNAL:
      return 3;
    case ResourceStorageType.UNRECOGNIZED:
    default:
      return -1;
  }
}

export interface ResourcePayload {
  s3Object?: ResourcePayload_S3Object | undefined;
}

export interface ResourcePayload_S3Object {
  s3Config?:
    | StorageS3Config
    | undefined;
  /** key is the S3 object key. */
  key: string;
  /**
   * last_presigned_time is the last time the object was presigned.
   * This is used to determine if the presigned URL is still valid.
   */
  lastPresignedTime?: Date | undefined;
}

function createBaseResourcePayload(): ResourcePayload {
  return { s3Object: undefined };
}

export const ResourcePayload: MessageFns<ResourcePayload> = {
  encode(message: ResourcePayload, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.s3Object !== undefined) {
      ResourcePayload_S3Object.encode(message.s3Object, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ResourcePayload {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseResourcePayload();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.s3Object = ResourcePayload_S3Object.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<ResourcePayload>): ResourcePayload {
    return ResourcePayload.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<ResourcePayload>): ResourcePayload {
    const message = createBaseResourcePayload();
    message.s3Object = (object.s3Object !== undefined && object.s3Object !== null)
      ? ResourcePayload_S3Object.fromPartial(object.s3Object)
      : undefined;
    return message;
  },
};

function createBaseResourcePayload_S3Object(): ResourcePayload_S3Object {
  return { s3Config: undefined, key: "", lastPresignedTime: undefined };
}

export const ResourcePayload_S3Object: MessageFns<ResourcePayload_S3Object> = {
  encode(message: ResourcePayload_S3Object, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.s3Config !== undefined) {
      StorageS3Config.encode(message.s3Config, writer.uint32(10).fork()).join();
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    if (message.lastPresignedTime !== undefined) {
      Timestamp.encode(toTimestamp(message.lastPresignedTime), writer.uint32(26).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ResourcePayload_S3Object {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseResourcePayload_S3Object();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.s3Config = StorageS3Config.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.key = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.lastPresignedTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  create(base?: DeepPartial<ResourcePayload_S3Object>): ResourcePayload_S3Object {
    return ResourcePayload_S3Object.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<ResourcePayload_S3Object>): ResourcePayload_S3Object {
    const message = createBaseResourcePayload_S3Object();
    message.s3Config = (object.s3Config !== undefined && object.s3Config !== null)
      ? StorageS3Config.fromPartial(object.s3Config)
      : undefined;
    message.key = object.key ?? "";
    message.lastPresignedTime = object.lastPresignedTime ?? undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function toTimestamp(date: Date): Timestamp {
  const seconds = Math.trunc(date.getTime() / 1_000);
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new globalThis.Date(millis);
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  create(base?: DeepPartial<T>): T;
  fromPartial(object: DeepPartial<T>): T;
}