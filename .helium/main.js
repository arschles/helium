const { events, Job, Group } = require("brigadier")

const { installGoDeps, buildCLIJob, testGo } = require("go");

event.on("build", (evt, project) => {
    var installDepsJob = installGoDeps();
    var buildCLIJob = buildCLI();
    var testAllGoJob = testGo();
    return Group.runAll([installDepsJob, buldCLIJob, testAllGoJob]);
});
