"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
require("mocha");
const chai_1 = require("chai");
const brigade = require("../src/brigadier");
const mock = require("./mock");
// These tests are largely designed to ensure that the objects a script is likely
// to use are indeed exposed. Tests for the actual functionality of each are found
// in their respective libraries.
describe("brigadier", function () {
    it("has #fire", function () {
        chai_1.assert.property(brigade, "fire");
    });
    it("has .Job", function () {
        chai_1.assert.property(brigade, "Job");
    });
    it("has .Group", function () {
        chai_1.assert.property(brigade, "Group");
    });
    it("has .events", function () {
        chai_1.assert.property(brigade, "events");
    });
    // Events tests
    describe("events", function () {
        it("has #on", function () {
            chai_1.assert.property(brigade.events, "on");
        });
    });
    // Group tests
    describe("Group", function () {
        let g;
        beforeEach(function () {
            g = new brigade.Group();
        });
        describe("#add", function () {
            it("adds a job", function () {
                chai_1.assert.equal(g.length(), 0);
                let j = new mock.MockJob("hello");
                let j2 = new mock.MockJob("goodbye");
                g.add(j);
                g.add(j2);
                chai_1.assert.equal(g.length(), 2);
            });
        });
        describe("#runEach", function () {
            it("runs each job in order", function (done) {
                let j1 = new mock.MockJob("first");
                let j2 = new mock.MockJob("second");
                let j3 = new mock.MockJob("third");
                // This ensures that if the jobs were not executed in sequence,
                // 1 and 2 would finish before 3.
                j3.delay = 50;
                g.add(j1, j2, j3);
                g.runEach().then((rez) => {
                    chai_1.assert.equal(rez[0], j1.name);
                    chai_1.assert.equal(rez[1], j2.name);
                    chai_1.assert.equal(rez[2], j3.name);
                    done();
                });
            });
            context("when job fails", function () {
                it("stops processing with an error", function (done) {
                    let j1 = new mock.MockJob("first");
                    let j2 = new mock.MockJob("second");
                    j2.fail = true;
                    let j3 = new mock.MockJob("third");
                    g.add(j1, j2, j3);
                    g.runEach().then((rez) => {
                        done("expected error on job 2");
                    }).catch((msg) => {
                        chai_1.assert.equal(msg, "Failed");
                        done();
                    });
                });
            });
        });
        describe("#runAll", function () {
            it("runs jobs asynchronously", function (done) {
                let j1 = new mock.MockJob("first");
                let j2 = new mock.MockJob("second");
                let j3 = new mock.MockJob("third");
                g.add(j1, j2, j3);
                g.runAll().then((rez) => {
                    chai_1.assert.equal(rez.length, 3);
                    done();
                });
            });
            context("when job fails", function () {
                it("stops processing with an error", function (done) {
                    let j1 = new mock.MockJob("first");
                    let j2 = new mock.MockJob("second");
                    j2.fail = true;
                    let j3 = new mock.MockJob("third");
                    g.add(j1, j2, j3);
                    g.runAll().then((rez) => {
                        done("expected error on job 2");
                    }).catch((msg) => {
                        chai_1.assert.equal(msg, "Failed");
                        done();
                    });
                });
            });
        });
    });
});
