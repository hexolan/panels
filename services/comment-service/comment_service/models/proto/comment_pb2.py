"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rcomment.proto\x12\x11panels.comment.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto"\xaa\x01\n\x07Comment\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0f\n\x07post_id\x18\x02 \x01(\t\x12\x11\n\tauthor_id\x18\x03 \x01(\t\x12\x0f\n\x07message\x18\x04 \x01(\t\x12.\n\ncreated_at\x18\x05 \x01(\x0b2\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x06 \x01(\x0b2\x1a.google.protobuf.Timestamp"!\n\x0eCommentMutable\x12\x0f\n\x07message\x18\x01 \x01(\t"k\n\x14CreateCommentRequest\x12\x0f\n\x07post_id\x18\x01 \x01(\t\x12\x11\n\tauthor_id\x18\x02 \x01(\t\x12/\n\x04data\x18\x03 \x01(\x0b2!.panels.comment.v1.CommentMutable"S\n\x14UpdateCommentRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12/\n\x04data\x18\x02 \x01(\x0b2!.panels.comment.v1.CommentMutable""\n\x14DeleteCommentRequest\x12\n\n\x02id\x18\x01 \x01(\t"\x1f\n\x11GetCommentRequest\x12\n\n\x02id\x18\x01 \x01(\t")\n\x16GetPostCommentsRequest\x12\x0f\n\x07post_id\x18\x01 \x01(\t"<\n\x0cPostComments\x12,\n\x08comments\x18\x01 \x03(\x0b2\x1a.panels.comment.v1.Comment"F\n\x0cCommentEvent\x12\x0c\n\x04type\x18\x01 \x01(\t\x12(\n\x04data\x18\x02 \x01(\x0b2\x1a.panels.comment.v1.Comment2\xc7\x03\n\x0eCommentService\x12V\n\rCreateComment\x12\'.panels.comment.v1.CreateCommentRequest\x1a\x1a.panels.comment.v1.Comment"\x00\x12V\n\rUpdateComment\x12\'.panels.comment.v1.UpdateCommentRequest\x1a\x1a.panels.comment.v1.Comment"\x00\x12R\n\rDeleteComment\x12\'.panels.comment.v1.DeleteCommentRequest\x1a\x16.google.protobuf.Empty"\x00\x12P\n\nGetComment\x12$.panels.comment.v1.GetCommentRequest\x1a\x1a.panels.comment.v1.Comment"\x00\x12_\n\x0fGetPostComments\x12).panels.comment.v1.GetPostCommentsRequest\x1a\x1f.panels.comment.v1.PostComments"\x00b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'comment_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    DESCRIPTOR._options = None
    _globals['_COMMENT']._serialized_start = 99
    _globals['_COMMENT']._serialized_end = 269
    _globals['_COMMENTMUTABLE']._serialized_start = 271
    _globals['_COMMENTMUTABLE']._serialized_end = 304
    _globals['_CREATECOMMENTREQUEST']._serialized_start = 306
    _globals['_CREATECOMMENTREQUEST']._serialized_end = 413
    _globals['_UPDATECOMMENTREQUEST']._serialized_start = 415
    _globals['_UPDATECOMMENTREQUEST']._serialized_end = 498
    _globals['_DELETECOMMENTREQUEST']._serialized_start = 500
    _globals['_DELETECOMMENTREQUEST']._serialized_end = 534
    _globals['_GETCOMMENTREQUEST']._serialized_start = 536
    _globals['_GETCOMMENTREQUEST']._serialized_end = 567
    _globals['_GETPOSTCOMMENTSREQUEST']._serialized_start = 569
    _globals['_GETPOSTCOMMENTSREQUEST']._serialized_end = 610
    _globals['_POSTCOMMENTS']._serialized_start = 612
    _globals['_POSTCOMMENTS']._serialized_end = 672
    _globals['_COMMENTEVENT']._serialized_start = 674
    _globals['_COMMENTEVENT']._serialized_end = 744
    _globals['_COMMENTSERVICE']._serialized_start = 747
    _globals['_COMMENTSERVICE']._serialized_end = 1202