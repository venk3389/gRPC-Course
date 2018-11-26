#! /bin/sh



/usr/local/bin/protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
/usr/local/bin/protoc calculator/sumpb/sum.proto --go_out=plugins=grpc:.

