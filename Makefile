image:
	docker build . \
		-t chevaliersoft/blackhole:latest \
		-t chevaliersoft/blackhole:0.2 \
		--platform=linux/amd64 \
		--platform=linux/arm64
