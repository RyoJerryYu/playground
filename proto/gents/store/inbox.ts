// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.2.0
//   protoc               unknown
// source: store/inbox.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "memos.store";

export interface InboxMessage {
}

export enum InboxMessage_Type {
  TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED",
  MEMO_COMMENT = "MEMO_COMMENT",
  VERSION_UPDATE = "VERSION_UPDATE",
  UNRECOGNIZED = "UNRECOGNIZED",
}

export function inboxMessage_TypeFromJSON(object: any): InboxMessage_Type {
  switch (object) {
    case 0:
    case "TYPE_UNSPECIFIED":
      return InboxMessage_Type.TYPE_UNSPECIFIED;
    case 1:
    case "MEMO_COMMENT":
      return InboxMessage_Type.MEMO_COMMENT;
    case 2:
    case "VERSION_UPDATE":
      return InboxMessage_Type.VERSION_UPDATE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return InboxMessage_Type.UNRECOGNIZED;
  }
}

export function inboxMessage_TypeToNumber(object: InboxMessage_Type): number {
  switch (object) {
    case InboxMessage_Type.TYPE_UNSPECIFIED:
      return 0;
    case InboxMessage_Type.MEMO_COMMENT:
      return 1;
    case InboxMessage_Type.VERSION_UPDATE:
      return 2;
    case InboxMessage_Type.UNRECOGNIZED:
    default:
      return -1;
  }
}

function createBaseInboxMessage(): InboxMessage {
  return {};
}

export const InboxMessage: MessageFns<InboxMessage> = {
  encode(_: InboxMessage, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): InboxMessage {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInboxMessage();
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

  create(base?: DeepPartial<InboxMessage>): InboxMessage {
    return InboxMessage.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<InboxMessage>): InboxMessage {
    const message = createBaseInboxMessage();
    return message;
  },
};

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
