// TODO: allow for multiple files
const { events, Job, Group } = require("./balloon")

events.on("build", (evt, project) => {
    console.log("this is the build event!")
    var installDepsJob = installGoDeps();
    var buildCLIJob = buildCLI();
    var testAllGoJob = testGo();
    return Group.runAll([installDepsJob, buildCLIJob, testAllGoJob]);
});

const path = "/go/src/github.com/kubehelium/helium"
function goJob(name, tasks) {
    var j = new Job(name, "quay.io/deis/go-dev:v1.6.0")
    j.tasks = ["mkdir -p " + path, "cd " + path];
    for (var i = 0; i < tasks.length; i++) {
        j.tasks.push(tasks[i]);
    }
    return j
}

function installGoDeps() {
    return goJob("install-go-dependencies", [
        "go get github.com/golang/dep/cmd/dep",
        "dep ensure"
    ])
}

function buildCLI() {
    return goJob("build-cli", ["go build -o ./he ./cli"]);
}

function testGo() {
    return goJob("test-go", ["go test ./..."]);
}
