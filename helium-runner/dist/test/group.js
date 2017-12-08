"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
require("mocha");
const chai_1 = require("chai");
const group = require("../src/group");
const mock = require("./mock");
describe("group", function () {
    // Group tests
    describe("Group", function () {
        let g;
        beforeEach(function () {
            g = new group.Group();
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
        describe("static #runEach", function () {
            it("runs each job in order", function (done) {
                let j1 = new mock.MockJob("first");
                let j2 = new mock.MockJob("second");
                let j3 = new mock.MockJob("third");
                // This ensures that if the jobs were not executed in sequence,
                // 1 and 2 would finish before 3.
                j3.delay = 5;
                group.Group.runEach([j1, j2, j3]).then((rez) => {
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
                    group.Group.runEach([j1, j2, j3]).then((rez) => {
                        done("expected error on job 2");
                    }).catch((msg) => {
                        chai_1.assert.equal(msg, "Failed");
                        done();
                    });
                });
            });
        });
        describe("static #runAll", function () {
            it("runs jobs asynchronously", function (done) {
                let j1 = new mock.MockJob("first");
                let j2 = new mock.MockJob("second");
                let j3 = new mock.MockJob("third");
                group.Group.runAll([j1, j2, j3]).then((rez) => {
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
                    group.Group.runAll([j1, j2, j3]).then((rez) => {
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
