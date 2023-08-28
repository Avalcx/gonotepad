function uploadContent() {
    // If textarea value changes.
    if (content !== textarea.value) {
        var temp = textarea.value;
        var request = new XMLHttpRequest();

        request.open('POST', window.location.href, true);
        request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded; charset=UTF-8');
        request.onload = function() {
            if (request.readyState === 4) {
                // Request has ended, check again after 1 second.
                content = temp;
                setTimeout(uploadContent, 200);
            }
        }
        request.onerror = function() {
            // Try again after 1 second.
            setTimeout(uploadContent, 1000);
        }
        request.send('text=' + encodeURIComponent(temp));

        // Make the content available to print.
        printable.removeChild(printable.firstChild);
        printable.appendChild(document.createTextNode(temp));
    }
    else {
        // Content has not changed, check again after 1 second.
        setTimeout(uploadContent, 200);
    }
}

function cleanData(){
    var textName = window.location.pathname
    textName = textName.slice(1)
    var request = new XMLHttpRequest();
    request.open('GET', '/cleanData?textName=' + textName);
    request.send();
}    

var textarea = document.getElementById('content');
var printable = document.getElementById('printable');
var content = textarea.value;
var CleanTime = document.getElementById("cleantime").innerHTML; 
var timer 

// Make the content available to print.
printable.appendChild(document.createTextNode(content));
textarea.focus();

// add clean
textarea.addEventListener('input',()=>{
    clearTimeout(timer)
    timer = setTimeout(cleanData,CleanTime)
})

uploadContent();