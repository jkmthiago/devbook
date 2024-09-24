$('#login_form').on('submit', loginUser)

function loginUser(event) {
    event.preventDefault();
    console.log("Formulário enviado");

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            password: $('#password').val()
        },
        dataType: "text"
    }).done(function () {
        window.location = "/home"
    }).fail(function (erro) {
        console.log(erro)
        Swal.fire(
            "Usuário ou Senha Incorreta!",
            "Por gentileza confirme se o usuáro ou a senha inserida estão corretos!",
            "error"
        )
    })
}
