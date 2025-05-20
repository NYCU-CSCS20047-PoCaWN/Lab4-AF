#!/bin/bash

# Color codes
YELLOW='\033[1;33m'
CYAN='\033[1;36m'
NC='\033[0m' # No Color


while true; do
    echo -e "${CYAN}[Get Warning Users]${NC}"
    curl -s -X GET http://af.free5gc.org:8000/oam/warning-users | jq
    printf "\n"
    sleep 5
done