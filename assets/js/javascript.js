


function myFunction() {

    console.log("working")
}
tiles = document.getElementsByClassName('col-md-3')
for (var i = 0; i < tiles.length; i++) {
    tiles[i].addEventListener('click', myFunction, false);
}