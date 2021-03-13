/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var types_pb = require('./types_pb.js');
goog.exportSymbol('proto.users.CreateUserRequest', null, global);
goog.exportSymbol('proto.users.DeleteUserRequest', null, global);
goog.exportSymbol('proto.users.FindByEmailRequest', null, global);
goog.exportSymbol('proto.users.FindByIDRequest', null, global);
goog.exportSymbol('proto.users.UpdateUserRequest', null, global);
goog.exportSymbol('proto.users.UserReply', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.users.CreateUserRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.users.CreateUserRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.users.CreateUserRequest.displayName = 'proto.users.CreateUserRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.users.CreateUserRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.users.CreateUserRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.users.CreateUserRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.users.CreateUserRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    newuser: (f = msg.getNewuser()) && types_pb.CreateUser.toObject(includeInstance, f),
    jwt: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.users.CreateUserRequest}
 */
proto.users.CreateUserRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.users.CreateUserRequest;
  return proto.users.CreateUserRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.users.CreateUserRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.users.CreateUserRequest}
 */
proto.users.CreateUserRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.CreateUser;
      reader.readMessage(value,types_pb.CreateUser.deserializeBinaryFromReader);
      msg.setNewuser(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setJwt(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.users.CreateUserRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.users.CreateUserRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.users.CreateUserRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.users.CreateUserRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNewuser();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      types_pb.CreateUser.serializeBinaryToWriter
    );
  }
  f = message.getJwt();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional types.CreateUser newUser = 1;
 * @return {?proto.types.CreateUser}
 */
proto.users.CreateUserRequest.prototype.getNewuser = function() {
  return /** @type{?proto.types.CreateUser} */ (
    jspb.Message.getWrapperField(this, types_pb.CreateUser, 1));
};


/** @param {?proto.types.CreateUser|undefined} value */
proto.users.CreateUserRequest.prototype.setNewuser = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.users.CreateUserRequest.prototype.clearNewuser = function() {
  this.setNewuser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.users.CreateUserRequest.prototype.hasNewuser = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string JWT = 2;
 * @return {string}
 */
proto.users.CreateUserRequest.prototype.getJwt = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.users.CreateUserRequest.prototype.setJwt = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.users.FindByIDRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.users.FindByIDRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.users.FindByIDRequest.displayName = 'proto.users.FindByIDRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.users.FindByIDRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.users.FindByIDRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.users.FindByIDRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.users.FindByIDRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    jwt: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.users.FindByIDRequest}
 */
proto.users.FindByIDRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.users.FindByIDRequest;
  return proto.users.FindByIDRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.users.FindByIDRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.users.FindByIDRequest}
 */
proto.users.FindByIDRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setJwt(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.users.FindByIDRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.users.FindByIDRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.users.FindByIDRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.users.FindByIDRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getJwt();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional int64 id = 1;
 * @return {number}
 */
proto.users.FindByIDRequest.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.users.FindByIDRequest.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string JWT = 2;
 * @return {string}
 */
proto.users.FindByIDRequest.prototype.getJwt = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.users.FindByIDRequest.prototype.setJwt = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.users.FindByEmailRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.users.FindByEmailRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.users.FindByEmailRequest.displayName = 'proto.users.FindByEmailRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.users.FindByEmailRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.users.FindByEmailRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.users.FindByEmailRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.users.FindByEmailRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    email: jspb.Message.getFieldWithDefault(msg, 1, ""),
    jwt: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.users.FindByEmailRequest}
 */
proto.users.FindByEmailRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.users.FindByEmailRequest;
  return proto.users.FindByEmailRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.users.FindByEmailRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.users.FindByEmailRequest}
 */
