"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
require("mocha");
const chai_1 = require("chai");
const mock = require("./mock");
const job_1 = require("../src/job");
describe("job", function () {
    describe("jobNameIsValid", () => {
        it('allows DNS-like names', function () {
            let legal = [
                "abcdef",
                "ab",
                "a-b",
                "a-9",
                "a12345678b",
                "A-B"
            ];
            for (let n of legal) {
                chai_1.assert.isTrue(job_1.jobNameIsValid(n), n);
            }
        });
        it('disallows non-DNS-like names', function () {
            let illegal = [
                "ab-",
                "-ab",
            ];
            for (let n of illegal) {
                chai_1.assert.isFalse(job_1.jobNameIsValid(n), "tested " + n);
            }
        });
    });
    describe("JobCache", function () {
        describe("#constructor", function () {
            it("correctly sets default values", function () {
                let c = new job_1.JobCache();
                chai_1.assert.equal(c.path, job_1.brigadeCachePath, "Dir is /brigade/cache");
                chai_1.assert.isFalse(c.enabled, "disabled by default");
                chai_1.assert.equal(c.size, "5Mi", "size is 5mi");
            });
        });
    });
    describe("JobStorage", function () {
        describe("#constructor", function () {
            it("correctly sets default values", function () {
                let c = new job_1.JobStorage();
                chai_1.assert.equal(c.path, job_1.brigadeStoragePath, "Dir is " + job_1.brigadeStoragePath);
                chai_1.assert.isFalse(c.enabled, "disabled by default");
            });
        });
    });
    describe("Job", function () {
        let j;
        describe("#constructor", function () {
            it("creates a named job", function () {
                j = new mock.MockJob("myName");
                chai_1.assert.equal(j.name, "myName");
            });
            context("when image is supplied", function () {
                it("sets image property", function () {
                    j = new mock.MockJob("myName", "alpine:3.4");
                    chai_1.assert.equal(j.image, "alpine:3.4");
                });
            });
            context("when imageForcePull is supplied", function () {
                it("sets imageForcePull property", function () {
                    j = new mock.MockJob("myName", "alpine:3.4", [], true);
                    chai_1.assert.isTrue(j.imageForcePull);
                });
            });
            context("when tasks are supplied", function () {
                it("sets task list", function () {
                    j = new mock.MockJob("my", "img", ["a", "b", "c"]);
                    chai_1.assert.deepEqual(j.tasks, ["a", "b", "c"]);
                });
            });
        });
        describe("#podName", function () {
            beforeEach(function () {
                j = new mock.MockJob("my-job");
            });
            context("before run", function () {
                it("is empty", function () {
                    chai_1.assert.isUndefined(j.podName);
                });
            });
            context("after run", function () {
                it("is accessible", function (done) {
                    j.run().then((rez) => {
                        chai_1.assert.equal(j.podName, "generated-fake-job-name");
                        done();
                    });
                });
            });
        });
        describe("#cache", function () {
            it("is disabled by default", function () {
                chai_1.assert.isFalse(j.cache.enabled);
            });
        });
        describe("#storage", function () {
            it("is disabled by default", function () {
                chai_1.assert.isFalse(j.storage.enabled);
            });
        });
    });
});
