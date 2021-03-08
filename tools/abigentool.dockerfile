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

ARG SOLC_VER
ARG BUILDER_IMAGE

# Solidity compiler
FROM ethereum/solc:${SOLC_VER} as solc

# Abigen executable image
FROM $BUILDER_IMAGE

ENV GOROOT=/usr/local/go
ENV GOPATH=/opt/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

COPY --from=solc /usr/bin/solc /usr/local/bin/solc

RUN cd /opt/go/src/github.com/ethereum/go-ethereum/; make && make devtools; cd / && \
    mv $GOPATH/bin/abigen /usr/local/bin/abigen

ENTRYPOINT [ "/usr/local/bin/abigen" ]