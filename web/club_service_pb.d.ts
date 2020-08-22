import * as jspb from "google-protobuf"

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_field_mask_pb from 'google-protobuf/google/protobuf/field_mask_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';

export class IdRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IdRequest.AsObject;
  static toObject(includeInstance: boolean, msg: IdRequest): IdRequest.AsObject;
  static serializeBinaryToWriter(message: IdRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IdRequest;
  static deserializeBinaryFromReader(message: IdRequest, reader: jspb.BinaryReader): IdRequest;
}

export namespace IdRequest {
  export type AsObject = {
    id: string,
  }
}

export class ListRequest extends jspb.Message {
  getPageSize(): number;
  setPageSize(value: number): void;

  getPageToken(): string;
  setPageToken(value: string): void;

  getFilter(): string;
  setFilter(value: string): void;

  getOrderBy(): string;
  setOrderBy(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRequest): ListRequest.AsObject;
  static serializeBinaryToWriter(message: ListRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRequest;
  static deserializeBinaryFromReader(message: ListRequest, reader: jspb.BinaryReader): ListRequest;
}

export namespace ListRequest {
  export type AsObject = {
    pageSize: number,
    pageToken: string,
    filter: string,
    orderBy: string,
  }
}

export class Member extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasCreatedAt(): boolean;
  clearCreatedAt(): void;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;

  getUserId(): string;
  setUserId(value: string): void;

