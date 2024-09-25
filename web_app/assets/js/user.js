$('#follow').on('click', follow);
$('#unfollow').on('click', unfollow);
$('#delete-user').on('click', deleteUser);
$('#edit-user').on('submit', edit);
$('#update-password').on('submit', updatePassword);

function follow() {
    const user_id = $(this).data('user-id')
    $(this).prop('disabled', true)

    $.ajax({
        url: `/users/${user_id}/follow`,
        method: "POST"
    }).done(function () {
        window.location.href = `/users/${user_id}`;
    }).fail(function (err) {
        console.log(err)
        Swal.fire(
            "Erro!",
            "Erro na tentativa de seguir o usuário",
            "error"
        )
        $('#follow').prop('disabled', false)
    })
}

function unfollow() {
    const user_id = $(this).data('user-id')
    $(this).prop('disabled', true)

    $.ajax({
        url: `/users/${user_id}/unfollow`,
        method: "POST"
    }).done(function () {
        window.location.href = `/users/${user_id}`;
    }).fail(function (err) {
        console.log(err)
        Swal.fire(
            "Erro!",
            "Erro na tentativa de deixar de seguir o usuário",
            "error"
        )
        $('#unfollow').prop('disabled', false)
    })
}

function edit(event) {
    const user_id = $(this).data('user-id')
    event.preventDefault();

    $.ajax({
        url: "/edit-user",
        method: "PUT",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
        }
    }).done(function () {
        Swal.fire(
            "Sucesso!",
            "Sua conta foi atualizada com sucesso",
            "success"
        ).then(function () {
            window.location.href = `/users/${user_id}`
        })
    }).fail(function () {
        Swal.fire(
            "Falha!",
            "Não foi possível atualizar os dados da conta",
            "error"
        )
    })
}

function updatePassword(event) {
    const user_id = $(this).data('user-id');
    event.preventDefault();

    if ($('#new-pass').val() != $('#conf-new-pass').val()) {
        Swal.fire(
            "Falha! As senhas não coincidem",
            "Verifique se a nova senha é igual a confirmada!",
            "error"
        )
        return;
    }

    $.ajax({
        url: "/updatePassword",
        method: "POST",
        data: {
            new_Password: $('#new-pass').val(),
            old_Password: $('#actual-pass').val(),
        }
    }).done(function () {
        Swal.fire(
            "Sucesso!",
            "Sua senha foi atualizada com sucesso",
            "success"
        ).then(function () {
            window.location.href = `/users/${user_id}`
        })
    }).fail(function (err) {
        console.log(err)
        Swal.fire(
            "Falha!",
            "Não foi possível atualizar a senha!",
            "error"
        )
    })
}

function deleteUser(event) {
    Swal.fire({
        title: "Atenção",
        text: "Deseja deletar sua conta? Essa é uma ação irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        confirmButtonText: "Sim, deletar!",
        icon: "warning"
    }).then(function (confirm) {
        if (confirm.isConfirmed) {
            $.ajax({
                url: "/delete-user",
                method: "DELETE",
                dataType: "text"
            }).done(function () {
                Swal.fire(
                    "Sucesso!",
                    "Sua conta foi excluida!",
                    "success"
                ).then(function () {
                    window.location.href = "/logout";
                })
            }).fail(function () {
                Swal.fire(
                    "Erro!",
                    "Erro ao excluir a conta!",
                    "error"
                )
            })
        }
    })
}