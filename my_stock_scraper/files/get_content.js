function get_cnt(what){
    var request = new XMLHttpRequest();
    request.open('GET', 'http://pawel22729.pythonanywhere.com/static/'+what+'.content', true);
    request.send();
    request.onreadystatechange = function () {
        if (request.readyState === 4 && request.status === 200) {
            var type = request.getResponseHeader('Content-Type');
            if (type.indexOf("text") !== 1) {
                document.getElementById("myself_descr").innerHTML = request.responseText;
            }
        }
    }
}