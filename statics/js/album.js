function getAllAlbums(e) {
    e.preventDefault();
    $.ajax({
        url:"/albums/data",
        type:"GET",
        async:true,
        contentType:"application/json",
        dataType:"json",
        success:function(datas){
            if(datas["code"]===666){

                var total=datas["nums"];

                var data=datas["data"];
                var output="";
                var top_init=datas["islog"]?-390:0;
                var left_init=datas["islog"]?380:0;
                var i=datas["islog"]?2:0;
                for(var j=0;j<total;j++){
                    //每行放6个
                    if(((j+i)%6)===0&&j>0){
                        top_init+=195;
                        left_init=0;
                    }
                    top_str=String(top_init)+"px";
                    left_str=String(left_init)+"px";
                    output+=`
                    <div class="album" style="top:${top_str};left:${left_str}">
                        <div class="albumDel"></div>
                        <a href="/albums/album/${data[j].id}">
                            <img src="${data[j].frontpage}" class="frontpage">
                        </a>
                        <div class="albumName">
                            <a href="/albums/album/${data[j].id}" style="text-decoration: none" index="${data[j].id}">${data[j].name}</a>
                        </div>
                    </div>
                    `;
                    left_init+=190;
                    top_init-=195;
                }
                document.getElementById("totalAlbum").innerHTML = output;
            }else{
                alert("错误："+datas["msg"]);
            }
        }
    })
}


function addAlbum(e){
    e.preventDefault();
    //弹窗获取相册名
    var name=window.prompt("请输入相册名");
    if(name==null)
        return;
    if(name.length==0) {
        alert("相册名为空");
    }else{
        //相册名传到服务器
        jsondata={
            "name":name,
        };
        $.ajax({
            url:"/albums/data",
            type:"POST",
            async:true,
            contentType:"application/json",
            data:JSON.stringify(jsondata),
            dataType:"json",
            success:function(datas){
                //获得返回值
                if(datas["code"]===666){
                    alert("创建成功");
                    getAllAlbums(e);
                }else{
                    alert("创建失败，错误：\n"+datas["msg"]);
                }
            }
        })
    }
}

var delCount=false;
function delAlbumSwitch(e){
    e.preventDefault();
    var targets=document.getElementsByClassName("albumDel");
    if(delCount===false){
        delCount=true;
        for(var i=0;i<targets.length;i++){
            targets[i].style.display="block";
        }
    }else{
        delCount=false;
        for(var j=0;j<targets.length;j++){
            targets[j].style.display="none";
        }
    }
}

var albumSet=document.getElementById("totalAlbum");
albumSet.onclick=function (e) {
    var event=e||window.event;
    var deltarget=event.target||event.srcElement;
    var delclass=deltarget.getAttribute("class");
    if(delclass==="albumDel"){
        var albumName=deltarget.nextElementSibling.nextElementSibling.firstElementChild.innerHTML;
        confirmMsg="即将删除相册：\n"+albumName;
        var r=confirm(confirmMsg);
        if (r===true){
            var id=deltarget.nextElementSibling.nextElementSibling.firstElementChild.getAttribute("index");
            id=parseInt(id);
            var jsondata={"id":id,"name":albumName};
            $.ajax({
                url:"/albums/data",
                type:"DELETE",
                async:true,
                contentType:"application/json",
                data:JSON.stringify(jsondata),
                dataType:"json",
                success:function(datas){
                    //获得返回值
                    if(datas["code"]===666){
                        alert("删除成功");
                        getAllAlbums(e);
                    }else{
                        alert("删除失败，错误：\n"+datas["msg"]);
                    }
                }
            })
        }
    }else{}
};