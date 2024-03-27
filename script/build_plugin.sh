#!/bin/bash
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

cmd="./answer build"

echo "Check, if file plugin_list exist"
sleep 1
if ! [ -f "plugin_list" ]; then
  echo "plugin_list is not exist"
  exit 0
fi

echo "Following plugins will be installed:"
echo "$(cat plugin_list)"
sleep 1

echo "Begin build plugin..."
sleep 1
for repo in $(cat plugin_list); do
  cmd+=" --with ${repo}"
done

$cmd

if ! [ -f "./new_answer" ]; then
  echo "File new_answer is not exist! Build failed"
  exit 1
fi

rm answer
mv new_answer answer

./answer plugin