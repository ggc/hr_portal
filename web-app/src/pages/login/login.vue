<template>
    <div class="login__form-block">
        <input placeholder="Username" @keyup.enter="focusNextInput" v-model="username"/>
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
        },
        focusNextInput(event) {
            do {
                const sibling = event.target.nextElementSibling;
                if (sibling.localName === 'input') {
                    sibling.focus();
                    return;
                }
            } while(sibling);
        }
    },
    computed: {
        creds() {
            return {
                username: this.username.trim(),
                password: this.password
            }
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<!-- enter-active: Enter transition -->
<!-- leave-active: Leave transition -->
<!-- enter/leave-to: Ironicaly, the off state -->
<!-- 
    .inputError-leave-active {
        transform: translateX(200%);
        transition: all 3s ease;
    }-->
<style scoped>
    .inputError-enter-active {
        transition: all .2s cubic-bezier(0,.77,.75,1.45);
    }
    .inputError-enter {
        transform: translateY(-100%);
    }
    .inputError-leave-to {
        opacity: 0;
        transform: translateY(-100%);
    }
    .inputError-leave-active {
        transition: all .3s cubic-bezier(.75,1.45, 0,.77);
    }
    input {
        border: 0;
        border-radius: 9px;
        font-size: 1em;
        margin: .4em;
        padding: 1em;
        background: '##f3cdd2';
    }
    button {
        margin: 1em .5em;
        padding: .5em 1em;
        font-size: 1em;
        border-radius: 2em;
        background: white;
        color: #f54a61;
        font-weight: bold;
        border-width: 0;
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
        color: white;
        margin: 0 1em;
        align-self: flex-start;
        border-bottom: 2px solid;
        padding: 2px;
    }
</style>
