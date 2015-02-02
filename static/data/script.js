function get(path, callback) {
    var xhr = new window.XMLHttpRequest()
    xhr.open('GET', path);
    xhr.onreadystatechange = function() {
        if ( xhr.readyState != 4)
            return;
            callback(xhr.responseText);
    };
    xhr.send();
}

function getJson(path, callback) {
	get(path, function(data) {
		callback(JSON.parse(data))
	})
}
