$(function () {
    var array = [];

    var elements = document.body.getElementsByTagName("h4");

    for (var i = 0; i < elements.length; i++) {
        var current = elements[i];
        //console.log(current, current.children)
        if (current.children.length === 0 && current.textContent.replace(/ |\n/g, '') !== '') {
            array.push(current.textContent);
        }
    }
    //console.log(array)
    $("#search").autocomplete({
        source: array
    });
});

function ouvrirPage() {
    var array = [];

    var elements = document.body.getElementsByTagName("h4");

    for (var i = 0; i < elements.length; i++) {
        var current = elements[i];
        //console.log(current, current.children)
        if (current.children.length === 0 && current.textContent.replace(/ |\n/g, '') !== '') {
            array.push(current.textContent);
        }
    }
    var a = document.getElementById("search").value;
    if(array.includes(a)){
        console.log(array.indexOf(a))
        window.open(`/artist/${array.indexOf(a)}`);
    }
    
}