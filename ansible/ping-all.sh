#!/bin/bash

ENVIRONMENT=${1:-ci}
ansible all -i environments/$ENVIRONMENT/inventory --private-key=environments/$ENVIRONMENT/.keys/dcos-bootstrap.pem -m ping
