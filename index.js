//上传本地图片并预览
function preview(file) {
    var prevDiv = document.getElementById('preview');
    if (file.files && file.files[0]) {
        var reader = new FileReader();
        reader.onload = function(evt) {
            prevDiv.innerHTML = '<img src="' + evt.target.result + '" />';
        }
        reader.readAsDataURL(file.files[0]);
    } else {
        prevDiv.innerHTML = '<div class="img" style="filter:progid:DXImageTransform.Microsoft.AlphaImageLoader(sizingMethod=scale,src=\'' + file.value + '\'"></div>';
    }
}

//todo 上传图片URL，并预览，测试未通过的
function urlPreview() {
    var urlInput = document.getElementById('avatarSlect')

    urlInput.addEventListener('change',function(event){
        var imgShow = document.getElementById('preview'),
            files = this.files,    //原书这里是event.files，经过测试，修正为this.files
            url = window.URL.createObjectURL(files[0])    //创建URL对象
        if (url) {
            if (/image/.test(files[0].type)) {    //如果上传的文件类型为image
                imgShow.innerHTML = '<img src="'+url+'">'    //将图片直接插入到页面中
            }else{
                imgShow.innerHTML = "not an image"
            }
        }else{
            imgShow.innerHTML = 'your browser doesnt support obj urls!'
        }
    })

}