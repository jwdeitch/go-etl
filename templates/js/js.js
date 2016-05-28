$(document).ready(function () {

    // Thanks! http://stackoverflow.com/a/16404503/4603498
    var clicked = false;
    $(".uploadInput").change(function (){
        if ($(this).val() != "") {
            $('.container').addClass('spinner').html("").height(207)
        }
    });
    $(".upload").click(function () {
        if (!clicked) {
            clicked = true;
            $('.uploadInput').click();
            clicked = false;
        }
    });

});