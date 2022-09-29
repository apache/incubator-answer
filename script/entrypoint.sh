#!/bin/bash
if [ ! -f "/data/config.yaml" ]; then
  /usr/bin/answer init
fi
/usr/bin/answer run -c  /data/config.yaml
