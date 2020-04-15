#!/bin/bash

openapi-generator generate \
  -i ./swagger.json \
  -g python \
  --library asyncio \
  --package-name openshift-release-apiclient \
  --model-package releaseapi.models \
  --api-package releaseapi.api \
  -o apiclient-python
