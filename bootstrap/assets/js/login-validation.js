let token = localStorage.getItem("token")

if(token == null){
    alert("Please login");
    window.location.href = "bootstrap/pages/login/login.html";
} else {
    let request = {
      //   method:"POST",
        credentials: "omit",
      // headers: {
        Authorization: "Bearer " + token,
        "Content-Type": "application/json",
      }
    }

