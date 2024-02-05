function loadPage(){
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4) {
            var response = JSON.parse(this.responseText)
            if (response.error){
                alert(response.error)
                return
            }
            console.log(response.message)
            placeInfo(response.info)
        }
    };
    xhttp.withCredentials = true;
    xhttp.open("get", "/info/getinfo");
    xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhttp.send();
    return;
}

function placeInfo(info){
    var currentInfo = document.getElementById("info")
    currentInfo.innerHTML = info
}

function addItem(){
    var JSONItem = {
        "NewItemName": document.getElementById("input_1").value,
        "NewItemModel": document.getElementById("input_2").value,
        "NewItemProdYear": document.getElementById("input_3").value,
        "NewItemDescription": document.getElementById("input_4").value
    }
    var finalJSON = JSON.stringify(JSONItem)

    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4) {
            var response = JSON.parse(this.responseText)
            if (response.error){
                alert(response.error)
                return
            }
            //var element = document.getElementById("form_1");
            //element.reset()
            alert(response.message)
        }
    };
    xhttp.withCredentials = true;
    xhttp.open("post", "/db/additem");
    xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhttp.send(finalJSON);
    return;
}