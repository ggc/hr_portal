<template>
    <div class="login__form-block">
        <input placeholder="Username" @keyup.enter="onLogin" v-model="username"/>
        <transition name="inputError">
            <span class="login__input-error" v-if="notifyErrors && !isValidUsername()"> Invalid username </span>
        </transition>
        <input type="password" placeholder="Password" @keyup.enter="onLogin" v-model="password"/>
        <transition name="inputError">
            <span class="login__input-error" v-if="notifyErrors && !isValidPassword()"> Invalid password </span>
        </transition>
        <div class="login__form__buttons">
            <button>Forget password</button>
            <button @click="onLogin">Login</button>
        </div>
    </div>
</template>

<script>
import { LoginService } from '../../services/login.service';
export default {
    name: 'LoginPage',

    data() {
        return {
            username: '',
            password: '',
            notifyErrors: false
        }
    },
    methods: {
        async onLogin() {
            this.notifyErrors = true;
            if (!this.isValidUsername() || !this.isValidPassword()) {
                console.log('[LOG] Cant move forward, sorry');
                return;
            };
            try {
                const token = await LoginService.login(this.creds);
                this.$router.push('create-user')
                console.log('[LOG] token', token);
            } catch(err) {
                console.log('[LOG] login error', err);
            }
        },
        isValidUsername() {
            // TODO: Subs by regex
            return this.username.trim().indexOf(' ') === -1 && this.username;
        },
        isValidPassword() {
            // TODO: Subs by regex
            return this.password.trim().indexOf(' ') === -1 && this.password;
        }
    },
    computed: {
        creds() {
            return {
                username: this.username.trim(),
                password: this.password
            }
        }
    },
    watch: {
        creds() {
            this.notifyErrors = false;
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<!-- enter-active: Enter transition -->
<!-- leave-active: Leave transition -->
<!-- enter/leave-to: Ironicaly, the off state -->
<style scoped>
    .inputError-enter-active {
        transition: all .3s cubic-bezier(.37,-0.06,.49,1.43);
    }    
    .inputError-leave-active {
        transition: all .3s ease-in;
    }
    .inputError-enter, .inputError-leave-to {
        opacity: 0;
        transform: translateX(-200%);
    }

    input {
        margin: .4em;
    }
    button {
        margin: .4em;
    }
    .login__form-block {
        display: flex;
        flex-direction: column;
    }
    .login__form__buttons {
        display: flex;
        justify-content: flex-end;
    }
    .login__input-error {
        color: #ff8080;
        margin: 0 1em;
        align-self: flex-start;
    }
</style>
