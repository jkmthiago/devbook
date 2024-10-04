<script setup>

function deleteCookie(name) {
    document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/; SameSite=Lax`;
}

function logout() {
    deleteCookie("UserId")
    deleteCookie("Token")
}
</script>

<!-- --- --- --- --- --- --- --- --- --- --- --- -->

<template>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbar-principal"
            aria-controls="navbar-principal" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbar-principal">
            <a class="navbar-brand" href="#">DevBook</a>
            <ul class="navbar-nav mr-auto mt-2 mt-lg-0">
                <li class="nav-item active">
                    <a class="nav-link" href="/home">Home <span class="sr-only">(current)</span></a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="/users/{{.}}">Meu Perfil</a>
                </li>
                <li class="nav-item">
                    <form class="form-inline my-2 my-lg-0" action="/search-users" method="GET">
                        <input class="form-control mr-sm-2" type="search" placeholder="Pesquisar Usuários"
                            aria-label="Search_Users" name="search_users" id="search_users" required="required">
                    </form>
                </li>
            </ul>

            <span class="navbar-text p-0">
                <a href="\" @click="logout()" style="text-decoration: none;">Sair
                    <i class="fa-solid fa-power-off fa-fade"></i>
                </a>
            </span>
        </div>
    </nav>
</template>

<!-- --- --- --- --- --- --- --- --- --- --- --- -->

<style>
</style>