let token1 = localStorage.getItem("token")

if(token1 == null){
    alert("Please login");
    window.location.href = "../../pages/login/login.html";
}

function logout(){
    localStorage.removeItem("token")
    localStorage.clear()
}