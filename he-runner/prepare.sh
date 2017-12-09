# This script copies the helium javascript build files from their locations on disk
# to the right place, so that they can serve as the entry point into the application.

# TODO: handle multiple scripts

cp ${HELIUM_SCRIPTS_DIR}/main.js ${HELIUM_TARGET_DIR}/main.js
