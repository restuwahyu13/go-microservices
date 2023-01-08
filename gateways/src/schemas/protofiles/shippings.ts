/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Any } from "../google/protobuf/any";
import { Empty } from "../google/protobuf/empty";

export const protobufPackage = "shippings";

export enum ServiceName {
  PingShippings = 0,
  UNRECOGNIZED = -1,
}

export function serviceNameFromJSON(object: any): ServiceName {
  switch (object) {
    case 0:
    case "PingShippings":
      return ServiceName.PingShippings;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ServiceName.UNRECOGNIZED;
  }
}

export function serviceNameToJSON(object: ServiceName): string {
  switch (object) {
    case ServiceName.PingShippings:
      return "PingShippings";
    case ServiceName.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface GrpcResponse {
  statCode: number;
  statMessage: string;
  data?: Any | undefined;
}

function createBaseGrpcResponse(): GrpcResponse {
  return { statCode: 0, statMessage: "", data: undefined };
}

export const GrpcResponse = {
  encode(message: GrpcResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.statCode !== 0) {
      writer.uint32(8).int32(message.statCode);
    }
    if (message.statMessage !== "") {
      writer.uint32(18).string(message.statMessage);
    }
    if (message.data !== undefined) {
      Any.encode(message.data, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GrpcResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGrpcResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.statCode = reader.int32();
          break;
        case 2:
          message.statMessage = reader.string();
          break;
        case 3:
          message.data = Any.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GrpcResponse {
    return {
      statCode: isSet(object.statCode) ? Number(object.statCode) : 0,
      statMessage: isSet(object.statMessage) ? String(object.statMessage) : "",
      data: isSet(object.data) ? Any.fromJSON(object.data) : undefined,
    };
  },

  toJSON(message: GrpcResponse): unknown {
    const obj: any = {};
    message.statCode !== undefined && (obj.statCode = Math.round(message.statCode));
    message.statMessage !== undefined && (obj.statMessage = message.statMessage);
    message.data !== undefined && (obj.data = message.data ? Any.toJSON(message.data) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GrpcResponse>, I>>(object: I): GrpcResponse {
    const message = createBaseGrpcResponse();
    message.statCode = object.statCode ?? 0;
    message.statMessage = object.statMessage ?? "";
    message.data = (object.data !== undefined && object.data !== null) ? Any.fromPartial(object.data) : undefined;
    return message;
  },
};

export interface ShippingsService {
  PingShippings(request: Empty): Promise<GrpcResponse>;
}

export class ShippingsServiceClientImpl implements ShippingsService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "shippings.ShippingsService";
    this.rpc = rpc;
    this.PingShippings = this.PingShippings.bind(this);
  }
  PingShippings(request: Empty): Promise<GrpcResponse> {
    const data = Empty.encode(request).finish();
    const promise = this.rpc.request(this.service, "PingShippings", data);
    return promise.then((data) => GrpcResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
