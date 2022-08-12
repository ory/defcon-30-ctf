#!/bin/sh
set -eux

# migrate the database
keto migrate up -y -c /home/ory/keto.yml
# start the server
keto serve -c /home/ory/keto.yml --sqa-opt-out &

# TODO remove testdata before deployment
keto status -b
echo '{"namespace": "districts", "object": "Testing", "relation": "submit", "subject_id": "8de7540d-b5d2-4dc2-bac0-34e9048a4d63"}' | keto relation-tuple create -
echo '{"namespace": "flags", "object": "new_flag", "relation": "create", "subject_id": "8de7540d-b5d2-4dc2-bac0-34e9048a4d63"}' | keto relation-tuple create -

wait
