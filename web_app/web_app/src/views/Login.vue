<script setup>

import { computed, reactive } from 'vue';

const form = reactive({
    email: "",
    password: ""
});

const isEmailValid = computed(() => {
    let regex = /^\S+@\S+\.\S+$/;
    return regex.test(form.email);
});

const isPasswordValid = computed(() => {
    return form.password.length >= 6;
})

const isFormValid = computed(() => {
    return isEmailValid.value && isPasswordValid.value;
});

function handleSubmit() {
    if (!isFormValid.value) {
        if (!isEmailValid.value) {
            Swal.fire(
                "E-mail Incorreto",
                "Favor verifique se o E-mail está digitado corretamente",
                "error"
            )
        } else if (!isPasswordValid.value) {
            Swal.fire(
                "Senha Incorreta",
                "Favor verifique se a Senha contém no mínimo 6 dígitos",
                "error"
            )
        }
        return
    }

    fetch("http://localhost:5000/login", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "email": form.email,
            "password": form.password
        })
    }).then(response => {
        if (!response.ok) {
            throw new Error(`Erro na requisição: ${response.status} - ${response.statusText}`);
        }

        return response.json();
    }).then(data => {
        document.cookie = `UserID=${data.id || ""}; path=/; SameSite=Lax`;
        document.cookie = `Token=${data.token || ""}; path=/; SameSite=Lax`;

        Swal.fire(
            "Sucesso!",
            "Login efetuado com sucesso!",
            "success"
        );

        const redirectTo = new URLSearchParams(window.location.search).get('redirect') || "/home";
        window.location = redirectTo;
    }).catch(error => {
        console.error('Erro na requisição: ', error);
        Swal.fire(
            "Erro",
            "Ocorreu um erro ao tentar fazer login. Tente novamente.",
            "error"
        );
    });
}

</script>

<!-- --- --- --- --- --- --- --- --- --- --- --- -->
col-xs-12 col-sm-12 col-md-7 col-lg-7 col-xl-7
<template>
    <div class="modalizado">
        <div id="form-ui">
            <form @submit.prevent="handleSubmit()" id="login_form">
                <div id="login_form-body">
                    <div id="welcome-lines">
                        <div id="welcome-line-1">
                            DevBook
                        </div>
                        <div id="welcome-line-2">Welcome Back, User</div>
                    </div>
                    <div id="input-area">
                        <div class="form-inp-login">
                            <input placeholder="Digite seu endereço de Email" type="email" name="email" id="email"
                                required="required" v-model="form.email">
                        </div>
                        <div class="form-inp-login">
                            <input placeholder="Digite sua Senha" type="password" name="password" id="password"
                                required="required" v-model="form.password">
                        </div>
                    </div>
                    <div id="submit-button-cvr">
                        <button id="submit-button" type="submit">Login</button>
                    </div>
                    <div id="forgot-pass">
                        <a href="/register">Deseja se Cadastrar?</a>
                    </div>
                </div>
            </form>
        </div>
    </div>
</template>

<!-- --- --- --- --- --- --- --- --- --- --- --- -->

<style scoped>
.modalizado {
    width: 100vw;
    height: 100vh;
    position: fixed;
    top: 0;
    left: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 999;
    background-color: rgba(20, 20, 20, 0.966);
}
</style>