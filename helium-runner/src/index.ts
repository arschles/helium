/**
 * The Brigade Worker is responsible for executing `brigade.js` files.
 *
 * When the Brigade Worker starts, it will look for the following environment
 * variables, which it will use to generate BrigadeEvent and Project configuration:
 *
 * - `BRIGADE_EVENT_TYPE`: The event type, such as `push`, `pull_request`
 * - `BRIGADE_EVENT_PROVIDER`: The name of the event provider, such as `github` or `dockerhub`
 * - `BRIGADE_PROJECT_ID`: The project ID. This is used to load the Project
 *   object from configuration.
 * - `BRIGADE_COMMIT`: The VCS commit ID (e.g. the Git commit)
 * - `BRIGADE_PAYLOAD`: The payload from the original event trigger.
 * - `BRIGADE_PROJECT_NAMESPACE`: For Kubernetes, this is the Kubernetes namespace in
 *   which new jobs should be created. The Brigade worker must have write access to
 *   this namespace.
 * - `BRIGADE_BUILD`: The ULID for the build. This is unique.
 * - `BRIGADE_BUILD_NAME`: This is actually the ID of the worker.
 *
 * Also, the Brigade script must be written to `brigade.js`.
 */

// Seems to be a bug in typedocs that requires this empty comment.
/** */

import * as fs from "fs"
import * as process from "process"

import * as ulid from "ulid"

import { Event } from "./events"
import * as brigadier from './balloon'
import { App } from "./app"

// helium is the result of compiling all of the helium scripts together. it only 
// registers events
//
// this module should match the "main.js" in the /helium-runner/prepare.sh file
import "./main"

const pkg = require('../../package.json')
console.log(`helium - runner version: ${pkg.version} `)

let defaultULID = ulid()
let e: Event = {
  id: process.env.HELIUM_BUILD_ID || defaultULID,
  workerID: process.env.HELIUM_BUILD_NAME || `unknown - ${defaultULID} `,
  type: process.env.HELIUM_EVENT_TYPE || "ping",
  provider: process.env.HELIUM_EVENT_PROVIDER || "unknown",
  metadata: process.env.HELIUM_EVENT_METADATA || ""
}

try {
  e.payload = fs.readFileSync("/etc/brigade/payload", "utf8")
} catch (e) {
  console.log("no payload loaded")
}

// Run the app.
(new App()).run(e)
