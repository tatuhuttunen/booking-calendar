#!/bin/sh
cat  pb/**/*.swagger.json | \
jq '.info.title = "Booking Calendar"' | \
jq '.info.version = "1.0.0"' | \
jq '.server.url = "http://localhost:8080"' | \
jq --slurp 'reduce .[] as $item ({}; . * $item)' > cmd/api/swagger.json