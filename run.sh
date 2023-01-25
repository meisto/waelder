#!/bin/zsh
# ======================================================================
# Author: Tobias Meisel (meisto)
# Creation Date: Thu 19 Jan 2023 02:48:12 PM CET
# Description: -
# ======================================================================
#

dir_path="/tmp/dntui"
db_name="db.sqlite"

export waelder_db_path="/tmp/dntui/db.sqlite"
test ! -d "$dir_path" && mkdir $dir_path

go run cmd/waelder/waelder.go
