function OpenLightbox(imgPath){
    let lightbox = document.getElementById("lightbox");
    let img = document.getElementById("lightboxImage");

    lightbox.style.display = "block";
    img.src = imgPath;
}

function CloseLightbox(){
    let lightbox = document.getElementById("lightbox");

    lightbox.style.display = "none";
}