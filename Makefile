build:
	buildctl build \
		--frontend dockerfile.v0 \
		--local dockerfile=. \
		--local context=. \
		--output type=image,name=registry.moonrhythm.io/h2cp,push=true

build-gcr:
	buildctl build \
		--frontend dockerfile.v0 \
		--local dockerfile=. \
		--local context=. \
		--output type=image,name=gcr.io/moonrhythm-containers/h2cp,push=true
