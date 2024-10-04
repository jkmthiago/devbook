<script setup>
import { computed, reactive } from 'vue';

const form = reactive({
    name: "",
    nick: "",
    email: "",
    password: "",
    confirmPassword: "",
})

const isEmailValid = computed(() => {
    let regex = /^\S+@\S+\.\S+$/;
    return regex.test(form.email);
});

const isPasswordValid = computed(() => {
    return form.password.length >= 6;
});

const isPasswordsSame = computed(() => {
    return form.password === form.confirmPassword;
});

const isFormValid = computed(() => {
    return isEmailValid.value && isPasswordValid.value && isPasswordsSame.value
})

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
        } else if (!isPasswordsSame.value) {
            Swal.fire(
                "Senha Diferentes",
                "Favor verifique se as Senha são iguais",
                "error"
            )
        }
        return
    }

    fetch("http://localhost:5000/users", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "name": form.name,
            "nick": form.nick,
            "email": form.email,
            "password": form.password
        })
    }).then(response => {
        if (!response.ok) {
            throw new Error(`Erro na requisição: ${response.status} - ${response.statusText}`);
        }

        return response.json();
    }).then(data => {
        console.log("Resposta da API:");
        console.log(data);

        Swal.fire(
            "Sucesso!",
            "Cadastro efetuado com sucesso!",
            "success"
        );

        window.location = "/login"
    }).catch(error => {
        console.error('Erro na requisição: ', error);
        Swal.fire(
            "Erro",
            "Ocorreu um erro ao tentar fazer o Cadastro. Tente novamente.",
            "error"
        );
    });
}

</script>

<!-- --- --- --- --- --- --- --- --- --- --- --- -->

<template>
    <div class="modalizado">
    <div id="form-ui">
        <form @submit.prevent="handleSubmit" id="register_form">
            <div id="register_form-body">
                <div id="welcome-lines">
                    <div id="welcome-line-1">
                        DevBook
                    </div>
                    <div id="welcome-line-2">Crie sua conta agora mesmo</div>
                </div>
                <div id="input-area">
                    <div class="form-inp-register">
                        <input placeholder="Digite seu Nome" type="text" name="name" id="name" required="required"
                            v-model="form.name">
                    </div>
                    <div class="form-inp-register">
                        <input placeholder="Digite seu Nome de Usuário" type="text" name="nick" id="nick"
                            required="required" v-model="form.nick">
                    </div>
                    <div class="form-inp-register">
                        <input placeholder="Digite seu endereço de Email" type="email" name="email" id="email"
                            required="required" v-model="form.email">
                    </div>
                    <div class="form-inp-register">
                        <input placeholder="Digite sua Senha" type="password" name="password" id="password"
                            required="required" v-model="form.password">
                    </div>
                    <div class="form-inp-register">
                        <input placeholder="Confirme sua Senha" type="password" name="confirme_password"
                            id="confirme_password" required="required" v-model="form.confirmPassword">
                    </div>
                </div>
                <div id="submit-button-cvr">
                    <button id="submit-button" type="submit">Registrar</button>
                </div>
                <div id="forgot-pass">
                    <a href="/login">Já possúi uma conta?</a>
                </div>
                <div id="bar-register"></div>
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