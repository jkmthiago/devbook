$('#new-post').on('submit', createPost)
$(document).on('click', '.like_post', like)
$(document).on('click', '.unlike_post', unlike)
$('#updatePost').on('click', updatePost)
$('.deletePost').on('click', deletePost)

function createPost(event) {
    event.preventDefault();

    $.ajax({
        url: "/posts",
        method: "POST",
        data: {
            title: $('#title').val(),
            content: $('#content').val(),
        },
        dataType: "text"
    }).done(function () {
        window.location = "/home";
    }).fail(function () {
        Swal.fire(
            "Erro",
            "Erro ao realizar a Publicação",
            "error"
        );
    })
}

function like(event) {
    event.preventDefault();

    const clickedElement = $(event.target)
    const post_id = clickedElement.closest('div').data("post-id")

    clickedElement.prop('disable', true);

    $.ajax({
        url: `/posts/${post_id}/like`,
        method: "POST",
        dataType: "text"
    }).done(function () {
        const likeCounter = clickedElement.next('span');
        const likeCount = parseInt(likeCounter.text());

        likeCounter.text(likeCount + 1);

        clickedElement.addClass('unlike_post');
        clickedElement.addClass('text-danger');
        clickedElement.removeClass('like_post');

    }).fail(function (err) {
        console.log(err)
        Swal.fire(
            "Erro",
            "Erro ao Curtir a Publicação",
            "error"
        );
    }).always(function () {
        clickedElement.prop('disable', false);
    })
}

function unlike(event) {
    event.preventDefault();

    const clickedElement = $(event.target)
    const post_id = clickedElement.closest('div').data("post-id")

    clickedElement.prop('disable', true);

    $.ajax({
        url: `/posts/${post_id}/unlike`,
        method: "POST",
        dataType: "text"
    }).done(function () {
        const likeCounter = clickedElement.next('span');
        const likeCount = parseInt(likeCounter.text());

        likeCounter.text(likeCount - 1);

        clickedElement.addClass('like_post');
        clickedElement.removeClass('text-danger');
        clickedElement.removeClass('unlike_post');

    }).fail(function (err) {
        console.log(err)
        Swal.fire(
            "Erro",
            "Erro ao Descurtir a Publicação",
            "error"
        );
    }).always(function () {
        clickedElement.prop('disable', false);
    })
}

function updatePost() {
    $(this).prop('disable', true);

    const post_id = $(this).data('post-id');

    $.ajax({
        url: `/posts/${post_id}`,
        method: "PUT",
        data: {
            title: $('#title').val(),
            content: $('#content').val(),
        },
        dataType: "text"
    }).done(function () {
        Swal.fire(
            'Sucesso',
            'Postagem Atualizada',
            'success'
        ).then(function () {
            window.location.href = "/home"
        })
    }).fail(function (erro) {
        Swal.fire(
            "Erro",
            "Erro ao Editar a Publicação",
            "error"
        );
        console.log(erro)
    }).always(function () {
        $('#updatePost').prop('disable', true)
    })
}

function deletePost(event) {
    event.preventDefault();

    Swal.fire({
        title: "Você tem Certeza?",
        text: "Em caso de exclusão da publicação você não poderá recuperá-la",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function (confirmation) {
        if (!confirmation.value) return;

        const clickedElement = $(event.target)
        const post = clickedElement.closest('div')
        const post_id = post.data("post-id")

        clickedElement.prop('disable', true);

        $.ajax({
            url: `/posts/${post_id}`,
            method: "DELETE",
            dataType: "text"
        }).done(function () {
            post.fadeOut("slow", function () {
                $(this).remove();
            });
        }).fail(function (erro) {
            Swal.fire(
                "Erro",
                "Erro ao Excluir a Publicação",
                "error"
            );
        })
    })
}