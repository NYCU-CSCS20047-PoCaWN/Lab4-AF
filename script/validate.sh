#!/bin/bash

# Color codes
YELLOW='\033[1;33m'
CYAN='\033[1;36m'
NC='\033[0m' # No Color

echo -e "${CYAN}[UE Usage]${NC}"
curl -s -X GET http://af.free5gc.org:8000/oam/ue-usage | jq

printf "\n======================================\n"

echo -e "${CYAN}[Get Warning Users]${NC}"
curl -s -X GET http://af.free5gc.org:8000/oam/warning-users | jq

printf "\n"
