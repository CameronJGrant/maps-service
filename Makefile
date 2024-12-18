.PHONY: maps-service web validation
maps-service:
	DOCKER_BUILDKIT=1 docker build . -f Containerfile -t maps-service \
		--secret id=netrc,src=$HOME/.netrc
web:
	docker build web -f web/Containerfile -t maps-service-web
validation:
	docker build validation -f validation/Containerfile -t maps-service-validation
