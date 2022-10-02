function tempAlert(msg, duration) {
    var el = document.createElement("div");
    el.setAttribute("class", "notification is-success");
    el.innerHTML = msg;
    setTimeout(function () {
        el.parentNode.removeChild(el);
    }, duration);
    document.getElementById("birthdays_box").before(el);
}

