user-rpc-dev:
	@make -f deploy/makefile/user-rpc.mk release-test

release-test: user-rpc-dev
