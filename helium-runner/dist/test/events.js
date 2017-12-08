"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
require("mocha");
const chai_1 = require("chai");
const mock = require("./mock");
const events = require("../src/events");
describe("events", function () {
    // Here, we just want to ensure that objects exported to brigadier are
    // available.
    it("has .BrigadeEvent", function () {
        chai_1.assert.property(events, "BrigadeEvent");
    });
    it("has .Project", function () {
        chai_1.assert.property(events, "Project");
    });
    it("has .EventRegistry", function () {
        chai_1.assert.property(events, "EventRegistry");
    });
    describe("EventRegistry", function () {
        let er;
        beforeEach(function () {
            er = new events.EventRegistry();
        });
        describe("#constructor", function () {
            it("registers 'ping' handler", function () {
                chai_1.assert.isTrue(er.has("ping"));
            });
        });
        describe("#on", function () {
            it("registers an event handler", function () {
                er.on("my-event", (e, p) => { });
                chai_1.assert.isTrue(er.has("my-event"));
            });
        });
        describe("#fire", function () {
            it("executes an event handler", function () {
                let fired = false;
                let ename = "my-event";
                let myEvent = mock.mockEvent();
                let myProj = mock.mockProject();
                myEvent.type = ename;
                er.on(ename, (e, p) => { fired = true; });
                er.fire(myEvent, myProj);
                chai_1.assert.isTrue(fired);
            });
            context("when calling an event with no handler", function () {
                it("does not cause an error (does nothing)", function () {
                    // We want this behavior because we don't want to force every brigade.js
                    // to implement every possible event.
                    let myEvent = mock.mockEvent();
                    let myProj = mock.mockProject();
                    myEvent.type = "no-such-event";
                    er.fire(myEvent, myProj);
                });
            });
        });
    });
});
