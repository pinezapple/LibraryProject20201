$(document).ready(function() {
    document.getElementById("delete").disabled = true;
    document.getElementById("edit").disabled = true;
  
  var table = $('#datatable-basic').DataTable();

  $('#datatable-basic tbody').on( 'click', 'tr', function () {
      if ( $(this).hasClass('selected') ) {
          $(this).removeClass('selected');
          document.getElementById("delete").disabled = true;
          document.getElementById("edit").disabled = true;
          // button.disabled = true
      }
      else {
          table.$('tr.selected').removeClass('selected');
          $(this).addClass('selected');
          document.getElementById("delete").disabled = false;
          document.getElementById("edit").disabled = false;
          // button.disabled = false
      }
  } );

   // Fetch data 
   fetch ('http://localhost:3000/users',{
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    },
  })
  .then(response => response.json())
  .then(usersList =>{
    usersList.forEach(user =>{

      let data =[]
      // get doc data
      let userid= `${user.id}`
      let username =`${user.username}`
      let name = `${user.name}`
      let sex = `${user.sex}`
      let phone = `${user.phone_number}`
      let role = `${user.role}`
     

      data.push (userid,username,name,sex,phone,role)

      // add to table
      table.row.add(data).draw()
    })
  })
  .catch(error => {
  alert('Something went wrong', error)
  });


  // Delete selected row
  $('#delete').click( function () {
    if (confirm('Are you sure you want to delete the row?')){
        let userid =  table.row('.selected').data()[0]
        let url = 'http://localhost:3000/users/' +userid
        fetch( url, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            }
            })
        .catch(error => alert('something went wrong',error))
        // delete from table
       table.row('.selected').remove().draw( false );
  } ;
  });

  // Edit selected row
  $('#edit').click( function () {
    
    let userid =  table.row('.selected').data()[0]
    
    let url = 'http://localhost:3000/users/' +userid

    fetch (url,{
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        },
      })
      .then(response => response.json())
      .then(function(user){

          let userId =`${user.id}`
          let username =`${user.username}`
          let password = `${user.password}`
          let name = `${user.name}`
          let dob = `${user.dob}`
          let sex = `${user.sex}`
          let phone = `${user.phone_number}`
          let role = `${user.role}`
        
        $("#userId").val(userId)
        $("#username").val(username)
        $("#password").val(password)
        $("#name").val(name)
        $("#DOB").val(dob)
        $("#sex").val(sex)
        $("#phone").val(phone)
        $("#role").val(role)

      })
      .catch(error => {
      alert('Something went wrong',error)
      });
  });
});

// sumbit form

$("#sumbit").on('click',function(){
    
    let userid = document.getElementById('userId').value;
    let username = document.getElementById('username').value;
    let password = document.getElementById('password').value;
    let name = document.getElementById('name').value;
    let dob = document.getElementById('DOB').value;
    let sex = document.getElementById('sex').value;
    let phone = document.getElementById('phone').value;
    let role = document.getElementById('role').value;

    let url = 'http://localhost:3000/users/' +userid

      fetch(url, {
         method: 'PUT',
         headers: {
          'Content-Type': 'application/json'
          },
         body:JSON.stringify({name:name,username:username,password:password,phone_number:phone,dob:dob,sex:sex,role:role})

     }).then((res) => res.json())
     .then(result => alert("Success", result))
     .catch((err)=>alert("Something went wrong",err))

    // e.preventDefault();

  })

  function newUser(){
    // event.preventDefault();
  
    let username = document.getElementById('username').value;
    let password = document.getElementById('password').value;
    let name = document.getElementById('name').value;
    let dob = document.getElementById('DOB').value;
    let sex = document.getElementById('sex').value;
    let phone = document.getElementById('phone').value;
    let role = document.getElementById('role').value;


    fetch('http://localhost:3000/users', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
          },
        body:JSON.stringify({username:username,password:password,name:name,phone_number:phone,dob:dob,sex:sex,role:role})
       
    }).then((res) => res.json())
    .then(result => alert("Added new user", result))
    .catch((err)=>alert("Something went wrong",err))
  
  }

