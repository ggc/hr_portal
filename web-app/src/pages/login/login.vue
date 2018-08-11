<template>
    <div class="login__form-block">
        <input placeholder="Username" v-model="username"/>
        <input placeholder="Password" v-model="password"/>
        <div class="login__form__buttons">
            <button>Forget password</button>
            <button @click="onLogin">Login</button>
        </div>
    </div>
</template>

<script>
export default {
    name: 'LoginPage',

    data() {
        return {
            username: '',
            password: ''
        }
    },
    methods: {
        onLogin() {
            console.log('[LOG] onLogin', this.username, this.password);
            this.$router.push('create-user');
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .login__form-block {
        display: flex;
    }
    .input {
        width: 10em;
    }
</style>
