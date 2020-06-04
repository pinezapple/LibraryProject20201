let btn = document.querySelector(".btn-outline-primary")
btn.disabled = true;

var table = $('#datatable-basic').DataTable();
var t = $('#datatable-buttons').DataTable();

$('#datatable-basic tbody').on( 'click', 'tr', function () {
    if ( $(this).hasClass('selected') ) {
        $(this).removeClass('selected');
          btn.disabled = true;
        // button.disabled = true
    }
    else {
        table.$('tr.selected').removeClass('selected');
        $(this).addClass('selected');
          btn.disabled = false;
        // button.disabled = false
    }
} );

// fetch data

fetch ('http://localhost:11001/doc/allform',  {
  method: 'POST',
  credentials: "omit",
  headers: {
     'Content-Type': 'application/json',
     Authorization: "Bearer " + token,
    },
  })
  .then(response => response.json())
  .then(formList =>{
    formList.forEach(form =>{

      let data =[]
      // get doc data
    
      let id = `${form.id_borrow}`
      let docId= `${form.id_doc}`
      let docName = `${form.doc_name}`
      let cusId = `${form.id_cus}`
      let status = `${form.status}`

      data.push (id,docId,docName,cusId,status)

      // add to table
      table.row.add(data).draw()
    })
  })
  .catch(error => {
  console.error('Error:', error);
  });

// update borrow form
$('#update').click(function() {
    let formId =  parseInt(table.row('.selected').data()[0],10)
    let docId = parseInt(table.row('.selected').data()[1],10)
    let docName = table.row('.selected').data()[2]
    let userId = table.row('.selected').data()[3]
    let status = table.row('.selected').data()[4]

    $("#formId").val(formId)
    $("#docId").val(docId)
    $("#name").val(docName)
    $("#userId").val(userId)
    $("#status").val(status)
})

$('#submit').on('click',function(){
    let formId = parseInt(document.getElementById('formId').value,10)
    let docId = parseInt(document.getElementById('docId').value,10)
    let e= document.getElementById('status').value
    // let e = $("#status").val()
    //let status = parseInt(e.options[e.selectedIndex].value,10)
    let option = $("#status").val()
    let status = parseInt(option,10)

    fetch('http://localhost:11001/doc/updateStatus',   {
      method: 'POST',
      credentials: "omit",
      headers: {
         'Content-Type': 'application/json',
         Authorization: "Bearer " + token,
        },
      body:JSON.stringify({form_id:formId, doc_id:docId, status: status})
     
  }).then((res) => res.json())
  .then(result => alert("Success", result))
  .catch((err)=>alert("Something went wrong",err))
})

