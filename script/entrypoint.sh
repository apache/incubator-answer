#!/bin/bash
if [ ! -f "/data/conf/config.yaml" ]; then
  /usr/bin/answer init
fi
/usr/bin/answer run -c  /data/conf/config.yaml
