function changeLocale(lang) {
    let currentLocation = location.href;
    $.ajax({
        type: 'GET',
        url: '/locale?locale='.concat(lang),
        dataType : "text",
        success: function () {
            location.href = currentLocation;
        }
    });
}

let state = false;

function showPassword() {
    let pass = $('#inputPassword');
    if (this.state){
        this.state = false;
        $('#show-password').removeClass('btn-primary');
        $('#show-password').addClass('btn-light');
        pass.attr('type', "password");

    }else{
        this.state = true;
        $('#show-password').removeClass('btn-light');
        $('#show-password').addClass('btn-primary');
        pass.attr('type', "text");
    }
}

function loadThing(thingURL){

    $.ajax({
        type: 'POST',
        url: thingURL,
        dataType : "html",
        success: function (data) {
            $('#mainPanel').empty();
            $('#mainPanel').append(data);
        }
    });

}

