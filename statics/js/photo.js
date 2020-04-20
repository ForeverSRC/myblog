document.getElementById("returnAlbum").addEventListener("click",returnAlbum);
function returnAlbum(e) {
    e.preventDefault();
    window.open("/albums","_self");
}

function getAllPhotos(e) {
    e.preventDefault();
    //获得相册id
    var u=window.location.href;
    var id=u.charAt(u.length-1);
    var jsondata={"id":parseInt(id)};
    //获取所有图片的路径
    $.ajax({
        url:"/photo/get",
        type:"POST",
        async:true,
        contentType:"application/json",
        data:JSON.stringify(jsondata),
        dataType:"json",
        success:function (datas) {
            //插入图片
            if(datas["code"]===666){
                var total=datas["nums"];
                var data=datas["data"];
                var output="";
                var top_init=datas["islog"]?-390:0;
                var left_init=0;
                for(var j=0;j<total;j++){
                    //每行放6个
                    if((j%6)===0&&j>0){
                        top_init+=195;
                        left_init=0;
                    }
                    top_str=String(top_init)+"px";
                    left_str=String(left_init)+"px";
                    output+=`
                    <div class="photodiv" style="top:${top_str};left:${left_str}">
                        <div class="albumDel"></div>
                        <img src="${data[j].path}" class="photo" index="${data[j].id}">
                    </div>
                    `;
                    left_init+=180;
                    top_init-=195;
                }
                document.getElementById("totalPhoto").innerHTML = output;
            }else{
                alert("错误："+datas["msg"]);
            }
        }
    })
}

function addPhoto(e) {
    e.preventDefault();
    if($("#file").val() != ""){
        var formdata=new FormData();
        formdata.append('pic',$("#file")[0].files[0]);
        var u=window.location.href;
        var id=u.charAt(u.length-1);
        formdata.append("albumid",id);
        $.ajax({
            type: "POST",
            url:"/photo/data",
            processData:false,
            contentType: false,
            data:formdata,
            dataType: "json",
            success: function(data){
                if(data["code"]===666){
                    alert("上传成功");
                    getAllPhotos(e);
                }else{
                    alert("上传失败，错误：\n"+data["msg"]);
                }
            },
            error: function () {
                alert("上传失败");

            },
        });
    }else{
        alert("未选择文件");
    }
}

var photoSet=document.getElementById("totalPhoto");
photoSet.onclick=function(e){
    var event=e||window.event;
    var target=event.target||event.srcElement;
    var targetclass=target.getAttribute("class");
    if(targetclass==="photo"){
        //查看图片
        //1.显示遮罩层
        var backdiv=document.getElementById("back");
        backdiv.style.display="block";
        $("#back").height($(document).height());
        $("#back").width($(document).width());
        $(document).resize(function () {
            $("#back").height($(document).height());
            $("#back").width($(document).width());
        });
        var windowH=$(window).height();//获取当前窗口的高度
        var windowW=$(window).width();//获取当前窗口的宽度
        var sH=$(window).scrollTop();//滚动条距离顶部的距离
        var backH=$("#back").height();//div的高度
        var vTop=(windowH-backH)/2+sH;//获取position的top
        $(window).resize(function(){
            //当窗口变化时候重新获取高度和宽度
            windowH=$(window).height();
            windowW=$(window).width();
            vTop=(windowH-backH)/2+sH;
        });
        $(window).scroll(function () {
            var sH=$(window).scrollTop();
            var windowH=$(window).height();//获取当前窗口的高度
            var windowW=$(window).width();//获取当前窗口的宽度
            var backH=$("#back").height();//div
            var vTop=(windowH-backH)/2+sH;//获取position的top
            var vLeft=(windowW-$("#back").width())/2;
            $("#back").css({
                top:vTop,
                bottom:vTop,
                left:vLeft,
                right:vLeft,
            })
        });

        //2.显示大图片
        var showPic=document.getElementById("showpic");
        showPic.style.display="table";
        $("#showpic").css({width:backdiv.style.width,height:backdiv.style.height});
        var showBig=document.getElementById("showBig");
        showBig.setAttribute("src",target.getAttribute("src"));
        var imgleft=($("#showpic").width()-$("#showBig").width())/2;
        $("#showBig").css({
            display:"table-cell",
            position:"relative",
            left:imgleft,
            top:150,
            verticalAlign:"middle",
        });

        var closeBig=document.getElementById("bigPicClose");
        closeBig.style.display="block";
        closeBig.style.top=showBig.style.top;
        closeBig.style.left=showBig.style.left;
        //3.关闭显示
        closeBig.onclick=function (e) {
            e.preventDefault();
            backdiv.style.display="none";
            showPic.style.display="none";
            closeBig.style.display="none";
        }
    }else if(targetclass==="albumDel"){
        var id=target.nextElementSibling.getAttribute("index");
        id=parseInt(id);
        var u=window.location.href;
        var aid=u.charAt(u.length-1);
        var jsondata={"id":id,"aid":parseInt(aid)};
        $.ajax({
            url:"/photo/data",
            type:"DELETE",
            async:true,
            contentType:"application/json",
            data:JSON.stringify(jsondata),
            dataType:"json",
            success:function(datas){
                //获得返回值
                if(datas["code"]===666){
                    alert("删除成功");
                    getAllPhotos(e);
                }else{
                    alert("删除失败，错误：\n"+datas["msg"]);
                }
            }
        })
    }
};

var delPhotoCount=false;
function delPhotoSwitch(e) {
    e.preventDefault();
    var targets=document.getElementsByClassName("albumDel");
    if(delPhotoCount===false){
        delPhotoCount=true;
        for(var i=0;i<targets.length;i++){
            if(targets[i].parentElement.getAttribute("class")==="photodiv"){
                targets[i].style.display="block";
            }

        }
    }else{
        delPhotoCount=false;
        for(var j=0;j<targets.length;j++){
            if(targets[j].parentElement.getAttribute("class")==="photodiv"){
                targets[j].style.display="none";
            }

        }
    }
}

document.getElementById("file").addEventListener("change",uploaded);
function uploaded(e) {
    e.preventDefault();
    alert("文件选择成功");
    addPhoto(e);
}