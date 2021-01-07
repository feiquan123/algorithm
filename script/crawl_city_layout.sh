#!/bin/bash

usage() {
  cat 1>&2 <<EOF
crawl.sh crawl layout map data from https://beijing.anjuke.com/v3/ajax/map/sale/boundary/

USAGE:
    crawl.sh [FLAGS] \$1
    crawl.sh [FLAGS] \$1 \$2

FLAGS:
    -h, --help    help
    -s, --sed     ses result
    -g            generate concat file
    -w,--web      see in web

OPTIONS:
    \$1   file path , content format [id name].
          or id where input \$1 and \$2
    \$2   name
EOF
}

re="./data/map/"
city="dg"

fetch() {
  mkdir "$re" 2 > /dev/null
  cat "$1" | awk -v re=${re} -v city="${city}" '{print "curl \"https://city.anjuke.com/v3/ajax/map/sale/boundary/?region_id="$1"\" | jq -r '\''.val.\""$1"\"[0]|.[]|.[1]+\",\"+.[0]'\'' | tr \"\\n\" \"\\\&\" | awk '\''{sub(/.{1}$/,\"\")}1'\'' | tr \"\\n\" \" \" > "re$2".txt && sleep 1"}'
}

see() {
  for f in `ls ${re}`
  do
    t=${re}${f}
    echo `echo $f | tr '.txt' " "`": "$(cat $t)
  done
}

web(){
  for f in `ls ${re}`
  do
    cat $re$f | tr '&' '\n' && echo -e "\n######"
  done
}

main() {
  need_cmd cat
  need_cmd awk
  need_cmd curl
  need_cmd jq
  need_cmd tr
  need_cmd sleep
  need_cmd mkdir
  need_cmd echo

  for arg in "$@";do
      case "$arg" in
          -h|--help)
              usage
              exit 0
              ;;
          -s|--see)
              see
              exit 0
              ;;
          -w|--web)
              web
              exit 0
              ;;
          -g)
              see > ./data/map.txt
              exit  0
              ;;
          *)
              ;;
      esac
  done

  file=""
  if [ -z "$2" ]; then
    if [ -z "$1" -o ! -f "$1" ]; then
      echo "plz input filename, [id->name]"
      exit 1
    fi
    file="$1"
  else
    file=/tmp/t.txt
    echo $1 $2 >"$file"
  fi

  fetch $file | sh
}

err() {
    echo "$1" >&2
    exit 1
}

need_cmd() {
    if ! check_cmd "$1"; then
        err "need '$1' (command not found)"
    fi
}

check_cmd() {
    command -v "$1" > /dev/null 2>&1
}

ensure() {
    if ! "$@"; then err "command failed: $*"; fi
}

main "$@" || exit 1
