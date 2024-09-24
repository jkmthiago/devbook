$('#register_form').on('submit', registerUser)

function registerUser(event) {
    event.preventDefault();
    console.log("Formulário enviado");

    if ($('#password').val() !== $('#confirme_password').val()) {
        Swal.fire(
            "Senhas não conferem",
            "Por gentileza confirme se a senha inserida é a mesma nos dois campos",
            "error"
        )
        $('#confirme_password').focus();
        return;
    }

    $.ajax({
        url: "/register",
        method: "POST",
        data: {
            name: $('#name').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            password: $('#password').val()
        },
        dataType: "text"
    }).done(function () {
        Swal.fire(
            'Sucesso',
            'Usuário Cadastrado com Sucesso',
            'success'
        ).then(function () {
            window.location.href = "/login"
        })
    }).fail(function (erro) {
        console.log(erro)
        Swal.fire(
            'Erro',
            'Erro ao cadastrar o usuário!',
            'error'
        );
    });       
}
