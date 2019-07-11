'use strict';

function SendText(text) {
    let req = new XMLHttpRequest();
    req.open('POST', location.protocol.concat("//", window.location.host, "/post/", text));
    req.onload = function() {
        if (req.status === 200) {
            alert('Status OK Response: ' + req.responseText);
        }
        else {
            alert('Request failed.  Returned status: ' + req.status);
        };
    };
    req.send();
}

window.addEventListener('load', function () {
    document.getElementById('submit_button').onclick = () => {
        SendText(document.getElementById("username").value)
    };
    document.getElementById('clear_button').onclick = () => {
        document.getElementById("username").value = ""
        document.getElementById("petname").value = ""
    };
}, false);
