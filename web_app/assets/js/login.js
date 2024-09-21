$('#login_form').on('submit', loginUser)

function loginUser(event) {
    event.preventDefault();
    console.log("Formulário enviado");

    $.ajax({
        uls: "/login",
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
        alert("Usuário ou Senha Inválidos!")
    })
}
