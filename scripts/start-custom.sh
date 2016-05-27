USAGE="Usage: ./start-custom [-d|--daemon] name"

DAEMON=

while [[ $# > 1 ]]
do
key="$1"

echo ${key}

case $key in
    -d|--daemon)
      DAEMON=true
    ;;
    -n|--name)
      NAME=$2
      shift
    ;;
    *)
      # unknown option
    ;;
esac
shift # past argument or value
done

NAME=$1
: "${NAME:?Need to set WEBSITE_ROOT non-empty

${USAGE}
}"

NAME_LOWERCASE=$(echo "$NAME" | tr '[:upper:]' '[:lower:]')
NAME_UPPERCASE=$(echo "$NAME" | tr '[:lower:]' '[:upper:]')

FINAL_PATH_NAME=${NAME_UPPERCASE}_PATH
FINAL_PATH=${!FINAL_PATH_NAME}

echo "base path: ${FINAL_PATH}"

: "${FINAL_PATH:?Need to set ${FINAL_PATH_NAME} non-empty}"

FINAL_BINARY=${FINAL_PATH}/bin/${NAME_LOWERCASE}

if [ ! -f "$FINAL_BINARY" ];
then
  echo "File $FINAL_BINARY does not exist."
  exit 1
fi

if [ -z "$DAEMON" ]; then
  ${FINAL_BINARY}
else
  screen -d -m -S ${NAME} ${FINAL_BINARY}
  echo "${FINAL_BINARY} started in screen as ${NAME}"
fi
