#!/usr/bin/env bash

echo "test whatsapp..."
curl --location --request POST 'http://localhost:8000/whatsapp' \
--header 'Content-type: application/json' \
--data @./testdata/post.json
echo ""
