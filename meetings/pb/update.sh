#!/bin/bash
protoc -I=./ -I=../../../booking-calendar-services/protos/ ../../../booking-calendar-services/protos/users.proto --go_out=plugins=grpc:./
protoc -I=./ -I=../../../booking-calendar-services/protos/ ../../../booking-calendar-services/protos/meetings.proto --go_out=plugins=grpc:./