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

file=$1
if [ "$file" == "" ]
then
    file=compose.yaml
fi

docker-compose version
docker-compose -f "$file" up --build -d

while true
do
    docker-compose -f "$file" ps | grep testing
    if [ $? -eq 1 ]
    then
        code=-1
        docker-compose -f "$file" logs | grep e2e-testing
        docker-compose -f "$file" logs | grep e2e-testing | grep Usage
        if [ $? -eq 1 ]
        then
            code=0
            echo "successed"
        fi

        docker-compose -f "$file" down
        set -e
        exit $code
    fi
    sleep 1
done
