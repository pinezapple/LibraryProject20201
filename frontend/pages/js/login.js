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
    '{\n  "username" : "' +
    user +
    '",\n  "password" : "' +
    pass +
    '"\n   \n}';

  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("http://25.43.134.201:8080/user/login", requestOptions)
    .then((response) => response.json())
    .then((result) => {
      if (result.message == "login success") {
        var token = result.token.access_token;
        var id = result.id;
        var name = result.name;
        sessionStorage.setItem("token", token);
        sessionStorage.setItem("id", token);
        sessionStorage.setItem("name", token);
        window.location.href = "../../pages/dashboards/dashboard.html";
      } else if (result.message == "Invalid login details") {
        alert("Thông tin đăng nhập chưa đúng");
      } else if (result.message == "Invalid form") {
        alert("Thông tin điền chưa hợp lệ!");
      } else {
        return;
      }
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