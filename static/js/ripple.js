let buttons = document.getElementsByClassName("dabi-ripple");

Array.prototype.forEach.call(buttons, (btn) => {
    btn.addEventListener("click", createRipple);
});

function createRipple(e){
    if(this.getElementsByClassName("ripple").length > 0)
        this.removeChild(this.childNodes[1]);

    let circle = document.createElement("div");
    this.appendChild(circle);

    var d = Math.max(this.clientWidth, this.clientHeight);
    circle.style.width = circle.style.height = d + "px";

    circle.style.left = e.clientX - this.offsetLeft - d / 2 + "px";
    circle.style.top = e.clientY - this.offsetTop - d / 2 + "px";

    circle.classList.add("ripple");
}

let buttonsDark = document.getElementsByClassName("dabi-ripple-dark");

Array.prototype.forEach.call(buttonsDark, (btn) => {
    btn.addEventListener("click", createDarkRipple);
});

function createDarkRipple(e){
    if(this.getElementsByClassName("ripple-dark").length > 0)
        this.removeChild(this.childNodes[1]);

    let circle = document.createElement("div");
    this.appendChild(circle);

    var d = Math.max(this.clientWidth, this.clientHeight);
    circle.style.width = circle.style.height = d + "px";

    circle.style.left = e.clientX - this.offsetLeft - d / 2 + "px";
    circle.style.top = e.clientY - this.offsetTop - d / 2 + "px";

    circle.classList.add("ripple-dark");
}