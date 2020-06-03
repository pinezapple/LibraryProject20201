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
  });

   // Fetch data 
   fetch ('http://localhost:11001/user/alluser',{
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
  })
  .then(response => response.json())
  .then(usersList =>{
    usersList.forEach(user =>{

      let data =[]
      // get doc data
      let userid= `${user.id_user}`
      let username =`${user.username}`
      let name = `${user.name}`
      let sex = `${user.sex}`
      let phone = `${user.phonenumber}`
      let dob = `${user.dob}`
      let role = `${user.role}`
     

      data.push (userid,username,name,sex,dob,phone,role)

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
        let userid = parseInt(table.row('.selected').data()[0],10)
        let url = 'http://localhost:11001/user/delete'
        fetch( url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({id_user:userid})
            })
        .catch(error => alert('something went wrong',error))
        // delete from table
       table.row('.selected').remove().draw( false );
  } ;
  });

  // Edit selected row
  $('#edit').click( function () {
    
    let userid =  table.row('.selected').data()[0]
    let username =  table.row('.selected').data()[1]
    let password =  ""
    let name =  table.row('.selected').data()[2]
    let sex =  table.row('.selected').data()[3]
    let dob =  table.row('.selected').data()[4]
    let phone =  table.row('.selected').data()[5]
    let role =  table.row('.selected').data()[6]
    
   
        
        $("#userId").val(userid)
        $("#username").val(username)
        $("#password").val(password)
        $("#name").val(name)
        $("#DOB").val(dob)
        $("#sex").val(sex)
        $("#phone").val(phone)
        $("#role").val(role)

      })

// sumbit form

$("#submit").on('click',function(){
    
    let userid = parseInt(document.getElementById('userId').value,10);
    let username = document.getElementById('username').value;
    let password = document.getElementById('password').value;
    let name = document.getElementById('name').value;
    let dob = document.getElementById('DOB').value;
    let sex = document.getElementById('sex').value;
    let phone = document.getElementById('phone').value;
    let role = document.getElementById('role').value;

    console.log(userid)
    console.log(username)
    console.log(password)
    console.log(name)
    console.log(dob)
    console.log(sex)
    console.log(phone)
    console.log(role)

    let url = 'http://localhost:11001/user/update' 

      fetch(url, {
         method: 'POST',
         headers: {
          'Content-Type': 'application/json'
          },
         body:JSON.stringify({id_user:userid,name:name,username:username,password:password,phonenumber:phone,dob:dob,sex:sex,role:role})

     }).then((res) => res.json())
     .then(result => alert("Success", result))
     .catch((err)=>alert("Something went wrong",err))

    // e.preventDefault();

  })
  // $("#addUser").on('click',function(){
  //   // event.preventDefault();
  
  //   let username = document.getElementById('username').value;
  //   let password = document.getElementById('password').value;
  //   let name = document.getElementById('name').value;
  //   let dob = document.getElementById('DOB').value;
  //   let sex = document.getElementById('sex').value;
  //   let phone = document.getElementById('phone').value;
  //   let role = document.getElementById('role').value;


  //   fetch('http://localhost:11001/user/save', {
  //       method: 'POST',
  //       headers: {
  //         'Content-Type': 'application/json'
  //         },
  //       body:JSON.stringify({username:username,password:password,name:name,phonenumber:phone,dob:dob,sex:sex,role:role})
       
  //   }).then((res) => res.json())
  //   .then(result => alert("Added new user", result))
  //   .catch((err)=>alert("Something went wrong",err))
  
})
