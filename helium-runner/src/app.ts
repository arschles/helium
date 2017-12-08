/**
 * Module app is the main application runner.
 */
import * as ulid from "ulid"

import * as events from "./events"
import * as process from "process"
import * as brigadier from './balloon'
import { loadEmptyProject } from './project';


/**
 * App is the main application.
 *
 * App assumes that it has full control of the process. It acts as a top-level
 * error handler and will exit the process with errors when uncaught resolutions
 * and errors occur.
 */
export class App {
  constructor() { }
  /**
   * exitOnError controls whether the app will exit when an uncaught exception or unhandled rejection occurs.
   *
   * exitOnError can be set to false in order to run tests on the error handling.
   * In general, though, it should be left on. In some cases, by the time the
   * process trap is invoked, the runtime is not in a good state to continue.
   */
  public exitOnError: boolean = true
  protected errorsHandled: boolean = false
  protected lastEvent: events.Event
  // On project loading error, this value may be passed. In all other cases,
  // it is overwritten by an actual project.
  protected proj: events.Project = new events.Project()

  // true if the "after" event has fired.
  protected afterHasFired: boolean = false
  protected storageIsDestroyed: boolean = false

  /**
   * run runs a particular event for this app.
   */
  public run(e: events.Event): Promise<boolean> {
    this.lastEvent = e

    // We need at least one error trap to avoid losing the error to a new
    // throw from EventEmitter.
    brigadier.events.once("error", () => {
      console.log("error handler is cleaning up")
      this.exitOnError && process.exit(1)
    })

    // We need to ensure that after is called exactly once. So we need an
    // empty after handler.
    brigadier.events.once("after", () => {
      this.afterHasFired = true

      // Delay long enough to cause beforeExit to be emitted again.
      setImmediate(() => {
        console.log("after: default event fired")
      }, 20)
    })

    // Run if an uncaught rejection happens.
    process.on("unhandledRejection", (reason: any, p: Promise<any>) => {
      var msg = reason
      // Kubernetes objects put error messages here:
      if (reason.body && reason.body.message) {
        msg = reason.body.message
      }
      console.log(`FATAL: ${msg} (rejection)`)
      this.fireError(reason, "unhandledRejection")
    })

    // Run at the end.
    process.on("beforeExit", (code) => {
      if (this.afterHasFired) {
        return
      }

      let after: events.Event = {
        id: e.id,
        workerID: e.workerID,
        type: "after",
        provider: "helium",
        metadata: e.metadata,
        cause: {
          event: e,
          trigger: code == 0 ? "success" : "failure"
        } as events.Cause
      }

      // Only fire an event if the top-level had a match.
      if (brigadier.events.has(e.type)) {
        brigadier.fire(after, this.proj)
      } else {
        this.afterHasFired = true
        setImmediate(() => {
          console.log("no-after: fired")
        }, 20)
      }
    })

    // Now that we have all the handlers registered, load the project and
    // execute the event.
    return new Promise((resolve, reject) => {
      brigadier.fire(e, this.proj)
      resolve(true);
    });
  }

  /**
   * fireError fires an "error" event when the top-level script catches an error.
   *
   * It is fired no more than once, and is only fired when the error bubbles all
   * the way to the top.
   */
  public fireError(reason?: any, errorType?: string): void {
    if (this.errorsHandled) {
      return
    }
    this.errorsHandled = true

    let errorEvent: events.Event = {
      id: this.lastEvent.id,
      workerID: this.lastEvent.workerID,
      type: "error",
      provider: "helium",
      metadata: this.lastEvent.metadata,
      cause: {
        event: this.lastEvent,
        reason: reason,
        trigger: errorType
      } as events.Cause
    }

    brigadier.fire(errorEvent, this.proj)
  }
}

