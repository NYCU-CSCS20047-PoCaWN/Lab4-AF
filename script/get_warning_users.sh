#!/bin/bash

# Yellow color
YELLOW='\033[1;33m'
CYAN='\033[1;36m'
NC='\033[0m' # No Color

echo -e "${CYAN}[Test Get UE Usage]${NC}"
curl -X GET http://af.free5gc.org:8000/oam/ue-usage

printf "\n\n==========================\n\n"

echo -e "${CYAN}[Test Get Warning Users]${NC}"
curl -X GET http://af.free5gc.org:8000/oam/warning-users

printf "\n"
