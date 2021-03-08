# Copyright 2021 The Open Consentia Contributors
# 
# Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM ubuntu:18.04 as gotool

ENV GOROOT=/usr/local/go
ENV GOPATH=/opt/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

RUN apt-get update && \
    apt-get -y install wget gcc g++ make git nodejs npm protobuf-compiler && \
    cd /tmp; wget https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz && \
    tar -xvf go1.15.6.linux-amd64.tar.gz; mv go /usr/local/; cd / && \
    go get github.com/ethereum/go-ethereum
