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
  var raw =
      '{\n  "username":"' +
      user +
      '",\n  "password":"' +
      pass +
      '"\n}';
  fetch('http://localhost:11001/p/login',{
        method: 'POST',
        headers: {
        'Content-Type': 'application/json'
        },
        body: raw
    }).then(response => response.json())
    .then(res =>{
      var user_id = result.user_id;        
      sessionStorage.setItem("id", user_id);
      window.location.href = "../../pages/dashboards/dashboard.html";
    })
    .catch(error => {
        alert("Something went wrong, please check your username or password", error)
        });
}

$("#loginButton").on("click", function () {
  let username = $("#email").val();
  let password = $("#password").val();

  console.log(username);
  console.log(password);

  loginRequest(username, password);
});
});