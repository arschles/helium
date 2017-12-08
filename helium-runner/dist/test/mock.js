"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const job_1 = require("../src/job");
// This package contains mocks of objects found elsewhere in Brigade.
function mockProject() {
    return {
        id: "brigade-c0ff33544b459e6ac0ffee",
        name: "deis/empty-testbed",
        repo: {
            name: "deis/empty-testbed",
            cloneURL: "https://github.com/deis/empty-testbed.git",
            token: "supersecret"
        },
        allowPrivilegedJobs: true,
        allowHostMounts: false,
    };
}
exports.mockProject = mockProject;
function mockEvent() {
    return {
        buildID: "1234567890abcdef",
        workerID: "test-1234567890abcdef-12345678",
        type: "push",
        provider: "github",
        commit: "c0ffee",
        payload: "{}"
    };
}
exports.mockEvent = mockEvent;
class MockResult {
    constructor(msg) {
        this.msg = "uninitialized";
        this.msg = msg;
    }
    toString() {
        return this.msg;
    }
}
exports.MockResult = MockResult;
// MockJob implements the run() method on Job with a resolved Promise<MockResult>.
//
// If 'MockJob.fail = true', the job will return a failure instead of a success.
//
// The MockJob.run method will sleep for one nanosecond (that is, give up at least
// one scheduler run). To set a longer delay, set MockJob.delay.
class MockJob extends job_1.Job {
    constructor() {
        super(...arguments);
        this.fail = false;
        this.delay = 1; // Just enough to cause the event loop to sleep it.
    }
    run() {
        let fail = this.fail;
        let delay = this.delay;
        this._podName = "generated-fake-job-name";
        return new Promise((resolve, reject) => {
            if (fail) {
                setTimeout(() => { reject("Failed"); }, delay);
                return;
            }
            setTimeout(resolve(new MockResult(this.name)), delay);
        });
    }
}
exports.MockJob = MockJob;
class MockBuildStorage {
    create(e, project, size) {
        return Promise.resolve(e.workerID);
    }
    destroy() {
        return Promise.resolve(true);
    }
}
exports.MockBuildStorage = MockBuildStorage;
