ROOT = .makefile_markers
MARKER_FOLDER = $(ROOT)/$(NAME)


NAME = python/falcon
LANGUAGE = $(split $(NAME) / 1)
LIBRARY = $(split $(NAME) / 2)
TYPE = http

WEB_ROOT = implementations/backend/$(LANGUAGE)
TEST_ROOT = test

WEB_DOCKERFILE = $(LIBRARY).Dockerfile

DOCKER_WEB_CREATED = $(MARKER_FOLDER)/created_image
DOCKER_TEST_CREATED = $(ROOT)/created_test_image

DOCKER_RUNNING = $(MARKER_FOLDER)/running_docker
DOCKER_IMAGE_NAME = ligting-$(NAME)
DOCKER_WEB_NAME = lighting-$(NAME)-web
DOCKER_TEST_NAME = lighting-$(NAME)-test
DOCKER_PORT = 8080

$(ROOT_FOLDER):
	mkdir -p $(ROOT)

$(MARKER_FOLDER): | $(ROOT_FOLDER)
	mkdir -p $(MARKER_FOLDER)

# Test
## Build dockerfile
$(DOCKER_WEB_CREATED): | $(MARKER_FOLDER)
	docker build -t $(DOCKER_IMAGE_NAME) --file="$(WEB_DOCKERFILE)" $(WEB_ROOT)
	touch $(DOCKER_WEB_CREATED)

## Run app
$(DOCKER_RUNNING): $(DOCKER_WEB_CREATED)
	docker rm --force $(DOCKER_WEB_NAME) | true
	docker run \
	--name $(DOCKER_WEB_NAME) \
	--publish=[$(DOCKER_PORT)] \
	-e PORT=$(DOCKER_PORT) \
	-d \
	$(DOCKER_IMAGE_NAME)
	sleep 5 # so that it has time to create

$(DOCKER_TEST_CREATED): | $(ROOT_FOLDER)
	docker build -t $(DOCKER_TEST_NAME) $(TEST_ROOT)

test:
	docker run \
	--name $(DOCKER_TEST_NAME) \
	--link $(DOCKER_WEB_NAME):web \
	-e PORT=$(PORT) \
	-e TYPE=$(TYPE) \
	--rm \
	lighting-test
