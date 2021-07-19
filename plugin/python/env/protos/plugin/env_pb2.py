# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/plugin/env.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from protos.plugin import common_pb2 as protos_dot_plugin_dot_common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='protos/plugin/env.proto',
  package='proto',
  syntax='proto3',
  serialized_options=b'Z\036code-analyser/protos/pb/plugin',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x17protos/plugin/env.proto\x12\x05proto\x1a\x1aprotos/plugin/common.proto\"\xba\x01\n\x10ServiceOutputEnv\x12?\n\x0c\x65nvKeyValues\x18\x01 \x03(\x0b\x32).proto.ServiceOutputEnv.EnvKeyValuesEntry\x12\x0c\n\x04keys\x18\x02 \x03(\t\x12\"\n\x05\x65rror\x18\x03 \x01(\x0b\x32\x13.proto.ServiceError\x1a\x33\n\x11\x45nvKeyValuesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x32\x44\n\nEnvService\x12\x36\n\x06\x44\x65tect\x12\x13.proto.ServiceInput\x1a\x17.proto.ServiceOutputEnvB Z\x1e\x63ode-analyser/protos/pb/pluginb\x06proto3'
  ,
  dependencies=[protos_dot_plugin_dot_common__pb2.DESCRIPTOR,])




_SERVICEOUTPUTENV_ENVKEYVALUESENTRY = _descriptor.Descriptor(
  name='EnvKeyValuesEntry',
  full_name='proto.ServiceOutputEnv.EnvKeyValuesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='proto.ServiceOutputEnv.EnvKeyValuesEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='proto.ServiceOutputEnv.EnvKeyValuesEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=b'8\001',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=198,
  serialized_end=249,
)

_SERVICEOUTPUTENV = _descriptor.Descriptor(
  name='ServiceOutputEnv',
  full_name='proto.ServiceOutputEnv',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='envKeyValues', full_name='proto.ServiceOutputEnv.envKeyValues', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='keys', full_name='proto.ServiceOutputEnv.keys', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='error', full_name='proto.ServiceOutputEnv.error', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_SERVICEOUTPUTENV_ENVKEYVALUESENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=63,
  serialized_end=249,
)

_SERVICEOUTPUTENV_ENVKEYVALUESENTRY.containing_type = _SERVICEOUTPUTENV
_SERVICEOUTPUTENV.fields_by_name['envKeyValues'].message_type = _SERVICEOUTPUTENV_ENVKEYVALUESENTRY
_SERVICEOUTPUTENV.fields_by_name['error'].message_type = protos_dot_plugin_dot_common__pb2._SERVICEERROR
DESCRIPTOR.message_types_by_name['ServiceOutputEnv'] = _SERVICEOUTPUTENV
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ServiceOutputEnv = _reflection.GeneratedProtocolMessageType('ServiceOutputEnv', (_message.Message,), {

  'EnvKeyValuesEntry' : _reflection.GeneratedProtocolMessageType('EnvKeyValuesEntry', (_message.Message,), {
    'DESCRIPTOR' : _SERVICEOUTPUTENV_ENVKEYVALUESENTRY,
    '__module__' : 'protos.plugin.env_pb2'
    # @@protoc_insertion_point(class_scope:proto.ServiceOutputEnv.EnvKeyValuesEntry)
    })
  ,
  'DESCRIPTOR' : _SERVICEOUTPUTENV,
  '__module__' : 'protos.plugin.env_pb2'
  # @@protoc_insertion_point(class_scope:proto.ServiceOutputEnv)
  })
_sym_db.RegisterMessage(ServiceOutputEnv)
_sym_db.RegisterMessage(ServiceOutputEnv.EnvKeyValuesEntry)


DESCRIPTOR._options = None
_SERVICEOUTPUTENV_ENVKEYVALUESENTRY._options = None

_ENVSERVICE = _descriptor.ServiceDescriptor(
  name='EnvService',
  full_name='proto.EnvService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=251,
  serialized_end=319,
  methods=[
  _descriptor.MethodDescriptor(
    name='Detect',
    full_name='proto.EnvService.Detect',
    index=0,
    containing_service=None,
    input_type=protos_dot_plugin_dot_common__pb2._SERVICEINPUT,
    output_type=_SERVICEOUTPUTENV,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_ENVSERVICE)

DESCRIPTOR.services_by_name['EnvService'] = _ENVSERVICE

# @@protoc_insertion_point(module_scope)