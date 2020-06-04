$("#login").on('click',function(){
    let username = document.getElementById('username').value
    let password = document.getElementById('password').value

    if (username == "" || password == ""){
        alert("Please fill all the input fields")
    }

    // get token?
    fetch('http://localhost:11001/p/login',{
        method: 'POST',
        headers: {
        'Content-Type': 'application/json'
        },
        body: JSON.stringify({username:username,password:password})
    }).then(response => response.json())
    .then(res =>{
        if (!res){
            return res.status(401).json({
                message: "auth failed"
            })
        } else {
            let token = res
            localStorage.setItem("token",token)
            window.location.href = "../dashboards/dashboard.html";
        }
    })
    .catch(error => {
        alert("Something went wrong, please check your username or password", error)
        });
})

function logout(){
    localStorage.removeItem('token')
}

