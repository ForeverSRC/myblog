document.getElementById("download_pdf").addEventListener("click",Download_pdf);

function Download_pdf(e){
    console.log("1");
    e.preventDefault();
    //获得文章id
    var u=window.location.href;
    console.log(u);
    var id=u.charAt(u.length-1);
    //打开对应的网页
    alert("在弹出的窗口中右键选择”打印网页，即可保存为pdf格式的文件");
    window.open("/topic/download_pdf/"+id,"_blank","fullscreen=yes,directories=no,toolbar=no, menubar=no,titlebar=no, location=no, status=no");

}