/**
 * @fileoverview gRPC-Web generated client stub for club.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_field_mask_pb from 'google-protobuf/google/protobuf/field_mask_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';

import {
  ChangePasswordRequest,
  ConfirmMailRequest,
  ContactRequest,
  ForgotPasswordRequest,
  Group,
  IdRequest,
  ListGroupsResponse,
  ListMembersResponse,
  ListRequest,
  Member,
  ResetPasswordRequest,
  UpdateGroupRequest,
  UpdateMemberRequest} from './club_service_pb';

export class ClubClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoContact = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_empty_pb.Empty,
    (request: ContactRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  contact(
    request: ContactRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/club.v1.Club/Contact',
      request,
      metadata || {},
      this.methodInfoContact,
      callback);
  }

  methodInfoListMembers = new grpcWeb.AbstractClientBase.MethodInfo(
    ListMembersResponse,
    (request: ListRequest) => {
      return request.serializeBinary();
    },
    ListMembersResponse.deserializeBinary
  );

  listMembers(
    request: ListRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ListMembersResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/club.v1.Club/ListMembers',
      request,
      metadata || {},
      this.methodInfoListMembers,
      callback);
  }

  methodInfoGetMember = new grpcWeb.AbstractClientBase.MethodInfo(
    Member,
    (request: IdRequest) => {
      return request.serializeBinary();
    },
    Member.deserializeBinary
  );

  getMember(
    request: IdRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Member) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/club.v1.Club/GetMember',
      request,
      metadata || {},
      this.methodInfoGetMember,
      callback);
  }

  methodInfoCreateMember = new grpcWeb.AbstractClientBase.MethodInfo(
    Member,
    (request: Member) => {
      return request.serializeBinary();
    },
    Member.deserializeBinary
  );

  createMember(
    request: Member,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Member) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/club.v1.Club/CreateMember',
      request,
      metadata || {},
      this.methodInfoCreateMember,
      callback);
  }

  methodInfoUpdateMember = new grpcWeb.AbstractClientBase.MethodInfo(
    Member,
    (request: UpdateMemberRequest) => {
      return request.serializeBinary();
    },
    Member.deserializeBinary
  );

  updateMember(
    request: UpdateMemberRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Member) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/club.v1.Club/UpdateMember',
      request,
      metadata || {},
      this.methodInfoUpdateMember,
      callback);
  }

  methodInfoDeleteMember = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_empty_pb.Empty,
    (request: IdRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  deleteMember(
    request: IdRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/club.v1.Club/DeleteMember',
      request,
      metadata || {},
      this.methodInfoDeleteMember,
      callback);
  }

  methodInfoChangePassword = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_empty_pb.Empty,
    (request: ChangePasswordRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  changePassword(
    request: ChangePasswordRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/club.v1.Club/ChangePassword',
      request,
      metadata || {},
      this.methodInfoChangePassword,
      callback);
  }

  methodInfoConfirmMail = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_empty_pb.Empty,
    (request: ConfirmMailRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  confirmMail(
    request: ConfirmMailRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/club.v1.Club/ConfirmMail',
      request,
      metadata || {},
      this.methodInfoConfirmMail,
      callback);
  }

  methodInfoForgotPassword = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_empty_pb.Empty,
    (request: ForgotPasswordRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  forgotPassword(
    request: ForgotPasswordRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/club.v1.Club/ForgotPassword',
      request,
      metadata || {},
      this.methodInfoForgotPassword,
      callback);
  }

  methodInfoResetPassword = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_empty_pb.Empty,
    (request: ResetPasswordRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  resetPassword(
    request: ResetPasswordRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/club.v1.Club/ResetPassword',
      request,
      metadata || {},
      this.methodInfoResetPassword,
      callback);
  }

  methodInfoListGroups = new grpcWeb.AbstractClientBase.MethodInfo(
    ListGroupsResponse,
    (request: ListRequest) => {
      return request.serializeBinary();
    },
    ListGroupsResponse.deserializeBinary
  );

  listGroups(
    request: ListRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ListGroupsResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/club.v1.Club/ListGroups',
      request,
      metadata || {},
      this.methodInfoListGroups,
      callback);
  }

  methodInfoGetGroup = new grpcWeb.AbstractClientBase.MethodInfo(
    Group,
    (request: IdRequest) => {
      return request.serializeBinary();
    },
    Group.deserializeBinary
  );

  getGroup(
    request: IdRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Group) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/club.v1.Club/GetGroup',
      request,
      metadata || {},
      this.methodInfoGetGroup,
      callback);
  }

  methodInfoCreateGroup = new grpcWeb.AbstractClientBase.MethodInfo(
    Group,
    (request: Group) => {
      return request.serializeBinary();
    },
    Group.deserializeBinary
  );

  createGroup(
    request: Group,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Group) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/club.v1.Club/CreateGroup',
      request,
      metadata || {},
      this.methodInfoCreateGroup,
      callback);
  }

  methodInfoUpdateGroup = new grpcWeb.AbstractClientBase.MethodInfo(
    Group,
    (request: UpdateGroupRequest) => {
      return request.serializeBinary();
    },
    Group.deserializeBinary
  );

  updateGroup(
    request: UpdateGroupRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: Group) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/club.v1.Club/UpdateGroup',
      request,
      metadata || {},
      this.methodInfoUpdateGroup,
      callback);
  }

  methodInfoDeleteGroup = new grpcWeb.AbstractClientBase.MethodInfo(
    google_protobuf_empty_pb.Empty,
    (request: IdRequest) => {
      return request.serializeBinary();
    },
    google_protobuf_empty_pb.Empty.deserializeBinary
  );

  deleteGroup(
    request: IdRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: google_protobuf_empty_pb.Empty) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/club.v1.Club/DeleteGroup',
      request,
      metadata || {},
      this.methodInfoDeleteGroup,
      callback);
  }

}

