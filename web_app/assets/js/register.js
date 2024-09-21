$('#register_form').on('submit', registerUser)

function registerUser(event) {
    event.preventDefault();
    console.log("Formulário enviado");

    if ($('#password').val() !== $('#confirme_password').val()) {
        alert("As senhas não coincidem");
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
        alert("Usuário cadastrado com sucesso!");
    }).fail(function (erro) {
        console.log(erro)
        alert("Erro ao cadastrar o usuário!");
    });       
}
