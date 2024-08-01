#!/usr/bin/env sh

docker run -it --rm -v "$(pwd):/src" -u "$(id -u):$(id -g)" --network host --workdir /src/webui node:lts /bin/bash
# THIS CODE IS PROPERTY OF ENRICO BASSETTI, Licensed under the MIT License.
# i did not wrote this and this is used only for development purposes. cause it uses npm:lts 
# and so it would be a better choice for compatibility with the project.