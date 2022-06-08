
run-docker-img-debug:
	docker run -p 8080:8080 -it --entrypoint /bin/sh my-http-server 

run-docker-img:
	docker run -p 8080:8080 my-http-server

build-docker-img:
	# docker image build . -t my-http-server
	docker buildx build --platform linux/amd64 -t my-http-server .

publish-docker-img:
	docker tag my-http-server us-central1-docker.pkg.dev/synthetic-grail-352701/my-repo/my-http-server
	docker push us-central1-docker.pkg.dev/synthetic-grail-352701/my-repo/my-http-server
