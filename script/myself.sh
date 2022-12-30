#!/bin/bash
#
# Script file to use curl command to get current user's info.
#

URL="${URL:-"http://127.0.0.1"}"
PORT="${PORT:-"8000"}"
METHOD="GET"
ENDPOINT='api/user/myself'
FILE_COOKIE=".cookie"
OPTS="-s"
NUM_ARGS=0
DEBUG=""

# Function
SCRIPT_NAME=${0##*/}
Usage () {
	echo
	echo "Description:"
	echo "Script file to use curl command to get current user's info."
	echo
	echo "Usage: $SCRIPT_NAME [username] [password]"
	echo "Options:"
	echo " -d                           Print curl statement output for debug purpose"
	echo " -k                           Allow https insecure connection"
	echo " -u  [url]                    iTMS API server URL"
	echo " -p  [port]                   iTMS API server port number"
	echo " -v                           Make the curl operation more talkative"
	echo " -h                           This help message"
	echo
}

# Parse input argument(s)
while [ "${1:0:1}" == "-" ]; do
	OPT=${1:1:1}
	case "$OPT" in
	"d")
		DEBUG="echo"
		;;
	"k")
		OPTS="$OPTS -k"
		;;
	"u")
		URL=$2
		shift
		;;
	"p")
		PORT=$2
		shift
		;;
	"v")
		OPTS="$OPTS -v"
		;;
	"h")
		Usage
		exit
		;;
	esac
	shift
done

if [ "$#" -ne "$NUM_ARGS" ]; then
    echo "Invalid parameter!"
	Usage
	exit 1
fi

# trim URL trailing "/"
if [ "$PORT" = "" ]; then
	URL="$(echo -e "${URL}" | sed -e 's/\/*$//')"
else
	URL="$(echo -e "${URL}" | sed -e 's/\/*$//')"
	URL="$(echo -e "${URL}:${PORT}" | sed -e 's/\/*$//')"
fi

# perform curl
$DEBUG curl $OPTS \
	-X $METHOD \
	--cookie-jar $FILE_COOKIE \
	--cookie $FILE_COOKIE \
	-H "Content-Type: application/json" \
	-H "Accept: application/json" \
	${URL}/${ENDPOINT}