  getUsername(): string;
  setUsername(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  getMail(): string;
  setMail(value: string): void;

  getFirstName(): string;
  setFirstName(value: string): void;

  getLastName(): string;
  setLastName(value: string): void;

  getDateOfBirth(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDateOfBirth(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasDateOfBirth(): boolean;
  clearDateOfBirth(): void;

  getPhone(): string;
  setPhone(value: string): void;

  getAddress(): string;
  setAddress(value: string): void;

  getAddress2(): string;
  setAddress2(value: string): void;

  getPostalCode(): string;
  setPostalCode(value: string): void;

  getCity(): string;
  setCity(value: string): void;

  getLanguage(): string;
  setLanguage(value: string): void;

  getJuristic(): boolean;
  setJuristic(value: boolean): void;

  getOrganisation(): string;
  setOrganisation(value: string): void;

  getWebsite(): string;
  setWebsite(value: string): void;

  getOrganisationType(): string;
  setOrganisationType(value: string): void;

  getInactive(): boolean;
  setInactive(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Member.AsObject;
  static toObject(includeInstance: boolean, msg: Member): Member.AsObject;
  static serializeBinaryToWriter(message: Member, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Member;
  static deserializeBinaryFromReader(message: Member, reader: jspb.BinaryReader): Member;
}

export namespace Member {
  export type AsObject = {
    id: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    userId: string,
    username: string,
    password: string,
    mail: string,
    firstName: string,
    lastName: string,
    dateOfBirth?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    phone: string,
    address: string,
    address2: string,
    postalCode: string,
    city: string,
    language: string,
    juristic: boolean,
    organisation: string,
    website: string,
    organisationType: string,
    inactive: boolean,
  }
}

export class ListMembersResponse extends jspb.Message {
  getMembersList(): Array<Member>;
  setMembersList(value: Array<Member>): void;
  clearMembersList(): void;
  addMembers(value?: Member, index?: number): Member;

  getNextPageToken(): string;
  setNextPageToken(value: string): void;

  getPageSize(): number;
  setPageSize(value: number): void;

  getTotalSize(): number;
  setTotalSize(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListMembersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListMembersResponse): ListMembersResponse.AsObject;
  static serializeBinaryToWriter(message: ListMembersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListMembersResponse;
  static deserializeBinaryFromReader(message: ListMembersResponse, reader: jspb.BinaryReader): ListMembersResponse;
}

export namespace ListMembersResponse {
  export type AsObject = {
    membersList: Array<Member.AsObject>,
    nextPageToken: string,
    pageSize: number,
    totalSize: number,
  }
}

export class UpdateMemberRequest extends jspb.Message {
  getMember(): Member | undefined;
  setMember(value?: Member): void;
  hasMember(): boolean;
  clearMember(): void;

  getFieldMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
  setFieldMask(value?: google_protobuf_field_mask_pb.FieldMask): void;
  hasFieldMask(): boolean;
  clearFieldMask(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateMemberRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateMemberRequest): UpdateMemberRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateMemberRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateMemberRequest;
  static deserializeBinaryFromReader(message: UpdateMemberRequest, reader: jspb.BinaryReader): UpdateMemberRequest;
}

export namespace UpdateMemberRequest {
  export type AsObject = {
    member?: Member.AsObject,
    fieldMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
  }
}

export class User extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasCreatedAt(): boolean;
  clearCreatedAt(): void;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;

  getUsername(): string;
  setUsername(value: string): void;

  getMail(): string;
  setMail(value: string): void;

  getLanguage(): string;
  setLanguage(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  getConfirmed(): boolean;
  setConfirmed(value: boolean): void;

  getRolesList(): Array<string>;
  setRolesList(value: Array<string>): void;
  clearRolesList(): void;
  addRoles(value: string, index?: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    username: string,
    mail: string,
    language: string,
    password: string,
    confirmed: boolean,
    rolesList: Array<string>,
  }
}

export class UpdateUserRequest extends jspb.Message {
  getUser(): User | undefined;
  setUser(value?: User): void;
  hasUser(): boolean;
  clearUser(): void;

  getFieldMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
  setFieldMask(value?: google_protobuf_field_mask_pb.FieldMask): void;
  hasFieldMask(): boolean;
  clearFieldMask(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUserRequest): UpdateUserRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUserRequest;
  static deserializeBinaryFromReader(message: UpdateUserRequest, reader: jspb.BinaryReader): UpdateUserRequest;
}

export namespace UpdateUserRequest {
  export type AsObject = {
    user?: User.AsObject,
    fieldMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
  }
}

export class ChangePasswordRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getOldPassword(): string;
  setOldPassword(value: string): void;

  getNewPassword(): string;
  setNewPassword(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChangePasswordRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ChangePasswordRequest): ChangePasswordRequest.AsObject;
  static serializeBinaryToWriter(message: ChangePasswordRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChangePasswordRequest;
  static deserializeBinaryFromReader(message: ChangePasswordRequest, reader: jspb.BinaryReader): ChangePasswordRequest;
}

export namespace ChangePasswordRequest {
  export type AsObject = {
    id: string,
    oldPassword: string,
    newPassword: string,
  }
}

export class ConfirmMailRequest extends jspb.Message {
  getToken(): string;
  setToken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ConfirmMailRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ConfirmMailRequest): ConfirmMailRequest.AsObject;
  static serializeBinaryToWriter(message: ConfirmMailRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ConfirmMailRequest;
  static deserializeBinaryFromReader(message: ConfirmMailRequest, reader: jspb.BinaryReader): ConfirmMailRequest;
}

export namespace ConfirmMailRequest {
  export type AsObject = {
    token: string,
  }
}

export class ForgotPasswordRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): void;

  getMail(): string;
  setMail(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ForgotPasswordRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ForgotPasswordRequest): ForgotPasswordRequest.AsObject;
  static serializeBinaryToWriter(message: ForgotPasswordRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ForgotPasswordRequest;
  static deserializeBinaryFromReader(message: ForgotPasswordRequest, reader: jspb.BinaryReader): ForgotPasswordRequest;
}

export namespace ForgotPasswordRequest {
  export type AsObject = {
    username: string,
    mail: string,
  }
}

export class ResetPasswordRequest extends jspb.Message {
  getToken(): string;
  setToken(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ResetPasswordRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ResetPasswordRequest): ResetPasswordRequest.AsObject;
  static serializeBinaryToWriter(message: ResetPasswordRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ResetPasswordRequest;
  static deserializeBinaryFromReader(message: ResetPasswordRequest, reader: jspb.BinaryReader): ResetPasswordRequest;
}

export namespace ResetPasswordRequest {
  export type AsObject = {
    token: string,
    password: string,
  }
}

export class Group extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasCreatedAt(): boolean;
  clearCreatedAt(): void;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): void;

  getName(): string;
  setName(value: string): void;

  getRolesList(): Array<string>;
  setRolesList(value: Array<string>): void;
  clearRolesList(): void;
  addRoles(value: string, index?: number): void;

  getMembersList(): Array<string>;
  setMembersList(value: Array<string>): void;
  clearMembersList(): void;
  addMembers(value: string, index?: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Group.AsObject;
  static toObject(includeInstance: boolean, msg: Group): Group.AsObject;
  static serializeBinaryToWriter(message: Group, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Group;
  static deserializeBinaryFromReader(message: Group, reader: jspb.BinaryReader): Group;
}

export namespace Group {
  export type AsObject = {
    id: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    name: string,
    rolesList: Array<string>,
    membersList: Array<string>,
  }
}

export class UpdateGroupRequest extends jspb.Message {
  getGroup(): Group | undefined;
  setGroup(value?: Group): void;
  hasGroup(): boolean;
  clearGroup(): void;

  getFieldMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
  setFieldMask(value?: google_protobuf_field_mask_pb.FieldMask): void;
  hasFieldMask(): boolean;
  clearFieldMask(): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateGroupRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateGroupRequest): UpdateGroupRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateGroupRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateGroupRequest;
  static deserializeBinaryFromReader(message: UpdateGroupRequest, reader: jspb.BinaryReader): UpdateGroupRequest;
}

export namespace UpdateGroupRequest {
  export type AsObject = {
    group?: Group.AsObject,
    fieldMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
  }
}

export class ListGroupsResponse extends jspb.Message {
  getGroupsList(): Array<Group>;
  setGroupsList(value: Array<Group>): void;
  clearGroupsList(): void;
  addGroups(value?: Group, index?: number): Group;

  getNextPageToken(): string;
  setNextPageToken(value: string): void;

  getPageSize(): number;
  setPageSize(value: number): void;

  getTotalSize(): number;
  setTotalSize(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListGroupsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListGroupsResponse): ListGroupsResponse.AsObject;
  static serializeBinaryToWriter(message: ListGroupsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListGroupsResponse;
  static deserializeBinaryFromReader(message: ListGroupsResponse, reader: jspb.BinaryReader): ListGroupsResponse;
}

export namespace ListGroupsResponse {
  export type AsObject = {
    groupsList: Array<Group.AsObject>,
    nextPageToken: string,
    pageSize: number,
    totalSize: number,
  }
}