proto.users.FindByEmailRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setEmail(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setJwt(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.users.FindByEmailRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.users.FindByEmailRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.users.FindByEmailRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.users.FindByEmailRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEmail();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getJwt();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string email = 1;
 * @return {string}
 */
proto.users.FindByEmailRequest.prototype.getEmail = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.users.FindByEmailRequest.prototype.setEmail = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string JWT = 2;
 * @return {string}
 */
proto.users.FindByEmailRequest.prototype.getJwt = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.users.FindByEmailRequest.prototype.setJwt = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.users.UpdateUserRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.users.UpdateUserRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.users.UpdateUserRequest.displayName = 'proto.users.UpdateUserRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.users.UpdateUserRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.users.UpdateUserRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.users.UpdateUserRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.users.UpdateUserRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    newpassword: jspb.Message.getFieldWithDefault(msg, 2, ""),
    firstname: jspb.Message.getFieldWithDefault(msg, 3, ""),
    lastname: jspb.Message.getFieldWithDefault(msg, 4, ""),
    jwt: jspb.Message.getFieldWithDefault(msg, 5, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.users.UpdateUserRequest}
 */
proto.users.UpdateUserRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.users.UpdateUserRequest;
  return proto.users.UpdateUserRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.users.UpdateUserRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.users.UpdateUserRequest}
 */
proto.users.UpdateUserRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setNewpassword(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setFirstname(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastname(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setJwt(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.users.UpdateUserRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.users.UpdateUserRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.users.UpdateUserRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.users.UpdateUserRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getNewpassword();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getFirstname();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getLastname();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getJwt();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
};


/**
 * optional int64 id = 1;
 * @return {number}
 */
proto.users.UpdateUserRequest.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.users.UpdateUserRequest.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string newPassword = 2;
 * @return {string}
 */
proto.users.UpdateUserRequest.prototype.getNewpassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.users.UpdateUserRequest.prototype.setNewpassword = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string firstName = 3;
 * @return {string}
 */
proto.users.UpdateUserRequest.prototype.getFirstname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.users.UpdateUserRequest.prototype.setFirstname = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string lastName = 4;
 * @return {string}
 */
proto.users.UpdateUserRequest.prototype.getLastname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.users.UpdateUserRequest.prototype.setLastname = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string JWT = 5;
 * @return {string}
 */
proto.users.UpdateUserRequest.prototype.getJwt = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.users.UpdateUserRequest.prototype.setJwt = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.users.DeleteUserRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.users.DeleteUserRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.users.DeleteUserRequest.displayName = 'proto.users.DeleteUserRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.users.DeleteUserRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.users.DeleteUserRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.users.DeleteUserRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.users.DeleteUserRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0),
    jwt: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.users.DeleteUserRequest}
 */
proto.users.DeleteUserRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.users.DeleteUserRequest;
  return proto.users.DeleteUserRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.users.DeleteUserRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.users.DeleteUserRequest}
 */
proto.users.DeleteUserRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setJwt(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.users.DeleteUserRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.users.DeleteUserRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.users.DeleteUserRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.users.DeleteUserRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getJwt();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional int64 id = 1;
 * @return {number}
 */
proto.users.DeleteUserRequest.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.users.DeleteUserRequest.prototype.setId = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string JWT = 2;
 * @return {string}
 */
proto.users.DeleteUserRequest.prototype.getJwt = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.users.DeleteUserRequest.prototype.setJwt = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.users.UserReply = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.users.UserReply, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.users.UserReply.displayName = 'proto.users.UserReply';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.users.UserReply.prototype.toObject = function(opt_includeInstance) {
  return proto.users.UserReply.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.users.UserReply} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.users.UserReply.toObject = function(includeInstance, msg) {
  var f, obj = {
    user: (f = msg.getUser()) && types_pb.User.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.users.UserReply}
 */
proto.users.UserReply.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.users.UserReply;
  return proto.users.UserReply.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.users.UserReply} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.users.UserReply}
 */
proto.users.UserReply.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.User;
      reader.readMessage(value,types_pb.User.deserializeBinaryFromReader);
      msg.setUser(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.users.UserReply.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.users.UserReply.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.users.UserReply} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.users.UserReply.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUser();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      types_pb.User.serializeBinaryToWriter
    );
  }
};


/**
 * optional types.User user = 1;
 * @return {?proto.types.User}
 */
proto.users.UserReply.prototype.getUser = function() {
  return /** @type{?proto.types.User} */ (
    jspb.Message.getWrapperField(this, types_pb.User, 1));
};


/** @param {?proto.types.User|undefined} value */
proto.users.UserReply.prototype.setUser = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.users.UserReply.prototype.clearUser = function() {
  this.setUser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.users.UserReply.prototype.hasUser = function() {
  return jspb.Message.getField(this, 1) != null;
};


goog.object.extend(exports, proto.users);