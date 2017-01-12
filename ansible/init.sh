#!/bin/bash

virtualenv wheel
source wheel/bin/activate
pip install -r requirements.txt
#ansible-galaxy install -r Ansiblefile -p roles
