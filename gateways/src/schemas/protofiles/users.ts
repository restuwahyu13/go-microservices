/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Any } from "../google/protobuf/any";

export const protobufPackage = "users";

export interface LoginDTO {
  email: string;
  password: string;
}

export interface RegisterDTO {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
}

export interface ApiResponse {
  statCode: number;
  statMessage: string;
  data?: Any | undefined;
}

function createBaseLoginDTO(): LoginDTO {
  return { email: "", password: "" };
}

export const LoginDTO = {
  encode(message: LoginDTO, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.email !== "") {
      writer.uint32(10).string(message.email);
    }
    if (message.password !== "") {
      writer.uint32(18).string(message.password);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginDTO {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginDTO();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.email = reader.string();
          break;
        case 2:
          message.password = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LoginDTO {
    return {
      email: isSet(object.email) ? String(object.email) : "",
      password: isSet(object.password) ? String(object.password) : "",
    };
  },

  toJSON(message: LoginDTO): unknown {
    const obj: any = {};
    message.email !== undefined && (obj.email = message.email);
    message.password !== undefined && (obj.password = message.password);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LoginDTO>, I>>(object: I): LoginDTO {
    const message = createBaseLoginDTO();
    message.email = object.email ?? "";
    message.password = object.password ?? "";
    return message;
  },
};

function createBaseRegisterDTO(): RegisterDTO {
  return { firstName: "", lastName: "", email: "", password: "" };
}

export const RegisterDTO = {
  encode(message: RegisterDTO, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.firstName !== "") {
      writer.uint32(10).string(message.firstName);
    }
    if (message.lastName !== "") {
      writer.uint32(18).string(message.lastName);
    }
    if (message.email !== "") {
      writer.uint32(34).string(message.email);
    }
    if (message.password !== "") {
      writer.uint32(42).string(message.password);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterDTO {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterDTO();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.firstName = reader.string();
          break;
        case 2:
          message.lastName = reader.string();
          break;
        case 4:
          message.email = reader.string();
          break;
        case 5:
          message.password = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegisterDTO {
    return {
      firstName: isSet(object.firstName) ? String(object.firstName) : "",
      lastName: isSet(object.lastName) ? String(object.lastName) : "",
      email: isSet(object.email) ? String(object.email) : "",
      password: isSet(object.password) ? String(object.password) : "",
    };
  },

  toJSON(message: RegisterDTO): unknown {
    const obj: any = {};
    message.firstName !== undefined && (obj.firstName = message.firstName);
    message.lastName !== undefined && (obj.lastName = message.lastName);
    message.email !== undefined && (obj.email = message.email);
    message.password !== undefined && (obj.password = message.password);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterDTO>, I>>(object: I): RegisterDTO {
    const message = createBaseRegisterDTO();
    message.firstName = object.firstName ?? "";
    message.lastName = object.lastName ?? "";
    message.email = object.email ?? "";
    message.password = object.password ?? "";
    return message;
  },
};

function createBaseApiResponse(): ApiResponse {
  return { statCode: 0, statMessage: "", data: undefined };
}

export const ApiResponse = {
  encode(message: ApiResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): ApiResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseApiResponse();
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

  fromJSON(object: any): ApiResponse {
    return {
      statCode: isSet(object.statCode) ? Number(object.statCode) : 0,
      statMessage: isSet(object.statMessage) ? String(object.statMessage) : "",
      data: isSet(object.data) ? Any.fromJSON(object.data) : undefined,
    };
  },

  toJSON(message: ApiResponse): unknown {
    const obj: any = {};
    message.statCode !== undefined && (obj.statCode = Math.round(message.statCode));
    message.statMessage !== undefined && (obj.statMessage = message.statMessage);
    message.data !== undefined && (obj.data = message.data ? Any.toJSON(message.data) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ApiResponse>, I>>(object: I): ApiResponse {
    const message = createBaseApiResponse();
    message.statCode = object.statCode ?? 0;
    message.statMessage = object.statMessage ?? "";
    message.data = (object.data !== undefined && object.data !== null) ? Any.fromPartial(object.data) : undefined;
    return message;
  },
};

export interface Users {
  LoginAuth(request: LoginDTO): Promise<ApiResponse>;
  RegisterAuth(request: RegisterDTO): Promise<ApiResponse>;
}

export class UsersClientImpl implements Users {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "users.Users";
    this.rpc = rpc;
    this.LoginAuth = this.LoginAuth.bind(this);
    this.RegisterAuth = this.RegisterAuth.bind(this);
  }
  LoginAuth(request: LoginDTO): Promise<ApiResponse> {
    const data = LoginDTO.encode(request).finish();
    const promise = this.rpc.request(this.service, "LoginAuth", data);
    return promise.then((data) => ApiResponse.decode(new _m0.Reader(data)));
  }

  RegisterAuth(request: RegisterDTO): Promise<ApiResponse> {
    const data = RegisterDTO.encode(request).finish();
    const promise = this.rpc.request(this.service, "RegisterAuth", data);
    return promise.then((data) => ApiResponse.decode(new _m0.Reader(data)));
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
