var allCheckboxes = document.querySelectorAll('input[type=checkbox]');
var allPeople = Array.from(document.querySelectorAll('.people'));

var checked = {};

getChecked('artistGroup');
getChecked('creationDate');
getChecked('country');

Array.prototype.forEach.call(allCheckboxes, function (el) {

    el.addEventListener('change', toggleCheckbox);
});

function
toggleCheckbox(e) {

    getChecked(e.target.name);
    setVisibility();
}

function
getChecked(name) {
    checked[name] = Array.from(document.querySelectorAll('input[name=' + name + ']:checked')).map(function (el) {
        return el.value;
    });
}

function setVisibility() {
    allPeople.map(function (el) {
        var artistGroup = checked.artistGroup.length ? _.intersection(Array.from(el.classList), checked.artistGroup).length : true;
        var creationDate = checked.creationDate.length ? _.intersection(Array.from(el.classList), checked.creationDate).length : true;
        var country = checked.country.length ? _.intersection(Array.from(el.classList), checked.country).length : true;
        if (artistGroup && creationDate && country) {
            el.style.display = 'block';
        }   else {
            el.style.display = "none";
        }
    });
}