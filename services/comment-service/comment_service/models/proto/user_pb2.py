"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nuser.proto\x12\x0epanels.user.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto"\x96\x01\n\x04User\x12\n\n\x02id\x18\x01 \x01(\t\x12\x10\n\x08username\x18\x02 \x01(\t\x12\x10\n\x08is_admin\x18\x03 \x01(\x08\x12.\n\ncreated_at\x18\x04 \x01(\x0b2\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x05 \x01(\x0b2\x1a.google.protobuf.Timestamp"1\n\x0bUserMutable\x12\x15\n\x08username\x18\x01 \x01(\tH\x00\x88\x01\x01B\x0b\n\t_username">\n\x11CreateUserRequest\x12)\n\x04data\x18\x01 \x01(\x0b2\x1b.panels.user.v1.UserMutable" \n\x12GetUserByIdRequest\x12\n\n\x02id\x18\x01 \x01(\t"(\n\x14GetUserByNameRequest\x12\x10\n\x08username\x18\x01 \x01(\t"N\n\x15UpdateUserByIdRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12)\n\x04data\x18\x02 \x01(\x0b2\x1b.panels.user.v1.UserMutable"V\n\x17UpdateUserByNameRequest\x12\x10\n\x08username\x18\x01 \x01(\t\x12)\n\x04data\x18\x02 \x01(\x0b2\x1b.panels.user.v1.UserMutable"#\n\x15DeleteUserByIdRequest\x12\n\n\x02id\x18\x01 \x01(\t"+\n\x17DeleteUserByNameRequest\x12\x10\n\x08username\x18\x01 \x01(\t"=\n\tUserEvent\x12\x0c\n\x04type\x18\x01 \x01(\t\x12"\n\x04data\x18\x02 \x01(\x0b2\x14.panels.user.v1.User2\xb4\x04\n\x0bUserService\x12G\n\nCreateUser\x12!.panels.user.v1.CreateUserRequest\x1a\x14.panels.user.v1.User"\x00\x12E\n\x07GetUser\x12".panels.user.v1.GetUserByIdRequest\x1a\x14.panels.user.v1.User"\x00\x12M\n\rGetUserByName\x12$.panels.user.v1.GetUserByNameRequest\x1a\x14.panels.user.v1.User"\x00\x12K\n\nUpdateUser\x12%.panels.user.v1.UpdateUserByIdRequest\x1a\x14.panels.user.v1.User"\x00\x12S\n\x10UpdateUserByName\x12\'.panels.user.v1.UpdateUserByNameRequest\x1a\x14.panels.user.v1.User"\x00\x12M\n\nDeleteUser\x12%.panels.user.v1.DeleteUserByIdRequest\x1a\x16.google.protobuf.Empty"\x00\x12U\n\x10DeleteUserByName\x12\'.panels.user.v1.DeleteUserByNameRequest\x1a\x16.google.protobuf.Empty"\x00b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'user_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    DESCRIPTOR._options = None
    _globals['_USER']._serialized_start = 93
    _globals['_USER']._serialized_end = 243
    _globals['_USERMUTABLE']._serialized_start = 245
    _globals['_USERMUTABLE']._serialized_end = 294
    _globals['_CREATEUSERREQUEST']._serialized_start = 296
    _globals['_CREATEUSERREQUEST']._serialized_end = 358
    _globals['_GETUSERBYIDREQUEST']._serialized_start = 360
    _globals['_GETUSERBYIDREQUEST']._serialized_end = 392
    _globals['_GETUSERBYNAMEREQUEST']._serialized_start = 394
    _globals['_GETUSERBYNAMEREQUEST']._serialized_end = 434
    _globals['_UPDATEUSERBYIDREQUEST']._serialized_start = 436
    _globals['_UPDATEUSERBYIDREQUEST']._serialized_end = 514
    _globals['_UPDATEUSERBYNAMEREQUEST']._serialized_start = 516
    _globals['_UPDATEUSERBYNAMEREQUEST']._serialized_end = 602
    _globals['_DELETEUSERBYIDREQUEST']._serialized_start = 604
    _globals['_DELETEUSERBYIDREQUEST']._serialized_end = 639
    _globals['_DELETEUSERBYNAMEREQUEST']._serialized_start = 641
    _globals['_DELETEUSERBYNAMEREQUEST']._serialized_end = 684
    _globals['_USEREVENT']._serialized_start = 686
    _globals['_USEREVENT']._serialized_end = 747
    _globals['_USERSERVICE']._serialized_start = 750
    _globals['_USERSERVICE']._serialized_end = 1314