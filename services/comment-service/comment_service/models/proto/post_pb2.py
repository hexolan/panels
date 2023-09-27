"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_sym_db = _symbol_database.Default()
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\npost.proto\x12\x0epanels.post.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto"\xb7\x01\n\x04Post\x12\n\n\x02id\x18\x01 \x01(\t\x12\x10\n\x08panel_id\x18\x02 \x01(\t\x12\x11\n\tauthor_id\x18\x03 \x01(\t\x12\r\n\x05title\x18\x04 \x01(\t\x12\x0f\n\x07content\x18\x05 \x01(\t\x12.\n\ncreated_at\x18\x06 \x01(\x0b2\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x07 \x01(\x0b2\x1a.google.protobuf.Timestamp"M\n\x0bPostMutable\x12\x12\n\x05title\x18\x01 \x01(\tH\x00\x88\x01\x01\x12\x14\n\x07content\x18\x02 \x01(\tH\x01\x88\x01\x01B\x08\n\x06_titleB\n\n\x08_content"a\n\x11CreatePostRequest\x12\x10\n\x08panel_id\x18\x01 \x01(\t\x12\x0f\n\x07user_id\x18\x02 \x01(\t\x12)\n\x04data\x18\x03 \x01(\x0b2\x1b.panels.post.v1.PostMutable"\x1c\n\x0eGetPostRequest\x12\n\n\x02id\x18\x01 \x01(\t"3\n\x13GetPanelPostRequest\x12\x10\n\x08panel_id\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t"J\n\x11UpdatePostRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12)\n\x04data\x18\x02 \x01(\x0b2\x1b.panels.post.v1.PostMutable"\x1f\n\x11DeletePostRequest\x12\n\n\x02id\x18\x01 \x01(\t"\x15\n\x13GetFeedPostsRequest"0\n\tFeedPosts\x12#\n\x05posts\x18\x01 \x03(\x0b2\x14.panels.post.v1.Post"&\n\x13GetUserPostsRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t"0\n\tUserPosts\x12#\n\x05posts\x18\x01 \x03(\x0b2\x14.panels.post.v1.Post"(\n\x14GetPanelPostsRequest\x12\x10\n\x08panel_id\x18\x01 \x01(\t"1\n\nPanelPosts\x12#\n\x05posts\x18\x01 \x03(\x0b2\x14.panels.post.v1.Post"=\n\tPostEvent\x12\x0c\n\x04type\x18\x01 \x01(\t\x12"\n\x04data\x18\x02 \x01(\x0b2\x14.panels.post.v1.Post2\xf3\x04\n\x0bPostService\x12G\n\nCreatePost\x12!.panels.post.v1.CreatePostRequest\x1a\x14.panels.post.v1.Post"\x00\x12A\n\x07GetPost\x12\x1e.panels.post.v1.GetPostRequest\x1a\x14.panels.post.v1.Post"\x00\x12K\n\x0cGetPanelPost\x12#.panels.post.v1.GetPanelPostRequest\x1a\x14.panels.post.v1.Post"\x00\x12G\n\nUpdatePost\x12!.panels.post.v1.UpdatePostRequest\x1a\x14.panels.post.v1.Post"\x00\x12I\n\nDeletePost\x12!.panels.post.v1.DeletePostRequest\x1a\x16.google.protobuf.Empty"\x00\x12P\n\x0cGetFeedPosts\x12#.panels.post.v1.GetFeedPostsRequest\x1a\x19.panels.post.v1.FeedPosts"\x00\x12P\n\x0cGetUserPosts\x12#.panels.post.v1.GetUserPostsRequest\x1a\x19.panels.post.v1.UserPosts"\x00\x12S\n\rGetPanelPosts\x12$.panels.post.v1.GetPanelPostsRequest\x1a\x1a.panels.post.v1.PanelPosts"\x00b\x06proto3')
_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'post_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
    DESCRIPTOR._options = None
    _globals['_POST']._serialized_start = 93
    _globals['_POST']._serialized_end = 276
    _globals['_POSTMUTABLE']._serialized_start = 278
    _globals['_POSTMUTABLE']._serialized_end = 355
    _globals['_CREATEPOSTREQUEST']._serialized_start = 357
    _globals['_CREATEPOSTREQUEST']._serialized_end = 454
    _globals['_GETPOSTREQUEST']._serialized_start = 456
    _globals['_GETPOSTREQUEST']._serialized_end = 484
    _globals['_GETPANELPOSTREQUEST']._serialized_start = 486
    _globals['_GETPANELPOSTREQUEST']._serialized_end = 537
    _globals['_UPDATEPOSTREQUEST']._serialized_start = 539
    _globals['_UPDATEPOSTREQUEST']._serialized_end = 613
    _globals['_DELETEPOSTREQUEST']._serialized_start = 615
    _globals['_DELETEPOSTREQUEST']._serialized_end = 646
    _globals['_GETFEEDPOSTSREQUEST']._serialized_start = 648
    _globals['_GETFEEDPOSTSREQUEST']._serialized_end = 669
    _globals['_FEEDPOSTS']._serialized_start = 671
    _globals['_FEEDPOSTS']._serialized_end = 719
    _globals['_GETUSERPOSTSREQUEST']._serialized_start = 721
    _globals['_GETUSERPOSTSREQUEST']._serialized_end = 759
    _globals['_USERPOSTS']._serialized_start = 761
    _globals['_USERPOSTS']._serialized_end = 809
    _globals['_GETPANELPOSTSREQUEST']._serialized_start = 811
    _globals['_GETPANELPOSTSREQUEST']._serialized_end = 851
    _globals['_PANELPOSTS']._serialized_start = 853
    _globals['_PANELPOSTS']._serialized_end = 902
    _globals['_POSTEVENT']._serialized_start = 904
    _globals['_POSTEVENT']._serialized_end = 965
    _globals['_POSTSERVICE']._serialized_start = 968
    _globals['_POSTSERVICE']._serialized_end = 1595