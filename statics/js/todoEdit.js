document.getElementById("todo_edit").addEventListener("click",editDone);
document.getElementById("todo_edit_return").addEventListener("click",editReturn);
function editDone(e) {
    e.preventDefault();//需要阻止默认事件，否则会发生跳转，url最后会多加上querystring

    //获得ID
    var u=window.location.href;
    var id=u.charAt(u.length-1);
    id=parseInt(id);
    //获得正文
    var content=document.getElementById("todoContent").innerText;
    //组装JSON
    var obj={"id":id,"content":content};
    $.ajax({
        url:"/todolist/todo/modify",
        type:"PUT",
        async:true,
        contentType:"application/json",
        data:JSON.stringify(obj),
        dataType:"json",
        success:function (msg) {
            alert(msg["msg"]);
            window.open("/todolist","_self");
        },
    });

}
function editReturn(e) {
    e.preventDefault();//需要阻止默认事件，否则会发生跳转，url最后会多加上querystring
    window.open("/todolist","_self");
}