#!/bin/sh
set -e

if [ -z "$PACKAGE" ]; then
  echo "PACKAGE is not set"
  exit 1
fi

if [ -z "$PACKAGE_TYPE" ]; then
  echo "PACKAGE_TYPE is not set"
  exit 1
fi

if [ -z "$BINTRAY_USER" ]; then
  echo "BINTRAY_USER is not set"
  exit 1
fi

if [ -z "$BINTRAY_API_KEY" ]; then
  echo "BINTRAY_API_KEY is not set"
  exit 1
fi

if [ -z "$BINTRAY_PACKAGE" ]; then
    BINTRAY_PACKAGE="Nightly"
fi

echo "Publishing package : $PACKAGE"

function getPlatformFromFileName () {
    if [ "$NOPLATFORM" == "1" ]; then
        echo ""
    else
        PLATFORM=$(echo $1 | sed "s/$PACKAGE_FILE_PREFIX-//" | rev | cut -d '-' -f 1 | rev | cut -d '.' -f 1);
        echo "$PLATFORM/"
    fi
}

PACKAGE_FILE_PREFIX=$(echo $PACKAGE | tr '[:upper:]' '[:lower:]')

for f in `ls`; do
    mv "$f" "`echo $f | tr '[:upper:]' '[:lower:]'`"
done

VERSION=$(ls $PACKAGE_FILE_PREFIX* | head -1 | sed "s/\.[^\.]*$//" | sed "s/$PACKAGE_FILE_PREFIX-//" | cut -d '-' -f 1);

if [ "$NOVERSION" == "1" ]; then
    VERSION="latest"
    echo "Not checking for package version"
    for f in $PACKAGE_FILE_PREFIX*;
      do mv "$f" "`echo $f | sed s/$PACKAGE_FILE_PREFIX/$PACKAGE_FILE_PREFIX-$VERSION/`";
    done
else
    if [ -z "$VERSION" ]; then
        echo "Could not determine $PACKAGE version"
        exit 1
    fi
fi

echo "Version to be uploaded: $VERSION"

CURR_DATE=$(date +"%Y-%m-%d")

for f in $PACKAGE_FILE_PREFIX*;
  do mv "$f" "`echo $f | sed s/$VERSION/$VERSION.$PACKAGE_TYPE-$CURR_DATE/`";
done

for i in `ls`; do
  PLATFORM=$( getPlatformFromFileName $i )
  URL="https://api.bintray.com/content/gauge/$PACKAGE/$BINTRAY_PACKAGE/$VERSION.$PACKAGE_TYPE-$CURR_DATE/$PLATFORM$i?publish=1&override=1"

  echo "Uploading to : $URL"

  RESPONSE_CODE=$(curl -T $i -u$BINTRAY_USER:$BINTRAY_API_KEY $URL -I -s -w "%{http_code}" -o /dev/null);
  if [[ "${RESPONSE_CODE:0:2}" != "20" ]]; then
    echo "Unable to upload, HTTP response code: $RESPONSE_CODE"
    exit 1
  fi
  echo "HTTP response code: $RESPONSE_CODE"
done;

echo "\nSleeping for 30 seconds. Have a coffee..."
sleep 30s;
echo "Done sleeping\n"

for i in `ls`; do
  PLATFORM=$( getPlatformFromFileName $i )
  URL="https://api.bintray.com/file_metadata/gauge/$PACKAGE/$PLATFORM$i"

  echo "Putting $i in $PACKAGE's download list"
  RESPONSE_CODE=$(curl -X PUT -d "{ \"list_in_downloads\": true }" -H "Content-Type: application/json" -u$BINTRAY_USER:$BINTRAY_API_KEY $URL -s -w "%{http_code}" -o /dev/null);
  if [[ "${RESPONSE_CODE:0:2}" != "20" ]]; then
    echo "Unable to put in download list, HTTP response code: $RESPONSE_CODE"
    exit 1
  fi
  echo "HTTP response code: $RESPONSE_CODE"
done
