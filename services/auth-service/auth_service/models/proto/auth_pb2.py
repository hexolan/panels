"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nauth.proto\x12\x0epanels.auth.v1\x1a\x1bgoogle/protobuf/empty.proto":\n\x15SetPasswordAuthMethod\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t"+\n\x18DeletePasswordAuthMethod\x12\x0f\n\x07user_id\x18\x01 \x01(\t"8\n\x13PasswordAuthRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t"`\n\tAuthToken\x12\x12\n\ntoken_type\x18\x01 \x01(\t\x12\x14\n\x0caccess_token\x18\x02 \x01(\t\x12\x15\n\rrefresh_token\x18\x03 \x01(\t\x12\x12\n\nexpires_in\x18\x04 \x01(\x032\x91\x02\n\x0bAuthService\x12T\n\x10AuthWithPassword\x12#.panels.auth.v1.PasswordAuthRequest\x1a\x19.panels.auth.v1.AuthToken"\x00\x12R\n\x0fSetPasswordAuth\x12%.panels.auth.v1.SetPasswordAuthMethod\x1a\x16.google.protobuf.Empty"\x00\x12X\n\x12DeletePasswordAuth\x12(.panels.auth.v1.DeletePasswordAuthMethod\x1a\x16.google.protobuf.Empty"\x00b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'auth_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    DESCRIPTOR._options = None
    _globals['_SETPASSWORDAUTHMETHOD']._serialized_start = 59
    _globals['_SETPASSWORDAUTHMETHOD']._serialized_end = 117
    _globals['_DELETEPASSWORDAUTHMETHOD']._serialized_start = 119
    _globals['_DELETEPASSWORDAUTHMETHOD']._serialized_end = 162
    _globals['_PASSWORDAUTHREQUEST']._serialized_start = 164
    _globals['_PASSWORDAUTHREQUEST']._serialized_end = 220
    _globals['_AUTHTOKEN']._serialized_start = 222
    _globals['_AUTHTOKEN']._serialized_end = 318
    _globals['_AUTHSERVICE']._serialized_start = 321
    _globals['_AUTHSERVICE']._serialized_end = 594