const path = "/go/src/github.com/kubehelium/helium"
function goJob(name, tasks) {
    var j = new Job(name, "quay.io/deis/go-dev:v1.6.0")
    j.tasks = [
        "mkdir -p " + path,
        "cd " + path
    ];
    for (var i = 0; i < tasks.length; i++) {
        j.tasks.append(tasks[i]);
    }
    return j
}

function installGoDeps() {
    return goJob("install-go-dependencies", [
        "go get github.com/golang/dep/cmd/dep",
        "dep ensure"
    ])
}
exports.installGoDeps = installGoDeps;

function buildCLI() {
    return goJob("build-cli", ["go build -o ./he ./cli"]);
}
exports.buildCLI = buildCLI;

function testGo() {
    return goJob("test-go", ["go test ./..."]);
}
exports.testGo = testGo;
