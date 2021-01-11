$(document).ready(function () {


//Check if enough input
function checkInput() {
  var username = $("#username").val();
  var password = $("#password").val();
  if (username != "" && password != "") {
    $("#loginButton").removeClass("disabled");
    $("#loginButton").prop('disabled', false);
  } else {
    $("#loginButton").addClass("disabled");
    $("#loginButton").prop('disabled', true);
  }
}
setInterval(checkInput, 300);

//Login request
function loginRequest(user, pass) {
  var myHeaders = new Headers();
  myHeaders.append("Content-Type", "text/plain");

  var raw =
    '{\n  "username":"'+
    user +
    '",\n  "password":"'+
    pass +
    '"\n   \n}';

  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("http://localhost:11001/p/login", requestOptions)
    .then((response) => response.json())
    .then((result) => {
        var user_id = result.user_id;        
        sessionStorage.setItem("id", user_id);
        window.location.href = "../../pages/dashboards/dashboard.html";
    })
    .catch((error) => {
      console.log("Không kết nối được tới máy chủ", error);
      alert("Không kết nối được tới máy chủ");
    });
}

$("#loginButton").on("click", function () {
  let username = $("#username").val();
  let password = $("#password").val();

  loginRequest(username, password);
});
});