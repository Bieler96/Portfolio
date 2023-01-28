window.onscroll = () => {
    if (document.body.scrollTop > 80 || document.documentElement.scrollTop > 80) {
        document.getElementsByClassName("dabi-appbar-title")[0].style.fontSize = "25px";
        document.getElementsByClassName("dabi-appbar-title")[0].style.transform = "translateX(-5%)";
        document.getElementsByClassName("dabi-appbar-title")[0].style.left = "5%";
        document.getElementsByClassName("dabi-appbar")[0].style.padding = "30px 10px";
    }
    else{
        document.getElementsByClassName("dabi-appbar-title")[0].style.fontSize = "35px";
        document.getElementsByClassName("dabi-appbar-title")[0].style.transform = "translateX(-50%)";
        document.getElementsByClassName("dabi-appbar-title")[0].style.left = "50%";
        document.getElementsByClassName("dabi-appbar")[0].style.padding = "180px 10px";
    }
}