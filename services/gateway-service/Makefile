protobufs-compile:
	protoc --proto_path=../../protobufs/ --go_out=. --go_opt=Mpanel.proto=./internal/rpc/panelv1 --go-grpc_out=. --go-grpc_opt=Mpanel.proto=./internal/rpc/panelv1 panel.proto
	protoc --proto_path=../../protobufs/ --go_out=. --go_opt=Mpost.proto=./internal/rpc/postv1 --go-grpc_out=. --go-grpc_opt=Mpost.proto=./internal/rpc/postv1 post.proto
	protoc --proto_path=../../protobufs/ --go_out=. --go_opt=Muser.proto=./internal/rpc/userv1 --go-grpc_out=. --go-grpc_opt=Muser.proto=./internal/rpc/userv1 user.proto
	protoc --proto_path=../../protobufs/ --go_out=. --go_opt=Mauth.proto=./internal/rpc/authv1 --go-grpc_out=. --go-grpc_opt=Mauth.proto=./internal/rpc/authv1 auth.proto
	protoc --proto_path=../../protobufs/ --go_out=. --go_opt=Mcomment.proto=./internal/rpc/commentv1 --go-grpc_out=. --go-grpc_opt=Mcomment.proto=./internal/rpc/commentv1 comment.proto