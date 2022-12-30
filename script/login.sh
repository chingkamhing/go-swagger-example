#!/bin/bash
#
# Script file to use curl command to login to wk-api. Upon successful login, it will response access token which can then be used as authen token for other authen restful api access.
# e.g.
# $ export TOKEN=$(./script/login.sh kamching password1234 | jq .data.access_token | tr -d '"')
# $ ./script/get-user.sh 2 | jq .
#

URL="${URL:-"http://127.0.0.1"}"
PORT="${PORT:-"8000"}"
METHOD="POST"
ENDPOINT='api/auth/login'
OPTS="-s"
NUM_ARGS=2
DEBUG=""

# Function
SCRIPT_NAME=${0##*/}
Usage () {
	echo
	echo "Description:"
	echo "Script file to use curl command to login to wk-api. Upon successful login, it will response access token which can then be used as authen token for other authen restful api access."
	echo "e.g."
	echo $'$ export TOKEN=$(./script/login.sh kamching password1234 | jq .data.access_token | tr -d \'"\')'
	echo '$ ./script/get-user.sh 2 | jq .'
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

USERNAME=$1
PASSWORD=$2

# perform curl
$DEBUG curl $OPTS -d \
	"{ \
		\"username\": \"$USERNAME\", \
		\"password\": \"$PASSWORD\" \
	}" \
	-X $METHOD \
	-H "Content-Type: application/json" \
	-H "Accept: application/json" \
	${URL}/${ENDPOINT}
