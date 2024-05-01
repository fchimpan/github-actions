#!/usr/bin/bash

hoge=true
fuga=false

targets=$(jq -cn \
--argjson development $hoge \
--argjson staging $fuga \
'{
    environments: [
    (if $development then "development" else empty end),
    (if $staging then "staging" else empty end)
    ]
}')

echo "$targets"
