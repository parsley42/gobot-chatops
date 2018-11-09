#!/bin/bash -e
# mkdist.sh - create a distributable .zip file

trap_handler()
{
    ERRLINE="$1"
    ERRVAL="$2"
    echo "line ${ERRLINE} exit status: ${ERRVAL}"
    rm -f commit.go
    exit $ERRVAL
}
trap 'trap_handler ${LINENO} $?' ERR

usage(){
	cat <<EOF
Usage: mkdist.sh (linux|darwin|windows)

Generate distributable .zip files for the given platform, or all platforms if
no argument given.
EOF
	exit 0
}

if [ "$1" = "-h" -o "$1" = "--help" ]
then
	usage
fi

git status | grep -qE "nothing to commit, working directory|tree clean" || { echo "Your working directory isn't clean, aborting build"; exit 1; }
COMMIT=$(git rev-parse HEAD)
cat >commit.go <<EOF
package main

func init(){
	versionInfo.Commit = "$COMMIT"
}
EOF

eval `go env`
PLATFORMS=${1:-linux darwin windows}
for BUILDOS in $PLATFORMS
do
	echo "Building gopherbot for $BUILDOS"
	OUTFILE=./gopherbot-$BUILDOS-$GOARCH.zip
	rm -f $OUTFILE
	if [ "$BUILDOS" = "windows" ]
	then
		GOOS=$BUILDOS go build -mod vendor -o gopherbot.exe
		echo "Creating $OUTFILE"
		zip -r $OUTFILE gopherbot.exe LICENSE README.md brain/ conf/ doc/ cfg/ lib/ licenses/ plugins/ resources/ jobs/ tasks/ scripts/ --exclude *.swp
	then
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -a -tags 'netgo osusergo static_build' -o gopherbot
		echo "Creating $OUTFILE"
		zip -r $OUTFILE gopherbot LICENSE README.md brain/ conf/ doc/ cfg/ lib/ licenses/ plugins/ resources/ jobs/ tasks/ scripts/ --exclude *.swp
	else
		GOOS=$BUILDOS go build -mod vendor
		echo "Creating $OUTFILE"
		zip -r $OUTFILE gopherbot LICENSE README.md brain/ conf/ doc/ cfg/ lib/ licenses/ plugins/ resources/ jobs/ tasks/ scripts/ --exclude *.swp
	fi
done
rm -f commit.go
