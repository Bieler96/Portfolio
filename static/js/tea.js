function checkElementLocation() {
    let w = window;
    let bottomOfWindow = getYPosition() + w.innerHeight;

    let elements = document.getElementsByClassName("dabi-tea");

    for (let i = 0; i < elements.length; i++) {
        let temp = elements[i];
        let bottomOfObject = getTopPos(temp) + temp.offsetHeight;

        if (bottomOfWindow > bottomOfObject) {
            temp.classList.add("dabi-tea-active");
        }
    }
}

function getTopPos(el) {
    for (var topPos = 0;
        el != null;
        topPos += el.offsetTop, el = el.offsetParent);
    return topPos;
}

function getLeftPos(el) {
    for (var leftPos = 0;
        el != null;
        leftPos += el.offsetLeft, el = el.offsetParent);
    return leftPos;
}

function getYPosition() {
    var top = window.pageYOffset || document.documentElement.scrollTop
    return top;
}

checkElementLocation();

document.addEventListener("scroll", (event) => {
    checkElementLocation();
});

function isInViewport(element) {
    const rect = element.getBoundingClientRect();
    return (
        rect.top >= 0 &&
        rect.left >= 0 &&
        rect.bottom <= (window.innerHeight || document.documentElement.clientHeight) &&
        rect.right <= (window.innerWidth || document.documentElement.clientWidth)
    );
}