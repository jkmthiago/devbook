<script setup>
import { onMounted, ref } from 'vue';

function getCookie(nome) {
    let nomeCookie = nome + "=";
    let ca = document.cookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') c = c.substring(1, c.length);
        if (c.indexOf(nomeCookie) == 0) return c.substring(nomeCookie.length, c.length);
    }
    return null;
}

const UserId = getCookie("UserID");
const token = getCookie("Token");

const posts = ref([]);

onMounted(() => {    
    fetch("http://localhost:5000/posts", {
        method: "GET",
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
        }
    }).then(response => {
        if (!response.ok) {
            throw new Error(`Erro na requisição: ${response.status} - ${response.statusText}`);
        }

        return response.json();
    }).then(data => {
        console.log(data);
        posts.value = data;
    }).catch(error => {
        console.error('Erro na requisição: ', error);
        Swal.fire(
            "Erro",
            "Ocorreu um erro ao tentar buscar as postagens. Tente novamente.",
            "error"
        );
    });
})

function deletePost(postId) {
    fetch(`http://localhost:5000/posts/${postId}`, {
        method: 'DELETE',
        headers: {
            'Authorization': `Bearer ${token}`,
        }
    }).then(response => {
        if (!response.ok) {
            throw new Error('Erro ao deletar o post.');
        }
        // Remova o post deletado da lista de posts
        posts.value = posts.value.filter(post => post.id !== postId);
        Swal.fire('Sucesso', 'Post deletado com sucesso!', 'success');
    }).catch(error => {
        Swal.fire('Erro', 'Ocorreu um erro ao tentar deletar o post.', 'error');
    });
}

</script>

<!-- --- --- --- --- --- --- --- --- --- --- --- -->

<template>
    <div class="container-fluid">
        <div class="row mt-4">
            <div class="col-xs-12 col-sm-12 col-md-5 col-lg-5 col-xl-5">
                <h3> O que está pensando hoje? </h3>
                <fieldset>
                    <form id="new-post">
                        <div class="form-group">
                            <label for="title"> Titulo </label>
                            <input type="text" class="form-control" id="title" name="title" required="required"
                                placeholder="Insira o título da sua publicação">
                        </div>
                        <div class="form-group">
                            <label for="content"> Conteúdo </label>
                            <textarea class="form-control" name="content" id="content" required="required"
                                placeholder="Insira o conteúdo da sua publicação"></textarea>
                        </div>
                        <button class="btn btn-primary" type="submit">
                            Publicar
                        </button>
                    </form>
                </fieldset>
            </div>
            <div class="col-xs-12 col-sm-12 col-md-6 col-lg-6 col-xl-6">
                <div v-for="post in posts" :key="post.id" class="jumbotron">
                    <a :href="'/users/' + post.autor_id"> {{ post.autor_nick }} - {{ post.created_in }} </a>
                    <h1 class="display-4">{{ post.title }}</h1>
                    <p class="lead">{{ post.content }}</p>
                    <hr class="my-4">

                    <p class="post">
                        <i class="fa-regular fa-heart like_post post_i"
                            style="cursor: pointer; text-decoration: none; color: rgb(27, 25, 25);"></i>
                        <span>{{ post.likes }}</span>

                        <a v-if="post.autor_id === Number(UserId)" href="posts/{{ post.id }}/edit"
                            style="text-decoration: none;">
                            <i class="fa-regular fa-pen-to-square post_i2"></i>
                        </a>

                        <i v-if="post.autor_id === Number(UserId)" class="fa-solid fa-trash deletePost"
                            style="cursor: pointer;" @click="deletePost(post.id)"></i>
                    </p>
                </div>
            </div>
        </div>
    </div>
</template>

<!-- --- --- --- --- --- --- --- --- --- --- --- -->

<style>
body {
    background-color: rgb(255, 255, 255);
}

.post_i{
    padding-right: 2px;
}

.post span, .post_i2{
    padding-right: 10px;
}
</style>