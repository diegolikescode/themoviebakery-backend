build:
	yarn && yarn build

full: build
	node dist/server.js

full-forever: build
	forever start -c "node dist/server.js" .
