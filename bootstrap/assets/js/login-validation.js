let token = localStorage.getItem("token")

if(token == null){
    alert("Please login");
    window.location.href = "../../pages/login/login.html";
}

function logout(){
    localStorage.removeItem("token")
    localStorage.clear()
}