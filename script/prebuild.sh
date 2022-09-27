#!/usr/bin/env bash

git clone https://git.backyard.segmentfault.com/opensource/pacman.git /tmp/sf-pacman

cat <<'EOF' > go.work
go 1.18

use (
	.
	/tmp/sf-pacman
	/tmp/sf-pacman/contrib/cache/redis
	/tmp/sf-pacman/contrib/cache/memory
	/tmp/sf-pacman/contrib/conf/viper
	/tmp/sf-pacman/contrib/log/zap
	/tmp/sf-pacman/contrib/i18n
	/tmp/sf-pacman/contrib/server/http
)
EOF

go work sync
