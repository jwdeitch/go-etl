$(document).ready(function () {

    // Thanks! http://stackoverflow.com/a/16404503/4603498
    var clicked = false;
    $(".uploadInput").change(function () {
        if ($(this).val() != "") {
            $('.container').addClass('spinner').html("").height(207);
            $('.upload').submit();
        }
    });

    $(".upload").click(function () {
        if (!clicked) {
            clicked = true;
            $('.uploadInput').click();
            clicked = false;
        }
    }).on('submit', function (e) {
        e.preventDefault();
        $.ajax({
            url: "http://localhost:9090/recieve",
            type: "POST",
            contentType: false,
            processData: false,
            data: $(this).serialize()
        });
    });
});