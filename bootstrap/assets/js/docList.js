//datatable basic docList 
    var btn = document.querySelectorAll(".btn-outline-primary")
    for (var i=0; i<btn.length; i++ ){
        btn[i].disabled = true;
    }
     
  $('table.display').DataTable();
  var table = $('#datatable-basic').DataTable();
  var t = $('#datatable-buttons').DataTable();


  $('#datatable-basic tbody').on( 'click', 'tr', function () {
      if ( $(this).hasClass('selected') ) {
          $(this).removeClass('selected');
        for (var i=0; i<btn.length; i++ ){
            btn[i].disabled = true;
        }
          // button.disabled = true
      }
      else {
          table.$('tr.selected').removeClass('selected');
          $(this).addClass('selected');
        for (var i=0; i<btn.length; i++ ){
            btn[i].disabled = false;
        }
          // button.disabled = false
      }
  } );

  docList = document.querySelector("#docList")

  fetch ('http://localhost:11001/doc/alldoc',{
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
  })
  .then(response => response.json())
  .then(docList =>{
    docList.forEach(doc =>{

      var data =[]
      // get doc data
      var docId= `${doc.id_doc}`
      var docName = `${doc.doc_name}`
      var docAuthor = `${doc.doc_author}`
      var docType = `${doc.doc_type}`
      var docDescription = `${doc.doc_description}`
      var docDate = `${doc.updated_at}`
      var docStatus = `${doc.status}`

      data.push (docId,docName,docAuthor,docType,docDescription,docDate,docStatus)

      // add to table
      table.row.add(data).draw()
    })
  })
  .catch(error => {
  console.error('Error:', error);
  });

  

  // Delete selected row
  $('#delete').click( function () {
    if (confirm('Are you sure you want to delete the row?')){
        var documentId = table.row('.selected').data()[0]
        var url = 'http://localhost:11001/doc/delete'
        fetch( url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            }
            // them json dang {id: id}
            })
        .catch(error => console.log(error))
       table.row('.selected').remove().draw( false );
  } ;
  });

  // Edit selected row
  $('#edit').click( function () {
    
    var id =  table.row('.selected').data()[0]
    var name =  table.row('.selected').data()[1]
    var author = table.row('.selected').data()[2]
    var type = table.row('.selected').data()[3]
    var status = table.row('.selected').data()[6]
    var description = table.row('.selected').data()[4]

    $("#id").val(id)
    $("#name").val(name)
    $("#author").val(author)
    $("#type").val(type)
    $("#status").val(status)
    $("#description").val(description)

  });
  
  $("#sumbit").on('click',function(){
    

    let id = document.getElementById('id').value;
    let name = document.getElementById('name').value;
    let author = document.getElementById('author').value;
    let type = document.getElementById('type').value;
    let status = document.getElementById('status').value;
    let description = document.getElementById('description').value;

    let url = 'http://localhost:11001/doc/update'

      fetch(url, {
         method: 'POST',
         headers: {
          'Content-Type': 'application/json; charset=UTF-8'
          },
         body:JSON.stringify({doc_name:name, doc_author:author, doc_type:type,status:status, doc_description:description})

     }).then((res) => res.json())
     .then(result => alert("Fixed document", result))
     .catch((err)=>console.log(err))

    // e.preventDefault();

  })

// reset borrow form
function reset() {
  document.getElementById("borrowForm").reset();
}

 // add docId to form
 $('#add').click( function () {
    var docId =  table.row('.selected').data()[0]
    var docName =  table.row('.selected').data()[1]
    var author =  table.row('.selected').data()[2]
    var docType =  table.row('.selected').data()[3]
    // display document
    $("#docId").val(docId)
    $("#docName").val(docName)
    $("#author").val(author)
    $("#docType").val(docType)
  });

// datatable-buttons docList
fetch ('http://localhost:11001/doc/alldoc',{
  method: 'POST',
  headers: {
    'Content-Type': 'application/json'
  },
}).then(response => response.json())
  .then(docList =>{
    console.log(docList)
    docList.forEach(doc =>{

      var data =[]
      // get doc data
      var docId= `${doc.id_doc}`
      var docName = `${doc.doc_name}`
      var docAuthor = `${doc.doc_author}`
      var docType = `${doc.doc_type}`
      var docDate = `${doc.updated_at}`
      var docStatus = `${doc.status}`

      data.push (docId,docName,docAuthor,docType,docDate,docStatus)

      // add to table
      t.row.add(data).draw()
    })
  })
  .catch(error => {
  console.error('Error:', error);
  });

// add new doc
// document.getElementById('newDoc').addEventListener('submit', newDoc);

function newDoc(event){
  event.preventDefault();

  let name = document.getElementById('name').value;
  let author = document.getElementById('author').value;
  let type = document.getElementById('type').value;
  let description = document.getElementById('description').value;

  fetch('http://localhost:11001/doc/save', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json; charset=UTF-8'
        },
      body:JSON.stringify({doc_name:name, doc_author:author, doc_type:type, doc_description:description})
     
  }).then((res) => res.json())
  .then(result => alert("Added new document", result))
  .catch((err)=>alert("Something went wrong",err))

  document.getElementById("newDoc").reset();
}

function newForm(e){
  e.preventDefault();

  let librarian = parseInt(document.getElementById('librarian').value,10);
  let student = parseInt(document.getElementById('studentId').value,10);
  let days = parseInt(document.getElementById('days').value,10);
  let docId = parseInt(document.getElementById('docId').value,10);

  fetch('http://localhost:11001/doc/saveForm', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json; charset=UTF-8'
        },
      body:JSON.stringify({id_doc:docId, id_cus:student, id_lib:librarian , status: 1, ttl:days})
     
  }).then((res) => res.json())
  .then(result => alert("Success", result))
  .catch((err)=>alert("Something went wrong",err))

  document.getElementById("borrowForm").reset();
}
