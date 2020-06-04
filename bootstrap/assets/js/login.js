$("#login").on('click',function(){
    let username = document.getElementById('username').value
    let password = document.getElementById('password').value

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
            window.location.href = "bootstrap/pages/dashboards/dashboard.html";
        }
    })
    .catch(error => {
        alert("error with fetch login", error)
        });
})

function logout(){
    localStorage.removeItem('token')
}

