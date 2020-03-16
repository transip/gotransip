#!/bin/sh

# define a list of packages to test
subpackages="authenticator availabilityzone colocation domain haip \
invoice ipaddress jwt mailservice product repository \
rest test traffic vps"

# prefix the folders with ./
subpackage_folders=$(echo $subpackages | sed 's/^/\.\//' | sed 's/ / \.\//g')

# test the root package and defined subpackages
go vet ./ $subpackage_folders
