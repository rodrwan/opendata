VERSION=v0.0.1
SVC=frontend
BIN=$(PWD)/bin/$(SVC)

REGISTRY_URL=rodrwan

docker d:
	@echo "[docker] Building image..."
	@docker build -t $(REGISTRY_URL)/$(SVC):$(VERSION) .

push: docker
	@echo "[docker] pushing $(REGISTRY_URL)/$(SVC):$(VERSION)"
	docker push $(REGISTRY_URL)/$(SVC):$(VERSION)