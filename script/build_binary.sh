docker run \
		--rm \
		-e CGO_ENABLED=1 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/github.com/answerdev/answer \
		-v `pwd`/sysroot:/sysroot \
		-w /go/src/github.com/answerdev/answer \
		goreleaser/goreleaser-cross \
		--rm-dist --skip-validate --skip-publish