$(document).ready(function () {
  let id = sessionStorage.getItem('id');
  if (id != null){
    $("#librarian").val(id);
}
else{
    alert("Vui lòng đăng nhập để truy cập trang!");
    window.location.href = "../../pages/account/login.html";
}
});
