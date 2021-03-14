# Overview

Using JSON-RPC to interact with Solidity is a fiddly prospect involving lots of boilerplate codes. Fortunately, Geth provides a nifty tool call `abigen` to generate Go contract binding of solidity contract.

Whilst `abigen` solves part of the problem, you will need to ensure that you have to install appropriate versions of the solidity compiler (`solc`). You will also have to build `abigen`.

In this project, we have created a docker based tool to help you manage multiple versions of the `solc` and `abigen`.

## Prerequisite

Please install docker and make.

## How to use this tool

1. Fork this code.

2. Add your contract under the folder `solidity`.

3. Update the `Makefile`, use the example for generating Go binding for the `trontoken` contract.

4. Run the command `make`.

## Copyright Notice

Copyright 2020 Open Consentia Contributors.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.