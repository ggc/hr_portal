<template>
    <div class="create-user__form-block">
        <input placeholder="Username" v-model="username"/>
        <input placeholder="Password" type="password" v-model="password"/>
        <password-quality-level :password="password"></password-quality-level>
        <input placeholder="First name" v-model="firstName"/>
        <input placeholder="Last name" v-model="lastName"/>
        <input placeholder="Project" v-model="project"/>
        <button @click="onSave">Save</button>
    </div>
</template>

<script>
import { UserService } from '../../services/user.service';
import passwordQualityLevel from '../../components/password-quality-level';

export default {
    name: 'CreateUserPage',
    components: {
        passwordQualityLevel
    },
    data() {
        return {
            username: '',
            password: '',
            firstName: '',
            lastName: '',
            project: ''
        }
    },
    computed: {
        userData() {
            return {
                username: this.username,
                password: this.password,
                firstName: this.firstName,
                lastName: this.lastName,
                project: this.project
            }
        }
    },
    methods: {
        async onSave() {
            console.log('[LOG] saving user...');
            const res = await UserService.createUser(this.userData);
            console.log('[LOG] user saved!', res);
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .create-user__form-block {
        display: flex;
        flex-direction: column;
    }
</style>
