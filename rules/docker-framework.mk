# docker image for framework

DOCKER_FRAMEWORK_STEM = docker-framework
DOCKER_FRAMEWORK = $(DOCKER_FRAMEWORK_STEM).gz
DOCKER_FRAMEWORK_DBG = $(DOCKER_FRAMEWORK_STEM)-$(DBG_IMAGE_MARK).gz

$(DOCKER_FRAMEWORK)_PATH = $(DOCKERS_PATH)/$(DOCKER_FRAMEWORK_STEM)

$(DOCKER_FRAMEWORK)_DEPENDS += $(FRAMEWORK)
$(DOCKER_FRAMEWORK)_DBG_DEPENDS = $($(DOCKER_CONFIG_ENGINE_BULLSEYE)_DBG_DEPENDS)
$(DOCKER_FRAMEWORK)_DBG_DEPENDS += $(FRAMEWORK_DBG) $(LIBSWSSCOMMON_DBG)
$(DOCKER_FRAMEWORK)_DBG_IMAGE_PACKAGES = $($(DOCKER_CONFIG_ENGINE_BULLSEYE)_DBG_IMAGE_PACKAGES)

$(DOCKER_FRAMEWORK)_LOAD_DOCKERS += $(DOCKER_CONFIG_ENGINE_BULLSEYE)
$(DOCKER_FRAMEWORK)_LOAD_DOCKERS += $($(DOCKER_CONFIG_ENGINE_BULLSEYE)_LOAD_DOCKERS)

$(DOCKER_FRAMEWORK)_VERSION = 1.0.0
$(DOCKER_FRAMEWORK)_PACKAGE_NAME = framework

SONIC_DOCKER_IMAGES += $(DOCKER_FRAMEWORK)
SONIC_INSTALL_DOCKER_IMAGES += $(DOCKER_FRAMEWORK)

SONIC_DOCKER_DBG_IMAGES += $(DOCKER_FRAMEWORK_DBG)
SONIC_INSTALL_DOCKER_DBG_IMAGES += $(DOCKER_FRAMEWORK_DBG)

$(DOCKER_FRAMEWORK)_CONTAINER_NAME = framework
$(DOCKER_FRAMEWORK)_RUN_OPT += --privileged -t
$(DOCKER_FRAMEWORK)_RUN_OPT += -v /etc/sonic:/etc/sonic:ro
$(DOCKER_FRAMEWORK)_GIT_REPOSITORIES += "sonic-swss"
$(DOCKER_FRAMEWORK)_GIT_REPOSITORIES += "sonic-swss-common"

$(DOCKER_FRAMEWORK)_FILES += $(SUPERVISOR_PROC_EXIT_LISTENER_SCRIPT)

SONIC_BULLSEYE_DOCKERS += $(DOCKER_FRAMEWORK)
SONIC_BULLSEYE_DBG_DOCKERS += $(DOCKER_FRAMEWORK_DBG)
