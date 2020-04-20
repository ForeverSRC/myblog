function checkInput() {
    var name=document.getElementById("name");
    if(name.value.length==0) {
        alert("请输入待办事项");
        return false
    }
    return true
}

document.getElementsByClassName("addbtn")[0].addEventListener("click",sendThings);
function sendThings(e){
    e.preventDefault();//需要阻止默认事件，否则会发生跳转，url最后会多加上querystring
    if(checkInput()){
        var xhr=new XMLHttpRequest();
        xhr.open("POST","/todolist/todo",true);
        xhr.setRequestHeader("Content-type","application/json charset=UTF-8");
        var things=document.getElementById("name").value;
        var obj={"title":things,"status":false};
        var data=JSON.stringify(obj);
        xhr.send(data);
        xhr.onreadystatechange=function () {
            if(this.status==200&&this.readyState==4){
                var msg=JSON.parse(this.responseText);
                alert(msg["msg"]);
                document.getElementById("name").value="";//清空输入
                getAll(e);
            }
        };
    }
}


function getAll(e) {
    e.preventDefault();
    var xhr=new XMLHttpRequest();
    xhr.open("GET","/todolist/todo",true);
    xhr.send();
    xhr.onload=function () {
        if(xhr.status==200) {
            var user = JSON.parse(xhr.responseText);
            if (user["code"] == 666) {
                var data = user["data"];
                output = "";
                for (var i = 0; i < data.length; i++) {
                    output += `
                        <tr class="my_table">
                            <th>${data[i].id}</th>`;

                    if(data[i].status){
                        output+=`
                            <th style="text-decoration: line-through red">${data[i].title}</th>
                            <th>${data[i].created}</th>
                            <th>
                                <div class="btn-group" role="group">
                                    <span class="done" style="display:inline-block"></span>`;
                    }else{
                        output+=`
                            <th>${data[i].title}</th>
                            <th>${data[i].created}</th>
                            <th>
                                <div class="btn-group" role="group">
                                    <span class="correct" style="display:inline-block"></span>`;
                    }
                    output+=`       <span class="note" style="display:inline-block"></span>
                                    <span class="incorrect" style="display:inline-block"></span>
                                </div>
                            </th>
                        </tr>
                       `;
                }
                document.getElementById("todolist").innerHTML = output;
            }
        }
    }
}

var tbody=document.getElementById("todolist");
tbody.onclick=function (e) {
    var event=e||window.event;
    var target=event.target||event.srcElement;
    var tagClass=target.getAttribute("class");
    if(tagClass==="correct"){
        var id=target.parentElement.parentElement.previousElementSibling.previousElementSibling.previousElementSibling.innerText;
        var jsondata={"id":parseInt(id),"status":true};
        $.ajax({
           url:"/todolist/todo",
           type:"PUT",
            async:true,
           contentType:"application/json",
            data:JSON.stringify(jsondata),
            dataType:"json",
            success:function (msg) {
               alert(msg["msg"]);
            },
        });
        target.setAttribute("class","done");
        target.parentElement.parentElement.previousElementSibling.previousElementSibling.style.textDecoration="line-through red";
    }else if(tagClass==="done"){
        id=target.parentElement.parentElement.previousElementSibling.previousElementSibling.previousElementSibling.innerText;
        jsondata={"id":parseInt(id),"status":false};
        $.ajax({
            url:"/todolist/todo",
            type:"PUT",
            async:true,
            contentType:"application/json",
            data:JSON.stringify(jsondata),
            dataType:"json",
            success:function (msg) {
                alert(msg["msg"]);
            },
        });
        target.setAttribute("class","correct");
        target.parentElement.parentElement.previousElementSibling.previousElementSibling.style.textDecoration="";
    }else if(tagClass==="incorrect"){
        id=target.parentElement.parentElement.previousElementSibling.previousElementSibling.previousElementSibling.innerText;
        jsondata={"id":parseInt(id)};
        $.ajax({
            url:"/todolist/todo",
            type:"DELETE",
            async:true,
            contentType:"application/json",
            data:JSON.stringify(jsondata),
            dataType:"json",
            success:function (msg) {
                //从表格中删除事项
                var child=target.parentElement.parentElement.parentElement;
                var parent=child.parentElement;
                parent.removeChild(child);
                alert(msg["msg"]);
            },
        });
    }else if(tagClass==="note"){
        id=target.parentElement.parentElement.previousElementSibling.previousElementSibling.previousElementSibling.innerText;
        var url="/todolist/todo/modify/"+id;
        window.open(url,"_self");
    }
    else{}
};