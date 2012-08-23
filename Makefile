SERVER=foreman start
JAMMIT=jammit
BUNDLE=bundle
GUARD=$(BUNDLE) exec guard
PUBLIC_DIR?=public
DEV_OPTS=
DEV?=0

all: help

help:
	@echo "Usage: make target [VARS...]"
	@echo "The targets are: server, assets, guard, deploy, prepare"
	@echo "Defaults: DEV=0; PUBLIC_DIR=public;"

server:
	$(SERVER)

jammit:
	DEV=$(DEV) $(JAMMIT) -f -o $(PUBLIC_DIR)/assets

guard:
	$(GUARD) start -i

deploy: all
	-git add $(PUBLIC_DIR) && git commit -qm "Recompiled assets"
	git push heroku master

prepare:
	$(BUNDLE) install
