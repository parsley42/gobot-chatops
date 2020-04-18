#!/bin/bash -e

# startbuild.sh - utility task to tell the user a build is starting

source $GOPHER_INSTALLDIR/lib/gopherbot_v1.sh

TELL="Starting build of $GOPHER_REPOSITORY, branch '$GOPHERCI_BRANCH'"
if [ "$GOPHERCI_CUSTOM_PIPELINE" ]
then
  TELL="Starting custom job for $GOPHER_REPOSITORY, branch '$GOPHERCI_BRANCH', pipeline '$GOPHERCI_CUSTOM_PIPELINE'"
fi

if [ "$GOPHER_LOG_REF" ] || [ "$GOPHER_LOG_LINK" ]
then
  if [ "$GOPHER_LOG_REF" ] && [ "$GOPHER_LOG_LINK" ]
  then
    TELL="$TELL (log $GOPHER_LOG_REF; link $GOPHER_LOG_LINK)"
  elif [ "$GOPHER_LOG_REF" ]
  then
    TELL="$TELL (log $GOPHER_LOG_REF)"
  else
    TELL="$TELL (link $GOPHER_LOG_LINK)"
  fi
fi

Say "$TELL"