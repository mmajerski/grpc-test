/**
 * @fileoverview gRPC-Web generated client stub for users
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var types_pb = require('./types_pb.js')
const proto = {};
proto.users = require('./users_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.users.UsersClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.users.UsersPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.users.CreateUserRequest,
 *   !proto.users.UserReply>}
 */
const methodDescriptor_Users_Create = new grpc.web.MethodDescriptor(
  '/users.Users/Create',
  grpc.web.MethodType.UNARY,
  proto.users.CreateUserRequest,
  proto.users.UserReply,
  /**
   * @param {!proto.users.CreateUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.users.CreateUserRequest,
 *   !proto.users.UserReply>}
 */
const methodInfo_Users_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.users.UserReply,
  /**
   * @param {!proto.users.CreateUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @param {!proto.users.CreateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.users.UserReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.users.UserReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.users.UsersClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/users.Users/Create',
      request,
      metadata || {},
      methodDescriptor_Users_Create,
      callback);
};


/**
 * @param {!proto.users.CreateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.users.UserReply>}
 *     Promise that resolves to the response
 */
proto.users.UsersPromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/users.Users/Create',
      request,
      metadata || {},
      methodDescriptor_Users_Create);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.users.FindByIDRequest,
 *   !proto.users.UserReply>}
 */
const methodDescriptor_Users_FindByID = new grpc.web.MethodDescriptor(
  '/users.Users/FindByID',
  grpc.web.MethodType.UNARY,
  proto.users.FindByIDRequest,
  proto.users.UserReply,
  /**
   * @param {!proto.users.FindByIDRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.users.FindByIDRequest,
 *   !proto.users.UserReply>}
 */
const methodInfo_Users_FindByID = new grpc.web.AbstractClientBase.MethodInfo(
  proto.users.UserReply,
  /**
   * @param {!proto.users.FindByIDRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @param {!proto.users.FindByIDRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.users.UserReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.users.UserReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.users.UsersClient.prototype.findByID =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/users.Users/FindByID',
      request,
      metadata || {},
      methodDescriptor_Users_FindByID,
      callback);
};


/**
 * @param {!proto.users.FindByIDRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.users.UserReply>}
 *     Promise that resolves to the response
 */
proto.users.UsersPromiseClient.prototype.findByID =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/users.Users/FindByID',
      request,
      metadata || {},
      methodDescriptor_Users_FindByID);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.users.FindByEmailRequest,
 *   !proto.users.UserReply>}
 */
const methodDescriptor_Users_FindByEmail = new grpc.web.MethodDescriptor(
  '/users.Users/FindByEmail',
  grpc.web.MethodType.UNARY,
  proto.users.FindByEmailRequest,
  proto.users.UserReply,
  /**
   * @param {!proto.users.FindByEmailRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.users.FindByEmailRequest,
 *   !proto.users.UserReply>}
 */
const methodInfo_Users_FindByEmail = new grpc.web.AbstractClientBase.MethodInfo(
  proto.users.UserReply,
  /**
   * @param {!proto.users.FindByEmailRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @param {!proto.users.FindByEmailRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.users.UserReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.users.UserReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.users.UsersClient.prototype.findByEmail =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/users.Users/FindByEmail',
      request,
      metadata || {},
      methodDescriptor_Users_FindByEmail,
      callback);
};


/**
 * @param {!proto.users.FindByEmailRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.users.UserReply>}
 *     Promise that resolves to the response
 */
proto.users.UsersPromiseClient.prototype.findByEmail =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/users.Users/FindByEmail',
      request,
      metadata || {},
      methodDescriptor_Users_FindByEmail);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.users.UpdateUserRequest,
 *   !proto.users.UserReply>}
 */
const methodDescriptor_Users_Update = new grpc.web.MethodDescriptor(
  '/users.Users/Update',
  grpc.web.MethodType.UNARY,
  proto.users.UpdateUserRequest,
  proto.users.UserReply,
  /**
   * @param {!proto.users.UpdateUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.users.UpdateUserRequest,
 *   !proto.users.UserReply>}
 */
const methodInfo_Users_Update = new grpc.web.AbstractClientBase.MethodInfo(
  proto.users.UserReply,
  /**
   * @param {!proto.users.UpdateUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @param {!proto.users.UpdateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.users.UserReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.users.UserReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.users.UsersClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/users.Users/Update',
      request,
      metadata || {},
      methodDescriptor_Users_Update,
      callback);
};


/**
 * @param {!proto.users.UpdateUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.users.UserReply>}
 *     Promise that resolves to the response
 */
proto.users.UsersPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/users.Users/Update',
      request,
      metadata || {},
      methodDescriptor_Users_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.users.DeleteUserRequest,
 *   !proto.users.UserReply>}
 */
const methodDescriptor_Users_Delete = new grpc.web.MethodDescriptor(
  '/users.Users/Delete',
  grpc.web.MethodType.UNARY,
  proto.users.DeleteUserRequest,
  proto.users.UserReply,
  /**
   * @param {!proto.users.DeleteUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.users.DeleteUserRequest,
 *   !proto.users.UserReply>}
 */
const methodInfo_Users_Delete = new grpc.web.AbstractClientBase.MethodInfo(
  proto.users.UserReply,
  /**
   * @param {!proto.users.DeleteUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.UserReply.deserializeBinary
);


/**
 * @param {!proto.users.DeleteUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.users.UserReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.users.UserReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.users.UsersClient.prototype.delete =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/users.Users/Delete',
      request,
      metadata || {},
      methodDescriptor_Users_Delete,
      callback);
};


/**
 * @param {!proto.users.DeleteUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.users.UserReply>}
 *     Promise that resolves to the response
 */
proto.users.UsersPromiseClient.prototype.delete =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/users.Users/Delete',
      request,
      metadata || {},
      methodDescriptor_Users_Delete);
};


module.exports = proto.users;

