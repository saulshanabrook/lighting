MARKER_FOLDER = .makefile_markers/$(NAME)


NAME = python/falcon
LANGUAGE = $(split $(NAME) / 1)
LIBRARY = $(split $(NAME) / 2)

CONTEXT_ROOT = implementations/backend/$(LANGUAGE)

DOCKER_IMAGE_CREATED = $(MARKER_FOLDER)/created_docker_image
DOCKER_RUNNING = $(MARKER_FOLDER)/running_docker
DOCKER_IMAGE_NAME = ligting-$(NAME)
DOCKER_WEB_NAME = lighting-$(NAME)-web
DOCKER_TEST_NAME = lighting-$(NAME)-test
DOCKER_PORT = 8080

$(MARKER_FOLDER):
	mkdir -p MARKER_FOLDER

# Test
## Build dockerfile
$(DOCKER_IMAGE_CREATED): | $(MARKER_FOLDER)
	docker build -t $(DOCKER_IMAGE_NAME) --file="$(LIBRARY).Dockerfile" $(CONTEXT_ROOT)
	touch $(DOCKER_IMAGE_CREATED)

## Run app
$(DOCKER_RUNNING): $(DOCKER_IMAGE_CREATED)
	docker rm --force $(DOCKER_WEB_NAME) | true
	docker run \
	--name $(DOCKER_WEB_NAME) \
	--publish=[$(DOCKER_PORT)] \
	-e PORT=$(DOCKER_PORT) \
	-d \
	$(DOCKER_IMAGE_NAME)
	sleep 5 # so that it has time to create

test:
	docker run \
	--name $(DOCKER_TEST_NAME) \
	--link $(DOCKER_WEB_NAME):web \
	--rm \
	lighting-test
