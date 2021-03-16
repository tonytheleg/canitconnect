function curlCompile() {
    var hostname = document.getElementById("hostname").value;
    var inputs = document.querySelectorAll('[data-input]');
    for (input in inputs) {
        console.log(input.values);
    }

    var curlOutput = document.getElementById("curlOutput").innerHTML = "<p>" +
    "Hostname: " + hostname + "<br>" + "</p>";
}

function tracerouteCompile() {
    var hostname = document.getElementById("hostname").value;
    var inputs = document.querySelectorAll('[data-input]');
    for (input in inputs) {
        console.log(input.values);
    }

    var tracerouteOutput = document.getElementById("tracerouteOutput").innerHTML = "<p>" +
    "Hostname: " + hostname + "<br>" + "</p>";
}

function netcatCompile() {
    var hostname = document.getElementById("hostname").value;
    var inputs = document.querySelectorAll('[data-input]');
    for (input in inputs) {
        console.log(input.values);
    }

    var netcatOutput = document.getElementById("netcatOutput").innerHTML = "<p>" +
    "Hostname: " + hostname + "<br>" + "</p>";
}