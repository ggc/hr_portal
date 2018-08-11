<template>
    <div class="create-user__form-block">
        <input placeholder="First name" v-model="firstName"/>
        <input placeholder="Last name" v-model="lastName"/>
        <input placeholder="Contractor" v-model="contractor"/>
        <input placeholder="Project" v-model="project"/>
        <button @click="onSave">Save</button>
    </div>
</template>

<script>
import { UserService } from '../../services/user.service'; 

export default {
    name: 'CreateUserPage',

    data() {
        return {
            firstName: '',
            lastName: '',
            contractor: '',
            project: ''
        }
    },
    methods: {
        onSave() {
            console.log('[LOG] userservice', UserService.createUser(this.userData));
            console.log('[LOG] onSave', this.firstName, this.lastName, this.contractor, this.project);
            this.firstName = '';
            this.lastName = '';
            this.contractor = '';
            this.project = '';
        }
    },
    computed: {
        userData() {
            return {
                firstName: this.firstName,
                lastName: this.lastName,
                contractor: this.contractor,
                project: this.project,
            }
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
